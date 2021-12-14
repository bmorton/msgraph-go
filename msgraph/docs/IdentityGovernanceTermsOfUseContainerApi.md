# \IdentityGovernanceTermsOfUseContainerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityGovernanceDeleteTermsOfUse**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceDeleteTermsOfUse) | **Delete** /identityGovernance/termsOfUse | Delete navigation property termsOfUse for identityGovernance
[**IdentityGovernanceGetTermsOfUse**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceGetTermsOfUse) | **Get** /identityGovernance/termsOfUse | Get termsOfUse from identityGovernance
[**IdentityGovernanceTermsOfUseCreateAgreementAcceptances**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseCreateAgreementAcceptances) | **Post** /identityGovernance/termsOfUse/agreementAcceptances | Create new navigation property to agreementAcceptances for identityGovernance
[**IdentityGovernanceTermsOfUseCreateAgreements**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseCreateAgreements) | **Post** /identityGovernance/termsOfUse/agreements | Create new navigation property to agreements for identityGovernance
[**IdentityGovernanceTermsOfUseDeleteAgreementAcceptances**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseDeleteAgreementAcceptances) | **Delete** /identityGovernance/termsOfUse/agreementAcceptances/{agreementAcceptance-id} | Delete navigation property agreementAcceptances for identityGovernance
[**IdentityGovernanceTermsOfUseDeleteAgreements**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseDeleteAgreements) | **Delete** /identityGovernance/termsOfUse/agreements/{agreement-id} | Delete navigation property agreements for identityGovernance
[**IdentityGovernanceTermsOfUseGetAgreementAcceptances**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseGetAgreementAcceptances) | **Get** /identityGovernance/termsOfUse/agreementAcceptances/{agreementAcceptance-id} | Get agreementAcceptances from identityGovernance
[**IdentityGovernanceTermsOfUseGetAgreements**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseGetAgreements) | **Get** /identityGovernance/termsOfUse/agreements/{agreement-id} | Get agreements from identityGovernance
[**IdentityGovernanceTermsOfUseListAgreementAcceptances**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseListAgreementAcceptances) | **Get** /identityGovernance/termsOfUse/agreementAcceptances | Get agreementAcceptances from identityGovernance
[**IdentityGovernanceTermsOfUseListAgreements**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseListAgreements) | **Get** /identityGovernance/termsOfUse/agreements | Get agreements from identityGovernance
[**IdentityGovernanceTermsOfUseUpdateAgreementAcceptances**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseUpdateAgreementAcceptances) | **Patch** /identityGovernance/termsOfUse/agreementAcceptances/{agreementAcceptance-id} | Update the navigation property agreementAcceptances in identityGovernance
[**IdentityGovernanceTermsOfUseUpdateAgreements**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceTermsOfUseUpdateAgreements) | **Patch** /identityGovernance/termsOfUse/agreements/{agreement-id} | Update the navigation property agreements in identityGovernance
[**IdentityGovernanceUpdateTermsOfUse**](IdentityGovernanceTermsOfUseContainerApi.md#IdentityGovernanceUpdateTermsOfUse) | **Patch** /identityGovernance/termsOfUse | Update the navigation property termsOfUse in identityGovernance



## IdentityGovernanceDeleteTermsOfUse

> IdentityGovernanceDeleteTermsOfUse(ctx).IfMatch(ifMatch).Execute()

Delete navigation property termsOfUse for identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceDeleteTermsOfUse(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceDeleteTermsOfUse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceDeleteTermsOfUseRequest struct via the builder pattern


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


## IdentityGovernanceGetTermsOfUse

> MicrosoftGraphTermsOfUseContainer IdentityGovernanceGetTermsOfUse(ctx).Select_(select_).Expand(expand).Execute()

Get termsOfUse from identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceGetTermsOfUse(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceGetTermsOfUse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceGetTermsOfUse`: MicrosoftGraphTermsOfUseContainer
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceGetTermsOfUse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceGetTermsOfUseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTermsOfUseContainer**](MicrosoftGraphTermsOfUseContainer.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceTermsOfUseCreateAgreementAcceptances

> MicrosoftGraphAgreementAcceptance IdentityGovernanceTermsOfUseCreateAgreementAcceptances(ctx).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()

Create new navigation property to agreementAcceptances for identityGovernance

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
    microsoftGraphAgreementAcceptance := *openapiclient.NewMicrosoftGraphAgreementAcceptance() // MicrosoftGraphAgreementAcceptance | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseCreateAgreementAcceptances(context.Background()).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseCreateAgreementAcceptances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceTermsOfUseCreateAgreementAcceptances`: MicrosoftGraphAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseCreateAgreementAcceptances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseCreateAgreementAcceptancesRequest struct via the builder pattern


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


## IdentityGovernanceTermsOfUseCreateAgreements

> MicrosoftGraphAgreement IdentityGovernanceTermsOfUseCreateAgreements(ctx).MicrosoftGraphAgreement(microsoftGraphAgreement).Execute()

Create new navigation property to agreements for identityGovernance

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
    microsoftGraphAgreement := *openapiclient.NewMicrosoftGraphAgreement() // MicrosoftGraphAgreement | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseCreateAgreements(context.Background()).MicrosoftGraphAgreement(microsoftGraphAgreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseCreateAgreements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceTermsOfUseCreateAgreements`: MicrosoftGraphAgreement
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseCreateAgreements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseCreateAgreementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAgreement** | [**MicrosoftGraphAgreement**](MicrosoftGraphAgreement.md) | New navigation property | 

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


## IdentityGovernanceTermsOfUseDeleteAgreementAcceptances

> IdentityGovernanceTermsOfUseDeleteAgreementAcceptances(ctx, agreementAcceptanceId).IfMatch(ifMatch).Execute()

Delete navigation property agreementAcceptances for identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseDeleteAgreementAcceptances(context.Background(), agreementAcceptanceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseDeleteAgreementAcceptances``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseDeleteAgreementAcceptancesRequest struct via the builder pattern


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


## IdentityGovernanceTermsOfUseDeleteAgreements

> IdentityGovernanceTermsOfUseDeleteAgreements(ctx, agreementId).IfMatch(ifMatch).Execute()

Delete navigation property agreements for identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseDeleteAgreements(context.Background(), agreementId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseDeleteAgreements``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseDeleteAgreementsRequest struct via the builder pattern


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


## IdentityGovernanceTermsOfUseGetAgreementAcceptances

> MicrosoftGraphAgreementAcceptance IdentityGovernanceTermsOfUseGetAgreementAcceptances(ctx, agreementAcceptanceId).Select_(select_).Expand(expand).Execute()

Get agreementAcceptances from identityGovernance

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
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseGetAgreementAcceptances(context.Background(), agreementAcceptanceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseGetAgreementAcceptances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceTermsOfUseGetAgreementAcceptances`: MicrosoftGraphAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseGetAgreementAcceptances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementAcceptanceId** | **string** | key: id of agreementAcceptance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseGetAgreementAcceptancesRequest struct via the builder pattern


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


## IdentityGovernanceTermsOfUseGetAgreements

> MicrosoftGraphAgreement IdentityGovernanceTermsOfUseGetAgreements(ctx, agreementId).Select_(select_).Expand(expand).Execute()

Get agreements from identityGovernance

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
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseGetAgreements(context.Background(), agreementId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseGetAgreements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceTermsOfUseGetAgreements`: MicrosoftGraphAgreement
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseGetAgreements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseGetAgreementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

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


## IdentityGovernanceTermsOfUseListAgreementAcceptances

> CollectionOfAgreementAcceptance IdentityGovernanceTermsOfUseListAgreementAcceptances(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get agreementAcceptances from identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseListAgreementAcceptances(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseListAgreementAcceptances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceTermsOfUseListAgreementAcceptances`: CollectionOfAgreementAcceptance
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseListAgreementAcceptances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseListAgreementAcceptancesRequest struct via the builder pattern


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


## IdentityGovernanceTermsOfUseListAgreements

> CollectionOfAgreement IdentityGovernanceTermsOfUseListAgreements(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get agreements from identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseListAgreements(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseListAgreements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceTermsOfUseListAgreements`: CollectionOfAgreement
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseListAgreements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseListAgreementsRequest struct via the builder pattern


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

[**CollectionOfAgreement**](CollectionOfAgreement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceTermsOfUseUpdateAgreementAcceptances

> IdentityGovernanceTermsOfUseUpdateAgreementAcceptances(ctx, agreementAcceptanceId).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()

Update the navigation property agreementAcceptances in identityGovernance

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
    microsoftGraphAgreementAcceptance := *openapiclient.NewMicrosoftGraphAgreementAcceptance() // MicrosoftGraphAgreementAcceptance | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseUpdateAgreementAcceptances(context.Background(), agreementAcceptanceId).MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseUpdateAgreementAcceptances``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseUpdateAgreementAcceptancesRequest struct via the builder pattern


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


## IdentityGovernanceTermsOfUseUpdateAgreements

> IdentityGovernanceTermsOfUseUpdateAgreements(ctx, agreementId).MicrosoftGraphAgreement(microsoftGraphAgreement).Execute()

Update the navigation property agreements in identityGovernance

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
    microsoftGraphAgreement := *openapiclient.NewMicrosoftGraphAgreement() // MicrosoftGraphAgreement | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseUpdateAgreements(context.Background(), agreementId).MicrosoftGraphAgreement(microsoftGraphAgreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceTermsOfUseUpdateAgreements``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIdentityGovernanceTermsOfUseUpdateAgreementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAgreement** | [**MicrosoftGraphAgreement**](MicrosoftGraphAgreement.md) | New navigation property values | 

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


## IdentityGovernanceUpdateTermsOfUse

> IdentityGovernanceUpdateTermsOfUse(ctx).MicrosoftGraphTermsOfUseContainer(microsoftGraphTermsOfUseContainer).Execute()

Update the navigation property termsOfUse in identityGovernance

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
    microsoftGraphTermsOfUseContainer := *openapiclient.NewMicrosoftGraphTermsOfUseContainer() // MicrosoftGraphTermsOfUseContainer | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceUpdateTermsOfUse(context.Background()).MicrosoftGraphTermsOfUseContainer(microsoftGraphTermsOfUseContainer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceTermsOfUseContainerApi.IdentityGovernanceUpdateTermsOfUse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceUpdateTermsOfUseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTermsOfUseContainer** | [**MicrosoftGraphTermsOfUseContainer**](MicrosoftGraphTermsOfUseContainer.md) | New navigation property values | 

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

