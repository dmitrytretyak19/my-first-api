package main

import (
	"encoding/json"
	"net/http"
)

// Структура для хранения целей
type Goals struct {
	Goal         string `json:"goal"`
	Timeline     string `json:"timeline"`
	Freedom      string `json:"freedom"`
	SalaryTarget int    `json:"salary_target_rub_per_hour"`
}

func main() {
	// Создаём "меню" (роутер)
	http.HandleFunc("/goals", func(w http.ResponseWriter, r *http.Request) {
		// Наши цели
		myGoals := Goals{
			Goal:         "Хочу работать из Бали",
			Timeline:     "6 месяцев",
			Freedom:      "Работать из любой точки мира",
			SalaryTarget: 3000,
		}

		// Преобразуем в JSON
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(myGoals)
	})

	// Запускаем "ресторан" на порту 8080
	println("Сервер запущен! Открой в браузере: http://localhost:8080/goals")
	http.ListenAndServe(":8080", nil)
}
