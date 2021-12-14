# \AgreementAcceptancesAgreementAcceptanceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance**](AgreementAcceptancesAgreementAcceptanceApi.md#AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance) | **Post** /agreementAcceptances | Add new entity to agreementAcceptances
[**AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance**](AgreementAcceptancesAgreementAcceptanceApi.md#AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance) | **Delete** /agreementAcceptances/{agreementAcceptance-id} | Delete entity from agreementAcceptances
[**AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance**](AgreementAcceptancesAgreementAcceptanceApi.md#AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance) | **Get** /agreementAcceptances/{agreementAcceptance-id} | Get entity from agreementAcceptances by key
[**AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance**](AgreementAcceptancesAgreementAcceptanceApi.md#AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance) | **Get** /agreementAcceptances | Get entities from agreementAcceptances
[**AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance**](AgreementAcceptancesAgreementAcceptanceApi.md#AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance) | **Patch** /agreementAcceptances/{agreementAcceptance-id} | Update entity in agreementAcceptances



## AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance

> MicrosoftGraphAgreementAcceptance AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance(ctx).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()

Add new entity to agreementAcceptances

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
    microsoftGraphAgreementAcceptance := *openapiclient.NewMicrosoftGraphAgreementAcceptance() // MicrosoftGraphAgreementAcceptance | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance(context.Background()).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance`: MicrosoftGraphAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAgreementAcceptance** | [**MicrosoftGraphAgreementAcceptance**](MicrosoftGraphAgreementAcceptance.md) | New entity | 

### Return type

[**MicrosoftGraphAgreementAcceptance**](MicrosoftGraphAgreementAcceptance.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance

> AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance(ctx, agreementAcceptanceId).IfMatch(ifMatch).Execute()

Delete entity from agreementAcceptances

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
    agreementAcceptanceId := "agreementAcceptanceId_example" // string | key: id of agreementAcceptance
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance(context.Background(), agreementAcceptanceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementAcceptanceId** | **string** | key: id of agreementAcceptance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest struct via the builder pattern


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


## AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance

> MicrosoftGraphAgreementAcceptance AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance(ctx, agreementAcceptanceId).Select_(select_).Execute()

Get entity from agreementAcceptances by key

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
    agreementAcceptanceId := "agreementAcceptanceId_example" // string | key: id of agreementAcceptance
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance(context.Background(), agreementAcceptanceId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance`: MicrosoftGraphAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementAcceptanceId** | **string** | key: id of agreementAcceptance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphAgreementAcceptance**](MicrosoftGraphAgreementAcceptance.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance

> CollectionOfAgreementAcceptance AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance(ctx).Search(search).Select_(select_).Execute()

Get entities from agreementAcceptances

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
    resp, r, err := api_client.AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance(context.Background()).Search(search).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance`: CollectionOfAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search items by search phrases | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfAgreementAcceptance**](CollectionOfAgreementAcceptance.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance

> AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance(ctx, agreementAcceptanceId).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()

Update entity in agreementAcceptances

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
    agreementAcceptanceId := "agreementAcceptanceId_example" // string | key: id of agreementAcceptance
    microsoftGraphAgreementAcceptance := *openapiclient.NewMicrosoftGraphAgreementAcceptance() // MicrosoftGraphAgreementAcceptance | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance(context.Background(), agreementAcceptanceId).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementAcceptancesAgreementAcceptanceApi.AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementAcceptanceId** | **string** | key: id of agreementAcceptance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAgreementAcceptance** | [**MicrosoftGraphAgreementAcceptance**](MicrosoftGraphAgreementAcceptance.md) | New property values | 

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

