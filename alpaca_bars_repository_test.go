// Copyright (c) 2021-2024 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package indicatoralpaca_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/cinar/indicatoralpaca"
)

func newClient(response string) *marketdata.Client {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	}))

	return marketdata.NewClient(
		marketdata.ClientOpts{
			APIKey:    "key",
			APISecret: "secret",
			BaseURL:   server.URL,
		},
	)
}

func TestNewAlpacaBarsRepository(t *testing.T) {
	repository := indicatoralpaca.NewAlpacaBarsRepository("", "")
	if repository == nil {
		t.Fatal("expected repository")
	}
}

func TestAlpacaBarsRepositoryAssets(t *testing.T) {
	repository := indicatoralpaca.NewAlpacaBarsRepositoryWithClient(newClient(""))

	_, err := repository.Assets()
	if err != errors.ErrUnsupported {
		t.Fatal(err)
	}
}

func TestAlpacaBarsRepositoryGet(t *testing.T) {
	repository := indicatoralpaca.NewAlpacaBarsRepositoryWithClient(newClient(
		`{"bars":{"A":[{"t":"2021-10-15T16:00:00Z","o":3378.14,"h":3380.815,"l":3376.3001,"c":3379.72,"v":211689,"n":5435,"vw":3379.041755}]},"next_page_token":null}`,
	))

	snapshots, err := repository.Get("A")
	if err != nil {
		t.Fatal(err)
	}

	snapshot := <-snapshots

	expected := time.Date(2021, 10, 15, 16, 0, 0, 0, time.UTC)
	actual := snapshot.Date

	if actual != expected {
		t.Fatalf("actual %v expected %v", actual, expected)
	}
}

func TestAlpacaBarsRepositoryGetFailed(t *testing.T) {
	repository := indicatoralpaca.NewAlpacaBarsRepositoryWithClient(newClient(""))

	_, err := repository.Get("A")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTiingoRepositoryLastDate(t *testing.T) {
	repository := indicatoralpaca.NewAlpacaBarsRepositoryWithClient(newClient(
		`{"bars":{"A":{"t":"2021-10-15T16:00:00Z","o":3378.14,"h":3380.815,"l":3376.3001,"c":3379.72,"v":211689,"n":5435,"vw":3379.041755}}}`,
	))

	actual, err := repository.LastDate("A")
	if err != nil {
		t.Fatal(err)
	}

	expected := time.Date(2021, 10, 15, 16, 0, 0, 0, time.UTC)

	if actual != expected {
		t.Fatalf("actual %v expected %v", actual, expected)
	}
}

func TestAlpacaRepositorLastDateFailed(t *testing.T) {
	repository := indicatoralpaca.NewAlpacaBarsRepositoryWithClient(newClient(""))

	_, err := repository.LastDate("A")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestAlpacaBarsRepositoryAppend(t *testing.T) {
	repository := indicatoralpaca.NewAlpacaBarsRepositoryWithClient(newClient(""))

	err := repository.Append("A", nil)
	if err != errors.ErrUnsupported {
		t.Fatal(err)
	}
}
