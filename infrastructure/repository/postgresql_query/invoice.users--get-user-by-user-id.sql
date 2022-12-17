SELECT 
    user_id, 
    fullname, 
    email, 
    password, 
    created_at, 
    updated_at
FROM invoice.users
WHERE user_id = $1;
