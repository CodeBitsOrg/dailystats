package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TelegramProvider interface {
	Send(chat_id, message string) error
}

type BotProvider interface {
	SendMessage(ctx context.Context, params *bot.SendMessageParams) (*models.Message, error)
}

type TelegramClient struct {
	Bot BotProvider
}

func New(bot BotProvider) *TelegramClient {
	return &TelegramClient{Bot: bot}
}

func (t *TelegramClient) Send(chat_id, message string) error {
	_, err := t.Bot.SendMessage(context.TODO(), &bot.SendMessageParams{
		ChatID: chat_id,
		Text:   message,
	})

	return err
}
