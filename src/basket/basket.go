package basket

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Atoo35/pantry_client/src/utils"
)

type BasketI interface {
	UpsertBasket(baseURL, pantryID, basketName, data interface{}) error
	GetContents(baseURL, pantryID, basketName string) (map[string]interface{}, error)
	UpdateContents(baseURL, pantryID, basketName string, data interface{}) error
	DeleteBasket(baseURL, pantryID, basketName string) error
}

type basket struct {
	client *http.Client
}

var Basket BasketI = &basket{
	client: &http.Client{},
}

func (b *basket) UpsertBasket(baseURL, pantryID, basketName, data interface{}) error {
	reqBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := utils.GetRequest(http.MethodPost, fmt.Sprintf("%s/pantry/%s/basket/%s", baseURL, pantryID, basketName), strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}

	resp, err := b.client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(respBody))
	}
	return nil
}

func (b *basket) GetContents(baseURL, pantryID, basketName string) (map[string]interface{}, error) {
	request, err := utils.GetRequest(http.MethodGet, fmt.Sprintf("%s/pantry/%s/basket/%s", baseURL, pantryID, basketName), nil)
	if err != nil {
		return nil, err
	}

	resp, err := b.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBody))
	}
	var response map[string]interface{}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (b *basket) UpdateContents(baseURL, pantryID, basketName string, data interface{}) error {
	reqBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := utils.GetRequest(http.MethodPut, fmt.Sprintf("%s/pantry/%s/basket/%s", baseURL, pantryID, basketName), strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}

	resp, err := b.client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(respBody))
	}
	return nil
}

func (b *basket) DeleteBasket(baseURL, pantryID, basketName string) error {
	request, err := utils.GetRequest(http.MethodDelete, fmt.Sprintf("%s/pantry/%s/basket/%s", baseURL, pantryID, basketName), nil)
	if err != nil {
		return err
	}

	resp, err := b.client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(respBody))
	}

	return nil
}
