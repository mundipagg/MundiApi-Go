/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package recipients_pkg

import "time"
import "mundiapi_lib/configuration_pkg"
import "mundiapi_lib/models_pkg"

/*
 * Interface for the RECIPIENTS_IMPL
 */
type RECIPIENTS interface {
    UpdateRecipientMetadata (string, *models_pkg.UpdateMetadataRequest, *string) (*models_pkg.GetRecipientResponse, error)

    UpdateRecipientTransferSettings (string, *models_pkg.UpdateTransferSettingsRequest, *string) (*models_pkg.GetRecipientResponse, error)

    GetAnticipation (string, string) (*models_pkg.GetAnticipationResponse, error)

    GetRecipients (*int64, *int64) (*models_pkg.ListRecipientResponse, error)

    GetBalance (string) (*models_pkg.GetBalanceResponse, error)

    GetAnticipations (string, *int64, *int64, *string, *string, *time.Time, *time.Time, *time.Time, *time.Time) (*models_pkg.ListAnticipationResponse, error)

    CreateAnticipation (string, *models_pkg.CreateAnticipationRequest, *string) (*models_pkg.GetAnticipationResponse, error)

    UpdateRecipientDefaultBankAccount (string, *models_pkg.UpdateRecipientBankAccountRequest, *string) (*models_pkg.GetRecipientResponse, error)

    GetRecipient (string) (*models_pkg.GetRecipientResponse, error)

    GetAnticipationLimits (string, string, *time.Time) (*models_pkg.GetAnticipationLimitResponse, error)

    GetTransfer (string, string) (*models_pkg.GetTransferResponse, error)

    GetTransfers (string, *int64, *int64, *string, *time.Time, *time.Time) (*models_pkg.ListTransferResponse, error)

    UpdateRecipient (string, *models_pkg.UpdateRecipientRequest, *string) (*models_pkg.GetRecipientResponse, error)

    CreateRecipient (*models_pkg.CreateRecipientRequest, *string) (*models_pkg.GetRecipientResponse, error)

    CreateTransfer (string, *models_pkg.CreateTransferRequest, *string) (*models_pkg.GetTransferResponse, error)

    CreateWithdraw (string, *models_pkg.CreateWithdrawRequest) (*models_pkg.GetWithdrawResponse, error)

    GetWithdrawById (string, string) (*models_pkg.GetWithdrawResponse, error)

    GetWithdrawals (string, *int64, *int64, *string, *time.Time, *time.Time) (*models_pkg.ListWithdrawals, error)

    UpdateAutomaticAnticipationSettings (string, *models_pkg.UpdateAutomaticAnticipationSettingsRequest, *string) (*models_pkg.GetRecipientResponse, error)

    GetRecipientByCode (string) (*models_pkg.GetRecipientResponse, error)
}

/*
 * Factory for the RECIPIENTS interaface returning RECIPIENTS_IMPL
 */
func NewRECIPIENTS(config configuration_pkg.CONFIGURATION) *RECIPIENTS_IMPL {
    client := new(RECIPIENTS_IMPL)
    client.config = config
    return client
}