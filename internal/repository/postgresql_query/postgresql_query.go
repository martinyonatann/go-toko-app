package postgresql_query

import _ "embed"

var (
	//go:embed invoice.users--insert.sql
	UserInsert string

	//go:embed invoice.users--get-user-by-user-id.sql
	GetUserById string
)
