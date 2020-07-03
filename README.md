[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/calendarheatmap)](https://goreportcard.com/report/github.com/nikolaydubina/calendarheatmap)

Self-contained, plain Go implementation of calendar heatmap inspired by Github contribution activity.

Colorscales
![PuBu9](example/chart_PuBu9.png)
![GnBu9](example/chart_GnBu9.png)
![YlGn9](example/chart_YlGn9.png)

Without month separator
![PuBu9_noseparator](example/chart_PuBu9_noseparator.png)

Without labels
![PuBu9_noseparator](example/chart_PuBu9_nolabels.png)

Without labels, without separator
![PuBu9_noseparator](example/chart_PuBu9_noseparator_nolabels.png)

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
Check full example at `example/main.go` and `input.txt` for more details.
Generate examples above with `./generate_examples.sh`.