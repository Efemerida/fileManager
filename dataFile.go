package main

import (
	"fmt"
	"sort"
)

type DataFile struct {
	FileType string
	FileSize int64
	FileName string
}

func NewDataFile(fileType string,
	fileSize int64,
	fileName string) *DataFile {
	return &DataFile{
		FileType: fileType,
		FileSize: fileSize,
		FileName: fileName,
	}
}

// Print - печать в консоли
func (dataFile *DataFile) Print() {
	newSize, newType := calcTypeSize(dataFile.FileSize)
	fmt.Printf("%-15s %-10.2f %-16s %-10s\n", dataFile.FileType, newSize, newType, dataFile.FileName)
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

func SortDataFiles(dataFiles []DataFile, ask bool) {
	if !ask {
		sort.Slice(dataFiles, func(i, j int) bool {
			return dataFiles[i].FileSize < dataFiles[j].FileSize
		})
	} else {
		sort.Slice(dataFiles, func(i, j int) bool {
			return dataFiles[i].FileSize > dataFiles[j].FileSize
		})
	}
}
