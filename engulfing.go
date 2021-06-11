package golangcandles

import (
	"github.com/go-gota/gota/dataframe"
)

func Engulfing(df dataframe.DataFrame) (output int) {

	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_1 := GenerateOhcl(df, 1)

	// BEAR
	if ohlc_2.Open < ohlc_2.Close {
		if ohlc_1.Open > ohlc_2.Close {
			if ohlc_1.Close < ohlc_2.Open {
				output = -100
			}
		}
	}

	// BULL
	if ohlc_2.Open > ohlc_2.Close {
		if ohlc_1.Open < ohlc_2.Close {
			if ohlc_1.Close > ohlc_2.Open {
				output = 100
			}
		}
	}

	return
}
