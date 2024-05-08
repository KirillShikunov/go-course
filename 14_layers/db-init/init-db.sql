CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    status INTEGER NOT NULL,
    total_price INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);