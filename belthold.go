package golangcandles

import (
	"github.com/go-gota/gota/dataframe"
)

func BeltHold(df dataframe.DataFrame) (output int) {

	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_1 := GenerateOhcl(df, 1)

	// BEAR
	if ohlc_2.Close > ohlc_2.Open {
		if ohlc_2.High < ohlc_1.Open {
			if ohlc_1.Open == ohlc_1.High {
				if ohlc_1.Close < ohlc_1.Open {
					output = -100
				}
			}
		}
	}

	// BULL
	if ohlc_2.Close < ohlc_2.Open {
		if ohlc_2.Low > ohlc_1.Open {
			if ohlc_1.Open == ohlc_1.Low {
				if ohlc_1.Close > ohlc_1.Open {
					output = 100
				}
			}
		}
	}

	return
}
