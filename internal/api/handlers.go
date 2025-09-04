package api

import (
	"encoding/json"
	"gnomik/internal/models"
	"net/http"
	"strconv"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Convert(w http.ResponseWriter, r *http.Request) {
	// Разрешаем CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	// Если это предварительный запрос
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}
	
	// Принимаем только POST запросы
	if r.Method != "POST" {
		http.Error(w, "Только POST", 405)
		return
	}
	
	// Читаем JSON из тела запроса
	var req models.ConvertRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := models.ConvertResponse{Error: "Неверный JSON"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Проверяем число
	if req.Number < 0 {
		response := models.ConvertResponse{Error: "Число должно быть положительным"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// КОНВЕРТИРУЕМ В ДВОИЧНУЮ СИСТЕМУ
	binary := strconv.FormatInt(int64(req.Number), 2)
	
	// Отправляем результат
	response := models.ConvertResponse{
		Decimal: req.Number,
		Binary:  binary,
	}
	
	json.NewEncoder(w).Encode(response)
}