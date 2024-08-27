package server

import (
	"context"
	"encoding/json"
	manager "fileService/data_manager"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func handleGetFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nЗапрос %s\n", r.RemoteAddr)
	begunTime := time.Now()

	dst := r.URL.Query().Get("dst") // dst - параметр пути
	sort := r.FormValue("sort")     // sort - параметр сортировки

	sortType := manager.GetSortType(sort)

	//чтение размеров файлов в директории
	filesData, errReaddir := manager.ReadDataFileOfDir(dst)

	// ошибка при выполнении
	if errReaddir != nil {
		resp := response{
			Status:    400,
			TextError: fmt.Sprintf("Ошибка работы программы: %s", errReaddir),
			Data:      "",
		}
		jsonResponce, err := json.MarshalIndent(resp, "", " ")
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(400)
		w.Write(jsonResponce)
		fmt.Printf("Время обработки запроса:%s\n", time.Since(begunTime))
		return
	}

	// сортировка
	manager.SortDataFiles(filesData, sortType)

	//перевод данных в транспортировочный вид
	responseData := []DataFileDto{}
	for _, val := range filesData {
		responseData = append(responseData, mapToDataFileDto(val))
	}

	//сериализация в json
	resp := response{
		Status:    200,
		TextError: "",
		Data:      responseData,
	}

	jsonResponce, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// формирование ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonResponce)
	fmt.Printf("Время обработки запроса:%s\n", time.Since(begunTime))
}

// StartServer - функция запуска сервера
func StartServer() {

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
		fmt.Printf("Сервер запущен на порту %s...\n", port)
		errService := server.ListenAndServe()
		if errService != nil && errService != http.ErrServerClosed {
			fmt.Printf("\nошибка сервера\n%s\n", errService)
		}
	}()

	// отслеживание сигналов
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTSTP)
	defer stop()
	<-ctx.Done()

	// закрытие сервера
	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Printf("ошибка закрытия сервера: %s\n", err)
		os.Exit(1)
	}

}
