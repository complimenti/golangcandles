# GOLANGCANDLES

The simplest candles pattern recognition library written purely in Golang-

## DISCLAIMER

This library has been developed by highly unexperinced traders, and it contains bugs.

Moreover, it is full of if-else statements.

On the other hand, it is extremely fast and easy to understand.

Any function is tested in both bull/bear cases.

## USAGE

This library uses golang dataframe coming from [gota](https://github.com/go-gota/gota).

The goal is to pass a dataframe containing OHLC data, and it returns wheater a bullish or bearish pattern has been identified:

- 100_: BULL
- -100: BEAR
- 0: NULL

## TEST

Run all tests:


```shell
go test ./...
```

Run a single test:

```shell
go test -v -run TestThreeLineStrike
```
