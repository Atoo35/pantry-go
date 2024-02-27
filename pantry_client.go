package pantry_client

import (
	"github/Atoo35/pantry_client/src/basket"
	"github/Atoo35/pantry_client/src/constants"
	"github/Atoo35/pantry_client/src/pantry"
	"github/Atoo35/pantry_client/src/types"
)

type PantryClient struct {
	BaseURL  string
	PantryID string
	pantry   pantry.PantryI
	basket   basket.BasketI
}

type Option func(*PantryClient)

func NewPantry(opts ...Option) *PantryClient {
	p := &PantryClient{
		BaseURL: constants.BASE_URL,
		pantry:  pantry.Pantry,
		basket:  basket.Basket,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithPantryID(id string) Option {
	return func(p *PantryClient) {
		p.PantryID = id
	}
}

func (p *PantryClient) SetPantryID(id string) {
	p.PantryID = id
}

func (p *PantryClient) GetPantryID() string {
	return p.PantryID
}

func (p *PantryClient) GetBaseURL() string {
	return p.BaseURL
}

func (p *PantryClient) GetPantry() (*types.GetPantryResponse, error) {
	return p.pantry.GetPantry(p.BaseURL, p.PantryID)
}

func (p *PantryClient) UpdatePantryDetails(input *types.UpdatePantryDetailsInput) (*types.GetPantryResponse, error) {
	return p.pantry.UpdatePantryDetails(p.BaseURL, p.PantryID, input)
}

func (p *PantryClient) UpsertBasket(basketName string, data interface{}) error {
	return p.basket.UpsertBasket(p.BaseURL, p.PantryID, basketName, data)
}

func (p *PantryClient) GetContents(basketName string) (map[string]interface{}, error) {
	return p.basket.GetContents(p.BaseURL, p.PantryID, basketName)
}

func (p *PantryClient) UpdateContents(basketName string, data interface{}) error {
	return p.basket.UpdateContents(p.BaseURL, p.PantryID, basketName, data)
}

func (p *PantryClient) DeleteBasket(basketName string) error {
	return p.basket.DeleteBasket(p.BaseURL, p.PantryID, basketName)
}
