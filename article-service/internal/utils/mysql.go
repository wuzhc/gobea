package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	database := viper.GetString("database.database")
	dsn := fmt.Sprintf("%s:%s@tcp/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, database)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		panic("open mysql error, " + err.Error())
	}
}

func DB() *gorm.DB {
	return db
}
