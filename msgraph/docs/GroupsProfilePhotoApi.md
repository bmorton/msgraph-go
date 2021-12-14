# \GroupsProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreatePhotos**](GroupsProfilePhotoApi.md#GroupsCreatePhotos) | **Post** /groups/{group-id}/photos | Create new navigation property to photos for groups
[**GroupsDeletePhoto**](GroupsProfilePhotoApi.md#GroupsDeletePhoto) | **Delete** /groups/{group-id}/photo | Delete navigation property photo for groups
[**GroupsDeletePhotos**](GroupsProfilePhotoApi.md#GroupsDeletePhotos) | **Delete** /groups/{group-id}/photos/{profilePhoto-id} | Delete navigation property photos for groups
[**GroupsGetPhoto**](GroupsProfilePhotoApi.md#GroupsGetPhoto) | **Get** /groups/{group-id}/photo | Get photo from groups
[**GroupsGetPhotoContent**](GroupsProfilePhotoApi.md#GroupsGetPhotoContent) | **Get** /groups/{group-id}/photo/$value | Get media content for the navigation property photo from groups
[**GroupsGetPhotos**](GroupsProfilePhotoApi.md#GroupsGetPhotos) | **Get** /groups/{group-id}/photos/{profilePhoto-id} | Get photos from groups
[**GroupsGetPhotosContent**](GroupsProfilePhotoApi.md#GroupsGetPhotosContent) | **Get** /groups/{group-id}/photos/{profilePhoto-id}/$value | Get media content for the navigation property photos from groups
[**GroupsListPhotos**](GroupsProfilePhotoApi.md#GroupsListPhotos) | **Get** /groups/{group-id}/photos | Get photos from groups
[**GroupsUpdatePhoto**](GroupsProfilePhotoApi.md#GroupsUpdatePhoto) | **Patch** /groups/{group-id}/photo | Update the navigation property photo in groups
[**GroupsUpdatePhotoContent**](GroupsProfilePhotoApi.md#GroupsUpdatePhotoContent) | **Put** /groups/{group-id}/photo/$value | Update media content for the navigation property photo in groups
[**GroupsUpdatePhotos**](GroupsProfilePhotoApi.md#GroupsUpdatePhotos) | **Patch** /groups/{group-id}/photos/{profilePhoto-id} | Update the navigation property photos in groups
[**GroupsUpdatePhotosContent**](GroupsProfilePhotoApi.md#GroupsUpdatePhotosContent) | **Put** /groups/{group-id}/photos/{profilePhoto-id}/$value | Update media content for the navigation property photos in groups



## GroupsCreatePhotos

> MicrosoftGraphProfilePhoto GroupsCreatePhotos(ctx, groupId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Create new navigation property to photos for groups



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
    groupId := "groupId_example" // string | key: id of group
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsCreatePhotos(context.Background(), groupId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsCreatePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreatePhotos`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `GroupsProfilePhotoApi.GroupsCreatePhotos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreatePhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md) | New navigation property | 

### Return type

[**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDeletePhoto

> GroupsDeletePhoto(ctx, groupId).IfMatch(ifMatch).Execute()

Delete navigation property photo for groups



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
    groupId := "groupId_example" // string | key: id of group
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsDeletePhoto(context.Background(), groupId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsDeletePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeletePhotoRequest struct via the builder pattern


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


## GroupsDeletePhotos

> GroupsDeletePhotos(ctx, groupId, profilePhotoId).IfMatch(ifMatch).Execute()

Delete navigation property photos for groups



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
    groupId := "groupId_example" // string | key: id of group
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsDeletePhotos(context.Background(), groupId, profilePhotoId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsDeletePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeletePhotosRequest struct via the builder pattern


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


## GroupsGetPhoto

> MicrosoftGraphProfilePhoto GroupsGetPhoto(ctx, groupId).Select_(select_).Execute()

Get photo from groups



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
    groupId := "groupId_example" // string | key: id of group
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsGetPhoto(context.Background(), groupId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsGetPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetPhoto`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `GroupsProfilePhotoApi.GroupsGetPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetPhotoContent

> *os.File GroupsGetPhotoContent(ctx, groupId).Execute()

Get media content for the navigation property photo from groups



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
    groupId := "groupId_example" // string | key: id of group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsGetPhotoContent(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsGetPhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetPhotoContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GroupsProfilePhotoApi.GroupsGetPhotoContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetPhotoContentRequest struct via the builder pattern


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


## GroupsGetPhotos

> MicrosoftGraphProfilePhoto GroupsGetPhotos(ctx, groupId, profilePhotoId).Select_(select_).Execute()

Get photos from groups



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
    groupId := "groupId_example" // string | key: id of group
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsGetPhotos(context.Background(), groupId, profilePhotoId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsGetPhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetPhotos`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `GroupsProfilePhotoApi.GroupsGetPhotos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetPhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGetPhotosContent

> *os.File GroupsGetPhotosContent(ctx, groupId, profilePhotoId).Execute()

Get media content for the navigation property photos from groups

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
    groupId := "groupId_example" // string | key: id of group
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsGetPhotosContent(context.Background(), groupId, profilePhotoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsGetPhotosContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetPhotosContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GroupsProfilePhotoApi.GroupsGetPhotosContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetPhotosContentRequest struct via the builder pattern


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


## GroupsListPhotos

> CollectionOfProfilePhoto GroupsListPhotos(ctx, groupId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get photos from groups



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
    groupId := "groupId_example" // string | key: id of group
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsListPhotos(context.Background(), groupId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsListPhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsListPhotos`: CollectionOfProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `GroupsProfilePhotoApi.GroupsListPhotos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListPhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfProfilePhoto**](CollectionOfProfilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdatePhoto

> GroupsUpdatePhoto(ctx, groupId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Update the navigation property photo in groups



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
    groupId := "groupId_example" // string | key: id of group
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsUpdatePhoto(context.Background(), groupId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsUpdatePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdatePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md) | New navigation property values | 

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


## GroupsUpdatePhotoContent

> GroupsUpdatePhotoContent(ctx, groupId).Body(body).Execute()

Update media content for the navigation property photo in groups



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
    groupId := "groupId_example" // string | key: id of group
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsUpdatePhotoContent(context.Background(), groupId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsUpdatePhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdatePhotoContentRequest struct via the builder pattern


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


## GroupsUpdatePhotos

> GroupsUpdatePhotos(ctx, groupId, profilePhotoId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Update the navigation property photos in groups



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
    groupId := "groupId_example" // string | key: id of group
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsUpdatePhotos(context.Background(), groupId, profilePhotoId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsUpdatePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdatePhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md) | New navigation property values | 

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


## GroupsUpdatePhotosContent

> GroupsUpdatePhotosContent(ctx, groupId, profilePhotoId).Body(body).Execute()

Update media content for the navigation property photos in groups

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
    groupId := "groupId_example" // string | key: id of group
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsProfilePhotoApi.GroupsUpdatePhotosContent(context.Background(), groupId, profilePhotoId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsProfilePhotoApi.GroupsUpdatePhotosContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdatePhotosContentRequest struct via the builder pattern


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

