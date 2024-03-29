/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package invoices_pkg

import "time"
import "mundiapi_lib/configuration_pkg"
import "mundiapi_lib/models_pkg"

/*
 * Interface for the INVOICES_IMPL
 */
type INVOICES interface {
    CreateInvoice (string, string, *string, *models_pkg.SubscriptionsCyclesPayRequest) (*models_pkg.SubscriptionsCyclesPayResponse, error)

    GetPartialInvoice (string) (*models_pkg.SubscriptionsPartialInvoiceResponse, error)

    UpdateInvoiceStatus (string, *models_pkg.UpdateCurrentCycleStatusRequest, *string) (*models_pkg.InvoicesStatusResponse, error)

    GetInvoice (string) (*models_pkg.InvoicesResponse, error)

    CancelInvoice (string, *string) (*models_pkg.InvoicesResponse, error)

    UpdateInvoiceMetadata (string, *models_pkg.InvoicesMetadataRequest, *string) (*models_pkg.InvoicesMetadataResponse, error)

    GetInvoices (*int64, *int64, *string, *string, *string, *time.Time, *time.Time, *string, *time.Time, *time.Time, *string) (*models_pkg.InvoicesResponse2, error)
}

/*
 * Factory for the INVOICES interaface returning INVOICES_IMPL
 */
func NewINVOICES(config configuration_pkg.CONFIGURATION) *INVOICES_IMPL {
    client := new(INVOICES_IMPL)
    client.config = config
    return client
}