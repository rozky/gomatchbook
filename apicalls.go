package gomatchbook

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/now"
)

func (credentials *Credentials) UnsafeLogin() (*Session) {
	session, err := credentials.Login()
	if err != nil {
		panic(err)
	}

	return session
}

func (session *Session) UnsafeLogout() {
	err := session.Logout()
	if err != nil {
		panic(err)
	}
}

func (session *Session) LogBalance()  {
	result, err := session.Balance()
	if err != nil {
		panic(err)
	}

	spew.Dump(result)
}

func (session *Session) LogSettledBets(monthsDelta int)  {
	filter := SettledBetsFilter{
		PerPage: 1000,
		After: now.BeginningOfMonth().AddDate(0, monthsDelta, 0).Format(TimeFormat),
		Before: now.EndOfMonth().AddDate(0, monthsDelta, 0).Format(TimeFormat),
	}
	result, err := session.SettledBets(filter)
	if err != nil {
		panic(err)
	}

	spew.Dump(result)
}

func (session *Session) GetSettledBetsGroupedByMonth(numberOfMonths int) map[string]*SettledBets {
	var result = make(map[string]*SettledBets)

	for i := 0; i < numberOfMonths; i++ {
		bets, month := session.GetMonthlySettledBets(-1 * i)
		result[month] = bets
	}

	return result
}

func (session *Session) GetMonthlySettledBets(monthsDelta int) (bets *SettledBets, month string) {
	beginOfMonth := now.BeginningOfMonth().AddDate(0, monthsDelta, 0)

	filter := SettledBetsFilter{
		PerPage: 1000,
		After:   beginOfMonth.Format(TimeFormat),
		Before:  now.New(beginOfMonth).EndOfMonth().Format(TimeFormat),
	}
	result, err := session.SettledBets(filter)
	if err != nil {
		panic(err)
	}

	return result, beginOfMonth.Format("2006-01-02")
}

/*
It's possible to get last 12 months
 */
func (session *Session) LogAllAvailableSettledBetsGroupedByMonth()  {
	for i := 0; i < 12; i++ {
		session.LogSettledBets(-1 * i)
	}
}

func (session *Session) LogOffers()  {
	filter := OffersFilter{
		PerPage: 1,
	}
	result, err := session.Offers(filter)
	if err != nil {
		panic(err)
	}

	spew.Dump(result)
}

func (session *Session) LogEvents()  {
	filter := EventsFilter{
		PerPage: 1,
		SportIds: TennisId,
		States: []eventState{EventStates.Open} ,
		PriceDepth: 3,
		IncludePrices: false,
		MinimumLiquidity:10,
		ExchangeType: BackLay,
		OddsType: Decimal,
	}
	result, err := session.Events(filter)
	if err != nil {
		panic(err)
	}

	spew.Dump(result)
}

func (session *Session) LogSports()  {
	filter := SportsFilter{
		PerPage:100,
	}

	result, err := session.Sports(filter)
	if err != nil {
		panic(err)
	}

	spew.Dump(result)
}