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