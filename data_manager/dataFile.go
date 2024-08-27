package data_manager

import (
	"sort"
)

type sortType string

const (
	Ask  sortType = "asc"
	Desk sortType = "desc"
)

// DataFile - структура описывающая файл
type DataFile struct {
	FileType string // тип файла(директория/файл)
	FileSize int64  // размер файла в байтах
	FileName string // название файла
}

// SortDataFiles - сортировка файлов по размеру
func SortDataFiles(dataFiles []DataFile, sortType sortType) {
	if sortType != Ask {
		sort.Slice(dataFiles, func(i, j int) bool {
			return dataFiles[i].FileSize < dataFiles[j].FileSize
		})
	} else {
		sort.Slice(dataFiles, func(i, j int) bool {
			return dataFiles[i].FileSize > dataFiles[j].FileSize
		})
	}
}

func GetSortType(sortType string) sortType {
	if sortType == "desc" {
		return Desk
	}
	return Ask
}
