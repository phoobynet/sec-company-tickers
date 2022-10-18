package tickers

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

type Ticker struct {
	CIK      int
	Symbol   string
	Name     string
	Exchange string
}

const Link = "https://www.sec.gov/files/company_tickers_exchange.json"

func Get(link *string) ([]Ticker, error) {
	var l string
	if link == nil {
		l = Link
	}

	response, httpError := resty.New().R().
		Get(l)

	if httpError != nil {
		log.Fatal(httpError)
	}

	if response.IsError() {
		return nil, fmt.Errorf("Failed to get data from SEC: status was %d\n%s", response.StatusCode(), string(response.Body()))
	}

	var r struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	}

	err := json.Unmarshal(response.Body(), &r)

	if err != nil {
		return nil, err
	}

	var tickers []Ticker

	for _, item := range r.Data {
		cik := item[0].(float64)
		name := item[1].(string)
		symbol := item[2].(string)
		exchange := item[3].(string)

		tickers = append(tickers, Ticker{
			CIK:      int(cik),
			Symbol:   symbol,
			Name:     name,
			Exchange: exchange,
		})
	}

	return tickers, nil
}
