package main

import (
	"math/rand"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	itemCntPie = 4
	seasons    = []string{"بهار", "تابستان", "پاییز ", "زمستان"}
)

func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < itemCntPie; i++ {
		items = append(items, opts.PieData{Name: seasons[i], Value: rand.Intn(100)})
	}
	return items
}

func pieShowLabel(w http.ResponseWriter, _ *http.Request) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "نمودار دایره ای سرانه شادابی مردم در فصل های مختلف",
			Subtitle: "اعداد به صورت تصادفی انتخاب شده اند. در هر بار اعداد جدید باز تولید می شوند",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)
	pie.Render(w)
}

func pieRoseRadius(w http.ResponseWriter, _ *http.Request) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "نمودار دایره ای سرانه شادابی مردم در فصل های مختلف",
			Subtitle: "اعداد به صورت تصادفی انتخاب شده اند. در هر بار اعداد جدید باز تولید می شوند",
		}),
	)

	pie.AddSeries("pie", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"30%", "75%"},
				RoseType: "radius",
			}),
		)
	pie.Render(w)
}
