package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn    *gorm.DB
	SQLDBConn *sql.DB
)

type Hero struct {
	// 	gorm.Model // enable this to get updated_at, created_at, deleted_at
	Id         int
	Name       string
	SecretName string
	Age        int
}

func create_db_and_tables(dsn string, dbName string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		fmt.Println(err)
	}
}

func establish_gorm_connection(dsn string, dbName string) {
	dsn = dsn + dbName + "?parseTime=true"
	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	SQLDBConn, _ = DBConn.DB()
	SQLDBConn.SetMaxIdleConns(1000)
	SQLDBConn.SetMaxOpenConns(1000)
	SQLDBConn.SetConnMaxLifetime(time.Minute)

}

func create_heroes() {
	DBConn.AutoMigrate(&Hero{})
	DBConn.Create(&Hero{Name: "Deadpond", SecretName: "Dive Wilson"})
	DBConn.Create(&Hero{Name: "Spider-Boy", SecretName: "Pedro Parqueador"})
	DBConn.Create(&Hero{Name: "Rusty-Man", SecretName: "Tommy Sharp", Age: 48})
}

func select_hero_by_id(id int) (hero Hero) {
	DBConn.Find(&hero, id)
	return
}
