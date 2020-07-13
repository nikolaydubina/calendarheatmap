[![Go Report Card](https://goreportcard.com/badge/github.com/nikolaydubina/calendarheatmap)](https://goreportcard.com/report/github.com/nikolaydubina/calendarheatmap)
[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/nikolaydubina/calendarheatmap/charts)
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

Example module, next save output in formats supported by `image` module (PNG, JPEG, GIF).

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

Example script
```
$ cat testdata/basic.json | ./calendarheatmap -output jpeg > chart.jpeg
$ cat testdata/custom.txt | ./calendarheatmap -input row-day-seconds-count > custom.png
$ ./calendarheatmap -h

Usage of ./calendarheatmap:
  -colorscale string
        refer to colorscales for examples (default "PuBu9")
  -input string
        format of input file, refer to parsers module (default "json-basic")
  -labels
        labels for weekday and months (default true)
  -monthsep
        render month separator (default true)
  -output string
        output format (png, jpeg, gif) (default "png")
```
