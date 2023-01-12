package gorm

import (
	"GolangNotes/gorm/model"
	"fmt"
)

func GormMain() {
	// AddUser
	user := model.User{Name: "Namae", Password: "pass", Age: 21}
	addErr := model.UserAdd(user)
	if addErr != nil {
		fmt.Println(addErr)
	}
	// GetUser(id)
	id := 1
	getWithId, getErr := model.GetUser(id)
	if getErr != nil {
		fmt.Println("GetUserに失敗:", getErr)
	} else {
		fmt.Println(getWithId)
	}
	// GetListUsers
	users, getListErr := model.GetListUsers()
	if getListErr != nil {
		fmt.Println(getListErr)
	} else {
		fmt.Println(users)
	}
	//UpdateUser
	oldId := 4
	oldUser, targetUserErr := model.GetUser(oldId)
	if targetUserErr != nil {
		fmt.Println("GetUserに失敗:", targetUserErr)
	} else {
		newUser := model.User{
			Name:     "id4User",
			Password: "4pass",
			Age:      444,
		}
		new, updateErr := model.UpdateUser(oldUser, newUser)
		if updateErr != nil {
			fmt.Println(updateErr)
		}
		fmt.Println("変更前:", oldUser)
		fmt.Println("変更後:", new)
	}
	//DeleteUser
	delID := 6
	delErr := model.DeleteUser(delID)
	if delErr != nil {
		fmt.Println(delErr)
	}

}
