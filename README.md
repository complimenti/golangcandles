# GolangCandles

The simplest candles pattern recognition library written purely in Golang.

## Disclaimer

This library has been developed by highly unexperienced traders, it might contain bugs, use it at your own risk.

As strenght it is extremely fast and easy to understand.

Any support is more than welcome!

## Dependencies

This library uses golang dataframe coming from [gota](https://github.com/go-gota/gota).

## Use

Pass a dataframe containing OHLC data and get as return wheater a bullish or bearish pattern is identified:

- 100: BULL
- -100: BEAR
- 0: NULL

**Please note, the library identifies pattern only about the last element of the dataframe and not the previous one!**

## Examples

#### Simple script

```golang
package main

import (
  "fmt"
  "github.com/go-gota/gota/dataframe"
  "github.com/go-gota/gota/series"
  "github.com/complimenti/golangcandles"
)

func main() {
  df := dataframe.New(
    series.New([]float64{90, 125, 110}, series.Float, "Open"),
    series.New([]float64{105, 135, 115}, series.Float, "High"),
    series.New([]float64{100, 130, 95}, series.Float, "Close"),
    series.New([]float64{85, 120, 92}, series.Float, "Low"),
  )
  result := golangcandles.AbandonedBaby(df)
  fmt.Println(result)
}
```

#### Integrate with Postgresql
```golang
package main

import (
    "database/sql"
    "fmt"
    "golangcandles"
    
    "github.com/go-gota/gota/dataframe"
    "github.com/go-gota/gota/series"
    "github.com/lib/pq"
    _ "github.com/lib/pq"
)

var dbURI = `
    user=YOURUSER 
    password=YOURPASSWORD
    dbname=YOURDB
    host=YOURHOST`

func main() {

    // INITIALIZE DB
    db, err := sql.Open("postgres", dbURI)
    defer db.Close()
    if err != nil {
    	fmt.Println(err)
    }

    var open, close, high, low []float64

    query := `
    	SELECT
    		ARRAY_AGG(open),
    		ARRAY_AGG(close),
    		ARRAY_AGG(high),
    		ARRAY_AGG(low)
    	FROM ohlc
    	WHERE exchange = 'binance'
    	AND symbol = 'ETHBTC'
    	AND createdat > (current_timestamp - interval '10 day');`

    err = db.QueryRow(query).Scan(
    	   pq.Array(&open),
    	   pq.Array(&close),
    	   pq.Array(&high),
    	   pq.Array(&low),
    )
    if err != nil {
    	   fmt.Println(err)
    }
    
    df := dataframe.New(
    	   series.New(open, series.Float, "Open"),
    	   series.New(close, series.Float, "Close"),
    	   series.New(high, series.Float, "High"),
    	   series.New(low, series.Float, "Low"),
    )
    result := golangcandles.AbandonedBaby(df)
    fmt.Println(result)
}
```

## TEST

Run all tests:

```shell
go test ./...
```

Run a single test:

```shell
go test -v -run TestThreeLineStrike
```
