CREATE TABLE user_token (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	user_id uuid NOT NULL,
	"token" text NOT NULL,
	refresh_token text NOT NULL,
	created_at timestamp DEFAULT current_timestamp NOT NULL,
	token_expires_at timestamp NOT NULL,
	refresh_token_expires_at timestamp NOT NULL,
	revoked bool DEFAULT false NOT NULL,
	updated_at timestamp DEFAULT current_timestamp NOT NULL,
	CONSTRAINT user_tokens_pkey PRIMARY KEY (id),
	CONSTRAINT user_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX idx_user_tokens_refresh_expires_at ON public.user_token USING btree (refresh_token_expires_at);
CREATE INDEX idx_user_tokens_refresh_token ON public.user_token USING btree (refresh_token);
CREATE INDEX idx_user_tokens_revoked ON public.user_token USING btree (revoked);
CREATE INDEX idx_user_tokens_token ON public.user_token USING btree (token);
CREATE INDEX idx_user_tokens_token_expires_at ON public.user_token USING btree (token_expires_at);
CREATE UNIQUE INDEX idx_user_tokens_user_id ON public.user_token USING btree (user_id);