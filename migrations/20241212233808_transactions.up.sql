CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_public_id varchar(50) REFERENCES users(public_id),
    total_price int not null,
    status int4 not null,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now()
);