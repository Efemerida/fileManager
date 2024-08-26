package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	manager "fileService/data_manager"

	"github.com/joho/godotenv"
)

func handleGetFiles(w http.ResponseWriter, r *http.Request) {
	dst := r.URL.Query().Get("dst") // dst - параметр пути
	sort := r.FormValue("sort")     // sort - параметр сортировки

	var sortFlag bool = true

	if sort == "asc" {
		sortFlag = true
	} else if sort == "desc" {
		sortFlag = false
	}

	//чтение размеров файлов в директории
	filesData, errReaddir := manager.ReadDataFileOfDir(dst)

	// ошибка при выполнении
	if errReaddir != nil {
		w.WriteHeader(400)
		errorStr := fmt.Sprintf("Ошибка работы программы: \n%s\n", errReaddir)
		w.Write([]byte(errorStr))
		return
	}

	// сортировка
	manager.SortDataFiles(filesData, sortFlag)

	responseData := []DataFileDto{}

	for _, val := range filesData {
		responseData = append(responseData, val.mapToDataFileDto())
	}

	//сериализация в json
	jsonResponce, err := json.MarshalIndent(responseData, "", " ")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(jsonResponce)
}

func mainn() {
	if err := godotenv.Load(); err != nil {
		log.Print("файл .env не найден")
	}

	port, find := os.LookupEnv("PORT")
	if !find {
		fmt.Printf("ошибка конфига: PORT не найден\n")
		return
	}

	server := http.Server{Addr: port}

	//обработка запроса по пути - /
	http.HandleFunc("/", handleGetFiles)

	//запуск сервера
	go func() {
		errService := server.ListenAndServe()
		if errService != nil {
			fmt.Printf("ошибка сервера\n%s\n", errService)
		}
	}()

	fmt.Printf("Сервер запущен на порту %s...\nВведите exit чтобы закрыть сервер\n", port)

	//ожидание закрытия сервера
	for {
		st := ""
		fmt.Scan(&st)
		if st == "exit" {
			server.Close()
			break
		}

	}
}
