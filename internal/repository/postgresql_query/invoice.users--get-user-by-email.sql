SELECT 
    user_id, 
    fullname, 
    email,
    password,
    created_at
FROM invoice.users
WHERE email = $1;
