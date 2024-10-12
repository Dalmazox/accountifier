CREATE TABLE users (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	email varchar(255) NOT NULL,
	username varchar(50) NOT NULL,
	password_hash text NOT NULL,
	created_at timestamptz DEFAULT current_timestamp NOT NULL,
	updated_at timestamptz DEFAULT current_timestamp NOT NULL,
	last_login timestamptz NULL,
	is_active bool DEFAULT true NOT NULL,
	failed_attempts int4 NULL,
	locked_until timestamptz NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_username_key UNIQUE (username)
);
CREATE INDEX idx_users_email ON public.users USING btree (email);
CREATE INDEX idx_users_is_active ON public.users USING btree (is_active);
CREATE INDEX idx_users_username ON public.users USING btree (username);