package repository

var catalog = map[int]struct {
	schema string
	table  string
}{
	1: {schema: "users", table: "login_info"},
	2: {schema: "users", table: "wallet"},
}
