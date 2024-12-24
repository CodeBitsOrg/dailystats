package telegram_test

import (
	"context"
	"github.com/CodeBitsOrg/dailystats/telegram"
	mocks "github.com/CodeBitsOrg/dailystats/telegram/mocks"
	"github.com/go-telegram/bot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSend(t *testing.T) {
	// Arrange
	chatID := "123456789"
	message := "Hello, World!"

	contextMatcher := mock.MatchedBy(func(ctx context.Context) bool {
		return true
	})

	botClient := mocks.NewBotProvider(t)
	botClient.On("SendMessage", contextMatcher, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   message,
	}).Return(nil, nil)

	// Act
	tClient := telegram.New(botClient)
	err := tClient.Send(chatID, message)

	// Assert
	assert.NoError(t, err)
}
