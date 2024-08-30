package server

import (
	"context"
	"encoding/json"
	manager "fileService/server/data_manager"
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

// doResponce - формирует ответ
func doResponce(w http.ResponseWriter, responseData *response, begunTime time.Time) {

	//формирование json
	jsonResponce, err := json.MarshalIndent(responseData, "", " ")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//формирование заголовков
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(responseData.Status)

	//формирование тела
	w.Write(jsonResponce)

	//выывод времени запроса
	fmt.Printf("Время обработки запроса:%s\n", time.Since(begunTime))
}

// handleGetFiles - обработка запроса по пути /fs
func handleGetFiles(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("\nЗапрос %s\n", r.RemoteAddr)
	begunTime := time.Now()

	//переменная содержащая ответ
	responseData := response{}

	// формирование ответа
	defer doResponce(w, &responseData, begunTime)

	//получение параметров из строки запроса
	dst := r.URL.Query().Get("dst") // dst - параметр пути
	sort := r.FormValue("sort")     // sort - параметр сортировки

	sortType := manager.GetSortType(sort)

	//чтение размеров файлов в директории
	filesData, errReaddir := manager.ReadDataFileOfDir(dst)

	// ошибка при выполнении
	if errReaddir != nil {

		responseData.Status = 400
		responseData.TextError = fmt.Sprintf("Ошибка работы программы: %s", errReaddir)
		responseData.Data = ""
		responseData.RootDirectory = ""
		return
	}

	// сортировка
	manager.SortDataFiles(filesData, sortType)

	//перевод данных в транспортировочный вид
	for i := 0; i < len(filesData); i++ {
		filesData[i].MapToDataFileWithTypeSize()
	}

	//запись результата работы
	responseData.Status = 200
	responseData.TextError = ""
	responseData.Data = filesData
	responseData.RootDirectory = rootDirectory

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

	// handleGetMainPage - обработка запроса по пути /
	fs := http.FileServer(http.Dir("./static/resource"))

	//обработка запроса по пути - /fs
	http.HandleFunc("/fs", handleGetFiles)
	http.Handle("/", fs)

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
