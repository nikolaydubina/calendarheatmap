package charts

const fullyear = `<svg width="722" height="112" xmlns="http://www.w3.org/2000/svg" xmlns:xlink= "http://www.w3.org/1999/xlink">
  <g transform="translate(10, 20)">
	{{range $w, $wo := $.Days}}<g transform="translate({{mul 14 $w}}, 0)">
		{{range $d, $do := $wo}}{{if $do.Show}}<rect class="day" width="10" height="10" x="{{sub 14 $w}}" y="{{mul 13 $d}}" fill="{{$do.Color}}" data-count="{{$do.Count}}" data-date="{{$do.Date}}"></rect>{{end}}
		{{end}}
	</g>	
	{{end}}
	
	{{range $i, $label := $.LabelsMonths}}<text x="{{add 14 (mul 52 $i)}}" y="-7" font-size="10px" fill="{{$.LabelColor}}">{{$label}}</text>
	{{end}}
	 
	{{range $i, $o := $.LabelsWeekdays}}<text text-anchor="start" font-size="9px" dx="-10" dy="{{add 8 (mul 13 $i)}}" fill="{{$.LabelColor}}" {{if not $o.Show}}style="display: none;"{{end}}>{{$o.Label}}</text>
	{{end}}
</g></svg>
`
