package pkg

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// MySQL struct
type mySQL struct {
	DB *gorm.DB
}

// MySQL mySQL struct
var MySQL *mySQL

// init resource MySQL struct
func init() {
	db, err := gorm.Open(viper.GetString("DB_DRIVER"), viper.GetString("DB_CONNECTION"))
	if err != nil {
		panic("MySQL open connection failed.")
	}

	db.DB().SetMaxIdleConns(viper.GetInt("MYSQL_MAX_IDLE"))
	db.DB().SetMaxOpenConns(viper.GetInt("MYSQL_MAX_OPEN_CONNS"))
	db.DB().SetConnMaxLifetime(time.Hour)
	MySQL = &mySQL{
		DB: db,
	}
}
