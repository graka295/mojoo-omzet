package helper

import (
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB contract for DBImpl
type DB interface {
	Connection() *gorm.DB
	Migrate()
}

// DBImpl struct for conneciton db
type DBImpl struct {
	DB *gorm.DB
}

// NewDB inital connection db
func NewDB(config Config) DB {
	var db *gorm.DB
	dsn := ""
	if config.GetConfig().Database.Dialect == "mysql" {
		dsn = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.GetConfig().Database.Username, config.GetConfig().Database.Password, config.GetConfig().Database.Host, config.GetConfig().Database.Port, config.GetConfig().Database.Dbname)
		dbMysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db = dbMysql
	} else if config.GetConfig().Database.Dialect == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.GetConfig().Database.Host, config.GetConfig().Database.Port, config.GetConfig().Database.Username, config.GetConfig().Database.Dbname, config.GetConfig().Database.Password)
		dbPostgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db = dbPostgres
	} else {
		panic("database not supported")
	}
	if config.GetConfig().GetConfig().Mode == "develop" {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	dbStruct := DBImpl{
		DB: db,
	}
	return &dbStruct
}

// Connection get connection
func (x *DBImpl) Connection() *gorm.DB {
	return x.DB
}

// Migrate migrate migration/datbase.sql
func (x *DBImpl) Migrate() {
	query, err := ioutil.ReadFile("migration/database.sql")
	if err != nil {
		panic(err)
	}
	if err := x.DB.Exec(string(query)).Error; err != nil {
		panic(err)
	}
}
