package app

//go:generate go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.2
//go:generate oapi-codegen --config=types.cfg.yaml docs/openapi.yaml
//go:generate oapi-codegen --config=server.cfg.yaml docs/openapi.yaml
import (
	"context"
	"encoding/json"
	"log"
	"net/http"

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
	var result openapi.GetItemsResponse

	arg := params.ToDbParams()
	items, err := i.queries.ListItems(i.ctx, arg)
	if err != nil {
		log.Print("Cannot retrieve items: ", err)
	}

	for _, dbitem := range items {
		item, err := openapi.NewItemFromDbItem(dbitem)
		if err != nil {
			log.Fatal("Cannot convert item: ", err)
		}

		result.Items = append(result.Items, item)
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, result)
}

func (i *ItemStore) PostItems(w http.ResponseWriter, r *http.Request) {
	var reqItem openapi.GetItemsRequest
	if err := json.NewDecoder(r.Body).Decode(&reqItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Invalid format for PostItem")
		return
	}

	params := openapi.NewCreateItemParams(reqItem)
	dbItem, err := i.queries.CreateItem(i.ctx, params)
	if err != nil {
		log.Fatal("Could not insert item ", err)
	}

	item, _ := openapi.NewItemFromDbItem(dbItem)
	resItem := openapi.GetItemResponse(item)

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, resItem)
}

func (i *ItemStore) DeleteItem(w http.ResponseWriter, r *http.Request, itemId openapi.ItemId) {
	var result openapi.Item

	render.JSON(w, r, result)
}

func (i *ItemStore) GetItemById(w http.ResponseWriter, r *http.Request, itemId openapi.ItemId) {
	var result openapi.Item

	render.JSON(w, r, result)
}

func (i *ItemStore) PatchItem(w http.ResponseWriter, r *http.Request, itemId openapi.ItemId) {
	var result openapi.Item

	render.JSON(w, r, result)
}
