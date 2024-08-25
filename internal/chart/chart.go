package chart

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"trading-app-service/internal/models"
)

const dateFormat = "2006/01/02 15:04"

type CandleStick struct {
	Page   *components.Page
	Charts []components.Charter
	Opts   []charts.SeriesOpts
	Klines []*charts.Kline
}

func NewCandleStickPage() *CandleStick {
	return &CandleStick{
		Page:   components.NewPage(),
		Charts: make([]components.Charter, 0),
	}
}

func (cs *CandleStick) AddCandleStickChart(name string, candles []models.ChartCandle) {
	kl := charts.NewKLine()

	kl.SetGlobalOptions(
		charts.WithToolboxOpts(opts.Toolbox{Show: newTrue()}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      newTrue(),
			Trigger:   "axis",
			TriggerOn: "mousemove|click",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: name,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 7,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale:     newTrue(),
			SplitLine: &opts.SplitLine{Show: newTrue()},
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "inside",
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "slider",
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	for _, c := range candles {
		date := c.Timestamp.Format(dateFormat)

		x = append(x, date)
		y = append(y, opts.KlineData{Name: c.Symbol, Value: [4]float64{c.Open, c.Close, c.Low, c.High}})

		if c.MarkPoint != "" && c.MarkPointPrice != 0 {
			cs.AddMarkPoint(c.MarkPoint, date, c.MarkPointPrice)
		}
	}

	styleOpts := charts.WithItemStyleOpts(opts.ItemStyle{
		Color:        "#00da3c",
		Color0:       "#ec0000",
		BorderColor:  "#008F28",
		BorderColor0: "#8A0000",
	})

	cs.Opts = append(cs.Opts, styleOpts)

	kl = kl.SetXAxis(x).AddSeries(candles[0].Symbol, y)

	cs.Klines = append(cs.Klines, kl)

}

func (cs *CandleStick) AddMarkPoint(name, date string, value float64) {
	markPoint := charts.WithMarkPointNameCoordItemOpts(opts.MarkPointNameCoordItem{
		Name:       name,
		Coordinate: []interface{}{date, value},
		Label: &opts.Label{
			Show:      newTrue(),
			Formatter: name,
			Color:     "black",
		},
	})
	cs.Opts = append(cs.Opts, markPoint)
}

func (cs *CandleStick) Render(w io.Writer) {
	for _, kl := range cs.Klines {
		kl.SetSeriesOptions(cs.Opts...)
		cs.Charts = append(cs.Charts, kl)
	}

	cs.Page.AddCharts(cs.Charts...)
	err := cs.Page.Render(w)
	if err != nil {
		return
	}
}

func newTrue() *bool {
	b := true
	return &b
}
