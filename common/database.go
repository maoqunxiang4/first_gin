package common

import (
	"fmt"
	"github.com/spf13/viper"
	"goPro/config"
	"goPro/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB = InitDB()

func InitDB() *gorm.DB {
	config.InitConfig()

	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	dbName := viper.GetString("datasource.dbName")
	charset := viper.GetString("datasource.charset")
	parseTime := viper.GetString("datasource.parseTime")
	loc := viper.GetString("datasource.loc")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		username, password, host, dbName, charset, parseTime, loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("there are some problems in the system now : \n", err)
	}

	db.Table("t_user").AutoMigrate(&model.User{})

	return db
}

func GetDB(tableName string) *gorm.DB {
	return DB.Table(tableName)
}
