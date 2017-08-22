/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package plans_pkg

import "mundiapi_lib/models_pkg"

/*
 * Interface for the PLANS_IMPL
 */
type PLANS interface {
    GetPlanItems (string) (*models_pkg.ListPlanItemsResponse, error)

    UpdatePlanItem (string, string, *models_pkg.UpdatePlanItemRequest) (*models_pkg.GetPlanItemResponse, error)

    GetPlan (string) (*models_pkg.GetPlanResponse, error)

    CreatePlanItem (string, *models_pkg.CreatePlanItemRequest) (*models_pkg.GetPlanItemResponse, error)

    UpdatePlan (string, *models_pkg.UpdatePlanRequest) (*models_pkg.GetPlanResponse, error)

    CreatePlan (*models_pkg.CreatePlanRequest) (*models_pkg.GetPlanResponse, error)

    GetPlans () (*models_pkg.ListPlansResponse, error)

    DeletePlan (string) (*models_pkg.GetPlanResponse, error)

    GetPlanItem (string, string) (*models_pkg.GetPlanItemResponse, error)

    DeletePlanItem (string, string) (*models_pkg.GetPlanItemResponse, error)
}

/*
 * Factory for the PLANS interaface returning PLANS_IMPL
 */
func NewPLANS() PLANS {
    return &PLANS_IMPL{}
}