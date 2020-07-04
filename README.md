[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/calendarheatmap)](https://goreportcard.com/report/github.com/nikolaydubina/calendarheatmap)
[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/nikolaydubina/calendarheatmap@v1.0.0/charts)
[![codecov](https://codecov.io/gh/nikolaydubina/calendarheatmap/branch/master/graph/badge.svg)](https://codecov.io/gh/nikolaydubina/calendarheatmap)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

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

Example module, save output in formats supported by `image` module (PNG, JPEG, GIF).

```go
countByDayOfYear := map[int]int{
    1: 10,
    22: 15,
    150: 22,
    366: 55,
    ...
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

Example script, output as PNG.
```
go run main.go -h
```

TODO:
- [ ] SVG support
- [ ] select start and end date
- [ ] JPEG, GIF in script output
- [ ] CSV, JSON in script input
