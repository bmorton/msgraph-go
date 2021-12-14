# \SitesContentTypeApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesContentTypesColumnsDeleteRefSourceColumn**](SitesContentTypeApi.md#SitesContentTypesColumnsDeleteRefSourceColumn) | **Delete** /sites/{site-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Delete ref of navigation property sourceColumn for sites
[**SitesContentTypesColumnsGetRefSourceColumn**](SitesContentTypeApi.md#SitesContentTypesColumnsGetRefSourceColumn) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Get ref of sourceColumn from sites
[**SitesContentTypesColumnsGetSourceColumn**](SitesContentTypeApi.md#SitesContentTypesColumnsGetSourceColumn) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id}/sourceColumn | Get sourceColumn from sites
[**SitesContentTypesColumnsUpdateRefSourceColumn**](SitesContentTypeApi.md#SitesContentTypesColumnsUpdateRefSourceColumn) | **Put** /sites/{site-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Update the ref of navigation property sourceColumn in sites
[**SitesContentTypesCreateColumnLinks**](SitesContentTypeApi.md#SitesContentTypesCreateColumnLinks) | **Post** /sites/{site-id}/contentTypes/{contentType-id}/columnLinks | Create new navigation property to columnLinks for sites
[**SitesContentTypesCreateColumns**](SitesContentTypeApi.md#SitesContentTypesCreateColumns) | **Post** /sites/{site-id}/contentTypes/{contentType-id}/columns | Create new navigation property to columns for sites
[**SitesContentTypesCreateRefBaseTypes**](SitesContentTypeApi.md#SitesContentTypesCreateRefBaseTypes) | **Post** /sites/{site-id}/contentTypes/{contentType-id}/baseTypes/$ref | Create new navigation property ref to baseTypes for sites
[**SitesContentTypesCreateRefColumnPositions**](SitesContentTypeApi.md#SitesContentTypesCreateRefColumnPositions) | **Post** /sites/{site-id}/contentTypes/{contentType-id}/columnPositions/$ref | Create new navigation property ref to columnPositions for sites
[**SitesContentTypesDeleteColumnLinks**](SitesContentTypeApi.md#SitesContentTypesDeleteColumnLinks) | **Delete** /sites/{site-id}/contentTypes/{contentType-id}/columnLinks/{columnLink-id} | Delete navigation property columnLinks for sites
[**SitesContentTypesDeleteColumns**](SitesContentTypeApi.md#SitesContentTypesDeleteColumns) | **Delete** /sites/{site-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id} | Delete navigation property columns for sites
[**SitesContentTypesDeleteRefBase**](SitesContentTypeApi.md#SitesContentTypesDeleteRefBase) | **Delete** /sites/{site-id}/contentTypes/{contentType-id}/base/$ref | Delete ref of navigation property base for sites
[**SitesContentTypesGetBase**](SitesContentTypeApi.md#SitesContentTypesGetBase) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/base | Get base from sites
[**SitesContentTypesGetColumnLinks**](SitesContentTypeApi.md#SitesContentTypesGetColumnLinks) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/columnLinks/{columnLink-id} | Get columnLinks from sites
[**SitesContentTypesGetColumns**](SitesContentTypeApi.md#SitesContentTypesGetColumns) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id} | Get columns from sites
[**SitesContentTypesGetRefBase**](SitesContentTypeApi.md#SitesContentTypesGetRefBase) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/base/$ref | Get ref of base from sites
[**SitesContentTypesListBaseTypes**](SitesContentTypeApi.md#SitesContentTypesListBaseTypes) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/baseTypes | Get baseTypes from sites
[**SitesContentTypesListColumnLinks**](SitesContentTypeApi.md#SitesContentTypesListColumnLinks) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/columnLinks | Get columnLinks from sites
[**SitesContentTypesListColumnPositions**](SitesContentTypeApi.md#SitesContentTypesListColumnPositions) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/columnPositions | Get columnPositions from sites
[**SitesContentTypesListColumns**](SitesContentTypeApi.md#SitesContentTypesListColumns) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/columns | Get columns from sites
[**SitesContentTypesListRefBaseTypes**](SitesContentTypeApi.md#SitesContentTypesListRefBaseTypes) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/baseTypes/$ref | Get ref of baseTypes from sites
[**SitesContentTypesListRefColumnPositions**](SitesContentTypeApi.md#SitesContentTypesListRefColumnPositions) | **Get** /sites/{site-id}/contentTypes/{contentType-id}/columnPositions/$ref | Get ref of columnPositions from sites
[**SitesContentTypesUpdateColumnLinks**](SitesContentTypeApi.md#SitesContentTypesUpdateColumnLinks) | **Patch** /sites/{site-id}/contentTypes/{contentType-id}/columnLinks/{columnLink-id} | Update the navigation property columnLinks in sites
[**SitesContentTypesUpdateColumns**](SitesContentTypeApi.md#SitesContentTypesUpdateColumns) | **Patch** /sites/{site-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id} | Update the navigation property columns in sites
[**SitesContentTypesUpdateRefBase**](SitesContentTypeApi.md#SitesContentTypesUpdateRefBase) | **Put** /sites/{site-id}/contentTypes/{contentType-id}/base/$ref | Update the ref of navigation property base in sites
[**SitesCreateContentTypes**](SitesContentTypeApi.md#SitesCreateContentTypes) | **Post** /sites/{site-id}/contentTypes | Create new navigation property to contentTypes for sites
[**SitesDeleteContentTypes**](SitesContentTypeApi.md#SitesDeleteContentTypes) | **Delete** /sites/{site-id}/contentTypes/{contentType-id} | Delete navigation property contentTypes for sites
[**SitesGetContentTypes**](SitesContentTypeApi.md#SitesGetContentTypes) | **Get** /sites/{site-id}/contentTypes/{contentType-id} | Get contentTypes from sites
[**SitesListContentTypes**](SitesContentTypeApi.md#SitesListContentTypes) | **Get** /sites/{site-id}/contentTypes | Get contentTypes from sites
[**SitesUpdateContentTypes**](SitesContentTypeApi.md#SitesUpdateContentTypes) | **Patch** /sites/{site-id}/contentTypes/{contentType-id} | Update the navigation property contentTypes in sites



## SitesContentTypesColumnsDeleteRefSourceColumn

> SitesContentTypesColumnsDeleteRefSourceColumn(ctx, siteId, contentTypeId, columnDefinitionId).IfMatch(ifMatch).Execute()

Delete ref of navigation property sourceColumn for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesColumnsDeleteRefSourceColumn(context.Background(), siteId, contentTypeId, columnDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesColumnsDeleteRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesColumnsDeleteRefSourceColumnRequest struct via the builder pattern


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


## SitesContentTypesColumnsGetRefSourceColumn

> string SitesContentTypesColumnsGetRefSourceColumn(ctx, siteId, contentTypeId, columnDefinitionId).Execute()

Get ref of sourceColumn from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesColumnsGetRefSourceColumn(context.Background(), siteId, contentTypeId, columnDefinitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesColumnsGetRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesColumnsGetRefSourceColumn`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesColumnsGetRefSourceColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesColumnsGetRefSourceColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesColumnsGetSourceColumn

> MicrosoftGraphColumnDefinition SitesContentTypesColumnsGetSourceColumn(ctx, siteId, contentTypeId, columnDefinitionId).Select_(select_).Expand(expand).Execute()

Get sourceColumn from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesColumnsGetSourceColumn(context.Background(), siteId, contentTypeId, columnDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesColumnsGetSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesColumnsGetSourceColumn`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesColumnsGetSourceColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesColumnsGetSourceColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesColumnsUpdateRefSourceColumn

> SitesContentTypesColumnsUpdateRefSourceColumn(ctx, siteId, contentTypeId, columnDefinitionId).RequestBody(requestBody).Execute()

Update the ref of navigation property sourceColumn in sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesColumnsUpdateRefSourceColumn(context.Background(), siteId, contentTypeId, columnDefinitionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesColumnsUpdateRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesColumnsUpdateRefSourceColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

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


## SitesContentTypesCreateColumnLinks

> MicrosoftGraphColumnLink SitesContentTypesCreateColumnLinks(ctx, siteId, contentTypeId).MicrosoftGraphColumnLink(microsoftGraphColumnLink).Execute()

Create new navigation property to columnLinks for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    microsoftGraphColumnLink := *openapiclient.NewMicrosoftGraphColumnLink() // MicrosoftGraphColumnLink | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesCreateColumnLinks(context.Background(), siteId, contentTypeId).MicrosoftGraphColumnLink(microsoftGraphColumnLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesCreateColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesCreateColumnLinks`: MicrosoftGraphColumnLink
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesCreateColumnLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesCreateColumnLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphColumnLink** | [**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md) | New navigation property | 

### Return type

[**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesCreateColumns

> MicrosoftGraphColumnDefinition SitesContentTypesCreateColumns(ctx, siteId, contentTypeId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()

Create new navigation property to columns for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    microsoftGraphColumnDefinition := *openapiclient.NewMicrosoftGraphColumnDefinition() // MicrosoftGraphColumnDefinition | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesCreateColumns(context.Background(), siteId, contentTypeId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesCreateColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesCreateColumns`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesCreateColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesCreateColumnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) | New navigation property | 

### Return type

[**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesCreateRefBaseTypes

> map[string]interface{} SitesContentTypesCreateRefBaseTypes(ctx, siteId, contentTypeId).RequestBody(requestBody).Execute()

Create new navigation property ref to baseTypes for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesCreateRefBaseTypes(context.Background(), siteId, contentTypeId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesCreateRefBaseTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesCreateRefBaseTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesCreateRefBaseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesCreateRefBaseTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesCreateRefColumnPositions

> map[string]interface{} SitesContentTypesCreateRefColumnPositions(ctx, siteId, contentTypeId).RequestBody(requestBody).Execute()

Create new navigation property ref to columnPositions for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesCreateRefColumnPositions(context.Background(), siteId, contentTypeId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesCreateRefColumnPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesCreateRefColumnPositions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesCreateRefColumnPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesCreateRefColumnPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesDeleteColumnLinks

> SitesContentTypesDeleteColumnLinks(ctx, siteId, contentTypeId, columnLinkId).IfMatch(ifMatch).Execute()

Delete navigation property columnLinks for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnLinkId := "columnLinkId_example" // string | key: id of columnLink
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesDeleteColumnLinks(context.Background(), siteId, contentTypeId, columnLinkId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesDeleteColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnLinkId** | **string** | key: id of columnLink | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesDeleteColumnLinksRequest struct via the builder pattern


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


## SitesContentTypesDeleteColumns

> SitesContentTypesDeleteColumns(ctx, siteId, contentTypeId, columnDefinitionId).IfMatch(ifMatch).Execute()

Delete navigation property columns for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesDeleteColumns(context.Background(), siteId, contentTypeId, columnDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesDeleteColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesDeleteColumnsRequest struct via the builder pattern


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


## SitesContentTypesDeleteRefBase

> SitesContentTypesDeleteRefBase(ctx, siteId, contentTypeId).IfMatch(ifMatch).Execute()

Delete ref of navigation property base for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesDeleteRefBase(context.Background(), siteId, contentTypeId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesDeleteRefBase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesDeleteRefBaseRequest struct via the builder pattern


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


## SitesContentTypesGetBase

> MicrosoftGraphContentType SitesContentTypesGetBase(ctx, siteId, contentTypeId).Select_(select_).Expand(expand).Execute()

Get base from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesGetBase(context.Background(), siteId, contentTypeId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesGetBase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesGetBase`: MicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesGetBase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesGetBaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesGetColumnLinks

> MicrosoftGraphColumnLink SitesContentTypesGetColumnLinks(ctx, siteId, contentTypeId, columnLinkId).Select_(select_).Expand(expand).Execute()

Get columnLinks from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnLinkId := "columnLinkId_example" // string | key: id of columnLink
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesGetColumnLinks(context.Background(), siteId, contentTypeId, columnLinkId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesGetColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesGetColumnLinks`: MicrosoftGraphColumnLink
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesGetColumnLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnLinkId** | **string** | key: id of columnLink | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesGetColumnLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesGetColumns

> MicrosoftGraphColumnDefinition SitesContentTypesGetColumns(ctx, siteId, contentTypeId, columnDefinitionId).Select_(select_).Expand(expand).Execute()

Get columns from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesGetColumns(context.Background(), siteId, contentTypeId, columnDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesGetColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesGetColumns`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesGetColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesGetColumnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesGetRefBase

> string SitesContentTypesGetRefBase(ctx, siteId, contentTypeId).Execute()

Get ref of base from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesGetRefBase(context.Background(), siteId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesGetRefBase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesGetRefBase`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesGetRefBase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesGetRefBaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesListBaseTypes

> CollectionOfContentType SitesContentTypesListBaseTypes(ctx, siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get baseTypes from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
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
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesListBaseTypes(context.Background(), siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesListBaseTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesListBaseTypes`: CollectionOfContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesListBaseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesListBaseTypesRequest struct via the builder pattern


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

[**CollectionOfContentType**](CollectionOfContentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesListColumnLinks

> CollectionOfColumnLink SitesContentTypesListColumnLinks(ctx, siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get columnLinks from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
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
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesListColumnLinks(context.Background(), siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesListColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesListColumnLinks`: CollectionOfColumnLink
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesListColumnLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesListColumnLinksRequest struct via the builder pattern


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

[**CollectionOfColumnLink**](CollectionOfColumnLink.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesListColumnPositions

> CollectionOfColumnDefinition SitesContentTypesListColumnPositions(ctx, siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get columnPositions from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
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
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesListColumnPositions(context.Background(), siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesListColumnPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesListColumnPositions`: CollectionOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesListColumnPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesListColumnPositionsRequest struct via the builder pattern


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

[**CollectionOfColumnDefinition**](CollectionOfColumnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesListColumns

> CollectionOfColumnDefinition SitesContentTypesListColumns(ctx, siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get columns from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
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
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesListColumns(context.Background(), siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesListColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesListColumns`: CollectionOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesListColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesListColumnsRequest struct via the builder pattern


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

[**CollectionOfColumnDefinition**](CollectionOfColumnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesListRefBaseTypes

> CollectionOfLinksOfContentType SitesContentTypesListRefBaseTypes(ctx, siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of baseTypes from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesListRefBaseTypes(context.Background(), siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesListRefBaseTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesListRefBaseTypes`: CollectionOfLinksOfContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesListRefBaseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesListRefBaseTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfContentType**](CollectionOfLinksOfContentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesListRefColumnPositions

> CollectionOfLinksOfColumnDefinition SitesContentTypesListRefColumnPositions(ctx, siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of columnPositions from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesListRefColumnPositions(context.Background(), siteId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesListRefColumnPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesContentTypesListRefColumnPositions`: CollectionOfLinksOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesContentTypesListRefColumnPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesListRefColumnPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfColumnDefinition**](CollectionOfLinksOfColumnDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesContentTypesUpdateColumnLinks

> SitesContentTypesUpdateColumnLinks(ctx, siteId, contentTypeId, columnLinkId).MicrosoftGraphColumnLink(microsoftGraphColumnLink).Execute()

Update the navigation property columnLinks in sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnLinkId := "columnLinkId_example" // string | key: id of columnLink
    microsoftGraphColumnLink := *openapiclient.NewMicrosoftGraphColumnLink() // MicrosoftGraphColumnLink | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesUpdateColumnLinks(context.Background(), siteId, contentTypeId, columnLinkId).MicrosoftGraphColumnLink(microsoftGraphColumnLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesUpdateColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnLinkId** | **string** | key: id of columnLink | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesUpdateColumnLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphColumnLink** | [**MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md) | New navigation property values | 

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


## SitesContentTypesUpdateColumns

> SitesContentTypesUpdateColumns(ctx, siteId, contentTypeId, columnDefinitionId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()

Update the navigation property columns in sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    microsoftGraphColumnDefinition := *openapiclient.NewMicrosoftGraphColumnDefinition() // MicrosoftGraphColumnDefinition | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesUpdateColumns(context.Background(), siteId, contentTypeId, columnDefinitionId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesUpdateColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesUpdateColumnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphColumnDefinition** | [**MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) | New navigation property values | 

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


## SitesContentTypesUpdateRefBase

> SitesContentTypesUpdateRefBase(ctx, siteId, contentTypeId).RequestBody(requestBody).Execute()

Update the ref of navigation property base in sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesContentTypesUpdateRefBase(context.Background(), siteId, contentTypeId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesContentTypesUpdateRefBase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesContentTypesUpdateRefBaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

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


## SitesCreateContentTypes

> MicrosoftGraphContentType SitesCreateContentTypes(ctx, siteId).MicrosoftGraphContentType(microsoftGraphContentType).Execute()

Create new navigation property to contentTypes for sites



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
    siteId := "siteId_example" // string | key: id of site
    microsoftGraphContentType := *openapiclient.NewMicrosoftGraphContentType() // MicrosoftGraphContentType | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesCreateContentTypes(context.Background(), siteId).MicrosoftGraphContentType(microsoftGraphContentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesCreateContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesCreateContentTypes`: MicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesCreateContentTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesCreateContentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphContentType** | [**MicrosoftGraphContentType**](MicrosoftGraphContentType.md) | New navigation property | 

### Return type

[**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesDeleteContentTypes

> SitesDeleteContentTypes(ctx, siteId, contentTypeId).IfMatch(ifMatch).Execute()

Delete navigation property contentTypes for sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesDeleteContentTypes(context.Background(), siteId, contentTypeId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesDeleteContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesDeleteContentTypesRequest struct via the builder pattern


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


## SitesGetContentTypes

> MicrosoftGraphContentType SitesGetContentTypes(ctx, siteId, contentTypeId).Select_(select_).Expand(expand).Execute()

Get contentTypes from sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesGetContentTypes(context.Background(), siteId, contentTypeId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesGetContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetContentTypes`: MicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesGetContentTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetContentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphContentType**](MicrosoftGraphContentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListContentTypes

> CollectionOfContentType SitesListContentTypes(ctx, siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get contentTypes from sites



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
    siteId := "siteId_example" // string | key: id of site
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
    resp, r, err := api_client.SitesContentTypeApi.SitesListContentTypes(context.Background(), siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesListContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListContentTypes`: CollectionOfContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesContentTypeApi.SitesListContentTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListContentTypesRequest struct via the builder pattern


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

[**CollectionOfContentType**](CollectionOfContentType.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesUpdateContentTypes

> SitesUpdateContentTypes(ctx, siteId, contentTypeId).MicrosoftGraphContentType(microsoftGraphContentType).Execute()

Update the navigation property contentTypes in sites



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
    siteId := "siteId_example" // string | key: id of site
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    microsoftGraphContentType := *openapiclient.NewMicrosoftGraphContentType() // MicrosoftGraphContentType | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesContentTypeApi.SitesUpdateContentTypes(context.Background(), siteId, contentTypeId).MicrosoftGraphContentType(microsoftGraphContentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesContentTypeApi.SitesUpdateContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesUpdateContentTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphContentType** | [**MicrosoftGraphContentType**](MicrosoftGraphContentType.md) | New navigation property values | 

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

