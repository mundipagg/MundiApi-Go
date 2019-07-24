/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package subscriptions_pkg

import "time"
import "mundiapi_lib/configuration_pkg"
import "mundiapi_lib/models_pkg"

/*
 * Interface for the SUBSCRIPTIONS_IMPL
 */
type SUBSCRIPTIONS interface {
    GetSubscriptionCycleById (string, string) (*models_pkg.GetPeriodResponse, error)

    RenewSubscription (string, *string) (*models_pkg.GetPeriodResponse, error)

    UpdateCurrentCycleStatus (string, *models_pkg.UpdateCurrentCycleStatusRequest, *string) (error)

    GetSubscriptionCycles (string, string, string) (*models_pkg.ListCyclesResponse, error)

    UpdateSubscriptionBillingDate (string, *models_pkg.UpdateSubscriptionBillingDateRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    UpdateLatestPeriodEndAt (string, *models_pkg.UpdateCurrentCycleEndDateRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionDueDays (string, *models_pkg.UpdateSubscriptionDueDaysRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionMiniumPrice (string, *models_pkg.UpdateSubscriptionMinimumPriceRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    GetUsages (string, string, *int64, *int64, *string, *string) (*models_pkg.ListUsagesResponse, error)

    GetUsagesDetails (string, *string, *int64, *int64, *string, *string) (*models_pkg.GetUsagesDetailsResponse, error)

    GetIncrements (string, *int64, *int64) (*models_pkg.ListIncrementsResponse, error)

    CreateIncrement (string, *models_pkg.CreateIncrementRequest, *string) (*models_pkg.GetIncrementResponse, error)

    DeleteIncrement (string, string, *string) (*models_pkg.GetIncrementResponse, error)

    GetDiscounts (string, int64, int64) (*models_pkg.ListDiscountsResponse, error)

    CancelSubscription (string, *models_pkg.CreateCancelSubscriptionRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    GetDiscountById (string, string) (*models_pkg.GetDiscountResponse, error)

    DeleteUsage (string, string, string, *string) (*models_pkg.GetUsageResponse, error)

    DeleteSubscriptionItem (string, string, *string) (*models_pkg.GetSubscriptionItemResponse, error)

    DeleteDiscount (string, string, *string) (*models_pkg.GetDiscountResponse, error)

    CreateAnUsage (string, string, *string) (*models_pkg.GetUsageResponse, error)

    UpdateSubscriptionMetadata (string, *models_pkg.UpdateMetadataRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    GetSubscriptionItem (string, string) (*models_pkg.GetSubscriptionItemResponse, error)

    UpdateSubscriptionAffiliationId (string, *models_pkg.UpdateSubscriptionAffiliationIdRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    CreateSubscriptionItem (string, *models_pkg.CreateSubscriptionItemRequest, *string) (*models_pkg.GetSubscriptionItemResponse, error)

    CreateDiscount (string, *models_pkg.CreateDiscountRequest, *string) (*models_pkg.GetDiscountResponse, error)

    UpdateSubscriptionPaymentMethod (string, *models_pkg.UpdateSubscriptionPaymentMethodRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    CreateSubscription (*models_pkg.CreateSubscriptionRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    CreateUsage (string, string, *models_pkg.CreateUsageRequest, *string) (*models_pkg.GetUsageResponse, error)

    UpdateSubscriptionItem (string, string, *models_pkg.UpdateSubscriptionItemRequest, *string) (*models_pkg.GetSubscriptionItemResponse, error)

    UpdateSubscriptionCard (string, *models_pkg.UpdateSubscriptionCardRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    GetSubscription (string) (*models_pkg.GetSubscriptionResponse, error)

    GetIncrementById (string, string) (*models_pkg.GetIncrementResponse, error)

    UpdateSubscriptionStartAt (string, *models_pkg.UpdateSubscriptionStartAtRequest, *string) (*models_pkg.GetSubscriptionResponse, error)

    GetSubscriptions (*int64, *int64, *string, *string, *string, *string, *string, *string, *time.Time, *time.Time, *time.Time, *time.Time) (*models_pkg.ListSubscriptionsResponse, error)

    GetSubscriptionItems (string, *int64, *int64, *string, *string, *string, *string, *string, *string) (*models_pkg.ListSubscriptionItemsResponse, error)
}

/*
 * Factory for the SUBSCRIPTIONS interaface returning SUBSCRIPTIONS_IMPL
 */
func NewSUBSCRIPTIONS(config configuration_pkg.CONFIGURATION) *SUBSCRIPTIONS_IMPL {
    client := new(SUBSCRIPTIONS_IMPL)
    client.config = config
    return client
}