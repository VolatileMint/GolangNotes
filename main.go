package main

import (
	"GolangNotes/echoNotes"
	"GolangNotes/gorm"
)

func main() {
	flag := "echo"

	if flag == "gorm" {
		gorm.GormMain()
	} else if flag == "echo" {
		echoNotes.EchoMain()
	}
}
