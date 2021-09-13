//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml docs/openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml docs/openapi.yaml
package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/kwryoh/oapi-sample/gen/db"
	"github.com/kwryoh/oapi-sample/gen/openapi"
)

type ItemStore struct {
	queries *db.Queries
	ctx     context.Context
}

var _ openapi.ServerInterface = (*ItemStore)(nil)

func NewItemStore(queries *db.Queries, ctx context.Context) *ItemStore {
	return &ItemStore{
		queries: queries,
		ctx:     ctx,
	}
}

func (i *ItemStore) GetItems(w http.ResponseWriter, r *http.Request, params openapi.GetItemsParams) {
	var result openapi.ResponseItems

	var limit int32 = 10
	if params.Limit != nil {
		limit = int32(*params.Limit)
	}
	var page int32 = 1
	if params.Page != nil {
		page = int32(*params.Page)
	}

	var arg db.ListItemsParams
	arg.Offset = limit * (page - 1)
	arg.Limit = limit

	var items []db.Item
	items, err := i.queries.ListItems(i.ctx, arg)
	if err != nil {
		log.Print("Cannot retrieve items: ", err)
	}

	for _, dbitem := range items {
		cost, err := strconv.ParseFloat(dbitem.Cost, 32)
		if err != nil {
			cost = 0.0
		}

		item := openapi.Item{
			Id:        openapi.Id(dbitem.ID),
			Code:      dbitem.Code,
			Name:      dbitem.Name,
			Unit:      dbitem.Unit,
			CreatedAt: dbitem.CreatedAt,
			UpdatedAt: dbitem.UpdatedAt,
			Cost:      float32(cost),
		}

		result.Items = append(result.Items, item)
	}

	render.JSON(w, r, result)
}

func (i *ItemStore) PostItems(w http.ResponseWriter, r *http.Request) {
	var reqItem openapi.RequestItem
	if err := json.NewDecoder(r.Body).Decode(&reqItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Invalid format for PostItem")
		return
	}

	params := db.CreateItemParams{
		Code: reqItem.Code,
		Name: reqItem.Name,
		Unit: reqItem.Unit,
		Cost: fmt.Sprintf("%g", reqItem.Cost),
	}

	dbItem, err := i.queries.CreateItem(i.ctx, params)
	if err != nil {
		log.Fatal("Could not insert item ", err)
	}

	cost, err := strconv.ParseFloat(dbItem.Cost, 32)
	if err != nil {
		cost = 0.0
	}
	result := openapi.ResponseItem{
		Id:        openapi.Id(dbItem.ID),
		Code:      dbItem.Code,
		Name:      dbItem.Name,
		Unit:      dbItem.Unit,
		Cost:      float32(cost),
		CreatedAt: dbItem.CreatedAt,
		UpdatedAt: dbItem.UpdatedAt,
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, result)
}

func (i *ItemStore) DeleteItemById(w http.ResponseWriter, r *http.Request, itemId openapi.ItemId) {
	var result openapi.Item

	render.JSON(w, r, result)
}

func (i *ItemStore) GetItemById(w http.ResponseWriter, r *http.Request, itemId openapi.ItemId) {
	var result openapi.Item

	render.JSON(w, r, result)
}

func (i *ItemStore) PatchItemById(w http.ResponseWriter, r *http.Request, itemId openapi.ItemId) {
	var result openapi.Item

	render.JSON(w, r, result)
}
