# \UsersProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreatePhotos**](UsersProfilePhotoApi.md#UsersCreatePhotos) | **Post** /users/{user-id}/photos | Create new navigation property to photos for users
[**UsersDeletePhoto**](UsersProfilePhotoApi.md#UsersDeletePhoto) | **Delete** /users/{user-id}/photo | Delete navigation property photo for users
[**UsersDeletePhotos**](UsersProfilePhotoApi.md#UsersDeletePhotos) | **Delete** /users/{user-id}/photos/{profilePhoto-id} | Delete navigation property photos for users
[**UsersGetPhoto**](UsersProfilePhotoApi.md#UsersGetPhoto) | **Get** /users/{user-id}/photo | Get photo from users
[**UsersGetPhotoContent**](UsersProfilePhotoApi.md#UsersGetPhotoContent) | **Get** /users/{user-id}/photo/$value | Get media content for the navigation property photo from users
[**UsersGetPhotos**](UsersProfilePhotoApi.md#UsersGetPhotos) | **Get** /users/{user-id}/photos/{profilePhoto-id} | Get photos from users
[**UsersGetPhotosContent**](UsersProfilePhotoApi.md#UsersGetPhotosContent) | **Get** /users/{user-id}/photos/{profilePhoto-id}/$value | Get media content for the navigation property photos from users
[**UsersListPhotos**](UsersProfilePhotoApi.md#UsersListPhotos) | **Get** /users/{user-id}/photos | Get photos from users
[**UsersUpdatePhoto**](UsersProfilePhotoApi.md#UsersUpdatePhoto) | **Patch** /users/{user-id}/photo | Update the navigation property photo in users
[**UsersUpdatePhotoContent**](UsersProfilePhotoApi.md#UsersUpdatePhotoContent) | **Put** /users/{user-id}/photo/$value | Update media content for the navigation property photo in users
[**UsersUpdatePhotos**](UsersProfilePhotoApi.md#UsersUpdatePhotos) | **Patch** /users/{user-id}/photos/{profilePhoto-id} | Update the navigation property photos in users
[**UsersUpdatePhotosContent**](UsersProfilePhotoApi.md#UsersUpdatePhotosContent) | **Put** /users/{user-id}/photos/{profilePhoto-id}/$value | Update media content for the navigation property photos in users



## UsersCreatePhotos

> MicrosoftGraphProfilePhoto UsersCreatePhotos(ctx, userId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Create new navigation property to photos for users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersCreatePhotos(context.Background(), userId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersCreatePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreatePhotos`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `UsersProfilePhotoApi.UsersCreatePhotos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreatePhotosRequest struct via the builder pattern


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


## UsersDeletePhoto

> UsersDeletePhoto(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property photo for users



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
    userId := "userId_example" // string | key: id of user
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersDeletePhoto(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersDeletePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeletePhotoRequest struct via the builder pattern


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


## UsersDeletePhotos

> UsersDeletePhotos(ctx, userId, profilePhotoId).IfMatch(ifMatch).Execute()

Delete navigation property photos for users



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
    userId := "userId_example" // string | key: id of user
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersDeletePhotos(context.Background(), userId, profilePhotoId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersDeletePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeletePhotosRequest struct via the builder pattern


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


## UsersGetPhoto

> MicrosoftGraphProfilePhoto UsersGetPhoto(ctx, userId).Select_(select_).Execute()

Get photo from users



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
    userId := "userId_example" // string | key: id of user
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersGetPhoto(context.Background(), userId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersGetPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetPhoto`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `UsersProfilePhotoApi.UsersGetPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetPhotoRequest struct via the builder pattern


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


## UsersGetPhotoContent

> *os.File UsersGetPhotoContent(ctx, userId).Execute()

Get media content for the navigation property photo from users



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
    userId := "userId_example" // string | key: id of user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersGetPhotoContent(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersGetPhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetPhotoContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UsersProfilePhotoApi.UsersGetPhotoContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetPhotoContentRequest struct via the builder pattern


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


## UsersGetPhotos

> MicrosoftGraphProfilePhoto UsersGetPhotos(ctx, userId, profilePhotoId).Select_(select_).Execute()

Get photos from users



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
    userId := "userId_example" // string | key: id of user
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersGetPhotos(context.Background(), userId, profilePhotoId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersGetPhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetPhotos`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `UsersProfilePhotoApi.UsersGetPhotos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetPhotosRequest struct via the builder pattern


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


## UsersGetPhotosContent

> *os.File UsersGetPhotosContent(ctx, userId, profilePhotoId).Execute()

Get media content for the navigation property photos from users

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
    userId := "userId_example" // string | key: id of user
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersGetPhotosContent(context.Background(), userId, profilePhotoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersGetPhotosContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetPhotosContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UsersProfilePhotoApi.UsersGetPhotosContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetPhotosContentRequest struct via the builder pattern


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


## UsersListPhotos

> CollectionOfProfilePhoto UsersListPhotos(ctx, userId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get photos from users



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
    userId := "userId_example" // string | key: id of user
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersListPhotos(context.Background(), userId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersListPhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListPhotos`: CollectionOfProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `UsersProfilePhotoApi.UsersListPhotos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListPhotosRequest struct via the builder pattern


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


## UsersUpdatePhoto

> UsersUpdatePhoto(ctx, userId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Update the navigation property photo in users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersUpdatePhoto(context.Background(), userId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersUpdatePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdatePhotoRequest struct via the builder pattern


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


## UsersUpdatePhotoContent

> UsersUpdatePhotoContent(ctx, userId).Body(body).Execute()

Update media content for the navigation property photo in users



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
    userId := "userId_example" // string | key: id of user
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersUpdatePhotoContent(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersUpdatePhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdatePhotoContentRequest struct via the builder pattern


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


## UsersUpdatePhotos

> UsersUpdatePhotos(ctx, userId, profilePhotoId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Update the navigation property photos in users



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
    userId := "userId_example" // string | key: id of user
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersUpdatePhotos(context.Background(), userId, profilePhotoId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersUpdatePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdatePhotosRequest struct via the builder pattern


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


## UsersUpdatePhotosContent

> UsersUpdatePhotosContent(ctx, userId, profilePhotoId).Body(body).Execute()

Update media content for the navigation property photos in users

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
    userId := "userId_example" // string | key: id of user
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersProfilePhotoApi.UsersUpdatePhotosContent(context.Background(), userId, profilePhotoId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersProfilePhotoApi.UsersUpdatePhotosContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdatePhotosContentRequest struct via the builder pattern


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

