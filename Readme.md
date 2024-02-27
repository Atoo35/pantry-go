# Pantry Client in Go

This is a Go client for interacting with the Pantry API.

## Installation

```bash
go get github.com/Atoo35/pantry_client
```

## Usage

First, import the package in your Go file.

```go
import "github.com/Atoo35/pantry_client"
```

Then, create a new client.

```go
client :=  pantry_client.NewPantry(pantry_client.WithPantryID("pantry_id"))
```

## Get Pantry

To get a pantry, use the GetPantry method.

```go
item, err := pantryClient.GetPantry()
if err != nil {
    log.Fatal(err)
}
fmt.Println(item)
```

## Update Pantry Details

To update pantry details, use the UpdatePantryDetails method.

```go
update := &types.UpdatePantryDetailsInput{
		Name:        "pantry_name",
		Description: "pantry_description",
	}
updatedItem, err := pantryClient.UpdatePantryDetails(update)
if err != nil {
    log.Fatal(err)
}
fmt.Println(updatedItem)
```

## Upset Basket

To update basket, use the UpdateBasket method.

```go
basketData := map[string]interface{}{
		"item": "item_name",
		"qty":  1,
	}
err := pantryClient.UpsertBasket("basket_name", basketData)
if err != nil {
    log.Fatal(err)
}
```

## Get Basket Contents

To get contents of the basket, use the GetContents method.

```go
data, err := pantryClient.GetContents("basket_name")
if err != nil {
    log.Fatal(err)
}
fmt.Println(data)
```

## Delete Basket

To delete a basket, use the DeleteBasket method.

```go
err := pantryClient.DeleteBasket("basket_name")
if err != nil {
    log.Fatal(err)
}
```

## Update Basket Contents

To update basket contents, use the UpdateBasketContents method.

```go
basketData := map[string]interface{}{
        "item": "item_name",
        "qty":  1,
    }
err := pantryClient.UpdateBasketContents("basket_name", basketData)
if err != nil {
    log.Fatal(err)
}
```
