package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	//метка старта программы
	begunTime := time.Now()

	//чтение флагов
	pathMainDirectory, err := readFlags()
	if err != nil {
		fmt.Printf("Ошибка запуска: \n%s\n", err)
		os.Exit(1)
	}

	//получение данных о файлах директории
	errReaddir := readDataFileOfDir(pathMainDirectory)
	if errReaddir != nil {
		fmt.Printf("Ошибка работы программы: \n%s\n", err)
		os.Exit(1)
	}

	//метка завершения программы
	endTime := time.Now()

	fmt.Printf("Время выполнения программы: %s\n", endTime.Sub(begunTime))
}

// readDataFileOfDir - получение данных о файлах директории
func readDataFileOfDir(pathDirectory string) error {

	//получение файлов директории
	files, err := os.ReadDir(pathDirectory)
	if err != nil {
		return fmt.Errorf("не удалось прочитать директорию: %s\nОшибка: %s", pathDirectory, err)
	}

	for _, file := range files {

		isDir := file.IsDir()

		var fileSize int64 = 0
		var fileType = ""

		//если файл - это директория, то подсчитывается общий размер файлов в этой директории, затем печатается
		if isDir {

			//формирование пути вложенной директории и получение ее размера
			newPath := fmt.Sprintf("%s/%s", pathDirectory, file.Name())
			dirSum, errCalcSumSizeDirectory := calcSumSizeDirectory(newPath)
			if errCalcSumSizeDirectory != nil {
				fmt.Printf("\nНе удалось открыть директорию: %s\nОшибка: %s\n\n", newPath, errCalcSumSizeDirectory)
				continue
			}

			fileSize += dirSum
			fileType = "Директория"
		} else {
			fileType = "Файл"
		}

		//печать файла
		info, errFileINfo := file.Info()
		if errFileINfo != nil {
			fmt.Printf("\nНе удалось получит данные о файле: %s\nОшибка: %s\n\n", file.Name(), errFileINfo)
			continue
		}

		fileSize += info.Size()

		size, typeSize := calcTypeSize(fileSize)

		fmt.Printf("%-15s %-5.1f %-10s%-10s\n", fileType, size, typeSize, file.Name())
	}
	return nil

}

// calcTypeSize - вычисление более подходящего вида размера и перевод байт в этот размер
func calcTypeSize(size int64) (float32, string) {

	sizetypes := []string{"байт", "КБ", "МБ", "ГБ", "ТБ"}
	typeSize := 0
	var newSize float32 = float32(size)

	for {
		var sizeTmp float32 = newSize / 1000.0

		if sizeTmp > 1 {
			newSize = sizeTmp
			typeSize++
		} else {
			break
		}

	}
	return newSize, sizetypes[typeSize]
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
func readFlags() (string, error) {
	directoryPath := flag.String("dst", "", "Путь целевой директории")

	flag.Parse()

	if *directoryPath == "" {
		flag.PrintDefaults()
		return "", fmt.Errorf("не указана целевая директория")
	}

	return *directoryPath, nil
}
