package dzh

import (
	"bytes"
	"encoding/binary"
	"os"
	"reflect"
)

type FinData struct {
	Code  string  `json:"code"`
	Fin1  int64   `json:"fin1"`  // 数据下载日期
	Fin2  int64   `json:"fin2"`  // 数据更新日期
	Fin3  int64   `json:"fin3"`  // 公司上市日期
	Fin4  float64 `json:"fin4"`  // 每股收益
	Fin5  float64 `json:"fin5"`  // 每股净资产
	Fin6  float64 `json:"fin6"`  // 净资产收率(%)
	Fin7  float64 `json:"fin7"`  // 每股经营现金
	Fin8  float64 `json:"fin8"`  // 每股公积金
	Fin9  float64 `json:"fin9"`  // 每股未分配
	Fin10 float64 `json:"fin10"` // 股东权益比
	Fin11 float64 `json:"fin11"` // 净利润同比
	Fin12 float64 `json:"fin12"` // 主营收同比
	Fin13 float64 `json:"fin13"` // 销售毛利率
	Fin14 float64 `json:"fin14"` // 调整每股净资产
	Fin15 float64 `json:"fin15"` // 总资产
	Fin16 float64 `json:"fin16"` // 流动资产
	Fin17 float64 `json:"fin17"` // 固定资产
	Fin18 float64 `json:"fin18"` // 无形资产
	Fin19 float64 `json:"fin19"` // 流动负债
	Fin20 float64 `json:"fin20"` // 长期负债
	Fin21 float64 `json:"fin21"` // 总负债
	Fin22 float64 `json:"fin22"` // 股东权益
	Fin23 float64 `json:"fin23"` // 资本公积金
	Fin24 float64 `json:"fin24"` // 经营现金流量
	Fin25 float64 `json:"fin25"` // 筹资现金流量
	Fin26 float64 `json:"fin26"` // 投资现金流量
	Fin27 float64 `json:"fin27"` // 现金增加额
	Fin28 float64 `json:"fin28"` // 主营收入
	Fin29 float64 `json:"fin29"` // 主营利润
	Fin30 float64 `json:"fin30"` // 营业利润
	Fin31 float64 `json:"fin31"` // 投资收益
	Fin32 float64 `json:"fin32"` // 营业外收支
	Fin33 float64 `json:"fin33"` // 利润总额
	Fin34 float64 `json:"fin34"` // 净利润
	Fin35 float64 `json:"fin35"` // 未分配利润
	Fin36 float64 `json:"fin36"` // 总股本
	Fin37 float64 `json:"fin37"` // 无限售股
	Fin38 float64 `json:"fin38"` // A股
	Fin39 float64 `json:"fin39"` // B股
	Fin40 float64 `json:"fin40"` // 境外上市股
	Fin41 float64 `json:"fin41"` // 其它流通股
	Fin42 float64 `json:"fin42"` // 限售股
	Fin43 float64 `json:"fin43"` // 国家持股
	Fin44 float64 `json:"fin44"` // 国有法人股
	Fin45 float64 `json:"fin45"` // 境内法人股
	Fin46 float64 `json:"fin46"` // 境内自然人股
	Fin47 float64 `json:"fin47"` // 其它发起人股
	Fin48 float64 `json:"fin48"` // 募集法人股
	Fin49 float64 `json:"fin49"` // 境外法人股
	Fin50 float64 `json:"fin50"` // 境外自然人股
	Fin51 float64 `json:"fin51"` // 优先股或其它
}

func ReadData(name string) ([]*FinData, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	src := make([]byte, 216)

	var dataset = make([]*FinData, 0)
	var i int = 0
	var xi int32
	var xf float32
	var buf *bytes.Buffer
	var data *FinData

	for err == nil {
		_, err = file.Read(src)
		data = &FinData{}
		data.Code = string(src[8:16])

		for i = 0; i < 49; i++ {
			buf = bytes.NewBuffer(src[20+i*4 : 24+i*4])
			switch i {
			case 0, 1, 2:
				binary.Read(buf, binary.LittleEndian, &xi)
				reflect.ValueOf(data).Elem().Field(i + 1).SetInt(int64(xi))

			default:
				binary.Read(buf, binary.LittleEndian, &xf)
				reflect.ValueOf(data).Elem().Field(i + 1).SetFloat(float64(xf))
			}
		}
		if data.Fin1 == 0 || data.Fin2 == 0 || data.Fin4 == 0 {
			continue
		}
		if data.Code[:2] == "SO" {
			continue
		}

		dataset = append(dataset, data)
	}

	return dataset, nil
}
