package data_manager

import (
	"fmt"
	"os"
	"sync"
)

// readDataFileOfDir - получение данных о файлах директории
func ReadDataFileOfDir(pathDirectory string) ([]DataFile, error) {

	//получение файлов директории
	files, err := os.ReadDir(pathDirectory)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать директорию: %s Ошибка: %s", pathDirectory, err)
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
				fmt.Printf("Не удалось получит данные о файле: %s Ошибка: %s ", file.Name(), errFileINfo)
				continue
			}
			filesData[i] = DataFile{"Файл", info.Size(), file.Name()}
		}
	}

	wg.Wait()

	return filesData, nil

}

// getSizeDirectory - получение и сохранение данных о директории
func getSizeDirectory(file os.DirEntry, pathDirectory string, index int, filesData []DataFile, wg *sync.WaitGroup) {

	defer wg.Done()

	fileInfo, errFileINfo := file.Info()
	if errFileINfo != nil {
		fmt.Printf("Не удалось получит данные о файле: %s Ошибка: %s ", fileInfo.Name(), errFileINfo)
		return
	}

	//получение размера директории
	newPath := fmt.Sprintf("%s/%s", pathDirectory, fileInfo.Name())
	dirSum, errCalcSumSizeDirectory := calcSumSizeDirectory(newPath)
	if errCalcSumSizeDirectory != nil {
		fmt.Printf("Не удалось открыть директорию: %s Ошибка: %s ", newPath, errCalcSumSizeDirectory)
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
				fmt.Printf("Не удалось получить данные о файле: %s Ошибка: %s ", newPath, errCalcSumSizeDirectory)
				continue
			}

			sum += dirSum
			continue
		}

		info, errFileINfo := file.Info()
		if errFileINfo != nil {
			fmt.Printf("Не удалось получит данные о файле: %s Ошибка: %s ", file.Name(), errFileINfo)
			continue
		}
		sum += info.Size()

	}
	return sum, nil

}
