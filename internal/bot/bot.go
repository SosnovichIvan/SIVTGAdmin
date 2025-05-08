package bot

import (
	"SIVTGAdmin/common/logger"
	"SIVTGAdmin/internal/config"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	BotService interface {
		Start()
	}

	botService struct {
		log logger.Logger
	}
)

const (
	botPrefix = "bot"
)

func InitBotService(log logger.Logger) BotService {
	return &botService{log: log}
}

func (b *botService) Start() {

	b.log.Info(botPrefix, "start connection to bot")

	cfg := config.GetConfig()

	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)

	if err != nil {
		b.log.Fatal(botPrefix, fmt.Sprintf("Init bot with error: %s", err))
	}

	bot.Debug = cfg.App.DebugMode
	b.log.Info(botPrefix, fmt.Sprintf("Authorized on account %s", bot.Self.UserName))
}

// func subscribeMessages(bot *tgbotapi.BotAPI, storage storage.Storage) {
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates := bot.GetUpdatesChan(u)

// 	for update := range updates {

// 		message := update.Message

// 		if update.MyChatMember != nil {
// 			handleMyChatMember(bot, update, storage)
// 		}

// 		if message != nil {
// 			isComand := update.Message.IsCommand()

// 			if isComand {
// 				executingCommands(bot, update, storage)
// 			}

// 		}
// 	}
// }

// func handleMyChatMember(bot *tgbotapi.BotAPI, update tgbotapi.Update, storage storage.Storage) {
// 	newStatus := update.MyChatMember.NewChatMember.Status
// 	oldStatus := update.MyChatMember.OldChatMember.Status

// 	switch newStatus {
// 	case "member":
// 		// Добавление бота в чат
// 		msg := tgbotapi.NewMessage(update.MyChatMember.Chat.ID, fmt.Sprintf("Привет! Меня добавили в этот чат."))
// 		bot.Send(msg)

// 	case "left":
// 		// Удаление бота из чата
// 		if oldStatus != "kicked" && newStatus == "left" {
// 			fmt.Printf("Меня удалили из чата %s\n", update.MyChatMember.Chat.Title)
// 		}
// 	default:
// 		// Другие случаи (например, ограничение отправки сообщений)
// 		fmt.Printf("Изменился статус в чате %s (%v → %v)\n",
// 			update.MyChatMember.Chat.Title, oldStatus, newStatus)
// 	}
// }
