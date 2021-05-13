package handlers_test

import (
	"net/http"
	"testing"

	"github.com/adria-stef/gopher-translator-service/pkg/server"
	"github.com/gorilla/mux"

	. "github.com/adria-stef/gopher-translator-service/pkg/handlers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

func createHandlerTestRouter(path string, translator Translator, history History, handleFunc func(response server.Response, request server.Request, translator Translator, history History) error) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
		request := server.Request{
			Request: r,
		}

		response := server.Response{
			ResponseWriter: rw,
		}
		handleFunc(response, request, translator, history)
	}).
		Methods(http.MethodPost)
	return router
}
