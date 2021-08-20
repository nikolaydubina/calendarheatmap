build:
	go build

test:
	go test ./charts

docs: build
	cat charts/testdata/basic.json | ./calendarheatmap > docs/basic.png
	cat charts/testdata/basic.json |  CALENDAR_HEATMAP_ASSETS_PATH=assets ./calendarheatmap -colorscale=purple-blue-9.csv > docs/colorscale-1.png
	cat charts/testdata/basic.json |  CALENDAR_HEATMAP_ASSETS_PATH=assets ./calendarheatmap -colorscale=green-blue-9.csv > docs/colorscale-2.png
	cat charts/testdata/basic.json |  CALENDAR_HEATMAP_ASSETS_PATH=assets ./calendarheatmap -colorscale=yellow-green-9.csv > docs/colorscale-3.png
	cat charts/testdata/basic.json | ./calendarheatmap -locale=ko_KR > docs/korean.png
	cat charts/testdata/basic.json | ./calendarheatmap -locale=ko_KR -output=svg > docs/korean.svg
	cat charts/testdata/basic.json | ./calendarheatmap -labels=false > docs/nolabels.png
	cat charts/testdata/basic.json | ./calendarheatmap -monthsep=false > docs/noseparator.png
	cat charts/testdata/basic.json | ./calendarheatmap -labels=false -monthsep=false > docs/noseparator_nolabels.png

build-web:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" web/
	cp -r assets web/assets
	cd web; GOARCH=wasm GOOS=js go build -o main.wasm main.go

run-web: build-web
	cd web; python3 -m http.server 8000

clean:
	-rm web/wasm_exec.js
	-rm web/main.wasm
	-rm -rf web/assets

.PHONY: build test docs build-web run-web clean
