package main

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var (
	bpX = [...]string{"Concorde", "Boeing 747", "Airbus a380", "Boeing 777", "Embraer 170"}
	bpY = [][]int{
		{850, 740, 900, 1070, 930, 850, 950, 980, 980, 880, 1000, 980, 930, 650, 760, 810, 1000, 1000, 960, 960},
		{960, 940, 960, 940, 880, 800, 850, 880, 900, 840, 830, 790, 810, 880, 880, 830, 800, 790, 760, 800},
		{880, 880, 880, 860, 720, 720, 620, 860, 970, 950, 880, 910, 850, 870, 840, 840, 850, 840, 840, 840},
		{890, 810, 810, 820, 800, 770, 760, 740, 750, 760, 910, 920, 890, 860, 880, 720, 840, 850, 850, 780},
		{890, 840, 780, 810, 760, 810, 790, 810, 820, 850, 870, 870, 810, 740, 810, 940, 950, 800, 810, 870},
	}
)

func generateBoxPlotItems(boxPlotData [][]int) []opts.BoxPlotData {
	items := make([]opts.BoxPlotData, 0)
	for i := 0; i < len(boxPlotData); i++ {
		items = append(items, opts.BoxPlotData{Value: boxPlotData[i]})
	}
	return items

}

func boxPlotBase(w http.ResponseWriter, _ *http.Request) {
	box := charts.NewBoxPlot()
	box.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "نمودار جعبه ای برای نشان دادن تغییرات سرعت هواپیماهای مختلف",
			Subtitle: "سرعت بر اساس کیلومتر بر ساعت محاسبه شده است",
		},
		),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemePurplePassion}),
	)

	box.SetXAxis(bpX).AddSeries("boxplot", generateBoxPlotItems(bpY))
	box.Render(w)
}
