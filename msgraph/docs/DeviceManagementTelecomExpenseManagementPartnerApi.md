# \DeviceManagementTelecomExpenseManagementPartnerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementCreateTelecomExpenseManagementPartners) | **Post** /deviceManagement/telecomExpenseManagementPartners | Create new navigation property to telecomExpenseManagementPartners for deviceManagement
[**DeviceManagementDeleteTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementDeleteTelecomExpenseManagementPartners) | **Delete** /deviceManagement/telecomExpenseManagementPartners/{telecomExpenseManagementPartner-id} | Delete navigation property telecomExpenseManagementPartners for deviceManagement
[**DeviceManagementGetTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementGetTelecomExpenseManagementPartners) | **Get** /deviceManagement/telecomExpenseManagementPartners/{telecomExpenseManagementPartner-id} | Get telecomExpenseManagementPartners from deviceManagement
[**DeviceManagementListTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementListTelecomExpenseManagementPartners) | **Get** /deviceManagement/telecomExpenseManagementPartners | Get telecomExpenseManagementPartners from deviceManagement
[**DeviceManagementUpdateTelecomExpenseManagementPartners**](DeviceManagementTelecomExpenseManagementPartnerApi.md#DeviceManagementUpdateTelecomExpenseManagementPartners) | **Patch** /deviceManagement/telecomExpenseManagementPartners/{telecomExpenseManagementPartner-id} | Update the navigation property telecomExpenseManagementPartners in deviceManagement



## DeviceManagementCreateTelecomExpenseManagementPartners

> MicrosoftGraphTelecomExpenseManagementPartner DeviceManagementCreateTelecomExpenseManagementPartners(ctx).MicrosoftGraphTelecomExpenseManagementPartner(microsoftGraphTelecomExpenseManagementPartner).Execute()

Create new navigation property to telecomExpenseManagementPartners for deviceManagement



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
    microsoftGraphTelecomExpenseManagementPartner := *openapiclient.NewMicrosoftGraphTelecomExpenseManagementPartner() // MicrosoftGraphTelecomExpenseManagementPartner | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementCreateTelecomExpenseManagementPartners(context.Background()).MicrosoftGraphTelecomExpenseManagementPartner(microsoftGraphTelecomExpenseManagementPartner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementCreateTelecomExpenseManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateTelecomExpenseManagementPartners`: MicrosoftGraphTelecomExpenseManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementCreateTelecomExpenseManagementPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateTelecomExpenseManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTelecomExpenseManagementPartner** | [**MicrosoftGraphTelecomExpenseManagementPartner**](MicrosoftGraphTelecomExpenseManagementPartner.md) | New navigation property | 

### Return type

[**MicrosoftGraphTelecomExpenseManagementPartner**](MicrosoftGraphTelecomExpenseManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteTelecomExpenseManagementPartners

> DeviceManagementDeleteTelecomExpenseManagementPartners(ctx, telecomExpenseManagementPartnerId).IfMatch(ifMatch).Execute()

Delete navigation property telecomExpenseManagementPartners for deviceManagement



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
    telecomExpenseManagementPartnerId := "telecomExpenseManagementPartnerId_example" // string | key: id of telecomExpenseManagementPartner
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementDeleteTelecomExpenseManagementPartners(context.Background(), telecomExpenseManagementPartnerId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementDeleteTelecomExpenseManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**telecomExpenseManagementPartnerId** | **string** | key: id of telecomExpenseManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteTelecomExpenseManagementPartnersRequest struct via the builder pattern


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


## DeviceManagementGetTelecomExpenseManagementPartners

> MicrosoftGraphTelecomExpenseManagementPartner DeviceManagementGetTelecomExpenseManagementPartners(ctx, telecomExpenseManagementPartnerId).Select_(select_).Expand(expand).Execute()

Get telecomExpenseManagementPartners from deviceManagement



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
    telecomExpenseManagementPartnerId := "telecomExpenseManagementPartnerId_example" // string | key: id of telecomExpenseManagementPartner
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementGetTelecomExpenseManagementPartners(context.Background(), telecomExpenseManagementPartnerId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementGetTelecomExpenseManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetTelecomExpenseManagementPartners`: MicrosoftGraphTelecomExpenseManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementGetTelecomExpenseManagementPartners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**telecomExpenseManagementPartnerId** | **string** | key: id of telecomExpenseManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetTelecomExpenseManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTelecomExpenseManagementPartner**](MicrosoftGraphTelecomExpenseManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListTelecomExpenseManagementPartners

> CollectionOfTelecomExpenseManagementPartner DeviceManagementListTelecomExpenseManagementPartners(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get telecomExpenseManagementPartners from deviceManagement



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
    resp, r, err := api_client.DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementListTelecomExpenseManagementPartners(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementListTelecomExpenseManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListTelecomExpenseManagementPartners`: CollectionOfTelecomExpenseManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementListTelecomExpenseManagementPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListTelecomExpenseManagementPartnersRequest struct via the builder pattern


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

[**CollectionOfTelecomExpenseManagementPartner**](CollectionOfTelecomExpenseManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateTelecomExpenseManagementPartners

> DeviceManagementUpdateTelecomExpenseManagementPartners(ctx, telecomExpenseManagementPartnerId).MicrosoftGraphTelecomExpenseManagementPartner(microsoftGraphTelecomExpenseManagementPartner).Execute()

Update the navigation property telecomExpenseManagementPartners in deviceManagement



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
    telecomExpenseManagementPartnerId := "telecomExpenseManagementPartnerId_example" // string | key: id of telecomExpenseManagementPartner
    microsoftGraphTelecomExpenseManagementPartner := *openapiclient.NewMicrosoftGraphTelecomExpenseManagementPartner() // MicrosoftGraphTelecomExpenseManagementPartner | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementUpdateTelecomExpenseManagementPartners(context.Background(), telecomExpenseManagementPartnerId).MicrosoftGraphTelecomExpenseManagementPartner(microsoftGraphTelecomExpenseManagementPartner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTelecomExpenseManagementPartnerApi.DeviceManagementUpdateTelecomExpenseManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**telecomExpenseManagementPartnerId** | **string** | key: id of telecomExpenseManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateTelecomExpenseManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTelecomExpenseManagementPartner** | [**MicrosoftGraphTelecomExpenseManagementPartner**](MicrosoftGraphTelecomExpenseManagementPartner.md) | New navigation property values | 

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

