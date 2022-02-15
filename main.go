package main

import "net/http"

func main() {
	http.HandleFunc("/line", line)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/box", boxPlotBase)
	http.HandleFunc("/pie-label", pieShowLabel)
	http.HandleFunc("/pie-radius", pieRoseRadius)
	http.HandleFunc("/3d-base", bar3DBase)
	http.HandleFunc("/3d-rotate", bar3DAutoRotate)
	http.HandleFunc("/heat", heatMapBase)
	http.ListenAndServe(":8081", nil)
}
