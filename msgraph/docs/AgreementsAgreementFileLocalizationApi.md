# \AgreementsAgreementFileLocalizationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgreementsCreateFiles**](AgreementsAgreementFileLocalizationApi.md#AgreementsCreateFiles) | **Post** /agreements/{agreement-id}/files | Create new navigation property to files for agreements
[**AgreementsDeleteFiles**](AgreementsAgreementFileLocalizationApi.md#AgreementsDeleteFiles) | **Delete** /agreements/{agreement-id}/files/{agreementFileLocalization-id} | Delete navigation property files for agreements
[**AgreementsFilesCreateVersions**](AgreementsAgreementFileLocalizationApi.md#AgreementsFilesCreateVersions) | **Post** /agreements/{agreement-id}/files/{agreementFileLocalization-id}/versions | Create new navigation property to versions for agreements
[**AgreementsFilesDeleteVersions**](AgreementsAgreementFileLocalizationApi.md#AgreementsFilesDeleteVersions) | **Delete** /agreements/{agreement-id}/files/{agreementFileLocalization-id}/versions/{agreementFileVersion-id} | Delete navigation property versions for agreements
[**AgreementsFilesGetVersions**](AgreementsAgreementFileLocalizationApi.md#AgreementsFilesGetVersions) | **Get** /agreements/{agreement-id}/files/{agreementFileLocalization-id}/versions/{agreementFileVersion-id} | Get versions from agreements
[**AgreementsFilesListVersions**](AgreementsAgreementFileLocalizationApi.md#AgreementsFilesListVersions) | **Get** /agreements/{agreement-id}/files/{agreementFileLocalization-id}/versions | Get versions from agreements
[**AgreementsFilesUpdateVersions**](AgreementsAgreementFileLocalizationApi.md#AgreementsFilesUpdateVersions) | **Patch** /agreements/{agreement-id}/files/{agreementFileLocalization-id}/versions/{agreementFileVersion-id} | Update the navigation property versions in agreements
[**AgreementsGetFiles**](AgreementsAgreementFileLocalizationApi.md#AgreementsGetFiles) | **Get** /agreements/{agreement-id}/files/{agreementFileLocalization-id} | Get files from agreements
[**AgreementsListFiles**](AgreementsAgreementFileLocalizationApi.md#AgreementsListFiles) | **Get** /agreements/{agreement-id}/files | Get files from agreements
[**AgreementsUpdateFiles**](AgreementsAgreementFileLocalizationApi.md#AgreementsUpdateFiles) | **Patch** /agreements/{agreement-id}/files/{agreementFileLocalization-id} | Update the navigation property files in agreements



## AgreementsCreateFiles

> MicrosoftGraphAgreementFileLocalization AgreementsCreateFiles(ctx, agreementId).MicrosoftGraphAgreementFileLocalization(microsoftGraphAgreementFileLocalization).Execute()

Create new navigation property to files for agreements



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
    microsoftGraphAgreementFileLocalization := *openapiclient.NewMicrosoftGraphAgreementFileLocalization() // MicrosoftGraphAgreementFileLocalization | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsCreateFiles(context.Background(), agreementId).MicrosoftGraphAgreementFileLocalization(microsoftGraphAgreementFileLocalization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsCreateFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsCreateFiles`: MicrosoftGraphAgreementFileLocalization
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileLocalizationApi.AgreementsCreateFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsCreateFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAgreementFileLocalization** | [**MicrosoftGraphAgreementFileLocalization**](MicrosoftGraphAgreementFileLocalization.md) | New navigation property | 

### Return type

[**MicrosoftGraphAgreementFileLocalization**](MicrosoftGraphAgreementFileLocalization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsDeleteFiles

> AgreementsDeleteFiles(ctx, agreementId, agreementFileLocalizationId).IfMatch(ifMatch).Execute()

Delete navigation property files for agreements



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
    agreementFileLocalizationId := "agreementFileLocalizationId_example" // string | key: id of agreementFileLocalization
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsDeleteFiles(context.Background(), agreementId, agreementFileLocalizationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsDeleteFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsDeleteFilesRequest struct via the builder pattern


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


## AgreementsFilesCreateVersions

> MicrosoftGraphAgreementFileVersion AgreementsFilesCreateVersions(ctx, agreementId, agreementFileLocalizationId).MicrosoftGraphAgreementFileVersion(microsoftGraphAgreementFileVersion).Execute()

Create new navigation property to versions for agreements

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
    agreementFileLocalizationId := "agreementFileLocalizationId_example" // string | key: id of agreementFileLocalization
    microsoftGraphAgreementFileVersion := *openapiclient.NewMicrosoftGraphAgreementFileVersion() // MicrosoftGraphAgreementFileVersion | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsFilesCreateVersions(context.Background(), agreementId, agreementFileLocalizationId).MicrosoftGraphAgreementFileVersion(microsoftGraphAgreementFileVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsFilesCreateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFilesCreateVersions`: MicrosoftGraphAgreementFileVersion
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileLocalizationApi.AgreementsFilesCreateVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFilesCreateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAgreementFileVersion** | [**MicrosoftGraphAgreementFileVersion**](MicrosoftGraphAgreementFileVersion.md) | New navigation property | 

### Return type

[**MicrosoftGraphAgreementFileVersion**](MicrosoftGraphAgreementFileVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsFilesDeleteVersions

> AgreementsFilesDeleteVersions(ctx, agreementId, agreementFileLocalizationId, agreementFileVersionId).IfMatch(ifMatch).Execute()

Delete navigation property versions for agreements

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
    agreementFileLocalizationId := "agreementFileLocalizationId_example" // string | key: id of agreementFileLocalization
    agreementFileVersionId := "agreementFileVersionId_example" // string | key: id of agreementFileVersion
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsFilesDeleteVersions(context.Background(), agreementId, agreementFileLocalizationId, agreementFileVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsFilesDeleteVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 
**agreementFileVersionId** | **string** | key: id of agreementFileVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFilesDeleteVersionsRequest struct via the builder pattern


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


## AgreementsFilesGetVersions

> MicrosoftGraphAgreementFileVersion AgreementsFilesGetVersions(ctx, agreementId, agreementFileLocalizationId, agreementFileVersionId).Select_(select_).Expand(expand).Execute()

Get versions from agreements

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
    agreementFileLocalizationId := "agreementFileLocalizationId_example" // string | key: id of agreementFileLocalization
    agreementFileVersionId := "agreementFileVersionId_example" // string | key: id of agreementFileVersion
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsFilesGetVersions(context.Background(), agreementId, agreementFileLocalizationId, agreementFileVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsFilesGetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFilesGetVersions`: MicrosoftGraphAgreementFileVersion
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileLocalizationApi.AgreementsFilesGetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 
**agreementFileVersionId** | **string** | key: id of agreementFileVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFilesGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAgreementFileVersion**](MicrosoftGraphAgreementFileVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsFilesListVersions

> CollectionOfAgreementFileVersion AgreementsFilesListVersions(ctx, agreementId, agreementFileLocalizationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get versions from agreements

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
    agreementFileLocalizationId := "agreementFileLocalizationId_example" // string | key: id of agreementFileLocalization
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
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsFilesListVersions(context.Background(), agreementId, agreementFileLocalizationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsFilesListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFilesListVersions`: CollectionOfAgreementFileVersion
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileLocalizationApi.AgreementsFilesListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFilesListVersionsRequest struct via the builder pattern


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

[**CollectionOfAgreementFileVersion**](CollectionOfAgreementFileVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsFilesUpdateVersions

> AgreementsFilesUpdateVersions(ctx, agreementId, agreementFileLocalizationId, agreementFileVersionId).MicrosoftGraphAgreementFileVersion(microsoftGraphAgreementFileVersion).Execute()

Update the navigation property versions in agreements

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
    agreementFileLocalizationId := "agreementFileLocalizationId_example" // string | key: id of agreementFileLocalization
    agreementFileVersionId := "agreementFileVersionId_example" // string | key: id of agreementFileVersion
    microsoftGraphAgreementFileVersion := *openapiclient.NewMicrosoftGraphAgreementFileVersion() // MicrosoftGraphAgreementFileVersion | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsFilesUpdateVersions(context.Background(), agreementId, agreementFileLocalizationId, agreementFileVersionId).MicrosoftGraphAgreementFileVersion(microsoftGraphAgreementFileVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsFilesUpdateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 
**agreementFileVersionId** | **string** | key: id of agreementFileVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFilesUpdateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphAgreementFileVersion** | [**MicrosoftGraphAgreementFileVersion**](MicrosoftGraphAgreementFileVersion.md) | New navigation property values | 

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


## AgreementsGetFiles

> MicrosoftGraphAgreementFileLocalization AgreementsGetFiles(ctx, agreementId, agreementFileLocalizationId).Select_(select_).Expand(expand).Execute()

Get files from agreements



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
    agreementFileLocalizationId := "agreementFileLocalizationId_example" // string | key: id of agreementFileLocalization
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsGetFiles(context.Background(), agreementId, agreementFileLocalizationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsGetFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsGetFiles`: MicrosoftGraphAgreementFileLocalization
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileLocalizationApi.AgreementsGetFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsGetFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAgreementFileLocalization**](MicrosoftGraphAgreementFileLocalization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsListFiles

> CollectionOfAgreementFileLocalization AgreementsListFiles(ctx, agreementId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get files from agreements



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
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsListFiles(context.Background(), agreementId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsListFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsListFiles`: CollectionOfAgreementFileLocalization
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileLocalizationApi.AgreementsListFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsListFilesRequest struct via the builder pattern


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

[**CollectionOfAgreementFileLocalization**](CollectionOfAgreementFileLocalization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsUpdateFiles

> AgreementsUpdateFiles(ctx, agreementId, agreementFileLocalizationId).MicrosoftGraphAgreementFileLocalization(microsoftGraphAgreementFileLocalization).Execute()

Update the navigation property files in agreements



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
    agreementFileLocalizationId := "agreementFileLocalizationId_example" // string | key: id of agreementFileLocalization
    microsoftGraphAgreementFileLocalization := *openapiclient.NewMicrosoftGraphAgreementFileLocalization() // MicrosoftGraphAgreementFileLocalization | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileLocalizationApi.AgreementsUpdateFiles(context.Background(), agreementId, agreementFileLocalizationId).MicrosoftGraphAgreementFileLocalization(microsoftGraphAgreementFileLocalization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileLocalizationApi.AgreementsUpdateFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsUpdateFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAgreementFileLocalization** | [**MicrosoftGraphAgreementFileLocalization**](MicrosoftGraphAgreementFileLocalization.md) | New navigation property values | 

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

