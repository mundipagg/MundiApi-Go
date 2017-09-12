# Getting started

Mundipagg API

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=MundiAPI-GoLang&projectName=mundiapi_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=mundiapi_lib)

## How to Use

The following section explains how to use the MundiapiLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=mundiapi_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=mundiapi_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=mundiapi_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=mundiapi_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=mundiapi_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=MundiAPI-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=mundiapi_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=mundiapi_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| basicAuthUserName | The username to use with basic authentication |
| basicAuthPassword | The password to use with basic authentication |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [charges_pkg](#charges_pkg)
* [customers_pkg](#customers_pkg)
* [subscriptions_pkg](#subscriptions_pkg)
* [plans_pkg](#plans_pkg)
* [invoices_pkg](#invoices_pkg)
* [orders_pkg](#orders_pkg)
* [tokens_pkg](#tokens_pkg)

## <a name="charges_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".charges_pkg") charges_pkg

### Get instance

Factory for the ``` CHARGES ``` interface can be accessed from the package charges_pkg.

```go
charges := charges_pkg.NewCHARGES()
```

### <a name="get_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetCharge") GetCharge

> Get a charge from its id


```go
func (me *CHARGES_IMPL) GetCharge(chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |


#### Example Usage

```go
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.GetCharge(chargeId)

```


### <a name="retry_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.RetryCharge") RetryCharge

> Retries a charge


```go
func (me *CHARGES_IMPL) RetryCharge(chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |


#### Example Usage

```go
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.RetryCharge(chargeId)

```


### <a name="get_charges"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetCharges") GetCharges

> Lists all charges


```go
func (me *CHARGES_IMPL) GetCharges()(*models_pkg.ListChargesResponse,error)
```

#### Example Usage

```go

var result *models_pkg.ListChargesResponse
result,_ = charges.GetCharges()

```


### <a name="create_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CreateCharge") CreateCharge

> Creates a new charge


```go
func (me *CHARGES_IMPL) CreateCharge(request *models_pkg.CreateChargeRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| request |  ``` Required ```  | Request for creating a charge |


#### Example Usage

```go
var request *models_pkg.CreateChargeRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.CreateCharge(request)

```


### <a name="update_charge_card"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargeCard") UpdateChargeCard

> Updates the card from a charge


```go
func (me *CHARGES_IMPL) UpdateChargeCard(
            chargeId string,
            request *models_pkg.UpdateChargeCardRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| request |  ``` Required ```  | Request for updating a charge's card |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.UpdateChargeCardRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargeCard(chargeId, request)

```


### <a name="update_charge_payment_method"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargePaymentMethod") UpdateChargePaymentMethod

> Updates a charge's payment method


```go
func (me *CHARGES_IMPL) UpdateChargePaymentMethod(
            chargeId string,
            request *models_pkg.UpdateChargePaymentMethodRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| request |  ``` Required ```  | Request for updating the payment method from a charge |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.UpdateChargePaymentMethodRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargePaymentMethod(chargeId, request)

```


### <a name="cancel_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CancelCharge") CancelCharge

> Cancel a charge


```go
func (me *CHARGES_IMPL) CancelCharge(
            chargeId string,
            request *models_pkg.CreateCancelChargeRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| request |  ``` Optional ```  | Request for cancelling a charge |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.CreateCancelChargeRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.CancelCharge(chargeId, request)

```


### <a name="capture_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CaptureCharge") CaptureCharge

> Captures a charge


```go
func (me *CHARGES_IMPL) CaptureCharge(
            chargeId string,
            request *models_pkg.CreateCaptureChargeRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| request |  ``` Optional ```  | Request for capturing a charge |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.CreateCaptureChargeRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.CaptureCharge(chargeId, request)

```


### <a name="update_charge_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargeMetadata") UpdateChargeMetadata

> Updates the metadata from a charge


```go
func (me *CHARGES_IMPL) UpdateChargeMetadata(
            chargeId string,
            request *models_pkg.UpdateMetadataRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | The charge id |
| request |  ``` Required ```  | Request for updating the charge metadata |


#### Example Usage

```go
chargeId := "charge_id"
var request *models_pkg.UpdateMetadataRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargeMetadata(chargeId, request)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="customers_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".customers_pkg") customers_pkg

### Get instance

Factory for the ``` CUSTOMERS ``` interface can be accessed from the package customers_pkg.

```go
customers := customers_pkg.NewCUSTOMERS()
```

### <a name="get_addresses"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAddresses") GetAddresses

> Gets all adressess from a customer


```go
func (me *CUSTOMERS_IMPL) GetAddresses(customerId string)(*models_pkg.ListAddressesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.ListAddressesResponse
result,_ = customers.GetAddresses(customerId)

```


### <a name="get_cards"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCards") GetCards

> Get all cards from a customer


```go
func (me *CUSTOMERS_IMPL) GetCards(customerId string)(*models_pkg.ListCardsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.ListCardsResponse
result,_ = customers.GetCards(customerId)

```


### <a name="get_customers"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCustomers") GetCustomers

> Get all Customers


```go
func (me *CUSTOMERS_IMPL) GetCustomers()(*models_pkg.ListCustomersResponse,error)
```

#### Example Usage

```go

var result *models_pkg.ListCustomersResponse
result,_ = customers.GetCustomers()

```


### <a name="create_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateCustomer") CreateCustomer

> Creates a new customer


```go
func (me *CUSTOMERS_IMPL) CreateCustomer(request *models_pkg.CreateCustomerRequest)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| request |  ``` Required ```  | Request for creating a customer |


#### Example Usage

```go
var request *models_pkg.CreateCustomerRequest

var result *models_pkg.GetCustomerResponse
result,_ = customers.CreateCustomer(request)

```


### <a name="get_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCustomer") GetCustomer

> Get a customer


```go
func (me *CUSTOMERS_IMPL) GetCustomer(customerId string)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.GetCustomerResponse
result,_ = customers.GetCustomer(customerId)

```


### <a name="update_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateAddress") UpdateAddress

> Updates an address


```go
func (me *CUSTOMERS_IMPL) UpdateAddress(
            customerId string,
            addressId string,
            request *models_pkg.UpdateAddressRequest)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| addressId |  ``` Required ```  | Address Id |
| request |  ``` Required ```  | Request for updating an address |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"
var request *models_pkg.UpdateAddressRequest

var result *models_pkg.GetAddressResponse
result,_ = customers.UpdateAddress(customerId, addressId, request)

```


### <a name="update_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCard") UpdateCard

> Updates a card


```go
func (me *CUSTOMERS_IMPL) UpdateCard(
            customerId string,
            cardId string,
            request *models_pkg.UpdateCardRequest)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| cardId |  ``` Required ```  | Card id |
| request |  ``` Required ```  | Request for updating a card |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"
var request *models_pkg.UpdateCardRequest

var result *models_pkg.GetCardResponse
result,_ = customers.UpdateCard(customerId, cardId, request)

```


### <a name="get_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAddress") GetAddress

> Get a customer's address


```go
func (me *CUSTOMERS_IMPL) GetAddress(
            customerId string,
            addressId string)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| addressId |  ``` Required ```  | Address Id |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"

var result *models_pkg.GetAddressResponse
result,_ = customers.GetAddress(customerId, addressId)

```


### <a name="delete_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAddress") DeleteAddress

> Delete a Customer's address


```go
func (me *CUSTOMERS_IMPL) DeleteAddress(
            customerId string,
            addressId string)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| addressId |  ``` Required ```  | Address Id |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"

var result *models_pkg.GetAddressResponse
result,_ = customers.DeleteAddress(customerId, addressId)

```


### <a name="delete_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteCard") DeleteCard

> Delete a customer's card


```go
func (me *CUSTOMERS_IMPL) DeleteCard(
            customerId string,
            cardId string)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| cardId |  ``` Required ```  | Card Id |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"

var result *models_pkg.GetCardResponse
result,_ = customers.DeleteCard(customerId, cardId)

```


### <a name="create_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateAddress") CreateAddress

> Creates a new address for a customer


```go
func (me *CUSTOMERS_IMPL) CreateAddress(
            customerId string,
            request *models_pkg.CreateAddressRequest)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| request |  ``` Required ```  | Request for creating an address |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.CreateAddressRequest

var result *models_pkg.GetAddressResponse
result,_ = customers.CreateAddress(customerId, request)

```


### <a name="get_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCard") GetCard

> Get a customer's card


```go
func (me *CUSTOMERS_IMPL) GetCard(
            customerId string,
            cardId string)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| cardId |  ``` Required ```  | Card id |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"

var result *models_pkg.GetCardResponse
result,_ = customers.GetCard(customerId, cardId)

```


### <a name="create_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateCard") CreateCard

> Creates a new card for a customer


```go
func (me *CUSTOMERS_IMPL) CreateCard(
            customerId string,
            request *models_pkg.CreateCardRequest)(*models_pkg.GetCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| request |  ``` Required ```  | Request for creating a card |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.CreateCardRequest

var result *models_pkg.GetCardResponse
result,_ = customers.CreateCard(customerId, request)

```


### <a name="update_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCustomer") UpdateCustomer

> Updates a customer


```go
func (me *CUSTOMERS_IMPL) UpdateCustomer(
            customerId string,
            request *models_pkg.UpdateCustomerRequest)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| request |  ``` Required ```  | Request for updating a customer |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.UpdateCustomerRequest

var result *models_pkg.GetCustomerResponse
result,_ = customers.UpdateCustomer(customerId, request)

```


### <a name="delete_access_tokens"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAccessTokens") DeleteAccessTokens

> Delete a Customer's access tokens


```go
func (me *CUSTOMERS_IMPL) DeleteAccessTokens(customerId string)(*models_pkg.ListAccessTokensResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.ListAccessTokensResponse
result,_ = customers.DeleteAccessTokens(customerId)

```


### <a name="get_access_tokens"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAccessTokens") GetAccessTokens

> Get all access tokens from a customer


```go
func (me *CUSTOMERS_IMPL) GetAccessTokens(customerId string)(*models_pkg.ListAccessTokensResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.ListAccessTokensResponse
result,_ = customers.GetAccessTokens(customerId)

```


### <a name="delete_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAccessToken") DeleteAccessToken

> Delete a customer's access token


```go
func (me *CUSTOMERS_IMPL) DeleteAccessToken(
            customerId string,
            tokenId string)(*models_pkg.GetAccessTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| tokenId |  ``` Required ```  | Token Id |


#### Example Usage

```go
customerId := "customer_id"
tokenId := "token_id"

var result *models_pkg.GetAccessTokenResponse
result,_ = customers.DeleteAccessToken(customerId, tokenId)

```


### <a name="create_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateAccessToken") CreateAccessToken

> Creates a access token for a customer


```go
func (me *CUSTOMERS_IMPL) CreateAccessToken(
            customerId string,
            request *models_pkg.CreateAccessTokenRequest)(*models_pkg.GetAccessTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| request |  ``` Required ```  | Request for creating a access token |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.CreateAccessTokenRequest

var result *models_pkg.GetAccessTokenResponse
result,_ = customers.CreateAccessToken(customerId, request)

```


### <a name="get_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAccessToken") GetAccessToken

> Get a Customer's access token


```go
func (me *CUSTOMERS_IMPL) GetAccessToken(
            customerId string,
            tokenId string)(*models_pkg.GetAccessTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| tokenId |  ``` Required ```  | Token Id |


#### Example Usage

```go
customerId := "customer_id"
tokenId := "token_id"

var result *models_pkg.GetAccessTokenResponse
result,_ = customers.GetAccessToken(customerId, tokenId)

```


### <a name="update_customer_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCustomerMetadata") UpdateCustomerMetadata

> Updates the metadata a customer


```go
func (me *CUSTOMERS_IMPL) UpdateCustomerMetadata(
            customerId string,
            request *models_pkg.UpdateMetadataRequest)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | The customer id |
| request |  ``` Required ```  | Request for updating the customer metadata |


#### Example Usage

```go
customerId := "customer_id"
var request *models_pkg.UpdateMetadataRequest

var result *models_pkg.GetCustomerResponse
result,_ = customers.UpdateCustomerMetadata(customerId, request)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="subscriptions_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".subscriptions_pkg") subscriptions_pkg

### Get instance

Factory for the ``` SUBSCRIPTIONS ``` interface can be accessed from the package subscriptions_pkg.

```go
subscriptions := subscriptions_pkg.NewSUBSCRIPTIONS()
```

### <a name="update_subscription_billing_date"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionBillingDate") UpdateSubscriptionBillingDate

> Updates the billing date from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionBillingDate(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionBillingDateRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| request |  ``` Required ```  | Request for updating the subscription billing date |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionBillingDateRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionBillingDate(subscriptionId, request)

```


### <a name="create_usage"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateUsage") CreateUsage

> Creates a usage


```go
func (me *SUBSCRIPTIONS_IMPL) CreateUsage(
            subscriptionId string,
            itemId string,
            body *models_pkg.CreateUsageRequest)(*models_pkg.GetUsageResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| itemId |  ``` Required ```  | Item id |
| body |  ``` Required ```  | Request for creating a usage |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
var body *models_pkg.CreateUsageRequest

var result *models_pkg.GetUsageResponse
result,_ = subscriptions.CreateUsage(subscriptionId, itemId, body)

```


### <a name="update_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionItem") UpdateSubscriptionItem

> Updates a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionItem(
            subscriptionId string,
            itemId string,
            body *models_pkg.UpdateSubscriptionItemRequest)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| itemId |  ``` Required ```  | Item id |
| body |  ``` Required ```  | Request for updating a subscription item |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
var body *models_pkg.UpdateSubscriptionItemRequest

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.UpdateSubscriptionItem(subscriptionId, itemId, body)

```


### <a name="get_subscriptions"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptions") GetSubscriptions

> Gets all subscriptions


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptions()(*models_pkg.ListSubscriptionsResponse,error)
```

#### Example Usage

```go

var result *models_pkg.ListSubscriptionsResponse
result,_ = subscriptions.GetSubscriptions()

```


### <a name="update_subscription_card"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionCard") UpdateSubscriptionCard

> Updates the credit card from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionCard(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionCardRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for updating a card |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionCardRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionCard(subscriptionId, request)

```


### <a name="create_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateSubscription") CreateSubscription

> Creates a new subscription


```go
func (me *SUBSCRIPTIONS_IMPL) CreateSubscription(body *models_pkg.CreateSubscriptionRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating a subscription |


#### Example Usage

```go
var body *models_pkg.CreateSubscriptionRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.CreateSubscription(body)

```


### <a name="create_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateSubscriptionItem") CreateSubscriptionItem

> Creates a new Subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) CreateSubscriptionItem(
            subscriptionId string,
            request *models_pkg.CreateSubscriptionItemRequest)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for creating a subscription item |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.CreateSubscriptionItemRequest

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.CreateSubscriptionItem(subscriptionId, request)

```


### <a name="create_discount"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateDiscount") CreateDiscount

> Creates a discount


```go
func (me *SUBSCRIPTIONS_IMPL) CreateDiscount(
            subscriptionId string,
            request *models_pkg.CreateDiscountRequest)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for creating a discount |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.CreateDiscountRequest

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.CreateDiscount(subscriptionId, request)

```


### <a name="get_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscription") GetSubscription

> Gets a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscription(subscriptionId string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.GetSubscription(subscriptionId)

```


### <a name="update_subscription_payment_method"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionPaymentMethod") UpdateSubscriptionPaymentMethod

> Updates the payment method from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionPaymentMethod(
            subscriptionId string,
            request *models_pkg.UpdateSubscriptionPaymentMethodRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Required ```  | Request for updating the paymentmethod from a subscription |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateSubscriptionPaymentMethodRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionPaymentMethod(subscriptionId, request)

```


### <a name="get_usages"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetUsages") GetUsages

> Lists all usages from a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) GetUsages(
            subscriptionId string,
            itemId string)(*models_pkg.ListUsagesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| itemId |  ``` Required ```  | The subscription item id |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"

var result *models_pkg.ListUsagesResponse
result,_ = subscriptions.GetUsages(subscriptionId, itemId)

```


### <a name="delete_usage"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteUsage") DeleteUsage

> Deletes a usage


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteUsage(
            subscriptionId string,
            itemId string,
            usageId string)(*models_pkg.GetUsageResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| itemId |  ``` Required ```  | The subscription item id |
| usageId |  ``` Required ```  | The usage id |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
usageId := "usage_id"

var result *models_pkg.GetUsageResponse
result,_ = subscriptions.DeleteUsage(subscriptionId, itemId, usageId)

```


### <a name="delete_discount"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteDiscount") DeleteDiscount

> Deletes a discount


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteDiscount(
            subscriptionId string,
            discountId string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| discountId |  ``` Required ```  | Discount Id |


#### Example Usage

```go
subscriptionId := "subscription_id"
discountId := "discount_id"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.DeleteDiscount(subscriptionId, discountId)

```


### <a name="cancel_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CancelSubscription") CancelSubscription

> Cancels a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) CancelSubscription(
            subscriptionId string,
            request *models_pkg.CreateCancelSubscriptionRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| request |  ``` Optional ```  | Request for cancelling a subscription |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.CreateCancelSubscriptionRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.CancelSubscription(subscriptionId, request)

```


### <a name="delete_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteSubscriptionItem") DeleteSubscriptionItem

> Deletes a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteSubscriptionItem(
            subscriptionId string,
            subscriptionItemId string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| subscriptionItemId |  ``` Required ```  | Subscription item id |


#### Example Usage

```go
subscriptionId := "subscription_id"
subscriptionItemId := "subscription_item_id"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.DeleteSubscriptionItem(subscriptionId, subscriptionItemId)

```


### <a name="update_subscription_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionMetadata") UpdateSubscriptionMetadata

> Updates the metadata from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionMetadata(
            subscriptionId string,
            request *models_pkg.UpdateMetadataRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| request |  ``` Required ```  | Request for updating the subscrption metadata |


#### Example Usage

```go
subscriptionId := "subscription_id"
var request *models_pkg.UpdateMetadataRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionMetadata(subscriptionId, request)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="plans_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".plans_pkg") plans_pkg

### Get instance

Factory for the ``` PLANS ``` interface can be accessed from the package plans_pkg.

```go
plans := plans_pkg.NewPLANS()
```

### <a name="update_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.UpdatePlanItem") UpdatePlanItem

> Updates a plan item


```go
func (me *PLANS_IMPL) UpdatePlanItem(
            planId string,
            planItemId string,
            body *models_pkg.UpdatePlanItemRequest)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| planItemId |  ``` Required ```  | Plan item id |
| body |  ``` Required ```  | Request for updating the plan item |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"
var body *models_pkg.UpdatePlanItemRequest

var result *models_pkg.GetPlanItemResponse
result,_ = plans.UpdatePlanItem(planId, planItemId, body)

```


### <a name="get_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlan") GetPlan

> Gets a plan


```go
func (me *PLANS_IMPL) GetPlan(planId string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |


#### Example Usage

```go
planId := "plan_id"

var result *models_pkg.GetPlanResponse
result,_ = plans.GetPlan(planId)

```


### <a name="create_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.CreatePlanItem") CreatePlanItem

> Adds a new item to a plan


```go
func (me *PLANS_IMPL) CreatePlanItem(
            planId string,
            request *models_pkg.CreatePlanItemRequest)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| request |  ``` Required ```  | Request for creating a plan item |


#### Example Usage

```go
planId := "plan_id"
var request *models_pkg.CreatePlanItemRequest

var result *models_pkg.GetPlanItemResponse
result,_ = plans.CreatePlanItem(planId, request)

```


### <a name="update_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.UpdatePlan") UpdatePlan

> Updates a plan


```go
func (me *PLANS_IMPL) UpdatePlan(
            planId string,
            request *models_pkg.UpdatePlanRequest)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| request |  ``` Required ```  | Request for updating a plan |


#### Example Usage

```go
planId := "plan_id"
var request *models_pkg.UpdatePlanRequest

var result *models_pkg.GetPlanResponse
result,_ = plans.UpdatePlan(planId, request)

```


### <a name="create_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.CreatePlan") CreatePlan

> Creates a new plan


```go
func (me *PLANS_IMPL) CreatePlan(body *models_pkg.CreatePlanRequest)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating a plan |


#### Example Usage

```go
var body *models_pkg.CreatePlanRequest

var result *models_pkg.GetPlanResponse
result,_ = plans.CreatePlan(body)

```


### <a name="get_plans"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlans") GetPlans

> Gets all plans


```go
func (me *PLANS_IMPL) GetPlans()(*models_pkg.ListPlansResponse,error)
```

#### Example Usage

```go

var result *models_pkg.ListPlansResponse
result,_ = plans.GetPlans()

```


### <a name="delete_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.DeletePlan") DeletePlan

> Deletes a plan


```go
func (me *PLANS_IMPL) DeletePlan(planId string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |


#### Example Usage

```go
planId := "plan_id"

var result *models_pkg.GetPlanResponse
result,_ = plans.DeletePlan(planId)

```


### <a name="get_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlanItem") GetPlanItem

> Gets a plan item


```go
func (me *PLANS_IMPL) GetPlanItem(
            planId string,
            planItemId string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| planItemId |  ``` Required ```  | Plan item id |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.GetPlanItem(planId, planItemId)

```


### <a name="delete_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.DeletePlanItem") DeletePlanItem

> Removes an item from a plan


```go
func (me *PLANS_IMPL) DeletePlanItem(
            planId string,
            planItemId string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| planItemId |  ``` Required ```  | Plan item id |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.DeletePlanItem(planId, planItemId)

```


### <a name="update_plan_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.UpdatePlanMetadata") UpdatePlanMetadata

> Updates the metadata from a plan


```go
func (me *PLANS_IMPL) UpdatePlanMetadata(
            planId string,
            request *models_pkg.UpdateMetadataRequest)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | The plan id |
| request |  ``` Required ```  | Request for updating the plan metadata |


#### Example Usage

```go
planId := "plan_id"
var request *models_pkg.UpdateMetadataRequest

var result *models_pkg.GetPlanResponse
result,_ = plans.UpdatePlanMetadata(planId, request)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="invoices_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".invoices_pkg") invoices_pkg

### Get instance

Factory for the ``` INVOICES ``` interface can be accessed from the package invoices_pkg.

```go
invoices := invoices_pkg.NewINVOICES()
```

### <a name="cancel_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.CancelInvoice") CancelInvoice

> Cancels an invoice


```go
func (me *INVOICES_IMPL) CancelInvoice(invoiceId string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice id |


#### Example Usage

```go
invoiceId := "invoice_id"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.CancelInvoice(invoiceId)

```


### <a name="get_last_invoice_charge"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetLastInvoiceCharge") GetLastInvoiceCharge

> Gets the last charge from an invoice


```go
func (me *INVOICES_IMPL) GetLastInvoiceCharge(invoiceId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice id |


#### Example Usage

```go
invoiceId := "invoice_id"

var result *models_pkg.GetChargeResponse
result,_ = invoices.GetLastInvoiceCharge(invoiceId)

```


### <a name="get_invoices"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetInvoices") GetInvoices

> Gets all invoices


```go
func (me *INVOICES_IMPL) GetInvoices()(*models_pkg.ListInvoicesResponse,error)
```

#### Example Usage

```go

var result *models_pkg.ListInvoicesResponse
result,_ = invoices.GetInvoices()

```


### <a name="get_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetInvoice") GetInvoice

> Gets an invoice


```go
func (me *INVOICES_IMPL) GetInvoice(invoiceId string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice Id |


#### Example Usage

```go
invoiceId := "invoice_id"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.GetInvoice(invoiceId)

```


### <a name="update_invoice_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.UpdateInvoiceMetadata") UpdateInvoiceMetadata

> Updates the metadata from an invoice


```go
func (me *INVOICES_IMPL) UpdateInvoiceMetadata(
            invoiceId string,
            request *models_pkg.UpdateMetadataRequest)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | The invoice id |
| request |  ``` Required ```  | Request for updating the invoice metadata |


#### Example Usage

```go
invoiceId := "invoice_id"
var request *models_pkg.UpdateMetadataRequest

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.UpdateInvoiceMetadata(invoiceId, request)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="orders_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".orders_pkg") orders_pkg

### Get instance

Factory for the ``` ORDERS ``` interface can be accessed from the package orders_pkg.

```go
orders := orders_pkg.NewORDERS()
```

### <a name="get_order"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.GetOrder") GetOrder

> Gets an order


```go
func (me *ORDERS_IMPL) GetOrder(orderId string)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order id |


#### Example Usage

```go
orderId := "order_id"

var result *models_pkg.GetOrderResponse
result,_ = orders.GetOrder(orderId)

```


### <a name="get_orders"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.GetOrders") GetOrders

> Gets all orders


```go
func (me *ORDERS_IMPL) GetOrders()(*models_pkg.ListOrderResponse,error)
```

#### Example Usage

```go

var result *models_pkg.ListOrderResponse
result,_ = orders.GetOrders()

```


### <a name="create_order"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.CreateOrder") CreateOrder

> Creates a new Order


```go
func (me *ORDERS_IMPL) CreateOrder(body *models_pkg.CreateOrderRequest)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating an order |


#### Example Usage

```go
var body *models_pkg.CreateOrderRequest

var result *models_pkg.GetOrderResponse
result,_ = orders.CreateOrder(body)

```


### <a name="update_order_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.UpdateOrderMetadata") UpdateOrderMetadata

> Updates the metadata from an order


```go
func (me *ORDERS_IMPL) UpdateOrderMetadata(
            orderId string,
            request *models_pkg.UpdateMetadataRequest)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | The order id |
| request |  ``` Required ```  | Request for updating the order metadata |


#### Example Usage

```go
orderId := "order_id"
var request *models_pkg.UpdateMetadataRequest

var result *models_pkg.GetOrderResponse
result,_ = orders.UpdateOrderMetadata(orderId, request)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="tokens_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".tokens_pkg") tokens_pkg

### Get instance

Factory for the ``` TOKENS ``` interface can be accessed from the package tokens_pkg.

```go
tokens := tokens_pkg.NewTOKENS()
```

### <a name="get_token"></a>![Method: ](https://apidocs.io/img/method.png ".tokens_pkg.GetToken") GetToken

> *Tags:*  ``` Skips Authentication ``` 

> Gets a token from its id


```go
func (me *TOKENS_IMPL) GetToken(
            id string,
            publicKey string)(*models_pkg.GetTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Token id |
| publicKey |  ``` Required ```  | Public key |


#### Example Usage

```go
id := "id"
publicKey := "public_key"

var result *models_pkg.GetTokenResponse
result,_ = tokens.GetToken(id, publicKey)

```


### <a name="create_token"></a>![Method: ](https://apidocs.io/img/method.png ".tokens_pkg.CreateToken") CreateToken

> *Tags:*  ``` Skips Authentication ``` 

> TODO: Add a method description


```go
func (me *TOKENS_IMPL) CreateToken(
            publicKey string,
            request *models_pkg.CreateTokenRequest)(*models_pkg.GetTokenResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| publicKey |  ``` Required ```  | Public key |
| request |  ``` Required ```  | Request for creating a token |


#### Example Usage

```go
publicKey := "public_key"
var request *models_pkg.CreateTokenRequest

var result *models_pkg.GetTokenResponse
result,_ = tokens.CreateToken(publicKey, request)

```


[Back to List of Controllers](#list_of_controllers)



