# \SitesColumnDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesColumnsDeleteRefSourceColumn**](SitesColumnDefinitionApi.md#SitesColumnsDeleteRefSourceColumn) | **Delete** /sites/{site-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Delete ref of navigation property sourceColumn for sites
[**SitesColumnsGetRefSourceColumn**](SitesColumnDefinitionApi.md#SitesColumnsGetRefSourceColumn) | **Get** /sites/{site-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Get ref of sourceColumn from sites
[**SitesColumnsGetSourceColumn**](SitesColumnDefinitionApi.md#SitesColumnsGetSourceColumn) | **Get** /sites/{site-id}/columns/{columnDefinition-id}/sourceColumn | Get sourceColumn from sites
[**SitesColumnsUpdateRefSourceColumn**](SitesColumnDefinitionApi.md#SitesColumnsUpdateRefSourceColumn) | **Put** /sites/{site-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Update the ref of navigation property sourceColumn in sites
[**SitesCreateColumns**](SitesColumnDefinitionApi.md#SitesCreateColumns) | **Post** /sites/{site-id}/columns | Create new navigation property to columns for sites
[**SitesCreateRefExternalColumns**](SitesColumnDefinitionApi.md#SitesCreateRefExternalColumns) | **Post** /sites/{site-id}/externalColumns/$ref | Create new navigation property ref to externalColumns for sites
[**SitesDeleteColumns**](SitesColumnDefinitionApi.md#SitesDeleteColumns) | **Delete** /sites/{site-id}/columns/{columnDefinition-id} | Delete navigation property columns for sites
[**SitesGetColumns**](SitesColumnDefinitionApi.md#SitesGetColumns) | **Get** /sites/{site-id}/columns/{columnDefinition-id} | Get columns from sites
[**SitesListColumns**](SitesColumnDefinitionApi.md#SitesListColumns) | **Get** /sites/{site-id}/columns | Get columns from sites
[**SitesListExternalColumns**](SitesColumnDefinitionApi.md#SitesListExternalColumns) | **Get** /sites/{site-id}/externalColumns | Get externalColumns from sites
[**SitesListRefExternalColumns**](SitesColumnDefinitionApi.md#SitesListRefExternalColumns) | **Get** /sites/{site-id}/externalColumns/$ref | Get ref of externalColumns from sites
[**SitesUpdateColumns**](SitesColumnDefinitionApi.md#SitesUpdateColumns) | **Patch** /sites/{site-id}/columns/{columnDefinition-id} | Update the navigation property columns in sites



## SitesColumnsDeleteRefSourceColumn

> SitesColumnsDeleteRefSourceColumn(ctx, siteId, columnDefinitionId).IfMatch(ifMatch).Execute()

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
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesColumnsDeleteRefSourceColumn(context.Background(), siteId, columnDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesColumnsDeleteRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesColumnsDeleteRefSourceColumnRequest struct via the builder pattern


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


## SitesColumnsGetRefSourceColumn

> string SitesColumnsGetRefSourceColumn(ctx, siteId, columnDefinitionId).Execute()

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
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesColumnsGetRefSourceColumn(context.Background(), siteId, columnDefinitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesColumnsGetRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesColumnsGetRefSourceColumn`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesColumnDefinitionApi.SitesColumnsGetRefSourceColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesColumnsGetRefSourceColumnRequest struct via the builder pattern


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


## SitesColumnsGetSourceColumn

> MicrosoftGraphColumnDefinition SitesColumnsGetSourceColumn(ctx, siteId, columnDefinitionId).Select_(select_).Expand(expand).Execute()

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
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesColumnsGetSourceColumn(context.Background(), siteId, columnDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesColumnsGetSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesColumnsGetSourceColumn`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesColumnDefinitionApi.SitesColumnsGetSourceColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesColumnsGetSourceColumnRequest struct via the builder pattern


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


## SitesColumnsUpdateRefSourceColumn

> SitesColumnsUpdateRefSourceColumn(ctx, siteId, columnDefinitionId).RequestBody(requestBody).Execute()

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
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesColumnsUpdateRefSourceColumn(context.Background(), siteId, columnDefinitionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesColumnsUpdateRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesColumnsUpdateRefSourceColumnRequest struct via the builder pattern


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


## SitesCreateColumns

> MicrosoftGraphColumnDefinition SitesCreateColumns(ctx, siteId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()

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
    microsoftGraphColumnDefinition := *openapiclient.NewMicrosoftGraphColumnDefinition() // MicrosoftGraphColumnDefinition | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesCreateColumns(context.Background(), siteId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesCreateColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesCreateColumns`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesColumnDefinitionApi.SitesCreateColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesCreateColumnsRequest struct via the builder pattern


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


## SitesCreateRefExternalColumns

> map[string]interface{} SitesCreateRefExternalColumns(ctx, siteId).RequestBody(requestBody).Execute()

Create new navigation property ref to externalColumns for sites



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesCreateRefExternalColumns(context.Background(), siteId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesCreateRefExternalColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesCreateRefExternalColumns`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SitesColumnDefinitionApi.SitesCreateRefExternalColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesCreateRefExternalColumnsRequest struct via the builder pattern


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


## SitesDeleteColumns

> SitesDeleteColumns(ctx, siteId, columnDefinitionId).IfMatch(ifMatch).Execute()

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
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesDeleteColumns(context.Background(), siteId, columnDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesDeleteColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesDeleteColumnsRequest struct via the builder pattern


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


## SitesGetColumns

> MicrosoftGraphColumnDefinition SitesGetColumns(ctx, siteId, columnDefinitionId).Select_(select_).Expand(expand).Execute()

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
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesGetColumns(context.Background(), siteId, columnDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesGetColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetColumns`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesColumnDefinitionApi.SitesGetColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetColumnsRequest struct via the builder pattern


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


## SitesListColumns

> CollectionOfColumnDefinition SitesListColumns(ctx, siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesListColumns(context.Background(), siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesListColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListColumns`: CollectionOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesColumnDefinitionApi.SitesListColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListColumnsRequest struct via the builder pattern


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


## SitesListExternalColumns

> CollectionOfColumnDefinition SitesListExternalColumns(ctx, siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get externalColumns from sites



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
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesListExternalColumns(context.Background(), siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesListExternalColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListExternalColumns`: CollectionOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesColumnDefinitionApi.SitesListExternalColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListExternalColumnsRequest struct via the builder pattern


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


## SitesListRefExternalColumns

> CollectionOfLinksOfColumnDefinition SitesListRefExternalColumns(ctx, siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of externalColumns from sites



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesListRefExternalColumns(context.Background(), siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesListRefExternalColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListRefExternalColumns`: CollectionOfLinksOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesColumnDefinitionApi.SitesListRefExternalColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListRefExternalColumnsRequest struct via the builder pattern


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


## SitesUpdateColumns

> SitesUpdateColumns(ctx, siteId, columnDefinitionId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()

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
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    microsoftGraphColumnDefinition := *openapiclient.NewMicrosoftGraphColumnDefinition() // MicrosoftGraphColumnDefinition | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesColumnDefinitionApi.SitesUpdateColumns(context.Background(), siteId, columnDefinitionId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesColumnDefinitionApi.SitesUpdateColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesUpdateColumnsRequest struct via the builder pattern


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

