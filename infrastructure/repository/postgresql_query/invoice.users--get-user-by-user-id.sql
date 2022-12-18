SELECT 
    user_id, 
    fullname, 
    email, 
    created_at
FROM invoice.users
WHERE user_id = $1;
