package pantry

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Atoo35/pantry_client/src/types"
	"github.com/Atoo35/pantry_client/src/utils"
)

type PantryI interface {
	GetPantry(baseURL, pantryID string) (*types.GetPantryResponse, error)
	UpdatePantryDetails(baseURL, pantryID string, input *types.UpdatePantryDetailsInput) (*types.GetPantryResponse, error)
}

type pantry struct {
	client *http.Client
}

var Pantry PantryI = &pantry{
	client: &http.Client{},
}

func (p *pantry) GetPantry(baseURL, pantryID string) (*types.GetPantryResponse, error) {
	request, err := utils.GetRequest(http.MethodGet, fmt.Sprintf("%s/pantry/%s", baseURL, pantryID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.New(string(respBody))
	}

	var pantry types.GetPantryResponse
	err = json.Unmarshal(respBody, &pantry)
	if err != nil {
		return nil, err
	}

	return &pantry, nil
}

func (p *pantry) UpdatePantryDetails(baseURL, pantryID string, input *types.UpdatePantryDetailsInput) (*types.GetPantryResponse, error) {
	requestBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	request, err := utils.GetRequest(http.MethodPut, fmt.Sprintf("%s/pantry/%s", baseURL, pantryID), strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.New(string(respBody))
	}

	var pantry types.GetPantryResponse
	err = json.Unmarshal(respBody, &pantry)
	if err != nil {
		return nil, err
	}

	return &pantry, nil
}
