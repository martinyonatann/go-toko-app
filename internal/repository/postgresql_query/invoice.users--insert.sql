INSERT INTO invoice.users (
    "fullname",
    "email",
    "password",
    "created_at",
    "updated_at"
) VALUES (
    $1, $2, $3, now(), now())
returning user_id, fullname, email, password, created_at;
