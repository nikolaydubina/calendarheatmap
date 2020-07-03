#!/bin/sh

set -x

go run ../main.go -colorscale=PuBu9 -output=chart_PuBu9.png
go run ../main.go -colorscale=GnBu9 -output=chart_GnBu9.png
go run ../main.go -colorscale=YlGn9 -output=chart_YlGn9.png

go run ../main.go -colorscale=PuBu9 -output=chart_PuBu9_nolabels.png -labels=false
go run ../main.go -colorscale=PuBu9 -output=chart_PuBu9_noseparator.png -monthsep=false
go run ../main.go -colorscale=PuBu9 -output=chart_PuBu9_noseparator_nolabels.png -monthsep=false -labels=false