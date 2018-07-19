/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package sellers_pkg

import "time"
import "mundiapi_lib/models_pkg"
import "mundiapi_lib/configuration_pkg"

/*
 * Interface for the SELLERS_IMPL
 */
type SELLERS interface {
    GetSellerById (string) (*models_pkg.GetSellerResponse, error)

    DeleteSeller (string) (*models_pkg.GetSellerResponse, error)

    CreateSeller (*models_pkg.CreateSellerRequest) (*models_pkg.GetSellerResponse, error)

    GetSellers (*int64, *int64, *string, *string, *string, *string, *string, *time.Time, *time.Time) (*models_pkg.ListSellerResponse, error)

    UpdateSeller (string, *models_pkg.UpdateSellerRequest) (*models_pkg.GetSellerResponse, error)

    UpdateSellerMetadata (string, *models_pkg.UpdateMetadataRequest) (*models_pkg.GetSellerResponse, error)
}

/*
 * Factory for the SELLERS interaface returning SELLERS_IMPL
 */
func NewSELLERS(config configuration_pkg.CONFIGURATION) *SELLERS_IMPL {
    client := new(SELLERS_IMPL)
    client.config = config
    return client
}