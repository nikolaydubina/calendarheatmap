package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"syscall/js"
	"time"

	_ "embed"

	"github.com/nikolaydubina/calendarheatmap/charts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed assets/fonts/Sunflower-Medium.ttf
var defaultFontFaceBytes []byte

//go:embed assets/colorscales/green-blue-9.csv
var defaultColorScaleBytes []byte

type Renderer struct {
	config charts.HeatmapConfig
	img    []byte
}

func (r *Renderer) PrettifyJSON(this js.Value, inputs []js.Value) interface{} {
	document := js.Global().Get("document")
	instr := document.Call("getElementById", "inputData").Get("value")
	data := map[string]interface{}{}
	if err := json.Unmarshal([]byte(instr.String()), &data); err == nil {
		if out, err := json.MarshalIndent(data, "", "    "); err == nil {
			document.Call("getElementById", "inputData").Set("value", string(out))
		}
	}
	return nil
}

func (r *Renderer) OnDataChange(this js.Value, inputs []js.Value) interface{} {
	document := js.Global().Get("document")
	instr := document.Call("getElementById", "inputData").Get("value")

	data := map[string]int{}
	if err := json.Unmarshal([]byte(instr.String()), &data); err == nil {
		r.config.Counts = data
		r.Render()
	}

	return nil
}

func (r *Renderer) GetFormatUpdater(format string) func(this js.Value, inputs []js.Value) interface{} {
	return func(this js.Value, inputs []js.Value) interface{} {
		r.config.Format = format
		r.Render()
		return nil
	}
}

func (r *Renderer) GetWeekdaySwitchUpdater(weekday time.Weekday) func(this js.Value, inputs []js.Value) interface{} {
	return func(this js.Value, inputs []js.Value) interface{} {
		r.config.ShowWeekdays[weekday] = !r.config.ShowWeekdays[weekday]
		r.Render()
		return nil
	}
}

func (r *Renderer) ToggleLabels(this js.Value, inputs []js.Value) interface{} {
	r.config.DrawLabels = !r.config.DrawLabels
	r.Render()
	return nil
}

func (r *Renderer) ToggleMonthSeparator(this js.Value, inputs []js.Value) interface{} {
	r.config.DrawMonthSeparator = !r.config.DrawMonthSeparator
	r.Render()
	return nil
}

func (r *Renderer) Render() {
	var output bytes.Buffer
	if err := charts.WriteHeatmap(r.config, &output); err != nil {
		log.Println(err)
		return
	}

	if r.config.Format == "svg" {
		img := output.String()
		js.Global().Get("document").Call("getElementById", "output-container").Set("innerHTML", img)
	} else {
		img := js.Global().Get("document").Call("createElement", "img")
		src := fmt.Sprintf("data:image/%s;base64,%s", r.config.Format, base64.StdEncoding.EncodeToString(output.Bytes()))
		img.Set("src", src)
		img.Set("style", "width: 100%;")

		container := js.Global().Get("document").Call("getElementById", "output-container")
		container.Set("innerHTML", "")
		container.Call("appendChild", img)
	}

	// download file update button
	src := fmt.Sprintf("data:image/%s;base64,%s", r.config.Format, base64.StdEncoding.EncodeToString(output.Bytes()))
	link := js.Global().Get("document").Call("getElementById", "downloadLink")
	link.Set("href", src)
	link.Set("download", fmt.Sprintf("calendar-heatmap.%s", r.config.Format))
}

func main() {
	c := make(chan bool)

	colorscale, _ := charts.NewBasicColorscaleFromCSV(bytes.NewBuffer(defaultColorScaleBytes))
	fontFace, _ := charts.LoadFontFace(defaultFontFaceBytes, opentype.FaceOptions{
		Size:    13,
		DPI:     80,
		Hinting: font.HintingNone,
	})
	renderer := Renderer{
		config: charts.HeatmapConfig{
			Counts:              nil,
			ColorScale:          colorscale,
			DrawMonthSeparator:  true,
			DrawLabels:          true,
			Margin:              5,
			BoxSize:             24,
			MonthSeparatorWidth: 1,
			MonthLabelYOffset:   5,
			TextWidthLeft:       40,
			TextHeightTop:       15,
			TextColor:           color.RGBA{100, 100, 100, 255},
			BorderColor:         color.RGBA{200, 200, 200, 255},
			Locale:              "en_US",
			Format:              "svg",
			FontFace:            fontFace,
			ShowWeekdays: map[time.Weekday]bool{
				time.Monday:    true,
				time.Wednesday: true,
				time.Friday:    true,
			},
		},
	}

	document := js.Global().Get("document")

	document.Call("getElementById", "inputConfig").Call("reset")

	document.Call("getElementById", "inputData").Set("onkeyup", js.FuncOf(renderer.OnDataChange))
	document.Call("getElementById", "btnPrettifyJSON").Set("onclick", js.FuncOf(renderer.PrettifyJSON))

	document.Call("getElementById", "formatSVG").Set("onclick", js.FuncOf(renderer.GetFormatUpdater("svg")))
	document.Call("getElementById", "formatPNG").Set("onclick", js.FuncOf(renderer.GetFormatUpdater("png")))
	document.Call("getElementById", "formatJPEG").Set("onclick", js.FuncOf(renderer.GetFormatUpdater("jpeg")))

	document.Call("getElementById", "switchMon").Set("onchange", js.FuncOf(renderer.GetWeekdaySwitchUpdater(time.Monday)))
	document.Call("getElementById", "switchTue").Set("onchange", js.FuncOf(renderer.GetWeekdaySwitchUpdater(time.Tuesday)))
	document.Call("getElementById", "switchWed").Set("onchange", js.FuncOf(renderer.GetWeekdaySwitchUpdater(time.Wednesday)))
	document.Call("getElementById", "switchThu").Set("onchange", js.FuncOf(renderer.GetWeekdaySwitchUpdater(time.Thursday)))
	document.Call("getElementById", "switchFri").Set("onchange", js.FuncOf(renderer.GetWeekdaySwitchUpdater(time.Friday)))
	document.Call("getElementById", "switchSat").Set("onchange", js.FuncOf(renderer.GetWeekdaySwitchUpdater(time.Saturday)))
	document.Call("getElementById", "switchSun").Set("onchange", js.FuncOf(renderer.GetWeekdaySwitchUpdater(time.Sunday)))

	document.Call("getElementById", "switchLabels").Set("onchange", js.FuncOf(renderer.ToggleLabels))
	document.Call("getElementById", "switchMonthSeparator").Set("onchange", js.FuncOf(renderer.ToggleMonthSeparator))

	renderer.OnDataChange(js.Value{}, nil)
	renderer.PrettifyJSON(js.Value{}, nil)
	renderer.Render()

	<-c
}
