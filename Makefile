build:
	go build

test:
	go test ./...

docs: build
	cat charts/testdata/basic.json | ./calendarheatmap > docs/basic.png
	CALENDAR_HEATMAP_ASSETS_PATH=assets cat charts/testdata/basic.json | ./calendarheatmap -colorscale=purple-blue-9.csv > docs/colorscale-1.png
	CALENDAR_HEATMAP_ASSETS_PATH=assets cat charts/testdata/basic.json | ./calendarheatmap -colorscale=green-blue-9.csv > docs/colorscale-2.png
	CALENDAR_HEATMAP_ASSETS_PATH=assets cat charts/testdata/basic.json | ./calendarheatmap -colorscale=yellow-green-9.csv > docs/colorscale-3.png
	cat charts/testdata/basic.json | ./calendarheatmap -locale=ko_KR > docs/korean.png
	cat charts/testdata/basic.json | ./calendarheatmap -locale=ko_KR -output=svg > docs/korean.svg
	cat charts/testdata/basic.json | ./calendarheatmap -labels=false > docs/nolabels.png
	cat charts/testdata/basic.json | ./calendarheatmap -monthsep=false > docs/noseparator.png
	cat charts/testdata/basic.json | ./calendarheatmap -labels=false -monthsep=false > docs/noseparator_nolabels.png

.PHONY: build test docs