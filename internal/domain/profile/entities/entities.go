package entities

import "github.com/google/uuid"

type Profile struct {
	ID        uuid.UUID
	Username  string
	Bio       string
	Image     string
	Following bool
}
