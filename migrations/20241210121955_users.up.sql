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