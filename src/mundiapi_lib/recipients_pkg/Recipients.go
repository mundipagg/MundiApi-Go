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
    UpdateRecipientMetadata (string, *models_pkg.RecipientsMetadataRequest, *string) (*models_pkg.RecipientsMetadataResponse, error)

    UpdateRecipientTransferSettings (string, *models_pkg.UpdateTransferSettingsRequest, *string) (*models_pkg.RecipientsTransferSettingsResponse, error)

    GetAnticipation (string, string) (*models_pkg.RecipientsAnticipationsResponse, error)

    GetRecipients (*int64, *int64) (*models_pkg.RecipientsResponse, error)

    CreateRecipient (*models_pkg.RecipientsRequest, *string) (*models_pkg.RecipientsResponse1, error)

    GetBalance (string) (*models_pkg.RecipientsBalanceResponse, error)

    GetAnticipations (string, *int64, *int64, *string, *string, *time.Time, *time.Time, *time.Time, *time.Time) (*models_pkg.RecipientsAnticipationsResponse1, error)

    CreateAnticipation (string, *models_pkg.RecipientsAnticipationsRequest, *string) (*models_pkg.RecipientsAnticipationsResponse, error)

    UpdateRecipientDefaultBankAccount (string, *models_pkg.RecipientsDefaultBankAccountRequest, *string) (*models_pkg.RecipientsDefaultBankAccountResponse, error)

    GetRecipient (string) (*models_pkg.RecipientsResponse1, error)

    UpdateRecipient (string, *models_pkg.RecipientsRequest1, *string) (*models_pkg.RecipientsResponse1, error)

    GetTransfer (string, string) (*models_pkg.RecipientsTransfersResponse, error)

    GetTransfers (string, *int64, *int64, *string, *time.Time, *time.Time) (*models_pkg.RecipientsTransfersResponse1, error)

    CreateTransfer (string, *models_pkg.RecipientsTransfersRequest, *string) (*models_pkg.RecipientsTransfersResponse, error)

    GetAnticipationLimits (string, string, *time.Time) (*models_pkg.RecipientsAnticipationLimitsResponse, error)

    CreateWithdraw (string, *models_pkg.CreateWithdrawRequest) (*models_pkg.GetWithdrawResponse, error)

    GetWithdrawals (string, *int64, *int64, *string, *time.Time, *time.Time) (*models_pkg.ListWithdrawals, error)

    GetWithdrawById (string, string) (*models_pkg.GetWithdrawResponse, error)

    UpdateAutomaticAnticipationSettings (string, *models_pkg.UpdateAutomaticAnticipationSettingsRequest, *string) (*models_pkg.RecipientsAutomaticAnticipationSettingsResponse, error)

    GetRecipientByCode (string) (*models_pkg.RecipientsCodeResponse, error)
}

/*
 * Factory for the RECIPIENTS interaface returning RECIPIENTS_IMPL
 */
func NewRECIPIENTS(config configuration_pkg.CONFIGURATION) *RECIPIENTS_IMPL {
    client := new(RECIPIENTS_IMPL)
    client.config = config
    return client
}