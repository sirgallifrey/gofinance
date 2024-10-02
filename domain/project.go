package domain

import "github.com/uptrace/bun"

type Project struct {
	bun.BaseModel `bun:"table:projects,alias:p"`
	Id            string `bun:",pk"`
	Name          string
}
