# \AgreementsAgreementApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgreementsAgreementCreateAgreement**](AgreementsAgreementApi.md#AgreementsAgreementCreateAgreement) | **Post** /agreements | Add new entity to agreements
[**AgreementsAgreementDeleteAgreement**](AgreementsAgreementApi.md#AgreementsAgreementDeleteAgreement) | **Delete** /agreements/{agreement-id} | Delete entity from agreements
[**AgreementsAgreementGetAgreement**](AgreementsAgreementApi.md#AgreementsAgreementGetAgreement) | **Get** /agreements/{agreement-id} | Get entity from agreements by key
[**AgreementsAgreementListAgreement**](AgreementsAgreementApi.md#AgreementsAgreementListAgreement) | **Get** /agreements | Get entities from agreements
[**AgreementsAgreementUpdateAgreement**](AgreementsAgreementApi.md#AgreementsAgreementUpdateAgreement) | **Patch** /agreements/{agreement-id} | Update entity in agreements



## AgreementsAgreementCreateAgreement

> MicrosoftGraphAgreement AgreementsAgreementCreateAgreement(ctx).MicrosoftGraphAgreement(microsoftGraphAgreement).Execute()

Add new entity to agreements

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
    microsoftGraphAgreement := *openapiclient.NewMicrosoftGraphAgreement() // MicrosoftGraphAgreement | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementApi.AgreementsAgreementCreateAgreement(context.Background()).MicrosoftGraphAgreement(microsoftGraphAgreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementApi.AgreementsAgreementCreateAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsAgreementCreateAgreement`: MicrosoftGraphAgreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementApi.AgreementsAgreementCreateAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsAgreementCreateAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAgreement** | [**MicrosoftGraphAgreement**](MicrosoftGraphAgreement.md) | New entity | 

### Return type

[**MicrosoftGraphAgreement**](MicrosoftGraphAgreement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsAgreementDeleteAgreement

> AgreementsAgreementDeleteAgreement(ctx, agreementId).IfMatch(ifMatch).Execute()

Delete entity from agreements

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
    agreementId := "agreementId_example" // string | key: id of agreement
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementApi.AgreementsAgreementDeleteAgreement(context.Background(), agreementId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementApi.AgreementsAgreementDeleteAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsAgreementDeleteAgreementRequest struct via the builder pattern


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


## AgreementsAgreementGetAgreement

> MicrosoftGraphAgreement AgreementsAgreementGetAgreement(ctx, agreementId).Select_(select_).Execute()

Get entity from agreements by key

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
    agreementId := "agreementId_example" // string | key: id of agreement
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementApi.AgreementsAgreementGetAgreement(context.Background(), agreementId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementApi.AgreementsAgreementGetAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsAgreementGetAgreement`: MicrosoftGraphAgreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementApi.AgreementsAgreementGetAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsAgreementGetAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphAgreement**](MicrosoftGraphAgreement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsAgreementListAgreement

> CollectionOfAgreement AgreementsAgreementListAgreement(ctx).Search(search).Select_(select_).Execute()

Get entities from agreements

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
    search := "search_example" // string | Search items by search phrases (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementApi.AgreementsAgreementListAgreement(context.Background()).Search(search).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementApi.AgreementsAgreementListAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsAgreementListAgreement`: CollectionOfAgreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementApi.AgreementsAgreementListAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsAgreementListAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search items by search phrases | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfAgreement**](CollectionOfAgreement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsAgreementUpdateAgreement

> AgreementsAgreementUpdateAgreement(ctx, agreementId).MicrosoftGraphAgreement(microsoftGraphAgreement).Execute()

Update entity in agreements

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
    agreementId := "agreementId_example" // string | key: id of agreement
    microsoftGraphAgreement := *openapiclient.NewMicrosoftGraphAgreement() // MicrosoftGraphAgreement | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementApi.AgreementsAgreementUpdateAgreement(context.Background(), agreementId).MicrosoftGraphAgreement(microsoftGraphAgreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementApi.AgreementsAgreementUpdateAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsAgreementUpdateAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAgreement** | [**MicrosoftGraphAgreement**](MicrosoftGraphAgreement.md) | New property values | 

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

