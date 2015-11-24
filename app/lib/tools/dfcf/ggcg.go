package dfcf

// 高管持股
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jiorry/gotock/app/lib/tools/wget"
	"github.com/jiorry/gotock/app/lib/util"
	"github.com/kere/gos"
	"github.com/kere/gos/db"
)

// FillYyb 抓取数据
func FillGgcg() error {
	v, err := FetchGgcg(1)
	if err != nil {
		return err
	}
	pages := v.Pages
	if pages > 20 {
		pages = 20
	}

	exists := db.NewExistsBuilder("ggcg")

	var l []*ggcgVO
	var item *ggcgVO
	for i := 2; i < pages+1; i++ {
		v, err = FetchGgcg(i)
		if err != nil {
			return err
		}

		l = v.Parse()
		for _, item = range l {
			item.Init(item)
			if exists.Where("stock_id=? and date=?", item.StockID, item.Date).Exists() {
				continue
			}
			item.Create()

			if err != nil {
				gos.DoError(err)
			}
		}

		time.Sleep(1 * time.Second)
	}

	return err
}

// fetchAndFillYyb 抓取数据
func FetchGgcg(page int) (*jsonGgcgData, error) {
	// http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=LHB&sty=YYTJ&stat=6&sr=0&st=1&p=2&ps=50&js=var%20XvAVhGPE={%22data%22:[(x)],%22pages%22:%22(pc)%22,%22update%22:%22(ud)%22}&rt=48257541
	formt := `http://datainterface.eastmoney.com/EM_DataCenter/JS.aspx?type=GG&sty=GGMX&p=%d&ps=%d&js=var%%20ftzqaLVS={pages:(pc),data:[(x)]}&rt=%d`
	pageLimit := 100

	body, err := wget.GetBody(fmt.Sprintf(formt, page, pageLimit, time.Now().Unix()))
	if err != nil {
		return nil, err
	}

	src := body[13:]
	src = bytes.Replace(src, []byte("pages:"), []byte(`"pages":`), 1)
	src = bytes.Replace(src, []byte(",data:"), []byte(`,"data":`), 1)

	v := &jsonGgcgData{}
	err = json.Unmarshal(src, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

type ggcgVO struct {
	db.BaseVO
	StockID       int64   `json:"stock_id" skip:"update"`
	Date          string  `json:"date"`
	Person        string  `json:"person"`
	Price         float64 `json:"price"`           //
	Amount        float64 `json:"amount"`          //
	Total         float64 `json:"total"`           //
	Remain        int64   `json:"total"`           //
	Proportion    float64 `json:"proportion"`      // 变动比例
	Reason        string  `json:"reason"`          // 变动原因
	PersonRel     string  `json:"person_rel"`      // 相关董监高人员姓名
	PersonRelType string  `json:"person_rel_type"` // 变动人与董监高的关系
	Job           string  `json:"job"`             // 职务
}

type jsonGgcgData struct {
	Pages int      `json:"pages"`
	Data  []string `json:"data"`
}

func (s *jsonGgcgData) Parse() []*ggcgVO {
	dataset := make([]*ggcgVO, 0)

	query := db.NewQueryBuilder("stock")
	upd := db.NewUpdateBuilder("stock")

	var tmp []string
	var arr []string
	var row db.DataRow
	var ggcg *ggcgVO

	for _, item := range s.Data {
		arr = strings.Split(item, ",")
		tmp = strings.Split(strings.ToLower(arr[15]), ".")
		row, _ = query.Where("code=? and ctype=?", tmp[0], tmp[1]).QueryOne()

		if row.Empty() {
			continue
		}
		upd.Where("id=?", row.GetInt64("id")).Update(db.DataRow{"name": arr[9]})
		// "0.00444,谢飞鹏,002420,谢杏思,A股,2015-11-16,-17800,0,9.44,毅昌股份,父母,YCGF,竞价交易,-168032,董事,002420.SZ"
		//     0      1     2      3    4      5         6  7   8     9    10   11    12       13   14    15
		ggcg = &ggcgVO{
			StockID:       row.GetInt64("id"),
			Date:          arr[5],
			Person:        arr[3],
			Price:         util.ParseMoney(arr[8]),
			Amount:        util.ParseMoney(arr[6]),
			Total:         util.ParseMoney(arr[13]),
			Remain:        int64(util.ParseMoney(arr[7])),
			Proportion:    util.ParseMoney(arr[0]),
			Reason:        arr[12],
			PersonRel:     arr[1],
			PersonRelType: arr[10],
			Job:           arr[14],
		}

		dataset = append(dataset, ggcg)
	}

	return dataset
}
