/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package charges_pkg

import "time"
import "mundiapi_lib/models_pkg"

/*
 * Interface for the CHARGES_IMPL
 */
type CHARGES interface {
    UpdateChargeCard (string, *models_pkg.UpdateChargeCardRequest) (*models_pkg.GetChargeResponse, error)

    UpdateChargePaymentMethod (string, *models_pkg.UpdateChargePaymentMethodRequest) (*models_pkg.GetChargeResponse, error)

    CreateCharge (*models_pkg.CreateChargeRequest) (*models_pkg.GetChargeResponse, error)

    GetCharge (string) (*models_pkg.GetChargeResponse, error)

    RetryCharge (string) (*models_pkg.GetChargeResponse, error)

    GetCharges (*int64, *int64, *string, *string, *string, *string, *string, *time.Time, *time.Time) (*models_pkg.ListChargesResponse, error)

    UpdateChargeMetadata (string, *models_pkg.UpdateMetadataRequest) (*models_pkg.GetChargeResponse, error)

    CancelCharge (string, *models_pkg.CreateCancelChargeRequest) (*models_pkg.GetChargeResponse, error)

    CaptureCharge (string, *models_pkg.CreateCaptureChargeRequest) (*models_pkg.GetChargeResponse, error)

    UpdateChargeDueDate (string, *models_pkg.UpdateChargeDueDateRequest) (*models_pkg.GetChargeResponse, error)
}

/*
 * Factory for the CHARGES interaface returning CHARGES_IMPL
 */
func NewCHARGES() CHARGES {
    return &CHARGES_IMPL{}
}