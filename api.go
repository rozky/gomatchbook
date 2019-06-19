package gomatchbook

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type Session struct {
	credentials Credentials
	token       SessionToken
	httpClient  *http.Client
}

func (credentials Credentials) Login() (*Session, error) {

	if credentials.Username == "" {
		return nil, errors.New("username is empty")
	}
	if credentials.Password == "" {
		return nil, errors.New("password is empty")
	}

	session := new(Session)
	session.credentials = credentials
	session.httpClient = createHttpClient()

	body, err := session.login()

	if err != nil {
		return nil, err
	}

	session.token = body.SessionToken

	return session, nil
}

func (s *Session) Logout() error {
	_, err := s.doRequest(SessionURL, http.MethodDelete, strings.NewReader(""))
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) Balance() (*Balance, error) {
	resp, err := s.doRequest(BalanceURL, http.MethodGet, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	return parseBalance(resp)
}

func (s *Session) SettledBets(filter SettledBetsFilter) (*SettledBets, error) {
	resp, err := s.doRequest(SettledBetsURL(filter), http.MethodGet, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	return parseSettledBets(resp)
}

func (s *Session) SubmitOffers(body SubmitOffersPayload) (*SubmitOffersResult, error) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := s.doRequest(OffersURL, http.MethodPost, strings.NewReader(string(bodyJson)))

	return parseSubmitOffersResult(resp)
}

func (s *Session) Offers(filter OffersFilter) (*Offers, error) {
	resp, err := s.doRequest(FilteredOffersURL(filter), http.MethodGet, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	return parseOffers(resp)
}

func (s *Session) Events(filter EventsFilter) (*EventsResult, error) {
	resp, err := s.doRequest(EventsURL(filter), http.MethodGet, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	return parseEventsResult(resp)
}

func (s *Session) Sports(filter SportsFilter) (*SportsResult, error) {
	resp, err := s.doRequest(SportsURL(filter), http.MethodGet, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	return parseSportsResult(resp)
}

func (s *Session) doRequest(url string, method string, body *strings.Reader) ([]byte, error) {

	fmt.Println("calling: " + url)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	if s.token != "" {
		req.Header.Set("session-token", string(s.token))
	}

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Session) login() (*LoginResponse, error) {
	bodyJson, err := json.Marshal(s.credentials)
	if err != nil {
		return nil, err
	}

	resp, err := s.doRequest(SessionURL, http.MethodPost, strings.NewReader(string(bodyJson)))
	if err != nil {
		return nil, err
	}

	return parseLogin(resp)
}

func parseLogin(data []byte) (*LoginResponse, error) {
	r := &LoginResponse{}
	return r, json.Unmarshal(data, r)
}

func parseBalance(data []byte) (*Balance, error) {
	r := &Balance{}
	return r, json.Unmarshal(data, r)
}

func parseSettledBets(data []byte) (*SettledBets, error) {
	r := &SettledBets{}
	return r, json.Unmarshal(data, r)
}

func parseOffers(data []byte) (*Offers, error) {
	r := &Offers{}
	return r, json.Unmarshal(data, r)
}

func parseEventsResult(data []byte) (*EventsResult, error) {
	r := &EventsResult{}
	return r, json.Unmarshal(data, r)
}

func parseSportsResult(data []byte) (*SportsResult, error) {
	r := &SportsResult{}
	return r, json.Unmarshal(data, r)
}

func parseSubmitOffersResult(data []byte) (*SubmitOffersResult, error) {
	r := &SubmitOffersResult{}
	return r, json.Unmarshal(data, r)
}

func createHttpClient() (*http.Client){
	ssl := &tls.Config{
		InsecureSkipVerify: true,
	}
	ssl.Rand = rand.Reader

	return &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(time.Second*3))
			},
			TLSClientConfig: ssl,
		},
	}
}
