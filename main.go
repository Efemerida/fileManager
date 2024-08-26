package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	//метка старта программы
	begunTime := time.Now()

	//чтение флагов
	pathMainDirectory, sortFlag, err := readFlags()
	if err != nil {
		fmt.Printf("Ошибка запуска: \n%s\n", err)
		os.Exit(1)
	}

	//получение данных о файлах директории
	filesData, errReaddir := readDataFileOfDir(pathMainDirectory)
	if errReaddir != nil {
		fmt.Printf("Ошибка работы программы: \n%s\n", errReaddir)
		os.Exit(1)
	}

	SortDataFiles(filesData, sortFlag)
	printFilesData(filesData)
	//метка завершения программы
	endTime := time.Now()

	fmt.Printf("Время выполнения программы: %s\n", endTime.Sub(begunTime))
}

// readDataFileOfDir - получение данных о файлах директории
func readDataFileOfDir(pathDirectory string) ([]DataFile, error) {

	//получение файлов директории
	files, err := os.ReadDir(pathDirectory)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать директорию: %s\nОшибка: %s", pathDirectory, err)
	}

	var wg sync.WaitGroup
	var filesData = make([]DataFile, len(files))

	for i, file := range files {

		//если файл - это директория, то подсчитывается общий размер файлов в этой директории, затем сохраняется
		if file.IsDir() {
			wg.Add(1)
			go getSizeDirectory(file, pathDirectory, i, filesData, &wg)
		} else {

			info, errFileINfo := file.Info()
			if errFileINfo != nil {
				fmt.Printf("\nНе удалось получит данные о файле: %s\nОшибка: %s\n\n", file.Name(), errFileINfo)
				continue
			}
			filesData[i] = DataFile{"Файл", info.Size(), file.Name()}
		}
	}

	wg.Wait()

	return filesData, nil

}

// printFilesData - печать данных о файлах в директории
func printFilesData(filesData []DataFile) {

	for _, dataFile := range filesData {
		dataFile.Print()
	}
}

// getSizeDirectory - получение и сохранение данных о директории
func getSizeDirectory(file os.DirEntry, pathDirectory string, index int, filesData []DataFile, wg *sync.WaitGroup) {

	defer wg.Done()

	fileInfo, errFileINfo := file.Info()
	if errFileINfo != nil {
		fmt.Printf("\nНе удалось получит данные о файле: %s\nОшибка: %s\n\n", fileInfo.Name(), errFileINfo)
		return
	}

	//получение размера директории
	newPath := fmt.Sprintf("%s/%s", pathDirectory, fileInfo.Name())
	dirSum, errCalcSumSizeDirectory := calcSumSizeDirectory(newPath)
	if errCalcSumSizeDirectory != nil {
		fmt.Printf("\nНе удалось открыть директорию: %s\nОшибка: %s\n\n", newPath, errCalcSumSizeDirectory)
		return
	}
	dirSum += fileInfo.Size()

	//сохранение данных о директории
	filesData[index] = DataFile{"Директория", dirSum, file.Name()}
}

// calcSumSizeDirectory - вычисление суммарного размера директории
func calcSumSizeDirectory(pathDirectory string) (int64, error) {

	//считывание содержания директории
	files, err := os.ReadDir(pathDirectory)
	if err != nil {
		return 0, err
	}

	//суммарный размер директории
	var sum int64 = 0

	//проход по кадому файлу
	for _, file := range files {

		//формирование нового пути к внутренней директории и рекурсивный вызов
		//с этим путем в качестве аргумента
		if file.IsDir() {
			newPath := fmt.Sprintf("%s/%s", pathDirectory, file.Name())
			dirSum, errCalcSumSizeDirectory := calcSumSizeDirectory(newPath)
			if errCalcSumSizeDirectory != nil {
				fmt.Printf("Не удалось получить данные о файле: %s\nОшибка: %s\n\n", newPath, errCalcSumSizeDirectory)
				continue
			}

			sum += dirSum
			continue
		}

		info, errFileINfo := file.Info()
		if errFileINfo != nil {
			fmt.Printf("\nНе удалось получит данные о файле: %s\nОшибка: %s\n\n", file.Name(), errFileINfo)
			continue
		}
		sum += info.Size()

	}
	return sum, nil

}

// readFlugs - считывание фалагов
func readFlags() (string, bool, error) {
	directoryPath := flag.String("dst", "", "Путь целевой директории")
	sort := flag.String("sort", "", "Сортировка по возрастанию/убыванию")

	flag.Parse()

	if *directoryPath == "" {
		flag.PrintDefaults()
		return "", false, fmt.Errorf("не указана целевая директория")
	}

	if *sort == "ask" {
		return *directoryPath, true, nil
	} else if *sort == "desk" {
		return *directoryPath, false, nil
	}
	flag.PrintDefaults()
	fmt.Printf("Вы не указали сортировку или указали не корректно\nБудет использоваться значение по умолчанию (asc)\n")
	return *directoryPath, true, nil

}
