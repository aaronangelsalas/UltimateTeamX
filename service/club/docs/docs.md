
# Project Tooling & Dependencies

This document tracks the main Go libraries and local tooling used in this project.

---

## Go Version

- **Go**: `1.22.12`

Managed via **mise**.

---

## Go Libraries

### Database

| Library | Version | Purpose |
|------|--------|--------|
| `github.com/jackc/pgx/v5` | v5.x | PostgreSQL driver & connection pool |
| `github.com/golang-migrate/migrate/v4` | v4.x | SQL migrations |
| `github.com/georgysavva/scany/pgxscan` | v2.x | Struct mapping from pgx rows |

### Standard Utilities

| Library | Version | Purpose |
|------|--------|--------|
| `github.com/google/uuid` | v1.x | UUID generation |
| `github.com/stretchr/testify` | v1.x | Testing helpers |

> Versions are pinned in `go.mod`.

---

## Tooling (mise)

Local development tools are managed using **mise**.

### Required Tools

| Tool | Version | Usage |
|----|--------|------|
| `go` | 1.22.12 | Application runtime |
| `postgres` | 15.x | Local database |
| `migrate` | 4.x | Run DB migrations |

### Install Tools

```bash
mise install
