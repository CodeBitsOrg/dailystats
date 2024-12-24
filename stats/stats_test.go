package stats_test

import (
	"bytes"
	"github.com/CodeBitsOrg/dailystats/stats"
	mocks "github.com/CodeBitsOrg/dailystats/stats/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestFetchData(t *testing.T) {
	// Arrange
	mockHTTPClient := mocks.NewHttpClient(t)

	mockResponse := expectGetIsCalled(mockHTTPClient)
	statsClient := stats.NewStatsClient(mockHTTPClient)

	// Act
	body, err := statsClient.FetchData("http://localhost:8080/stats")

	// Assert
	assert.NoErrorf(t, err, "Error fetching data: %s", err)
	assert.Equalf(t, mockResponse, body, "Expected %s, got %s", mockResponse, body)
}

func TestGetStats(t *testing.T) {
	// Arrange
	mockHttpClient := mocks.NewHttpClient(t)
	expectGetIsCalled(mockHttpClient)

	statsClient := stats.NewStatsClient(mockHttpClient)

	// Act
	s, err := statsClient.GetStats()

	// Assert
	assert.NoErrorf(t, err, "Error fetching data: %s", err)

	user := 121
	assert.Equalf(t, user, s.Data.Users, "Expected %d users, got %d", user, s.Data.Users)
}

func expectGetIsCalled(mockHttpClient *mocks.HttpClient) []byte {
	mockResponse := `{"data":{"users":121, "comments": 260, "news_subscribers": 334, "monthly_plan": 40, "yearly_plan": 50}}`
	mockHttpClient.On("Get", "http://localhost:8080/stats").Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
	}, nil)

	return []byte(mockResponse)
}
