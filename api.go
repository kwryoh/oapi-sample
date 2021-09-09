package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/render"
	"github.com/kwryoh/oapi-sample/gen/db"
	"github.com/kwryoh/oapi-sample/gen/openapi"
)

type ItemStore struct {
	Items  map[openapi.Id]openapi.Item
	NextId openapi.Id
	Lock   sync.Mutex
}

var _ openapi.ServerInterface = (*ItemStore)(nil)

func NewItemStore() *ItemStore {
	return &ItemStore{
		Items:  make(map[openapi.Id]openapi.Item),
		NextId: 1000,
	}
}

func (i *ItemStore) GetItems(w http.ResponseWriter, r *http.Request, params openapi.GetItemsParams) {
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
	items, err := queries.ListItems(ctx, arg)
	if err != nil {
		log.Fatal("Cannot retrieve items: ", err)
	}

	var result []openapi.Item
	for _, dbitem := range items {
		var item openapi.Item

		j, _ := json.Marshal(dbitem)
		if err := json.Unmarshal(j, &item); err != nil {
			log.Fatal("cannot convert DB to RESTAPI: ", err)
		}
	}

	render.JSON(w, r, result)
}

func (i *ItemStore) PostItems(w http.ResponseWriter, r *http.Request) {
	var newItem openapi.RequestItem
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Invalid format for PostItem")
		return
	}

	i.Lock.Lock()
	defer i.Lock.Unlock()

	var item openapi.Item
	item.Code = newItem.Code
	item.Name = newItem.Name
	item.Unit = newItem.Unit
	item.Cost = newItem.Cost
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	item.Id = i.NextId
	i.NextId = i.NextId + 1

	i.Items[item.Id] = item

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, item)
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
