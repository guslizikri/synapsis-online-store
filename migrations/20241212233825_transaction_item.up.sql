CREATE TABLE transaction_item (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id),
    product_id INT REFERENCES products(id),
    quantity INT not null,
    product_price int not null,
    created_at timestamp NOT NULL DEFAULT now()
);