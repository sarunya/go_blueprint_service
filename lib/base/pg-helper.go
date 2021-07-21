package base

import (
	"bp_service/lib/helper"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type PGHelper struct {
	database string
	table    string
	schema   helper.Schema
	pgx      *pgxpool.Pool
}

func (pg *PGHelper) filter(query string, args interface{}, sc helper.Schema) {
	// sqlscan.Scan
}

func (pg *PGHelper) insert(query string, args interface{}) {

}
