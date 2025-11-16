package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// Service представляет собой основную структуру бота
type Service struct {
	api *tgbotapi.BotAPI
}

// NewService инициализирует новое подключение к Telegram API
func NewService(token string) (*Service, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	// botAPI.Debug = true // Раскомментируйте для режима отладки

	return &Service{api: botAPI}, nil
}

// GetUserName возвращает имя пользователя бота
func (s *Service) GetUserName() string {
	return s.api.Self.UserName
}

// Run запускает цикл получения обновлений
func (s *Service) Run() {
	log.Println("Начало получения обновлений...")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 // Долгий опрос

	updates := s.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Передаем сообщение в обработчик
		s.handleMessage(update.Message)
	}
}

// handleMessage содержит логику обработки входящих сообщений
// Бот отправляет то же сообщение, прибавляя имя пользователя.
func (s *Service) handleMessage(msg *tgbotapi.Message) {
	// Игнорируем пустые сообщения (например, стикеры без подписи)
	if msg.Text == "" {
		return
	}

	originalText := msg.Text

	// Получаем имя пользователя (First Name)
	userName := msg.From.FirstName

	// Формируем ответ
	responseText := originalText + ", " + userName

	log.Printf("[%s] %s -> %s", msg.From.UserName, originalText, responseText)

	// Создаем и отправляем сообщение
	replyMsg := tgbotapi.NewMessage(msg.Chat.ID, responseText)

	if _, err := s.api.Send(replyMsg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}
