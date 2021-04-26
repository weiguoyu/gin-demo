// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package db

import (
	"fmt"
	"gin-demo/pkg/config"
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
	"time"
)

var db *Database

func GetMysqlInstance() *Database {
	err := db.Conn.Ping()
	if err != nil {
		logrus.WithField("global", "database").Fatalf("failed to connect mysql")
		panic(err)
	}
	return db
}

func OpenDatabase(config *config.Config) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Common.Mysql.Username, config.Common.Mysql.Password,
		config.Common.Mysql.Address, config.Common.Mysql.Port, config.Common.Mysql.Dbname)

	// https://github.com/go-sql-driver/mysql/issues/9
	conn, err := dbr.Open("mysql", url+"?parseTime=1&multiStatements=1&charset=utf8mb4&collation=utf8mb4_unicode_ci", nil)
	if err != nil {
		panic(err)
	}
	conn.SetMaxIdleConns(100)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Duration(10) * time.Second)

	db = &Database{
		Conn: conn,
	}
}
