package server

// stats - структура для отправки статистики
type stats struct {
	Root        string  `json:"root"`         //обрабатываемая директория
	Size        float32 `json:"size"`         //размер в байтах
	ElapsedTime float64 `json:"elapsed_time"` //время выполнения
}
