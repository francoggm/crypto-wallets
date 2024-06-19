package tickers

import (
	"context"
	"encoding/json"
	"github/francoggm/crypto-wallets/internal/models"
	"github/francoggm/crypto-wallets/pkg/utils"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3/log"
)

type TickerData struct {
	SearchName   string `json:"id"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	MarketCapUSD string `json:"marketCapUsd"`
	VolumeUSD    string `json:"volumeUsd24Hr"`
	PriceUSD     string `json:"priceUsd"`
}

type TickerDataList struct {
	Data []TickerData `json:"data"`
}

func (tr *TickersRoutine) GetTickersData() {
	for {
		ctx := context.Background()

		assets, err := tr.Repository.GetAllAssets(ctx)
		if err != nil {
			log.Error(err)
			time.Sleep(30 * time.Second)

			continue
		}

		assetsTickers, err := getAssetsTickers(tr.TickersURL, assets)
		if err != nil {
			log.Error(err)
			time.Sleep(30 * time.Second)

			continue
		}

		for _, at := range assetsTickers {
			go func(assetTicker *models.AssetTicker) {
				if err := tr.Repository.SaveAssetTicker(ctx, assetTicker); err != nil {
					log.Error(err)
				}
			}(at)
		}

		time.Sleep(tr.TickersInterval)
	}
}

func getAssetsTickers(url string, assets []*models.Asset) ([]*models.AssetTicker, error) {
	tl, err := makeTickersRequest(url, utils.GetAssetsListQueryParam(assets))
	if err != nil {
		return nil, err
	}

	assetsTickers := make([]*models.AssetTicker, 0, len(tl.Data))

	for _, ticker := range tl.Data {
		asset := utils.GetAssetFromAssets(assets, ticker.SearchName, ticker.Name, ticker.Symbol)
		if asset == nil {
			continue
		}

		assetTicker := new(models.AssetTicker)
		assetTicker.Asset = *asset

		marketCapUSD, err := strconv.ParseFloat(ticker.MarketCapUSD, 64)
		if err != nil {
			continue
		}

		volumeUSD, err := strconv.ParseFloat(ticker.VolumeUSD, 64)
		if err != nil {
			continue
		}

		priceUSD, err := strconv.ParseFloat(ticker.PriceUSD, 64)
		if err != nil {
			continue
		}

		assetTicker.MarketCapUSD = marketCapUSD
		assetTicker.VolumeUSD = volumeUSD
		assetTicker.PriceUSD = priceUSD

		assetsTickers = append(assetsTickers, assetTicker)
	}

	return assetsTickers, nil
}

func makeTickersRequest(url, assetsQueryParam string) (*TickerDataList, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("ids", assetsQueryParam)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tl *TickerDataList
	if err := json.Unmarshal(responseBody, &tl); err != nil {
		return nil, err
	}

	return tl, nil
}
