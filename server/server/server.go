package server

import (
	"context"
	"encoding/json"
	manager "fileService/data_manager"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

// rootDirectory - переменная содержащая корневой каталог из конфига
var rootDirectory string

// handleGetFiles - обработка запроса по пути /fs
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
		responseData := response{
			Status:        400,
			TextError:     fmt.Sprintf("Ошибка работы программы: %s", errReaddir),
			Data:          "",
			RootDirectory: "",
		}
		jsonResponce, err := json.MarshalIndent(responseData, "", " ")
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
	for i := 0; i < len(filesData); i++ {
		filesData[i].MapToDataFileWithTypeSize()
	}

	//сериализация в json
	responseData := response{
		Status:        200,
		TextError:     "",
		Data:          filesData,
		RootDirectory: rootDirectory,
	}

	jsonResponce, err := json.MarshalIndent(responseData, "", " ")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// формирование ответа
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(jsonResponce)
	fmt.Printf("Время обработки запроса:%s\n", time.Since(begunTime))
}

// handleGetMainPage - обработка запроса по пути /
func handleGetMainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./../../front/index.html")
}

// StartServer - функция запуска сервера
func StartServer() error {

	// загрузка данных из .env файлов
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("ошибка загрузки .emv файла:%s", err)
	}

	//загрузка порта
	port, findPort := os.LookupEnv("PORT")
	if !findPort {
		return fmt.Errorf("ошибка конфига: PORT не найден")
	}

	//загрузка корневой директории
	rootDirectoryTmp, findRootDirectory := os.LookupEnv("ROOTDIR")
	if !findRootDirectory {
		return fmt.Errorf("ошибка конфига: ROOTDIR не найден")
	}

	rootDirectory = rootDirectoryTmp

	server := http.Server{Addr: port}

	//обработка запроса по пути - /fs
	http.HandleFunc("/fs", handleGetFiles)
	http.HandleFunc("/", handleGetMainPage)

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
		return fmt.Errorf("ошибка закрытия сервера: %s", err)
	}

	return nil

}
