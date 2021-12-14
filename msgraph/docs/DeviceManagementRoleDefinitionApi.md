# \DeviceManagementRoleDefinitionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementCreateRoleDefinitions) | **Post** /deviceManagement/roleDefinitions | Create new navigation property to roleDefinitions for deviceManagement
[**DeviceManagementDeleteRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementDeleteRoleDefinitions) | **Delete** /deviceManagement/roleDefinitions/{roleDefinition-id} | Delete navigation property roleDefinitions for deviceManagement
[**DeviceManagementGetRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementGetRoleDefinitions) | **Get** /deviceManagement/roleDefinitions/{roleDefinition-id} | Get roleDefinitions from deviceManagement
[**DeviceManagementListRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementListRoleDefinitions) | **Get** /deviceManagement/roleDefinitions | Get roleDefinitions from deviceManagement
[**DeviceManagementRoleDefinitionsCreateRoleAssignments**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsCreateRoleAssignments) | **Post** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments | Create new navigation property to roleAssignments for deviceManagement
[**DeviceManagementRoleDefinitionsDeleteRoleAssignments**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsDeleteRoleAssignments) | **Delete** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments/{roleAssignment-id} | Delete navigation property roleAssignments for deviceManagement
[**DeviceManagementRoleDefinitionsGetRoleAssignments**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsGetRoleAssignments) | **Get** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments/{roleAssignment-id} | Get roleAssignments from deviceManagement
[**DeviceManagementRoleDefinitionsListRoleAssignments**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsListRoleAssignments) | **Get** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments | Get roleAssignments from deviceManagement
[**DeviceManagementRoleDefinitionsRoleAssignmentsDeleteRefRoleDefinition**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsRoleAssignmentsDeleteRefRoleDefinition) | **Delete** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments/{roleAssignment-id}/roleDefinition/$ref | Delete ref of navigation property roleDefinition for deviceManagement
[**DeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinition**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinition) | **Get** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments/{roleAssignment-id}/roleDefinition/$ref | Get ref of roleDefinition from deviceManagement
[**DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition) | **Get** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments/{roleAssignment-id}/roleDefinition | Get roleDefinition from deviceManagement
[**DeviceManagementRoleDefinitionsRoleAssignmentsUpdateRefRoleDefinition**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsRoleAssignmentsUpdateRefRoleDefinition) | **Put** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments/{roleAssignment-id}/roleDefinition/$ref | Update the ref of navigation property roleDefinition in deviceManagement
[**DeviceManagementRoleDefinitionsUpdateRoleAssignments**](DeviceManagementRoleDefinitionApi.md#DeviceManagementRoleDefinitionsUpdateRoleAssignments) | **Patch** /deviceManagement/roleDefinitions/{roleDefinition-id}/roleAssignments/{roleAssignment-id} | Update the navigation property roleAssignments in deviceManagement
[**DeviceManagementUpdateRoleDefinitions**](DeviceManagementRoleDefinitionApi.md#DeviceManagementUpdateRoleDefinitions) | **Patch** /deviceManagement/roleDefinitions/{roleDefinition-id} | Update the navigation property roleDefinitions in deviceManagement



## DeviceManagementCreateRoleDefinitions

> MicrosoftGraphRoleDefinition DeviceManagementCreateRoleDefinitions(ctx).MicrosoftGraphRoleDefinition(microsoftGraphRoleDefinition).Execute()

Create new navigation property to roleDefinitions for deviceManagement



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
    microsoftGraphRoleDefinition := *openapiclient.NewMicrosoftGraphRoleDefinition() // MicrosoftGraphRoleDefinition | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementCreateRoleDefinitions(context.Background()).MicrosoftGraphRoleDefinition(microsoftGraphRoleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementCreateRoleDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateRoleDefinitions`: MicrosoftGraphRoleDefinition
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRoleDefinitionApi.DeviceManagementCreateRoleDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateRoleDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphRoleDefinition** | [**MicrosoftGraphRoleDefinition**](MicrosoftGraphRoleDefinition.md) | New navigation property | 

### Return type

[**MicrosoftGraphRoleDefinition**](MicrosoftGraphRoleDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteRoleDefinitions

> DeviceManagementDeleteRoleDefinitions(ctx, roleDefinitionId).IfMatch(ifMatch).Execute()

Delete navigation property roleDefinitions for deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementDeleteRoleDefinitions(context.Background(), roleDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementDeleteRoleDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteRoleDefinitionsRequest struct via the builder pattern


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


## DeviceManagementGetRoleDefinitions

> MicrosoftGraphRoleDefinition DeviceManagementGetRoleDefinitions(ctx, roleDefinitionId).Select_(select_).Expand(expand).Execute()

Get roleDefinitions from deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementGetRoleDefinitions(context.Background(), roleDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementGetRoleDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetRoleDefinitions`: MicrosoftGraphRoleDefinition
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRoleDefinitionApi.DeviceManagementGetRoleDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetRoleDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphRoleDefinition**](MicrosoftGraphRoleDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListRoleDefinitions

> CollectionOfRoleDefinition DeviceManagementListRoleDefinitions(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get roleDefinitions from deviceManagement



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
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementListRoleDefinitions(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementListRoleDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListRoleDefinitions`: CollectionOfRoleDefinition
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRoleDefinitionApi.DeviceManagementListRoleDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListRoleDefinitionsRequest struct via the builder pattern


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

[**CollectionOfRoleDefinition**](CollectionOfRoleDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementRoleDefinitionsCreateRoleAssignments

> MicrosoftGraphRoleAssignment DeviceManagementRoleDefinitionsCreateRoleAssignments(ctx, roleDefinitionId).MicrosoftGraphRoleAssignment(microsoftGraphRoleAssignment).Execute()

Create new navigation property to roleAssignments for deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    microsoftGraphRoleAssignment := *openapiclient.NewMicrosoftGraphRoleAssignment() // MicrosoftGraphRoleAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsCreateRoleAssignments(context.Background(), roleDefinitionId).MicrosoftGraphRoleAssignment(microsoftGraphRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsCreateRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementRoleDefinitionsCreateRoleAssignments`: MicrosoftGraphRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsCreateRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsCreateRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphRoleAssignment** | [**MicrosoftGraphRoleAssignment**](MicrosoftGraphRoleAssignment.md) | New navigation property | 

### Return type

[**MicrosoftGraphRoleAssignment**](MicrosoftGraphRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementRoleDefinitionsDeleteRoleAssignments

> DeviceManagementRoleDefinitionsDeleteRoleAssignments(ctx, roleDefinitionId, roleAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property roleAssignments for deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    roleAssignmentId := "roleAssignmentId_example" // string | key: id of roleAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsDeleteRoleAssignments(context.Background(), roleDefinitionId, roleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsDeleteRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 
**roleAssignmentId** | **string** | key: id of roleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsDeleteRoleAssignmentsRequest struct via the builder pattern


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


## DeviceManagementRoleDefinitionsGetRoleAssignments

> MicrosoftGraphRoleAssignment DeviceManagementRoleDefinitionsGetRoleAssignments(ctx, roleDefinitionId, roleAssignmentId).Select_(select_).Expand(expand).Execute()

Get roleAssignments from deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    roleAssignmentId := "roleAssignmentId_example" // string | key: id of roleAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsGetRoleAssignments(context.Background(), roleDefinitionId, roleAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsGetRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementRoleDefinitionsGetRoleAssignments`: MicrosoftGraphRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsGetRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 
**roleAssignmentId** | **string** | key: id of roleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsGetRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphRoleAssignment**](MicrosoftGraphRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementRoleDefinitionsListRoleAssignments

> CollectionOfRoleAssignment DeviceManagementRoleDefinitionsListRoleAssignments(ctx, roleDefinitionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get roleAssignments from deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
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
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsListRoleAssignments(context.Background(), roleDefinitionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsListRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementRoleDefinitionsListRoleAssignments`: CollectionOfRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsListRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsListRoleAssignmentsRequest struct via the builder pattern


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

[**CollectionOfRoleAssignment**](CollectionOfRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementRoleDefinitionsRoleAssignmentsDeleteRefRoleDefinition

> DeviceManagementRoleDefinitionsRoleAssignmentsDeleteRefRoleDefinition(ctx, roleDefinitionId, roleAssignmentId).IfMatch(ifMatch).Execute()

Delete ref of navigation property roleDefinition for deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    roleAssignmentId := "roleAssignmentId_example" // string | key: id of roleAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsDeleteRefRoleDefinition(context.Background(), roleDefinitionId, roleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsDeleteRefRoleDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 
**roleAssignmentId** | **string** | key: id of roleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsRoleAssignmentsDeleteRefRoleDefinitionRequest struct via the builder pattern


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


## DeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinition

> string DeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinition(ctx, roleDefinitionId, roleAssignmentId).Execute()

Get ref of roleDefinition from deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    roleAssignmentId := "roleAssignmentId_example" // string | key: id of roleAssignment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinition(context.Background(), roleDefinitionId, roleAssignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinition`: string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 
**roleAssignmentId** | **string** | key: id of roleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsRoleAssignmentsGetRefRoleDefinitionRequest struct via the builder pattern


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


## DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition

> MicrosoftGraphRoleDefinition DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition(ctx, roleDefinitionId, roleAssignmentId).Select_(select_).Expand(expand).Execute()

Get roleDefinition from deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    roleAssignmentId := "roleAssignmentId_example" // string | key: id of roleAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition(context.Background(), roleDefinitionId, roleAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition`: MicrosoftGraphRoleDefinition
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 
**roleAssignmentId** | **string** | key: id of roleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsRoleAssignmentsGetRoleDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphRoleDefinition**](MicrosoftGraphRoleDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementRoleDefinitionsRoleAssignmentsUpdateRefRoleDefinition

> DeviceManagementRoleDefinitionsRoleAssignmentsUpdateRefRoleDefinition(ctx, roleDefinitionId, roleAssignmentId).RequestBody(requestBody).Execute()

Update the ref of navigation property roleDefinition in deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    roleAssignmentId := "roleAssignmentId_example" // string | key: id of roleAssignment
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsUpdateRefRoleDefinition(context.Background(), roleDefinitionId, roleAssignmentId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsRoleAssignmentsUpdateRefRoleDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 
**roleAssignmentId** | **string** | key: id of roleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsRoleAssignmentsUpdateRefRoleDefinitionRequest struct via the builder pattern


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


## DeviceManagementRoleDefinitionsUpdateRoleAssignments

> DeviceManagementRoleDefinitionsUpdateRoleAssignments(ctx, roleDefinitionId, roleAssignmentId).MicrosoftGraphRoleAssignment(microsoftGraphRoleAssignment).Execute()

Update the navigation property roleAssignments in deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    roleAssignmentId := "roleAssignmentId_example" // string | key: id of roleAssignment
    microsoftGraphRoleAssignment := *openapiclient.NewMicrosoftGraphRoleAssignment() // MicrosoftGraphRoleAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsUpdateRoleAssignments(context.Background(), roleDefinitionId, roleAssignmentId).MicrosoftGraphRoleAssignment(microsoftGraphRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementRoleDefinitionsUpdateRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 
**roleAssignmentId** | **string** | key: id of roleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementRoleDefinitionsUpdateRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphRoleAssignment** | [**MicrosoftGraphRoleAssignment**](MicrosoftGraphRoleAssignment.md) | New navigation property values | 

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


## DeviceManagementUpdateRoleDefinitions

> DeviceManagementUpdateRoleDefinitions(ctx, roleDefinitionId).MicrosoftGraphRoleDefinition(microsoftGraphRoleDefinition).Execute()

Update the navigation property roleDefinitions in deviceManagement



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
    roleDefinitionId := "roleDefinitionId_example" // string | key: id of roleDefinition
    microsoftGraphRoleDefinition := *openapiclient.NewMicrosoftGraphRoleDefinition() // MicrosoftGraphRoleDefinition | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementRoleDefinitionApi.DeviceManagementUpdateRoleDefinitions(context.Background(), roleDefinitionId).MicrosoftGraphRoleDefinition(microsoftGraphRoleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementRoleDefinitionApi.DeviceManagementUpdateRoleDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleDefinitionId** | **string** | key: id of roleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateRoleDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphRoleDefinition** | [**MicrosoftGraphRoleDefinition**](MicrosoftGraphRoleDefinition.md) | New navigation property values | 

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

