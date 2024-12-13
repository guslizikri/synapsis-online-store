CREATE TABLE public.cart_item (
	id serial4 PRIMARY KEY NOT NULL,
	product_id int NOT NULL,
	cart_id int not null,
	quantity int not null,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT itemcart_product_fk FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE on update cascade,
	CONSTRAINT itemcart_cart_fk FOREIGN KEY (cart_id) REFERENCES cart(id) ON DELETE CASCADE on update CASCADE
);