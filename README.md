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
| serviceRefererName | TODO: add a description |
| basicAuthUserName | The username to use with basic authentication |
| basicAuthPassword | The password to use with basic authentication |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [subscriptions_pkg](#subscriptions_pkg)
* [orders_pkg](#orders_pkg)
* [plans_pkg](#plans_pkg)
* [invoices_pkg](#invoices_pkg)
* [customers_pkg](#customers_pkg)
* [charges_pkg](#charges_pkg)
* [recipients_pkg](#recipients_pkg)
* [tokens_pkg](#tokens_pkg)
* [transactions_pkg](#transactions_pkg)
* [transfers_pkg](#transfers_pkg)

## <a name="subscriptions_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".subscriptions_pkg") subscriptions_pkg

### Get instance

Factory for the ``` SUBSCRIPTIONS ``` interface can be accessed from the package subscriptions_pkg.

```go
subscriptions := subscriptions_pkg.NewSUBSCRIPTIONS()
```

### <a name="create_discount"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateDiscount") CreateDiscount

> Creates a discount


```go
func (me *SUBSCRIPTIONS_IMPL) CreateDiscount(
            subscriptionId string,
            body *models_pkg.SubscriptionsDiscountsRequest,
            idempotencyKey *string)(*models_pkg.SubscriptionsDiscountsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| body |  ``` Required ```  | Request for creating a discount |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsDiscountsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.SubscriptionsDiscountsResponse
result,_ = subscriptions.CreateDiscount(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionItem") GetSubscriptionItem

> Get Subscription Item


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionItem(
            subscriptionId string,
            itemId string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| itemId |  ``` Required ```  | Item id |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.GetSubscriptionItem(subscriptionId, itemId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionItem") UpdateSubscriptionItem

> Updates a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionItem(
            subscriptionId string,
            itemId string,
            body *models_pkg.SubscriptionsItemsRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| itemId |  ``` Required ```  | Item id |
| body |  ``` Required ```  | Request for updating a subscription item |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
var body *models_pkg.SubscriptionsItemsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.UpdateSubscriptionItem(subscriptionId, itemId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_usage"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteUsage") DeleteUsage

> Deletes a usage


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteUsage(
            subscriptionId string,
            itemId string,
            usageId string,
            idempotencyKey *string)(*models_pkg.SubscriptionsItemsUsagesUsageIdResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| itemId |  ``` Required ```  | The subscription item id |
| usageId |  ``` Required ```  | The usage id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
usageId := "usage_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.SubscriptionsItemsUsagesUsageIdResponse
result,_ = subscriptions.DeleteUsage(subscriptionId, itemId, usageId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="cancel_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CancelSubscription") CancelSubscription

> Cancels a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) CancelSubscription(
            subscriptionId string,
            idempotencyKey *string,
            body *models_pkg.SubscriptionsRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | Request for cancelling a subscription |


#### Example Usage

```go
subscriptionId := "subscription_id"
idempotencyKey := "idempotency-key"
var body *models_pkg.SubscriptionsRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.CancelSubscription(subscriptionId, idempotencyKey, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



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

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_increment"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteIncrement") DeleteIncrement

> Deletes a increment


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteIncrement(
            subscriptionId string,
            incrementId string,
            idempotencyKey *string)(*models_pkg.SubscriptionsIncrementsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| incrementId |  ``` Required ```  | Increment id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
incrementId := "increment_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.SubscriptionsIncrementsResponse
result,_ = subscriptions.DeleteIncrement(subscriptionId, incrementId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_increment_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetIncrementById") GetIncrementById

> GetIncrementById


```go
func (me *SUBSCRIPTIONS_IMPL) GetIncrementById(
            subscriptionId string,
            incrementId string)(*models_pkg.SubscriptionsIncrementsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription Id |
| incrementId |  ``` Required ```  | The increment Id |


#### Example Usage

```go
subscriptionId := "subscription_id"
incrementId := "increment_id"

var result *models_pkg.SubscriptionsIncrementsResponse
result,_ = subscriptions.GetIncrementById(subscriptionId, incrementId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_subscription_cycle_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionCycleById") GetSubscriptionCycleById

> GetSubscriptionCycleById


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionCycleById(
            subscriptionId string,
            cycleId string)(*models_pkg.SubscriptionsCyclesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| cycleId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
cycleId := "cycleId"

var result *models_pkg.SubscriptionsCyclesResponse
result,_ = subscriptions.GetSubscriptionCycleById(subscriptionId, cycleId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_start_at"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionStartAt") UpdateSubscriptionStartAt

> Updates the start at date from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionStartAt(
            subscriptionId string,
            body *models_pkg.SubscriptionsStartAtRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| body |  ``` Required ```  | Request for updating the subscription start date |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsStartAtRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionStartAt(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_payment_method"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionPaymentMethod") UpdateSubscriptionPaymentMethod

> Updates the payment method from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionPaymentMethod(
            subscriptionId string,
            body *models_pkg.SubscriptionsPaymentMethodRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| body |  ``` Required ```  | Request for updating the paymentmethod from a subscription |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsPaymentMethodRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionPaymentMethod(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_current_cycle_status"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateCurrentCycleStatus") UpdateCurrentCycleStatus

> UpdateCurrentCycleStatus


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateCurrentCycleStatus(
            subscriptionId string,
            body *models_pkg.UpdateCurrentCycleStatusRequest,
            idempotencyKey *string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| body |  ``` Required ```  | Request for updating the end date of the subscription current status |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.UpdateCurrentCycleStatusRequest
idempotencyKey := "idempotency-key"

var result 
result,_ = subscriptions.UpdateCurrentCycleStatus(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateSubscription") CreateSubscription

> Creates a new subscription


```go
func (me *SUBSCRIPTIONS_IMPL) CreateSubscription(
            body *models_pkg.SubscriptionsRequest1,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating a subscription |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.SubscriptionsRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.CreateSubscription(body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_subscriptions"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptions") GetSubscriptions

> Gets all subscriptions


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptions(
            page *int64,
            size *int64,
            code *string,
            billingType *string,
            customerId *string,
            planId *string,
            cardId *string,
            status *string,
            nextBillingSince *time.Time,
            nextBillingUntil *time.Time,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.SubscriptionsResponse3,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Filter for subscription's code |
| billingType |  ``` Optional ```  | Filter for subscription's billing type |
| customerId |  ``` Optional ```  | Filter for subscription's customer id |
| planId |  ``` Optional ```  | Filter for subscription's plan id |
| cardId |  ``` Optional ```  | Filter for subscription's card id |
| status |  ``` Optional ```  | Filter for subscription's status |
| nextBillingSince |  ``` Optional ```  | Filter for subscription's next billing date start range |
| nextBillingUntil |  ``` Optional ```  | Filter for subscription's next billing date end range |
| createdSince |  ``` Optional ```  | Filter for subscription's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for subscriptions creation date end range |


#### Example Usage

```go
page,_ := strconv.ParseInt("87", 10, 8)
size,_ := strconv.ParseInt("87", 10, 8)
code := "code"
billingType := "billing_type"
customerId := "customer_id"
planId := "plan_id"
cardId := "card_id"
status := "status"
nextBillingSince := time.Now()
nextBillingUntil := time.Now()
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.SubscriptionsResponse3
result,_ = subscriptions.GetSubscriptions(page, size, code, billingType, customerId, planId, cardId, status, nextBillingSince, nextBillingUntil, createdSince, createdUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_usages_details"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetUsagesDetails") GetUsagesDetails

> GetUsagesDetails


```go
func (me *SUBSCRIPTIONS_IMPL) GetUsagesDetails(
            subscriptionId string,
            cycleId *string,
            size *int64,
            page *int64,
            itemId *string,
            group *string)(*models_pkg.GetUsagesDetailsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Identifier |
| cycleId |  ``` Optional ```  | Cycle id |
| size |  ``` Optional ```  | Page size |
| page |  ``` Optional ```  | Page number |
| itemId |  ``` Optional ```  | Identificador do item |
| group |  ``` Optional ```  | identificador da loja (account) de cada item |


#### Example Usage

```go
subscriptionId := "subscription_id"
cycleId := "cycle_id"
size,_ := strconv.ParseInt("87", 10, 8)
page,_ := strconv.ParseInt("87", 10, 8)
itemId := "item_id"
group := "group"

var result *models_pkg.GetUsagesDetailsResponse
result,_ = subscriptions.GetUsagesDetails(subscriptionId, cycleId, size, page, itemId, group)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="renew_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.RenewSubscription") RenewSubscription

> RenewSubscription


```go
func (me *SUBSCRIPTIONS_IMPL) RenewSubscription(
            subscriptionId string,
            idempotencyKey *string)(*models_pkg.SubscriptionsCyclesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.SubscriptionsCyclesResponse
result,_ = subscriptions.RenewSubscription(subscriptionId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_subscription_cycles"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionCycles") GetSubscriptionCycles

> GetSubscriptionCycles


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionCycles(
            subscriptionId string,
            page string,
            size string)(*models_pkg.SubscriptionsCyclesResponse2,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| page |  ``` Required ```  | Page number |
| size |  ``` Required ```  | Page size |


#### Example Usage

```go
subscriptionId := "subscription_id"
page := "page"
size := "size"

var result *models_pkg.SubscriptionsCyclesResponse2
result,_ = subscriptions.GetSubscriptionCycles(subscriptionId, page, size)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_an_usage"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateAnUsage") CreateAnUsage

> Create Usage


```go
func (me *SUBSCRIPTIONS_IMPL) CreateAnUsage(
            subscriptionId string,
            itemId string,
            idempotencyKey *string)(*models_pkg.SubscriptionsItemsUsagesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| itemId |  ``` Required ```  | Item id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.SubscriptionsItemsUsagesResponse
result,_ = subscriptions.CreateAnUsage(subscriptionId, itemId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_usages"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetUsages") GetUsages

> Lists all usages from a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) GetUsages(
            subscriptionId string,
            itemId string,
            page *int64,
            size *int64,
            code *string,
            group *string,
            usedSince *time.Time,
            usedUntil *time.Time)(*models_pkg.SubscriptionsItemsUsagesResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| itemId |  ``` Required ```  | The subscription item id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Identification code in the client system |
| group |  ``` Optional ```  | Identification group in the client system |
| usedSince |  ``` Optional ```  | TODO: Add a parameter description |
| usedUntil |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
page,_ := strconv.ParseInt("87", 10, 8)
size,_ := strconv.ParseInt("87", 10, 8)
code := "code"
group := "group"
usedSince := time.Now()
usedUntil := time.Now()

var result *models_pkg.SubscriptionsItemsUsagesResponse1
result,_ = subscriptions.GetUsages(subscriptionId, itemId, page, size, code, group, usedSince, usedUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_discount"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteDiscount") DeleteDiscount

> Deletes a discount


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteDiscount(
            subscriptionId string,
            discountId string,
            idempotencyKey *string)(*models_pkg.SubscriptionsDiscountsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| discountId |  ``` Required ```  | Discount Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
discountId := "discount_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.SubscriptionsDiscountsResponse
result,_ = subscriptions.DeleteDiscount(subscriptionId, discountId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_increments"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetIncrements") GetIncrements

> GetIncrements


```go
func (me *SUBSCRIPTIONS_IMPL) GetIncrements(
            subscriptionId string,
            page *int64,
            size *int64)(*models_pkg.ListIncrementsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
subscriptionId := "subscription_id"
page,_ := strconv.ParseInt("87", 10, 8)
size,_ := strconv.ParseInt("87", 10, 8)

var result *models_pkg.ListIncrementsResponse
result,_ = subscriptions.GetIncrements(subscriptionId, page, size)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateSubscriptionItem") CreateSubscriptionItem

> Creates a new Subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) CreateSubscriptionItem(
            subscriptionId string,
            body *models_pkg.SubscriptionsItemsRequest1,
            idempotencyKey *string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| body |  ``` Required ```  | Request for creating a subscription item |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsItemsRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.CreateSubscriptionItem(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_subscription_items"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionItems") GetSubscriptionItems

> Get Subscription Items


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionItems(
            subscriptionId string,
            page *int64,
            size *int64,
            name *string,
            code *string,
            status *string,
            description *string,
            createdSince *string,
            createdUntil *string)(*models_pkg.SubscriptionsItemsResponse3,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| name |  ``` Optional ```  | The item name |
| code |  ``` Optional ```  | Identification code in the client system |
| status |  ``` Optional ```  | The item statis |
| description |  ``` Optional ```  | The item description |
| createdSince |  ``` Optional ```  | Filter for item's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for item's creation date end range |


#### Example Usage

```go
subscriptionId := "subscription_id"
page,_ := strconv.ParseInt("87", 10, 8)
size,_ := strconv.ParseInt("87", 10, 8)
name := "name"
code := "code"
status := "status"
description := "description"
createdSince := "created_since"
createdUntil := "created_until"

var result *models_pkg.SubscriptionsItemsResponse3
result,_ = subscriptions.GetSubscriptionItems(subscriptionId, page, size, name, code, status, description, createdSince, createdUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_billing_date"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionBillingDate") UpdateSubscriptionBillingDate

> Updates the billing date from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionBillingDate(
            subscriptionId string,
            body *models_pkg.SubscriptionsBillingDateRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| body |  ``` Required ```  | Request for updating the subscription billing date |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsBillingDateRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionBillingDate(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_latest_period_end_at"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateLatestPeriodEndAt") UpdateLatestPeriodEndAt

> UpdateLatestPeriodEndAt


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateLatestPeriodEndAt(
            subscriptionId string,
            body *models_pkg.SubscriptionsPeriodsLatestEndAtRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | Request for updating the end date of the current signature cycle |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsPeriodsLatestEndAtRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateLatestPeriodEndAt(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_affiliation_id"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionAffiliationId") UpdateSubscriptionAffiliationId

> UpdateSubscriptionAffiliationId


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionAffiliationId(
            subscriptionId string,
            body *models_pkg.SubscriptionsGatewayAffiliationIdRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | Request for updating a subscription affiliation id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsGatewayAffiliationIdRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionAffiliationId(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_subscription_item"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.DeleteSubscriptionItem") DeleteSubscriptionItem

> Deletes a subscription item


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteSubscriptionItem(
            subscriptionId string,
            subscriptionItemId string,
            idempotencyKey *string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| subscriptionItemId |  ``` Required ```  | Subscription item id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
subscriptionItemId := "subscription_item_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.DeleteSubscriptionItem(subscriptionId, subscriptionItemId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_card"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionCard") UpdateSubscriptionCard

> Updates the credit card from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionCard(
            subscriptionId string,
            body *models_pkg.SubscriptionsCardRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| body |  ``` Required ```  | Request for updating a card |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsCardRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionCard(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionMetadata") UpdateSubscriptionMetadata

> Updates the metadata from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionMetadata(
            subscriptionId string,
            body *models_pkg.SubscriptionsMetadataRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| body |  ``` Required ```  | Request for updating the subscrption metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionMetadata(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_due_days"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionDueDays") UpdateSubscriptionDueDays

> Updates the boleto due days from a subscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionDueDays(
            subscriptionId string,
            body *models_pkg.UpdateSubscriptionDueDaysRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| body |  ``` Required ```  | TODO: Add a parameter description |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.UpdateSubscriptionDueDaysRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionDueDays(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_discounts"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetDiscounts") GetDiscounts

> GetDiscounts


```go
func (me *SUBSCRIPTIONS_IMPL) GetDiscounts(
            subscriptionId string,
            page int64,
            size int64)(*models_pkg.ListDiscountsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| page |  ``` Required ```  | Page number |
| size |  ``` Required ```  | Page size |


#### Example Usage

```go
subscriptionId := "subscription_id"
page,_ := strconv.ParseInt("87", 10, 8)
size,_ := strconv.ParseInt("87", 10, 8)

var result *models_pkg.ListDiscountsResponse
result,_ = subscriptions.GetDiscounts(subscriptionId, page, size)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_increment"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.CreateIncrement") CreateIncrement

> Creates a increment


```go
func (me *SUBSCRIPTIONS_IMPL) CreateIncrement(
            subscriptionId string,
            body *models_pkg.SubscriptionsIncrementsRequest,
            idempotencyKey *string)(*models_pkg.SubscriptionsIncrementsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription id |
| body |  ``` Required ```  | Request for creating a increment |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsIncrementsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.SubscriptionsIncrementsResponse
result,_ = subscriptions.CreateIncrement(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_discount_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetDiscountById") GetDiscountById

> GetDiscountById


```go
func (me *SUBSCRIPTIONS_IMPL) GetDiscountById(
            subscriptionId string,
            discountId string)(*models_pkg.SubscriptionsDiscountsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription id |
| discountId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
discountId := "discountId"

var result *models_pkg.SubscriptionsDiscountsResponse
result,_ = subscriptions.GetDiscountById(subscriptionId, discountId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_subscription_minium_price"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionMiniumPrice") UpdateSubscriptionMiniumPrice

> Atualização do valor mínimo da assinatura


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionMiniumPrice(
            subscriptionId string,
            body *models_pkg.SubscriptionsMinimumPriceRequest,
            idempotencyKey *string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| body |  ``` Required ```  | Request da requisição com o valor mínimo que será configurado |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.SubscriptionsMinimumPriceRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionMiniumPrice(subscriptionId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_usage_report"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.GetUsageReport") GetUsageReport

> GetUsageReport


```go
func (me *SUBSCRIPTIONS_IMPL) GetUsageReport(
            subscriptionId string,
            periodId string)(*models_pkg.GetUsageReportResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | The subscription Id |
| periodId |  ``` Required ```  | The period Id |


#### Example Usage

```go
subscriptionId := "subscription_id"
periodId := "period_id"

var result *models_pkg.GetUsageReportResponse
result,_ = subscriptions.GetUsageReport(subscriptionId, periodId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_split_subscription"></a>![Method: ](https://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSplitSubscription") UpdateSplitSubscription

> UpdateSplitSubscription


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSplitSubscription(
            id string,
            body *models_pkg.UpdateSubscriptionSplitRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Subscription's id |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
id := "id"
var body *models_pkg.UpdateSubscriptionSplitRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSplitSubscription(id, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="orders_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".orders_pkg") orders_pkg

### Get instance

Factory for the ``` ORDERS ``` interface can be accessed from the package orders_pkg.

```go
orders := orders_pkg.NewORDERS()
```

### <a name="update_order_status"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.UpdateOrderStatus") UpdateOrderStatus

> UpdateOrderStatus


```go
func (me *ORDERS_IMPL) UpdateOrderStatus(
            id string,
            body *models_pkg.UpdateOrderStatusRequest,
            idempotencyKey *string)(*models_pkg.OrdersClosedResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Order Id |
| body |  ``` Required ```  | Update Order Model |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
id := "id"
var body *models_pkg.UpdateOrderStatusRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.OrdersClosedResponse
result,_ = orders.UpdateOrderStatus(id, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_all_order_items"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.DeleteAllOrderItems") DeleteAllOrderItems

> DeleteAllOrderItems


```go
func (me *ORDERS_IMPL) DeleteAllOrderItems(
            orderId string,
            idempotencyKey *string)(*models_pkg.OrdersItemsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "orderId"
idempotencyKey := "idempotency-key"

var result *models_pkg.OrdersItemsResponse
result,_ = orders.DeleteAllOrderItems(orderId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_order_item"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.CreateOrderItem") CreateOrderItem

> CreateOrderItem


```go
func (me *ORDERS_IMPL) CreateOrderItem(
            orderId string,
            body *models_pkg.OrdersItemsRequest,
            idempotencyKey *string)(*models_pkg.OrdersItemsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| body |  ``` Required ```  | Order Item Model |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "orderId"
var body *models_pkg.OrdersItemsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.OrdersItemsResponse1
result,_ = orders.CreateOrderItem(orderId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_order_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.UpdateOrderMetadata") UpdateOrderMetadata

> Updates the metadata from an order


```go
func (me *ORDERS_IMPL) UpdateOrderMetadata(
            orderId string,
            body *models_pkg.OrdersMetadataRequest,
            idempotencyKey *string)(*models_pkg.OrdersMetadataResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | The order id |
| body |  ``` Required ```  | Request for updating the order metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "order_id"
var body *models_pkg.OrdersMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.OrdersMetadataResponse
result,_ = orders.UpdateOrderMetadata(orderId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_orders"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.GetOrders") GetOrders

> Gets all orders


```go
func (me *ORDERS_IMPL) GetOrders(
            page *int64,
            size *int64,
            code *string,
            status *string,
            createdSince *time.Time,
            createdUntil *time.Time,
            customerId *string)(*models_pkg.OrdersResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Filter for order's code |
| status |  ``` Optional ```  | Filter for order's status |
| createdSince |  ``` Optional ```  | Filter for order's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for order's creation date end range |
| customerId |  ``` Optional ```  | Filter for order's customer id |


#### Example Usage

```go
page,_ := strconv.ParseInt("179", 10, 8)
size,_ := strconv.ParseInt("179", 10, 8)
code := "code"
status := "status"
createdSince := time.Now()
createdUntil := time.Now()
customerId := "customer_id"

var result *models_pkg.OrdersResponse
result,_ = orders.GetOrders(page, size, code, status, createdSince, createdUntil, customerId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_order"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.CreateOrder") CreateOrder

> Creates a new Order


```go
func (me *ORDERS_IMPL) CreateOrder(
            body *models_pkg.OrdersRequest,
            idempotencyKey *string)(*models_pkg.OrdersResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating an order |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.OrdersRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.OrdersResponse1
result,_ = orders.CreateOrder(body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_order_item"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.DeleteOrderItem") DeleteOrderItem

> DeleteOrderItem


```go
func (me *ORDERS_IMPL) DeleteOrderItem(
            orderId string,
            itemId string,
            idempotencyKey *string)(*models_pkg.OrdersItemsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| itemId |  ``` Required ```  | Item Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "orderId"
itemId := "itemId"
idempotencyKey := "idempotency-key"

var result *models_pkg.OrdersItemsResponse1
result,_ = orders.DeleteOrderItem(orderId, itemId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_order_item"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.GetOrderItem") GetOrderItem

> GetOrderItem


```go
func (me *ORDERS_IMPL) GetOrderItem(
            orderId string,
            itemId string)(*models_pkg.OrdersItemsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| itemId |  ``` Required ```  | Item Id |


#### Example Usage

```go
orderId := "orderId"
itemId := "itemId"

var result *models_pkg.OrdersItemsResponse1
result,_ = orders.GetOrderItem(orderId, itemId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_order_item"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.UpdateOrderItem") UpdateOrderItem

> UpdateOrderItem


```go
func (me *ORDERS_IMPL) UpdateOrderItem(
            orderId string,
            itemId string,
            body *models_pkg.OrdersItemsRequest1,
            idempotencyKey *string)(*models_pkg.OrdersItemsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order Id |
| itemId |  ``` Required ```  | Item Id |
| body |  ``` Required ```  | Item Model |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "orderId"
itemId := "itemId"
var body *models_pkg.OrdersItemsRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.OrdersItemsResponse1
result,_ = orders.UpdateOrderItem(orderId, itemId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_order"></a>![Method: ](https://apidocs.io/img/method.png ".orders_pkg.GetOrder") GetOrder

> Gets an order


```go
func (me *ORDERS_IMPL) GetOrder(orderId string)(*models_pkg.OrdersResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | Order id |


#### Example Usage

```go
orderId := "order_id"

var result *models_pkg.OrdersResponse1
result,_ = orders.GetOrder(orderId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



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
            body *models_pkg.PlansItemsRequest,
            idempotencyKey *string)(*models_pkg.PlansItemsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| planItemId |  ``` Required ```  | Plan item id |
| body |  ``` Required ```  | Request for updating the plan item |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"
var body *models_pkg.PlansItemsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.PlansItemsResponse
result,_ = plans.UpdatePlanItem(planId, planItemId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.DeletePlanItem") DeletePlanItem

> Removes an item from a plan


```go
func (me *PLANS_IMPL) DeletePlanItem(
            planId string,
            planItemId string,
            idempotencyKey *string)(*models_pkg.PlansItemsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| planItemId |  ``` Required ```  | Plan item id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.PlansItemsResponse
result,_ = plans.DeletePlanItem(planId, planItemId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlanItem") GetPlanItem

> Gets a plan item


```go
func (me *PLANS_IMPL) GetPlanItem(
            planId string,
            planItemId string)(*models_pkg.PlansItemsResponse,error)
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

var result *models_pkg.PlansItemsResponse
result,_ = plans.GetPlanItem(planId, planItemId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_plan_item"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.CreatePlanItem") CreatePlanItem

> Adds a new item to a plan


```go
func (me *PLANS_IMPL) CreatePlanItem(
            planId string,
            body *models_pkg.PlansItemsRequest1,
            idempotencyKey *string)(*models_pkg.PlansItemsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| body |  ``` Required ```  | Request for creating a plan item |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
var body *models_pkg.PlansItemsRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.PlansItemsResponse
result,_ = plans.CreatePlanItem(planId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_plans"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlans") GetPlans

> Gets all plans


```go
func (me *PLANS_IMPL) GetPlans(
            page *int64,
            size *int64,
            name *string,
            status *string,
            billingType *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.PlansResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| name |  ``` Optional ```  | Filter for Plan's name |
| status |  ``` Optional ```  | Filter for Plan's status |
| billingType |  ``` Optional ```  | Filter for plan's billing type |
| createdSince |  ``` Optional ```  | Filter for plan's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for plan's creation date end range |


#### Example Usage

```go
page,_ := strconv.ParseInt("179", 10, 8)
size,_ := strconv.ParseInt("179", 10, 8)
name := "name"
status := "status"
billingType := "billing_type"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.PlansResponse
result,_ = plans.GetPlans(page, size, name, status, billingType, createdSince, createdUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.CreatePlan") CreatePlan

> Creates a new plan


```go
func (me *PLANS_IMPL) CreatePlan(
            body *models_pkg.PlansRequest,
            idempotencyKey *string)(*models_pkg.PlansResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating a plan |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.PlansRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.PlansResponse1
result,_ = plans.CreatePlan(body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.GetPlan") GetPlan

> Gets a plan


```go
func (me *PLANS_IMPL) GetPlan(planId string)(*models_pkg.PlansResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |


#### Example Usage

```go
planId := "plan_id"

var result *models_pkg.PlansResponse1
result,_ = plans.GetPlan(planId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.UpdatePlan") UpdatePlan

> Updates a plan


```go
func (me *PLANS_IMPL) UpdatePlan(
            planId string,
            body *models_pkg.PlansRequest1,
            idempotencyKey *string)(*models_pkg.PlansResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| body |  ``` Required ```  | Request for updating a plan |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
var body *models_pkg.PlansRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.PlansResponse1
result,_ = plans.UpdatePlan(planId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_plan"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.DeletePlan") DeletePlan

> Deletes a plan


```go
func (me *PLANS_IMPL) DeletePlan(
            planId string,
            idempotencyKey *string)(*models_pkg.PlansResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | Plan id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.PlansResponse1
result,_ = plans.DeletePlan(planId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_plan_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".plans_pkg.UpdatePlanMetadata") UpdatePlanMetadata

> Updates the metadata from a plan


```go
func (me *PLANS_IMPL) UpdatePlanMetadata(
            planId string,
            body *models_pkg.PlansMetadataRequest,
            idempotencyKey *string)(*models_pkg.PlansMetadataResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | The plan id |
| body |  ``` Required ```  | Request for updating the plan metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
var body *models_pkg.PlansMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.PlansMetadataResponse
result,_ = plans.UpdatePlanMetadata(planId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="invoices_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".invoices_pkg") invoices_pkg

### Get instance

Factory for the ``` INVOICES ``` interface can be accessed from the package invoices_pkg.

```go
invoices := invoices_pkg.NewINVOICES()
```

### <a name="create_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.CreateInvoice") CreateInvoice

> Create an Invoice


```go
func (me *INVOICES_IMPL) CreateInvoice(
            subscriptionId string,
            cycleId string,
            idempotencyKey *string,
            body *models_pkg.SubscriptionsCyclesPayRequest)(*models_pkg.SubscriptionsCyclesPayResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |
| cycleId |  ``` Required ```  | Cycle Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
cycleId := "cycle_id"
idempotencyKey := "idempotency-key"
var body *models_pkg.SubscriptionsCyclesPayRequest

var result *models_pkg.SubscriptionsCyclesPayResponse
result,_ = invoices.CreateInvoice(subscriptionId, cycleId, idempotencyKey, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_partial_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetPartialInvoice") GetPartialInvoice

> GetPartialInvoice


```go
func (me *INVOICES_IMPL) GetPartialInvoice(subscriptionId string)(*models_pkg.SubscriptionsPartialInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | Subscription Id |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result *models_pkg.SubscriptionsPartialInvoiceResponse
result,_ = invoices.GetPartialInvoice(subscriptionId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_invoice_status"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.UpdateInvoiceStatus") UpdateInvoiceStatus

> Updates the status from an invoice


```go
func (me *INVOICES_IMPL) UpdateInvoiceStatus(
            invoiceId string,
            body *models_pkg.UpdateCurrentCycleStatusRequest,
            idempotencyKey *string)(*models_pkg.InvoicesStatusResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice Id |
| body |  ``` Required ```  | Request for updating an invoice's status |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"
var body *models_pkg.UpdateCurrentCycleStatusRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.InvoicesStatusResponse
result,_ = invoices.UpdateInvoiceStatus(invoiceId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetInvoice") GetInvoice

> Gets an invoice


```go
func (me *INVOICES_IMPL) GetInvoice(invoiceId string)(*models_pkg.InvoicesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice Id |


#### Example Usage

```go
invoiceId := "invoice_id"

var result *models_pkg.InvoicesResponse
result,_ = invoices.GetInvoice(invoiceId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="cancel_invoice"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.CancelInvoice") CancelInvoice

> Cancels an invoice


```go
func (me *INVOICES_IMPL) CancelInvoice(
            invoiceId string,
            idempotencyKey *string)(*models_pkg.InvoicesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | Invoice id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.InvoicesResponse
result,_ = invoices.CancelInvoice(invoiceId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_invoice_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.UpdateInvoiceMetadata") UpdateInvoiceMetadata

> Updates the metadata from an invoice


```go
func (me *INVOICES_IMPL) UpdateInvoiceMetadata(
            invoiceId string,
            body *models_pkg.InvoicesMetadataRequest,
            idempotencyKey *string)(*models_pkg.InvoicesMetadataResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | The invoice id |
| body |  ``` Required ```  | Request for updating the invoice metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"
var body *models_pkg.InvoicesMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.InvoicesMetadataResponse
result,_ = invoices.UpdateInvoiceMetadata(invoiceId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_invoices"></a>![Method: ](https://apidocs.io/img/method.png ".invoices_pkg.GetInvoices") GetInvoices

> Gets all invoices


```go
func (me *INVOICES_IMPL) GetInvoices(
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
            customerDocument *string)(*models_pkg.InvoicesResponse2,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Filter for Invoice's code |
| customerId |  ``` Optional ```  | Filter for Invoice's customer id |
| subscriptionId |  ``` Optional ```  | Filter for Invoice's subscription id |
| createdSince |  ``` Optional ```  | Filter for Invoice's creation date start range |
| createdUntil |  ``` Optional ```  | Filter for Invoices creation date end range |
| status |  ``` Optional ```  | Filter for Invoice's status |
| dueSince |  ``` Optional ```  | Filter for Invoice's due date start range |
| dueUntil |  ``` Optional ```  | Filter for Invoice's due date end range |
| customerDocument |  ``` Optional ```  | Fillter for invoice's document |


#### Example Usage

```go
page,_ := strconv.ParseInt("179", 10, 8)
size,_ := strconv.ParseInt("179", 10, 8)
code := "code"
customerId := "customer_id"
subscriptionId := "subscription_id"
createdSince := time.Now()
createdUntil := time.Now()
status := "status"
dueSince := time.Now()
dueUntil := time.Now()
customerDocument := "customer_document"

var result *models_pkg.InvoicesResponse2
result,_ = invoices.GetInvoices(page, size, code, customerId, subscriptionId, createdSince, createdUntil, status, dueSince, dueUntil, customerDocument)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="customers_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".customers_pkg") customers_pkg

### Get instance

Factory for the ``` CUSTOMERS ``` interface can be accessed from the package customers_pkg.

```go
customers := customers_pkg.NewCUSTOMERS()
```

### <a name="create_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateAccessToken") CreateAccessToken

> Creates a access token for a customer


```go
func (me *CUSTOMERS_IMPL) CreateAccessToken(
            customerId string,
            body *models_pkg.CustomersAccessTokensRequest,
            idempotencyKey *string)(*models_pkg.CustomersAccessTokensResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| body |  ``` Required ```  | Request for creating a access token |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var body *models_pkg.CustomersAccessTokensRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersAccessTokensResponse
result,_ = customers.CreateAccessToken(customerId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_access_tokens"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAccessTokens") GetAccessTokens

> Get all access tokens from a customer


```go
func (me *CUSTOMERS_IMPL) GetAccessTokens(
            customerId string,
            page *int64,
            size *int64)(*models_pkg.CustomersAccessTokensResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
customerId := "customer_id"
page,_ := strconv.ParseInt("179", 10, 8)
size,_ := strconv.ParseInt("179", 10, 8)

var result *models_pkg.CustomersAccessTokensResponse1
result,_ = customers.GetAccessTokens(customerId, page, size)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCustomer") UpdateCustomer

> Updates a customer


```go
func (me *CUSTOMERS_IMPL) UpdateCustomer(
            customerId string,
            body *models_pkg.CustomersRequest,
            idempotencyKey *string)(*models_pkg.CustomersResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| body |  ``` Required ```  | Request for updating a customer |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var body *models_pkg.CustomersRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersResponse
result,_ = customers.UpdateCustomer(customerId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCustomer") GetCustomer

> Get a customer


```go
func (me *CUSTOMERS_IMPL) GetCustomer(customerId string)(*models_pkg.CustomersResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.CustomersResponse
result,_ = customers.GetCustomer(customerId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_access_tokens"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAccessTokens") DeleteAccessTokens

> Delete a Customer's access tokens


```go
func (me *CUSTOMERS_IMPL) DeleteAccessTokens(customerId string)(*models_pkg.CustomersAccessTokensResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.CustomersAccessTokensResponse1
result,_ = customers.DeleteAccessTokens(customerId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_addresses"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAddresses") GetAddresses

> Gets all adressess from a customer


```go
func (me *CUSTOMERS_IMPL) GetAddresses(
            customerId string,
            page *int64,
            size *int64)(*models_pkg.CustomersAddressesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
customerId := "customer_id"
page,_ := strconv.ParseInt("179", 10, 8)
size,_ := strconv.ParseInt("179", 10, 8)

var result *models_pkg.CustomersAddressesResponse
result,_ = customers.GetAddresses(customerId, page, size)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateAddress") CreateAddress

> Creates a new address for a customer


```go
func (me *CUSTOMERS_IMPL) CreateAddress(
            customerId string,
            body *models_pkg.CustomersAddressesRequest,
            idempotencyKey *string)(*models_pkg.CustomersAddressesResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| body |  ``` Required ```  | Request for creating an address |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var body *models_pkg.CustomersAddressesRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersAddressesResponse1
result,_ = customers.CreateAddress(customerId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAccessToken") GetAccessToken

> Get a Customer's access token


```go
func (me *CUSTOMERS_IMPL) GetAccessToken(
            customerId string,
            tokenId string)(*models_pkg.CustomersAccessTokensResponse,error)
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

var result *models_pkg.CustomersAccessTokensResponse
result,_ = customers.GetAccessToken(customerId, tokenId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_access_token"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAccessToken") DeleteAccessToken

> Delete a customer's access token


```go
func (me *CUSTOMERS_IMPL) DeleteAccessToken(
            customerId string,
            tokenId string,
            idempotencyKey *string)(*models_pkg.CustomersAccessTokensResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| tokenId |  ``` Required ```  | Token Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
tokenId := "token_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersAccessTokensResponse
result,_ = customers.DeleteAccessToken(customerId, tokenId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetAddress") GetAddress

> Get a customer's address


```go
func (me *CUSTOMERS_IMPL) GetAddress(
            customerId string,
            addressId string)(*models_pkg.CustomersAddressesResponse1,error)
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

var result *models_pkg.CustomersAddressesResponse1
result,_ = customers.GetAddress(customerId, addressId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateAddress") UpdateAddress

> Updates an address


```go
func (me *CUSTOMERS_IMPL) UpdateAddress(
            customerId string,
            addressId string,
            body *models_pkg.CustomersAddressesRequest1,
            idempotencyKey *string)(*models_pkg.CustomersAddressesResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| addressId |  ``` Required ```  | Address Id |
| body |  ``` Required ```  | Request for updating an address |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"
var body *models_pkg.CustomersAddressesRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersAddressesResponse1
result,_ = customers.UpdateAddress(customerId, addressId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_address"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteAddress") DeleteAddress

> Delete a Customer's address


```go
func (me *CUSTOMERS_IMPL) DeleteAddress(
            customerId string,
            addressId string,
            idempotencyKey *string)(*models_pkg.CustomersAddressesResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| addressId |  ``` Required ```  | Address Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersAddressesResponse1
result,_ = customers.DeleteAddress(customerId, addressId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateCard") CreateCard

> Creates a new card for a customer


```go
func (me *CUSTOMERS_IMPL) CreateCard(
            customerId string,
            body *models_pkg.CustomersCardsRequest,
            idempotencyKey *string)(*models_pkg.CustomersCardsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| body |  ``` Required ```  | Request for creating a card |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var body *models_pkg.CustomersCardsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersCardsResponse
result,_ = customers.CreateCard(customerId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_cards"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCards") GetCards

> Get all cards from a customer


```go
func (me *CUSTOMERS_IMPL) GetCards(
            customerId string,
            page *int64,
            size *int64)(*models_pkg.CustomersCardsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
customerId := "customer_id"
page,_ := strconv.ParseInt("15", 10, 8)
size,_ := strconv.ParseInt("15", 10, 8)

var result *models_pkg.CustomersCardsResponse1
result,_ = customers.GetCards(customerId, page, size)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="renew_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.RenewCard") RenewCard

> Renew a card


```go
func (me *CUSTOMERS_IMPL) RenewCard(
            customerId string,
            cardId string,
            idempotencyKey *string)(*models_pkg.CustomersCardsRenewResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer id |
| cardId |  ``` Required ```  | Card Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersCardsRenewResponse
result,_ = customers.RenewCard(customerId, cardId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_customer"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.CreateCustomer") CreateCustomer

> Creates a new customer


```go
func (me *CUSTOMERS_IMPL) CreateCustomer(
            body *models_pkg.CustomersRequest1,
            idempotencyKey *string)(*models_pkg.CustomersResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating a customer |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CustomersRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersResponse
result,_ = customers.CreateCustomer(body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_customers"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCustomers") GetCustomers

> Get all Customers


```go
func (me *CUSTOMERS_IMPL) GetCustomers(
            name *string,
            document *string,
            page *int64,
            size *int64,
            email *string,
            code *string)(*models_pkg.CustomersResponse3,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| name |  ``` Optional ```  | Name of the Customer |
| document |  ``` Optional ```  | Document of the Customer |
| page |  ``` Optional ```  ``` DefaultValue ```  | Current page the the search |
| size |  ``` Optional ```  ``` DefaultValue ```  | Quantity pages of the search |
| email |  ``` Optional ```  | Customer's email |
| code |  ``` Optional ```  | Customer's code |


#### Example Usage

```go
name := "name"
document := "document"
page,_ := strconv.ParseInt("1", 10, 8)
size,_ := strconv.ParseInt("10", 10, 8)
email := "email"
code := "Code"

var result *models_pkg.CustomersResponse3
result,_ = customers.GetCustomers(name, document, page, size, email, code)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_customer_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCustomerMetadata") UpdateCustomerMetadata

> Updates the metadata a customer


```go
func (me *CUSTOMERS_IMPL) UpdateCustomerMetadata(
            customerId string,
            body *models_pkg.CustomersMetadataRequest,
            idempotencyKey *string)(*models_pkg.CustomersMetadataResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | The customer id |
| body |  ``` Required ```  | Request for updating the customer metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var body *models_pkg.CustomersMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersMetadataResponse
result,_ = customers.UpdateCustomerMetadata(customerId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.UpdateCard") UpdateCard

> Updates a card


```go
func (me *CUSTOMERS_IMPL) UpdateCard(
            customerId string,
            cardId string,
            body *models_pkg.CustomersCardsRequest1,
            idempotencyKey *string)(*models_pkg.CustomersCardsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| cardId |  ``` Required ```  | Card id |
| body |  ``` Required ```  | Request for updating a card |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"
var body *models_pkg.CustomersCardsRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersCardsResponse
result,_ = customers.UpdateCard(customerId, cardId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="delete_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.DeleteCard") DeleteCard

> Delete a customer's card


```go
func (me *CUSTOMERS_IMPL) DeleteCard(
            customerId string,
            cardId string,
            idempotencyKey *string)(*models_pkg.CustomersCardsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | Customer Id |
| cardId |  ``` Required ```  | Card Id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.CustomersCardsResponse
result,_ = customers.DeleteCard(customerId, cardId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_card"></a>![Method: ](https://apidocs.io/img/method.png ".customers_pkg.GetCard") GetCard

> Get a customer's card


```go
func (me *CUSTOMERS_IMPL) GetCard(
            customerId string,
            cardId string)(*models_pkg.CustomersCardsResponse,error)
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

var result *models_pkg.CustomersCardsResponse
result,_ = customers.GetCard(customerId, cardId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="charges_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".charges_pkg") charges_pkg

### Get instance

Factory for the ``` CHARGES ``` interface can be accessed from the package charges_pkg.

```go
charges := charges_pkg.NewCHARGES()
```

### <a name="get_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetCharge") GetCharge

> Get a charge from its id


```go
func (me *CHARGES_IMPL) GetCharge(chargeId string)(*models_pkg.ChargesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |


#### Example Usage

```go
chargeId := "charge_id"

var result *models_pkg.ChargesResponse
result,_ = charges.GetCharge(chargeId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="cancel_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CancelCharge") CancelCharge

> Cancel a charge


```go
func (me *CHARGES_IMPL) CancelCharge(
            chargeId string,
            idempotencyKey *string,
            body *models_pkg.ChargesRequest)(*models_pkg.ChargesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | Request for cancelling a charge |


#### Example Usage

```go
chargeId := "charge_id"
idempotencyKey := "idempotency-key"
var body *models_pkg.ChargesRequest

var result *models_pkg.ChargesResponse
result,_ = charges.CancelCharge(chargeId, idempotencyKey, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="confirm_payment"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.ConfirmPayment") ConfirmPayment

> ConfirmPayment


```go
func (me *CHARGES_IMPL) ConfirmPayment(
            chargeId string,
            idempotencyKey *string,
            body *models_pkg.CreateConfirmPaymentRequest)(*models_pkg.ChargesConfirmPaymentResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | TODO: Add a parameter description |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | Request for confirm payment |


#### Example Usage

```go
chargeId := "charge_id"
idempotencyKey := "idempotency-key"
var body *models_pkg.CreateConfirmPaymentRequest

var result *models_pkg.ChargesConfirmPaymentResponse
result,_ = charges.ConfirmPayment(chargeId, idempotencyKey, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_charge_card"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargeCard") UpdateChargeCard

> Updates the card from a charge


```go
func (me *CHARGES_IMPL) UpdateChargeCard(
            chargeId string,
            body *models_pkg.ChargesCardRequest,
            idempotencyKey *string)(*models_pkg.ChargesCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| body |  ``` Required ```  | Request for updating a charge's card |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var body *models_pkg.ChargesCardRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.ChargesCardResponse
result,_ = charges.UpdateChargeCard(chargeId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_charges"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetCharges") GetCharges

> Lists all charges


```go
func (me *CHARGES_IMPL) GetCharges(
            page *int64,
            size *int64,
            code *string,
            status *string,
            paymentMethod *string,
            customerId *string,
            orderId *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ChargesResponse2,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| code |  ``` Optional ```  | Filter for charge's code |
| status |  ``` Optional ```  | Filter for charge's status |
| paymentMethod |  ``` Optional ```  | Filter for charge's payment method |
| customerId |  ``` Optional ```  | Filter for charge's customer id |
| orderId |  ``` Optional ```  | Filter for charge's order id |
| createdSince |  ``` Optional ```  | Filter for the beginning of the range for charge's creation |
| createdUntil |  ``` Optional ```  | Filter for the end of the range for charge's creation |


#### Example Usage

```go
page,_ := strconv.ParseInt("15", 10, 8)
size,_ := strconv.ParseInt("15", 10, 8)
code := "code"
status := "status"
paymentMethod := "payment_method"
customerId := "customer_id"
orderId := "order_id"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ChargesResponse2
result,_ = charges.GetCharges(page, size, code, status, paymentMethod, customerId, orderId, createdSince, createdUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="retry_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.RetryCharge") RetryCharge

> Retries a charge


```go
func (me *CHARGES_IMPL) RetryCharge(
            chargeId string,
            idempotencyKey *string)(*models_pkg.ChargesRetryResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
idempotencyKey := "idempotency-key"

var result *models_pkg.ChargesRetryResponse
result,_ = charges.RetryCharge(chargeId, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_charge_payment_method"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargePaymentMethod") UpdateChargePaymentMethod

> Updates a charge's payment method


```go
func (me *CHARGES_IMPL) UpdateChargePaymentMethod(
            chargeId string,
            body *models_pkg.ChargesPaymentMethodRequest,
            idempotencyKey *string)(*models_pkg.ChargesPaymentMethodResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| body |  ``` Required ```  | Request for updating the payment method from a charge |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var body *models_pkg.ChargesPaymentMethodRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.ChargesPaymentMethodResponse
result,_ = charges.UpdateChargePaymentMethod(chargeId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_charge_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargeMetadata") UpdateChargeMetadata

> Updates the metadata from a charge


```go
func (me *CHARGES_IMPL) UpdateChargeMetadata(
            chargeId string,
            body *models_pkg.ChargesMetadataRequest,
            idempotencyKey *string)(*models_pkg.ChargesMetadataResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | The charge id |
| body |  ``` Required ```  | Request for updating the charge metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var body *models_pkg.ChargesMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.ChargesMetadataResponse
result,_ = charges.UpdateChargeMetadata(chargeId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="capture_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CaptureCharge") CaptureCharge

> Captures a charge


```go
func (me *CHARGES_IMPL) CaptureCharge(
            chargeId string,
            idempotencyKey *string,
            body *models_pkg.ChargesCaptureRequest)(*models_pkg.ChargesCaptureResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge id |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | Request for capturing a charge |


#### Example Usage

```go
chargeId := "charge_id"
idempotencyKey := "idempotency-key"
var body *models_pkg.ChargesCaptureRequest

var result *models_pkg.ChargesCaptureResponse
result,_ = charges.CaptureCharge(chargeId, idempotencyKey, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_charge_due_date"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.UpdateChargeDueDate") UpdateChargeDueDate

> Updates the due date from a charge


```go
func (me *CHARGES_IMPL) UpdateChargeDueDate(
            chargeId string,
            body *models_pkg.ChargesDueDateRequest,
            idempotencyKey *string)(*models_pkg.ChargesDueDateResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge Id |
| body |  ``` Required ```  | Request for updating the due date |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var body *models_pkg.ChargesDueDateRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.ChargesDueDateResponse
result,_ = charges.UpdateChargeDueDate(chargeId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_charge"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.CreateCharge") CreateCharge

> Creates a new charge


```go
func (me *CHARGES_IMPL) CreateCharge(
            body *models_pkg.ChargesRequest1,
            idempotencyKey *string)(*models_pkg.ChargesResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Request for creating a charge |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.ChargesRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.ChargesResponse
result,_ = charges.CreateCharge(body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_charge_transactions"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetChargeTransactions") GetChargeTransactions

> GetChargeTransactions


```go
func (me *CHARGES_IMPL) GetChargeTransactions(
            chargeId string,
            page *int64,
            size *int64)(*models_pkg.ChargesTransactionsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | Charge Id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
chargeId := "charge_id"
page,_ := strconv.ParseInt("15", 10, 8)
size,_ := strconv.ParseInt("15", 10, 8)

var result *models_pkg.ChargesTransactionsResponse
result,_ = charges.GetChargeTransactions(chargeId, page, size)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_charges_summary"></a>![Method: ](https://apidocs.io/img/method.png ".charges_pkg.GetChargesSummary") GetChargesSummary

> GetChargesSummary


```go
func (me *CHARGES_IMPL) GetChargesSummary(
            status string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.GetChargesSummaryResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| status |  ``` Required ```  | TODO: Add a parameter description |
| createdSince |  ``` Optional ```  | TODO: Add a parameter description |
| createdUntil |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
status := "status"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.GetChargesSummaryResponse
result,_ = charges.GetChargesSummary(status, createdSince, createdUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="recipients_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".recipients_pkg") recipients_pkg

### Get instance

Factory for the ``` RECIPIENTS ``` interface can be accessed from the package recipients_pkg.

```go
recipients := recipients_pkg.NewRECIPIENTS()
```

### <a name="update_recipient_metadata"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateRecipientMetadata") UpdateRecipientMetadata

> Updates recipient metadata


```go
func (me *RECIPIENTS_IMPL) UpdateRecipientMetadata(
            recipientId string,
            body *models_pkg.RecipientsMetadataRequest,
            idempotencyKey *string)(*models_pkg.RecipientsMetadataResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| body |  ``` Required ```  | Metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var body *models_pkg.RecipientsMetadataRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.RecipientsMetadataResponse
result,_ = recipients.UpdateRecipientMetadata(recipientId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_recipient_transfer_settings"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateRecipientTransferSettings") UpdateRecipientTransferSettings

> UpdateRecipientTransferSettings


```go
func (me *RECIPIENTS_IMPL) UpdateRecipientTransferSettings(
            recipientId string,
            body *models_pkg.UpdateTransferSettingsRequest,
            idempotencyKey *string)(*models_pkg.RecipientsTransferSettingsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient Identificator |
| body |  ``` Required ```  | TODO: Add a parameter description |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var body *models_pkg.UpdateTransferSettingsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.RecipientsTransferSettingsResponse
result,_ = recipients.UpdateRecipientTransferSettings(recipientId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_anticipation"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetAnticipation") GetAnticipation

> Gets an anticipation


```go
func (me *RECIPIENTS_IMPL) GetAnticipation(
            recipientId string,
            anticipationId string)(*models_pkg.RecipientsAnticipationsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| anticipationId |  ``` Required ```  | Anticipation id |


#### Example Usage

```go
recipientId := "recipient_id"
anticipationId := "anticipation_id"

var result *models_pkg.RecipientsAnticipationsResponse
result,_ = recipients.GetAnticipation(recipientId, anticipationId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_recipients"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetRecipients") GetRecipients

> Retrieves paginated recipients information


```go
func (me *RECIPIENTS_IMPL) GetRecipients(
            page *int64,
            size *int64)(*models_pkg.RecipientsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |


#### Example Usage

```go
page,_ := strconv.ParseInt("15", 10, 8)
size,_ := strconv.ParseInt("15", 10, 8)

var result *models_pkg.RecipientsResponse
result,_ = recipients.GetRecipients(page, size)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_recipient"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.CreateRecipient") CreateRecipient

> Creates a new recipient


```go
func (me *RECIPIENTS_IMPL) CreateRecipient(
            body *models_pkg.RecipientsRequest,
            idempotencyKey *string)(*models_pkg.RecipientsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | Recipient data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.RecipientsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.RecipientsResponse1
result,_ = recipients.CreateRecipient(body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_balance"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetBalance") GetBalance

> Get balance information for a recipient


```go
func (me *RECIPIENTS_IMPL) GetBalance(recipientId string)(*models_pkg.RecipientsBalanceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |


#### Example Usage

```go
recipientId := "recipient_id"

var result *models_pkg.RecipientsBalanceResponse
result,_ = recipients.GetBalance(recipientId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_anticipations"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetAnticipations") GetAnticipations

> Retrieves a paginated list of anticipations from a recipient


```go
func (me *RECIPIENTS_IMPL) GetAnticipations(
            recipientId string,
            page *int64,
            size *int64,
            status *string,
            timeframe *string,
            paymentDateSince *time.Time,
            paymentDateUntil *time.Time,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.RecipientsAnticipationsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| status |  ``` Optional ```  | Filter for anticipation status |
| timeframe |  ``` Optional ```  | Filter for anticipation timeframe |
| paymentDateSince |  ``` Optional ```  | Filter for start range for anticipation payment date |
| paymentDateUntil |  ``` Optional ```  | Filter for end range for anticipation payment date |
| createdSince |  ``` Optional ```  | Filter for start range for anticipation creation date |
| createdUntil |  ``` Optional ```  | Filter for end range for anticipation creation date |


#### Example Usage

```go
recipientId := "recipient_id"
page,_ := strconv.ParseInt("15", 10, 8)
size,_ := strconv.ParseInt("15", 10, 8)
status := "status"
timeframe := "timeframe"
paymentDateSince := time.Now()
paymentDateUntil := time.Now()
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.RecipientsAnticipationsResponse1
result,_ = recipients.GetAnticipations(recipientId, page, size, status, timeframe, paymentDateSince, paymentDateUntil, createdSince, createdUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_anticipation"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.CreateAnticipation") CreateAnticipation

> Creates an anticipation


```go
func (me *RECIPIENTS_IMPL) CreateAnticipation(
            recipientId string,
            body *models_pkg.RecipientsAnticipationsRequest,
            idempotencyKey *string)(*models_pkg.RecipientsAnticipationsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| body |  ``` Required ```  | Anticipation data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var body *models_pkg.RecipientsAnticipationsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.RecipientsAnticipationsResponse
result,_ = recipients.CreateAnticipation(recipientId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_recipient_default_bank_account"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateRecipientDefaultBankAccount") UpdateRecipientDefaultBankAccount

> Updates the default bank account from a recipient


```go
func (me *RECIPIENTS_IMPL) UpdateRecipientDefaultBankAccount(
            recipientId string,
            body *models_pkg.RecipientsDefaultBankAccountRequest,
            idempotencyKey *string)(*models_pkg.RecipientsDefaultBankAccountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| body |  ``` Required ```  | Bank account data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var body *models_pkg.RecipientsDefaultBankAccountRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.RecipientsDefaultBankAccountResponse
result,_ = recipients.UpdateRecipientDefaultBankAccount(recipientId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_recipient"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetRecipient") GetRecipient

> Retrieves recipient information


```go
func (me *RECIPIENTS_IMPL) GetRecipient(recipientId string)(*models_pkg.RecipientsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipiend id |


#### Example Usage

```go
recipientId := "recipient_id"

var result *models_pkg.RecipientsResponse1
result,_ = recipients.GetRecipient(recipientId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_recipient"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateRecipient") UpdateRecipient

> Updates a recipient


```go
func (me *RECIPIENTS_IMPL) UpdateRecipient(
            recipientId string,
            body *models_pkg.RecipientsRequest1,
            idempotencyKey *string)(*models_pkg.RecipientsResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| body |  ``` Required ```  | Recipient data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var body *models_pkg.RecipientsRequest1
idempotencyKey := "idempotency-key"

var result *models_pkg.RecipientsResponse1
result,_ = recipients.UpdateRecipient(recipientId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_transfer"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetTransfer") GetTransfer

> Gets a transfer


```go
func (me *RECIPIENTS_IMPL) GetTransfer(
            recipientId string,
            transferId string)(*models_pkg.RecipientsTransfersResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| transferId |  ``` Required ```  | Transfer id |


#### Example Usage

```go
recipientId := "recipient_id"
transferId := "transfer_id"

var result *models_pkg.RecipientsTransfersResponse
result,_ = recipients.GetTransfer(recipientId, transferId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_transfers"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetTransfers") GetTransfers

> Gets a paginated list of transfers for the recipient


```go
func (me *RECIPIENTS_IMPL) GetTransfers(
            recipientId string,
            page *int64,
            size *int64,
            status *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.RecipientsTransfersResponse1,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| page |  ``` Optional ```  | Page number |
| size |  ``` Optional ```  | Page size |
| status |  ``` Optional ```  | Filter for transfer status |
| createdSince |  ``` Optional ```  | Filter for start range of transfer creation date |
| createdUntil |  ``` Optional ```  | Filter for end range of transfer creation date |


#### Example Usage

```go
recipientId := "recipient_id"
page,_ := strconv.ParseInt("15", 10, 8)
size,_ := strconv.ParseInt("15", 10, 8)
status := "status"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.RecipientsTransfersResponse1
result,_ = recipients.GetTransfers(recipientId, page, size, status, createdSince, createdUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_transfer"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.CreateTransfer") CreateTransfer

> Creates a transfer for a recipient


```go
func (me *RECIPIENTS_IMPL) CreateTransfer(
            recipientId string,
            body *models_pkg.RecipientsTransfersRequest,
            idempotencyKey *string)(*models_pkg.RecipientsTransfersResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient Id |
| body |  ``` Required ```  | Transfer data |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var body *models_pkg.RecipientsTransfersRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.RecipientsTransfersResponse
result,_ = recipients.CreateTransfer(recipientId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_anticipation_limits"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetAnticipationLimits") GetAnticipationLimits

> Gets the anticipation limits for a recipient


```go
func (me *RECIPIENTS_IMPL) GetAnticipationLimits(
            recipientId string,
            timeframe string,
            paymentDate *time.Time)(*models_pkg.RecipientsAnticipationLimitsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| timeframe |  ``` Required ```  | Timeframe |
| paymentDate |  ``` Required ```  | Anticipation payment date |


#### Example Usage

```go
recipientId := "recipient_id"
timeframe := "timeframe"
paymentDate := time.Now()

var result *models_pkg.RecipientsAnticipationLimitsResponse
result,_ = recipients.GetAnticipationLimits(recipientId, timeframe, paymentDate)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="create_withdraw"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.CreateWithdraw") CreateWithdraw

> CreateWithdraw


```go
func (me *RECIPIENTS_IMPL) CreateWithdraw(
            recipientId string,
            body *models_pkg.CreateWithdrawRequest)(*models_pkg.GetWithdrawResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var body *models_pkg.CreateWithdrawRequest

var result *models_pkg.GetWithdrawResponse
result,_ = recipients.CreateWithdraw(recipientId, body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_withdrawals"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetWithdrawals") GetWithdrawals

> Gets a paginated list of transfers for the recipient


```go
func (me *RECIPIENTS_IMPL) GetWithdrawals(
            recipientId string,
            page *int64,
            size *int64,
            status *string,
            createdSince *time.Time,
            createdUntil *time.Time)(*models_pkg.ListWithdrawals,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | TODO: Add a parameter description |
| page |  ``` Optional ```  | TODO: Add a parameter description |
| size |  ``` Optional ```  | TODO: Add a parameter description |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| createdSince |  ``` Optional ```  | TODO: Add a parameter description |
| createdUntil |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
page,_ := strconv.ParseInt("15", 10, 8)
size,_ := strconv.ParseInt("15", 10, 8)
status := "status"
createdSince := time.Now()
createdUntil := time.Now()

var result *models_pkg.ListWithdrawals
result,_ = recipients.GetWithdrawals(recipientId, page, size, status, createdSince, createdUntil)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_withdraw_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetWithdrawById") GetWithdrawById

> GetWithdrawById


```go
func (me *RECIPIENTS_IMPL) GetWithdrawById(
            recipientId string,
            withdrawalId string)(*models_pkg.GetWithdrawResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | TODO: Add a parameter description |
| withdrawalId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
withdrawalId := "withdrawal_id"

var result *models_pkg.GetWithdrawResponse
result,_ = recipients.GetWithdrawById(recipientId, withdrawalId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="update_automatic_anticipation_settings"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.UpdateAutomaticAnticipationSettings") UpdateAutomaticAnticipationSettings

> Updates recipient metadata


```go
func (me *RECIPIENTS_IMPL) UpdateAutomaticAnticipationSettings(
            recipientId string,
            body *models_pkg.UpdateAutomaticAnticipationSettingsRequest,
            idempotencyKey *string)(*models_pkg.RecipientsAutomaticAnticipationSettingsResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recipientId |  ``` Required ```  | Recipient id |
| body |  ``` Required ```  | Metadata |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
recipientId := "recipient_id"
var body *models_pkg.UpdateAutomaticAnticipationSettingsRequest
idempotencyKey := "idempotency-key"

var result *models_pkg.RecipientsAutomaticAnticipationSettingsResponse
result,_ = recipients.UpdateAutomaticAnticipationSettings(recipientId, body, idempotencyKey)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_recipient_by_code"></a>![Method: ](https://apidocs.io/img/method.png ".recipients_pkg.GetRecipientByCode") GetRecipientByCode

> Retrieves recipient information


```go
func (me *RECIPIENTS_IMPL) GetRecipientByCode(code string)(*models_pkg.RecipientsCodeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| code |  ``` Required ```  | Recipient code |


#### Example Usage

```go
code := "code"

var result *models_pkg.RecipientsCodeResponse
result,_ = recipients.GetRecipientByCode(code)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="tokens_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".tokens_pkg") tokens_pkg

### Get instance

Factory for the ``` TOKENS ``` interface can be accessed from the package tokens_pkg.

```go
tokens := tokens_pkg.NewTOKENS()
```

### <a name="create_token"></a>![Method: ](https://apidocs.io/img/method.png ".tokens_pkg.CreateToken") CreateToken

> *Tags:*  ``` Skips Authentication ``` 

> CreateToken


```go
func (me *TOKENS_IMPL) CreateToken(
            publicKey string,
            body *models_pkg.TokensRequest,
            idempotencyKey *string,
            appId *string)(*models_pkg.TokensResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| publicKey |  ``` Required ```  | Public key |
| body |  ``` Required ```  | Request for creating a token |
| idempotencyKey |  ``` Optional ```  | TODO: Add a parameter description |
| appId |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
publicKey := "public_key"
var body *models_pkg.TokensRequest
idempotencyKey := "idempotency-key"
appId := "appId"

var result *models_pkg.TokensResponse
result,_ = tokens.CreateToken(publicKey, body, idempotencyKey, appId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_token"></a>![Method: ](https://apidocs.io/img/method.png ".tokens_pkg.GetToken") GetToken

> *Tags:*  ``` Skips Authentication ``` 

> Gets a token from its id


```go
func (me *TOKENS_IMPL) GetToken(
            id string,
            publicKey string,
            appId *string)(*models_pkg.TokensResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Token id |
| publicKey |  ``` Required ```  | Public key |
| appId |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
id := "id"
publicKey := "public_key"
appId := "appId"

var result *models_pkg.TokensResponse
result,_ = tokens.GetToken(id, publicKey, appId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="transactions_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".transactions_pkg") transactions_pkg

### Get instance

Factory for the ``` TRANSACTIONS ``` interface can be accessed from the package transactions_pkg.

```go
transactions := transactions_pkg.NewTRANSACTIONS()
```

### <a name="get_transaction"></a>![Method: ](https://apidocs.io/img/method.png ".transactions_pkg.GetTransaction") GetTransaction

> GetTransaction


```go
func (me *TRANSACTIONS_IMPL) GetTransaction(transactionId string)(*models_pkg.GetTransactionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transactionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
transactionId := "transaction_id"

var result *models_pkg.GetTransactionResponse
result,_ = transactions.GetTransaction(transactionId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)

## <a name="transfers_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".transfers_pkg") transfers_pkg

### Get instance

Factory for the ``` TRANSFERS ``` interface can be accessed from the package transfers_pkg.

```go
transfers := transfers_pkg.NewTRANSFERS()
```

### <a name="post_create_transfer"></a>![Method: ](https://apidocs.io/img/method.png ".transfers_pkg.PostCreateTransfer") PostCreateTransfer

> CreateTransfer


```go
func (me *TRANSFERS_IMPL) PostCreateTransfer(body *models_pkg.CreateTransfer)(*models_pkg.GetTransfer,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateTransfer

var result *models_pkg.GetTransfer
result,_ = transfers.PostCreateTransfer(body)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_transfer_by_id"></a>![Method: ](https://apidocs.io/img/method.png ".transfers_pkg.GetTransferById") GetTransferById

> GetTransferById


```go
func (me *TRANSFERS_IMPL) GetTransferById(transferId string)(*models_pkg.GetTransfer,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transferId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
transferId := "transfer_id"

var result *models_pkg.GetTransfer
result,_ = transfers.GetTransferById(transferId)

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



### <a name="get_transfers1"></a>![Method: ](https://apidocs.io/img/method.png ".transfers_pkg.GetTransfers1") GetTransfers1

> Gets all transfers


```go
func (me *TRANSFERS_IMPL) GetTransfers1()(*models_pkg.ListTransfers,error)
```

#### Example Usage

```go

var result *models_pkg.ListTransfers
result,_ = transfers.GetTransfers1()

```

#### Errors
 
| Error Code | Error Description |
|------------|-------------------|
| 400 | Invalid request |
| 401 | Invalid API key |
| 404 | An informed resource was not found |
| 412 | Business validation error |
| 422 | Contract validation error |
| 500 | Internal server error |



[Back to List of Controllers](#list_of_controllers)



