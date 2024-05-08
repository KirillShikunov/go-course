package functional

import (
	"14_layers/internal/models"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestGetOrders(t *testing.T) {
	test := NewAPITest()
	test.setUp()
	defer func() {
		test.truncateTables([]string{"orders"})
		test.tearDown()
	}()

	order := models.Order{CustomerID: 2, TotalPrice: 100, Status: models.OrderStatusCreated}

	ctx := context.Background()
	err := test.container.OrderRepository().Create(ctx, &order)
	require.NoError(t, err)

	response, err := test.client.Get(test.getAbsolutePath("/orders"))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	responseData, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	err = response.Body.Close()
	require.NoError(t, err)

	var orders []models.Order
	err = json.Unmarshal(responseData, &orders)
	require.NoError(t, err)

	require.Len(t, orders, 1)
	actualOrder := orders[0]
	require.Equal(t, order.CustomerID, actualOrder.CustomerID)
	require.Equal(t, order.TotalPrice, actualOrder.TotalPrice)
	require.Equal(t, order.Status, actualOrder.Status)
}
