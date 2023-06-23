select
	p.product_id,
	p.name,
	p.description,
	p.price,
	p.category_id
from
	toko.products p
where
	p.product_id = $1;
