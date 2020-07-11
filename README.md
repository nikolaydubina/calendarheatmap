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

Example script,
```
$ go build

$ ./calendarheatmap -input testdata/basic.json -output basicjson.png

$ ./calendarheatmap -input testdata/custom.txt -output custom.png -input-format row-day-seconds-count

$ ./calendarheatmap -h

Usage of ./calendarheatmap:
  -colorscale string
        refer to colorscales for examples (default "PuBu9")
  -input string
        input filenam (default "input.txt")
  -intput-format /parsers
        format of input file refer to /parsers for examples (default "json-basic")
  -labels
        labels for weekday and months (default true)
  -monthsep
        render month separator (default true)
  -output string
        output filename, will export as PNG (default "chart.png")
  -output-format string
        output format (png, jpeg, gif) (default "png")
```

TODO:
- [ ] SVG support
- [ ] select start and end date
