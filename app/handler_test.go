package app_test

import (
	"fmt"
	"github.com/CodeBitsOrg/dailystats/app"
	"github.com/CodeBitsOrg/dailystats/stats"
	smocks "github.com/CodeBitsOrg/dailystats/stats/mocks"
	tmocks "github.com/CodeBitsOrg/dailystats/telegram/mocks"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestStatsEndpoint(t *testing.T) {
	// Arrange
	statsClient := smocks.NewStatsProvider(t)
	statsClient.On("GetStats").Return(stats.DailyStats{}, nil)

	tClient := tmocks.NewTelegramProvider(t)
	tClient.On("Send", mock.Anything, mock.Anything).Return(nil)

	router := app.Router(app.NewHandler(statsClient, tClient))

	// Act & Assert
	req, err := http.NewRequest("GET", "/stats", nil)
	require.Nil(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file provided")
	}

	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	// Arrange
	mockStats := stats.DailyStats{
		Data: stats.StatsData{
			Users: 1,
		},
	}

	statsClient := smocks.NewStatsProvider(t)
	statsClient.On("GetStats").Return(mockStats, nil)

	tClient := tmocks.NewTelegramProvider(t)
	tClient.On(
		"Send",
		mock.MatchedBy(func(msg string) bool {
			return msg != ""
		}),
		fmt.Sprintf("Users: %d", mockStats.Data.Users),
	).Return(nil)

	h := app.NewHandler(statsClient, tClient)

	svr := httptest.NewServer(http.HandlerFunc(h.Get))
	defer svr.Close()

	// Act
	r, err := http.Get(svr.URL)

	// Assert
	require.Nil(t, err)
	assert.Equal(t, http.StatusOK, r.StatusCode)
}
