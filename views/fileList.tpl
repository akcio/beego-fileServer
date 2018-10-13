{{.Val}}


{{range $key, $val := .Files}}
    {{$val.Stored}}
{{end}}
