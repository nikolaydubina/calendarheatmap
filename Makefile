build:
	go build

test:
	go test ./...

docs: build
	cat testdata/basic.json | ./calendarheatmap > docs/basic.png
	cat testdata/basic.json | ./calendarheatmap -colorscale=purple-blue-9.csv > docs/colorscale-1.png
	cat testdata/basic.json | ./calendarheatmap -colorscale=green-blue-9.csv > docs/colorscale-2.png
	cat testdata/basic.json | ./calendarheatmap -colorscale=yellow-green-9.csv > docs/colorscale-3.png
	cat testdata/basic.json | ./calendarheatmap -locale=ko_KR > docs/korean.png
	cat testdata/basic.json | ./calendarheatmap -locale=ko_KR -output=svg > docs/korean.svg
	cat testdata/basic.json | ./calendarheatmap -labels=false > docs/nolabels.png
	cat testdata/basic.json | ./calendarheatmap -monthsep=false > docs/noseparator.png
	cat testdata/basic.json | ./calendarheatmap -labels=false -monthsep=false > docs/noseparator_nolabels.png

.PHONY: build test docs