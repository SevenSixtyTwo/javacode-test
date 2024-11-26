package structs

import "github.com/google/uuid"

type Wallet struct {
	ID        uuid.UUID `json:"valletid"`
	Operation string    `json:"operationType"`
	Balance   float64   `json:"amount"`
}
