package golangcandles

import (
	"github.com/go-gota/gota/dataframe"
)

func Kicking(df dataframe.DataFrame) (output int) {

	ohlc_2 := GenerateOhcl(df, 2)
	ohlc_1 := GenerateOhcl(df, 1)

	// BEAR
	if ohlc_2.Open < ohlc_2.Close {
		if ohlc_1.Open > ohlc_1.Close {
			if ohlc_2.Open > ohlc_1.Open {
				if ohlc_2.Open == ohlc_2.Low && ohlc_2.Close == ohlc_2.High {
					if ohlc_1.Open == ohlc_1.High && ohlc_1.Close == ohlc_1.Low {
						output = -100
					}
				}
			}
		}
	}

	// BULL
	if ohlc_2.Open > ohlc_2.Close {
		if ohlc_1.Open < ohlc_1.Close {
			if ohlc_2.Open < ohlc_1.Open {
				if ohlc_2.Open == ohlc_2.High && ohlc_2.Close == ohlc_2.Low {
					if ohlc_1.Open == ohlc_1.Low && ohlc_1.Close == ohlc_1.High {
						output = 100
					}
				}
			}
		}
	}

	return
}
