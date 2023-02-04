SELECT 
    user_id, 
    fullname, 
    email, 
    created_at
FROM toko.users
WHERE user_id = $1;
