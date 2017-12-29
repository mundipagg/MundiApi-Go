/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package recipients_pkg

import "time"
import "mundiapi_lib/models_pkg"

/*
 * Interface for the RECIPIENTS_IMPL
 */
type RECIPIENTS interface {
    UpdateRecipientMetadata (string, *models_pkg.UpdateMetadataRequest) (*models_pkg.GetRecipientResponse, error)

    GetTransfer (string, string) (*models_pkg.GetTransferResponse, error)

    GetTransfers (string, *int64, *int64, *string, *time.Time, *time.Time) (*models_pkg.ListTransferResponse, error)

    CreateAnticipation (string, *models_pkg.CreateAnticipationRequest) (*models_pkg.GetAnticipationResponse, error)

    GetAnticipation (string, string) (*models_pkg.GetAnticipationResponse, error)

    GetAnticipationLimits (string, string, *time.Time) (*models_pkg.GetAnticipationLimitResponse, error)

    GetAnticipations (string, *int64, *int64, *string, *string, *time.Time, *time.Time, *time.Time, *time.Time) (*models_pkg.ListAnticipationResponse, error)

    UpdateRecipient (string, *models_pkg.UpdateRecipientRequest) (*models_pkg.GetRecipientResponse, error)

    UpdateRecipientDefaultBankAccount (string, *models_pkg.UpdateRecipientBankAccountRequest) (*models_pkg.GetRecipientResponse, error)

    GetRecipient (string) (*models_pkg.GetRecipientResponse, error)

    GetRecipients (*int64, *int64) (*models_pkg.ListRecipientResponse, error)

    GetBalance (string) (*models_pkg.GetBalanceResponse, error)

    CreateTransfer (string, *models_pkg.CreateTransferRequest) (*models_pkg.GetTransferResponse, error)

    CreateRecipient (*models_pkg.CreateRecipientRequest) (*models_pkg.GetRecipientResponse, error)
}

/*
 * Factory for the RECIPIENTS interaface returning RECIPIENTS_IMPL
 */
func NewRECIPIENTS() RECIPIENTS {
    return &RECIPIENTS_IMPL{}
}