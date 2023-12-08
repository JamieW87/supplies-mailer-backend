CREATE TABLE IF NOT EXISTS categories (
    "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name" VARCHAR(255),
    "notes" VARCHAR(255),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
    )