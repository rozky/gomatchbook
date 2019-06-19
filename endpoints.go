package gomatchbook

import "github.com/google/go-querystring/query"

const (
	apiURL         = "https://api.matchbook.com/edge/rest"
	SessionURL     = apiURL + "/security/session"
	BalanceURL     = apiURL + "/account/balance"
	settledBetsURL = apiURL + "/reports/v2/bets/settled"
	OffersURL      = apiURL + "/v2/offers"
	eventsURL      = apiURL + "/events"
	sportsURL      = apiURL + "/lookups/sports"
)

const (
	TennisId = "9"
	TimeFormat = "2006-01-02T15:04:05.000Z00:00"
)

func SportsURL(filter SportsFilter) string {
	v, err := query.Values(filter)
	if err != nil {
		panic(err)
	}

	return sportsURL + "?" + v.Encode();
}

func SettledBetsURL(filter SettledBetsFilter) string {
	v, err := query.Values(filter)
	if err != nil {
		panic(err)
	}

	return settledBetsURL + "?" + v.Encode();
}


func OfferURL(offerId int64, filter GetOfferFilter) string {
	v, err := query.Values(filter)
	if err != nil {
		panic(err)
	}

	return OffersURL + "/" + string(offerId) + "?" + v.Encode()
}

func FilteredOffersURL(filter OffersFilter) string {
	v, err := query.Values(filter)
	if err != nil {
		panic(err)
	}

	return OffersURL + "?" + v.Encode()
}

func EventsURL(filter EventsFilter) string {
	v, err := query.Values(filter)
	if err != nil {
		panic(err)
	}

	return eventsURL + "?" + v.Encode();
}

type OffersFilter struct {
	Offset             int    `url:"offset,omitempty"`
	PerPage            int    `url:"per-page,omitempty"`
	After              string `url:"after,omitempty"`
	Before             string `url:"before,omitempty"`
	EventIds           string `url:"event-ids,omitempty"`
	SportIds           string `url:"sport-ids,omitempty"`
	MarketIds          string `url:"market-ids,omitempty"`
	RunnerIds          string `url:"runner-ids,omitempty"`
	Side               string `url:"side,omitempty"`
	Status             string `url:"status,omitempty"`
	Interval           int    `url:"interval,omitempty"`
	IncludeEdits       bool   `url:"include-edits,omitempty"`
	CancellationReason bool   `url:"cancellation-reason,omitempty"`
}

type GetOfferFilter struct {
	IncludeEdits bool `url:"include-edits,omitempty"`
}

type SettledBetsFilter struct {
	Offset    int    `url:"offset,omitempty"`
	PerPage   int    `url:"per-page,omitempty"`
	After     string `url:"after,omitempty"`
	Before    string `url:"before,omitempty"`
	EventIds  string `url:"event-ids,omitempty"`
	SportIds  string `url:"sport-ids,omitempty"`
	MarketIds string `url:"market-ids,omitempty"`
}

type MarketsFilter struct {
	Offset                   int     `url:"offset,omitempty"`
	PerPage                  int     `url:"per-page,omitempty"`
	After                    string  `url:"after,omitempty"`
	Before                   string  `url:"before,omitempty"`
	CategoryIds              string  `url:"category-ids,omitempty"`
	SportIds                 string  `url:"sport-ids,omitempty"`
	Ids                      string  `url:"ids,omitempty"`
	States                   string  `url:"states,omitempty"`
	TagUrlNames              string  `url:"tag-url-names,omitempty"`
	ExchangeType             string  `url:"exchange-type,omitempty"`
	OddsType                 string  `url:"odds-type,omitempty"`
	IncludePrices            bool    `url:"include-prices,omitempty"`
	PriceDepth               int     `url:"price-depth,omitempty"`
	PriceMode                string  `url:"price-mode,omitempty"`
	Side                     string  `url:"side,omitempty"`
	Currency                 string  `url:"currency,omitempty"`
	MinimumLiquidity         float64 `url:"minimum-liquidity,omitempty"`
	IncludeEventParticipants bool    `url:"include-event-participants,omitempty"`
}

type SportsFilter struct {
	Offset  int    `url:"offset,omitempty"`
	PerPage int    `url:"per-page,omitempty"`
	Status  string `url:"status,omitempty"`
	Order   string `url:"order,omitempty"`
}

type EventsFilter struct {
	Offset                   int          `url:"offset,omitempty"`
	PerPage                  int          `url:"per-page,omitempty"`
	After                    string       `url:"after,omitempty"`
	Before                   string       `url:"before,omitempty"`
	CategoryIds              string       `url:"category-ids,omitempty"`
	SportIds                 string       `url:"sport-ids,omitempty"`
	Ids                      string       `url:"ids,omitempty"`
	States                   []eventState `url:"states,omitempty,comma"`
	TagUrlNames              string       `url:"tag-url-names,omitempty"`
	ExchangeType             string       `url:"exchange-type,omitempty"`
	OddsType                 string       `url:"odds-type,omitempty"`
	IncludePrices            bool         `url:"include-prices,omitempty"`
	PriceDepth               int          `url:"price-depth,omitempty"`
	PriceMode                string       `url:"price-mode,omitempty"`
	Side                     string       `url:"side,omitempty"`
	Currency                 string       `url:"currency,omitempty"`
	MinimumLiquidity         float64      `url:"minimum-liquidity,omitempty"`
	IncludeEventParticipants bool         `url:"include-event-participants,omitempty"`
}

type EventFilter struct {
	ExchangeType             string  `url:"exchange-type,omitempty"`
	OddsType                 string  `url:"odds-type,omitempty"`
	IncludePrices            bool    `url:"include-prices,omitempty"`
	PriceDepth               int     `url:"price-depth,omitempty"`
	PriceMode                string  `url:"price-mode,omitempty"`
	Side                     string  `url:"side,omitempty"`
	Currency                 string  `url:"currency,omitempty"`
	MinimumLiquidity         float64 `url:"minimum-liquidity,omitempty"`
	IncludeEventParticipants bool    `url:"include-event-participants,omitempty"`
}

var EventStates = eventStates{
	Open:      "open",
	Suspended: "suspended",
	Closed:    "closed",
	Graded:    "graded",
}

type eventState string

type eventStates struct {
	Open      eventState
	Suspended eventState
	Closed    eventState
	Graded    eventState
}
