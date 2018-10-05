/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package tokens_pkg

import "mundiapi_lib/models_pkg"
import "mundiapi_lib/configuration_pkg"

/*
 * Interface for the TOKENS_IMPL
 */
type TOKENS interface {
    GetToken (string, string) (*models_pkg.GetTokenResponse, error)

    CreateToken (string, *models_pkg.CreateTokenRequest) (*models_pkg.GetTokenResponse, error)
}

/*
 * Factory for the TOKENS interaface returning TOKENS_IMPL
 */
func NewTOKENS(config configuration_pkg.CONFIGURATION) *TOKENS_IMPL {
    client := new(TOKENS_IMPL)
    client.config = config
    return client
}