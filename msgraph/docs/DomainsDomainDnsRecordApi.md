# \DomainsDomainDnsRecordApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsCreateServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsCreateServiceConfigurationRecords) | **Post** /domains/{domain-id}/serviceConfigurationRecords | Create new navigation property to serviceConfigurationRecords for domains
[**DomainsCreateVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsCreateVerificationDnsRecords) | **Post** /domains/{domain-id}/verificationDnsRecords | Create new navigation property to verificationDnsRecords for domains
[**DomainsDeleteServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsDeleteServiceConfigurationRecords) | **Delete** /domains/{domain-id}/serviceConfigurationRecords/{domainDnsRecord-id} | Delete navigation property serviceConfigurationRecords for domains
[**DomainsDeleteVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsDeleteVerificationDnsRecords) | **Delete** /domains/{domain-id}/verificationDnsRecords/{domainDnsRecord-id} | Delete navigation property verificationDnsRecords for domains
[**DomainsGetServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsGetServiceConfigurationRecords) | **Get** /domains/{domain-id}/serviceConfigurationRecords/{domainDnsRecord-id} | Get serviceConfigurationRecords from domains
[**DomainsGetVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsGetVerificationDnsRecords) | **Get** /domains/{domain-id}/verificationDnsRecords/{domainDnsRecord-id} | Get verificationDnsRecords from domains
[**DomainsListServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsListServiceConfigurationRecords) | **Get** /domains/{domain-id}/serviceConfigurationRecords | Get serviceConfigurationRecords from domains
[**DomainsListVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsListVerificationDnsRecords) | **Get** /domains/{domain-id}/verificationDnsRecords | Get verificationDnsRecords from domains
[**DomainsUpdateServiceConfigurationRecords**](DomainsDomainDnsRecordApi.md#DomainsUpdateServiceConfigurationRecords) | **Patch** /domains/{domain-id}/serviceConfigurationRecords/{domainDnsRecord-id} | Update the navigation property serviceConfigurationRecords in domains
[**DomainsUpdateVerificationDnsRecords**](DomainsDomainDnsRecordApi.md#DomainsUpdateVerificationDnsRecords) | **Patch** /domains/{domain-id}/verificationDnsRecords/{domainDnsRecord-id} | Update the navigation property verificationDnsRecords in domains



## DomainsCreateServiceConfigurationRecords

> MicrosoftGraphDomainDnsRecord DomainsCreateServiceConfigurationRecords(ctx, domainId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()

Create new navigation property to serviceConfigurationRecords for domains



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
    domainId := "domainId_example" // string | key: id of domain
    microsoftGraphDomainDnsRecord := *openapiclient.NewMicrosoftGraphDomainDnsRecord() // MicrosoftGraphDomainDnsRecord | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsCreateServiceConfigurationRecords(context.Background(), domainId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsCreateServiceConfigurationRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsCreateServiceConfigurationRecords`: MicrosoftGraphDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsDomainDnsRecordApi.DomainsCreateServiceConfigurationRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCreateServiceConfigurationRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md) | New navigation property | 

### Return type

[**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsCreateVerificationDnsRecords

> MicrosoftGraphDomainDnsRecord DomainsCreateVerificationDnsRecords(ctx, domainId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()

Create new navigation property to verificationDnsRecords for domains



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
    domainId := "domainId_example" // string | key: id of domain
    microsoftGraphDomainDnsRecord := *openapiclient.NewMicrosoftGraphDomainDnsRecord() // MicrosoftGraphDomainDnsRecord | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsCreateVerificationDnsRecords(context.Background(), domainId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsCreateVerificationDnsRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsCreateVerificationDnsRecords`: MicrosoftGraphDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsDomainDnsRecordApi.DomainsCreateVerificationDnsRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCreateVerificationDnsRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md) | New navigation property | 

### Return type

[**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDeleteServiceConfigurationRecords

> DomainsDeleteServiceConfigurationRecords(ctx, domainId, domainDnsRecordId).IfMatch(ifMatch).Execute()

Delete navigation property serviceConfigurationRecords for domains



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
    domainId := "domainId_example" // string | key: id of domain
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsDeleteServiceConfigurationRecords(context.Background(), domainId, domainDnsRecordId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsDeleteServiceConfigurationRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDeleteServiceConfigurationRecordsRequest struct via the builder pattern


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


## DomainsDeleteVerificationDnsRecords

> DomainsDeleteVerificationDnsRecords(ctx, domainId, domainDnsRecordId).IfMatch(ifMatch).Execute()

Delete navigation property verificationDnsRecords for domains



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
    domainId := "domainId_example" // string | key: id of domain
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsDeleteVerificationDnsRecords(context.Background(), domainId, domainDnsRecordId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsDeleteVerificationDnsRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDeleteVerificationDnsRecordsRequest struct via the builder pattern


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


## DomainsGetServiceConfigurationRecords

> MicrosoftGraphDomainDnsRecord DomainsGetServiceConfigurationRecords(ctx, domainId, domainDnsRecordId).Select_(select_).Expand(expand).Execute()

Get serviceConfigurationRecords from domains



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
    domainId := "domainId_example" // string | key: id of domain
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsGetServiceConfigurationRecords(context.Background(), domainId, domainDnsRecordId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsGetServiceConfigurationRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsGetServiceConfigurationRecords`: MicrosoftGraphDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsDomainDnsRecordApi.DomainsGetServiceConfigurationRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsGetServiceConfigurationRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsGetVerificationDnsRecords

> MicrosoftGraphDomainDnsRecord DomainsGetVerificationDnsRecords(ctx, domainId, domainDnsRecordId).Select_(select_).Expand(expand).Execute()

Get verificationDnsRecords from domains



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
    domainId := "domainId_example" // string | key: id of domain
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsGetVerificationDnsRecords(context.Background(), domainId, domainDnsRecordId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsGetVerificationDnsRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsGetVerificationDnsRecords`: MicrosoftGraphDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsDomainDnsRecordApi.DomainsGetVerificationDnsRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsGetVerificationDnsRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsListServiceConfigurationRecords

> CollectionOfDomainDnsRecord DomainsListServiceConfigurationRecords(ctx, domainId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get serviceConfigurationRecords from domains



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
    domainId := "domainId_example" // string | key: id of domain
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
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsListServiceConfigurationRecords(context.Background(), domainId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsListServiceConfigurationRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsListServiceConfigurationRecords`: CollectionOfDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsDomainDnsRecordApi.DomainsListServiceConfigurationRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsListServiceConfigurationRecordsRequest struct via the builder pattern


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

[**CollectionOfDomainDnsRecord**](CollectionOfDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsListVerificationDnsRecords

> CollectionOfDomainDnsRecord DomainsListVerificationDnsRecords(ctx, domainId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get verificationDnsRecords from domains



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
    domainId := "domainId_example" // string | key: id of domain
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
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsListVerificationDnsRecords(context.Background(), domainId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsListVerificationDnsRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsListVerificationDnsRecords`: CollectionOfDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainsDomainDnsRecordApi.DomainsListVerificationDnsRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsListVerificationDnsRecordsRequest struct via the builder pattern


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

[**CollectionOfDomainDnsRecord**](CollectionOfDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsUpdateServiceConfigurationRecords

> DomainsUpdateServiceConfigurationRecords(ctx, domainId, domainDnsRecordId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()

Update the navigation property serviceConfigurationRecords in domains



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
    domainId := "domainId_example" // string | key: id of domain
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    microsoftGraphDomainDnsRecord := *openapiclient.NewMicrosoftGraphDomainDnsRecord() // MicrosoftGraphDomainDnsRecord | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsUpdateServiceConfigurationRecords(context.Background(), domainId, domainDnsRecordId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsUpdateServiceConfigurationRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsUpdateServiceConfigurationRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md) | New navigation property values | 

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


## DomainsUpdateVerificationDnsRecords

> DomainsUpdateVerificationDnsRecords(ctx, domainId, domainDnsRecordId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()

Update the navigation property verificationDnsRecords in domains



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
    domainId := "domainId_example" // string | key: id of domain
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    microsoftGraphDomainDnsRecord := *openapiclient.NewMicrosoftGraphDomainDnsRecord() // MicrosoftGraphDomainDnsRecord | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsDomainDnsRecordApi.DomainsUpdateVerificationDnsRecords(context.Background(), domainId, domainDnsRecordId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsDomainDnsRecordApi.DomainsUpdateVerificationDnsRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsUpdateVerificationDnsRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md) | New navigation property values | 

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

