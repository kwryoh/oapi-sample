// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// GetItems
	// (GET /items)
	GetItems(w http.ResponseWriter, r *http.Request, params GetItemsParams)
	// PostItems
	// (POST /items)
	PostItems(w http.ResponseWriter, r *http.Request)
	// DeleteItemById
	// (DELETE /items/{item_id})
	DeleteItemById(w http.ResponseWriter, r *http.Request, itemId ItemId)
	// Get Item
	// (GET /items/{item_id})
	GetItemById(w http.ResponseWriter, r *http.Request, itemId ItemId)
	// PatchItemById
	// (PATCH /items/{item_id})
	PatchItemById(w http.ResponseWriter, r *http.Request, itemId ItemId)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetItems operation middleware
func (siw *ServerInterfaceWrapper) GetItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetItemsParams

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter limit: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetItems(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostItems operation middleware
func (siw *ServerInterfaceWrapper) PostItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostItems(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteItemById operation middleware
func (siw *ServerInterfaceWrapper) DeleteItemById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "item_id" -------------
	var itemId ItemId

	err = runtime.BindStyledParameter("simple", false, "item_id", chi.URLParam(r, "item_id"), &itemId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter item_id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteItemById(w, r, itemId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetItemById operation middleware
func (siw *ServerInterfaceWrapper) GetItemById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "item_id" -------------
	var itemId ItemId

	err = runtime.BindStyledParameter("simple", false, "item_id", chi.URLParam(r, "item_id"), &itemId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter item_id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetItemById(w, r, itemId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PatchItemById operation middleware
func (siw *ServerInterfaceWrapper) PatchItemById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "item_id" -------------
	var itemId ItemId

	err = runtime.BindStyledParameter("simple", false, "item_id", chi.URLParam(r, "item_id"), &itemId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter item_id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchItemById(w, r, itemId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/items", wrapper.GetItems)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/items", wrapper.PostItems)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/items/{item_id}", wrapper.DeleteItemById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/items/{item_id}", wrapper.GetItemById)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/items/{item_id}", wrapper.PatchItemById)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RXW2/cVBD+K6sBiQec2CkRUv1E2pRqhSBR4YloqU7s2V23to9zfNx2tfJDdsWlQKGi",
	"akoAcROXB0RVCg+VevkzGy/kX6DxsXd9W9KkIF6iPZs5M998883M2SFY3Au4j74MwRxCwATzUKJIT45E",
	"76Jj00cbQ0s4gXS4DyYkt99Lbu2210EDh84Bk33QwGcegjm7poHAncgRaIMpRYQahFYfPUb+nhfYBROe",
	"0+fxdfXfUG/bEMcauI7nyHrsyfiLyfjRZPRgevteHn8nQjGYA1A3i+Fs7LLIlWCuGBrIQZDC9CX2UKSx",
	"AtbDeqjDvft//fTzLOCCaOnd5mANsWJiJQy4H2LK8YXs0Jbo0dnivkQ/zZsFgetYjMDol0JCNHxaBslZ",
	"Gquc0LlrzAtcbOUIINZKAMJjIQgED1BIB2diKX84GiFxwewN3x3kCsn4YkKwASiycgltZZ47Myu+fQkt",
	"+XSJxnmFUmjtBlEfPHg4Gf06GT9Kvv4INOhy4TGpavfyKjRAdaRLKNqk9EqdNbi21ONL2beRchFrkFe5",
	"qZ0m4+8n4/cn419Aq1BrcRsX3hr9TuocXwcNUGUNJpzdWD9nGCvkiUmJguzf2VpbertDf4yl050X56BD",
	"KRy/R/AsHsqSgo0qsckn3xw8+a5IT9flTM6d+ZG3TULXwBLIJNoXWUMTHzz+avrBzemdH6f7o6Izm0lc",
	"ko6HsFAac7RqMh01SfI+LWQFkW9j1/GRKtdEanLzRhM9kd80kJIbnx88brYP7IUMTL/8Y7p375kYqHYH",
	"ZZNKJUs5A5yVtVSQErZOQcukz2p/kZgzaZEcYzpL1gsLLZnOkZ0IQ/lPCv9z/+Hhx/en+6PJ7l31zXT8",
	"bvLtbycW/AnFfUJBz3X0HymmUtDFtSxUrEh742B0/C7PxzqzUjToMccFEy6zq8wRA95/pUffLFvcm6+1",
	"CwPeb71GFhRckH1fyiA0db3nyH60Teb65at0H2oDuC1fCFuhGsJrm+3WOrcij7YK7XULaSSbwzzU6+23",
	"yEOe0pvqGkmrtRGgv7bZBg2uoAiV75Vlg8x5gD4LHDDhpWVj2VBy6Kf60Wc7qIdyUcmU9iajz5JP95In",
	"dyD1KNJlR+sBzqNUO1ErPYu2mkfO3ERXz49YO9IwfTnEncqL4JRhLJprMzu9vLXT7RZ5HhODMvB6mwbN",
	"XVBkZPeuatQaI5s8nHkWSndnuD34154tRS3H9afSKWPleMRUeCnCrxITa5lo9GH2fI0VSy5KPFpB1NC3",
	"kusfHu7/UGNtPfVBgc8M2upVXEpqte7+Dd46mxEaa7DabCJbr/LItytJ1qLVJfDsPdGciHHc6hwztfMo",
	"W/mcq+v6WC2a/0Sh5guYtPpHEqL2db0l6PL/R0k1fE3WZI3iSk7LfI6buu5yi7l9HkrztGEYKRfZ/WHx",
	"p1wIcSf+OwAA///PvJbRKQ4AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
