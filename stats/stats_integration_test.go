//go:build integration

package stats_test

import (
	"github.com/CodeBitsOrg/dailystats/stats"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetStats_ItChecksRealEndpoint(t *testing.T) {
	// Arrange
	statsClient := stats.NewStatsClient(http.DefaultClient)

	// Act
	_, err := statsClient.GetStats()

	// Assert
	assert.NoErrorf(t, err, "Error while getting stats: %v", err)
}
