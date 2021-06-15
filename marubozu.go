package golangcandles

import (
	"github.com/go-gota/gota/dataframe"
)

func Marubozu(df dataframe.DataFrame) (output int) {

	ohlc_1 := GenerateOhcl(df, 1)

	// BEAR
	if ohlc_1.Open > ohlc_1.Close {
		if ohlc_1.High == ohlc_1.Open && ohlc_1.Low == ohlc_1.Close {
			output = -100
		}
	}

	// BULL
	if ohlc_1.Open < ohlc_1.Close {
		if ohlc_1.Low == ohlc_1.Open && ohlc_1.High == ohlc_1.Close {
			output = 100
		}
	}

	return
}
