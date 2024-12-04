module github.com/Rohanrevanth/e-store-go/database

go 1.23.1

require github.com/Rohanrevanth/e-store-go/models v0.0.0-00010101000000-000000000000

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)

require (
	github.com/go-redis/redis/v8 v8.11.5
	gorm.io/driver/sqlite v1.5.6
	gorm.io/gorm v1.25.12
)

replace github.com/Rohanrevanth/e-store-go/models => ../models
