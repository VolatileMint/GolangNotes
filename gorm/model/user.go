package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Name       string `json:"name" xml:"name" form:"name" query:"name"`
	Password   string `json:"password" xml:"password" form:"password" query:"password"`
	Age        int    `json:"age" xml:"age" form:"age" query:"age"`
	gorm.Model        // ID uint, CreatedAt time.Time, UpdatedAt time.Time, DeletedAt gorm.DeletedAt
}

// Userを引数に Userとerrorを返す
func UserAdd(u User) error {
	db := Connect()
	result := db.Debug().Create(&u)
	return result.Error
}

// intを引数に、Userを検索
func GetUser(id int) (User, error) {
	db := Connect()
	u := User{}
	result := db.Debug().Where("id = ?", id).First(&u)
	//fmt.Println(u, result.Error)
	return u, result.Error
}

/* 条件を指定してUserを検索
 * 例：Age < 20
 * Created_at 2022-01-01 ~ 2022-04-01など
 */

// 引数無しに、[]Userとerrorを返す
func GetListUsers() ([]User, error) {
	db := Connect()
	var users []User
	result := db.Find(&users)
	fmt.Println(users, result.Error)
	return users, result.Error
}

// idを引数にユーザーを削除(deleted_atに日付を追加)
func DeleteUser(id int) error {
	db := Connect()
	result := db.Debug().Delete(&User{}, id)
	return result.Error
}

// 引数のUserで、更新する
func UpdateUser(user User, new User) (User, error) {
	db := Connect()

	db.First(&user).Debug()
	user.Name = new.Name
	user.Password = new.Password
	user.Age = new.Age
	fmt.Println("更新直前", user)
	result := db.Save(&user).Debug()
	fmt.Println("アップデートの成否", result.Error)

	return user, result.Error
}
