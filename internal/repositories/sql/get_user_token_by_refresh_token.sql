select
	ut.*
from
	public.user_token ut
inner join public.users u 
	on
	u.id = ut.user_id
where
	ut.refresh_token = $1
	and ut.revoked = false
	and ut.token_expires_at > current_timestamp
	and u.is_active = true
	and (
		u.locked_until is null or
		u.locked_until < current_timestamp
	);