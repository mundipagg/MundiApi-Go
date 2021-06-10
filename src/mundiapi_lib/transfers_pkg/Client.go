/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package transfers_pkg


import(
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"mundiapi_lib/apihelper_pkg"
	"mundiapi_lib/configuration_pkg"
	"mundiapi_lib/models_pkg"
)
/*
 * Client structure as interface implementation
 */
type TRANSFERS_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * TODO: type endpoint description here
 * @param    *models_pkg.CreateTransfer        request     parameter: Required
 * @return	Returns the *models_pkg.GetTransfer response from the API call
 */
func (me *TRANSFERS_IMPL) CreateTransfer (
            request *models_pkg.CreateTransfer) (*models_pkg.GetTransfer, error) {
    //the endpoint path uri
    _pathUrl := "/transfers/recipients"

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
        "user-agent" : "MundiSDK - Go 2.0.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
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
    var retVal *models_pkg.GetTransfer = &models_pkg.GetTransfer{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * TODO: type endpoint description here
 * @param    string        transferId      parameter: Required
 * @return	Returns the *models_pkg.GetTransfer response from the API call
 */
func (me *TRANSFERS_IMPL) GetTransferById (
            transferId string) (*models_pkg.GetTransfer, error) {
    //the endpoint path uri
    _pathUrl := "/transfers/{transfer_id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "transfer_id" : transferId,
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
        "user-agent" : "MundiSDK - Go 2.0.0",
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
    var retVal *models_pkg.GetTransfer = &models_pkg.GetTransfer{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets all transfers
 * @return	Returns the *models_pkg.ListTransfers response from the API call
 */
func (me *TRANSFERS_IMPL) GetTransfers () (*models_pkg.ListTransfers, error) {
    //the endpoint path uri
    _pathUrl := "/transfers"

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
        "user-agent" : "MundiSDK - Go 2.0.0",
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
    var retVal *models_pkg.ListTransfers = &models_pkg.ListTransfers{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

