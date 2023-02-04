SELECT 
    user_id, 
    fullname, 
    email,
    password,
    created_at
FROM toko.users
WHERE email = $1;
