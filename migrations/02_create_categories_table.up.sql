CREATE TABLE IF NOT EXISTS categories (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255),
    "notes" VARCHAR(255),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
    )