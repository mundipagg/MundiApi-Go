/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package subscriptions_pkg

import "time"
import "mundiapi_lib/models_pkg"
import "mundiapi_lib/configuration_pkg"

/*
 * Interface for the SUBSCRIPTIONS_IMPL
 */
type SUBSCRIPTIONS interface {
    GetIncrementById (string, string) (*models_pkg.GetIncrementResponse, error)

    UpdateSubscriptionStartAt (string, *models_pkg.UpdateSubscriptionStartAtRequest) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionCard (string, *models_pkg.UpdateSubscriptionCardRequest) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionItem (string, string, *models_pkg.UpdateSubscriptionItemRequest) (*models_pkg.GetSubscriptionItemResponse, error)

    CreateUsage (string, string, *models_pkg.CreateUsageRequest) (*models_pkg.GetUsageResponse, error)

    GetSubscription (string) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionPaymentMethod (string, *models_pkg.UpdateSubscriptionPaymentMethodRequest) (*models_pkg.GetSubscriptionResponse, error)

    CreateSubscription (*models_pkg.CreateSubscriptionRequest) (*models_pkg.GetSubscriptionResponse, error)

    CreateSubscriptionItem (string, *models_pkg.CreateSubscriptionItemRequest) (*models_pkg.GetSubscriptionItemResponse, error)

    CreateDiscount (string, *models_pkg.CreateDiscountRequest) (*models_pkg.GetDiscountResponse, error)

    GetSubscriptionItem (string, string) (*models_pkg.GetSubscriptionItemResponse, error)

    UpdateSubscriptionAffiliationId (string, *models_pkg.UpdateSubscriptionAffiliationIdRequest) (*models_pkg.GetSubscriptionResponse, error)

    CreateAnUsage (string, string) (*models_pkg.GetUsageResponse, error)

    GetSubscriptions (*int64, *int64, *string, *string, *string, *string, *string, *string, *time.Time, *time.Time, *time.Time, *time.Time) (*models_pkg.ListSubscriptionsResponse, error)

    UpdateSubscriptionMetadata (string, *models_pkg.UpdateMetadataRequest) (*models_pkg.GetSubscriptionResponse, error)

    DeleteSubscriptionItem (string, string) (*models_pkg.GetSubscriptionItemResponse, error)

    DeleteUsage (string, string, string) (*models_pkg.GetUsageResponse, error)

    DeleteDiscount (string, string) (*models_pkg.GetDiscountResponse, error)

    CancelSubscription (string, *models_pkg.CreateCancelSubscriptionRequest) (*models_pkg.GetSubscriptionResponse, error)

    GetDiscountById (string, string) (*models_pkg.GetDiscountResponse, error)

    GetDiscounts (string, int64, int64) (*models_pkg.ListDiscountsResponse, error)

    CreateIncrement (string, *models_pkg.CreateIncrementRequest) (*models_pkg.GetIncrementResponse, error)

    GetIncrements (string, *int64, *int64) (*models_pkg.ListIncrementsResponse, error)

    DeleteIncrement (string, string) (*models_pkg.GetIncrementResponse, error)

    GetUsagesDetails (string, *string, *int64, *int64, *string, *string) (*models_pkg.GetUsagesDetailsResponse, error)

    GetUsages (string, string, *int64, *int64, *string, *string) (*models_pkg.ListUsagesResponse, error)

    GetSubscriptionItems (string, *int64, *int64, *string, *string, *string, *string, *string, *string) (*models_pkg.ListSubscriptionItemsResponse, error)

    UpdateSubscriptionDueDays (string, *models_pkg.UpdateSubscriptionDueDaysRequest) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionMiniumPrice (string, *models_pkg.UpdateSubscriptionMinimumPriceRequest) (*models_pkg.GetSubscriptionResponse, error)

    UpdateSubscriptionBillingDate (string, *models_pkg.UpdateSubscriptionBillingDateRequest) (*models_pkg.GetSubscriptionResponse, error)

    UpdateCurrentCycleEndDate (string, *models_pkg.UpdateCurrentCycleEndDateRequest) (*models_pkg.GetSubscriptionItemResponse, error)

    UpdateCurrentCycleStatus (string, *models_pkg.UpdateCurrentCycleStatusRequest) (error)

    GetSubscriptionCycles (string, string, string) (*models_pkg.ListCyclesResponse, error)

    GetSubscriptionCycleById (string, string) (*models_pkg.GetPeriodResponse, error)
}

/*
 * Factory for the SUBSCRIPTIONS interaface returning SUBSCRIPTIONS_IMPL
 */
func NewSUBSCRIPTIONS(config configuration_pkg.CONFIGURATION) *SUBSCRIPTIONS_IMPL {
    client := new(SUBSCRIPTIONS_IMPL)
    client.config = config
    return client
}