package db

import (
	sbd "github.com/authink/sqlbuilder"
)

type Table interface {
	Tname() sbd.Table
}
