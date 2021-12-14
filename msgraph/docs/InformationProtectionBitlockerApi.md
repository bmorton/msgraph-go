# \InformationProtectionBitlockerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InformationProtectionBitlockerCreateRecoveryKeys**](InformationProtectionBitlockerApi.md#InformationProtectionBitlockerCreateRecoveryKeys) | **Post** /informationProtection/bitlocker/recoveryKeys | Create new navigation property to recoveryKeys for informationProtection
[**InformationProtectionBitlockerDeleteRecoveryKeys**](InformationProtectionBitlockerApi.md#InformationProtectionBitlockerDeleteRecoveryKeys) | **Delete** /informationProtection/bitlocker/recoveryKeys/{bitlockerRecoveryKey-id} | Delete navigation property recoveryKeys for informationProtection
[**InformationProtectionBitlockerGetRecoveryKeys**](InformationProtectionBitlockerApi.md#InformationProtectionBitlockerGetRecoveryKeys) | **Get** /informationProtection/bitlocker/recoveryKeys/{bitlockerRecoveryKey-id} | Get recoveryKeys from informationProtection
[**InformationProtectionBitlockerListRecoveryKeys**](InformationProtectionBitlockerApi.md#InformationProtectionBitlockerListRecoveryKeys) | **Get** /informationProtection/bitlocker/recoveryKeys | Get recoveryKeys from informationProtection
[**InformationProtectionBitlockerUpdateRecoveryKeys**](InformationProtectionBitlockerApi.md#InformationProtectionBitlockerUpdateRecoveryKeys) | **Patch** /informationProtection/bitlocker/recoveryKeys/{bitlockerRecoveryKey-id} | Update the navigation property recoveryKeys in informationProtection
[**InformationProtectionDeleteBitlocker**](InformationProtectionBitlockerApi.md#InformationProtectionDeleteBitlocker) | **Delete** /informationProtection/bitlocker | Delete navigation property bitlocker for informationProtection
[**InformationProtectionGetBitlocker**](InformationProtectionBitlockerApi.md#InformationProtectionGetBitlocker) | **Get** /informationProtection/bitlocker | Get bitlocker from informationProtection
[**InformationProtectionUpdateBitlocker**](InformationProtectionBitlockerApi.md#InformationProtectionUpdateBitlocker) | **Patch** /informationProtection/bitlocker | Update the navigation property bitlocker in informationProtection



## InformationProtectionBitlockerCreateRecoveryKeys

> MicrosoftGraphBitlockerRecoveryKey InformationProtectionBitlockerCreateRecoveryKeys(ctx).MicrosoftGraphBitlockerRecoveryKey(microsoftGraphBitlockerRecoveryKey).Execute()

Create new navigation property to recoveryKeys for informationProtection



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
    microsoftGraphBitlockerRecoveryKey := *openapiclient.NewMicrosoftGraphBitlockerRecoveryKey() // MicrosoftGraphBitlockerRecoveryKey | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionBitlockerApi.InformationProtectionBitlockerCreateRecoveryKeys(context.Background()).MicrosoftGraphBitlockerRecoveryKey(microsoftGraphBitlockerRecoveryKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionBitlockerApi.InformationProtectionBitlockerCreateRecoveryKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionBitlockerCreateRecoveryKeys`: MicrosoftGraphBitlockerRecoveryKey
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionBitlockerApi.InformationProtectionBitlockerCreateRecoveryKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionBitlockerCreateRecoveryKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphBitlockerRecoveryKey** | [**MicrosoftGraphBitlockerRecoveryKey**](MicrosoftGraphBitlockerRecoveryKey.md) | New navigation property | 

### Return type

[**MicrosoftGraphBitlockerRecoveryKey**](MicrosoftGraphBitlockerRecoveryKey.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionBitlockerDeleteRecoveryKeys

> InformationProtectionBitlockerDeleteRecoveryKeys(ctx, bitlockerRecoveryKeyId).IfMatch(ifMatch).Execute()

Delete navigation property recoveryKeys for informationProtection



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
    bitlockerRecoveryKeyId := "bitlockerRecoveryKeyId_example" // string | key: id of bitlockerRecoveryKey
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionBitlockerApi.InformationProtectionBitlockerDeleteRecoveryKeys(context.Background(), bitlockerRecoveryKeyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionBitlockerApi.InformationProtectionBitlockerDeleteRecoveryKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlockerRecoveryKeyId** | **string** | key: id of bitlockerRecoveryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionBitlockerDeleteRecoveryKeysRequest struct via the builder pattern


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


## InformationProtectionBitlockerGetRecoveryKeys

> MicrosoftGraphBitlockerRecoveryKey InformationProtectionBitlockerGetRecoveryKeys(ctx, bitlockerRecoveryKeyId).Select_(select_).Expand(expand).Execute()

Get recoveryKeys from informationProtection



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
    bitlockerRecoveryKeyId := "bitlockerRecoveryKeyId_example" // string | key: id of bitlockerRecoveryKey
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionBitlockerApi.InformationProtectionBitlockerGetRecoveryKeys(context.Background(), bitlockerRecoveryKeyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionBitlockerApi.InformationProtectionBitlockerGetRecoveryKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionBitlockerGetRecoveryKeys`: MicrosoftGraphBitlockerRecoveryKey
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionBitlockerApi.InformationProtectionBitlockerGetRecoveryKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlockerRecoveryKeyId** | **string** | key: id of bitlockerRecoveryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionBitlockerGetRecoveryKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphBitlockerRecoveryKey**](MicrosoftGraphBitlockerRecoveryKey.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionBitlockerListRecoveryKeys

> CollectionOfBitlockerRecoveryKey InformationProtectionBitlockerListRecoveryKeys(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get recoveryKeys from informationProtection



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
    resp, r, err := api_client.InformationProtectionBitlockerApi.InformationProtectionBitlockerListRecoveryKeys(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionBitlockerApi.InformationProtectionBitlockerListRecoveryKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionBitlockerListRecoveryKeys`: CollectionOfBitlockerRecoveryKey
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionBitlockerApi.InformationProtectionBitlockerListRecoveryKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionBitlockerListRecoveryKeysRequest struct via the builder pattern


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

[**CollectionOfBitlockerRecoveryKey**](CollectionOfBitlockerRecoveryKey.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionBitlockerUpdateRecoveryKeys

> InformationProtectionBitlockerUpdateRecoveryKeys(ctx, bitlockerRecoveryKeyId).MicrosoftGraphBitlockerRecoveryKey(microsoftGraphBitlockerRecoveryKey).Execute()

Update the navigation property recoveryKeys in informationProtection



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
    bitlockerRecoveryKeyId := "bitlockerRecoveryKeyId_example" // string | key: id of bitlockerRecoveryKey
    microsoftGraphBitlockerRecoveryKey := *openapiclient.NewMicrosoftGraphBitlockerRecoveryKey() // MicrosoftGraphBitlockerRecoveryKey | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionBitlockerApi.InformationProtectionBitlockerUpdateRecoveryKeys(context.Background(), bitlockerRecoveryKeyId).MicrosoftGraphBitlockerRecoveryKey(microsoftGraphBitlockerRecoveryKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionBitlockerApi.InformationProtectionBitlockerUpdateRecoveryKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bitlockerRecoveryKeyId** | **string** | key: id of bitlockerRecoveryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionBitlockerUpdateRecoveryKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphBitlockerRecoveryKey** | [**MicrosoftGraphBitlockerRecoveryKey**](MicrosoftGraphBitlockerRecoveryKey.md) | New navigation property values | 

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


## InformationProtectionDeleteBitlocker

> InformationProtectionDeleteBitlocker(ctx).IfMatch(ifMatch).Execute()

Delete navigation property bitlocker for informationProtection

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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionBitlockerApi.InformationProtectionDeleteBitlocker(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionBitlockerApi.InformationProtectionDeleteBitlocker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionDeleteBitlockerRequest struct via the builder pattern


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


## InformationProtectionGetBitlocker

> MicrosoftGraphBitlocker InformationProtectionGetBitlocker(ctx).Select_(select_).Expand(expand).Execute()

Get bitlocker from informationProtection

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionBitlockerApi.InformationProtectionGetBitlocker(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionBitlockerApi.InformationProtectionGetBitlocker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionGetBitlocker`: MicrosoftGraphBitlocker
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionBitlockerApi.InformationProtectionGetBitlocker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionGetBitlockerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphBitlocker**](MicrosoftGraphBitlocker.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionUpdateBitlocker

> InformationProtectionUpdateBitlocker(ctx).MicrosoftGraphBitlocker(microsoftGraphBitlocker).Execute()

Update the navigation property bitlocker in informationProtection

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
    microsoftGraphBitlocker := *openapiclient.NewMicrosoftGraphBitlocker() // MicrosoftGraphBitlocker | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionBitlockerApi.InformationProtectionUpdateBitlocker(context.Background()).MicrosoftGraphBitlocker(microsoftGraphBitlocker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionBitlockerApi.InformationProtectionUpdateBitlocker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionUpdateBitlockerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphBitlocker** | [**MicrosoftGraphBitlocker**](MicrosoftGraphBitlocker.md) | New navigation property values | 

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

