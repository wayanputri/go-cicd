package database

import (
	"be17/main/app/config"
	project "be17/main/feature/project/data"
	task "be17/main/feature/task/data"
	user "be17/main/feature/user/data"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB{
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME,cfg.DB_PASSWORD,cfg.DB_HOSTNAME,cfg.DB_PORT,cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString),&gorm.Config{})
	if err != nil{
		panic(err)
	}
	return db
}

func InitialMigration(db *gorm.DB){

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
}