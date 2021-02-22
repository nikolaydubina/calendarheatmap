[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/calendarheatmap)](https://goreportcard.com/report/github.com/nikolaydubina/calendarheatmap)
[![Go Reference](https://pkg.go.dev/badge/github.com/nikolaydubina/calendarheatmap/charts.svg)](https://pkg.go.dev/github.com/nikolaydubina/calendarheatmap/charts)
[![codecov](https://codecov.io/gh/nikolaydubina/calendarheatmap/branch/master/graph/badge.svg)](https://codecov.io/gh/nikolaydubina/calendarheatmap)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

Self-contained, plain Go implementation of calendar heatmap inspired by Github contribution activity.

```
$ go build
$ echo '{
    "2020-05-16": 8,
    "2020-05-17": 13,
    "2020-05-18": 5,
    "2020-05-19": 8,
    "2020-05-20": 5
}' | ./calendarheatmap > chart.png
```

- Basic

![basic](charts/testdata/basic.png)

- Negative values presented via an alternative color scheme

![basic_alt](charts/testdata/basic_alt_YlGn9.png)

- Colorscales
![col1](charts/testdata/colorscale_1.png)
![col2](charts/testdata/colorscale_2.png)

- UTF-8
![col1](charts/testdata/korean.png)

- SVG

![svg](charts/testdata/korean.svg)

- Without month separator
![nosep](charts/testdata/noseparator.png)

- Without labels
![nolab](charts/testdata/nolabels.png)

- Without labels, without separator
![nosep_nolab](charts/testdata/noseparator_nolabels.png)


## GitHub stars over time

[![GitHub stars over time](https://starchart.cc/nikolaydubina/calendarheatmap.svg)](https://starchart.cc/nikolaydubina/calendarheatmap)
