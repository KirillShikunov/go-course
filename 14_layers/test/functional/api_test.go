package functional

import (
	"14_layers/internal/api"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
)

type APITest struct {
	client *http.Client
	server *httptest.Server
}

func (t *APITest) setUp() {
	router := mux.NewRouter()
	api.RegisterRoutes(router)

	t.client = http.DefaultClient
	t.server = httptest.NewServer(router)
}

func (t *APITest) tearDown() {
	t.server.Close()
}

func (t *APITest) getAbsolutePath(path string) string {
	return t.server.URL + path
}

func NewAPITest() *APITest {
	return &APITest{}
}
