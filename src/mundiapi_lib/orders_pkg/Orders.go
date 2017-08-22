/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package orders_pkg

import "mundiapi_lib/models_pkg"

/*
 * Interface for the ORDERS_IMPL
 */
type ORDERS interface {
    GetOrder (string) (*models_pkg.GetOrderResponse, error)

    GetOrders () (*models_pkg.ListOrderResponse, error)

    CreateOrder (*models_pkg.CreateOrderRequest) (*models_pkg.GetOrderResponse, error)
}

/*
 * Factory for the ORDERS interaface returning ORDERS_IMPL
 */
func NewORDERS() ORDERS {
    return &ORDERS_IMPL{}
}