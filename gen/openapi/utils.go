package openapi

import (
	"strconv"

	"github.com/kwryoh/oapi-sample/gen/db"
)

func NewItemFromDbItem(dbitem db.Item) (Item, error) {
	cost, err := strconv.ParseFloat(dbitem.Cost, 32)
	if err != nil {
		return Item{}, err
	}

	item := Item{
		Id:        Id(dbitem.ID),
		Code:      dbitem.Code,
		Name:      dbitem.Name,
		Unit:      dbitem.Unit,
		CreatedAt: dbitem.CreatedAt,
		UpdatedAt: dbitem.UpdatedAt,
		Cost:      float32(cost),
	}

	return item, nil
}

func NewCreateItemParams(r PostItemsRequest) db.CreateItemParams {
	result := db.CreateItemParams{
		Code: r.Value.Code,
		Name: r.Value.Name,
		Unit: r.Value.Unit,
		Cost: strconv.FormatFloat(float64(r.Value.Cost), 'f', -1, 32),
	}

	return result
}

func (p GetItemsParams) ToDbParams() db.ListItemsParams {
	var arg db.ListItemsParams
	var limit int32 = 10
	var page int32 = 1

	if p.Limit != nil {
		limit = int32(*p.Limit)
	}

	if p.Page != nil {
		page = int32(*p.Page)
	}

	arg.Offset = limit * (page - 1)
	arg.Limit = limit

	return arg
}
