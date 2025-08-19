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
	content := "не ругайся"
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
