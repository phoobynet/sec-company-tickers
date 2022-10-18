package tickers

import "testing"

func TestGetWithDefaultLink(t *testing.T) {
	tickers, err := Get(nil)

	if err != nil {
		t.Error("Error: ", err)
	}

	if tickers == nil {
		t.Error("tickers was unexpectedly nil")
	}

	if len(tickers) == 0 {
		t.Error("tickers was unexpectedly empty")
	}

	t.Log("tickers length: ", len(tickers))

	t.Log("", tickers[0])
}
