/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package charges_pkg

import "mundiapi_lib/models_pkg"

/*
 * Interface for the CHARGES_IMPL
 */
type CHARGES interface {
    GetCharge (string) (*models_pkg.GetChargeResponse, error)

    RetryCharge (string) (*models_pkg.GetChargeResponse, error)

    GetCharges () (*models_pkg.ListChargesResponse, error)

    CreateCharge (*models_pkg.CreateChargeRequest) (*models_pkg.GetChargeResponse, error)

    UpdateChargeCard (string, *models_pkg.UpdateChargeCardRequest) (*models_pkg.GetChargeResponse, error)

    UpdateChargePaymentMethod (string, *models_pkg.UpdateChargePaymentMethodRequest) (*models_pkg.GetChargeResponse, error)

    CancelCharge (string, *models_pkg.CreateCancelChargeRequest) (*models_pkg.GetChargeResponse, error)

    CaptureCharge (string, *models_pkg.CreateCaptureChargeRequest) (*models_pkg.GetChargeResponse, error)

    UpdateChargeMetadata (string, *models_pkg.UpdateMetadataRequest) (*models_pkg.GetChargeResponse, error)
}

/*
 * Factory for the CHARGES interaface returning CHARGES_IMPL
 */
func NewCHARGES() CHARGES {
    return &CHARGES_IMPL{}
}