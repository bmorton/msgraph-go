# \DeviceManagementRemoteAssistancePartnerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementCreateRemoteAssistancePartners) | **Post** /deviceManagement/remoteAssistancePartners | Create new navigation property to remoteAssistancePartners for deviceManagement
[**DeviceManagementDeleteRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementDeleteRemoteAssistancePartners) | **Delete** /deviceManagement/remoteAssistancePartners/{remoteAssistancePartner-id} | Delete navigation property remoteAssistancePartners for deviceManagement
[**DeviceManagementGetRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementGetRemoteAssistancePartners) | **Get** /deviceManagement/remoteAssistancePartners/{remoteAssistancePartner-id} | Get remoteAssistancePartners from deviceManagement
[**DeviceManagementListRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementListRemoteAssistancePartners) | **Get** /deviceManagement/remoteAssistancePartners | Get remoteAssistancePartners from deviceManagement
[**DeviceManagementUpdateRemoteAssistancePartners**](DeviceManagementRemoteAssistancePartnerApi.md#DeviceManagementUpdateRemoteAssistancePartners) | **Patch** /deviceManagement/remoteAssistancePartners/{remoteAssistancePartner-id} | Update the navigation property remoteAssistancePartners in deviceManagement



## DeviceManagementCreateRemoteAssistancePartners

> MicrosoftGraphRemoteAssistancePartner DeviceManagementCreateRemoteAssistancePartners(ctx).MicrosoftGraphRemoteAssistancePartner(microsoftGraphRemoteAssistancePartner).Execute()

Create new navigation property to remoteAssistancePartners for deviceManagement



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
    microsoftGraphRemoteAssistancePartner := *openapiclient.NewMicrosoftGraphRemoteAssistancePartner() // MicrosoftGraphRemoteAssistancePartner | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRemoteAssistancePartnerApi.DeviceManagementCreateRemoteAssistancePartners(context.Background()).MicrosoftGraphRemoteAssistancePartner(microsoftGraphRemoteAssistancePartner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRemoteAssistancePartnerApi.DeviceManagementCreateRemoteAssistancePartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateRemoteAssistancePartners`: MicrosoftGraphRemoteAssistancePartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRemoteAssistancePartnerApi.DeviceManagementCreateRemoteAssistancePartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateRemoteAssistancePartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphRemoteAssistancePartner** | [**MicrosoftGraphRemoteAssistancePartner**](MicrosoftGraphRemoteAssistancePartner.md) | New navigation property | 

### Return type

[**MicrosoftGraphRemoteAssistancePartner**](MicrosoftGraphRemoteAssistancePartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteRemoteAssistancePartners

> DeviceManagementDeleteRemoteAssistancePartners(ctx, remoteAssistancePartnerId).IfMatch(ifMatch).Execute()

Delete navigation property remoteAssistancePartners for deviceManagement



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
    remoteAssistancePartnerId := "remoteAssistancePartnerId_example" // string | key: id of remoteAssistancePartner
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRemoteAssistancePartnerApi.DeviceManagementDeleteRemoteAssistancePartners(context.Background(), remoteAssistancePartnerId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRemoteAssistancePartnerApi.DeviceManagementDeleteRemoteAssistancePartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remoteAssistancePartnerId** | **string** | key: id of remoteAssistancePartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteRemoteAssistancePartnersRequest struct via the builder pattern


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


## DeviceManagementGetRemoteAssistancePartners

> MicrosoftGraphRemoteAssistancePartner DeviceManagementGetRemoteAssistancePartners(ctx, remoteAssistancePartnerId).Select_(select_).Expand(expand).Execute()

Get remoteAssistancePartners from deviceManagement



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
    remoteAssistancePartnerId := "remoteAssistancePartnerId_example" // string | key: id of remoteAssistancePartner
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRemoteAssistancePartnerApi.DeviceManagementGetRemoteAssistancePartners(context.Background(), remoteAssistancePartnerId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRemoteAssistancePartnerApi.DeviceManagementGetRemoteAssistancePartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetRemoteAssistancePartners`: MicrosoftGraphRemoteAssistancePartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRemoteAssistancePartnerApi.DeviceManagementGetRemoteAssistancePartners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remoteAssistancePartnerId** | **string** | key: id of remoteAssistancePartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetRemoteAssistancePartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphRemoteAssistancePartner**](MicrosoftGraphRemoteAssistancePartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListRemoteAssistancePartners

> CollectionOfRemoteAssistancePartner DeviceManagementListRemoteAssistancePartners(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get remoteAssistancePartners from deviceManagement



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
    resp, r, err := api_client.DeviceManagementRemoteAssistancePartnerApi.DeviceManagementListRemoteAssistancePartners(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRemoteAssistancePartnerApi.DeviceManagementListRemoteAssistancePartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListRemoteAssistancePartners`: CollectionOfRemoteAssistancePartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRemoteAssistancePartnerApi.DeviceManagementListRemoteAssistancePartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListRemoteAssistancePartnersRequest struct via the builder pattern


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

[**CollectionOfRemoteAssistancePartner**](CollectionOfRemoteAssistancePartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateRemoteAssistancePartners

> DeviceManagementUpdateRemoteAssistancePartners(ctx, remoteAssistancePartnerId).MicrosoftGraphRemoteAssistancePartner(microsoftGraphRemoteAssistancePartner).Execute()

Update the navigation property remoteAssistancePartners in deviceManagement



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
    remoteAssistancePartnerId := "remoteAssistancePartnerId_example" // string | key: id of remoteAssistancePartner
    microsoftGraphRemoteAssistancePartner := *openapiclient.NewMicrosoftGraphRemoteAssistancePartner() // MicrosoftGraphRemoteAssistancePartner | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRemoteAssistancePartnerApi.DeviceManagementUpdateRemoteAssistancePartners(context.Background(), remoteAssistancePartnerId).MicrosoftGraphRemoteAssistancePartner(microsoftGraphRemoteAssistancePartner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRemoteAssistancePartnerApi.DeviceManagementUpdateRemoteAssistancePartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remoteAssistancePartnerId** | **string** | key: id of remoteAssistancePartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateRemoteAssistancePartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphRemoteAssistancePartner** | [**MicrosoftGraphRemoteAssistancePartner**](MicrosoftGraphRemoteAssistancePartner.md) | New navigation property values | 

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

