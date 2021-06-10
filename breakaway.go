package golangcandles

import (
	"github.com/go-gota/gota/dataframe"
)

func BreakAway(df dataframe.DataFrame) (output int) {

	ohlc_5 := GenerateOhcl(df, 5)
	ohlc_4 := GenerateOhcl(df, 4)
	ohlc_3 := GenerateOhcl(df, 3)
	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_1 := GenerateOhcl(df, 1)

	// BEAR
	if ohlc_5.Open < ohlc_5.Close {
		if ohlc_4.Open < ohlc_4.Close && ohlc_5.Close < ohlc_4.Open {
			if ohlc_4.Close < ohlc_3.Close && ohlc_3.Close < ohlc_2.Close {
				if ohlc_1.Open > ohlc_1.Close && ohlc_1.Close > ohlc_5.Close {
					output = -100
				}
			}
		}
	}

	// BULL
	if ohlc_5.Open > ohlc_5.Close {
		if ohlc_4.Open > ohlc_4.Close && ohlc_5.Close > ohlc_4.Open {
			if ohlc_4.Close > ohlc_3.Close && ohlc_3.Close > ohlc_2.Close {
				if ohlc_1.Open < ohlc_1.Close && ohlc_1.Close < ohlc_5.Close {
					output = 100
				}
			}
		}
	}

	return
}
