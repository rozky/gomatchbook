package gomatchbook

import (
	"time"
)

const (
	BackLay = "back-lay"
	Decimal = "DECIMAL"
)

type Username string

type Password string

type SessionToken string

type Credentials struct {
	Username Username `json:"username"`
	Password Password `json:"password"`
}

type LoginResponse struct {
	SessionToken SessionToken   `json:"session-token"`
	UserId       int            `json:"user-id"`
	Role         string         `json:"role"`
	Account      AccountDetails `json:"account"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone-number"`
	Address      AddressDetails `json:"address"`
}

type AccountDetails struct {
	Id       int         `json:"id"`
	Username string      `json:"username"`
	Name     NameDetails `json:"name"`
}

type NameDetails struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

type BaseResponse struct {
	Total   int `json:"total"`
	PerPage int `json:"per-page"`
	Offset  int `json:"offset"`
}

type AddressDetails struct {
}

type SportsResult struct {
	BaseResponse
	Sports []Sport `json:"sports"`
}

type Sport struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Id   int    `json:"id"`
}

type EventsResult struct {
	BaseResponse
	Events []Event `json:"events"`
}

type Event struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	SportId          int       `json:"sport-id"`
	Start            time.Time `json:"start"`
	InRunningFlag    bool      `json:"in-running-flag"`
	AllowLiveBetting bool      `json:"allow-live-betting"`
	CategoryId       []int     `json:"category-id"`
	Status           string    `json:"status"`
	Volume           float32   `json:"volume"`
	Markets          []Market  `json:"markets"`
	MetaTags         []MetaTag `json:"meta-tags"`
}

type Market struct {
	Live             bool      `json:"live"`
	EventId          int       `json:"event-id"`
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Runners          []Runner  `json:"runners"`
	Start            time.Time `json:"start"`
	InRunningFlag    bool      `json:"in-running-flag"`
	AllowLiveBetting bool      `json:"allow-live-betting"`
	Status           string    `json:"status"`
	MarketType       string    `json:"market-type"`
	Type             string    `json:"type"`
	Volume           float32   `json:"volume"`
	BackOverround    float32   `json:"back-overround"`
	LayOverround     float32   `json:"lay-overround"`
}

type Runner struct {
	Prices             []PriceDetail `json:"prices"`
	EventId            int           `json:"event-id"`
	Id                 int           `json:"id"`
	MarketId           int           `json:"market-id"`
	Name               string        `json:"name"`
	Status             string        `json:"status"`
	Volume             float32       `json:"volume"`
	EventParticipantId int           `json:"event-participant-id"`
}

type PriceDetail struct {
	AvailableAmount float32 `json:"available-amount"`
	Currency        string  `json:"currency"`
	OddsType        string  `json:"odds-type"`
	Odds            float32 `json:"odds"`
	DecimalOdds     float32 `json:"decimal-odds"`
	Side            string  `json:"side"`
	ExchangeType    string  `json:"exchange-type"`
}

type MetaTag struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	UrlName string `json:"url-names"`
}

type Balance struct {
	Balance           float32 `json:"balance"`
	Exposure          float32 `json:"exposure"`
	FreeFunds         float32 `json:"free-funds"`
	CommissionReserve float32 `json:"commission-reserve"`
}

type SettledBets struct {
	Offset           int             `json:"offset"`
	PerPage          int             `json:"per-page"`
	Total            int             `json:"total"`
	Currency         string          `json:"currency"`
	Language         string          `json:"language"`
	OddsType         string          `json:"odds-type"`
	ProfitAndLoss    float32         `json:"profit-and-loss"`
	Commission       float32         `json:"commission"`
	NetProfitAndLoss float32         `json:"net-profit-and-loss"`
	Markets          []SettledMarket `json:"markets"`
}

type SettledMarket struct {
	Id               int64              `json:"id"`
	EventId          int64              `json:"event-id"`
	EventName        string             `json:"event-name"`
	SportId          int64              `json:"sport-id"`
	Name             string             `json:"name"`
	StartTime        string             `json:"start-time"`
	SettledTime      string             `json:"settled-time"`
	CommissionType   string             `json:"commission-type"`
	ProfitAndLoss    float32            `json:"profit-and-loss"`
	Commission       float32            `json:"commission"`
	NetProfitAndLoss float32            `json:"net-profit-and-loss"`
	Selections       []SettledSelection `json:"selections"`
}

type SettledSelection struct {
	RunnerId         int64                 `json:"runner-id"`
	RunnerName       string                `json:"runner-name"`
	Side             string                `json:"side"`
	Odds             float32               `json:"odds"`
	Stake            float32               `json:"stake"`
	ProfitAndLoss    float32               `json:"profit-and-loss"`
	Commission       float32               `json:"commission"`
	NetProfitAndLoss float32               `json:"net-profit-and-loss"`
	Bets             []SettledSelectionBet `json:"bets"`
}

type SettledSelectionBet struct {
	Id               int64   `json:"id"`
	OfferId          int64   `json:"offer-id"`
	Result           string  `json:"result"`
	Odds             float32 `json:"odds"`
	Stake            float32 `json:"stake"`
	Adjusted         bool    `json:"adjusted"`
	Originator       bool    `json:"originator"`
	InPlay           bool    `json:"in-play"`
	SubmittedTime    string  `json:"submitted-time"`
	MatchedTime      string  `json:"matched-time"`
	SettledTime      string  `json:"settled-time"`
	ProfitAndLoss    float32 `json:"profit-and-loss"`
	Commission       float32 `json:"commission"`
	CommissionType   string  `json:"commission-type"`
	NetProfitAndLoss float32 `json:"net-profit-and-loss"`
}

type Offers struct {
	Currency     string  `json:"currency"`
	Language     string  `json:"language"`
	OddsType     string  `json:"odds-type"`
	ExchangeType string  `json:"exchange-type"`
	Offers       []Offer `json:"offers"`
}

// todo
type Offer struct {
	Id        int64  `json:"id"`
	EventId   int64  `json:"event-id"`
	EventName string `json:"event-name"`
}

type SubmitOffersPayload struct {
	OddsType     string        `json:"odds-type"`
	ExchangeType string        `json:"exchange-type"`
	Offers       []SubmitOffer `json:"offers"`
}

type SubmitOffer struct {
	RunnerId   int64   `json:"runner-id"`
	Side       string  `json:"side"`
	Odds       float32 `json:"odds"`
	Stake      float32 `json:"stake"`
	KeepInPlay bool    `json:"keep-in-play"`
}

type SubmitOffersResult struct {
	Language     string `json:"language"`
	Currency     string `json:"currency"`
	ExchangeType string `json:"exchange-type"`
	OddsType     string `json:"odds-type"`
	Offers       []struct {
		Id                       int64   `json:"id"`
		EventId                  int64   `json:"event-id"`
		EventName                string  `json:"event-name"`
		MarketId                 int64   `json:"market-id"`
		MarketType               string  `json:"market-type"`
		RunnerId                 int64   `json:"runner-id"`
		RunnerName               string  `json:"runner-name"`
		ExchangeType             string  `json:"exchange-type"`
		Side                     string  `json:"side"`
		Odds                     float32 `json:"odds"`
		OddsType                 string  `json:"odds-type"`
		DecimalOdds              float32 `json:"decimal-odds"`
		Stake                    float32 `json:"stake"`
		Remaining                float32 `json:"remaining"`
		PotentialProfit          float32 `json:"potential-profit"`
		RemainingPotentialProfit float32 `json:"remaining-potential-profit"`
		CommissionType           string  `json:"commission-type"`
		Currency                 string  `json:"currency"`
		CreatedAt                string  `json:"created-at"`
		Status                   string  `json:"status"`
		InPlay                   bool    `json:"in-play"`
		MatchedBets              []struct {
			Id      int64 `json:"id"`
			OfferId int64 `json:"offer-id"`
			// todo
		} `json:"matched-bets"`
	} `json:"offers"`
}
