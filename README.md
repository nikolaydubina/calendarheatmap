[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/calendarheatmap)](https://goreportcard.com/report/github.com/nikolaydubina/calendarheatmap)
[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/nikolaydubina/calendarheatmap/charts)
[![GoDoc](https://godoc.org/github.com/nikolaydubina/calendarheatmap/charts?status.svg)](https://godoc.org/github.com/nikolaydubina/calendarheatmap/charts)

Self-contained, plain Go implementation of calendar heatmap inspired by Github contribution activity.

Colorscales
![PuBu9](examples/chart_PuBu9.png)
![GnBu9](examples/chart_GnBu9.png)
![YlGn9](examples/chart_YlGn9.png)

Without month separator
![PuBu9_noseparator](examples/chart_PuBu9_noseparator.png)

Without labels
![PuBu9_noseparator](examples/chart_PuBu9_nolabels.png)

Without labels, without separator
![PuBu9_noseparator](examples/chart_PuBu9_noseparator_nolabels.png)

Example:

```go
countByDayOfYear := map[int]int{
    1: 10,
    22: 15,
    150: 22,
    366: 55,
}

img := charts.NewHeatmap(charts.HeatmapConfig{
    Year:               2020,
    CountByDay:         countByDay,
    ColorScale:         colorscales.PuBu9,
    DrawMonthSeparator: true,
    DrawLabels:         true,
    ...
})
```

You can run it in your Go code or as standalone script. 
Check full example at `examples/main.go` and `input.txt` for more details.
Generate examples above with `./generate_examples.sh`.
