package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_test/model"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open("root:123456@(192.168.1.103:3306)/aaa?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})

}

func main() {

	//db.Create(&u)
	//u.Model.ID = 2
	//db.Model(&u).Update("name", "goubi")

	db.Model(&model.User{}).Create([]map[string]interface{}{
		{"Name": sql.NullString{
			String: "hello",
			Valid:  true,
		}, "Age": 20}, {
			"Name": sql.NullString{
				String: "",
				Valid:  false,
			}, "Age": 30,
		},
	})

	/*db.Debug().First(&u)
	fmt.Println(u.Model.ID, u.Name.String, u.Age)*/
}
