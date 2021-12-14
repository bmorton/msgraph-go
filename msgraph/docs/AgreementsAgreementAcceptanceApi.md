# \AgreementsAgreementAcceptanceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgreementsCreateAcceptances**](AgreementsAgreementAcceptanceApi.md#AgreementsCreateAcceptances) | **Post** /agreements/{agreement-id}/acceptances | Create new navigation property to acceptances for agreements
[**AgreementsDeleteAcceptances**](AgreementsAgreementAcceptanceApi.md#AgreementsDeleteAcceptances) | **Delete** /agreements/{agreement-id}/acceptances/{agreementAcceptance-id} | Delete navigation property acceptances for agreements
[**AgreementsGetAcceptances**](AgreementsAgreementAcceptanceApi.md#AgreementsGetAcceptances) | **Get** /agreements/{agreement-id}/acceptances/{agreementAcceptance-id} | Get acceptances from agreements
[**AgreementsListAcceptances**](AgreementsAgreementAcceptanceApi.md#AgreementsListAcceptances) | **Get** /agreements/{agreement-id}/acceptances | Get acceptances from agreements
[**AgreementsUpdateAcceptances**](AgreementsAgreementAcceptanceApi.md#AgreementsUpdateAcceptances) | **Patch** /agreements/{agreement-id}/acceptances/{agreementAcceptance-id} | Update the navigation property acceptances in agreements



## AgreementsCreateAcceptances

> MicrosoftGraphAgreementAcceptance AgreementsCreateAcceptances(ctx, agreementId).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()

Create new navigation property to acceptances for agreements



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
    microsoftGraphAgreementAcceptance := *openapiclient.NewMicrosoftGraphAgreementAcceptance() // MicrosoftGraphAgreementAcceptance | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementAcceptanceApi.AgreementsCreateAcceptances(context.Background(), agreementId).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementAcceptanceApi.AgreementsCreateAcceptances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsCreateAcceptances`: MicrosoftGraphAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementAcceptanceApi.AgreementsCreateAcceptances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsCreateAcceptancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAgreementAcceptance** | [**MicrosoftGraphAgreementAcceptance**](MicrosoftGraphAgreementAcceptance.md) | New navigation property | 

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


## AgreementsDeleteAcceptances

> AgreementsDeleteAcceptances(ctx, agreementId, agreementAcceptanceId).IfMatch(ifMatch).Execute()

Delete navigation property acceptances for agreements



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
    agreementAcceptanceId := "agreementAcceptanceId_example" // string | key: id of agreementAcceptance
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementAcceptanceApi.AgreementsDeleteAcceptances(context.Background(), agreementId, agreementAcceptanceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementAcceptanceApi.AgreementsDeleteAcceptances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementAcceptanceId** | **string** | key: id of agreementAcceptance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsDeleteAcceptancesRequest struct via the builder pattern


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


## AgreementsGetAcceptances

> MicrosoftGraphAgreementAcceptance AgreementsGetAcceptances(ctx, agreementId, agreementAcceptanceId).Select_(select_).Expand(expand).Execute()

Get acceptances from agreements



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
    agreementAcceptanceId := "agreementAcceptanceId_example" // string | key: id of agreementAcceptance
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementAcceptanceApi.AgreementsGetAcceptances(context.Background(), agreementId, agreementAcceptanceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementAcceptanceApi.AgreementsGetAcceptances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsGetAcceptances`: MicrosoftGraphAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementAcceptanceApi.AgreementsGetAcceptances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementAcceptanceId** | **string** | key: id of agreementAcceptance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsGetAcceptancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

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


## AgreementsListAcceptances

> CollectionOfAgreementAcceptance AgreementsListAcceptances(ctx, agreementId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get acceptances from agreements



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
    resp, r, err := api_client.AgreementsAgreementAcceptanceApi.AgreementsListAcceptances(context.Background(), agreementId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementAcceptanceApi.AgreementsListAcceptances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsListAcceptances`: CollectionOfAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementAcceptanceApi.AgreementsListAcceptances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsListAcceptancesRequest struct via the builder pattern


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

[**CollectionOfAgreementAcceptance**](CollectionOfAgreementAcceptance.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsUpdateAcceptances

> AgreementsUpdateAcceptances(ctx, agreementId, agreementAcceptanceId).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()

Update the navigation property acceptances in agreements



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
    agreementAcceptanceId := "agreementAcceptanceId_example" // string | key: id of agreementAcceptance
    microsoftGraphAgreementAcceptance := *openapiclient.NewMicrosoftGraphAgreementAcceptance() // MicrosoftGraphAgreementAcceptance | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementAcceptanceApi.AgreementsUpdateAcceptances(context.Background(), agreementId, agreementAcceptanceId).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementAcceptanceApi.AgreementsUpdateAcceptances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementAcceptanceId** | **string** | key: id of agreementAcceptance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsUpdateAcceptancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAgreementAcceptance** | [**MicrosoftGraphAgreementAcceptance**](MicrosoftGraphAgreementAcceptance.md) | New navigation property values | 

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

