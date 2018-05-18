/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package MundiAPIClient

import(
	"mundiapi_lib/configuration_pkg"
	"mundiapi_lib/subscriptions_pkg"
	"mundiapi_lib/charges_pkg"
	"mundiapi_lib/customers_pkg"
	"mundiapi_lib/invoices_pkg"
	"mundiapi_lib/plans_pkg"
	"mundiapi_lib/orders_pkg"
	"mundiapi_lib/tokens_pkg"
	"mundiapi_lib/recipients_pkg"
	"mundiapi_lib/sellers_pkg"
)


/*
 * Interface for the MUNDIAPI_IMPL
 */
type MUNDIAPI interface {
     Subscriptions()         subscriptions_pkg.SUBSCRIPTIONS
     Charges()               charges_pkg.CHARGES
     Customers()             customers_pkg.CUSTOMERS
     Invoices()              invoices_pkg.INVOICES
     Plans()                 plans_pkg.PLANS
     Orders()                orders_pkg.ORDERS
     Tokens()                tokens_pkg.TOKENS
     Recipients()            recipients_pkg.RECIPIENTS
     Sellers()               sellers_pkg.SELLERS
     Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the MUNDIAPI interaface returning MUNDIAPI_IMPL
 */
func NewMUNDIAPI(basicAuthUserName string, basicAuthPassword string) MUNDIAPI {
    mundiAPIClient := new(MUNDIAPI_IMPL)
    mundiAPIClient.config = configuration_pkg.NewCONFIGURATION()
    mundiAPIClient.config.SetBasicAuthUserName(basicAuthUserName)
    mundiAPIClient.config.SetBasicAuthPassword(basicAuthPassword)
    return mundiAPIClient
}
