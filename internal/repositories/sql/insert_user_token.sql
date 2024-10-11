INSERT INTO user_token (user_id, token, refresh_token, token_expires_at, refresh_token_expires_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id) 
DO UPDATE SET
    token = EXCLUDED.token,
    refresh_token = EXCLUDED.refresh_token,
    token_expires_at = EXCLUDED.token_expires_at,
    refresh_token_expires_at = EXCLUDED.refresh_token_expires_at,
    updated_at = now();
	