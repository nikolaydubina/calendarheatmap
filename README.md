[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/calendarheatmap)](https://goreportcard.com/report/github.com/nikolaydubina/calendarheatmap)
[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/nikolaydubina/calendarheatmap@v1.0.0/charts)
[![codecov](https://codecov.io/gh/nikolaydubina/calendarheatmap/branch/master/graph/badge.svg)](https://codecov.io/gh/nikolaydubina/calendarheatmap)

Self-contained, plain Go implementation of calendar heatmap inspired by Github contribution activity.

Basic
![basic](charts/testdata/basic.png)

Colorscales
![col1](charts/testdata/colorscale_1.png)
![col2](charts/testdata/colorscale_2.png)

Without month separator
![nosep](charts/testdata/noseparator.png)

Without labels
![nolab](charts/testdata/nolabels.png)

Without labels, without separator
![nosep_nolab](charts/testdata/noseparator_nolabels.png)

Example module:

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

Example script:
```
go run main.go -h
```

You can run it in your Go code or as standalone script. 
Check output examples in `charts/testdata/` and optional input data in `testdata/input.txt` for more details.
