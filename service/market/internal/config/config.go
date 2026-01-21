package config

import "os"

// Config contiene le impostazioni runtime per market-svc.
type Config struct {
	GRPCAddr      string
	DBDSN         string
	ClubGRPCAddr  string
}

// Load legge le variabili d'ambiente con default minimi.
func Load() Config {
	return Config{
		GRPCAddr:     getEnv("GRPC_ADDR", ":50053"),
		DBDSN:        os.Getenv("DB_DSN"),
		ClubGRPCAddr: os.Getenv("CLUB_GRPC_ADDR"),
	}
}

// getEnv ritorna il fallback quando la variabile non è presente.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
