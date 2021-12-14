# \PrintPrintServiceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrintCreateServices**](PrintPrintServiceApi.md#PrintCreateServices) | **Post** /print/services | Create new navigation property to services for print
[**PrintDeleteServices**](PrintPrintServiceApi.md#PrintDeleteServices) | **Delete** /print/services/{printService-id} | Delete navigation property services for print
[**PrintGetServices**](PrintPrintServiceApi.md#PrintGetServices) | **Get** /print/services/{printService-id} | Get services from print
[**PrintListServices**](PrintPrintServiceApi.md#PrintListServices) | **Get** /print/services | Get services from print
[**PrintServicesCreateEndpoints**](PrintPrintServiceApi.md#PrintServicesCreateEndpoints) | **Post** /print/services/{printService-id}/endpoints | Create new navigation property to endpoints for print
[**PrintServicesDeleteEndpoints**](PrintPrintServiceApi.md#PrintServicesDeleteEndpoints) | **Delete** /print/services/{printService-id}/endpoints/{printServiceEndpoint-id} | Delete navigation property endpoints for print
[**PrintServicesGetEndpoints**](PrintPrintServiceApi.md#PrintServicesGetEndpoints) | **Get** /print/services/{printService-id}/endpoints/{printServiceEndpoint-id} | Get endpoints from print
[**PrintServicesListEndpoints**](PrintPrintServiceApi.md#PrintServicesListEndpoints) | **Get** /print/services/{printService-id}/endpoints | Get endpoints from print
[**PrintServicesUpdateEndpoints**](PrintPrintServiceApi.md#PrintServicesUpdateEndpoints) | **Patch** /print/services/{printService-id}/endpoints/{printServiceEndpoint-id} | Update the navigation property endpoints in print
[**PrintUpdateServices**](PrintPrintServiceApi.md#PrintUpdateServices) | **Patch** /print/services/{printService-id} | Update the navigation property services in print



## PrintCreateServices

> MicrosoftGraphPrintService PrintCreateServices(ctx).MicrosoftGraphPrintService(microsoftGraphPrintService).Execute()

Create new navigation property to services for print



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
    microsoftGraphPrintService := *openapiclient.NewMicrosoftGraphPrintService() // MicrosoftGraphPrintService | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintServiceApi.PrintCreateServices(context.Background()).MicrosoftGraphPrintService(microsoftGraphPrintService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintCreateServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintCreateServices`: MicrosoftGraphPrintService
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintServiceApi.PrintCreateServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintCreateServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrintService** | [**MicrosoftGraphPrintService**](MicrosoftGraphPrintService.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintService**](MicrosoftGraphPrintService.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintDeleteServices

> PrintDeleteServices(ctx, printServiceId).IfMatch(ifMatch).Execute()

Delete navigation property services for print



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
    printServiceId := "printServiceId_example" // string | key: id of printService
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintServiceApi.PrintDeleteServices(context.Background(), printServiceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintDeleteServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printServiceId** | **string** | key: id of printService | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintDeleteServicesRequest struct via the builder pattern


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


## PrintGetServices

> MicrosoftGraphPrintService PrintGetServices(ctx, printServiceId).Select_(select_).Expand(expand).Execute()

Get services from print



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
    printServiceId := "printServiceId_example" // string | key: id of printService
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintServiceApi.PrintGetServices(context.Background(), printServiceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintGetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintGetServices`: MicrosoftGraphPrintService
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintServiceApi.PrintGetServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printServiceId** | **string** | key: id of printService | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintService**](MicrosoftGraphPrintService.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintListServices

> CollectionOfPrintService PrintListServices(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get services from print



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
    resp, r, err := api_client.PrintPrintServiceApi.PrintListServices(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintListServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintListServices`: CollectionOfPrintService
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintServiceApi.PrintListServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintListServicesRequest struct via the builder pattern


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

[**CollectionOfPrintService**](CollectionOfPrintService.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintServicesCreateEndpoints

> MicrosoftGraphPrintServiceEndpoint PrintServicesCreateEndpoints(ctx, printServiceId).MicrosoftGraphPrintServiceEndpoint(microsoftGraphPrintServiceEndpoint).Execute()

Create new navigation property to endpoints for print



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
    printServiceId := "printServiceId_example" // string | key: id of printService
    microsoftGraphPrintServiceEndpoint := *openapiclient.NewMicrosoftGraphPrintServiceEndpoint() // MicrosoftGraphPrintServiceEndpoint | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintServiceApi.PrintServicesCreateEndpoints(context.Background(), printServiceId).MicrosoftGraphPrintServiceEndpoint(microsoftGraphPrintServiceEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintServicesCreateEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintServicesCreateEndpoints`: MicrosoftGraphPrintServiceEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintServiceApi.PrintServicesCreateEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printServiceId** | **string** | key: id of printService | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintServicesCreateEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintServiceEndpoint** | [**MicrosoftGraphPrintServiceEndpoint**](MicrosoftGraphPrintServiceEndpoint.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintServiceEndpoint**](MicrosoftGraphPrintServiceEndpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintServicesDeleteEndpoints

> PrintServicesDeleteEndpoints(ctx, printServiceId, printServiceEndpointId).IfMatch(ifMatch).Execute()

Delete navigation property endpoints for print



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
    printServiceId := "printServiceId_example" // string | key: id of printService
    printServiceEndpointId := "printServiceEndpointId_example" // string | key: id of printServiceEndpoint
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintServiceApi.PrintServicesDeleteEndpoints(context.Background(), printServiceId, printServiceEndpointId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintServicesDeleteEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printServiceId** | **string** | key: id of printService | 
**printServiceEndpointId** | **string** | key: id of printServiceEndpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintServicesDeleteEndpointsRequest struct via the builder pattern


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


## PrintServicesGetEndpoints

> MicrosoftGraphPrintServiceEndpoint PrintServicesGetEndpoints(ctx, printServiceId, printServiceEndpointId).Select_(select_).Expand(expand).Execute()

Get endpoints from print



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
    printServiceId := "printServiceId_example" // string | key: id of printService
    printServiceEndpointId := "printServiceEndpointId_example" // string | key: id of printServiceEndpoint
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintServiceApi.PrintServicesGetEndpoints(context.Background(), printServiceId, printServiceEndpointId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintServicesGetEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintServicesGetEndpoints`: MicrosoftGraphPrintServiceEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintServiceApi.PrintServicesGetEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printServiceId** | **string** | key: id of printService | 
**printServiceEndpointId** | **string** | key: id of printServiceEndpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintServicesGetEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintServiceEndpoint**](MicrosoftGraphPrintServiceEndpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintServicesListEndpoints

> CollectionOfPrintServiceEndpoint PrintServicesListEndpoints(ctx, printServiceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get endpoints from print



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
    printServiceId := "printServiceId_example" // string | key: id of printService
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
    resp, r, err := api_client.PrintPrintServiceApi.PrintServicesListEndpoints(context.Background(), printServiceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintServicesListEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintServicesListEndpoints`: CollectionOfPrintServiceEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintServiceApi.PrintServicesListEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printServiceId** | **string** | key: id of printService | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintServicesListEndpointsRequest struct via the builder pattern


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

[**CollectionOfPrintServiceEndpoint**](CollectionOfPrintServiceEndpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintServicesUpdateEndpoints

> PrintServicesUpdateEndpoints(ctx, printServiceId, printServiceEndpointId).MicrosoftGraphPrintServiceEndpoint(microsoftGraphPrintServiceEndpoint).Execute()

Update the navigation property endpoints in print



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
    printServiceId := "printServiceId_example" // string | key: id of printService
    printServiceEndpointId := "printServiceEndpointId_example" // string | key: id of printServiceEndpoint
    microsoftGraphPrintServiceEndpoint := *openapiclient.NewMicrosoftGraphPrintServiceEndpoint() // MicrosoftGraphPrintServiceEndpoint | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintServiceApi.PrintServicesUpdateEndpoints(context.Background(), printServiceId, printServiceEndpointId).MicrosoftGraphPrintServiceEndpoint(microsoftGraphPrintServiceEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintServicesUpdateEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printServiceId** | **string** | key: id of printService | 
**printServiceEndpointId** | **string** | key: id of printServiceEndpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintServicesUpdateEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphPrintServiceEndpoint** | [**MicrosoftGraphPrintServiceEndpoint**](MicrosoftGraphPrintServiceEndpoint.md) | New navigation property values | 

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


## PrintUpdateServices

> PrintUpdateServices(ctx, printServiceId).MicrosoftGraphPrintService(microsoftGraphPrintService).Execute()

Update the navigation property services in print



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
    printServiceId := "printServiceId_example" // string | key: id of printService
    microsoftGraphPrintService := *openapiclient.NewMicrosoftGraphPrintService() // MicrosoftGraphPrintService | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintServiceApi.PrintUpdateServices(context.Background(), printServiceId).MicrosoftGraphPrintService(microsoftGraphPrintService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintServiceApi.PrintUpdateServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printServiceId** | **string** | key: id of printService | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintUpdateServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintService** | [**MicrosoftGraphPrintService**](MicrosoftGraphPrintService.md) | New navigation property values | 

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

