/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package customers_pkg

import "mundiapi_lib/configuration_pkg"
import "mundiapi_lib/models_pkg"

/*
 * Interface for the CUSTOMERS_IMPL
 */
type CUSTOMERS interface {
    CreateAccessToken (string, *models_pkg.CreateAccessTokenRequest, *string) (*models_pkg.GetAccessTokenResponse, error)

    UpdateCustomer (string, *models_pkg.UpdateCustomerRequest, *string) (*models_pkg.GetCustomerResponse, error)

    DeleteAccessTokens (string) (*models_pkg.ListAccessTokensResponse, error)

    GetCustomer (string) (*models_pkg.GetCustomerResponse, error)

    GetAddresses (string, *int64, *int64) (*models_pkg.ListAddressesResponse, error)

    GetAccessToken (string, string) (*models_pkg.GetAccessTokenResponse, error)

    GetAddress (string, string) (*models_pkg.GetAddressResponse, error)

    CreateCard (string, *models_pkg.CreateCardRequest, *string) (*models_pkg.GetCardResponse, error)

    RenewCard (string, string, *string) (*models_pkg.GetCardResponse, error)

    CreateCustomer (*models_pkg.CreateCustomerRequest, *string) (*models_pkg.GetCustomerResponse, error)

    UpdateAddress (string, string, *models_pkg.UpdateAddressRequest, *string) (*models_pkg.GetAddressResponse, error)

    GetAccessTokens (string, *int64, *int64) (*models_pkg.ListAccessTokensResponse, error)

    UpdateCustomerMetadata (string, *models_pkg.UpdateMetadataRequest, *string) (*models_pkg.GetCustomerResponse, error)

    DeleteAddress (string, string, *string) (*models_pkg.GetAddressResponse, error)

    UpdateCard (string, string, *models_pkg.UpdateCardRequest, *string) (*models_pkg.GetCardResponse, error)

    GetCards (string, *int64, *int64) (*models_pkg.ListCardsResponse, error)

    DeleteCard (string, string, *string) (*models_pkg.GetCardResponse, error)

    GetCustomers (*string, *string, *int64, *int64, *string, *string) (*models_pkg.ListCustomersResponse, error)

    DeleteAccessToken (string, string, *string) (*models_pkg.GetAccessTokenResponse, error)

    CreateAddress (string, *models_pkg.CreateAddressRequest, *string) (*models_pkg.GetAddressResponse, error)

    GetCard (string, string) (*models_pkg.GetCardResponse, error)
}

/*
 * Factory for the CUSTOMERS interaface returning CUSTOMERS_IMPL
 */
func NewCUSTOMERS(config configuration_pkg.CONFIGURATION) *CUSTOMERS_IMPL {
    client := new(CUSTOMERS_IMPL)
    client.config = config
    return client
}