package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// Структура для хранения целей
type Goals struct {
	Goal         string `json:"goal"`
	Timeline     string `json:"timeline"`
	Freedom      string `json:"freedom"`
	SalaryTarget int    `json:"salary_target_rub_per_hour"`
}

func main() {
	// Обработчик для главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Привет! Мой API работает. Перейди по ссылке /goals"))
	})

	// Обработчик для /goals
	http.HandleFunc("/goals", func(w http.ResponseWriter, r *http.Request) {
		myGoals := Goals{
			Goal:         "Хочу работать из Бали",
			Timeline:     "6 месяцев",
			Freedom:      "Работать из любой точки мира",
			SalaryTarget: 3000,
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(myGoals)
	})

	// Получаем порт от Heroku или используем 8080 локально
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Сервер запущен на порту: " + port)
	http.ListenAndServe(":"+port, nil)
}
