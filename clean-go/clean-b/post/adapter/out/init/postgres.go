package init

// for setup rdbms conenction

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=admin dbname=post-msa port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return DB, err
}
