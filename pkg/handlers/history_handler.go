package handlers

import (
	"net/http"

	"github.com/adria-stef/gopher-translator-service/pkg/model"
	"github.com/adria-stef/gopher-translator-service/pkg/server"
	"github.com/pkg/errors"
)

type HistoryHandler struct{}

func NewHistoryHandler() *HistoryHandler {
	return &HistoryHandler{}
}

func (h *HistoryHandler) Handle(response server.Response, request server.Request, history History) error {

	if err := response.WriteJSON(http.StatusOK, model.History{
		History: history.GetArranged(),
	}); err != nil {
		return errors.Wrap(err, "failed to write response")
	}
	return nil
}
