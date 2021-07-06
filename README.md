[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/calendarheatmap)](https://goreportcard.com/report/github.com/nikolaydubina/calendarheatmap)
[![Go Reference](https://pkg.go.dev/badge/github.com/nikolaydubina/calendarheatmap/charts.svg)](https://pkg.go.dev/github.com/nikolaydubina/calendarheatmap/charts)
[![codecov](https://codecov.io/gh/nikolaydubina/calendarheatmap/branch/master/graph/badge.svg)](https://codecov.io/gh/nikolaydubina/calendarheatmap)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go#science-and-data-analysis)

Self-contained, plain Go implementation of calendar heatmap inspired by GitHub contribution activity.

```
$ go install github.com/nikolaydubina/calendarheatmap@latest 
$ echo '{
    "2020-05-16": 8,
    "2020-05-17": 13,
    "2020-05-18": 5,
    "2020-05-19": 8,
    "2020-05-20": 5
}' | calendarheatmap > chart.png
```

Basic

![basic](docs/basic.png)

Colorscales

![col1](docs/colorscale-1.png)
![col2](docs/colorscale-2.png)
![col2](docs/colorscale-3.png)

UTF-8
![col1](docs/korean.png)

SVG

![svg](docs/korean.svg)

Without month separator
![nosep](docs/noseparator.png)

Without labels
![nolab](docs/nolabels.png)

Without labels, without separator
![nosep_nolab](docs/noseparator_nolabels.png)

## GitHub stars over time

[![GitHub stars over time](https://starchart.cc/nikolaydubina/calendarheatmap.svg)](https://starchart.cc/nikolaydubina/calendarheatmap)
