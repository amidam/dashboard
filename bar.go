package main

import (
	"math/rand"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 12; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func bar(w http.ResponseWriter, _ *http.Request) {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemePurplePassion}),
		charts.WithTitleOpts(opts.Title{
			Title:    "نمودار ستونی درآمد ماهیانه سه نفر بر پایه دلار",
			Subtitle: "مفادیر به صورت تصادفی تعریف شده اند. در هر بار مقادیر جدید بازتعریف می شوند",
		}))

	// Put data into instance
	bar.SetXAxis([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems()).
		AddSeries("Category C", generateBarItems())
	bar.Render(w)
}
