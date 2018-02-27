/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package customers_pkg

import "mundiapi_lib/models_pkg"

/*
 * Interface for the CUSTOMERS_IMPL
 */
type CUSTOMERS interface {
    UpdateCard (string, string, *models_pkg.UpdateCardRequest) (*models_pkg.GetCardResponse, error)

    UpdateAddress (string, string, *models_pkg.UpdateAddressRequest) (*models_pkg.GetAddressResponse, error)

    CreateCustomer (*models_pkg.CreateCustomerRequest) (*models_pkg.GetCustomerResponse, error)

    GetCustomer (string) (*models_pkg.GetCustomerResponse, error)

    GetAccessTokens (string, *int64, *int64) (*models_pkg.ListAccessTokensResponse, error)

    GetAddresses (string, *int64, *int64) (*models_pkg.ListAddressesResponse, error)

    GetCards (string, *int64, *int64) (*models_pkg.ListCardsResponse, error)

    DeleteAccessTokens (string) (*models_pkg.ListAccessTokensResponse, error)

    GetAccessToken (string, string) (*models_pkg.GetAccessTokenResponse, error)

    CreateAccessToken (string, *models_pkg.CreateAccessTokenRequest) (*models_pkg.GetAccessTokenResponse, error)

    DeleteAccessToken (string, string) (*models_pkg.GetAccessTokenResponse, error)

    UpdateCustomerMetadata (string, *models_pkg.UpdateMetadataRequest) (*models_pkg.GetCustomerResponse, error)

    UpdateCustomer (string, *models_pkg.UpdateCustomerRequest) (*models_pkg.GetCustomerResponse, error)

    GetAddress (string, string) (*models_pkg.GetAddressResponse, error)

    DeleteAddress (string, string) (*models_pkg.GetAddressResponse, error)

    DeleteCard (string, string) (*models_pkg.GetCardResponse, error)

    CreateAddress (string, *models_pkg.CreateAddressRequest) (*models_pkg.GetAddressResponse, error)

    GetCard (string, string) (*models_pkg.GetCardResponse, error)

    CreateCard (string, *models_pkg.CreateCardRequest) (*models_pkg.GetCardResponse, error)

    GetCustomers (string, *string, *string, *int64, *int64, *string) (*models_pkg.ListCustomersResponse, error)
}

/*
 * Factory for the CUSTOMERS interaface returning CUSTOMERS_IMPL
 */
func NewCUSTOMERS() CUSTOMERS {
    return &CUSTOMERS_IMPL{}
}