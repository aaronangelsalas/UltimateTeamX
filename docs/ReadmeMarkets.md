Market DB Notes

This doc records schema decisions for market-svc.

Scope
- market-svc owns listings/bids state and audit only.
- economic ownership stays in club-svc (holds/settlement).

Concurrency
- DB is not used for locking; Redis handles distributed locks.
- All listing state transitions must happen under the Redis lock.

Schema
- listings: current state of each listing (never deleted).
- bids: immutable history of bids for audit/recovery.
- best_bid/best_bidder_club_id are stored on listings for fast reads and are
  updated only under lock.
- bids.hold_id is used to release the previous credit hold on outbid.

Rules
- Allowed listing status values: ACTIVE, SOLD, EXPIRED.
- No foreign keys to other services' databases.
- No deletes for listings/bids; only updates and inserts.

CreateListing Flow (market-svc)
- Validate request fields (ids, prices, expiry).
- Check for an existing ACTIVE listing for the same card in market DB.
- Call club-svc LockCard; lock is the ownership/availability check.
- Insert listing with status ACTIVE in market DB.
- On DB failure, release the card lock in club-svc.

PlaceBid Flow (market-svc)
- Acquisisce un lock Redis su `lock:listing:{listing_id}`.
- Verifica listing (ACTIVE, non scaduta, importo valido).
- Crea hold crediti nel club-svc per il bidder.
- Inserisce il bid e aggiorna best_bid in transazione DB.
- Rilascia l'hold precedente (se presente).
- Rilascia il lock Redis.

Osservabilita
- Log strutturati nel server per errori e successi del flusso CreateListing.

Configurazione
- Le variabili vengono caricate da `.env` (sovrascrivono quelle gia' presenti) usando `GO_DOTENV_PATH` se valorizzato.

Esempi pratici (senza codice complesso)
- Creare un listing = "mettere una carta in vendita".
  - Risultato: il marketplace crea un annuncio con prezzo di partenza e scadenza.
  - Comando (da terminale):
    grpcurl -plaintext -d '{
      "seller_user_id": "11111111-1111-1111-1111-111111111111",
      "user_card_id": "22222222-2222-2222-2222-222222222222",
      "start_price": 1000,
      "buy_now_price": 2000,
      "expires_at_unix": 1893456000
    }' localhost:50053 market.v1.MarketService/CreateListing

- Fare un'offerta (bid) = "rilanciare su un annuncio".
  - Risultato: viene registrata l'offerta piu' alta e bloccati i crediti del bidder.
  - Comando (da terminale):
    grpcurl -plaintext -d '{
      "listing_id": "<LISTING_ID>",
      "bidder_user_id": "33333333-3333-3333-3333-333333333333",
      "bid_amount": 1500
    }' localhost:50053 market.v1.MarketService/PlaceBid
