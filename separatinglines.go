package golangcandles

import (
	"math"

	"github.com/go-gota/gota/dataframe"
)

func SeparatingLines(df dataframe.DataFrame) (output int) {

	ohlc_1 := GenerateOhcl(df, 1)
	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_3 := GenerateOhcl(df, 3)

	// BEAR
	if ohlc_3.Open > ohlc_3.Close {
		if ohlc_2.Open < ohlc_2.Close {
			if ohlc_1.Open > ohlc_1.Close {
				if math.Abs((ohlc_2.Open-ohlc_1.Open)/ohlc_1.Open) < 0.1 {
					if ohlc_1.Open == ohlc_1.High {
						output = -100
					}
				}
			}
		}
	}

	// BULL
	if ohlc_3.Open < ohlc_3.Close {
		if ohlc_2.Open > ohlc_2.Close {
			if ohlc_1.Open < ohlc_1.Close {
				if math.Abs((ohlc_2.Open-ohlc_1.Open)/ohlc_1.Open) < 0.1 {
					if ohlc_1.Open == ohlc_1.Low {
						output = 100
					}
				}
			}
		}
	}

	return
}
