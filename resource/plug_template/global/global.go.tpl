package global

{{- if .HasGlobal }}

import "github.com/wangrui19970405/wu-shi-admin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}