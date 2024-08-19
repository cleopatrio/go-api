package dtos

import "time"

type User struct {
	id         string     `json:"id"`
	name       string     `json:"name"`
	created_at *time.Time `json:"created_at"`
	updated_at *time.Time `json:"updated_at"`
}
