package main

import (
	"log"
	"maria-english-go/internal/bot"
	"maria-english-go/internal/config"
)

func main() {
	// 1. Загрузка конфигурации
	cfg := config.LoadConfig()

	// 2. Инициализация сервиса бота
	telegramBot, err := bot.NewService(cfg.Token)
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}

	// 3. Запуск
	log.Printf("Бот запущен. Авторизован как %s", telegramBot.GetUserName())
	telegramBot.Run()
}
