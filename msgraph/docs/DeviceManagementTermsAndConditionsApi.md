# \DeviceManagementTermsAndConditionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementCreateTermsAndConditions) | **Post** /deviceManagement/termsAndConditions | Create new navigation property to termsAndConditions for deviceManagement
[**DeviceManagementDeleteTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementDeleteTermsAndConditions) | **Delete** /deviceManagement/termsAndConditions/{termsAndConditions-id} | Delete navigation property termsAndConditions for deviceManagement
[**DeviceManagementGetTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementGetTermsAndConditions) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id} | Get termsAndConditions from deviceManagement
[**DeviceManagementListTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementListTermsAndConditions) | **Get** /deviceManagement/termsAndConditions | Get termsAndConditions from deviceManagement
[**DeviceManagementTermsAndConditionsAcceptanceStatusesDeleteRefTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsAcceptanceStatusesDeleteRefTermsAndConditions) | **Delete** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id}/termsAndConditions/$ref | Delete ref of navigation property termsAndConditions for deviceManagement
[**DeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditions) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id}/termsAndConditions/$ref | Get ref of termsAndConditions from deviceManagement
[**DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id}/termsAndConditions | Get termsAndConditions from deviceManagement
[**DeviceManagementTermsAndConditionsAcceptanceStatusesUpdateRefTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsAcceptanceStatusesUpdateRefTermsAndConditions) | **Put** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id}/termsAndConditions/$ref | Update the ref of navigation property termsAndConditions in deviceManagement
[**DeviceManagementTermsAndConditionsCreateAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsCreateAcceptanceStatuses) | **Post** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses | Create new navigation property to acceptanceStatuses for deviceManagement
[**DeviceManagementTermsAndConditionsCreateAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsCreateAssignments) | **Post** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments | Create new navigation property to assignments for deviceManagement
[**DeviceManagementTermsAndConditionsDeleteAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsDeleteAcceptanceStatuses) | **Delete** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id} | Delete navigation property acceptanceStatuses for deviceManagement
[**DeviceManagementTermsAndConditionsDeleteAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsDeleteAssignments) | **Delete** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments/{termsAndConditionsAssignment-id} | Delete navigation property assignments for deviceManagement
[**DeviceManagementTermsAndConditionsGetAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsGetAcceptanceStatuses) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id} | Get acceptanceStatuses from deviceManagement
[**DeviceManagementTermsAndConditionsGetAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsGetAssignments) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments/{termsAndConditionsAssignment-id} | Get assignments from deviceManagement
[**DeviceManagementTermsAndConditionsListAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsListAcceptanceStatuses) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses | Get acceptanceStatuses from deviceManagement
[**DeviceManagementTermsAndConditionsListAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsListAssignments) | **Get** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments | Get assignments from deviceManagement
[**DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses) | **Patch** /deviceManagement/termsAndConditions/{termsAndConditions-id}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus-id} | Update the navigation property acceptanceStatuses in deviceManagement
[**DeviceManagementTermsAndConditionsUpdateAssignments**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementTermsAndConditionsUpdateAssignments) | **Patch** /deviceManagement/termsAndConditions/{termsAndConditions-id}/assignments/{termsAndConditionsAssignment-id} | Update the navigation property assignments in deviceManagement
[**DeviceManagementUpdateTermsAndConditions**](DeviceManagementTermsAndConditionsApi.md#DeviceManagementUpdateTermsAndConditions) | **Patch** /deviceManagement/termsAndConditions/{termsAndConditions-id} | Update the navigation property termsAndConditions in deviceManagement



## DeviceManagementCreateTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementCreateTermsAndConditions(ctx).MicrosoftGraphTermsAndConditions(microsoftGraphTermsAndConditions).Execute()

Create new navigation property to termsAndConditions for deviceManagement



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
    microsoftGraphTermsAndConditions := *openapiclient.NewMicrosoftGraphTermsAndConditions() // MicrosoftGraphTermsAndConditions | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementCreateTermsAndConditions(context.Background()).MicrosoftGraphTermsAndConditions(microsoftGraphTermsAndConditions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementCreateTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateTermsAndConditions`: MicrosoftGraphTermsAndConditions
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementCreateTermsAndConditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateTermsAndConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTermsAndConditions** | [**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md) | New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteTermsAndConditions

> DeviceManagementDeleteTermsAndConditions(ctx, termsAndConditionsId).IfMatch(ifMatch).Execute()

Delete navigation property termsAndConditions for deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementDeleteTermsAndConditions(context.Background(), termsAndConditionsId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementDeleteTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteTermsAndConditionsRequest struct via the builder pattern


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


## DeviceManagementGetTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementGetTermsAndConditions(ctx, termsAndConditionsId).Select_(select_).Expand(expand).Execute()

Get termsAndConditions from deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementGetTermsAndConditions(context.Background(), termsAndConditionsId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementGetTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetTermsAndConditions`: MicrosoftGraphTermsAndConditions
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementGetTermsAndConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetTermsAndConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListTermsAndConditions

> CollectionOfTermsAndConditions DeviceManagementListTermsAndConditions(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get termsAndConditions from deviceManagement



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
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementListTermsAndConditions(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementListTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListTermsAndConditions`: CollectionOfTermsAndConditions
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementListTermsAndConditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListTermsAndConditionsRequest struct via the builder pattern


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

[**CollectionOfTermsAndConditions**](CollectionOfTermsAndConditions.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsAcceptanceStatusesDeleteRefTermsAndConditions

> DeviceManagementTermsAndConditionsAcceptanceStatusesDeleteRefTermsAndConditions(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId).IfMatch(ifMatch).Execute()

Delete ref of navigation property termsAndConditions for deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAcceptanceStatusId := "termsAndConditionsAcceptanceStatusId_example" // string | key: id of termsAndConditionsAcceptanceStatus
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesDeleteRefTermsAndConditions(context.Background(), termsAndConditionsId, termsAndConditionsAcceptanceStatusId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesDeleteRefTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string** | key: id of termsAndConditionsAcceptanceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsAcceptanceStatusesDeleteRefTermsAndConditionsRequest struct via the builder pattern


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


## DeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditions

> string DeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditions(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId).Execute()

Get ref of termsAndConditions from deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAcceptanceStatusId := "termsAndConditionsAcceptanceStatusId_example" // string | key: id of termsAndConditionsAcceptanceStatus

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditions(context.Background(), termsAndConditionsId, termsAndConditionsAcceptanceStatusId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditions`: string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string** | key: id of termsAndConditionsAcceptanceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsAcceptanceStatusesGetRefTermsAndConditionsRequest struct via the builder pattern


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


## DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions

> MicrosoftGraphTermsAndConditions DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId).Select_(select_).Expand(expand).Execute()

Get termsAndConditions from deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAcceptanceStatusId := "termsAndConditionsAcceptanceStatusId_example" // string | key: id of termsAndConditionsAcceptanceStatus
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions(context.Background(), termsAndConditionsId, termsAndConditionsAcceptanceStatusId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions`: MicrosoftGraphTermsAndConditions
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string** | key: id of termsAndConditionsAcceptanceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsAcceptanceStatusesUpdateRefTermsAndConditions

> DeviceManagementTermsAndConditionsAcceptanceStatusesUpdateRefTermsAndConditions(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId).RequestBody(requestBody).Execute()

Update the ref of navigation property termsAndConditions in deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAcceptanceStatusId := "termsAndConditionsAcceptanceStatusId_example" // string | key: id of termsAndConditionsAcceptanceStatus
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesUpdateRefTermsAndConditions(context.Background(), termsAndConditionsId, termsAndConditionsAcceptanceStatusId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsAcceptanceStatusesUpdateRefTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string** | key: id of termsAndConditionsAcceptanceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsAcceptanceStatusesUpdateRefTermsAndConditionsRequest struct via the builder pattern


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


## DeviceManagementTermsAndConditionsCreateAcceptanceStatuses

> MicrosoftGraphTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsCreateAcceptanceStatuses(ctx, termsAndConditionsId).MicrosoftGraphTermsAndConditionsAcceptanceStatus(microsoftGraphTermsAndConditionsAcceptanceStatus).Execute()

Create new navigation property to acceptanceStatuses for deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    microsoftGraphTermsAndConditionsAcceptanceStatus := *openapiclient.NewMicrosoftGraphTermsAndConditionsAcceptanceStatus() // MicrosoftGraphTermsAndConditionsAcceptanceStatus | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsCreateAcceptanceStatuses(context.Background(), termsAndConditionsId).MicrosoftGraphTermsAndConditionsAcceptanceStatus(microsoftGraphTermsAndConditionsAcceptanceStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsCreateAcceptanceStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementTermsAndConditionsCreateAcceptanceStatuses`: MicrosoftGraphTermsAndConditionsAcceptanceStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsCreateAcceptanceStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsCreateAcceptanceStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTermsAndConditionsAcceptanceStatus** | [**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md) | New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsCreateAssignments

> MicrosoftGraphTermsAndConditionsAssignment DeviceManagementTermsAndConditionsCreateAssignments(ctx, termsAndConditionsId).MicrosoftGraphTermsAndConditionsAssignment(microsoftGraphTermsAndConditionsAssignment).Execute()

Create new navigation property to assignments for deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    microsoftGraphTermsAndConditionsAssignment := *openapiclient.NewMicrosoftGraphTermsAndConditionsAssignment() // MicrosoftGraphTermsAndConditionsAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsCreateAssignments(context.Background(), termsAndConditionsId).MicrosoftGraphTermsAndConditionsAssignment(microsoftGraphTermsAndConditionsAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsCreateAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementTermsAndConditionsCreateAssignments`: MicrosoftGraphTermsAndConditionsAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsCreateAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsCreateAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTermsAndConditionsAssignment** | [**MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md) | New navigation property | 

### Return type

[**MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsDeleteAcceptanceStatuses

> DeviceManagementTermsAndConditionsDeleteAcceptanceStatuses(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId).IfMatch(ifMatch).Execute()

Delete navigation property acceptanceStatuses for deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAcceptanceStatusId := "termsAndConditionsAcceptanceStatusId_example" // string | key: id of termsAndConditionsAcceptanceStatus
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsDeleteAcceptanceStatuses(context.Background(), termsAndConditionsId, termsAndConditionsAcceptanceStatusId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsDeleteAcceptanceStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string** | key: id of termsAndConditionsAcceptanceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsDeleteAcceptanceStatusesRequest struct via the builder pattern


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


## DeviceManagementTermsAndConditionsDeleteAssignments

> DeviceManagementTermsAndConditionsDeleteAssignments(ctx, termsAndConditionsId, termsAndConditionsAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property assignments for deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAssignmentId := "termsAndConditionsAssignmentId_example" // string | key: id of termsAndConditionsAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsDeleteAssignments(context.Background(), termsAndConditionsId, termsAndConditionsAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsDeleteAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAssignmentId** | **string** | key: id of termsAndConditionsAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsDeleteAssignmentsRequest struct via the builder pattern


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


## DeviceManagementTermsAndConditionsGetAcceptanceStatuses

> MicrosoftGraphTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsGetAcceptanceStatuses(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId).Select_(select_).Expand(expand).Execute()

Get acceptanceStatuses from deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAcceptanceStatusId := "termsAndConditionsAcceptanceStatusId_example" // string | key: id of termsAndConditionsAcceptanceStatus
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsGetAcceptanceStatuses(context.Background(), termsAndConditionsId, termsAndConditionsAcceptanceStatusId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsGetAcceptanceStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementTermsAndConditionsGetAcceptanceStatuses`: MicrosoftGraphTermsAndConditionsAcceptanceStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsGetAcceptanceStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string** | key: id of termsAndConditionsAcceptanceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsGetAcceptanceStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsGetAssignments

> MicrosoftGraphTermsAndConditionsAssignment DeviceManagementTermsAndConditionsGetAssignments(ctx, termsAndConditionsId, termsAndConditionsAssignmentId).Select_(select_).Expand(expand).Execute()

Get assignments from deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAssignmentId := "termsAndConditionsAssignmentId_example" // string | key: id of termsAndConditionsAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsGetAssignments(context.Background(), termsAndConditionsId, termsAndConditionsAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsGetAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementTermsAndConditionsGetAssignments`: MicrosoftGraphTermsAndConditionsAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsGetAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAssignmentId** | **string** | key: id of termsAndConditionsAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsGetAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsListAcceptanceStatuses

> CollectionOfTermsAndConditionsAcceptanceStatus DeviceManagementTermsAndConditionsListAcceptanceStatuses(ctx, termsAndConditionsId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get acceptanceStatuses from deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
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
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsListAcceptanceStatuses(context.Background(), termsAndConditionsId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsListAcceptanceStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementTermsAndConditionsListAcceptanceStatuses`: CollectionOfTermsAndConditionsAcceptanceStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsListAcceptanceStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsListAcceptanceStatusesRequest struct via the builder pattern


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

[**CollectionOfTermsAndConditionsAcceptanceStatus**](CollectionOfTermsAndConditionsAcceptanceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsListAssignments

> CollectionOfTermsAndConditionsAssignment DeviceManagementTermsAndConditionsListAssignments(ctx, termsAndConditionsId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get assignments from deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
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
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsListAssignments(context.Background(), termsAndConditionsId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsListAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementTermsAndConditionsListAssignments`: CollectionOfTermsAndConditionsAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsListAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsListAssignmentsRequest struct via the builder pattern


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

[**CollectionOfTermsAndConditionsAssignment**](CollectionOfTermsAndConditionsAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses

> DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses(ctx, termsAndConditionsId, termsAndConditionsAcceptanceStatusId).MicrosoftGraphTermsAndConditionsAcceptanceStatus(microsoftGraphTermsAndConditionsAcceptanceStatus).Execute()

Update the navigation property acceptanceStatuses in deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAcceptanceStatusId := "termsAndConditionsAcceptanceStatusId_example" // string | key: id of termsAndConditionsAcceptanceStatus
    microsoftGraphTermsAndConditionsAcceptanceStatus := *openapiclient.NewMicrosoftGraphTermsAndConditionsAcceptanceStatus() // MicrosoftGraphTermsAndConditionsAcceptanceStatus | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses(context.Background(), termsAndConditionsId, termsAndConditionsAcceptanceStatusId).MicrosoftGraphTermsAndConditionsAcceptanceStatus(microsoftGraphTermsAndConditionsAcceptanceStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsUpdateAcceptanceStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAcceptanceStatusId** | **string** | key: id of termsAndConditionsAcceptanceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsUpdateAcceptanceStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTermsAndConditionsAcceptanceStatus** | [**MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md) | New navigation property values | 

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


## DeviceManagementTermsAndConditionsUpdateAssignments

> DeviceManagementTermsAndConditionsUpdateAssignments(ctx, termsAndConditionsId, termsAndConditionsAssignmentId).MicrosoftGraphTermsAndConditionsAssignment(microsoftGraphTermsAndConditionsAssignment).Execute()

Update the navigation property assignments in deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    termsAndConditionsAssignmentId := "termsAndConditionsAssignmentId_example" // string | key: id of termsAndConditionsAssignment
    microsoftGraphTermsAndConditionsAssignment := *openapiclient.NewMicrosoftGraphTermsAndConditionsAssignment() // MicrosoftGraphTermsAndConditionsAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsUpdateAssignments(context.Background(), termsAndConditionsId, termsAndConditionsAssignmentId).MicrosoftGraphTermsAndConditionsAssignment(microsoftGraphTermsAndConditionsAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementTermsAndConditionsUpdateAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 
**termsAndConditionsAssignmentId** | **string** | key: id of termsAndConditionsAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementTermsAndConditionsUpdateAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTermsAndConditionsAssignment** | [**MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md) | New navigation property values | 

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


## DeviceManagementUpdateTermsAndConditions

> DeviceManagementUpdateTermsAndConditions(ctx, termsAndConditionsId).MicrosoftGraphTermsAndConditions(microsoftGraphTermsAndConditions).Execute()

Update the navigation property termsAndConditions in deviceManagement



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
    termsAndConditionsId := "termsAndConditionsId_example" // string | key: id of termsAndConditions
    microsoftGraphTermsAndConditions := *openapiclient.NewMicrosoftGraphTermsAndConditions() // MicrosoftGraphTermsAndConditions | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementTermsAndConditionsApi.DeviceManagementUpdateTermsAndConditions(context.Background(), termsAndConditionsId).MicrosoftGraphTermsAndConditions(microsoftGraphTermsAndConditions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementTermsAndConditionsApi.DeviceManagementUpdateTermsAndConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**termsAndConditionsId** | **string** | key: id of termsAndConditions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateTermsAndConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTermsAndConditions** | [**MicrosoftGraphTermsAndConditions**](MicrosoftGraphTermsAndConditions.md) | New navigation property values | 

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

