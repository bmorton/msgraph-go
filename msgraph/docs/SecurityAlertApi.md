# \SecurityAlertApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityCreateAlerts**](SecurityAlertApi.md#SecurityCreateAlerts) | **Post** /security/alerts | Create new navigation property to alerts for security
[**SecurityDeleteAlerts**](SecurityAlertApi.md#SecurityDeleteAlerts) | **Delete** /security/alerts/{alert-id} | Delete navigation property alerts for security
[**SecurityGetAlerts**](SecurityAlertApi.md#SecurityGetAlerts) | **Get** /security/alerts/{alert-id} | Get alerts from security
[**SecurityListAlerts**](SecurityAlertApi.md#SecurityListAlerts) | **Get** /security/alerts | Get alerts from security
[**SecurityUpdateAlerts**](SecurityAlertApi.md#SecurityUpdateAlerts) | **Patch** /security/alerts/{alert-id} | Update the navigation property alerts in security



## SecurityCreateAlerts

> MicrosoftGraphAlert SecurityCreateAlerts(ctx).MicrosoftGraphAlert(microsoftGraphAlert).Execute()

Create new navigation property to alerts for security



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    microsoftGraphAlert := *openapiclient.NewMicrosoftGraphAlert() // MicrosoftGraphAlert | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityAlertApi.SecurityCreateAlerts(context.Background()).MicrosoftGraphAlert(microsoftGraphAlert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityAlertApi.SecurityCreateAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityCreateAlerts`: MicrosoftGraphAlert
    fmt.Fprintf(os.Stdout, "Response from `SecurityAlertApi.SecurityCreateAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityCreateAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAlert** | [**MicrosoftGraphAlert**](MicrosoftGraphAlert.md) | New navigation property | 

### Return type

[**MicrosoftGraphAlert**](MicrosoftGraphAlert.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityDeleteAlerts

> SecurityDeleteAlerts(ctx, alertId).IfMatch(ifMatch).Execute()

Delete navigation property alerts for security



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    alertId := "alertId_example" // string | key: id of alert
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityAlertApi.SecurityDeleteAlerts(context.Background(), alertId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityAlertApi.SecurityDeleteAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | key: id of alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityDeleteAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityGetAlerts

> MicrosoftGraphAlert SecurityGetAlerts(ctx, alertId).Select_(select_).Expand(expand).Execute()

Get alerts from security



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    alertId := "alertId_example" // string | key: id of alert
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityAlertApi.SecurityGetAlerts(context.Background(), alertId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityAlertApi.SecurityGetAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityGetAlerts`: MicrosoftGraphAlert
    fmt.Fprintf(os.Stdout, "Response from `SecurityAlertApi.SecurityGetAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | key: id of alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityGetAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAlert**](MicrosoftGraphAlert.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityListAlerts

> CollectionOfAlert SecurityListAlerts(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get alerts from security



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityAlertApi.SecurityListAlerts(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityAlertApi.SecurityListAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityListAlerts`: CollectionOfAlert
    fmt.Fprintf(os.Stdout, "Response from `SecurityAlertApi.SecurityListAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityListAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfAlert**](CollectionOfAlert.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityUpdateAlerts

> SecurityUpdateAlerts(ctx, alertId).MicrosoftGraphAlert(microsoftGraphAlert).Execute()

Update the navigation property alerts in security



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    alertId := "alertId_example" // string | key: id of alert
    microsoftGraphAlert := *openapiclient.NewMicrosoftGraphAlert() // MicrosoftGraphAlert | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityAlertApi.SecurityUpdateAlerts(context.Background(), alertId).MicrosoftGraphAlert(microsoftGraphAlert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityAlertApi.SecurityUpdateAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | key: id of alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityUpdateAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAlert** | [**MicrosoftGraphAlert**](MicrosoftGraphAlert.md) | New navigation property values | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

