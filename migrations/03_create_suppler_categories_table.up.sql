CREATE TABLE IF NOT EXISTS supplier_categories (
    "id" uuid PRIMARY KEY DEFAULT (uuid()),
    "supplier_id" uuid VARCHAR(255),
    "category_id" uuid VARCHAR(255),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
)