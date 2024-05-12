package dto

import "github.com/naufalfmm/cryptocurrency-price-api/utils/listener/ws"

type SyncCoinPriceRequest struct {
	CoinPriceMap map[string]string
}

func (req *SyncCoinPriceRequest) FromWsContext(wsc ws.Context) error {
	if err := wsc.Bind(&req.CoinPriceMap); err != nil {
		return err
	}

	return nil
}

func (req SyncCoinPriceRequest) ToCoincapIDs() []string {
	coincapIDs := make([]string, len(req.CoinPriceMap))
	i := 0
	for coincapID := range req.CoinPriceMap {
		coincapIDs[i] = coincapID
		i++
	}

	return coincapIDs
}
