/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package tokens_pkg


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
type TOKENS_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * CreateToken
 * @param    string                           publicKey           parameter: Required
 * @param    *models_pkg.TokensRequest        body                parameter: Required
 * @param    *string                          idempotencyKey      parameter: Optional
 * @param    *string                          appId               parameter: Optional
 * @return	Returns the *models_pkg.TokensResponse response from the API call
 */
func (me *TOKENS_IMPL) CreateToken (
            publicKey string,
            body *models_pkg.TokensRequest,
            idempotencyKey *string,
            appId *string) (*models_pkg.TokensResponse, error) {
    //the endpoint path uri
    _pathUrl := "/tokens"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "public_key" : publicKey,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "appId" : appId,
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
        "user-agent" : "MundiSDK - Go 2.4.2",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Content-Type" : "application/json",
        "idempotency-key" : apihelper_pkg.ToString(idempotencyKey, ""),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
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
    var retVal *models_pkg.TokensResponse = &models_pkg.TokensResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets a token from its id
 * @param    string         id             parameter: Required
 * @param    string         publicKey      parameter: Required
 * @param    *string        appId          parameter: Optional
 * @return	Returns the *models_pkg.TokensResponse response from the API call
 */
func (me *TOKENS_IMPL) GetToken (
            id string,
            publicKey string,
            appId *string) (*models_pkg.TokensResponse, error) {
    //the endpoint path uri
    _pathUrl := "/tokens/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
        "public_key" : publicKey,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "appId" : appId,
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
        "user-agent" : "MundiSDK - Go 2.4.2",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
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
    var retVal *models_pkg.TokensResponse = &models_pkg.TokensResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

