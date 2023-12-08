CREATE TABLE IF NOT EXISTS categories (
    "id" uuid PRIMARY KEY DEFAULT (uuid()),
    "name" string VARCHAR(255),
    "notes" string VARCHAR(255),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
    )