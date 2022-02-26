module alert

go 1.17

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.1
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
)
