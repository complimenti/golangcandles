# GOLANGCANDLES

The simplest candles pattern recognition library written purely in Golang.

## DISCLAIMER

This library has been developed by highly unexperienced traders, and it contains bugs.

Moreover it is code full of if-else statements.

On the other hand, it is extremely fast and easy to understand.

Any function is tested in both bull/bear cases.

Any support is more than welcome!

## USAGE

This library uses golang dataframe coming from [gota](https://github.com/go-gota/gota).

The goal is to pass a dataframe containing OHLC data, getting as return wheater a bullish or bearish pattern is identified:

- 100: BULL
- -100: BEAR
- 0: NULL

Please note, the library identifies pattern only about the last element of the dataframe and not the previous one!

### EXAMPLE

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
  result := AbandonedBaby(df)
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
