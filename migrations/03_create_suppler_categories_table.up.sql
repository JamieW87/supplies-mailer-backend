CREATE TABLE IF NOT EXISTS supplier_categories (
    "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "supplier_id" uuid,
    "category_id" uuid,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
)