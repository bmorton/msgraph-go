# \DeviceManagementComplianceManagementPartnerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementCreateComplianceManagementPartners) | **Post** /deviceManagement/complianceManagementPartners | Create new navigation property to complianceManagementPartners for deviceManagement
[**DeviceManagementDeleteComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementDeleteComplianceManagementPartners) | **Delete** /deviceManagement/complianceManagementPartners/{complianceManagementPartner-id} | Delete navigation property complianceManagementPartners for deviceManagement
[**DeviceManagementGetComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementGetComplianceManagementPartners) | **Get** /deviceManagement/complianceManagementPartners/{complianceManagementPartner-id} | Get complianceManagementPartners from deviceManagement
[**DeviceManagementListComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementListComplianceManagementPartners) | **Get** /deviceManagement/complianceManagementPartners | Get complianceManagementPartners from deviceManagement
[**DeviceManagementUpdateComplianceManagementPartners**](DeviceManagementComplianceManagementPartnerApi.md#DeviceManagementUpdateComplianceManagementPartners) | **Patch** /deviceManagement/complianceManagementPartners/{complianceManagementPartner-id} | Update the navigation property complianceManagementPartners in deviceManagement



## DeviceManagementCreateComplianceManagementPartners

> MicrosoftGraphComplianceManagementPartner DeviceManagementCreateComplianceManagementPartners(ctx).MicrosoftGraphComplianceManagementPartner(microsoftGraphComplianceManagementPartner).Execute()

Create new navigation property to complianceManagementPartners for deviceManagement



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
    microsoftGraphComplianceManagementPartner := *openapiclient.NewMicrosoftGraphComplianceManagementPartner() // MicrosoftGraphComplianceManagementPartner | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementComplianceManagementPartnerApi.DeviceManagementCreateComplianceManagementPartners(context.Background()).MicrosoftGraphComplianceManagementPartner(microsoftGraphComplianceManagementPartner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementComplianceManagementPartnerApi.DeviceManagementCreateComplianceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateComplianceManagementPartners`: MicrosoftGraphComplianceManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementComplianceManagementPartnerApi.DeviceManagementCreateComplianceManagementPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateComplianceManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphComplianceManagementPartner** | [**MicrosoftGraphComplianceManagementPartner**](MicrosoftGraphComplianceManagementPartner.md) | New navigation property | 

### Return type

[**MicrosoftGraphComplianceManagementPartner**](MicrosoftGraphComplianceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteComplianceManagementPartners

> DeviceManagementDeleteComplianceManagementPartners(ctx, complianceManagementPartnerId).IfMatch(ifMatch).Execute()

Delete navigation property complianceManagementPartners for deviceManagement



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
    complianceManagementPartnerId := "complianceManagementPartnerId_example" // string | key: id of complianceManagementPartner
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementComplianceManagementPartnerApi.DeviceManagementDeleteComplianceManagementPartners(context.Background(), complianceManagementPartnerId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementComplianceManagementPartnerApi.DeviceManagementDeleteComplianceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complianceManagementPartnerId** | **string** | key: id of complianceManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteComplianceManagementPartnersRequest struct via the builder pattern


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


## DeviceManagementGetComplianceManagementPartners

> MicrosoftGraphComplianceManagementPartner DeviceManagementGetComplianceManagementPartners(ctx, complianceManagementPartnerId).Select_(select_).Expand(expand).Execute()

Get complianceManagementPartners from deviceManagement



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
    complianceManagementPartnerId := "complianceManagementPartnerId_example" // string | key: id of complianceManagementPartner
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementComplianceManagementPartnerApi.DeviceManagementGetComplianceManagementPartners(context.Background(), complianceManagementPartnerId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementComplianceManagementPartnerApi.DeviceManagementGetComplianceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetComplianceManagementPartners`: MicrosoftGraphComplianceManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementComplianceManagementPartnerApi.DeviceManagementGetComplianceManagementPartners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complianceManagementPartnerId** | **string** | key: id of complianceManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetComplianceManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphComplianceManagementPartner**](MicrosoftGraphComplianceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListComplianceManagementPartners

> CollectionOfComplianceManagementPartner DeviceManagementListComplianceManagementPartners(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get complianceManagementPartners from deviceManagement



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
    resp, r, err := api_client.DeviceManagementComplianceManagementPartnerApi.DeviceManagementListComplianceManagementPartners(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementComplianceManagementPartnerApi.DeviceManagementListComplianceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListComplianceManagementPartners`: CollectionOfComplianceManagementPartner
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementComplianceManagementPartnerApi.DeviceManagementListComplianceManagementPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListComplianceManagementPartnersRequest struct via the builder pattern


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

[**CollectionOfComplianceManagementPartner**](CollectionOfComplianceManagementPartner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateComplianceManagementPartners

> DeviceManagementUpdateComplianceManagementPartners(ctx, complianceManagementPartnerId).MicrosoftGraphComplianceManagementPartner(microsoftGraphComplianceManagementPartner).Execute()

Update the navigation property complianceManagementPartners in deviceManagement



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
    complianceManagementPartnerId := "complianceManagementPartnerId_example" // string | key: id of complianceManagementPartner
    microsoftGraphComplianceManagementPartner := *openapiclient.NewMicrosoftGraphComplianceManagementPartner() // MicrosoftGraphComplianceManagementPartner | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementComplianceManagementPartnerApi.DeviceManagementUpdateComplianceManagementPartners(context.Background(), complianceManagementPartnerId).MicrosoftGraphComplianceManagementPartner(microsoftGraphComplianceManagementPartner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementComplianceManagementPartnerApi.DeviceManagementUpdateComplianceManagementPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**complianceManagementPartnerId** | **string** | key: id of complianceManagementPartner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateComplianceManagementPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphComplianceManagementPartner** | [**MicrosoftGraphComplianceManagementPartner**](MicrosoftGraphComplianceManagementPartner.md) | New navigation property values | 

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

