package domain

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	Id            string `bun:",pk"`
	Name          string
	Email         string
	PasswordHash  string
}
