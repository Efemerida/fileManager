package server

import (
	"bytes"
	"context"
	"encoding/json"
	manager "fileService/server/data_manager"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

// rootDirectory - переменная содержащая корневой каталог из конфига
var rootDirectory string

// statsURL - переменная содержащая URL для отправки статистики
var statsURL string

// doResponse - формирует ответ
func doResponse(w http.ResponseWriter, responseData *response, begunTime time.Time, statsData *stats) {

	//формирование json
	jsonResponse, err := json.MarshalIndent(responseData, "", " ")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//формирование заголовков
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(responseData.Status)

	//формирование тела
	w.Write(jsonResponse)

	if responseData.Data != "" {
		statsData.ElapsedTime = time.Since(begunTime).Seconds()
		sendStats(*statsData)
	}

	//выывод времени запроса
	fmt.Printf("Время обработки запроса:%s\n", time.Since(begunTime))
}

// sendStats - отправление статистики
func sendStats(statsData stats) {

	// кодируем структуру statsData в JSON
	dataJson, err := json.Marshal(statsData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Создаем новый запрос
	request, err := http.NewRequest("POST", statsURL, bytes.NewBuffer(dataJson))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Устанавливаем заголовок с типом данных в теле запроса
	request.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	client := &http.Client{}
	responseStat, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer responseStat.Body.Close()

	// Читаем ответ
	responseBody, err := io.ReadAll(responseStat.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	// Декодируем JSON-ответ
	var responseData map[string]interface{}
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON-ответа:", err)
		return
	}

	// Проверяем статус ответа
	status := responseData["status"].(int)
	message := responseData["message"].(string)
	if status == 200 {
		fmt.Println("Статистика успешно отправлена:", message)
	} else {
		fmt.Println("Ошибка отправки статистики:", message)
	}
}

// handleGetFiles - обработка запроса по пути /fs
func handleGetFiles(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("\nЗапрос %s\n", r.RemoteAddr)
	begunTime := time.Now()

	//переменная содержащая ответ
	responseData := response{}
	statsData := stats{}

	// формирование ответа
	defer doResponse(w, &responseData, begunTime, &statsData)

	//получение параметров из строки запроса
	dst := r.URL.Query().Get("dst") // dst - параметр пути
	sort := r.FormValue("sort")     // sort - параметр сортировки

	if dst == "" {
		responseData.Status = 200
		responseData.TextError = ""
		responseData.Data = ""
		responseData.RootDirectory = rootDirectory
		return
	}

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
	var sumSize float32
	for i := 0; i < len(filesData); i++ {
		sumSize += filesData[i].FileSize
		filesData[i].MapToDataFileWithTypeSize()
	}

	statsData.Root = dst
	statsData.Size = sumSize

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

	statsURLTmp, findStatsURL := os.LookupEnv("STATS_URL")
	if !findStatsURL {
		return fmt.Errorf("ошибка конфига: STATS_URL не найден")
	}
	statsURL = statsURLTmp

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
