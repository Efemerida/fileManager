package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	//метка старта программы
	begunTime := time.Now()

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	port, find := os.LookupEnv("PORT")
	if !find {
		fmt.Printf("ошибка конфига: PORT не найден\n")
		return
	}

	server := http.Server{Addr: port}

	//обработка запроса по пути - /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		dst := r.URL.Query().Get("dst") // dst - параметр пути
		sort := r.FormValue("sort")     // sort - параметр сортировки

		var sortFlag bool = true

		if sort == "asc" {
			sortFlag = true
		} else if sort == "desc" {
			sortFlag = false
		}

		//чтение размеров файлов в директории
		filesData, errReaddir := ReadDataFileOfDir(dst)

		// ошибка при выполнении
		if errReaddir != nil {
			w.WriteHeader(404)
			errorStr := fmt.Sprintf("Ошибка работы программы: \n%s\n", errReaddir)
			w.Write([]byte(errorStr))
			return
		}

		// сортировка
		SortDataFiles(filesData, sortFlag)

		//сериализация в json
		jsonResponce, err := json.MarshalIndent(filesData, "", " ")
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(jsonResponce)

	})

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

	//метка завершения программы
	endTime := time.Now()
	fmt.Printf("Время выполнения программы: %s\n", endTime.Sub(begunTime))
}
