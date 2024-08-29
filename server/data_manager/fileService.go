package data_manager

import (
	"fmt"
	"os"
	"sync"
)

// ReadDataFileOfDir - получение данных о файлах директории
func ReadDataFileOfDir(pathDirectory string) ([]DataFile, error) {

	//получение файлов директории
	files, err := os.ReadDir(pathDirectory)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать директорию: %s Ошибка: %s", pathDirectory, err)
	}

	var waitGroup sync.WaitGroup
	var filesData = make([]DataFile, len(files))

	for i, file := range files {

		//если файл - это директория, то подсчитывается общий размер файлов в этой директории, затем сохраняется
		if file.IsDir() {
			waitGroup.Add(1)
			go getSizeDirectory(file, pathDirectory, i, filesData, &waitGroup)
		} else {

			fileInfo, errFileINfo := file.Info()
			if errFileINfo != nil {
				fmt.Printf("Не удалось получит данные о файле: %s Ошибка: %s \n", file.Name(), errFileINfo)
				continue
			}
			filesData[i] = DataFile{"Файл", float32(fileInfo.Size()), "байт", file.Name()}
		}
	}

	waitGroup.Wait()

	return filesData, nil

}

// getSizeDirectory - получение и сохранение данных о директории
func getSizeDirectory(dirEnrty os.DirEntry, pathDirectory string, index int, filesData []DataFile, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	fileInfo, errFileINfo := dirEnrty.Info()
	if errFileINfo != nil {
		fmt.Printf("Не удалось получит данные о файле: %s Ошибка: %s \n", fileInfo.Name(), errFileINfo)
		return
	}

	//получение размера директории
	newPath := fmt.Sprintf("%s/%s", pathDirectory, fileInfo.Name())
	directorySum, errCalcSumSizeDirectory := calcSumSizeDirectory(newPath)
	if errCalcSumSizeDirectory != nil {
		fmt.Printf("Не удалось открыть директорию: %s Ошибка: %s \n", newPath, errCalcSumSizeDirectory)
		return
	}
	directorySum += fileInfo.Size()

	//сохранение данных о директории
	filesData[index] = DataFile{"Директория", float32(directorySum), "байт", dirEnrty.Name()}
}

// calcSumSizeDirectory - вычисление суммарного размера директории
func calcSumSizeDirectory(pathDirectory string) (int64, error) {

	//считывание содержания директории
	files, err := os.ReadDir(pathDirectory)
	if err != nil {
		return 4, err //если не удалось открыть папку возвращается ее размер 4 байта
	}

	//суммарный размер директории
	var directorySum int64 = 0

	//проход по кадому файлу
	for _, file := range files {

		//формирование нового пути к внутренней директории и рекурсивный вызов
		//с этим путем в качестве аргумента
		if file.IsDir() {
			newPath := fmt.Sprintf("%s/%s", pathDirectory, file.Name())
			subDirectorySum, errCalcSumSizeDirectory := calcSumSizeDirectory(newPath)
			if errCalcSumSizeDirectory != nil {
				fmt.Printf("Не удалось получить данные о файле: %s Ошибка: %s \n", newPath, errCalcSumSizeDirectory)
				continue
			}

			directorySum += subDirectorySum
		}

		fileInfo, errFileINfo := file.Info()
		if errFileINfo != nil {
			fmt.Printf("Не удалось получит данные о файле: %s Ошибка: %s \n", file.Name(), errFileINfo)
			continue
		}
		directorySum += fileInfo.Size()

	}
	return directorySum, nil

}
