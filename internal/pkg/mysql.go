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
package pkg

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type mySQL struct {
	DB     *gorm.DB
	IsOpen bool
}

var mySQLdb *mySQL

func NewMySQL() *mySQL {

	if mySQLdb != nil && mySQLdb.IsOpen {
		return mySQLdb
	}

	database, err := gorm.Open(viper.GetString("DB_DRIVER"), viper.GetString("DB_CONNECTION"))
	if err != nil {
		panic(fmt.Sprintf("MySQL open connection %s error %s.", viper.GetString("DB_CONNECTION"), err))
	}

	database.Set("gorm:table_options", "ENGINE=InnoDB")
	database.DB().SetMaxIdleConns(viper.GetInt("MYSQL_MAX_IDLE"))
	database.DB().SetMaxOpenConns(viper.GetInt("MYSQL_MAX_OPEN_CONNS"))
	database.DB().SetConnMaxLifetime(time.Hour)

	mySQLdb = &mySQL{
		DB:     database,
		IsOpen: true,
	}
	return mySQLdb
}
