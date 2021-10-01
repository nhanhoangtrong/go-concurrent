package dto

import (
	entities "deals-sync/entities"
)

type DealPayload struct {
	After entities.DealEntity `json:"after"`
}

type DealMessage struct {
	Payload DealPayload `json:"payload"`
}

type SaveDealResult struct {
	Deal *entities.DealEntity
	Err  error
}
