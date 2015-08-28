package stock

import (
	"../../../lib/tools/dfcf"

	"github.com/kere/gos"
)

type RZRQApi struct {
	gos.WebApi
}

func (a *RZRQApi) SumData() ([]*dfcf.RzrqSumItemData, error) {
	result, err := dfcf.RzrqSumData()
	if err != nil {
		return nil, err
	}

	return result, nil
}
