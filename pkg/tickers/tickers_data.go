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
	tc := time.NewTicker(tr.TickersInterval)
	defer tc.Stop()

	for range tc.C {
		tr.tickersRoutine()
	}
}

func (tr *TickersRoutine) tickersRoutine() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	assets, err := tr.Repository.GetAllAssets(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	assetsTickers, err := getAssetsTickers(ctx, tr.TickersURL, assets)
	if err != nil {
		log.Error(err)
		return
	}

	for _, at := range assetsTickers {
		go func(ctx context.Context, assetTicker models.AssetTicker) {
			if err := tr.Repository.SaveAssetTicker(ctx, &assetTicker); err != nil {
				log.Error(err)
			}
		}(ctx, *at)
	}
}

func getAssetsTickers(ctx context.Context, url string, assets []*models.Asset) ([]*models.AssetTicker, error) {
	tdl, err := makeTickersRequest(ctx, url, utils.GetAssetsListQueryParam(assets))
	if err != nil {
		return nil, err
	}

	assetsTickers := make([]*models.AssetTicker, 0, len(tdl.Data))

	for _, ticker := range tdl.Data {
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

func makeTickersRequest(ctx context.Context, url, assetsQueryParam string) (*TickerDataList, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("ids", assetsQueryParam)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tdl TickerDataList
	if err := json.Unmarshal(body, &tdl); err != nil {
		return nil, err
	}

	return &tdl, nil
}
