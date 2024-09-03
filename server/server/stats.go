package server

// stats - структура для отправки статистики
type stats struct {
	Root        string  `json:"root"` //обрабатываемая директория
	Size        float64 `json:"size"`
	ElapsedTime float64 `json:"elapsed_time"` //время выполнения
}
