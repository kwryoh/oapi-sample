// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

import (
	"time"
)

const (
	APIKeyScopes = "APIKey.Scopes"
)

// 商品登録時の商品情報
type GetItemsRequest struct {
	// 商品コード
	Code string `json:"code"`

	// 原価
	Cost float32 `json:"cost"`

	// 商品名
	Name string `json:"name"`

	// 単位
	Unit string `json:"unit"`
}

// 主キー型
type Id uint64

// 商品モデル
type Item struct {
	// 商品コード
	Code string `json:"code"`

	// 原価
	Cost float32 `json:"cost"`

	// 作成日時
	CreatedAt time.Time `json:"created_at"`

	// 主キー型
	Id Id `json:"id"`

	// 商品名
	Name string `json:"name"`

	// 単位
	Unit string `json:"unit"`

	// 更新日時
	UpdatedAt time.Time `json:"updated_at"`
}

// 主キー型
type ItemId Id

// Limit defines model for limit.
type Limit int

// Page defines model for page.
type Page int

// ErrorResponse defines model for errorResponse.
type ErrorResponse struct {
	// HTTP status code
	Code int64 `json:"code"`

	// Error message
	Message string `json:"message"`
}

// 商品モデル
type GetItemResponse Item

// GetItemsResponse defines model for getItemsResponse.
type GetItemsResponse struct {
	Items []Item `json:"items"`
}

// GetItemsParams defines parameters for GetItems.
type GetItemsParams struct {
	// ページ数
	Limit *Limit `json:"limit,omitempty"`

	// 閲覧ページ
	Page *Page `json:"page,omitempty"`
}

// PostItemsJSONBody defines parameters for PostItems.
type PostItemsJSONBody GetItemsRequest

// PostItemsJSONRequestBody defines body for PostItems for application/json ContentType.
type PostItemsJSONRequestBody PostItemsJSONBody
