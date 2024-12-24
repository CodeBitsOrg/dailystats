package stats

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type StatsProvider interface {
	GetStats() (DailyStats, error)
}

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type StatsClient struct {
	HttpClient HttpClient
}

type DailyStats struct {
	Data StatsData `json:"data"`
}

type StatsData struct {
	Users            int `json:"users"`
	Comments         int `json:"comments"`
	NewsSubscription int `json:"news_subscription"`
	MonthlyPlan      int `json:"monthly_plan"`
	YearlyPlan       int `json:"yearly_plan"`
}

func NewStatsClient(httpClient HttpClient) *StatsClient {
	return &StatsClient{
		HttpClient: httpClient,
	}
}

func (sc *StatsClient) FetchData(url string) ([]byte, error) {
	response, err := sc.HttpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting stats: %s", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	return body, nil
}

func (sc *StatsClient) GetStats() (DailyStats, error) {
	var ds DailyStats

	body, err := sc.FetchData(os.Getenv("STATS_API_URL"))
	if err != nil {
		return ds, fmt.Errorf("error fetching stats: %s", err)
	}

	err = json.Unmarshal(body, &ds)
	if err != nil {
		return ds, fmt.Errorf("error unmarshalling stats response: %s", err)
	}

	return ds, nil
}
