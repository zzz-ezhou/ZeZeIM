package test

import (
	"ZeZeIM/ZeZeIM_Server/ZeZeIM_Server/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMig() {
	dsn := "root:wpzyyds@tcp(127.0.0.1:3306)/zezeim?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移 UserBasic 结构体
	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		panic("自动迁移数据库失败")
	}

	// 打印生成的 SQL 语句
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(&models.UserBasic{})
	fmt.Println(stmt.Schema.Table)
	for _, field := range stmt.Schema.Fields {
		fmt.Println(field.DBName, field.DataType, field.Tag)
	}
}
