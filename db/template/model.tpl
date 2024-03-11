// Package db Code generated by authink/orm. DO NOT EDIT
package db

import (
	{{range .Imports}}"{{.}}"{{end}}
	"github.com/authink/orm/db"
	sbd "github.com/authink/sqlbuilder"
)

type {{.Name}} struct {
	{{if .AtEmbed}}
		{{.EmbedName}}
	{{end}}
	{{range .Fields}}
		{{.}} sbd.Field
	{{end}}
}

{{if .AtDB}}
// Tname implements db.Table.
func (*{{.Name}}) Tname() sbd.Table {
	return "{{.Tname}}"
}

var _ db.Table = (*{{.Name}})(nil)
{{end}}

var {{.Model}} {{.Name}}

func init() {
	db.Bind(&{{.Model}}, &models.{{.Model}}{})
}
