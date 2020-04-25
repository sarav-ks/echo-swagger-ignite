// Package model provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"gwapp/model"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// (GET /events)
	FindEvents(ctx echo.Context, params model.FindEventsParams) error
	// (POST /events)
	AddEvent(ctx echo.Context) error
	// (DELETE /events/{id})
	DeleteEvent(ctx echo.Context, id int64) error
	// (GET /events/{id})
	FindEventById(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindEvents converts echo context to params.
func (w *ServerInterfaceWrapper) FindEvents(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params model.FindEventsParams
	// ------------- Optional query parameter "tags" -------------
	if paramValue := ctx.QueryParam("tags"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := ctx.QueryParam("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindEvents(ctx, params)
	return err
}

// AddEvent converts echo context to params.
func (w *ServerInterfaceWrapper) AddEvent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddEvent(ctx)
	return err
}

// DeleteEvent converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteEvent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteEvent(ctx, id)
	return err
}

// FindEventById converts echo context to params.
func (w *ServerInterfaceWrapper) FindEventById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindEventById(ctx, id)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/events", wrapper.FindEvents)
	router.POST("/events", wrapper.AddEvent)
	router.DELETE("/events/:id", wrapper.DeleteEvent)
	router.GET("/events/:id", wrapper.FindEventById)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RWTW/cNhD9K8S0R1XaOkEPujnxtligSIumRQ+pD7Q40rIRPzwcebMw9N8LkpI/VrKD",
	"pEURICdJ/Bi+N/PmUbfQOOOdRcsB6lsIzR6NTK9bIkfxxZPzSKwxDTdOYXwqDA1pz9pZqPNikeYKaB0Z",
	"yVCDtvziDArgo8f8iR0SjAUYDEF2Twaap++2BiZtOxjHAgivB02ooH4H04Hz8suxgO0NWl7C1mp51h9W",
	"Xw8otBKuFbxHkfc+JvDDy1UCvWtkDnMa9fejxzkgTgFPWBRgpVkh/0aaj29l2X3WoSepSwguxzisbeuW",
	"Ic9/3QXROhI/Ib4XfyK+74/i3HsooNcN2pAIZCJw7mWzR3FWbqCAgXqoYc/sQ11Vh8OhlGm6dNRV095Q",
	"/bx7vX3zdvvdWbkp92z6xA3JhF/at0g3usEpSF1V4SC7DqnUrkpLqkhPcx+XPIAHBdwghYz/+3JTbmJQ",
	"59FKr6GGF2moAC95n1RRpVyl1w55mYPfkAeyQci+z2kNoiVnUprDMTAa8ZeFdAQlOewU1PCjtmqbA8ez",
	"SBpkpAD1u9P4LLsg2IlW94wkriIDHSeuB6T4MaU3roNi6s4kZ0aTYC/lkQckkTzG78DHlKaoaRiLUwRG",
	"ftBmMMIO5gopSogwDD0nWJToP4Gp10bzI1Af7fvxMmoweBcFEHecbTbZUyxPXSu973VurervkPtrhfa3",
	"hC3U8E11b1/V5F1V7uJFKqLSH3NPFRUzIEjzrRx6/iRMz0JJHrpy9GDxg8eGUQm8X+NdWBHha0LJGIQU",
	"Fg9ZhkLbrEF2hKUQF0NGGFcRCus4StYdUC3Eea7UbHPRDjDwK6eO/x3jnPwl4zQRRSWVio87+PDQl5gG",
	"HP+lSD4L3hemhbGYvam61WrMouiRVy6NPB7lEbTt+sn/xZUMqITLOtldiDBE5CuCuEj7Z008a1e7i2gQ",
	"OFdyQjTZQzTVe3fQalHXp6xi9YZdsYqXS+4ZSsahvoAGfv4SsaeluSvZ7uLpS+TVcac+sTAtcrP/3+ry",
	"FXZnvFeRbuZi3P/w1FXVHaT35YM/lvjvMV6O/wQAAP//WTSRW2sLAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
