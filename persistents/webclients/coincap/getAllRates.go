package coincap

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (c coincap) GetAllRates(ctx context.Context) (dao.GetAllRates, error) {
	u, err := url.Parse(c.basePath)
	if err != nil {
		c.log.Error(ctx, GetAllRatesLogMessage).Err(err).Send()
		return dao.GetAllRates{}, err
	}

	u = u.JoinPath(RatesCoincapPath)

	cl := http.Client{}

	resp, err := cl.Get(u.String())
	if err != nil {
		c.log.Error(ctx, GetAllRatesLogMessage).Err(err).Send()
		return dao.GetAllRates{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.log.Error(ctx, GetAllRatesLogMessage).Err(err).Send()
		return dao.GetAllRates{}, err
	}

	var data dao.GetAllRates
	if err := jsoniter.Unmarshal(body, &data); err != nil {
		return dao.GetAllRates{}, err
	}

	return data, nil
}
