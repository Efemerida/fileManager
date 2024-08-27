package main

import (
	"fileService/server"
	"fmt"
	"os"
)

func main() {

	err := server.StartServer()
	if err != nil {
		fmt.Printf("Ошибка сервера: %s\n", err)
		os.Exit(1)
	}

}
