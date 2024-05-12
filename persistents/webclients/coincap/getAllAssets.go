package coincap

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
)

func (c coincap) GetAllAssets(ctx context.Context, req dto.AllAssetsCoincapRequest) (dao.AllAsset, error) {
	var data dao.AllAsset

	u, err := url.Parse(c.basePath)
	if err != nil {
		c.log.Error(ctx, GetAllAssetLogMessage).Err(err).Any(consts.ReqLogKey, req).Send()
		return dao.AllAsset{}, err
	}

	u = u.JoinPath(AssetCoincapPath)

	q := u.Query()
	if req.Search != "" {
		q.Set("search", req.Search)
	}

	if req.Limit != 0 {
		q.Set("limit", strconv.Itoa(req.Limit))
	}

	if req.IDs != "" {
		q.Set("ids", req.IDs)
	}

	u.RawQuery = q.Encode()

	httpReq, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		c.log.Error(ctx, GetAllAssetLogMessage).Err(err).Any(consts.ReqLogKey, req).Send()
		return dao.AllAsset{}, err
	}

	cl := http.Client{}

	resp, err := cl.Do(httpReq)
	if err != nil {
		c.log.Error(ctx, GetAllAssetLogMessage).Err(err).Any(consts.ReqLogKey, req).Send()
		return dao.AllAsset{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.log.Error(ctx, GetAllAssetLogMessage).Err(err).Any(consts.ReqLogKey, req).Send()
		return dao.AllAsset{}, err
	}

	if err := jsoniter.Unmarshal(body, &data); err != nil {
		return dao.AllAsset{}, err
	}

	return data, nil
}
