# SEC Company Tickers

Pulls ticker, name and exchange information from the [SEC website](https://www.sec.gov/os/accessing-edgar-data).

## Installation
```bash
go get github.com/phoobynet/sec-company-tickers
```

## Usage

```go
package main
import (
	"fmt"
	"github.com/phoobynet/sec-company-tickers"
)

func main () {
	// it is possible to override the URL, or pass nil to use the default
	companyTickers, err := tickers.Get(nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Tickers found: %d", len(companyTickers))
}
```

```text
{320193 AAPL Apple Inc. Nasdaq}
```