select 
	id,
	email,
	username,
	password_hash,
	created_at,
	updated_at,
	last_login,
	is_active,
	failed_attempts,
	locked_until
from
	public.users
where email = $1 
and is_active = true;