package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

/*
	type User struct {
		gorm.Model
		Name         string
		Age          sql.NullInt64
		Birthday     *time.Time
		Email        string  `gorm:"type:varchar(100);unique_index"`
		Role         string  `gorm:"size:255"`
		MemberNumber *string `gorm:"unique;not null"`
		Num          int     `gorm:"AUTO_INCREMENT"`
		Address      string  `gorm:"index:addr"`
		IgnoreMe     int     `gorm:"-"`
	}
*/
type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}
type User struct {
	gorm.Model
	Name string
	//Age  *int `gorm:"default:18"`
	Age sql.NullInt64 `gorm:"default:18"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	var user = User{Name: "ssgh1"}
	db.Find(&user).Row().Scan(&user.Name, &user.Age)
	fmt.Printf("name:%s, age:%d\n", user.Name, user.Age)
	//db.Create(&user)
	/*	var animal = Animal{Age: 99, Name: ""}
		db.AutoMigrate(&Animal{})
		db.Create(&animal)
	*/
	/*	//自动检查Product结构是否变化，变化则进行迁移
		db.AutoMigrate(&Product{})
		db.AutoMigrate(&User{})
		str := "ssgh11"
		user := User{Name: "ssgh", Email: "upupsgh@16311.com", Role: "student", Address: "china", MemberNumber: &str}
		println(db.NewRecord(user))
		//增
		db.Create(&Product{Code: "l1212", Price: 10000})

		db.Create(&user)
		println(db.NewRecord(user))
		//查
		var product Product
		db.First(&product, 1)
		db.First(&product, "code = ?", "l1212")
		//改
		db.Model(&product).Update("Price", 20000)
		//删
		//db.Delete(&product)
	*/
}
