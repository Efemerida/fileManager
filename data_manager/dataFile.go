package data_manager

import (
	"sort"
)

// sortType - тип для описания констант, используемых для сортировки
type sortType string

const (
	Asс  sortType = "asc"  //по возрастанию
	Desс sortType = "desc" //по убыванию
)

const (
	systemSize float32 = 1000 //множитель перевода типа
)

// DataFile - структура описывающая файл
type DataFile struct {
	FileType     string  `json:"file_type"`      // тип файла(директория/файл)
	FileSize     float32 `json:"file_size"`      // размер файла в байтах
	FileSizeType string  `json:"file_size_type"` // тип размера файла
	FileName     string  `json:"file_name"`      // название файла
}

// MapToDataFileWithTypeSize - перевод размера файла из байтов в другие типы
func (dataFile *DataFile) MapToDataFileWithTypeSize() {
	newSize, newType := calcTypeSize(dataFile.FileSize)
	dataFile.FileSize = newSize
	dataFile.FileSizeType = newType

}

// calcTypeSize - вычисление более подходящего вида размера и перевод байт в этот размер
func calcTypeSize(size float32) (float32, string) {

	sizeTypes := []string{"байт", "КБ", "МБ", "ГБ", "ТБ"}
	sizeTypesIndex := 0
	var newSize float32 = float32(size)

	for {
		var sizeTmp float32 = newSize / systemSize

		if sizeTmp > 1 {
			newSize = sizeTmp
			sizeTypesIndex++
		} else {
			break
		}

	}
	return newSize, sizeTypes[sizeTypesIndex]
}

// SortDataFiles - сортировка файлов по размеру
func SortDataFiles(dataFiles []DataFile, sortType sortType) {
	if sortType != Asс { // по возрастанию
		sort.Slice(dataFiles, func(i, j int) bool {
			return dataFiles[i].FileSize < dataFiles[j].FileSize
		})
	} else { // по убыванию
		sort.Slice(dataFiles, func(i, j int) bool {
			return dataFiles[i].FileSize > dataFiles[j].FileSize
		})
	}
}

// GetSortType - получение константы, которая характерезует сортировку
func GetSortType(sortType string) sortType {
	if sortType == "desc" {
		return Desс
	}
	return Asс
}
