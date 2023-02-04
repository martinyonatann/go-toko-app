SELECT 
    item_id, 
    item_name, 
    description,
    price,
    created_at
FROM toko.items
WHERE item_id = $1;
