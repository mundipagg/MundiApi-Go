/*
 * mundiapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package mundiapi

/*
 * Interface for the mundipagg
 */
type Client interface {
	Subscriptions() Subscriptions
	Orders() Orders
	Plans() Plans
	Invoices() Invoices
	Customers() Customers
	Charges() Charges
	Recipients() Recipients
	Tokens() Tokens
	Sellers() Sellers
	Transactions() Transactions
	Transfers() Transfers
	Configuration() Configuration
}

/*
 * Factory for the Client interface returning mundipagg
 */
func NewClient(basicAuthUserName string, basicAuthPassword string) Client {
	config := NewConfiguration()
	config.SetBasicAuthUserName(basicAuthUserName)
	config.SetBasicAuthPassword(basicAuthPassword)

	mundiAPIClient := new(mundipagg)
	mundiAPIClient.config = config
	return mundiAPIClient
}

/*
 * Client structure as interface implementation
 */
type mundipagg struct {
	config        Configuration
	subscriptions Subscriptions
	orders        Orders
	plans         Plans
	invoices      Invoices
	customers     Customers
	charges       Charges
	recipients    Recipients
	tokens        Tokens
	sellers       Sellers
	transactions  Transactions
	transfers     Transfers
}

/**
 * Access to Configuration
 * @return Returns the Configuration instance
 */
func (m *mundipagg) Configuration() Configuration {
	return m.config
}

/**
 * Access to Subscriptions controller
 * @return Returns the Subscriptions() instance
 */
func (m *mundipagg) Subscriptions() Subscriptions {
	if (m.subscriptions) == nil {
		m.subscriptions = NewSubscriptions(m.config)
	}
	return m.subscriptions
}

/**
 * Access to Orders controller
 * @return Returns the Orders() instance
 */
func (m *mundipagg) Orders() Orders {
	if (m.orders) == nil {
		m.orders = NewOrders(m.config)
	}
	return m.orders
}

/**
 * Access to Plans controller
 * @return Returns the Plans() instance
 */
func (m *mundipagg) Plans() Plans {
	if (m.plans) == nil {
		m.plans = NewPlans(m.config)
	}
	return m.plans
}

/**
 * Access to Invoices controller
 * @return Returns the Invoices() instance
 */
func (m *mundipagg) Invoices() Invoices {
	if (m.invoices) == nil {
		m.invoices = NewInvoices(m.config)
	}
	return m.invoices
}

/**
 * Access to Customers controller
 * @return Returns the Customers() instance
 */
func (m *mundipagg) Customers() Customers {
	if (m.customers) == nil {
		m.customers = NewCustomers(m.config)
	}
	return m.customers
}

/**
 * Access to Charges controller
 * @return Returns the Charges() instance
 */
func (m *mundipagg) Charges() Charges {
	if (m.charges) == nil {
		m.charges = NewCharges(m.config)
	}
	return m.charges
}

/**
 * Access to Recipients controller
 * @return Returns the Recipients() instance
 */
func (m *mundipagg) Recipients() Recipients {
	if (m.recipients) == nil {
		m.recipients = NewRecipients(m.config)
	}
	return m.recipients
}

/**
 * Access to Tokens controller
 * @return Returns the Tokens() instance
 */
func (m *mundipagg) Tokens() Tokens {
	if (m.tokens) == nil {
		m.tokens = NewTokens(m.config)
	}
	return m.tokens
}

/**
 * Access to Sellers controller
 * @return Returns the Sellers() instance
 */
func (m *mundipagg) Sellers() Sellers {
	if (m.sellers) == nil {
		m.sellers = NewSellers(m.config)
	}
	return m.sellers
}

/**
 * Access to Transactions controller
 * @return Returns the Transactions() instance
 */
func (m *mundipagg) Transactions() Transactions {
	if (m.transactions) == nil {
		m.transactions = NewTransactions(m.config)
	}
	return m.transactions
}

/**
 * Access to Transfers controller
 * @return Returns the Transfers() instance
 */
func (m *mundipagg) Transfers() Transfers {
	if (m.transfers) == nil {
		m.transfers = NewTransfers(m.config)
	}
	return m.transfers
}
