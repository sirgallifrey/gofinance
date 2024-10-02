package domain

import (
	"strconv"

	"github.com/uptrace/bun"
)

const FeaturePrefix = "FE"

type Feature struct {
	bun.BaseModel `bun:"table:features,alias:f"`
	Id            string `bun:",pk"`
	ProjectId     string
	Name          string
	Num           int
}

func (f Feature) DisplayNum() string {
	return FeaturePrefix + strconv.Itoa(f.Num)
}
