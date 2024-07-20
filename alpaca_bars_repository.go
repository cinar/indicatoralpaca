// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicatoralpaca

package indicatoralpaca

import (
	"errors"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/helper"
)

const (
	// DefaultAlpacaBarsRepositoryTimeFrameUnit is the default time frame unit of a day.
	DefaultAlpacaBarsRepositoryTimeFrameUnit = marketdata.Day
)

// AlpacaBarsRepository provides access to financial market data, retrieving asset snapshots, by interacting with the
// Alpaca Markets API. To use this repository, you'll need a valid API key from https://alpaca.markets.
type AlpacaBarsRepository struct {
	// Client is the Alpaca Markets client.
	client *marketdata.Client

	// GetBarsRequestTemplate is the request template used to get the bars.
	GetBarsRequestTemplate marketdata.GetBarsRequest
}

// ToSnapshot converts the Alpaca Markets bar to a snapshot.
func barToSnapshot(bar marketdata.Bar) *asset.Snapshot {
	return &asset.Snapshot{
		Date:   bar.Timestamp,
		Open:   bar.Open,
		High:   bar.High,
		Low:    bar.Low,
		Close:  bar.Close,
		Volume: int64(bar.Volume),
	}
}

// NewAlpacaBarsRepository initializes an Alpaca Markets repository with the given API key and API secret.
func NewAlpacaBarsRepository(apiKey, apiSecret string) *AlpacaBarsRepository {
	return NewAlpacaBarsRepositoryWithClient(marketdata.NewClient(
		marketdata.ClientOpts{
			APIKey:    apiKey,
			APISecret: apiSecret,
		},
	))
}

// NewAlpacaBarsRepositoryWithClient initializes an Alpaca Markets repository with the given client.
func NewAlpacaBarsRepositoryWithClient(client *marketdata.Client) *AlpacaBarsRepository {
	return &AlpacaBarsRepository{
		client: client,
		GetBarsRequestTemplate: marketdata.GetBarsRequest{
			TimeFrame: marketdata.NewTimeFrame(1, DefaultAlpacaBarsRepositoryTimeFrameUnit),
		},
	}
}

// Assets returns the names of all assets in the repository.
func (*AlpacaBarsRepository) Assets() ([]string, error) {
	return nil, errors.ErrUnsupported
}

// Get attempts to return a channel of snapshots for the asset with the given name.
func (r *AlpacaBarsRepository) Get(name string) (<-chan *asset.Snapshot, error) {
	return r.GetSince(name, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
}

// GetSince attempts to return a channel of snapshots for the asset with the given name since the given date.
func (r *AlpacaBarsRepository) GetSince(name string, date time.Time) (<-chan *asset.Snapshot, error) {
	request := r.GetBarsRequestTemplate
	request.Start = date

	bars, err := r.client.GetBars(name, request)
	if err != nil {
		return nil, err
	}

	snapshots := helper.Map(
		helper.SliceToChan(bars),
		barToSnapshot,
	)

	return snapshots, nil
}

// LastDate returns the date of the last snapshot for the asset with the given name.
func (r *AlpacaBarsRepository) LastDate(name string) (time.Time, error) {
	request := marketdata.GetLatestBarRequest{}

	bar, err := r.client.GetLatestBar(name, request)
	if err != nil {
		return time.Time{}, err
	}

	return bar.Timestamp, nil
}

// Append adds the given snapshows to the asset with the given name.
func (*AlpacaBarsRepository) Append(_ string, _ <-chan *asset.Snapshot) error {
	return errors.ErrUnsupported
}
