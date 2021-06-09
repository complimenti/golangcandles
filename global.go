package golangcandles

import "github.com/go-gota/gota/dataframe"

type Ohcl struct {
	CreatedAt string
	Symbol    string
	Open      float64
	High      float64
	Close     float64
	Low       float64
}

func GenerateOhcl(df dataframe.DataFrame, row_number int) (ohcl Ohcl) {
	len_df := df.Col("Open").Len()
	row := df.Subset([]int{len_df - row_number})
	ohlc := Ohcl{}
	ohlc.Open = row.Col("Open").Val(0).(float64)
	ohlc.High = row.Col("High").Val(0).(float64)
	ohlc.Close = row.Col("Close").Val(0).(float64)
	ohlc.Low = row.Col("Low").Val(0).(float64)
	return ohlc
}
