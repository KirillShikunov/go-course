package functional

import (
	"14_layers/internal/dto"
	"14_layers/internal/models"
	"14_layers/internal/repositories"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestGetOrders(t *testing.T) {
	test := NewAPITest()
	test.setUp()
	defer test.tearDown()

	order := models.Order{Name: "Order #1", UserID: 1}
	repositories.NewOrderRepository().Create(&order)

	response, err := test.client.Get(test.getAbsolutePath("/orders"))
	require.NoError(t, err)
	require.Equal(t, response.StatusCode, http.StatusOK)

	expectedOrders, err := json.Marshal([]dto.Order{
		{
			ID:     order.ID,
			Name:   order.Name,
			UserID: order.UserID,
		},
	})
	require.NoError(t, err)

	actualOrders, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	assert.JSONEq(t, string(expectedOrders), string(actualOrders))
}
