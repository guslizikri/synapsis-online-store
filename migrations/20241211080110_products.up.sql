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