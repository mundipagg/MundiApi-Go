/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package plans_pkg


import(
	"time"
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"mundiapi_lib/apihelper_pkg"
	"mundiapi_lib/configuration_pkg"
	"mundiapi_lib/models_pkg"
)
/*
 * Client structure as interface implementation
 */
type PLANS_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Adds a new item to a plan
 * @param    string                                   planId              parameter: Required
 * @param    *models_pkg.CreatePlanItemRequest        request             parameter: Required
 * @param    *string                                  idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetPlanItemResponse response from the API call
 */
func (me *PLANS_IMPL) CreatePlanItem (
            planId string,
            request *models_pkg.CreatePlanItemRequest,
            idempotencyKey *string) (*models_pkg.GetPlanItemResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans/{plan_id}/items"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "plan_id" : planId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, request, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanItemResponse = &models_pkg.GetPlanItemResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Updates a plan item
 * @param    string                                   planId              parameter: Required
 * @param    string                                   planItemId          parameter: Required
 * @param    *models_pkg.UpdatePlanItemRequest        body                parameter: Required
 * @param    *string                                  idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetPlanItemResponse response from the API call
 */
func (me *PLANS_IMPL) UpdatePlanItem (
            planId string,
            planItemId string,
            body *models_pkg.UpdatePlanItemRequest,
            idempotencyKey *string) (*models_pkg.GetPlanItemResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans/{plan_id}/items/{plan_item_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "plan_id" : planId,
        "plan_item_id" : planItemId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PutWithAuth(_queryBuilder, headers, body, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanItemResponse = &models_pkg.GetPlanItemResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets a plan
 * @param    string        planId      parameter: Required
 * @return	Returns the *models_pkg.GetPlanResponse response from the API call
 */
func (me *PLANS_IMPL) GetPlan (
            planId string) (*models_pkg.GetPlanResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans/{plan_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "plan_id" : planId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanResponse = &models_pkg.GetPlanResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Deletes a plan
 * @param    string         planId              parameter: Required
 * @param    *string        idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetPlanResponse response from the API call
 */
func (me *PLANS_IMPL) DeletePlan (
            planId string,
            idempotencyKey *string) (*models_pkg.GetPlanResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans/{plan_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "plan_id" : planId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.DeleteWithAuth(_queryBuilder, headers, nil, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanResponse = &models_pkg.GetPlanResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Updates a plan
 * @param    string                               planId              parameter: Required
 * @param    *models_pkg.UpdatePlanRequest        request             parameter: Required
 * @param    *string                              idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetPlanResponse response from the API call
 */
func (me *PLANS_IMPL) UpdatePlan (
            planId string,
            request *models_pkg.UpdatePlanRequest,
            idempotencyKey *string) (*models_pkg.GetPlanResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans/{plan_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "plan_id" : planId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PutWithAuth(_queryBuilder, headers, request, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanResponse = &models_pkg.GetPlanResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Creates a new plan
 * @param    *models_pkg.CreatePlanRequest        body                parameter: Required
 * @param    *string                              idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetPlanResponse response from the API call
 */
func (me *PLANS_IMPL) CreatePlan (
            body *models_pkg.CreatePlanRequest,
            idempotencyKey *string) (*models_pkg.GetPlanResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, body, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanResponse = &models_pkg.GetPlanResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets all plans
 * @param    *int64            page              parameter: Optional
 * @param    *int64            size              parameter: Optional
 * @param    *string           name              parameter: Optional
 * @param    *string           status            parameter: Optional
 * @param    *string           billingType       parameter: Optional
 * @param    *time.Time        createdSince      parameter: Optional
 * @param    *time.Time        createdUntil      parameter: Optional
 * @return	Returns the *models_pkg.ListPlansResponse response from the API call
 */
func (me *PLANS_IMPL) GetPlans (
            page *int64,
            size *int64,
            name *string,
            status *string,
            billingType *string,
            createdSince *time.Time,
            createdUntil *time.Time) (*models_pkg.ListPlansResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "page" : page,
        "size" : size,
        "name" : name,
        "status" : status,
        "billing_type" : billingType,
        "created_since" : createdSince,
        "created_until" : createdUntil,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.ListPlansResponse = &models_pkg.ListPlansResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Updates the metadata from a plan
 * @param    string                                   planId              parameter: Required
 * @param    *models_pkg.UpdateMetadataRequest        request             parameter: Required
 * @param    *string                                  idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetPlanResponse response from the API call
 */
func (me *PLANS_IMPL) UpdatePlanMetadata (
            planId string,
            request *models_pkg.UpdateMetadataRequest,
            idempotencyKey *string) (*models_pkg.GetPlanResponse, error) {
    //the endpoint path uri
    _pathUrl := "/Plans/{plan_id}/metadata"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "plan_id" : planId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PatchWithAuth(_queryBuilder, headers, request, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanResponse = &models_pkg.GetPlanResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets a plan item
 * @param    string        planId           parameter: Required
 * @param    string        planItemId       parameter: Required
 * @return	Returns the *models_pkg.GetPlanItemResponse response from the API call
 */
func (me *PLANS_IMPL) GetPlanItem (
            planId string,
            planItemId string) (*models_pkg.GetPlanItemResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans/{plan_id}/items/{plan_item_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "plan_id" : planId,
        "plan_item_id" : planItemId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanItemResponse = &models_pkg.GetPlanItemResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Removes an item from a plan
 * @param    string         planId              parameter: Required
 * @param    string         planItemId          parameter: Required
 * @param    *string        idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.GetPlanItemResponse response from the API call
 */
func (me *PLANS_IMPL) DeletePlanItem (
            planId string,
            planItemId string,
            idempotencyKey *string) (*models_pkg.GetPlanItemResponse, error) {
    //the endpoint path uri
    _pathUrl := "/plans/{plan_id}/items/{plan_item_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "plan_id" : planId,
        "plan_item_id" : planItemId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "MundiSDK - Go 0.16.5",
        "accept" : "application/json",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.DeleteWithAuth(_queryBuilder, headers, nil, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Invalid request", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Invalid API key", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("An informed resource was not found", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Business validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Contract validation error", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Internal server error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetPlanItemResponse = &models_pkg.GetPlanItemResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

