package server

import (
	manager "fileService/data_manager"
)

const (
	systemSize float32 = 1000
)

// DataFileDto - структура для передачи данных о файле
type DataFileDto struct {
	FileType     string  // тип файла(директория/файл)
	FileSize     float32 // размер файла в байтах
	FileSizeType string  // тип размера файла
	FileName     string  // название файла
}

// перевод из структуры работы в структуру для передачи
func mapToDataFileDto(dataFile manager.DataFile) DataFileDto {
	newSize, newType := calcTypeSize(dataFile.FileSize)

	return DataFileDto{
		FileType:     dataFile.FileType,
		FileSize:     newSize,
		FileSizeType: newType,
		FileName:     dataFile.FileName,
	}

}

// calcTypeSize - вычисление более подходящего вида размера и перевод байт в этот размер
func calcTypeSize(size int64) (float32, string) {

	sizetypes := []string{"байт", "КБ", "МБ", "ГБ", "ТБ"}
	typeSize := 0
	var newSize float32 = float32(size)

	for {
		var sizeTmp float32 = newSize / systemSize

		if sizeTmp > 1 {
			newSize = sizeTmp
			typeSize++
		} else {
			break
		}

	}
	return newSize, sizetypes[typeSize]
}
