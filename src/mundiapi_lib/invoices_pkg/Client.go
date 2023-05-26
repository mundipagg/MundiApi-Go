/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package invoices_pkg


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
type INVOICES_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Create an Invoice
 * @param    string                                           subscriptionId      parameter: Required
 * @param    string                                           cycleId             parameter: Required
 * @param    *string                                          idempotencyKey      parameter: Optional
 * @param    *models_pkg.SubscriptionsCyclesPayRequest        body                parameter: Optional
 * @return	Returns the *models_pkg.SubscriptionsCyclesPayResponse response from the API call
 */
func (me *INVOICES_IMPL) CreateInvoice (
            subscriptionId string,
            cycleId string,
            idempotencyKey *string,
            body *models_pkg.SubscriptionsCyclesPayRequest) (*models_pkg.SubscriptionsCyclesPayResponse, error) {
    //the endpoint path uri
    _pathUrl := "/subscriptions/{subscription_id}/cycles/{cycle_id}/pay"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "subscription_id" : subscriptionId,
        "cycle_id" : cycleId,
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
        "user-agent" : "MundiSDK - Go 2.4.4",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Content-Type" : "application/json",
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
    var retVal *models_pkg.SubscriptionsCyclesPayResponse = &models_pkg.SubscriptionsCyclesPayResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * GetPartialInvoice
 * @param    string        subscriptionId      parameter: Required
 * @return	Returns the *models_pkg.SubscriptionsPartialInvoiceResponse response from the API call
 */
func (me *INVOICES_IMPL) GetPartialInvoice (
            subscriptionId string) (*models_pkg.SubscriptionsPartialInvoiceResponse, error) {
    //the endpoint path uri
    _pathUrl := "/subscriptions/{subscription_id}/partial-invoice"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "subscription_id" : subscriptionId,
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
        "user-agent" : "MundiSDK - Go 2.4.4",
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
    var retVal *models_pkg.SubscriptionsPartialInvoiceResponse = &models_pkg.SubscriptionsPartialInvoiceResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Updates the status from an invoice
 * @param    string                                             invoiceId           parameter: Required
 * @param    *models_pkg.UpdateCurrentCycleStatusRequest        body                parameter: Required
 * @param    *string                                            idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.InvoicesStatusResponse response from the API call
 */
func (me *INVOICES_IMPL) UpdateInvoiceStatus (
            invoiceId string,
            body *models_pkg.UpdateCurrentCycleStatusRequest,
            idempotencyKey *string) (*models_pkg.InvoicesStatusResponse, error) {
    //the endpoint path uri
    _pathUrl := "/invoices/{invoice_id}/status"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "invoice_id" : invoiceId,
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
        "user-agent" : "MundiSDK - Go 2.4.4",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Content-Type" : "application/json",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PatchWithAuth(_queryBuilder, headers, body, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
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
    var retVal *models_pkg.InvoicesStatusResponse = &models_pkg.InvoicesStatusResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets an invoice
 * @param    string        invoiceId      parameter: Required
 * @return	Returns the *models_pkg.InvoicesResponse response from the API call
 */
func (me *INVOICES_IMPL) GetInvoice (
            invoiceId string) (*models_pkg.InvoicesResponse, error) {
    //the endpoint path uri
    _pathUrl := "/invoices/{invoice_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "invoice_id" : invoiceId,
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
        "user-agent" : "MundiSDK - Go 2.4.4",
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
    var retVal *models_pkg.InvoicesResponse = &models_pkg.InvoicesResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Cancels an invoice
 * @param    string         invoiceId           parameter: Required
 * @param    *string        idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.InvoicesResponse response from the API call
 */
func (me *INVOICES_IMPL) CancelInvoice (
            invoiceId string,
            idempotencyKey *string) (*models_pkg.InvoicesResponse, error) {
    //the endpoint path uri
    _pathUrl := "/invoices/{invoice_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "invoice_id" : invoiceId,
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
        "user-agent" : "MundiSDK - Go 2.4.4",
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
    var retVal *models_pkg.InvoicesResponse = &models_pkg.InvoicesResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Updates the metadata from an invoice
 * @param    string                                     invoiceId           parameter: Required
 * @param    *models_pkg.InvoicesMetadataRequest        body                parameter: Required
 * @param    *string                                    idempotencyKey      parameter: Optional
 * @return	Returns the *models_pkg.InvoicesMetadataResponse response from the API call
 */
func (me *INVOICES_IMPL) UpdateInvoiceMetadata (
            invoiceId string,
            body *models_pkg.InvoicesMetadataRequest,
            idempotencyKey *string) (*models_pkg.InvoicesMetadataResponse, error) {
    //the endpoint path uri
    _pathUrl := "/invoices/{invoice_id}/metadata"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "invoice_id" : invoiceId,
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
        "user-agent" : "MundiSDK - Go 2.4.4",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Content-Type" : "application/json",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.PatchWithAuth(_queryBuilder, headers, body, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
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
    var retVal *models_pkg.InvoicesMetadataResponse = &models_pkg.InvoicesMetadataResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets all invoices
 * @param    *int64            page                  parameter: Optional
 * @param    *int64            size                  parameter: Optional
 * @param    *string           code                  parameter: Optional
 * @param    *string           customerId            parameter: Optional
 * @param    *string           subscriptionId        parameter: Optional
 * @param    *time.Time        createdSince          parameter: Optional
 * @param    *time.Time        createdUntil          parameter: Optional
 * @param    *string           status                parameter: Optional
 * @param    *time.Time        dueSince              parameter: Optional
 * @param    *time.Time        dueUntil              parameter: Optional
 * @param    *string           customerDocument      parameter: Optional
 * @return	Returns the *models_pkg.InvoicesResponse2 response from the API call
 */
func (me *INVOICES_IMPL) GetInvoices (
            page *int64,
            size *int64,
            code *string,
            customerId *string,
            subscriptionId *string,
            createdSince *time.Time,
            createdUntil *time.Time,
            status *string,
            dueSince *time.Time,
            dueUntil *time.Time,
            customerDocument *string) (*models_pkg.InvoicesResponse2, error) {
    //the endpoint path uri
    _pathUrl := "/invoices"

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
        "code" : code,
        "customer_id" : customerId,
        "subscription_id" : subscriptionId,
        "created_since" : createdSince,
        "created_until" : createdUntil,
        "status" : status,
        "due_since" : dueSince,
        "due_until" : dueUntil,
        "customer_document" : customerDocument,
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
        "user-agent" : "MundiSDK - Go 2.4.4",
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
    var retVal *models_pkg.InvoicesResponse2 = &models_pkg.InvoicesResponse2{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

