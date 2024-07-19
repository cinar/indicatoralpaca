[![GoDoc](https://godoc.org/github.com/cinar/indicatoralpaca?status.svg)](https://godoc.org/github.com/cinar/indicatoralpaca) [![License](https://img.shields.io/badge/License-AGPLv3-blue.svg)](https://opensource.org/licenses/AGPLv3) [![Go Report Card](https://goreportcard.com/badge/github.com/cinar/indicatoralpaca)](https://goreportcard.com/report/github.com/cinar/indicatoralpaca) ![Go CI](https://github.com/cinar/indicatoralpaca/actions/workflows/ci.yml/badge.svg) [![codecov](https://codecov.io/gh/cinar/indicatoralpaca/graph/badge.svg?token=AFZLMEBTVM)](https://codecov.io/gh/cinar/indicatoralpaca)

Indicator Alpaca
================

Indicator Alpaca provides [Alpaca Markets](https://alpaca.markets) integration for [Indicator](https://github.com/cinar/indicator) Golang module that provides a rich set of technical analysis indicators, strategies, and a framework for backtesting.

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

indicatoralpaca
===============

```go
import "github.com/cinar/indicatoralpaca"
```

Package indicatoralpaca provides Alpaca Markets inintegration for the Indicator technical analysis library.

Index
-----

-	[Constants](#constants)
-	[type AlpacaRepository](#AlpacaRepository)
	-	[func NewAlpacaRepository\(apiKey, apiSecret string\) \*AlpacaRepository](#NewAlpacaRepository)
	-	[func NewAlpacaRepositoryWithClient\(client \*marketdata.Client\) \*AlpacaRepository](#NewAlpacaRepositoryWithClient)
	-	[func \(\*AlpacaRepository\) Append\(\_ string, \_ \<\-chan \*asset.Snapshot\) error](#AlpacaRepository.Append)
	-	[func \(\*AlpacaRepository\) Assets\(\) \(\[\]string, error\)](#AlpacaRepository.Assets)
	-	[func \(r \*AlpacaRepository\) Get\(name string\) \(\<\-chan \*asset.Snapshot, error\)](#AlpacaRepository.Get)
	-	[func \(r \*AlpacaRepository\) GetSince\(name string, date time.Time\) \(\<\-chan \*asset.Snapshot, error\)](#AlpacaRepository.GetSince)
	-	[func \(r \*AlpacaRepository\) LastDate\(name string\) \(time.Time, error\)](#AlpacaRepository.LastDate)

Constants
---------

<a name="DefaultAlpacaRepositoryTimeFrameUnit"></a>

```go
const (
	// DefaultAlpacaRepositoryTimeFrameUnit is the default time frame unit of a day.
	DefaultAlpacaRepositoryTimeFrameUnit = marketdata.Day
)
```

<a name="AlpacaRepository"></a>

type [AlpacaRepository](https://github.com/cinar/indicatoralpaca/blob/main/alpaca_repository.go#L23-L29)
--------------------------------------------------------------------------------------------------------

AlpacaRepository provides access to financial market data, retrieving asset snapshots, by interacting with the Alpaca Markets API. To use this repository, you'll need a valid API key from https://alpaca.markets.

```go
type AlpacaRepository struct {

	// GetBarsRequestTemplate is the request template used to get the bars.
	GetBarsRequestTemplate marketdata.GetBarsRequest
	// contains filtered or unexported fields
}
```

<a name="NewAlpacaRepository"></a>

### func [NewAlpacaRepository](https://github.com/cinar/indicatoralpaca/blob/main/alpaca_repository.go#L44)

```go
func NewAlpacaRepository(apiKey, apiSecret string) *AlpacaRepository
```

NewAlpacaRepository initializes an Alpaca Markets repository with the given API key and API secret.

<a name="NewAlpacaRepositoryWithClient"></a>

### func [NewAlpacaRepositoryWithClient](https://github.com/cinar/indicatoralpaca/blob/main/alpaca_repository.go#L54)

```go
func NewAlpacaRepositoryWithClient(client *marketdata.Client) *AlpacaRepository
```

NewAlpacaRepositoryWithClient initializes an Alpaca Markets repository with the given client.

<a name="AlpacaRepository.Append"></a>

### func \(\*AlpacaRepository\) [Append](https://github.com/cinar/indicatoralpaca/blob/main/alpaca_repository.go#L104)

```go
func (*AlpacaRepository) Append(_ string, _ <-chan *asset.Snapshot) error
```

Append adds the given snapshows to the asset with the given name.

<a name="AlpacaRepository.Assets"></a>

### func \(\*AlpacaRepository\) [Assets](https://github.com/cinar/indicatoralpaca/blob/main/alpaca_repository.go#L64)

```go
func (*AlpacaRepository) Assets() ([]string, error)
```

Assets returns the names of all assets in the repository.

<a name="AlpacaRepository.Get"></a>

### func \(\*AlpacaRepository\) [Get](https://github.com/cinar/indicatoralpaca/blob/main/alpaca_repository.go#L69)

```go
func (r *AlpacaRepository) Get(name string) (<-chan *asset.Snapshot, error)
```

Get attempts to return a channel of snapshots for the asset with the given name.

<a name="AlpacaRepository.GetSince"></a>

### func \(\*AlpacaRepository\) [GetSince](https://github.com/cinar/indicatoralpaca/blob/main/alpaca_repository.go#L74)

```go
func (r *AlpacaRepository) GetSince(name string, date time.Time) (<-chan *asset.Snapshot, error)
```

GetSince attempts to return a channel of snapshots for the asset with the given name since the given date.

<a name="AlpacaRepository.LastDate"></a>

### func \(\*AlpacaRepository\) [LastDate](https://github.com/cinar/indicatoralpaca/blob/main/alpaca_repository.go#L92)

```go
func (r *AlpacaRepository) LastDate(name string) (time.Time, error)
```

LastDate returns the date of the last snapshot for the asset with the given name.

Generated by [gomarkdoc](https://github.com/princjef/gomarkdoc)

<!-- gomarkdoc:embed:end -->

Contributing to the Project
---------------------------

Anyone can contribute to Indicator library. Please make sure to read our [Contributor Covenant Code of Conduct](./CODE_OF_CONDUCT.md) guide first. Follow the [How to Contribute to Indicator](./CONTRIBUTING.md) to contribute. Signining a [Contributor Agreement](./CLA.md) is also required to contribute to the project.

Disclaimer
----------

The information provided on this project is strictly for informational purposes and is not to be construed as advice or solicitation to buy or sell any security.

License
-------

Indicator libary dual-licensed under GNU AGPLv3 License and a commercial license. For free use and modifications of the code, you can use the AGPLv3 license. If you require commercial license with different terms, please contact me.

```
Copyright (c) 2021-2024 Onur Cinar.    
The source code is provided under GNU AGPLv3 License.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```