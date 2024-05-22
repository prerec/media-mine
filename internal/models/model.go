package models

// Request модель для обработки запроса пользователя
type Request struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}

// Response модель для ответа пользователю
type Response struct {
	Exchanges [][]int `json:"exchanges"`
}

// ErrorResponse модель для возврата ошибок
type ErrorResponse struct {
	Error string `json:"error"`
}
