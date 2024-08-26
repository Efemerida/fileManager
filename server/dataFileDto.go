package server

type DataFileDto struct {
	FileType     string  // тип файла(директория/файл)
	FileSize     float32 // размер файла в байтах
	FileSizeType string  // тип размера файла
	FileName     string  // название файла
}
