package pkg

/*
Copyright © 2020 pochonlee@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type mySQL struct {
	DB *gorm.DB
}

var (
	mySQLdb   *mySQL
	onceMySQL sync.Once
)

// NewMySQL init mySQL
func NewMySQL() *mySQL {

	onceMySQL.Do(func() {
		driver, conn := viper.GetString("DB_DRIVER"), viper.GetString("DB_CONNECTION")
		if driver == "" || conn == "" {
			NewZapLogger().DPanic("DB_DIRVER/DB_CONNECTION must not empty.",
				zap.String("DB_DRIVER", driver),
				zap.String("DB_CONNECTION", driver))
		}

		database, err := gorm.Open(driver, conn)
		if err != nil {
			NewZapLogger().DPanic(err.Error(),
				zap.String("conn", conn))
		}

		database.LogMode(viper.GetBool("MYSQL_LOG_MODE"))
		database.Set("gorm:table_options", "ENGINE=InnoDB")
		database.DB().SetMaxIdleConns(viper.GetInt("MYSQL_MAX_IDLE"))
		database.DB().SetMaxOpenConns(viper.GetInt("MYSQL_MAX_OPEN_CONNS"))
		database.DB().SetConnMaxLifetime(time.Hour)

		mySQLdb = &mySQL{
			DB: database,
		}
	})
	return mySQLdb
}
