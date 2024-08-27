package server

// response - структура для ответа на запрос
type response struct {
	Status    int    `json:"status"`
	TextError string `json:"text_error"`
	Data      any    `json:"data"`
}
