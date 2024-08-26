package data_manager

import (
	"fileService/server"
	"fmt"
	"sort"
)

// DataFile - структура описывающая файл
type DataFile struct {
	FileType string // тип файла(директория/файл)
	FileSize int64  // размер файла в байтах
	FileName string // название файла
}

// Print - печать в консоли
func (dataFile *DataFile) Print() {
	newSize, newType := calcTypeSize(dataFile.FileSize)
	fmt.Printf("%-15s %-10.2f %-16s %-10s\n", dataFile.FileType, newSize, newType, dataFile.FileName)
}

func (dataFile *DataFile) mapToDataFileDto() server.DataFileDto {
	newSize, newType := calcTypeSize(dataFile.FileSize)

	return server.DataFileDto{
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

// SortDataFiles - сортировка файлов по размеру
func SortDataFiles(dataFiles []DataFile, typeSort bool) {
	if !typeSort {
		sort.Slice(dataFiles, func(i, j int) bool {
			return dataFiles[i].FileSize < dataFiles[j].FileSize
		})
	} else {
		sort.Slice(dataFiles, func(i, j int) bool {
			return dataFiles[i].FileSize > dataFiles[j].FileSize
		})
	}
}
