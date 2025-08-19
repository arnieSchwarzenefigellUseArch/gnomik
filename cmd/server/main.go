package main

import (
	"fmt"
	"os"
)

func main() {
	createFile()
	readfile()
}

func createFile() {
<<<<<<< HEAD
	content := "ЖРИ СВОЙ ЕБАНЫЙ ФАЙЛ, кстати да, это из файла тебе привет и кстати опять привет папа гыгыгыгыг, не ругайся" &&
=======


	content := "не ругайся"
>>>>>>> 6c591b899b3cd968824486a491ab86d6c674f5a2
	err := os.WriteFile("data.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Ошибка создания: %v\n", err)
		return
	}
}

func readfile() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Printf("АХТУНГ БЛЯДЬ: %v\n", err)
		return
	}
	fmt.Printf("держи бля: %s\n", data)
}
