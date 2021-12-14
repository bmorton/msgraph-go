# \MeProfilePhotoApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreatePhotos**](MeProfilePhotoApi.md#MeCreatePhotos) | **Post** /me/photos | Create new navigation property to photos for me
[**MeDeletePhoto**](MeProfilePhotoApi.md#MeDeletePhoto) | **Delete** /me/photo | Delete navigation property photo for me
[**MeDeletePhotos**](MeProfilePhotoApi.md#MeDeletePhotos) | **Delete** /me/photos/{profilePhoto-id} | Delete navigation property photos for me
[**MeGetPhoto**](MeProfilePhotoApi.md#MeGetPhoto) | **Get** /me/photo | Get photo from me
[**MeGetPhotoContent**](MeProfilePhotoApi.md#MeGetPhotoContent) | **Get** /me/photo/$value | Get media content for the navigation property photo from me
[**MeGetPhotos**](MeProfilePhotoApi.md#MeGetPhotos) | **Get** /me/photos/{profilePhoto-id} | Get photos from me
[**MeGetPhotosContent**](MeProfilePhotoApi.md#MeGetPhotosContent) | **Get** /me/photos/{profilePhoto-id}/$value | Get media content for the navigation property photos from me
[**MeListPhotos**](MeProfilePhotoApi.md#MeListPhotos) | **Get** /me/photos | Get photos from me
[**MeUpdatePhoto**](MeProfilePhotoApi.md#MeUpdatePhoto) | **Patch** /me/photo | Update the navigation property photo in me
[**MeUpdatePhotoContent**](MeProfilePhotoApi.md#MeUpdatePhotoContent) | **Put** /me/photo/$value | Update media content for the navigation property photo in me
[**MeUpdatePhotos**](MeProfilePhotoApi.md#MeUpdatePhotos) | **Patch** /me/photos/{profilePhoto-id} | Update the navigation property photos in me
[**MeUpdatePhotosContent**](MeProfilePhotoApi.md#MeUpdatePhotosContent) | **Put** /me/photos/{profilePhoto-id}/$value | Update media content for the navigation property photos in me



## MeCreatePhotos

> MicrosoftGraphProfilePhoto MeCreatePhotos(ctx).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Create new navigation property to photos for me



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
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeCreatePhotos(context.Background()).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeCreatePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreatePhotos`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `MeProfilePhotoApi.MeCreatePhotos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreatePhotosRequest struct via the builder pattern


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


## MeDeletePhoto

> MeDeletePhoto(ctx).IfMatch(ifMatch).Execute()

Delete navigation property photo for me



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
    resp, r, err := api_client.MeProfilePhotoApi.MeDeletePhoto(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeDeletePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeletePhotoRequest struct via the builder pattern


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


## MeDeletePhotos

> MeDeletePhotos(ctx, profilePhotoId).IfMatch(ifMatch).Execute()

Delete navigation property photos for me



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
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeDeletePhotos(context.Background(), profilePhotoId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeDeletePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeletePhotosRequest struct via the builder pattern


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


## MeGetPhoto

> MicrosoftGraphProfilePhoto MeGetPhoto(ctx).Select_(select_).Execute()

Get photo from me



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeGetPhoto(context.Background()).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeGetPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetPhoto`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `MeProfilePhotoApi.MeGetPhoto`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetPhotoRequest struct via the builder pattern


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


## MeGetPhotoContent

> *os.File MeGetPhotoContent(ctx).Execute()

Get media content for the navigation property photo from me



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeGetPhotoContent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeGetPhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetPhotoContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MeProfilePhotoApi.MeGetPhotoContent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetPhotoContentRequest struct via the builder pattern


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


## MeGetPhotos

> MicrosoftGraphProfilePhoto MeGetPhotos(ctx, profilePhotoId).Select_(select_).Execute()

Get photos from me



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
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeGetPhotos(context.Background(), profilePhotoId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeGetPhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetPhotos`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `MeProfilePhotoApi.MeGetPhotos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetPhotosRequest struct via the builder pattern


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


## MeGetPhotosContent

> *os.File MeGetPhotosContent(ctx, profilePhotoId).Execute()

Get media content for the navigation property photos from me

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
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeGetPhotosContent(context.Background(), profilePhotoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeGetPhotosContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetPhotosContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MeProfilePhotoApi.MeGetPhotosContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetPhotosContentRequest struct via the builder pattern


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


## MeListPhotos

> CollectionOfProfilePhoto MeListPhotos(ctx).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get photos from me



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
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeListPhotos(context.Background()).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeListPhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListPhotos`: CollectionOfProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `MeProfilePhotoApi.MeListPhotos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListPhotosRequest struct via the builder pattern


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


## MeUpdatePhoto

> MeUpdatePhoto(ctx).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Update the navigation property photo in me



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
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeUpdatePhoto(context.Background()).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeUpdatePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdatePhotoRequest struct via the builder pattern


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


## MeUpdatePhotoContent

> MeUpdatePhotoContent(ctx).Body(body).Execute()

Update media content for the navigation property photo in me



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
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeUpdatePhotoContent(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeUpdatePhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdatePhotoContentRequest struct via the builder pattern


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


## MeUpdatePhotos

> MeUpdatePhotos(ctx, profilePhotoId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Update the navigation property photos in me



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
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeUpdatePhotos(context.Background(), profilePhotoId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeUpdatePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdatePhotosRequest struct via the builder pattern


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


## MeUpdatePhotosContent

> MeUpdatePhotosContent(ctx, profilePhotoId).Body(body).Execute()

Update media content for the navigation property photos in me

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
    profilePhotoId := "profilePhotoId_example" // string | key: id of profilePhoto
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeProfilePhotoApi.MeUpdatePhotosContent(context.Background(), profilePhotoId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeProfilePhotoApi.MeUpdatePhotosContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profilePhotoId** | **string** | key: id of profilePhoto | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdatePhotosContentRequest struct via the builder pattern


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

