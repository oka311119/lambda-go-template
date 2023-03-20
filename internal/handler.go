package internal

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

const (
	ErrMethodNotAllowed = "status method not allowed"
	ErrIDRequired       = "ID must be provided for update"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

type Handler struct {
	repository ItemRepository
}

func NewHandler() *Handler {
	return &Handler{
		repository: NewDdbItemRepository(),
	}
}

// Invoke processes the request and calls the corresponding method.
func (h *Handler) Invoke(ctx context.Context, req Request) (Response, error) {
	switch req.HTTPMethod {
	case http.MethodGet:
		return h.findAll()
	case http.MethodPost:
		return h.create(ctx, req)
	case http.MethodPut:
		return h.update(ctx, req)
	case http.MethodDelete:
		return h.delete(ctx, req.PathParameters["id"])
	default:
		return jsonResponse(http.StatusMethodNotAllowed, ErrMethodNotAllowed)
	}
}

func (h *Handler) findAll() (Response, error) {
	items, err := h.repository.FindAll()
	if err != nil {
		return Response{}, err
	}
	return jsonResponse(http.StatusOK, items)
}

func (h *Handler) create(ctx context.Context, req Request) (Response, error) {
	var item Item
	if err := json.Unmarshal([]byte(req.Body), &item); err != nil {
		return Response{}, err
	}
	item.Id = generateID()
	if err := h.repository.Save(ctx, item); err != nil {
		return Response{}, err
	}
	return jsonResponse(http.StatusCreated, nil)
}

func (h *Handler) update(ctx context.Context, req Request) (Response, error) {
	var item Item
	if err := json.Unmarshal([]byte(req.Body), &item); err != nil {
		return Response{}, err
	}
	if id, ok := req.PathParameters["id"]; ok {
		item.Id = id
	} else {
		return jsonResponse(http.StatusBadRequest, ErrIDRequired)
	}
	if err := h.repository.Save(ctx, item); err != nil {
		return Response{}, err
	}
	return jsonResponse(http.StatusOK, nil)
}

func (h *Handler) delete(ctx context.Context, id string) (Response, error) {
	if err := h.repository.Delete(ctx, id); err != nil {
		return Response{}, err
	}
	return jsonResponse(http.StatusOK, nil)
}

func jsonResponse(statusCode int, body interface{}) (Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return Response{}, err
	}
	return Response{
		StatusCode: statusCode,
		Body:       string(jsonBody),
	}, nil
}

func generateID() string {
	return uuid.New().String()
}
