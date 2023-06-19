/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package configuration_pkg


/** The base Uri for API calls */
const BASEURI string = "https://api.mundipagg.com/core/v1"



type CONFIGURATION_IMPL struct {
    /** Replace the value of ServiceRefererName with an appropriate value */
    servicereferername string
    /** The username to use with basic authentication */
    /** Replace the value of basicauthusername with SetBasicAuthUserName function */
    basicauthusername string
    /** The password to use with basic authentication */
    /** Replace the value of basicauthpassword with SetBasicAuthPassword function */
    basicauthpassword string
}

/*
 * Getter function returning servicereferername
 */
func (me *CONFIGURATION_IMPL) ServiceRefererName() string{
    return me.servicereferername
}

/*
 * Setter function setting up servicereferername
 */
func (me *CONFIGURATION_IMPL) SetServiceRefererName(serviceRefererName string) {
    me.servicereferername = serviceRefererName
}

/*
 * Getter function returning basicauthusername
 */
func (me *CONFIGURATION_IMPL) BasicAuthUserName() string{
    return me.basicauthusername
}

/*
 * Setter function setting up basicauthusername
 */
func (me *CONFIGURATION_IMPL) SetBasicAuthUserName(basicAuthUserName string) {
    me.basicauthusername = basicAuthUserName
}

/*
 * Getter function returning basicauthpassword
 */
func (me *CONFIGURATION_IMPL) BasicAuthPassword() string{
    return me.basicauthpassword
}

/*
 * Setter function setting up basicauthpassword
 */
func (me *CONFIGURATION_IMPL) SetBasicAuthPassword(basicAuthPassword string) {
    me.basicauthpassword = basicAuthPassword
}

