CREATE TABLE IF NOT EXISTS supplier_categories (
    "supplier_id" uuid,
    "category_id" uuid,
    PRIMARY KEY (supplier_id, category_id),
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
)

