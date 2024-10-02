package domain

import "github.com/uptrace/bun"

type Organization struct {
	bun.BaseModel `bun:"table:organizations,alias:o"`
	Id            string `bun:",pk"`
	Name          string
}
