package sql

import (
	sbd "github.com/authink/sqlbuilder"
)

const (
	Id        = sbd.Field("id")
	CreatedAt = sbd.Field("created_at")
	UpdatedAt = sbd.Field("updated_at")
)
