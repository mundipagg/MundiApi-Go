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
    CancelInvoice (string) (*models_pkg.GetInvoiceResponse, error)

    GetInvoice (string) (*models_pkg.GetInvoiceResponse, error)

    CreateInvoice (string, string, *models_pkg.CreateInvoiceRequest) (*models_pkg.GetInvoiceResponse, error)

    UpdateInvoiceStatus (string, *models_pkg.UpdateInvoiceStatusRequest) (*models_pkg.GetInvoiceResponse, error)

    GetInvoices (*int64, *int64, *string, *string, *string, *time.Time, *time.Time, *string, *time.Time, *time.Time) (*models_pkg.ListInvoicesResponse, error)

    UpdateInvoiceMetadata (string, *models_pkg.UpdateMetadataRequest) (*models_pkg.GetInvoiceResponse, error)

    GetPartialInvoice (string) (*models_pkg.GetInvoiceResponse, error)
}

/*
 * Factory for the INVOICES interaface returning INVOICES_IMPL
 */
func NewINVOICES(config configuration_pkg.CONFIGURATION) *INVOICES_IMPL {
    client := new(INVOICES_IMPL)
    client.config = config
    return client
}