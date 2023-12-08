CREATE TABLE IF NOT EXISTS suppliers (
    "id" uuid PRIMARY KEY DEFAULT (UUID()),
    "name" string VARCHAR(255),
    "email" string VARCHAR(255),
    "notes" string VARCHAR(255),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (UUID())
)