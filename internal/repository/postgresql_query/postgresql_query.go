package postgresql_query

import _ "embed"

var (
	//go:embed toko.users--insert.sql
	UserInsert string
	//go:embed toko.users--get-user-by-email.sql
	GetUserByEmail string
	//go:embed toko.users--get-user-by-user-id.sql
	GetUserById string
	//go:embed toko.users--list-users.sql
	ListUsers string

	//go:embed toko.items--list-items.sql
	ListProduct string
	//go:embed toko.products-get-product-by-id.sql
	GetProductByID string
)
