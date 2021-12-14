# \SitesListApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesCreateLists**](SitesListApi.md#SitesCreateLists) | **Post** /sites/{site-id}/lists | Create new navigation property to lists for sites
[**SitesDeleteLists**](SitesListApi.md#SitesDeleteLists) | **Delete** /sites/{site-id}/lists/{list-id} | Delete navigation property lists for sites
[**SitesGetLists**](SitesListApi.md#SitesGetLists) | **Get** /sites/{site-id}/lists/{list-id} | Get lists from sites
[**SitesListLists**](SitesListApi.md#SitesListLists) | **Get** /sites/{site-id}/lists | Get lists from sites
[**SitesListsColumnsDeleteRefSourceColumn**](SitesListApi.md#SitesListsColumnsDeleteRefSourceColumn) | **Delete** /sites/{site-id}/lists/{list-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Delete ref of navigation property sourceColumn for sites
[**SitesListsColumnsGetRefSourceColumn**](SitesListApi.md#SitesListsColumnsGetRefSourceColumn) | **Get** /sites/{site-id}/lists/{list-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Get ref of sourceColumn from sites
[**SitesListsColumnsGetSourceColumn**](SitesListApi.md#SitesListsColumnsGetSourceColumn) | **Get** /sites/{site-id}/lists/{list-id}/columns/{columnDefinition-id}/sourceColumn | Get sourceColumn from sites
[**SitesListsColumnsUpdateRefSourceColumn**](SitesListApi.md#SitesListsColumnsUpdateRefSourceColumn) | **Put** /sites/{site-id}/lists/{list-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Update the ref of navigation property sourceColumn in sites
[**SitesListsContentTypesColumnsDeleteRefSourceColumn**](SitesListApi.md#SitesListsContentTypesColumnsDeleteRefSourceColumn) | **Delete** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Delete ref of navigation property sourceColumn for sites
[**SitesListsContentTypesColumnsGetRefSourceColumn**](SitesListApi.md#SitesListsContentTypesColumnsGetRefSourceColumn) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Get ref of sourceColumn from sites
[**SitesListsContentTypesColumnsGetSourceColumn**](SitesListApi.md#SitesListsContentTypesColumnsGetSourceColumn) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id}/sourceColumn | Get sourceColumn from sites
[**SitesListsContentTypesColumnsUpdateRefSourceColumn**](SitesListApi.md#SitesListsContentTypesColumnsUpdateRefSourceColumn) | **Put** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id}/sourceColumn/$ref | Update the ref of navigation property sourceColumn in sites
[**SitesListsContentTypesCreateColumnLinks**](SitesListApi.md#SitesListsContentTypesCreateColumnLinks) | **Post** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columnLinks | Create new navigation property to columnLinks for sites
[**SitesListsContentTypesCreateColumns**](SitesListApi.md#SitesListsContentTypesCreateColumns) | **Post** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns | Create new navigation property to columns for sites
[**SitesListsContentTypesCreateRefBaseTypes**](SitesListApi.md#SitesListsContentTypesCreateRefBaseTypes) | **Post** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/baseTypes/$ref | Create new navigation property ref to baseTypes for sites
[**SitesListsContentTypesCreateRefColumnPositions**](SitesListApi.md#SitesListsContentTypesCreateRefColumnPositions) | **Post** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columnPositions/$ref | Create new navigation property ref to columnPositions for sites
[**SitesListsContentTypesDeleteColumnLinks**](SitesListApi.md#SitesListsContentTypesDeleteColumnLinks) | **Delete** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columnLinks/{columnLink-id} | Delete navigation property columnLinks for sites
[**SitesListsContentTypesDeleteColumns**](SitesListApi.md#SitesListsContentTypesDeleteColumns) | **Delete** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id} | Delete navigation property columns for sites
[**SitesListsContentTypesDeleteRefBase**](SitesListApi.md#SitesListsContentTypesDeleteRefBase) | **Delete** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/base/$ref | Delete ref of navigation property base for sites
[**SitesListsContentTypesGetBase**](SitesListApi.md#SitesListsContentTypesGetBase) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/base | Get base from sites
[**SitesListsContentTypesGetColumnLinks**](SitesListApi.md#SitesListsContentTypesGetColumnLinks) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columnLinks/{columnLink-id} | Get columnLinks from sites
[**SitesListsContentTypesGetColumns**](SitesListApi.md#SitesListsContentTypesGetColumns) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id} | Get columns from sites
[**SitesListsContentTypesGetRefBase**](SitesListApi.md#SitesListsContentTypesGetRefBase) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/base/$ref | Get ref of base from sites
[**SitesListsContentTypesListBaseTypes**](SitesListApi.md#SitesListsContentTypesListBaseTypes) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/baseTypes | Get baseTypes from sites
[**SitesListsContentTypesListColumnLinks**](SitesListApi.md#SitesListsContentTypesListColumnLinks) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columnLinks | Get columnLinks from sites
[**SitesListsContentTypesListColumnPositions**](SitesListApi.md#SitesListsContentTypesListColumnPositions) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columnPositions | Get columnPositions from sites
[**SitesListsContentTypesListColumns**](SitesListApi.md#SitesListsContentTypesListColumns) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns | Get columns from sites
[**SitesListsContentTypesListRefBaseTypes**](SitesListApi.md#SitesListsContentTypesListRefBaseTypes) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/baseTypes/$ref | Get ref of baseTypes from sites
[**SitesListsContentTypesListRefColumnPositions**](SitesListApi.md#SitesListsContentTypesListRefColumnPositions) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columnPositions/$ref | Get ref of columnPositions from sites
[**SitesListsContentTypesUpdateColumnLinks**](SitesListApi.md#SitesListsContentTypesUpdateColumnLinks) | **Patch** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columnLinks/{columnLink-id} | Update the navigation property columnLinks in sites
[**SitesListsContentTypesUpdateColumns**](SitesListApi.md#SitesListsContentTypesUpdateColumns) | **Patch** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/columns/{columnDefinition-id} | Update the navigation property columns in sites
[**SitesListsContentTypesUpdateRefBase**](SitesListApi.md#SitesListsContentTypesUpdateRefBase) | **Put** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id}/base/$ref | Update the ref of navigation property base in sites
[**SitesListsCreateColumns**](SitesListApi.md#SitesListsCreateColumns) | **Post** /sites/{site-id}/lists/{list-id}/columns | Create new navigation property to columns for sites
[**SitesListsCreateContentTypes**](SitesListApi.md#SitesListsCreateContentTypes) | **Post** /sites/{site-id}/lists/{list-id}/contentTypes | Create new navigation property to contentTypes for sites
[**SitesListsCreateItems**](SitesListApi.md#SitesListsCreateItems) | **Post** /sites/{site-id}/lists/{list-id}/items | Create new navigation property to items for sites
[**SitesListsCreateSubscriptions**](SitesListApi.md#SitesListsCreateSubscriptions) | **Post** /sites/{site-id}/lists/{list-id}/subscriptions | Create new navigation property to subscriptions for sites
[**SitesListsDeleteColumns**](SitesListApi.md#SitesListsDeleteColumns) | **Delete** /sites/{site-id}/lists/{list-id}/columns/{columnDefinition-id} | Delete navigation property columns for sites
[**SitesListsDeleteContentTypes**](SitesListApi.md#SitesListsDeleteContentTypes) | **Delete** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id} | Delete navigation property contentTypes for sites
[**SitesListsDeleteDrive**](SitesListApi.md#SitesListsDeleteDrive) | **Delete** /sites/{site-id}/lists/{list-id}/drive | Delete navigation property drive for sites
[**SitesListsDeleteItems**](SitesListApi.md#SitesListsDeleteItems) | **Delete** /sites/{site-id}/lists/{list-id}/items/{listItem-id} | Delete navigation property items for sites
[**SitesListsDeleteSubscriptions**](SitesListApi.md#SitesListsDeleteSubscriptions) | **Delete** /sites/{site-id}/lists/{list-id}/subscriptions/{subscription-id} | Delete navigation property subscriptions for sites
[**SitesListsGetColumns**](SitesListApi.md#SitesListsGetColumns) | **Get** /sites/{site-id}/lists/{list-id}/columns/{columnDefinition-id} | Get columns from sites
[**SitesListsGetContentTypes**](SitesListApi.md#SitesListsGetContentTypes) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id} | Get contentTypes from sites
[**SitesListsGetDrive**](SitesListApi.md#SitesListsGetDrive) | **Get** /sites/{site-id}/lists/{list-id}/drive | Get drive from sites
[**SitesListsGetItems**](SitesListApi.md#SitesListsGetItems) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id} | Get items from sites
[**SitesListsGetSubscriptions**](SitesListApi.md#SitesListsGetSubscriptions) | **Get** /sites/{site-id}/lists/{list-id}/subscriptions/{subscription-id} | Get subscriptions from sites
[**SitesListsItemsCreateVersions**](SitesListApi.md#SitesListsItemsCreateVersions) | **Post** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/versions | Create new navigation property to versions for sites
[**SitesListsItemsDeleteDriveItem**](SitesListApi.md#SitesListsItemsDeleteDriveItem) | **Delete** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/driveItem | Delete navigation property driveItem for sites
[**SitesListsItemsDeleteFields**](SitesListApi.md#SitesListsItemsDeleteFields) | **Delete** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/fields | Delete navigation property fields for sites
[**SitesListsItemsDeleteRefAnalytics**](SitesListApi.md#SitesListsItemsDeleteRefAnalytics) | **Delete** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/analytics/$ref | Delete ref of navigation property analytics for sites
[**SitesListsItemsDeleteVersions**](SitesListApi.md#SitesListsItemsDeleteVersions) | **Delete** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/versions/{listItemVersion-id} | Delete navigation property versions for sites
[**SitesListsItemsGetAnalytics**](SitesListApi.md#SitesListsItemsGetAnalytics) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/analytics | Get analytics from sites
[**SitesListsItemsGetDriveItem**](SitesListApi.md#SitesListsItemsGetDriveItem) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/driveItem | Get driveItem from sites
[**SitesListsItemsGetDriveItemContent**](SitesListApi.md#SitesListsItemsGetDriveItemContent) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/driveItem/content | Get media content for the navigation property driveItem from sites
[**SitesListsItemsGetFields**](SitesListApi.md#SitesListsItemsGetFields) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/fields | Get fields from sites
[**SitesListsItemsGetRefAnalytics**](SitesListApi.md#SitesListsItemsGetRefAnalytics) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/analytics/$ref | Get ref of analytics from sites
[**SitesListsItemsGetVersions**](SitesListApi.md#SitesListsItemsGetVersions) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/versions/{listItemVersion-id} | Get versions from sites
[**SitesListsItemsListVersions**](SitesListApi.md#SitesListsItemsListVersions) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/versions | Get versions from sites
[**SitesListsItemsUpdateDriveItem**](SitesListApi.md#SitesListsItemsUpdateDriveItem) | **Patch** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/driveItem | Update the navigation property driveItem in sites
[**SitesListsItemsUpdateDriveItemContent**](SitesListApi.md#SitesListsItemsUpdateDriveItemContent) | **Put** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/driveItem/content | Update media content for the navigation property driveItem in sites
[**SitesListsItemsUpdateFields**](SitesListApi.md#SitesListsItemsUpdateFields) | **Patch** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/fields | Update the navigation property fields in sites
[**SitesListsItemsUpdateRefAnalytics**](SitesListApi.md#SitesListsItemsUpdateRefAnalytics) | **Put** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/analytics/$ref | Update the ref of navigation property analytics in sites
[**SitesListsItemsUpdateVersions**](SitesListApi.md#SitesListsItemsUpdateVersions) | **Patch** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/versions/{listItemVersion-id} | Update the navigation property versions in sites
[**SitesListsItemsVersionsDeleteFields**](SitesListApi.md#SitesListsItemsVersionsDeleteFields) | **Delete** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/versions/{listItemVersion-id}/fields | Delete navigation property fields for sites
[**SitesListsItemsVersionsGetFields**](SitesListApi.md#SitesListsItemsVersionsGetFields) | **Get** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/versions/{listItemVersion-id}/fields | Get fields from sites
[**SitesListsItemsVersionsUpdateFields**](SitesListApi.md#SitesListsItemsVersionsUpdateFields) | **Patch** /sites/{site-id}/lists/{list-id}/items/{listItem-id}/versions/{listItemVersion-id}/fields | Update the navigation property fields in sites
[**SitesListsListColumns**](SitesListApi.md#SitesListsListColumns) | **Get** /sites/{site-id}/lists/{list-id}/columns | Get columns from sites
[**SitesListsListContentTypes**](SitesListApi.md#SitesListsListContentTypes) | **Get** /sites/{site-id}/lists/{list-id}/contentTypes | Get contentTypes from sites
[**SitesListsListItems**](SitesListApi.md#SitesListsListItems) | **Get** /sites/{site-id}/lists/{list-id}/items | Get items from sites
[**SitesListsListSubscriptions**](SitesListApi.md#SitesListsListSubscriptions) | **Get** /sites/{site-id}/lists/{list-id}/subscriptions | Get subscriptions from sites
[**SitesListsUpdateColumns**](SitesListApi.md#SitesListsUpdateColumns) | **Patch** /sites/{site-id}/lists/{list-id}/columns/{columnDefinition-id} | Update the navigation property columns in sites
[**SitesListsUpdateContentTypes**](SitesListApi.md#SitesListsUpdateContentTypes) | **Patch** /sites/{site-id}/lists/{list-id}/contentTypes/{contentType-id} | Update the navigation property contentTypes in sites
[**SitesListsUpdateDrive**](SitesListApi.md#SitesListsUpdateDrive) | **Patch** /sites/{site-id}/lists/{list-id}/drive | Update the navigation property drive in sites
[**SitesListsUpdateItems**](SitesListApi.md#SitesListsUpdateItems) | **Patch** /sites/{site-id}/lists/{list-id}/items/{listItem-id} | Update the navigation property items in sites
[**SitesListsUpdateSubscriptions**](SitesListApi.md#SitesListsUpdateSubscriptions) | **Patch** /sites/{site-id}/lists/{list-id}/subscriptions/{subscription-id} | Update the navigation property subscriptions in sites
[**SitesUpdateLists**](SitesListApi.md#SitesUpdateLists) | **Patch** /sites/{site-id}/lists/{list-id} | Update the navigation property lists in sites



## SitesCreateLists

> MicrosoftGraphList SitesCreateLists(ctx, siteId).MicrosoftGraphList(microsoftGraphList).Execute()

Create new navigation property to lists for sites



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
    microsoftGraphList := *openapiclient.NewMicrosoftGraphList() // MicrosoftGraphList | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesCreateLists(context.Background(), siteId).MicrosoftGraphList(microsoftGraphList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesCreateLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesCreateLists`: MicrosoftGraphList
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesCreateLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesCreateListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphList** | [**MicrosoftGraphList**](MicrosoftGraphList.md) | New navigation property | 

### Return type

[**MicrosoftGraphList**](MicrosoftGraphList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesDeleteLists

> SitesDeleteLists(ctx, siteId, listId).IfMatch(ifMatch).Execute()

Delete navigation property lists for sites



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
    listId := "listId_example" // string | key: id of list
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesDeleteLists(context.Background(), siteId, listId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesDeleteLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesDeleteListsRequest struct via the builder pattern


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


## SitesGetLists

> MicrosoftGraphList SitesGetLists(ctx, siteId, listId).Select_(select_).Expand(expand).Execute()

Get lists from sites



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
    listId := "listId_example" // string | key: id of list
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesGetLists(context.Background(), siteId, listId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesGetLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetLists`: MicrosoftGraphList
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesGetLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphList**](MicrosoftGraphList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListLists

> CollectionOfList SitesListLists(ctx, siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get lists from sites



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
    resp, r, err := api_client.SitesListApi.SitesListLists(context.Background(), siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListLists`: CollectionOfList
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListListsRequest struct via the builder pattern


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

[**CollectionOfList**](CollectionOfList.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsColumnsDeleteRefSourceColumn

> SitesListsColumnsDeleteRefSourceColumn(ctx, siteId, listId, columnDefinitionId).IfMatch(ifMatch).Execute()

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
    listId := "listId_example" // string | key: id of list
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsColumnsDeleteRefSourceColumn(context.Background(), siteId, listId, columnDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsColumnsDeleteRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsColumnsDeleteRefSourceColumnRequest struct via the builder pattern


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


## SitesListsColumnsGetRefSourceColumn

> string SitesListsColumnsGetRefSourceColumn(ctx, siteId, listId, columnDefinitionId).Execute()

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
    listId := "listId_example" // string | key: id of list
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsColumnsGetRefSourceColumn(context.Background(), siteId, listId, columnDefinitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsColumnsGetRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsColumnsGetRefSourceColumn`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsColumnsGetRefSourceColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsColumnsGetRefSourceColumnRequest struct via the builder pattern


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


## SitesListsColumnsGetSourceColumn

> MicrosoftGraphColumnDefinition SitesListsColumnsGetSourceColumn(ctx, siteId, listId, columnDefinitionId).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsColumnsGetSourceColumn(context.Background(), siteId, listId, columnDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsColumnsGetSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsColumnsGetSourceColumn`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsColumnsGetSourceColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsColumnsGetSourceColumnRequest struct via the builder pattern


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


## SitesListsColumnsUpdateRefSourceColumn

> SitesListsColumnsUpdateRefSourceColumn(ctx, siteId, listId, columnDefinitionId).RequestBody(requestBody).Execute()

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
    listId := "listId_example" // string | key: id of list
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsColumnsUpdateRefSourceColumn(context.Background(), siteId, listId, columnDefinitionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsColumnsUpdateRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsColumnsUpdateRefSourceColumnRequest struct via the builder pattern


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


## SitesListsContentTypesColumnsDeleteRefSourceColumn

> SitesListsContentTypesColumnsDeleteRefSourceColumn(ctx, siteId, listId, contentTypeId, columnDefinitionId).IfMatch(ifMatch).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesColumnsDeleteRefSourceColumn(context.Background(), siteId, listId, contentTypeId, columnDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesColumnsDeleteRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesColumnsDeleteRefSourceColumnRequest struct via the builder pattern


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


## SitesListsContentTypesColumnsGetRefSourceColumn

> string SitesListsContentTypesColumnsGetRefSourceColumn(ctx, siteId, listId, contentTypeId, columnDefinitionId).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesColumnsGetRefSourceColumn(context.Background(), siteId, listId, contentTypeId, columnDefinitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesColumnsGetRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesColumnsGetRefSourceColumn`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesColumnsGetRefSourceColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesColumnsGetRefSourceColumnRequest struct via the builder pattern


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


## SitesListsContentTypesColumnsGetSourceColumn

> MicrosoftGraphColumnDefinition SitesListsContentTypesColumnsGetSourceColumn(ctx, siteId, listId, contentTypeId, columnDefinitionId).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesColumnsGetSourceColumn(context.Background(), siteId, listId, contentTypeId, columnDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesColumnsGetSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesColumnsGetSourceColumn`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesColumnsGetSourceColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesColumnsGetSourceColumnRequest struct via the builder pattern


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


## SitesListsContentTypesColumnsUpdateRefSourceColumn

> SitesListsContentTypesColumnsUpdateRefSourceColumn(ctx, siteId, listId, contentTypeId, columnDefinitionId).RequestBody(requestBody).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesColumnsUpdateRefSourceColumn(context.Background(), siteId, listId, contentTypeId, columnDefinitionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesColumnsUpdateRefSourceColumn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesColumnsUpdateRefSourceColumnRequest struct via the builder pattern


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


## SitesListsContentTypesCreateColumnLinks

> MicrosoftGraphColumnLink SitesListsContentTypesCreateColumnLinks(ctx, siteId, listId, contentTypeId).MicrosoftGraphColumnLink(microsoftGraphColumnLink).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    microsoftGraphColumnLink := *openapiclient.NewMicrosoftGraphColumnLink() // MicrosoftGraphColumnLink | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesCreateColumnLinks(context.Background(), siteId, listId, contentTypeId).MicrosoftGraphColumnLink(microsoftGraphColumnLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesCreateColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesCreateColumnLinks`: MicrosoftGraphColumnLink
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesCreateColumnLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesCreateColumnLinksRequest struct via the builder pattern


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


## SitesListsContentTypesCreateColumns

> MicrosoftGraphColumnDefinition SitesListsContentTypesCreateColumns(ctx, siteId, listId, contentTypeId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    microsoftGraphColumnDefinition := *openapiclient.NewMicrosoftGraphColumnDefinition() // MicrosoftGraphColumnDefinition | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesCreateColumns(context.Background(), siteId, listId, contentTypeId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesCreateColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesCreateColumns`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesCreateColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesCreateColumnsRequest struct via the builder pattern


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


## SitesListsContentTypesCreateRefBaseTypes

> map[string]interface{} SitesListsContentTypesCreateRefBaseTypes(ctx, siteId, listId, contentTypeId).RequestBody(requestBody).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesCreateRefBaseTypes(context.Background(), siteId, listId, contentTypeId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesCreateRefBaseTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesCreateRefBaseTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesCreateRefBaseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesCreateRefBaseTypesRequest struct via the builder pattern


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


## SitesListsContentTypesCreateRefColumnPositions

> map[string]interface{} SitesListsContentTypesCreateRefColumnPositions(ctx, siteId, listId, contentTypeId).RequestBody(requestBody).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesCreateRefColumnPositions(context.Background(), siteId, listId, contentTypeId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesCreateRefColumnPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesCreateRefColumnPositions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesCreateRefColumnPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesCreateRefColumnPositionsRequest struct via the builder pattern


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


## SitesListsContentTypesDeleteColumnLinks

> SitesListsContentTypesDeleteColumnLinks(ctx, siteId, listId, contentTypeId, columnLinkId).IfMatch(ifMatch).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnLinkId := "columnLinkId_example" // string | key: id of columnLink
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesDeleteColumnLinks(context.Background(), siteId, listId, contentTypeId, columnLinkId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesDeleteColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnLinkId** | **string** | key: id of columnLink | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesDeleteColumnLinksRequest struct via the builder pattern


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


## SitesListsContentTypesDeleteColumns

> SitesListsContentTypesDeleteColumns(ctx, siteId, listId, contentTypeId, columnDefinitionId).IfMatch(ifMatch).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesDeleteColumns(context.Background(), siteId, listId, contentTypeId, columnDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesDeleteColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesDeleteColumnsRequest struct via the builder pattern


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


## SitesListsContentTypesDeleteRefBase

> SitesListsContentTypesDeleteRefBase(ctx, siteId, listId, contentTypeId).IfMatch(ifMatch).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesDeleteRefBase(context.Background(), siteId, listId, contentTypeId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesDeleteRefBase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesDeleteRefBaseRequest struct via the builder pattern


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


## SitesListsContentTypesGetBase

> MicrosoftGraphContentType SitesListsContentTypesGetBase(ctx, siteId, listId, contentTypeId).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesGetBase(context.Background(), siteId, listId, contentTypeId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesGetBase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesGetBase`: MicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesGetBase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesGetBaseRequest struct via the builder pattern


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


## SitesListsContentTypesGetColumnLinks

> MicrosoftGraphColumnLink SitesListsContentTypesGetColumnLinks(ctx, siteId, listId, contentTypeId, columnLinkId).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnLinkId := "columnLinkId_example" // string | key: id of columnLink
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesGetColumnLinks(context.Background(), siteId, listId, contentTypeId, columnLinkId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesGetColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesGetColumnLinks`: MicrosoftGraphColumnLink
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesGetColumnLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnLinkId** | **string** | key: id of columnLink | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesGetColumnLinksRequest struct via the builder pattern


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


## SitesListsContentTypesGetColumns

> MicrosoftGraphColumnDefinition SitesListsContentTypesGetColumns(ctx, siteId, listId, contentTypeId, columnDefinitionId).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesGetColumns(context.Background(), siteId, listId, contentTypeId, columnDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesGetColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesGetColumns`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesGetColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesGetColumnsRequest struct via the builder pattern


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


## SitesListsContentTypesGetRefBase

> string SitesListsContentTypesGetRefBase(ctx, siteId, listId, contentTypeId).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesGetRefBase(context.Background(), siteId, listId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesGetRefBase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesGetRefBase`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesGetRefBase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesGetRefBaseRequest struct via the builder pattern


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


## SitesListsContentTypesListBaseTypes

> CollectionOfContentType SitesListsContentTypesListBaseTypes(ctx, siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
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
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesListBaseTypes(context.Background(), siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesListBaseTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesListBaseTypes`: CollectionOfContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesListBaseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesListBaseTypesRequest struct via the builder pattern


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


## SitesListsContentTypesListColumnLinks

> CollectionOfColumnLink SitesListsContentTypesListColumnLinks(ctx, siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
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
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesListColumnLinks(context.Background(), siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesListColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesListColumnLinks`: CollectionOfColumnLink
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesListColumnLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesListColumnLinksRequest struct via the builder pattern


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


## SitesListsContentTypesListColumnPositions

> CollectionOfColumnDefinition SitesListsContentTypesListColumnPositions(ctx, siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
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
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesListColumnPositions(context.Background(), siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesListColumnPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesListColumnPositions`: CollectionOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesListColumnPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesListColumnPositionsRequest struct via the builder pattern


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


## SitesListsContentTypesListColumns

> CollectionOfColumnDefinition SitesListsContentTypesListColumns(ctx, siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
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
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesListColumns(context.Background(), siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesListColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesListColumns`: CollectionOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesListColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesListColumnsRequest struct via the builder pattern


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


## SitesListsContentTypesListRefBaseTypes

> CollectionOfLinksOfContentType SitesListsContentTypesListRefBaseTypes(ctx, siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesListRefBaseTypes(context.Background(), siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesListRefBaseTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesListRefBaseTypes`: CollectionOfLinksOfContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesListRefBaseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesListRefBaseTypesRequest struct via the builder pattern


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


## SitesListsContentTypesListRefColumnPositions

> CollectionOfLinksOfColumnDefinition SitesListsContentTypesListRefColumnPositions(ctx, siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesListRefColumnPositions(context.Background(), siteId, listId, contentTypeId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesListRefColumnPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsContentTypesListRefColumnPositions`: CollectionOfLinksOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsContentTypesListRefColumnPositions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesListRefColumnPositionsRequest struct via the builder pattern


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


## SitesListsContentTypesUpdateColumnLinks

> SitesListsContentTypesUpdateColumnLinks(ctx, siteId, listId, contentTypeId, columnLinkId).MicrosoftGraphColumnLink(microsoftGraphColumnLink).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnLinkId := "columnLinkId_example" // string | key: id of columnLink
    microsoftGraphColumnLink := *openapiclient.NewMicrosoftGraphColumnLink() // MicrosoftGraphColumnLink | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesUpdateColumnLinks(context.Background(), siteId, listId, contentTypeId, columnLinkId).MicrosoftGraphColumnLink(microsoftGraphColumnLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesUpdateColumnLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnLinkId** | **string** | key: id of columnLink | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesUpdateColumnLinksRequest struct via the builder pattern


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


## SitesListsContentTypesUpdateColumns

> SitesListsContentTypesUpdateColumns(ctx, siteId, listId, contentTypeId, columnDefinitionId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    microsoftGraphColumnDefinition := *openapiclient.NewMicrosoftGraphColumnDefinition() // MicrosoftGraphColumnDefinition | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesUpdateColumns(context.Background(), siteId, listId, contentTypeId, columnDefinitionId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesUpdateColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesUpdateColumnsRequest struct via the builder pattern


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


## SitesListsContentTypesUpdateRefBase

> SitesListsContentTypesUpdateRefBase(ctx, siteId, listId, contentTypeId).RequestBody(requestBody).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsContentTypesUpdateRefBase(context.Background(), siteId, listId, contentTypeId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsContentTypesUpdateRefBase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsContentTypesUpdateRefBaseRequest struct via the builder pattern


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


## SitesListsCreateColumns

> MicrosoftGraphColumnDefinition SitesListsCreateColumns(ctx, siteId, listId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()

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
    listId := "listId_example" // string | key: id of list
    microsoftGraphColumnDefinition := *openapiclient.NewMicrosoftGraphColumnDefinition() // MicrosoftGraphColumnDefinition | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsCreateColumns(context.Background(), siteId, listId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsCreateColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsCreateColumns`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsCreateColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsCreateColumnsRequest struct via the builder pattern


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


## SitesListsCreateContentTypes

> MicrosoftGraphContentType SitesListsCreateContentTypes(ctx, siteId, listId).MicrosoftGraphContentType(microsoftGraphContentType).Execute()

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
    listId := "listId_example" // string | key: id of list
    microsoftGraphContentType := *openapiclient.NewMicrosoftGraphContentType() // MicrosoftGraphContentType | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsCreateContentTypes(context.Background(), siteId, listId).MicrosoftGraphContentType(microsoftGraphContentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsCreateContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsCreateContentTypes`: MicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsCreateContentTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsCreateContentTypesRequest struct via the builder pattern


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


## SitesListsCreateItems

> MicrosoftGraphListItem SitesListsCreateItems(ctx, siteId, listId).MicrosoftGraphListItem(microsoftGraphListItem).Execute()

Create new navigation property to items for sites



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
    listId := "listId_example" // string | key: id of list
    microsoftGraphListItem := *openapiclient.NewMicrosoftGraphListItem() // MicrosoftGraphListItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsCreateItems(context.Background(), siteId, listId).MicrosoftGraphListItem(microsoftGraphListItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsCreateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsCreateItems`: MicrosoftGraphListItem
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsCreateItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsCreateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphListItem** | [**MicrosoftGraphListItem**](MicrosoftGraphListItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphListItem**](MicrosoftGraphListItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsCreateSubscriptions

> MicrosoftGraphSubscription SitesListsCreateSubscriptions(ctx, siteId, listId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()

Create new navigation property to subscriptions for sites



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
    listId := "listId_example" // string | key: id of list
    microsoftGraphSubscription := *openapiclient.NewMicrosoftGraphSubscription() // MicrosoftGraphSubscription | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsCreateSubscriptions(context.Background(), siteId, listId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsCreateSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsCreateSubscriptions`: MicrosoftGraphSubscription
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsCreateSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsCreateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md) | New navigation property | 

### Return type

[**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsDeleteColumns

> SitesListsDeleteColumns(ctx, siteId, listId, columnDefinitionId).IfMatch(ifMatch).Execute()

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
    listId := "listId_example" // string | key: id of list
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsDeleteColumns(context.Background(), siteId, listId, columnDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsDeleteColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsDeleteColumnsRequest struct via the builder pattern


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


## SitesListsDeleteContentTypes

> SitesListsDeleteContentTypes(ctx, siteId, listId, contentTypeId).IfMatch(ifMatch).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsDeleteContentTypes(context.Background(), siteId, listId, contentTypeId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsDeleteContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsDeleteContentTypesRequest struct via the builder pattern


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


## SitesListsDeleteDrive

> SitesListsDeleteDrive(ctx, siteId, listId).IfMatch(ifMatch).Execute()

Delete navigation property drive for sites



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
    listId := "listId_example" // string | key: id of list
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsDeleteDrive(context.Background(), siteId, listId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsDeleteDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsDeleteDriveRequest struct via the builder pattern


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


## SitesListsDeleteItems

> SitesListsDeleteItems(ctx, siteId, listId, listItemId).IfMatch(ifMatch).Execute()

Delete navigation property items for sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsDeleteItems(context.Background(), siteId, listId, listItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsDeleteItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsDeleteItemsRequest struct via the builder pattern


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


## SitesListsDeleteSubscriptions

> SitesListsDeleteSubscriptions(ctx, siteId, listId, subscriptionId).IfMatch(ifMatch).Execute()

Delete navigation property subscriptions for sites



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
    listId := "listId_example" // string | key: id of list
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsDeleteSubscriptions(context.Background(), siteId, listId, subscriptionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsDeleteSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsDeleteSubscriptionsRequest struct via the builder pattern


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


## SitesListsGetColumns

> MicrosoftGraphColumnDefinition SitesListsGetColumns(ctx, siteId, listId, columnDefinitionId).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsGetColumns(context.Background(), siteId, listId, columnDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsGetColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsGetColumns`: MicrosoftGraphColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsGetColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsGetColumnsRequest struct via the builder pattern


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


## SitesListsGetContentTypes

> MicrosoftGraphContentType SitesListsGetContentTypes(ctx, siteId, listId, contentTypeId).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsGetContentTypes(context.Background(), siteId, listId, contentTypeId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsGetContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsGetContentTypes`: MicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsGetContentTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsGetContentTypesRequest struct via the builder pattern


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


## SitesListsGetDrive

> MicrosoftGraphDrive SitesListsGetDrive(ctx, siteId, listId).Select_(select_).Expand(expand).Execute()

Get drive from sites



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
    listId := "listId_example" // string | key: id of list
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsGetDrive(context.Background(), siteId, listId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsGetDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsGetDrive`: MicrosoftGraphDrive
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsGetDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsGetDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsGetItems

> MicrosoftGraphListItem SitesListsGetItems(ctx, siteId, listId, listItemId).Select_(select_).Expand(expand).Execute()

Get items from sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsGetItems(context.Background(), siteId, listId, listItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsGetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsGetItems`: MicrosoftGraphListItem
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsGetItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphListItem**](MicrosoftGraphListItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsGetSubscriptions

> MicrosoftGraphSubscription SitesListsGetSubscriptions(ctx, siteId, listId, subscriptionId).Select_(select_).Expand(expand).Execute()

Get subscriptions from sites



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
    listId := "listId_example" // string | key: id of list
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsGetSubscriptions(context.Background(), siteId, listId, subscriptionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsGetSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsGetSubscriptions`: MicrosoftGraphSubscription
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsGetSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsCreateVersions

> MicrosoftGraphListItemVersion SitesListsItemsCreateVersions(ctx, siteId, listId, listItemId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()

Create new navigation property to versions for sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    microsoftGraphListItemVersion := *openapiclient.NewMicrosoftGraphListItemVersion() // MicrosoftGraphListItemVersion | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsCreateVersions(context.Background(), siteId, listId, listItemId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsCreateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsCreateVersions`: MicrosoftGraphListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsCreateVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsCreateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md) | New navigation property | 

### Return type

[**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsDeleteDriveItem

> SitesListsItemsDeleteDriveItem(ctx, siteId, listId, listItemId).IfMatch(ifMatch).Execute()

Delete navigation property driveItem for sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsDeleteDriveItem(context.Background(), siteId, listId, listItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsDeleteDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsDeleteDriveItemRequest struct via the builder pattern


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


## SitesListsItemsDeleteFields

> SitesListsItemsDeleteFields(ctx, siteId, listId, listItemId).IfMatch(ifMatch).Execute()

Delete navigation property fields for sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsDeleteFields(context.Background(), siteId, listId, listItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsDeleteFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsDeleteFieldsRequest struct via the builder pattern


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


## SitesListsItemsDeleteRefAnalytics

> SitesListsItemsDeleteRefAnalytics(ctx, siteId, listId, listItemId).IfMatch(ifMatch).Execute()

Delete ref of navigation property analytics for sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsDeleteRefAnalytics(context.Background(), siteId, listId, listItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsDeleteRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsDeleteRefAnalyticsRequest struct via the builder pattern


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


## SitesListsItemsDeleteVersions

> SitesListsItemsDeleteVersions(ctx, siteId, listId, listItemId, listItemVersionId).IfMatch(ifMatch).Execute()

Delete navigation property versions for sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsDeleteVersions(context.Background(), siteId, listId, listItemId, listItemVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsDeleteVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsDeleteVersionsRequest struct via the builder pattern


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


## SitesListsItemsGetAnalytics

> MicrosoftGraphItemAnalytics SitesListsItemsGetAnalytics(ctx, siteId, listId, listItemId).Select_(select_).Expand(expand).Execute()

Get analytics from sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsGetAnalytics(context.Background(), siteId, listId, listItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsGetAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsGetAnalytics`: MicrosoftGraphItemAnalytics
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsGetAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsGetAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphItemAnalytics**](MicrosoftGraphItemAnalytics.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsGetDriveItem

> MicrosoftGraphDriveItem SitesListsItemsGetDriveItem(ctx, siteId, listId, listItemId).Select_(select_).Expand(expand).Execute()

Get driveItem from sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsGetDriveItem(context.Background(), siteId, listId, listItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsGetDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsGetDriveItem`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsGetDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsGetDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsGetDriveItemContent

> *os.File SitesListsItemsGetDriveItemContent(ctx, siteId, listId, listItemId).Execute()

Get media content for the navigation property driveItem from sites

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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsGetDriveItemContent(context.Background(), siteId, listId, listItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsGetDriveItemContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsGetDriveItemContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsGetDriveItemContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsGetDriveItemContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsGetFields

> MicrosoftGraphFieldValueSet SitesListsItemsGetFields(ctx, siteId, listId, listItemId).Select_(select_).Expand(expand).Execute()

Get fields from sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsGetFields(context.Background(), siteId, listId, listItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsGetFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsGetFields`: MicrosoftGraphFieldValueSet
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsGetFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsGetFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsGetRefAnalytics

> string SitesListsItemsGetRefAnalytics(ctx, siteId, listId, listItemId).Execute()

Get ref of analytics from sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsGetRefAnalytics(context.Background(), siteId, listId, listItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsGetRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsGetRefAnalytics`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsGetRefAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsGetRefAnalyticsRequest struct via the builder pattern


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


## SitesListsItemsGetVersions

> MicrosoftGraphListItemVersion SitesListsItemsGetVersions(ctx, siteId, listId, listItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()

Get versions from sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsGetVersions(context.Background(), siteId, listId, listItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsGetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsGetVersions`: MicrosoftGraphListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsGetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsListVersions

> CollectionOfListItemVersion SitesListsItemsListVersions(ctx, siteId, listId, listItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get versions from sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
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
    resp, r, err := api_client.SitesListApi.SitesListsItemsListVersions(context.Background(), siteId, listId, listItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsListVersions`: CollectionOfListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsListVersionsRequest struct via the builder pattern


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

[**CollectionOfListItemVersion**](CollectionOfListItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsUpdateDriveItem

> SitesListsItemsUpdateDriveItem(ctx, siteId, listId, listItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property driveItem in sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsUpdateDriveItem(context.Background(), siteId, listId, listItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsUpdateDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsUpdateDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property values | 

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


## SitesListsItemsUpdateDriveItemContent

> SitesListsItemsUpdateDriveItemContent(ctx, siteId, listId, listItemId).Body(body).Execute()

Update media content for the navigation property driveItem in sites

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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsUpdateDriveItemContent(context.Background(), siteId, listId, listItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsUpdateDriveItemContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsUpdateDriveItemContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | ***os.File** | New media content. | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsUpdateFields

> SitesListsItemsUpdateFields(ctx, siteId, listId, listItemId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()

Update the navigation property fields in sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    microsoftGraphFieldValueSet := *openapiclient.NewMicrosoftGraphFieldValueSet() // MicrosoftGraphFieldValueSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsUpdateFields(context.Background(), siteId, listId, listItemId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsUpdateFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsUpdateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphFieldValueSet** | [**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md) | New navigation property values | 

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


## SitesListsItemsUpdateRefAnalytics

> SitesListsItemsUpdateRefAnalytics(ctx, siteId, listId, listItemId).RequestBody(requestBody).Execute()

Update the ref of navigation property analytics in sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsUpdateRefAnalytics(context.Background(), siteId, listId, listItemId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsUpdateRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsUpdateRefAnalyticsRequest struct via the builder pattern


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


## SitesListsItemsUpdateVersions

> SitesListsItemsUpdateVersions(ctx, siteId, listId, listItemId, listItemVersionId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()

Update the navigation property versions in sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    microsoftGraphListItemVersion := *openapiclient.NewMicrosoftGraphListItemVersion() // MicrosoftGraphListItemVersion | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsUpdateVersions(context.Background(), siteId, listId, listItemId, listItemVersionId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsUpdateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsUpdateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md) | New navigation property values | 

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


## SitesListsItemsVersionsDeleteFields

> SitesListsItemsVersionsDeleteFields(ctx, siteId, listId, listItemId, listItemVersionId).IfMatch(ifMatch).Execute()

Delete navigation property fields for sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsVersionsDeleteFields(context.Background(), siteId, listId, listItemId, listItemVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsVersionsDeleteFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsVersionsDeleteFieldsRequest struct via the builder pattern


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


## SitesListsItemsVersionsGetFields

> MicrosoftGraphFieldValueSet SitesListsItemsVersionsGetFields(ctx, siteId, listId, listItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()

Get fields from sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsVersionsGetFields(context.Background(), siteId, listId, listItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsVersionsGetFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsItemsVersionsGetFields`: MicrosoftGraphFieldValueSet
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsItemsVersionsGetFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsVersionsGetFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsItemsVersionsUpdateFields

> SitesListsItemsVersionsUpdateFields(ctx, siteId, listId, listItemId, listItemVersionId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()

Update the navigation property fields in sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    microsoftGraphFieldValueSet := *openapiclient.NewMicrosoftGraphFieldValueSet() // MicrosoftGraphFieldValueSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsItemsVersionsUpdateFields(context.Background(), siteId, listId, listItemId, listItemVersionId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsItemsVersionsUpdateFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsItemsVersionsUpdateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphFieldValueSet** | [**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md) | New navigation property values | 

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


## SitesListsListColumns

> CollectionOfColumnDefinition SitesListsListColumns(ctx, siteId, listId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
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
    resp, r, err := api_client.SitesListApi.SitesListsListColumns(context.Background(), siteId, listId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsListColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsListColumns`: CollectionOfColumnDefinition
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsListColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsListColumnsRequest struct via the builder pattern


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


## SitesListsListContentTypes

> CollectionOfContentType SitesListsListContentTypes(ctx, siteId, listId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    listId := "listId_example" // string | key: id of list
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
    resp, r, err := api_client.SitesListApi.SitesListsListContentTypes(context.Background(), siteId, listId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsListContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsListContentTypes`: CollectionOfContentType
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsListContentTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsListContentTypesRequest struct via the builder pattern


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


## SitesListsListItems

> CollectionOfListItem SitesListsListItems(ctx, siteId, listId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get items from sites



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
    listId := "listId_example" // string | key: id of list
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
    resp, r, err := api_client.SitesListApi.SitesListsListItems(context.Background(), siteId, listId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsListItems`: CollectionOfListItem
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsListItemsRequest struct via the builder pattern


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

[**CollectionOfListItem**](CollectionOfListItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsListSubscriptions

> CollectionOfSubscription SitesListsListSubscriptions(ctx, siteId, listId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get subscriptions from sites



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
    listId := "listId_example" // string | key: id of list
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
    resp, r, err := api_client.SitesListApi.SitesListsListSubscriptions(context.Background(), siteId, listId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsListSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListsListSubscriptions`: CollectionOfSubscription
    fmt.Fprintf(os.Stdout, "Response from `SitesListApi.SitesListsListSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsListSubscriptionsRequest struct via the builder pattern


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

[**CollectionOfSubscription**](CollectionOfSubscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListsUpdateColumns

> SitesListsUpdateColumns(ctx, siteId, listId, columnDefinitionId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()

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
    listId := "listId_example" // string | key: id of list
    columnDefinitionId := "columnDefinitionId_example" // string | key: id of columnDefinition
    microsoftGraphColumnDefinition := *openapiclient.NewMicrosoftGraphColumnDefinition() // MicrosoftGraphColumnDefinition | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsUpdateColumns(context.Background(), siteId, listId, columnDefinitionId).MicrosoftGraphColumnDefinition(microsoftGraphColumnDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsUpdateColumns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**columnDefinitionId** | **string** | key: id of columnDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsUpdateColumnsRequest struct via the builder pattern


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


## SitesListsUpdateContentTypes

> SitesListsUpdateContentTypes(ctx, siteId, listId, contentTypeId).MicrosoftGraphContentType(microsoftGraphContentType).Execute()

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
    listId := "listId_example" // string | key: id of list
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    microsoftGraphContentType := *openapiclient.NewMicrosoftGraphContentType() // MicrosoftGraphContentType | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsUpdateContentTypes(context.Background(), siteId, listId, contentTypeId).MicrosoftGraphContentType(microsoftGraphContentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsUpdateContentTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsUpdateContentTypesRequest struct via the builder pattern


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


## SitesListsUpdateDrive

> SitesListsUpdateDrive(ctx, siteId, listId).MicrosoftGraphDrive(microsoftGraphDrive).Execute()

Update the navigation property drive in sites



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
    listId := "listId_example" // string | key: id of list
    microsoftGraphDrive := *openapiclient.NewMicrosoftGraphDrive() // MicrosoftGraphDrive | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsUpdateDrive(context.Background(), siteId, listId).MicrosoftGraphDrive(microsoftGraphDrive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsUpdateDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsUpdateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md) | New navigation property values | 

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


## SitesListsUpdateItems

> SitesListsUpdateItems(ctx, siteId, listId, listItemId).MicrosoftGraphListItem(microsoftGraphListItem).Execute()

Update the navigation property items in sites



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
    listId := "listId_example" // string | key: id of list
    listItemId := "listItemId_example" // string | key: id of listItem
    microsoftGraphListItem := *openapiclient.NewMicrosoftGraphListItem() // MicrosoftGraphListItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsUpdateItems(context.Background(), siteId, listId, listItemId).MicrosoftGraphListItem(microsoftGraphListItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsUpdateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsUpdateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphListItem** | [**MicrosoftGraphListItem**](MicrosoftGraphListItem.md) | New navigation property values | 

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


## SitesListsUpdateSubscriptions

> SitesListsUpdateSubscriptions(ctx, siteId, listId, subscriptionId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()

Update the navigation property subscriptions in sites



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
    listId := "listId_example" // string | key: id of list
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    microsoftGraphSubscription := *openapiclient.NewMicrosoftGraphSubscription() // MicrosoftGraphSubscription | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesListsUpdateSubscriptions(context.Background(), siteId, listId, subscriptionId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesListsUpdateSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListsUpdateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md) | New navigation property values | 

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


## SitesUpdateLists

> SitesUpdateLists(ctx, siteId, listId).MicrosoftGraphList(microsoftGraphList).Execute()

Update the navigation property lists in sites



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
    listId := "listId_example" // string | key: id of list
    microsoftGraphList := *openapiclient.NewMicrosoftGraphList() // MicrosoftGraphList | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesListApi.SitesUpdateLists(context.Background(), siteId, listId).MicrosoftGraphList(microsoftGraphList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesListApi.SitesUpdateLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**listId** | **string** | key: id of list | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesUpdateListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphList** | [**MicrosoftGraphList**](MicrosoftGraphList.md) | New navigation property values | 

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

