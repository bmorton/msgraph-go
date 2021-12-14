# \AgreementsAgreementFileApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgreementsDeleteFile**](AgreementsAgreementFileApi.md#AgreementsDeleteFile) | **Delete** /agreements/{agreement-id}/file | Delete navigation property file for agreements
[**AgreementsFileCreateLocalizations**](AgreementsAgreementFileApi.md#AgreementsFileCreateLocalizations) | **Post** /agreements/{agreement-id}/file/localizations | Create new navigation property to localizations for agreements
[**AgreementsFileDeleteLocalizations**](AgreementsAgreementFileApi.md#AgreementsFileDeleteLocalizations) | **Delete** /agreements/{agreement-id}/file/localizations/{agreementFileLocalization-id} | Delete navigation property localizations for agreements
[**AgreementsFileGetLocalizations**](AgreementsAgreementFileApi.md#AgreementsFileGetLocalizations) | **Get** /agreements/{agreement-id}/file/localizations/{agreementFileLocalization-id} | Get localizations from agreements
[**AgreementsFileListLocalizations**](AgreementsAgreementFileApi.md#AgreementsFileListLocalizations) | **Get** /agreements/{agreement-id}/file/localizations | Get localizations from agreements
[**AgreementsFileLocalizationsCreateVersions**](AgreementsAgreementFileApi.md#AgreementsFileLocalizationsCreateVersions) | **Post** /agreements/{agreement-id}/file/localizations/{agreementFileLocalization-id}/versions | Create new navigation property to versions for agreements
[**AgreementsFileLocalizationsDeleteVersions**](AgreementsAgreementFileApi.md#AgreementsFileLocalizationsDeleteVersions) | **Delete** /agreements/{agreement-id}/file/localizations/{agreementFileLocalization-id}/versions/{agreementFileVersion-id} | Delete navigation property versions for agreements
[**AgreementsFileLocalizationsGetVersions**](AgreementsAgreementFileApi.md#AgreementsFileLocalizationsGetVersions) | **Get** /agreements/{agreement-id}/file/localizations/{agreementFileLocalization-id}/versions/{agreementFileVersion-id} | Get versions from agreements
[**AgreementsFileLocalizationsListVersions**](AgreementsAgreementFileApi.md#AgreementsFileLocalizationsListVersions) | **Get** /agreements/{agreement-id}/file/localizations/{agreementFileLocalization-id}/versions | Get versions from agreements
[**AgreementsFileLocalizationsUpdateVersions**](AgreementsAgreementFileApi.md#AgreementsFileLocalizationsUpdateVersions) | **Patch** /agreements/{agreement-id}/file/localizations/{agreementFileLocalization-id}/versions/{agreementFileVersion-id} | Update the navigation property versions in agreements
[**AgreementsFileUpdateLocalizations**](AgreementsAgreementFileApi.md#AgreementsFileUpdateLocalizations) | **Patch** /agreements/{agreement-id}/file/localizations/{agreementFileLocalization-id} | Update the navigation property localizations in agreements
[**AgreementsGetFile**](AgreementsAgreementFileApi.md#AgreementsGetFile) | **Get** /agreements/{agreement-id}/file | Get file from agreements
[**AgreementsUpdateFile**](AgreementsAgreementFileApi.md#AgreementsUpdateFile) | **Patch** /agreements/{agreement-id}/file | Update the navigation property file in agreements



## AgreementsDeleteFile

> AgreementsDeleteFile(ctx, agreementId).IfMatch(ifMatch).Execute()

Delete navigation property file for agreements



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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsDeleteFile(context.Background(), agreementId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsDeleteFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAgreementsDeleteFileRequest struct via the builder pattern


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


## AgreementsFileCreateLocalizations

> MicrosoftGraphAgreementFileLocalization AgreementsFileCreateLocalizations(ctx, agreementId).MicrosoftGraphAgreementFileLocalization(microsoftGraphAgreementFileLocalization).Execute()

Create new navigation property to localizations for agreements

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileCreateLocalizations(context.Background(), agreementId).MicrosoftGraphAgreementFileLocalization(microsoftGraphAgreementFileLocalization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileCreateLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFileCreateLocalizations`: MicrosoftGraphAgreementFileLocalization
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileApi.AgreementsFileCreateLocalizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFileCreateLocalizationsRequest struct via the builder pattern


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


## AgreementsFileDeleteLocalizations

> AgreementsFileDeleteLocalizations(ctx, agreementId, agreementFileLocalizationId).IfMatch(ifMatch).Execute()

Delete navigation property localizations for agreements

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileDeleteLocalizations(context.Background(), agreementId, agreementFileLocalizationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileDeleteLocalizations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAgreementsFileDeleteLocalizationsRequest struct via the builder pattern


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


## AgreementsFileGetLocalizations

> MicrosoftGraphAgreementFileLocalization AgreementsFileGetLocalizations(ctx, agreementId, agreementFileLocalizationId).Select_(select_).Expand(expand).Execute()

Get localizations from agreements

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileGetLocalizations(context.Background(), agreementId, agreementFileLocalizationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileGetLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFileGetLocalizations`: MicrosoftGraphAgreementFileLocalization
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileApi.AgreementsFileGetLocalizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFileGetLocalizationsRequest struct via the builder pattern


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


## AgreementsFileListLocalizations

> CollectionOfAgreementFileLocalization AgreementsFileListLocalizations(ctx, agreementId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get localizations from agreements

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileListLocalizations(context.Background(), agreementId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileListLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFileListLocalizations`: CollectionOfAgreementFileLocalization
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileApi.AgreementsFileListLocalizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFileListLocalizationsRequest struct via the builder pattern


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


## AgreementsFileLocalizationsCreateVersions

> MicrosoftGraphAgreementFileVersion AgreementsFileLocalizationsCreateVersions(ctx, agreementId, agreementFileLocalizationId).MicrosoftGraphAgreementFileVersion(microsoftGraphAgreementFileVersion).Execute()

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileLocalizationsCreateVersions(context.Background(), agreementId, agreementFileLocalizationId).MicrosoftGraphAgreementFileVersion(microsoftGraphAgreementFileVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileLocalizationsCreateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFileLocalizationsCreateVersions`: MicrosoftGraphAgreementFileVersion
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileApi.AgreementsFileLocalizationsCreateVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFileLocalizationsCreateVersionsRequest struct via the builder pattern


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


## AgreementsFileLocalizationsDeleteVersions

> AgreementsFileLocalizationsDeleteVersions(ctx, agreementId, agreementFileLocalizationId, agreementFileVersionId).IfMatch(ifMatch).Execute()

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileLocalizationsDeleteVersions(context.Background(), agreementId, agreementFileLocalizationId, agreementFileVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileLocalizationsDeleteVersions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAgreementsFileLocalizationsDeleteVersionsRequest struct via the builder pattern


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


## AgreementsFileLocalizationsGetVersions

> MicrosoftGraphAgreementFileVersion AgreementsFileLocalizationsGetVersions(ctx, agreementId, agreementFileLocalizationId, agreementFileVersionId).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileLocalizationsGetVersions(context.Background(), agreementId, agreementFileLocalizationId, agreementFileVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileLocalizationsGetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFileLocalizationsGetVersions`: MicrosoftGraphAgreementFileVersion
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileApi.AgreementsFileLocalizationsGetVersions`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAgreementsFileLocalizationsGetVersionsRequest struct via the builder pattern


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


## AgreementsFileLocalizationsListVersions

> CollectionOfAgreementFileVersion AgreementsFileLocalizationsListVersions(ctx, agreementId, agreementFileLocalizationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileLocalizationsListVersions(context.Background(), agreementId, agreementFileLocalizationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileLocalizationsListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsFileLocalizationsListVersions`: CollectionOfAgreementFileVersion
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileApi.AgreementsFileLocalizationsListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 
**agreementFileLocalizationId** | **string** | key: id of agreementFileLocalization | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsFileLocalizationsListVersionsRequest struct via the builder pattern


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


## AgreementsFileLocalizationsUpdateVersions

> AgreementsFileLocalizationsUpdateVersions(ctx, agreementId, agreementFileLocalizationId, agreementFileVersionId).MicrosoftGraphAgreementFileVersion(microsoftGraphAgreementFileVersion).Execute()

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileLocalizationsUpdateVersions(context.Background(), agreementId, agreementFileLocalizationId, agreementFileVersionId).MicrosoftGraphAgreementFileVersion(microsoftGraphAgreementFileVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileLocalizationsUpdateVersions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAgreementsFileLocalizationsUpdateVersionsRequest struct via the builder pattern


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


## AgreementsFileUpdateLocalizations

> AgreementsFileUpdateLocalizations(ctx, agreementId, agreementFileLocalizationId).MicrosoftGraphAgreementFileLocalization(microsoftGraphAgreementFileLocalization).Execute()

Update the navigation property localizations in agreements

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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsFileUpdateLocalizations(context.Background(), agreementId, agreementFileLocalizationId).MicrosoftGraphAgreementFileLocalization(microsoftGraphAgreementFileLocalization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsFileUpdateLocalizations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAgreementsFileUpdateLocalizationsRequest struct via the builder pattern


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


## AgreementsGetFile

> MicrosoftGraphAgreementFile AgreementsGetFile(ctx, agreementId).Select_(select_).Expand(expand).Execute()

Get file from agreements



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
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsGetFile(context.Background(), agreementId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsGetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AgreementsGetFile`: MicrosoftGraphAgreementFile
    fmt.Fprintf(os.Stdout, "Response from `AgreementsAgreementFileApi.AgreementsGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agreementId** | **string** | key: id of agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgreementsGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAgreementFile**](MicrosoftGraphAgreementFile.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgreementsUpdateFile

> AgreementsUpdateFile(ctx, agreementId).MicrosoftGraphAgreementFile(microsoftGraphAgreementFile).Execute()

Update the navigation property file in agreements



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
    microsoftGraphAgreementFile := *openapiclient.NewMicrosoftGraphAgreementFile() // MicrosoftGraphAgreementFile | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AgreementsAgreementFileApi.AgreementsUpdateFile(context.Background(), agreementId).MicrosoftGraphAgreementFile(microsoftGraphAgreementFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementsAgreementFileApi.AgreementsUpdateFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAgreementsUpdateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAgreementFile** | [**MicrosoftGraphAgreementFile**](MicrosoftGraphAgreementFile.md) | New navigation property values | 

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

