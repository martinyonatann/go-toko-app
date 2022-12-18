//go:build dev
// +build dev

package config

const (
	DB_USER                = "postgres"
	DB_PASSWORD            = "password"
	DB_DATABASE            = "invoice"
	DB_HOST                = "127.0.0.1"
	API_PORT               = 5432
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
)
