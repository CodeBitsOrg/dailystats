package app

import (
	"fmt"
	"github.com/CodeBitsOrg/dailystats/stats"
	"github.com/CodeBitsOrg/dailystats/telegram"
	"net/http"
	"os"
)

type Handler struct {
	statsClient stats.StatsProvider
	tClient     telegram.TelegramProvider
}

func NewHandler(sc stats.StatsProvider, tc telegram.TelegramProvider) *Handler {
	return &Handler{
		statsClient: sc,
		tClient:     tc,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	stats, err := h.statsClient.GetStats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.tClient.Send(
		os.Getenv("BOT_CHAT_ID"),
		fmt.Sprintf("Users: %d", stats.Data.Users),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
