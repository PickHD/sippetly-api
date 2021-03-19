package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var DB *sql.DB

func Init(){
	db, _ := sql.Open("mysql",os.Getenv("DB_DSN_URI"))
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
			"file://db/migrations",
			"mysql", 
			driver,
	)
	defer db.Close()

	m.Steps(2)

	log.Println("Databases Successfully Migrated!")
   
}

func OpenDB()error{
	var err error

	DB,err =sql.Open("mysql",os.Getenv("DB_DSN_URI"))

	if err != nil {
    	return err
  	}

	DB.SetMaxOpenConns(10)
  	DB.SetMaxIdleConns(10)
  	DB.SetConnMaxLifetime(time.Minute * 5)

	return nil
}

func CloseDB() error{
	return DB.Close()
}
