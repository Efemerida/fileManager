package server

type response struct {
	Status    int    //код ответа
	TextError string // текст ошибки
	Data      any    // данные
}
