package functional

import (
	"testing"
)

func TestGetOrders(t *testing.T) {
	//test := NewAPITest()
	//test.setUp()
	//defer test.tearDown()
	//
	//ctx := context.Background()
	//
	//order := models.Order{CustomerID: 2, TotalPrice: 100, Status: models.OrderStatusCreated}
	//repositories.NewOrderRepository().Create(ctx, &order)
	//
	//response, err := test.client.Get(test.getAbsolutePath("/orders"))
	//require.NoError(t, err)
	//require.Equal(t, response.StatusCode, http.StatusOK)
	//
	//expectedOrders, err := json.Marshal([]dto.Order{
	//	{
	//		ID:     order.ID,
	//		Name:   order.Name,
	//		UserID: order.UserID,
	//	},
	//})
	//require.NoError(t, err)
	//
	//actualOrders, err := io.ReadAll(response.Body)
	//require.NoError(t, err)
	//
	//assert.JSONEq(t, string(expectedOrders), string(actualOrders))
}
