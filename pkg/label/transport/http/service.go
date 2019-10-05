package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	_ "github.com/eiji03aero/todos-los-dias/pkg/label"
	"github.com/eiji03aero/todos-los-dias/pkg/label/transport"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func NewService(
	svgEndpoints transport.Endpoints, logger log.Logger,
) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeErrorResponse),
	}

	r.Methods("POST").Path("/labels").Handler(kithttp.NewServer(
		svgEndpoints.CreateLabel,
		decodeCreateLabelRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeCreateLabelRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateLabelRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Label); e != nil {
		return nil, e
	}
	return req, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
