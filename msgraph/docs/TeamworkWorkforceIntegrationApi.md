# \TeamworkWorkforceIntegrationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamworkCreateWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkCreateWorkforceIntegrations) | **Post** /teamwork/workforceIntegrations | Create new navigation property to workforceIntegrations for teamwork
[**TeamworkDeleteWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkDeleteWorkforceIntegrations) | **Delete** /teamwork/workforceIntegrations/{workforceIntegration-id} | Delete navigation property workforceIntegrations for teamwork
[**TeamworkGetWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkGetWorkforceIntegrations) | **Get** /teamwork/workforceIntegrations/{workforceIntegration-id} | Get workforceIntegrations from teamwork
[**TeamworkListWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkListWorkforceIntegrations) | **Get** /teamwork/workforceIntegrations | Get workforceIntegrations from teamwork
[**TeamworkUpdateWorkforceIntegrations**](TeamworkWorkforceIntegrationApi.md#TeamworkUpdateWorkforceIntegrations) | **Patch** /teamwork/workforceIntegrations/{workforceIntegration-id} | Update the navigation property workforceIntegrations in teamwork



## TeamworkCreateWorkforceIntegrations

> MicrosoftGraphWorkforceIntegration TeamworkCreateWorkforceIntegrations(ctx).MicrosoftGraphWorkforceIntegration(microsoftGraphWorkforceIntegration).Execute()

Create new navigation property to workforceIntegrations for teamwork

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
    microsoftGraphWorkforceIntegration := *openapiclient.NewMicrosoftGraphWorkforceIntegration() // MicrosoftGraphWorkforceIntegration | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamworkWorkforceIntegrationApi.TeamworkCreateWorkforceIntegrations(context.Background()).MicrosoftGraphWorkforceIntegration(microsoftGraphWorkforceIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamworkWorkforceIntegrationApi.TeamworkCreateWorkforceIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamworkCreateWorkforceIntegrations`: MicrosoftGraphWorkforceIntegration
    fmt.Fprintf(os.Stdout, "Response from `TeamworkWorkforceIntegrationApi.TeamworkCreateWorkforceIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamworkCreateWorkforceIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphWorkforceIntegration** | [**MicrosoftGraphWorkforceIntegration**](MicrosoftGraphWorkforceIntegration.md) | New navigation property | 

### Return type

[**MicrosoftGraphWorkforceIntegration**](MicrosoftGraphWorkforceIntegration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamworkDeleteWorkforceIntegrations

> TeamworkDeleteWorkforceIntegrations(ctx, workforceIntegrationId).IfMatch(ifMatch).Execute()

Delete navigation property workforceIntegrations for teamwork

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
    workforceIntegrationId := "workforceIntegrationId_example" // string | key: id of workforceIntegration
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamworkWorkforceIntegrationApi.TeamworkDeleteWorkforceIntegrations(context.Background(), workforceIntegrationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamworkWorkforceIntegrationApi.TeamworkDeleteWorkforceIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workforceIntegrationId** | **string** | key: id of workforceIntegration | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamworkDeleteWorkforceIntegrationsRequest struct via the builder pattern


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


## TeamworkGetWorkforceIntegrations

> MicrosoftGraphWorkforceIntegration TeamworkGetWorkforceIntegrations(ctx, workforceIntegrationId).Select_(select_).Expand(expand).Execute()

Get workforceIntegrations from teamwork

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
    workforceIntegrationId := "workforceIntegrationId_example" // string | key: id of workforceIntegration
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamworkWorkforceIntegrationApi.TeamworkGetWorkforceIntegrations(context.Background(), workforceIntegrationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamworkWorkforceIntegrationApi.TeamworkGetWorkforceIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamworkGetWorkforceIntegrations`: MicrosoftGraphWorkforceIntegration
    fmt.Fprintf(os.Stdout, "Response from `TeamworkWorkforceIntegrationApi.TeamworkGetWorkforceIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workforceIntegrationId** | **string** | key: id of workforceIntegration | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamworkGetWorkforceIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphWorkforceIntegration**](MicrosoftGraphWorkforceIntegration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamworkListWorkforceIntegrations

> CollectionOfWorkforceIntegration TeamworkListWorkforceIntegrations(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get workforceIntegrations from teamwork

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
    resp, r, err := api_client.TeamworkWorkforceIntegrationApi.TeamworkListWorkforceIntegrations(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamworkWorkforceIntegrationApi.TeamworkListWorkforceIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamworkListWorkforceIntegrations`: CollectionOfWorkforceIntegration
    fmt.Fprintf(os.Stdout, "Response from `TeamworkWorkforceIntegrationApi.TeamworkListWorkforceIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamworkListWorkforceIntegrationsRequest struct via the builder pattern


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

[**CollectionOfWorkforceIntegration**](CollectionOfWorkforceIntegration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamworkUpdateWorkforceIntegrations

> TeamworkUpdateWorkforceIntegrations(ctx, workforceIntegrationId).MicrosoftGraphWorkforceIntegration(microsoftGraphWorkforceIntegration).Execute()

Update the navigation property workforceIntegrations in teamwork

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
    workforceIntegrationId := "workforceIntegrationId_example" // string | key: id of workforceIntegration
    microsoftGraphWorkforceIntegration := *openapiclient.NewMicrosoftGraphWorkforceIntegration() // MicrosoftGraphWorkforceIntegration | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamworkWorkforceIntegrationApi.TeamworkUpdateWorkforceIntegrations(context.Background(), workforceIntegrationId).MicrosoftGraphWorkforceIntegration(microsoftGraphWorkforceIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamworkWorkforceIntegrationApi.TeamworkUpdateWorkforceIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workforceIntegrationId** | **string** | key: id of workforceIntegration | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamworkUpdateWorkforceIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphWorkforceIntegration** | [**MicrosoftGraphWorkforceIntegration**](MicrosoftGraphWorkforceIntegration.md) | New navigation property values | 

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

