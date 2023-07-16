package iadzacksclientgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

//nolint:lll
const (
	baseURL = "https://quote-feed.zacks.com/"
)

type Ranks map[string]*Rank

//nolint:tagliatelle
type Rank struct {
	Source struct {
		Sungard struct {
			Bidasksize          string `json:"bidasksize"`
			DividendFreq        string `json:"dividend_freq"`
			PrevCloseDate       string `json:"prev_close_date"`
			Timestamp           string `json:"timestamp"`
			Exchange            string `json:"exchange"`
			Shares              string `json:"shares"`
			Volatility          string `json:"volatility"`
			ZacksRecommendation string `json:"zacks_recommendation"`
			PosSize             string `json:"pos_size"`
			Open                string `json:"open"`
			Yrlow               string `json:"yrlow"`
			Type                string `json:"type"`
			Yield               string `json:"yield"`
			MarketCap           string `json:"market_cap"`
			Ask                 string `json:"ask"`
			Dividend            string `json:"dividend"`
			DividendDate        string `json:"dividend_date"`
			Earnings            string `json:"earnings"`
			Close               string `json:"close"`
			DayLow              string `json:"day_low"`
			LastTradeDatetime   string `json:"last_trade_datetime"`
			Volume              string `json:"volume"`
			Yrhigh              string `json:"yrhigh"`
			DayHigh             string `json:"day_high"`
			Bid                 string `json:"bid"`
			Name                string `json:"name"`
			PeRatio             string `json:"pe_ratio"`
			Updated             string `json:"updated"`
		} `json:"sungard"`
		Bats struct {
			AskSize           string `json:"ask_size"`
			Routed            string `json:"routed"`
			LastTradeDatetime string `json:"last_trade_datetime"`
			Matched           string `json:"matched"`
			BidSize           string `json:"bid_size"`
			NetPctChange      string `json:"net_pct_change"`
			Updated           string `json:"updated"`
			EndMktDayPrice    string `json:"end_mkt_day_price"`
			AskPrice          string `json:"ask_price"`
			BidPrice          string `json:"bid_price"`
			Last              string `json:"last"`
			PreAfterUpdated   string `json:"pre_after_updated"`
			NetPriceChange    string `json:"net_price_change"`
			PreAfterPrice     string `json:"pre_after_price"`
			NetChange         string `json:"net_change"`
		} `json:"bats"`
		Pre struct {
			AfterPercentNetChange string `json:"after_percent_net_change"`
			AfterNetChange        string `json:"after_net_change"`
		} `json:"pre"`
	} `json:"source"`
	Exchange                   string `json:"exchange"`
	DividendYield              string `json:"dividend_yield"`
	Ticker                     string `json:"ticker"`
	Last                       string `json:"last"`
	TickerType                 string `json:"ticker_type"`
	ZacksRankText              string `json:"zacks_rank_text"`
	Volume                     string `json:"volume"`
	Updated                    string `json:"updated"`
	PercentNetChange           string `json:"percent_net_change"`
	ZacksRank                  string `json:"zacks_rank"`
	Name                       string `json:"name"`
	NetChange                  string `json:"net_change"`
	MarketTime                 string `json:"market_time"`
	PreviousClose              string `json:"previous_close"`
	SUNGARDBID                 string `json:"SUNGARD_BID"`
	SUNGARDYRLOW               string `json:"SUNGARD_YRLOW"`
	SUNGARDPERATIO             string `json:"SUNGARD_PE_RATIO"`
	SUNGARDDAYLOW              string `json:"SUNGARD_DAY_LOW"`
	SUNGARDMARKETCAP           string `json:"SUNGARD_MARKET_CAP"`
	FEEDNETCHANGE              string `json:"FEED_NET_CHANGE"`
	BATSPREAFTERUPDATED        string `json:"BATS_PRE_AFTER_UPDATED"`
	SUNGARDEARNINGS            string `json:"SUNGARD_EARNINGS"`
	SUNGARDVOLATILITY          string `json:"SUNGARD_VOLATILITY"`
	SUNGARDPREVCLOSEDATE       string `json:"SUNGARD_PREV_CLOSE_DATE"`
	BATSASKPRICE               string `json:"BATS_ASK_PRICE"`
	SUNGARDYRHIGH              string `json:"SUNGARD_YRHIGH"`
	SUNGARDDIVIDENDFREQ        string `json:"SUNGARD_DIVIDEND_FREQ"`
	SUNGARDVOLUME              string `json:"SUNGARD_VOLUME"`
	SUNGARDBIDASKSIZE          string `json:"SUNGARD_BIDASKSIZE"`
	SUNGARDYIELD               string `json:"SUNGARD_YIELD"`
	SUNGARDDAYHIGH             string `json:"SUNGARD_DAY_HIGH"`
	SUNGARDZACKSRECOMMENDATION string `json:"SUNGARD_ZACKS_RECOMMENDATION"`
	SUNGARDNAME                string `json:"SUNGARD_NAME"`
	SUNGARDTIMESTAMP           string `json:"SUNGARD_TIMESTAMP"`
	SUNGARDDIVIDENDDATE        string `json:"SUNGARD_DIVIDEND_DATE"`
	BATSBIDSIZE                string `json:"BATS_BID_SIZE"`
	BATSBIDPRICE               string `json:"BATS_BID_PRICE"`
	BATSLASTTRADEDATETIME      string `json:"BATS_LAST_TRADE_DATETIME"`
	FEEDVOLUME                 string `json:"FEED_VOLUME"`
	SUNGARDSHARES              string `json:"SUNGARD_SHARES"`
	SUNGARDDIVIDEND            string `json:"SUNGARD_DIVIDEND"`
	SUNGARDEXCHANGE            string `json:"SUNGARD_EXCHANGE"`
	SUNGARDTYPE                string `json:"SUNGARD_TYPE"`
	SUNGARDLASTTRADEDATETIME   string `json:"SUNGARD_LAST_TRADE_DATETIME"`
	SUNGARDUPDATED             string `json:"SUNGARD_UPDATED"`
	BATSROUTED                 string `json:"BATS_ROUTED"`
	BATSASKSIZE                string `json:"BATS_ASK_SIZE"`
	FEEDTICKER                 string `json:"FEED_TICKER"`
	SUNGARDPOSSIZE             string `json:"SUNGARD_POS_SIZE"`
	FEEDPERCENTNETCHANGE       string `json:"FEED_PERCENT_NET_CHANGE"`
	FEEDSOURCE                 string `json:"FEED_SOURCE"`
	BATSPREAFTERPRICE          string `json:"BATS_PRE_AFTER_PRICE"`
	BATSUPDATED                string `json:"BATS_UPDATED"`
	FEEDLAST                   string `json:"FEED_LAST"`
	SUNGARDASK                 string `json:"SUNGARD_ASK"`
	BATSMATCHED                string `json:"BATS_MATCHED"`
	FEEDUPDATED                string `json:"FEED_UPDATED"`
	SUNGARDOPEN                string `json:"SUNGARD_OPEN"`
	SUNGARDCLOSE               string `json:"SUNGARD_CLOSE"`
	PreAfterNetChange          string `json:"pre_after_net_change"`
	PreAfterPercentNetChange   string `json:"pre_after_percent_net_change"`
	CompanyShortName           string `json:"company_short_name"`
	TickerMarketStatus         string `json:"ticker_market_status"`
	ApShortName                string `json:"ap_short_name"`
	CompanyLogoURL             string `json:"company_logo_url"`
	ConfirmedReportingDate     string `json:"confirmed_reporting_date"`
	ExpectedReportingDate      string `json:"expected_reporting_date"`
	PeF1                       string `json:"pe_f1"`
	MarketStatus               string `json:"market_status"`
	PreviousCloseDate          string `json:"previous_close_date"`
	Error                      string `json:"error,omitempty"`
	Reason                     string `json:"reason,omitempty"`
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		timeout: timeout,
	}
}

type Client struct {
	timeout time.Duration
}

func (c *Client) GetRatings(ctx context.Context, ticker []string) (Ranks, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	body, err := c.Request(ctxWithTimeout, ticker)
	if err != nil {
		return nil, err
	}

	ranks := make(Ranks, 0)

	err = json.Unmarshal(body, &ranks)
	if err != nil {
		return nil, err
	}

	for id, rank := range ranks {
		if rank.Error == "true" {
			delete(ranks, id)
		}
	}

	return ranks, nil
}

func (c *Client) GetRating(ctx context.Context, ticker string) (*Rank, error) {
	ranks, err := c.GetRatings(ctx, []string{ticker})
	if err != nil {
		return nil, err
	}

	if item, exists := ranks[ticker]; exists {
		return item, nil
	}

	return nil, nil
}

func (c *Client) Request(ctx context.Context, tickers []string) ([]byte, error) {
	url := fmt.Sprintf("%sindex.php?t=%s", baseURL, strings.Join(tickers, ","))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
