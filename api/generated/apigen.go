// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package generated

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ArticleDto defines model for Article.
type ArticleDto struct {
	CreatedAt   *time.Time         `json:"created_at,omitempty"`
	Description *string            `json:"description,omitempty"`
	Id          openapi_types.UUID `json:"id"`
	Title       string             `json:"title"`
	UserId      openapi_types.UUID `json:"user_id"`
}

// CreateUserDto defines model for CreateUser.
type CreateUserDto struct {
	CreatedAt time.Time          `json:"created_at"`
	Email     string             `json:"email"`
	FirstName string             `json:"first_name"`
	Id        openapi_types.UUID `json:"id"`
	LastName  string             `json:"last_name"`
}

// UserDto defines model for User.
type UserDto struct {
	CreatedAt time.Time          `json:"created_at"`
	Email     *string            `json:"email,omitempty"`
	FirstName *string            `json:"first_name,omitempty"`
	Id        openapi_types.UUID `json:"id"`
	LastName  *string            `json:"last_name,omitempty"`
}

// CreateArticleJSONRequestBody defines body for CreateArticle for application/json ContentType.
type CreateArticleJSONRequestBody = ArticleDto

// UpdateArticleJSONRequestBody defines body for UpdateArticle for application/json ContentType.
type UpdateArticleJSONRequestBody = ArticleDto

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = CreateUserDto

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UserDto

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /articles)
	GetArticles(w http.ResponseWriter, r *http.Request)

	// (POST /articles)
	CreateArticle(w http.ResponseWriter, r *http.Request)

	// (GET /articles/{id})
	GetArticlesId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (PUT /articles/{id})
	UpdateArticle(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (GET /users)
	GetUsers(w http.ResponseWriter, r *http.Request)

	// (POST /users)
	CreateUser(w http.ResponseWriter, r *http.Request)

	// (GET /users/{id})
	GetUsersId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (PUT /users/{id})
	UpdateUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetArticles operation middleware
func (siw *ServerInterfaceWrapper) GetArticles(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetArticles(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateArticle operation middleware
func (siw *ServerInterfaceWrapper) CreateArticle(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateArticle(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetArticlesId operation middleware
func (siw *ServerInterfaceWrapper) GetArticlesId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", mux.Vars(r)["id"], &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetArticlesId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateArticle operation middleware
func (siw *ServerInterfaceWrapper) UpdateArticle(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", mux.Vars(r)["id"], &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateArticle(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsers(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUsersId operation middleware
func (siw *ServerInterfaceWrapper) GetUsersId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", mux.Vars(r)["id"], &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", mux.Vars(r)["id"], &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/articles", wrapper.GetArticles).Methods("GET")

	r.HandleFunc(options.BaseURL+"/articles", wrapper.CreateArticle).Methods("POST")

	r.HandleFunc(options.BaseURL+"/articles/{id}", wrapper.GetArticlesId).Methods("GET")

	r.HandleFunc(options.BaseURL+"/articles/{id}", wrapper.UpdateArticle).Methods("PUT")

	r.HandleFunc(options.BaseURL+"/users", wrapper.GetUsers).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users", wrapper.CreateUser).Methods("POST")

	r.HandleFunc(options.BaseURL+"/users/{id}", wrapper.GetUsersId).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users/{id}", wrapper.UpdateUser).Methods("PUT")

	return r
}

type GetArticlesRequestObject struct {
}

type GetArticlesResponseObject interface {
	VisitGetArticlesResponse(w http.ResponseWriter) error
}

type GetArticles200JSONResponse []ArticleDto

func (response GetArticles200JSONResponse) VisitGetArticlesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateArticleRequestObject struct {
	Body *CreateArticleJSONRequestBody
}

type CreateArticleResponseObject interface {
	VisitCreateArticleResponse(w http.ResponseWriter) error
}

type CreateArticle201JSONResponse ArticleDto

func (response CreateArticle201JSONResponse) VisitCreateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type GetArticlesIdRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type GetArticlesIdResponseObject interface {
	VisitGetArticlesIdResponse(w http.ResponseWriter) error
}

type GetArticlesId200JSONResponse ArticleDto

func (response GetArticlesId200JSONResponse) VisitGetArticlesIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateArticleRequestObject struct {
	Id   openapi_types.UUID `json:"id"`
	Body *UpdateArticleJSONRequestBody
}

type UpdateArticleResponseObject interface {
	VisitUpdateArticleResponse(w http.ResponseWriter) error
}

type UpdateArticle200JSONResponse ArticleDto

func (response UpdateArticle200JSONResponse) VisitUpdateArticleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []UserDto

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateUserRequestObject struct {
	Body *CreateUserJSONRequestBody
}

type CreateUserResponseObject interface {
	VisitCreateUserResponse(w http.ResponseWriter) error
}

type CreateUser201JSONResponse CreateUserDto

func (response CreateUser201JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersIdRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type GetUsersIdResponseObject interface {
	VisitGetUsersIdResponse(w http.ResponseWriter) error
}

type GetUsersId200JSONResponse UserDto

func (response GetUsersId200JSONResponse) VisitGetUsersIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateUserRequestObject struct {
	Id   openapi_types.UUID `json:"id"`
	Body *UpdateUserJSONRequestBody
}

type UpdateUserResponseObject interface {
	VisitUpdateUserResponse(w http.ResponseWriter) error
}

type UpdateUser200JSONResponse UserDto

func (response UpdateUser200JSONResponse) VisitUpdateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /articles)
	GetArticles(ctx context.Context, request GetArticlesRequestObject) (GetArticlesResponseObject, error)

	// (POST /articles)
	CreateArticle(ctx context.Context, request CreateArticleRequestObject) (CreateArticleResponseObject, error)

	// (GET /articles/{id})
	GetArticlesId(ctx context.Context, request GetArticlesIdRequestObject) (GetArticlesIdResponseObject, error)

	// (PUT /articles/{id})
	UpdateArticle(ctx context.Context, request UpdateArticleRequestObject) (UpdateArticleResponseObject, error)

	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)

	// (POST /users)
	CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error)

	// (GET /users/{id})
	GetUsersId(ctx context.Context, request GetUsersIdRequestObject) (GetUsersIdResponseObject, error)

	// (PUT /users/{id})
	UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetArticles operation middleware
func (sh *strictHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	var request GetArticlesRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetArticles(ctx, request.(GetArticlesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetArticles")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetArticlesResponseObject); ok {
		if err := validResponse.VisitGetArticlesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateArticle operation middleware
func (sh *strictHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var request CreateArticleRequestObject

	var body CreateArticleJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateArticle(ctx, request.(CreateArticleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateArticle")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateArticleResponseObject); ok {
		if err := validResponse.VisitCreateArticleResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetArticlesId operation middleware
func (sh *strictHandler) GetArticlesId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request GetArticlesIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetArticlesId(ctx, request.(GetArticlesIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetArticlesId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetArticlesIdResponseObject); ok {
		if err := validResponse.VisitGetArticlesIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateArticle operation middleware
func (sh *strictHandler) UpdateArticle(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request UpdateArticleRequestObject

	request.Id = id

	var body UpdateArticleJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateArticle(ctx, request.(UpdateArticleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateArticle")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(UpdateArticleResponseObject); ok {
		if err := validResponse.VisitUpdateArticleResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var request GetUsersRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx, request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		if err := validResponse.VisitGetUsersResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateUser operation middleware
func (sh *strictHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequestObject

	var body CreateUserJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUser(ctx, request.(CreateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateUserResponseObject); ok {
		if err := validResponse.VisitCreateUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetUsersId operation middleware
func (sh *strictHandler) GetUsersId(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request GetUsersIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersId(ctx, request.(GetUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetUsersIdResponseObject); ok {
		if err := validResponse.VisitGetUsersIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateUser operation middleware
func (sh *strictHandler) UpdateUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request UpdateUserRequestObject

	request.Id = id

	var body UpdateUserJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateUser(ctx, request.(UpdateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateUser")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(UpdateUserResponseObject); ok {
		if err := validResponse.VisitUpdateUserResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
