CREATE TABLE IF NOT EXISTS suppliers (
    "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name" VARCHAR(255),
    "email" VARCHAR(255) NOT NULL,
    "notes" VARCHAR(255),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW()),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
)