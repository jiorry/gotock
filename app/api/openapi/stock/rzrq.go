package stock

import (
	"github.com/jiorry/gotock/app/lib/tools/dfcf"

	"github.com/kere/gos"
	"github.com/kere/gos/lib/util"
)

type RZRQApi struct {
	gos.WebApi
}

func (a *RZRQApi) SumData() ([]*dfcf.RzrqSumItemData, error) {
	return dfcf.GetRzrqSumData()
}

func (a *RZRQApi) StockData(args util.MapData) ([]*dfcf.RzrqStockData, error) {
	return dfcf.GetRzrqStockData(args.GetString("code"))
}

func (a *RZRQApi) CachedStockData() ([]string, error) {
	return dfcf.StockCachedList(), nil
}
