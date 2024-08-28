package server

// response - структура для ответа на запрос
type response struct {
	Status        int    `json:"status"`     //статус ответа
	TextError     string `json:"text_error"` //текст ошибки
	RootDirectory string `json:"root_dir"`   //корневая директория
	Data          any    `json:"data"`       //данные
}
