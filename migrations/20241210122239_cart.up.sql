CREATE TABLE public.cart (
	id serial4 PRIMARY KEY NOT NULL,
	user_public_id varchar(50) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	CONSTRAINT user_fk FOREIGN KEY (user_public_id) REFERENCES users(public_id) ON DELETE CASCADE on update CASCADE
);
