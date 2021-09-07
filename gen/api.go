package gen

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/render"
)

type ItemStore struct {
	Items  map[Id]Item
	NextId Id
	Lock   sync.Mutex
}

var _ ServerInterface = (*ItemStore)(nil)

func NewItemStore() *ItemStore {
	return &ItemStore{
		Items:  make(map[Id]Item),
		NextId: 1000,
	}
}

func (i *ItemStore) GetItems(w http.ResponseWriter, r *http.Request, params GetItemsParams) {
	i.Lock.Lock()
	defer i.Lock.Unlock()

	var result []Item

	for _, item := range i.Items {
		result = append(result, item)

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				break
			}
		}
	}

	render.JSON(w, r, result)
}

func (i *ItemStore) PostItems(w http.ResponseWriter, r *http.Request) {
	var newItem RequestItem
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Invalid format for PostItem")
		return
	}

	i.Lock.Lock()
	defer i.Lock.Unlock()

	var item Item
	item.Code = newItem.Code
	item.Name = newItem.Name
	item.Unit = newItem.Unit
	item.Cost = newItem.Cost
	item.Created = time.Now()
	item.Updated = time.Now()
	item.Id = i.NextId
	i.NextId = i.NextId + 1

	i.Items[item.Id] = item

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, item)
}

func (i *ItemStore) DeleteItemById(w http.ResponseWriter, r *http.Request, itemId ItemId) {
	var result Item

	render.JSON(w, r, result)
}

func (i *ItemStore) GetItemById(w http.ResponseWriter, r *http.Request, itemId ItemId) {
	var result Item

	render.JSON(w, r, result)
}

func (i *ItemStore) PatchItemById(w http.ResponseWriter, r *http.Request, itemId ItemId) {
	var result Item

	render.JSON(w, r, result)
}
