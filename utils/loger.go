package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	fileName string = "log.txt"
	appName  string = "LuaApp"
)

//Init функция инициализации лог файла
func Init() {
	err := os.Remove(fileName)

	if err != nil {
		fmt.Println(err)
		file, err := os.Create(fileName)
		defer file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
	file, err := os.Create(fileName)
	defer file.Close()
}

//FatalMessage функция записи в лог сообщения о фатальной ошибке
func FatalMessage(errMsg string) {
	time := time.Now().Format("2006-01-02 15:04:05")
	msg := appName + " " + time + " FATAL ERROR: " + errMsg
	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(msg)
	os.Exit(1)
}

//InfoMessage функция для записи в лог информации
func InfoMessage(errMsg string) {
	msg := appName + " INFO: " + errMsg
	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(msg)
}
