CREATE TABLE IF NOT EXISTS user_categories (
    "user_id" uuid,
    "category_id" uuid,
    PRIMARY KEY (user_id, category_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
)