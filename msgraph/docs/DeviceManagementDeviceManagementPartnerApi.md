# \DeviceManagementDeviceManagementPartnerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementCreateDeviceManagementPartners) | **Post** /deviceManagement/deviceManagementPartners | Create new navigation property to deviceManagementPartners for deviceManagement
[**DeviceManagementDeleteDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementDeleteDeviceManagementPartners) | **Delete** /deviceManagement/deviceManagementPartners/{deviceManagementPartner-id} | Delete navigation property deviceManagementPartners for deviceManagement
[**DeviceManagementGetDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementGetDeviceManagementPartners) | **Get** /deviceManagement/deviceManagementPartners/{deviceManagementPartner-id} | Get deviceManagementPartners from deviceManagement
[**DeviceManagementListDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementListDeviceManagementPartners) | **Get** /deviceManagement/deviceManagementPartners | Get deviceManagementPartners from deviceManagement
[**DeviceManagementUpdateDeviceManagementPartners**](DeviceManagementDeviceManagementPartnerApi.md#DeviceManagementUpdateDeviceManagementPartners) | **Patch** /deviceManagement/deviceManagementPartners/{deviceManagementPartner-id} | Update the navigation property deviceManagementPartners in deviceManagement



## DeviceManagementCreateDeviceManagementPartners

> MicrosoftGraphDeviceManagementPartner DeviceManagementCreateDeviceManagementPartners(ctx).MicrosoftGraphDeviceManagementPartner(microsoftGraphDeviceManagementPartner).Execute()

Create new navigation property to deviceManagementPartners for deviceManagement



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
    microsoftGraphDeviceManagementPartner := *openapiclient.NewMicrosoftGraphDeviceManagementPartner() // MicrosoftGraphDeviceManagementPartner | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementPartnerApi.DeviceManagementCreateDeviceManagementPartners(context.Background()).MicrosoftGraphDeviceManagementPartner(microsoftGraphDeviceManagementPartner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementPartnerApi.DeviceManagementCreateDeviceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateDeviceManagementPartners`: MicrosoftGraphDeviceManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementPartnerApi.DeviceManagementCreateDeviceManagementPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateDeviceManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceManagementPartner** | [**MicrosoftGraphDeviceManagementPartner**](MicrosoftGraphDeviceManagementPartner.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceManagementPartner**](MicrosoftGraphDeviceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteDeviceManagementPartners

> DeviceManagementDeleteDeviceManagementPartners(ctx, deviceManagementPartnerId).IfMatch(ifMatch).Execute()

Delete navigation property deviceManagementPartners for deviceManagement



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
    deviceManagementPartnerId := "deviceManagementPartnerId_example" // string | key: id of deviceManagementPartner
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementPartnerApi.DeviceManagementDeleteDeviceManagementPartners(context.Background(), deviceManagementPartnerId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementPartnerApi.DeviceManagementDeleteDeviceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementPartnerId** | **string** | key: id of deviceManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteDeviceManagementPartnersRequest struct via the builder pattern


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


## DeviceManagementGetDeviceManagementPartners

> MicrosoftGraphDeviceManagementPartner DeviceManagementGetDeviceManagementPartners(ctx, deviceManagementPartnerId).Select_(select_).Expand(expand).Execute()

Get deviceManagementPartners from deviceManagement



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
    deviceManagementPartnerId := "deviceManagementPartnerId_example" // string | key: id of deviceManagementPartner
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementPartnerApi.DeviceManagementGetDeviceManagementPartners(context.Background(), deviceManagementPartnerId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementPartnerApi.DeviceManagementGetDeviceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetDeviceManagementPartners`: MicrosoftGraphDeviceManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementPartnerApi.DeviceManagementGetDeviceManagementPartners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementPartnerId** | **string** | key: id of deviceManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetDeviceManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagementPartner**](MicrosoftGraphDeviceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceManagementPartners

> CollectionOfDeviceManagementPartner DeviceManagementListDeviceManagementPartners(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceManagementPartners from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementPartnerApi.DeviceManagementListDeviceManagementPartners(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementPartnerApi.DeviceManagementListDeviceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListDeviceManagementPartners`: CollectionOfDeviceManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementPartnerApi.DeviceManagementListDeviceManagementPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListDeviceManagementPartnersRequest struct via the builder pattern


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

[**CollectionOfDeviceManagementPartner**](CollectionOfDeviceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceManagementPartners

> DeviceManagementUpdateDeviceManagementPartners(ctx, deviceManagementPartnerId).MicrosoftGraphDeviceManagementPartner(microsoftGraphDeviceManagementPartner).Execute()

Update the navigation property deviceManagementPartners in deviceManagement



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
    deviceManagementPartnerId := "deviceManagementPartnerId_example" // string | key: id of deviceManagementPartner
    microsoftGraphDeviceManagementPartner := *openapiclient.NewMicrosoftGraphDeviceManagementPartner() // MicrosoftGraphDeviceManagementPartner | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementPartnerApi.DeviceManagementUpdateDeviceManagementPartners(context.Background(), deviceManagementPartnerId).MicrosoftGraphDeviceManagementPartner(microsoftGraphDeviceManagementPartner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementPartnerApi.DeviceManagementUpdateDeviceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementPartnerId** | **string** | key: id of deviceManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateDeviceManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceManagementPartner** | [**MicrosoftGraphDeviceManagementPartner**](MicrosoftGraphDeviceManagementPartner.md) | New navigation property values | 

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

