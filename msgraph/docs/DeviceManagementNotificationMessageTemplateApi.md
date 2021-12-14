# \DeviceManagementNotificationMessageTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementCreateNotificationMessageTemplates) | **Post** /deviceManagement/notificationMessageTemplates | Create new navigation property to notificationMessageTemplates for deviceManagement
[**DeviceManagementDeleteNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementDeleteNotificationMessageTemplates) | **Delete** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id} | Delete navigation property notificationMessageTemplates for deviceManagement
[**DeviceManagementGetNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementGetNotificationMessageTemplates) | **Get** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id} | Get notificationMessageTemplates from deviceManagement
[**DeviceManagementListNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementListNotificationMessageTemplates) | **Get** /deviceManagement/notificationMessageTemplates | Get notificationMessageTemplates from deviceManagement
[**DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages) | **Post** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages | Create new navigation property to localizedNotificationMessages for deviceManagement
[**DeviceManagementNotificationMessageTemplatesDeleteLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesDeleteLocalizedNotificationMessages) | **Delete** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages/{localizedNotificationMessage-id} | Delete navigation property localizedNotificationMessages for deviceManagement
[**DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages) | **Get** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages/{localizedNotificationMessage-id} | Get localizedNotificationMessages from deviceManagement
[**DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages) | **Get** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages | Get localizedNotificationMessages from deviceManagement
[**DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages) | **Patch** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id}/localizedNotificationMessages/{localizedNotificationMessage-id} | Update the navigation property localizedNotificationMessages in deviceManagement
[**DeviceManagementUpdateNotificationMessageTemplates**](DeviceManagementNotificationMessageTemplateApi.md#DeviceManagementUpdateNotificationMessageTemplates) | **Patch** /deviceManagement/notificationMessageTemplates/{notificationMessageTemplate-id} | Update the navigation property notificationMessageTemplates in deviceManagement



## DeviceManagementCreateNotificationMessageTemplates

> MicrosoftGraphNotificationMessageTemplate DeviceManagementCreateNotificationMessageTemplates(ctx).MicrosoftGraphNotificationMessageTemplate(microsoftGraphNotificationMessageTemplate).Execute()

Create new navigation property to notificationMessageTemplates for deviceManagement



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
    microsoftGraphNotificationMessageTemplate := *openapiclient.NewMicrosoftGraphNotificationMessageTemplate() // MicrosoftGraphNotificationMessageTemplate | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementCreateNotificationMessageTemplates(context.Background()).MicrosoftGraphNotificationMessageTemplate(microsoftGraphNotificationMessageTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementCreateNotificationMessageTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateNotificationMessageTemplates`: MicrosoftGraphNotificationMessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementNotificationMessageTemplateApi.DeviceManagementCreateNotificationMessageTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateNotificationMessageTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphNotificationMessageTemplate** | [**MicrosoftGraphNotificationMessageTemplate**](MicrosoftGraphNotificationMessageTemplate.md) | New navigation property | 

### Return type

[**MicrosoftGraphNotificationMessageTemplate**](MicrosoftGraphNotificationMessageTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteNotificationMessageTemplates

> DeviceManagementDeleteNotificationMessageTemplates(ctx, notificationMessageTemplateId).IfMatch(ifMatch).Execute()

Delete navigation property notificationMessageTemplates for deviceManagement



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
    notificationMessageTemplateId := "notificationMessageTemplateId_example" // string | key: id of notificationMessageTemplate
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementDeleteNotificationMessageTemplates(context.Background(), notificationMessageTemplateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementDeleteNotificationMessageTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string** | key: id of notificationMessageTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteNotificationMessageTemplatesRequest struct via the builder pattern


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


## DeviceManagementGetNotificationMessageTemplates

> MicrosoftGraphNotificationMessageTemplate DeviceManagementGetNotificationMessageTemplates(ctx, notificationMessageTemplateId).Select_(select_).Expand(expand).Execute()

Get notificationMessageTemplates from deviceManagement



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
    notificationMessageTemplateId := "notificationMessageTemplateId_example" // string | key: id of notificationMessageTemplate
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementGetNotificationMessageTemplates(context.Background(), notificationMessageTemplateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementGetNotificationMessageTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetNotificationMessageTemplates`: MicrosoftGraphNotificationMessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementNotificationMessageTemplateApi.DeviceManagementGetNotificationMessageTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string** | key: id of notificationMessageTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetNotificationMessageTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphNotificationMessageTemplate**](MicrosoftGraphNotificationMessageTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListNotificationMessageTemplates

> CollectionOfNotificationMessageTemplate DeviceManagementListNotificationMessageTemplates(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get notificationMessageTemplates from deviceManagement



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
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementListNotificationMessageTemplates(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementListNotificationMessageTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListNotificationMessageTemplates`: CollectionOfNotificationMessageTemplate
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementNotificationMessageTemplateApi.DeviceManagementListNotificationMessageTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListNotificationMessageTemplatesRequest struct via the builder pattern


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

[**CollectionOfNotificationMessageTemplate**](CollectionOfNotificationMessageTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages

> MicrosoftGraphLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages(ctx, notificationMessageTemplateId).MicrosoftGraphLocalizedNotificationMessage(microsoftGraphLocalizedNotificationMessage).Execute()

Create new navigation property to localizedNotificationMessages for deviceManagement



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
    notificationMessageTemplateId := "notificationMessageTemplateId_example" // string | key: id of notificationMessageTemplate
    microsoftGraphLocalizedNotificationMessage := *openapiclient.NewMicrosoftGraphLocalizedNotificationMessage() // MicrosoftGraphLocalizedNotificationMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages(context.Background(), notificationMessageTemplateId).MicrosoftGraphLocalizedNotificationMessage(microsoftGraphLocalizedNotificationMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages`: MicrosoftGraphLocalizedNotificationMessage
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string** | key: id of notificationMessageTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementNotificationMessageTemplatesCreateLocalizedNotificationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphLocalizedNotificationMessage** | [**MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesDeleteLocalizedNotificationMessages

> DeviceManagementNotificationMessageTemplatesDeleteLocalizedNotificationMessages(ctx, notificationMessageTemplateId, localizedNotificationMessageId).IfMatch(ifMatch).Execute()

Delete navigation property localizedNotificationMessages for deviceManagement



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
    notificationMessageTemplateId := "notificationMessageTemplateId_example" // string | key: id of notificationMessageTemplate
    localizedNotificationMessageId := "localizedNotificationMessageId_example" // string | key: id of localizedNotificationMessage
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesDeleteLocalizedNotificationMessages(context.Background(), notificationMessageTemplateId, localizedNotificationMessageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesDeleteLocalizedNotificationMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string** | key: id of notificationMessageTemplate | 
**localizedNotificationMessageId** | **string** | key: id of localizedNotificationMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementNotificationMessageTemplatesDeleteLocalizedNotificationMessagesRequest struct via the builder pattern


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


## DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages

> MicrosoftGraphLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages(ctx, notificationMessageTemplateId, localizedNotificationMessageId).Select_(select_).Expand(expand).Execute()

Get localizedNotificationMessages from deviceManagement



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
    notificationMessageTemplateId := "notificationMessageTemplateId_example" // string | key: id of notificationMessageTemplate
    localizedNotificationMessageId := "localizedNotificationMessageId_example" // string | key: id of localizedNotificationMessage
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages(context.Background(), notificationMessageTemplateId, localizedNotificationMessageId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages`: MicrosoftGraphLocalizedNotificationMessage
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string** | key: id of notificationMessageTemplate | 
**localizedNotificationMessageId** | **string** | key: id of localizedNotificationMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementNotificationMessageTemplatesGetLocalizedNotificationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages

> CollectionOfLocalizedNotificationMessage DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages(ctx, notificationMessageTemplateId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get localizedNotificationMessages from deviceManagement



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
    notificationMessageTemplateId := "notificationMessageTemplateId_example" // string | key: id of notificationMessageTemplate
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
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages(context.Background(), notificationMessageTemplateId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages`: CollectionOfLocalizedNotificationMessage
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string** | key: id of notificationMessageTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementNotificationMessageTemplatesListLocalizedNotificationMessagesRequest struct via the builder pattern


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

[**CollectionOfLocalizedNotificationMessage**](CollectionOfLocalizedNotificationMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages

> DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages(ctx, notificationMessageTemplateId, localizedNotificationMessageId).MicrosoftGraphLocalizedNotificationMessage(microsoftGraphLocalizedNotificationMessage).Execute()

Update the navigation property localizedNotificationMessages in deviceManagement



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
    notificationMessageTemplateId := "notificationMessageTemplateId_example" // string | key: id of notificationMessageTemplate
    localizedNotificationMessageId := "localizedNotificationMessageId_example" // string | key: id of localizedNotificationMessage
    microsoftGraphLocalizedNotificationMessage := *openapiclient.NewMicrosoftGraphLocalizedNotificationMessage() // MicrosoftGraphLocalizedNotificationMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages(context.Background(), notificationMessageTemplateId, localizedNotificationMessageId).MicrosoftGraphLocalizedNotificationMessage(microsoftGraphLocalizedNotificationMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string** | key: id of notificationMessageTemplate | 
**localizedNotificationMessageId** | **string** | key: id of localizedNotificationMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementNotificationMessageTemplatesUpdateLocalizedNotificationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphLocalizedNotificationMessage** | [**MicrosoftGraphLocalizedNotificationMessage**](MicrosoftGraphLocalizedNotificationMessage.md) | New navigation property values | 

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


## DeviceManagementUpdateNotificationMessageTemplates

> DeviceManagementUpdateNotificationMessageTemplates(ctx, notificationMessageTemplateId).MicrosoftGraphNotificationMessageTemplate(microsoftGraphNotificationMessageTemplate).Execute()

Update the navigation property notificationMessageTemplates in deviceManagement



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
    notificationMessageTemplateId := "notificationMessageTemplateId_example" // string | key: id of notificationMessageTemplate
    microsoftGraphNotificationMessageTemplate := *openapiclient.NewMicrosoftGraphNotificationMessageTemplate() // MicrosoftGraphNotificationMessageTemplate | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementNotificationMessageTemplateApi.DeviceManagementUpdateNotificationMessageTemplates(context.Background(), notificationMessageTemplateId).MicrosoftGraphNotificationMessageTemplate(microsoftGraphNotificationMessageTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementNotificationMessageTemplateApi.DeviceManagementUpdateNotificationMessageTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationMessageTemplateId** | **string** | key: id of notificationMessageTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateNotificationMessageTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphNotificationMessageTemplate** | [**MicrosoftGraphNotificationMessageTemplate**](MicrosoftGraphNotificationMessageTemplate.md) | New navigation property values | 

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

