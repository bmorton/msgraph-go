# \ConnectionsExternalGroupApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionsCreateGroups**](ConnectionsExternalGroupApi.md#ConnectionsCreateGroups) | **Post** /connections/{externalConnection-id}/groups | Create new navigation property to groups for connections
[**ConnectionsDeleteGroups**](ConnectionsExternalGroupApi.md#ConnectionsDeleteGroups) | **Delete** /connections/{externalConnection-id}/groups/{externalGroup-id} | Delete navigation property groups for connections
[**ConnectionsGetGroups**](ConnectionsExternalGroupApi.md#ConnectionsGetGroups) | **Get** /connections/{externalConnection-id}/groups/{externalGroup-id} | Get groups from connections
[**ConnectionsGroupsCreateMembers**](ConnectionsExternalGroupApi.md#ConnectionsGroupsCreateMembers) | **Post** /connections/{externalConnection-id}/groups/{externalGroup-id}/members | Create new navigation property to members for connections
[**ConnectionsGroupsDeleteMembers**](ConnectionsExternalGroupApi.md#ConnectionsGroupsDeleteMembers) | **Delete** /connections/{externalConnection-id}/groups/{externalGroup-id}/members/{identity-id} | Delete navigation property members for connections
[**ConnectionsGroupsGetMembers**](ConnectionsExternalGroupApi.md#ConnectionsGroupsGetMembers) | **Get** /connections/{externalConnection-id}/groups/{externalGroup-id}/members/{identity-id} | Get members from connections
[**ConnectionsGroupsListMembers**](ConnectionsExternalGroupApi.md#ConnectionsGroupsListMembers) | **Get** /connections/{externalConnection-id}/groups/{externalGroup-id}/members | Get members from connections
[**ConnectionsGroupsUpdateMembers**](ConnectionsExternalGroupApi.md#ConnectionsGroupsUpdateMembers) | **Patch** /connections/{externalConnection-id}/groups/{externalGroup-id}/members/{identity-id} | Update the navigation property members in connections
[**ConnectionsListGroups**](ConnectionsExternalGroupApi.md#ConnectionsListGroups) | **Get** /connections/{externalConnection-id}/groups | Get groups from connections
[**ConnectionsUpdateGroups**](ConnectionsExternalGroupApi.md#ConnectionsUpdateGroups) | **Patch** /connections/{externalConnection-id}/groups/{externalGroup-id} | Update the navigation property groups in connections



## ConnectionsCreateGroups

> MicrosoftGraphExternalConnectorsExternalGroup ConnectionsCreateGroups(ctx, externalConnectionId).MicrosoftGraphExternalConnectorsExternalGroup(microsoftGraphExternalConnectorsExternalGroup).Execute()

Create new navigation property to groups for connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    microsoftGraphExternalConnectorsExternalGroup := *openapiclient.NewMicrosoftGraphExternalConnectorsExternalGroup() // MicrosoftGraphExternalConnectorsExternalGroup | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsCreateGroups(context.Background(), externalConnectionId).MicrosoftGraphExternalConnectorsExternalGroup(microsoftGraphExternalConnectorsExternalGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsCreateGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsCreateGroups`: MicrosoftGraphExternalConnectorsExternalGroup
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalGroupApi.ConnectionsCreateGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsCreateGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExternalConnectorsExternalGroup** | [**MicrosoftGraphExternalConnectorsExternalGroup**](MicrosoftGraphExternalConnectorsExternalGroup.md) | New navigation property | 

### Return type

[**MicrosoftGraphExternalConnectorsExternalGroup**](MicrosoftGraphExternalConnectorsExternalGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsDeleteGroups

> ConnectionsDeleteGroups(ctx, externalConnectionId, externalGroupId).IfMatch(ifMatch).Execute()

Delete navigation property groups for connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalGroupId := "externalGroupId_example" // string | key: id of externalGroup
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsDeleteGroups(context.Background(), externalConnectionId, externalGroupId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsDeleteGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalGroupId** | **string** | key: id of externalGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsDeleteGroupsRequest struct via the builder pattern


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


## ConnectionsGetGroups

> MicrosoftGraphExternalConnectorsExternalGroup ConnectionsGetGroups(ctx, externalConnectionId, externalGroupId).Select_(select_).Expand(expand).Execute()

Get groups from connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalGroupId := "externalGroupId_example" // string | key: id of externalGroup
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsGetGroups(context.Background(), externalConnectionId, externalGroupId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsGetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsGetGroups`: MicrosoftGraphExternalConnectorsExternalGroup
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalGroupApi.ConnectionsGetGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalGroupId** | **string** | key: id of externalGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExternalConnectorsExternalGroup**](MicrosoftGraphExternalConnectorsExternalGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsGroupsCreateMembers

> MicrosoftGraphExternalConnectorsIdentity ConnectionsGroupsCreateMembers(ctx, externalConnectionId, externalGroupId).MicrosoftGraphExternalConnectorsIdentity(microsoftGraphExternalConnectorsIdentity).Execute()

Create new navigation property to members for connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalGroupId := "externalGroupId_example" // string | key: id of externalGroup
    microsoftGraphExternalConnectorsIdentity := *openapiclient.NewMicrosoftGraphExternalConnectorsIdentity() // MicrosoftGraphExternalConnectorsIdentity | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsGroupsCreateMembers(context.Background(), externalConnectionId, externalGroupId).MicrosoftGraphExternalConnectorsIdentity(microsoftGraphExternalConnectorsIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsGroupsCreateMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsGroupsCreateMembers`: MicrosoftGraphExternalConnectorsIdentity
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalGroupApi.ConnectionsGroupsCreateMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalGroupId** | **string** | key: id of externalGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGroupsCreateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphExternalConnectorsIdentity** | [**MicrosoftGraphExternalConnectorsIdentity**](MicrosoftGraphExternalConnectorsIdentity.md) | New navigation property | 

### Return type

[**MicrosoftGraphExternalConnectorsIdentity**](MicrosoftGraphExternalConnectorsIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsGroupsDeleteMembers

> ConnectionsGroupsDeleteMembers(ctx, externalConnectionId, externalGroupId, identityId).IfMatch(ifMatch).Execute()

Delete navigation property members for connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalGroupId := "externalGroupId_example" // string | key: id of externalGroup
    identityId := "identityId_example" // string | key: id of identity
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsGroupsDeleteMembers(context.Background(), externalConnectionId, externalGroupId, identityId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsGroupsDeleteMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalGroupId** | **string** | key: id of externalGroup | 
**identityId** | **string** | key: id of identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGroupsDeleteMembersRequest struct via the builder pattern


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


## ConnectionsGroupsGetMembers

> MicrosoftGraphExternalConnectorsIdentity ConnectionsGroupsGetMembers(ctx, externalConnectionId, externalGroupId, identityId).Select_(select_).Expand(expand).Execute()

Get members from connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalGroupId := "externalGroupId_example" // string | key: id of externalGroup
    identityId := "identityId_example" // string | key: id of identity
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsGroupsGetMembers(context.Background(), externalConnectionId, externalGroupId, identityId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsGroupsGetMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsGroupsGetMembers`: MicrosoftGraphExternalConnectorsIdentity
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalGroupApi.ConnectionsGroupsGetMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalGroupId** | **string** | key: id of externalGroup | 
**identityId** | **string** | key: id of identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGroupsGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExternalConnectorsIdentity**](MicrosoftGraphExternalConnectorsIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsGroupsListMembers

> CollectionOfIdentity ConnectionsGroupsListMembers(ctx, externalConnectionId, externalGroupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get members from connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalGroupId := "externalGroupId_example" // string | key: id of externalGroup
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
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsGroupsListMembers(context.Background(), externalConnectionId, externalGroupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsGroupsListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsGroupsListMembers`: CollectionOfIdentity
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalGroupApi.ConnectionsGroupsListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalGroupId** | **string** | key: id of externalGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGroupsListMembersRequest struct via the builder pattern


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

[**CollectionOfIdentity**](CollectionOfIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsGroupsUpdateMembers

> ConnectionsGroupsUpdateMembers(ctx, externalConnectionId, externalGroupId, identityId).MicrosoftGraphExternalConnectorsIdentity(microsoftGraphExternalConnectorsIdentity).Execute()

Update the navigation property members in connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalGroupId := "externalGroupId_example" // string | key: id of externalGroup
    identityId := "identityId_example" // string | key: id of identity
    microsoftGraphExternalConnectorsIdentity := *openapiclient.NewMicrosoftGraphExternalConnectorsIdentity() // MicrosoftGraphExternalConnectorsIdentity | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsGroupsUpdateMembers(context.Background(), externalConnectionId, externalGroupId, identityId).MicrosoftGraphExternalConnectorsIdentity(microsoftGraphExternalConnectorsIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsGroupsUpdateMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalGroupId** | **string** | key: id of externalGroup | 
**identityId** | **string** | key: id of identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGroupsUpdateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphExternalConnectorsIdentity** | [**MicrosoftGraphExternalConnectorsIdentity**](MicrosoftGraphExternalConnectorsIdentity.md) | New navigation property values | 

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


## ConnectionsListGroups

> CollectionOfExternalGroup ConnectionsListGroups(ctx, externalConnectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get groups from connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
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
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsListGroups(context.Background(), externalConnectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsListGroups`: CollectionOfExternalGroup
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalGroupApi.ConnectionsListGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsListGroupsRequest struct via the builder pattern


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

[**CollectionOfExternalGroup**](CollectionOfExternalGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsUpdateGroups

> ConnectionsUpdateGroups(ctx, externalConnectionId, externalGroupId).MicrosoftGraphExternalConnectorsExternalGroup(microsoftGraphExternalConnectorsExternalGroup).Execute()

Update the navigation property groups in connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalGroupId := "externalGroupId_example" // string | key: id of externalGroup
    microsoftGraphExternalConnectorsExternalGroup := *openapiclient.NewMicrosoftGraphExternalConnectorsExternalGroup() // MicrosoftGraphExternalConnectorsExternalGroup | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalGroupApi.ConnectionsUpdateGroups(context.Background(), externalConnectionId, externalGroupId).MicrosoftGraphExternalConnectorsExternalGroup(microsoftGraphExternalConnectorsExternalGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalGroupApi.ConnectionsUpdateGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalGroupId** | **string** | key: id of externalGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsUpdateGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphExternalConnectorsExternalGroup** | [**MicrosoftGraphExternalConnectorsExternalGroup**](MicrosoftGraphExternalConnectorsExternalGroup.md) | New navigation property values | 

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

