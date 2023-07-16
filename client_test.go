package iadzacksclientgo_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"iadzacksclientgo"
)

func TestClient_GetRating(t *testing.T) {
	t.Parallel()

	c := iadzacksclientgo.NewClient(time.Second * 5)

	ticker := "ZIM"

	rank, err := c.GetRating(context.Background(), ticker)
	require.NoError(t, err)
	require.Equal(t, ticker, rank.Ticker)
	require.Contains(t, []string{"1", "2", "3", "4", "5"}, rank.ZacksRank)
	require.Contains(t, []string{"A", "B", "C", "D", "E"}, rank.ZacksRankText)
}

func TestClient_GetRatings(t *testing.T) {
	t.Parallel()

	c := iadzacksclientgo.NewClient(time.Second * 5)

	tickers := []string{"ZIM", "BAT", "BTI"}

	ranks, err := c.GetRatings(context.Background(), tickers)
	require.NoError(t, err)
	require.Len(t, ranks, 2)

	for _, rank := range ranks {
		require.Contains(t, tickers, rank.Ticker)

		require.Contains(t, []string{"1", "2", "3", "4", "5"}, rank.ZacksRank, rank.Ticker)
		require.Contains(t, []string{"Strong Buy", "Buy", "Hold", "Sell", "Strong Sell"}, rank.ZacksRankText, rank.Ticker)
	}
}
