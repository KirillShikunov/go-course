package functional

import (
	"14_layers/internal/app"
	_ "14_layers/internal/config"
	"14_layers/internal/di"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
)

type APITest struct {
	client       *http.Client
	server       *httptest.Server
	container    *di.ServiceContainer
	dbConnection *gorm.DB
}

func (t *APITest) setUp() {
	router := mux.NewRouter()

	container := di.NewServiceContainer()
	app.NewApp(container.Load()).RegisterRoutes(router)

	t.client = http.DefaultClient
	t.server = httptest.NewServer(router)
	t.container = container
	t.dbConnection = container.DbConnection()
}

func (t *APITest) tearDown() {
	t.server.Close()
}

func (t *APITest) getAbsolutePath(path string) string {
	return t.server.URL + path
}

func (t *APITest) truncateTables(tables []string) {
	for _, table := range tables {
		t.dbConnection.Exec(fmt.Sprintf("TRUNCATE %s", table))
	}
}

func NewAPITest() *APITest {
	return &APITest{}
}
