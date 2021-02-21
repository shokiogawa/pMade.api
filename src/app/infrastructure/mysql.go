package infrastructure

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type Mysql struct {
	DB *sqlx.DB
}

func NweMysql() *Mysql {
	mysql := new(Mysql)
	db, err := tryConnect()
	if err != nil {
		fmt.Println("database作成")
	}
	fmt.Println("database作成")
	mysql.DB = db
	return mysql

}

func tryConnect() (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST_NAME"), os.Getenv("MYSQL_DATABASE"))
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
