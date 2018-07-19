/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
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
	"mundiapi_lib/sellers_pkg"
)
/*
 * Client structure as interface implementation
 */
type MUNDIAPI_IMPL struct {
     subscriptions subscriptions_pkg.SUBSCRIPTIONS
     orders orders_pkg.ORDERS
     plans plans_pkg.PLANS
     invoices invoices_pkg.INVOICES
     customers customers_pkg.CUSTOMERS
     charges charges_pkg.CHARGES
     recipients recipients_pkg.RECIPIENTS
     tokens tokens_pkg.TOKENS
     sellers sellers_pkg.SELLERS
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *MUNDIAPI_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to Subscriptions controller
     * @return Returns the Subscriptions() instance
*/
func (me * MUNDIAPI_IMPL) Subscriptions() subscriptions_pkg.SUBSCRIPTIONS {
    if(me.subscriptions) == nil {
        me.subscriptions = subscriptions_pkg.NewSUBSCRIPTIONS(me.config)
    }
    return me.subscriptions
}
/**
     * Access to Orders controller
     * @return Returns the Orders() instance
*/
func (me * MUNDIAPI_IMPL) Orders() orders_pkg.ORDERS {
    if(me.orders) == nil {
        me.orders = orders_pkg.NewORDERS(me.config)
    }
    return me.orders
}
/**
     * Access to Plans controller
     * @return Returns the Plans() instance
*/
func (me * MUNDIAPI_IMPL) Plans() plans_pkg.PLANS {
    if(me.plans) == nil {
        me.plans = plans_pkg.NewPLANS(me.config)
    }
    return me.plans
}
/**
     * Access to Invoices controller
     * @return Returns the Invoices() instance
*/
func (me * MUNDIAPI_IMPL) Invoices() invoices_pkg.INVOICES {
    if(me.invoices) == nil {
        me.invoices = invoices_pkg.NewINVOICES(me.config)
    }
    return me.invoices
}
/**
     * Access to Customers controller
     * @return Returns the Customers() instance
*/
func (me * MUNDIAPI_IMPL) Customers() customers_pkg.CUSTOMERS {
    if(me.customers) == nil {
        me.customers = customers_pkg.NewCUSTOMERS(me.config)
    }
    return me.customers
}
/**
     * Access to Charges controller
     * @return Returns the Charges() instance
*/
func (me * MUNDIAPI_IMPL) Charges() charges_pkg.CHARGES {
    if(me.charges) == nil {
        me.charges = charges_pkg.NewCHARGES(me.config)
    }
    return me.charges
}
/**
     * Access to Recipients controller
     * @return Returns the Recipients() instance
*/
func (me * MUNDIAPI_IMPL) Recipients() recipients_pkg.RECIPIENTS {
    if(me.recipients) == nil {
        me.recipients = recipients_pkg.NewRECIPIENTS(me.config)
    }
    return me.recipients
}
/**
     * Access to Tokens controller
     * @return Returns the Tokens() instance
*/
func (me * MUNDIAPI_IMPL) Tokens() tokens_pkg.TOKENS {
    if(me.tokens) == nil {
        me.tokens = tokens_pkg.NewTOKENS(me.config)
    }
    return me.tokens
}
/**
     * Access to Sellers controller
     * @return Returns the Sellers() instance
*/
func (me * MUNDIAPI_IMPL) Sellers() sellers_pkg.SELLERS {
    if(me.sellers) == nil {
        me.sellers = sellers_pkg.NewSELLERS(me.config)
    }
    return me.sellers
}
