package service

import (
	"context"
	"fmt"
	"math/rand"
	"real-holat/storage/repo"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type VerificationCode struct {
	ID         uint  `gorm:"primaryKey"`
	TelegramID int64 `gorm:"index"`
	Phone      string
	Code       string `gorm:"size:6"`
	ExpiresAt  time.Time
	CreatedAt  time.Time
}

type TelegramBot struct {
	Api   *tgbotapi.BotAPI
	Verif *VerificationService
}

func NewTelegramBot(token string, verif *VerificationService) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{Api: bot, Verif: verif}, nil
}

func (b *TelegramBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.Api.GetUpdatesChan(u)

	for update := range updates {
		// handle callback button presses (inline keyboard)
		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "login" {
				if update.CallbackQuery.Message != nil {
					b.handleAuthRequest(update.CallbackQuery.Message)
				}
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		// 1. Handle /start and /login commands
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start", "login":
				b.handleAuthRequest(update.Message)
			}
			continue
		}

		// 2. Handle Contact sharing
		if update.Message.Contact != nil {
			b.handleContact(update.Message)
			continue
		}
	}
}

func (b *TelegramBot) handleAuthRequest(msg *tgbotapi.Message) {
	// Check if a valid code already exists for this user (within 1 minute)
	lastCode, err := b.Verif.GetValid(context.Background(), msg.From.ID)
	if err == nil && lastCode != nil {
		// Code still valid
		text := "🇺🇿 Eski kodingiz hali ham amal qiladi, iltimos uni ishlating yoki 1 daqiqa kuting.\n" +
			"🇺🇸 Your old code is still valid, please use it or wait 1 minute."

		reply := tgbotapi.NewMessage(msg.Chat.ID, text)
		// Add a "Renew" button just like the example
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🔄 Yangilash / Renew", "login"),
			),
		)
		reply.ReplyMarkup = keyboard
		b.Api.Send(reply)
		return
	}

	// Ask for contact if no valid code exists
	text := fmt.Sprintf("🇺🇿\nSalom %s 👋\n@safar ning rasmiy botiga xush kelibsiz\n\n⬇️ Kontaktingizni yuboring (tugmani bosib)\n\n"+
		"🇺🇸\nHi %s 👋\nWelcome to @safar's official bot\n\n⬇️ Send your contact (by clicking button)", msg.From.FirstName, msg.From.FirstName)

	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	btn := tgbotapi.NewKeyboardButtonContact("📱 Kontaktni yuborish / Send Contact")
	row := tgbotapi.NewKeyboardButtonRow(btn)

	keyboard := tgbotapi.NewReplyKeyboard(row)

	keyboard.OneTimeKeyboard = true
	keyboard.ResizeKeyboard = true

	reply.ReplyMarkup = keyboard

	b.Api.Send(reply)
}

func (b *TelegramBot) handleContact(msg *tgbotapi.Message) {
	contact := msg.Contact

	fmt.Println("---- MESSAGE INFO ----")
	fmt.Println("MessageID:", msg.MessageID)
	fmt.Println("Date:", msg.Date)

	if msg.Chat != nil {
		fmt.Println("ChatID:", msg.Chat.ID)
		fmt.Println("ChatType:", msg.Chat.Type)
		fmt.Println("ChatTitle:", msg.Chat.Title)
	}

	if msg.From != nil {
		fmt.Println("---- USER INFO ----")
		fmt.Println("UserID:", msg.From.ID)
		fmt.Println("Username:", msg.From.UserName)
		fmt.Println("FirstName:", msg.From.FirstName)
		fmt.Println("LastName:", msg.From.LastName)
		fmt.Println("LanguageCode:", msg.From.LanguageCode)
		fmt.Println("IsBot:", msg.From.IsBot)
	}

	if contact != nil {
		fmt.Println("---- CONTACT INFO ----")
		fmt.Println("PhoneNumber:", contact.PhoneNumber)
		fmt.Println("FirstName:", contact.FirstName)
		fmt.Println("LastName:", contact.LastName)
		fmt.Println("UserID:", contact.UserID)
		fmt.Println("VCard:", contact.VCard)
	}

	// Generate 6-digit code
	code := fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000))
	expiresAt := time.Now().Add(1 * time.Minute)

	// Save to DB via verification service
	err := b.Verif.Create(context.Background(), &repo.VerificationModel{
		TelegramID:     contact.UserID,
		Phone:          contact.PhoneNumber,
		Code:           code,
		ExpiresAt:      expiresAt,
		TgUserName:     msg.From.UserName,
		TgFirstName:    msg.From.FirstName,
		TgLanguageCode: msg.From.LanguageCode,
	})
	if err != nil {
		b.Api.Send(tgbotapi.NewMessage(msg.Chat.ID, "Internal error saving code"))
		return
	}

	responseText := fmt.Sprintf("🔒 Code: %s\n\n🇺🇿\n🔑 Yangi kod olish uchun /login ni bosing\n\n🇺🇸\n🔑 To get a new code click /login", code)

	reply := tgbotapi.NewMessage(msg.Chat.ID, responseText)
	reply.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	b.Api.Send(reply)
}
