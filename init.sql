CREATE TABLE public.users (
	id serial4 NOT NULL,
	email varchar(100) NOT NULL,
	"password" varchar(150) NOT NULL,
	"role" varchar(20) NOT NULL DEFAULT 'user'::character varying,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL DEFAULT now(),
	public_id varchar(50) unique NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE public.categories (
	id serial4 PRIMARY KEY NOT NULL,
	categorie varchar(50) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now()
);

INSERT INTO categories (categorie) VALUES
('Electronics'),
('Fashion'),
('Home Appliances'),
('Books'),
('Toys');

CREATE TABLE public.products (
	id serial4 NOT NULL,
	sku varchar(100) NOT NULL,
	"name" varchar(100) NOT NULL,
	price int4 NOT NULL DEFAULT 0,
	stock int4 NOT NULL DEFAULT 0,
    id_categorie int not NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT products_pkey PRIMARY KEY (id),
 	CONSTRAINT categories_fk FOREIGN KEY (id_categorie) REFERENCES categories(id) ON DELETE CASCADE on update CASCADE
);

CREATE TABLE public.cart (
	id serial4 PRIMARY KEY NOT NULL,
	user_public_id varchar(50) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT user_fk FOREIGN KEY (user_public_id) REFERENCES users(public_id) ON DELETE CASCADE on update CASCADE
);

CREATE TABLE public.cart_item (
	id serial4 PRIMARY KEY NOT NULL,
	product_id int NOT NULL,
	cart_id int not null,
	quantity int not null,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT itemcart_product_fk FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE on update cascade,
	CONSTRAINT itemcart_cart_fk FOREIGN KEY (cart_id) REFERENCES cart(id) ON DELETE CASCADE on update CASCADE
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_public_id varchar(50) REFERENCES users(public_id),
    total_price int not null,
    status int4 not null,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE transaction_item (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id),
    product_id INT REFERENCES products(id),
    quantity INT not null,
    product_price int not null,
    created_at timestamp NOT NULL DEFAULT now()
);