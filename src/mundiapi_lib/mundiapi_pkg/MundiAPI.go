/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package MundiAPIClient

import(
	"mundiapi_lib/configuration_pkg"
	"mundiapi_lib/subscriptions_pkg"
	"mundiapi_lib/orders_pkg"
	"mundiapi_lib/plans_pkg"
	"mundiapi_lib/invoices_pkg"
	"mundiapi_lib/customers_pkg"
	"mundiapi_lib/charges_pkg"
	"mundiapi_lib/recipients_pkg"
	"mundiapi_lib/tokens_pkg"
	"mundiapi_lib/transactions_pkg"
	"mundiapi_lib/transfers_pkg"
)


/*
 * Interface for the MUNDIAPI_IMPL
 */
type MUNDIAPI interface {
    Subscriptions()         subscriptions_pkg.SUBSCRIPTIONS
    Orders()                orders_pkg.ORDERS
    Plans()                 plans_pkg.PLANS
    Invoices()              invoices_pkg.INVOICES
    Customers()             customers_pkg.CUSTOMERS
    Charges()               charges_pkg.CHARGES
    Recipients()            recipients_pkg.RECIPIENTS
    Tokens()                tokens_pkg.TOKENS
    Transactions()          transactions_pkg.TRANSACTIONS
    Transfers()             transfers_pkg.TRANSFERS
    Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the MUNDIAPI interface returning MUNDIAPI_IMPL
 */
func NewMUNDIAPI(basicAuthUserName string, basicAuthPassword string) MUNDIAPI {
    mundiAPIClient := new(MUNDIAPI_IMPL)
    mundiAPIClient.config = configuration_pkg.NewCONFIGURATION()

    mundiAPIClient.config.SetBasicAuthUserName(basicAuthUserName)
    mundiAPIClient.config.SetBasicAuthPassword(basicAuthPassword)

    return mundiAPIClient
}
