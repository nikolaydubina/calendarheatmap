package charts

const fullyear = `<svg width="752" height="112" xmlns="http://www.w3.org/2000/svg" xmlns:xlink= "http://www.w3.org/1999/xlink">
  <g transform="translate(10, 20)">
	{{range $w, $wo := $.Days}}<g transform="translate({{mul 14 $w}}, 0)">
		{{range $d, $do := $wo}}{{if $do.Show}}<rect class="day" width="11" height="11" x="0" y="{{mul 13 $d}}" fill="{{$do.Color}}" data-count="{{$do.Count}}" data-date="{{$do.Date}}"></rect>{{end}}
		{{end}}
	</g>	
	{{end}}

	{{range $i, $label := $.LabelsMonths}}<text x="{{mul 14 $label.XOffset}}" y="-7" font-size="8px" fill="{{$.LabelsColor}}">{{$label.Label}}</text>
	{{end}}
	 
	{{range $i, $o := $.LabelsWeekdays}}<text text-anchor="start" font-size="8px" dx="-10" dy="{{add 8 (mul 13 $i)}}" fill="{{$.LabelsColor}}" {{if not $o.Show}}style="display: none;"{{end}}>{{$o.Label}}</text>
	{{end}}
</g></svg>
`
