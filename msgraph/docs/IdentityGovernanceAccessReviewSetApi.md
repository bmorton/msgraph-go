# \IdentityGovernanceAccessReviewSetApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityGovernanceAccessReviewsCreateDefinitions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsCreateDefinitions) | **Post** /identityGovernance/accessReviews/definitions | Create new navigation property to definitions for identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsCreateInstances**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsCreateInstances) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances | Create new navigation property to instances for identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsDeleteInstances**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsDeleteInstances) | **Delete** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id} | Delete navigation property instances for identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsGetInstances**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsGetInstances) | **Get** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id} | Get instances from identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisions) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/decisions | Create new navigation property to decisions for identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsInstancesDeleteDecisions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsInstancesDeleteDecisions) | **Delete** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/decisions/{accessReviewInstanceDecisionItem-id} | Delete navigation property decisions for identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisions) | **Get** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/decisions/{accessReviewInstanceDecisionItem-id} | Get decisions from identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsInstancesListDecisions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsInstancesListDecisions) | **Get** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/decisions | Get decisions from identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsInstancesUpdateDecisions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsInstancesUpdateDecisions) | **Patch** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/decisions/{accessReviewInstanceDecisionItem-id} | Update the navigation property decisions in identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsListInstances**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsListInstances) | **Get** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances | Get instances from identityGovernance
[**IdentityGovernanceAccessReviewsDefinitionsUpdateInstances**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDefinitionsUpdateInstances) | **Patch** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id} | Update the navigation property instances in identityGovernance
[**IdentityGovernanceAccessReviewsDeleteDefinitions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsDeleteDefinitions) | **Delete** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id} | Delete navigation property definitions for identityGovernance
[**IdentityGovernanceAccessReviewsGetDefinitions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsGetDefinitions) | **Get** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id} | Get definitions from identityGovernance
[**IdentityGovernanceAccessReviewsListDefinitions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsListDefinitions) | **Get** /identityGovernance/accessReviews/definitions | Get definitions from identityGovernance
[**IdentityGovernanceAccessReviewsUpdateDefinitions**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceAccessReviewsUpdateDefinitions) | **Patch** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id} | Update the navigation property definitions in identityGovernance
[**IdentityGovernanceDeleteAccessReviews**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceDeleteAccessReviews) | **Delete** /identityGovernance/accessReviews | Delete navigation property accessReviews for identityGovernance
[**IdentityGovernanceGetAccessReviews**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceGetAccessReviews) | **Get** /identityGovernance/accessReviews | Get accessReviews from identityGovernance
[**IdentityGovernanceUpdateAccessReviews**](IdentityGovernanceAccessReviewSetApi.md#IdentityGovernanceUpdateAccessReviews) | **Patch** /identityGovernance/accessReviews | Update the navigation property accessReviews in identityGovernance



## IdentityGovernanceAccessReviewsCreateDefinitions

> MicrosoftGraphAccessReviewScheduleDefinition IdentityGovernanceAccessReviewsCreateDefinitions(ctx).MicrosoftGraphAccessReviewScheduleDefinition(microsoftGraphAccessReviewScheduleDefinition).Execute()

Create new navigation property to definitions for identityGovernance

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
    microsoftGraphAccessReviewScheduleDefinition := *openapiclient.NewMicrosoftGraphAccessReviewScheduleDefinition() // MicrosoftGraphAccessReviewScheduleDefinition | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsCreateDefinitions(context.Background()).MicrosoftGraphAccessReviewScheduleDefinition(microsoftGraphAccessReviewScheduleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsCreateDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsCreateDefinitions`: MicrosoftGraphAccessReviewScheduleDefinition
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsCreateDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsCreateDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAccessReviewScheduleDefinition** | [**MicrosoftGraphAccessReviewScheduleDefinition**](MicrosoftGraphAccessReviewScheduleDefinition.md) | New navigation property | 

### Return type

[**MicrosoftGraphAccessReviewScheduleDefinition**](MicrosoftGraphAccessReviewScheduleDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsCreateInstances

> MicrosoftGraphAccessReviewInstance IdentityGovernanceAccessReviewsDefinitionsCreateInstances(ctx, accessReviewScheduleDefinitionId).MicrosoftGraphAccessReviewInstance(microsoftGraphAccessReviewInstance).Execute()

Create new navigation property to instances for identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    microsoftGraphAccessReviewInstance := *openapiclient.NewMicrosoftGraphAccessReviewInstance() // MicrosoftGraphAccessReviewInstance | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsCreateInstances(context.Background(), accessReviewScheduleDefinitionId).MicrosoftGraphAccessReviewInstance(microsoftGraphAccessReviewInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsCreateInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsCreateInstances`: MicrosoftGraphAccessReviewInstance
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsCreateInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsCreateInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAccessReviewInstance** | [**MicrosoftGraphAccessReviewInstance**](MicrosoftGraphAccessReviewInstance.md) | New navigation property | 

### Return type

[**MicrosoftGraphAccessReviewInstance**](MicrosoftGraphAccessReviewInstance.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsDeleteInstances

> IdentityGovernanceAccessReviewsDefinitionsDeleteInstances(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).IfMatch(ifMatch).Execute()

Delete navigation property instances for identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsDeleteInstances(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsDeleteInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsDeleteInstancesRequest struct via the builder pattern


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


## IdentityGovernanceAccessReviewsDefinitionsGetInstances

> MicrosoftGraphAccessReviewInstance IdentityGovernanceAccessReviewsDefinitionsGetInstances(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).Select_(select_).Expand(expand).Execute()

Get instances from identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsGetInstances(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsGetInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsGetInstances`: MicrosoftGraphAccessReviewInstance
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsGetInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsGetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAccessReviewInstance**](MicrosoftGraphAccessReviewInstance.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisions

> MicrosoftGraphAccessReviewInstanceDecisionItem IdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisions(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).MicrosoftGraphAccessReviewInstanceDecisionItem(microsoftGraphAccessReviewInstanceDecisionItem).Execute()

Create new navigation property to decisions for identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
    microsoftGraphAccessReviewInstanceDecisionItem := *openapiclient.NewMicrosoftGraphAccessReviewInstanceDecisionItem() // MicrosoftGraphAccessReviewInstanceDecisionItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisions(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).MicrosoftGraphAccessReviewInstanceDecisionItem(microsoftGraphAccessReviewInstanceDecisionItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisions`: MicrosoftGraphAccessReviewInstanceDecisionItem
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsInstancesCreateDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAccessReviewInstanceDecisionItem** | [**MicrosoftGraphAccessReviewInstanceDecisionItem**](MicrosoftGraphAccessReviewInstanceDecisionItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphAccessReviewInstanceDecisionItem**](MicrosoftGraphAccessReviewInstanceDecisionItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsInstancesDeleteDecisions

> IdentityGovernanceAccessReviewsDefinitionsInstancesDeleteDecisions(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId, accessReviewInstanceDecisionItemId).IfMatch(ifMatch).Execute()

Delete navigation property decisions for identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
    accessReviewInstanceDecisionItemId := "accessReviewInstanceDecisionItemId_example" // string | key: id of accessReviewInstanceDecisionItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesDeleteDecisions(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId, accessReviewInstanceDecisionItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesDeleteDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 
**accessReviewInstanceDecisionItemId** | **string** | key: id of accessReviewInstanceDecisionItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsInstancesDeleteDecisionsRequest struct via the builder pattern


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


## IdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisions

> MicrosoftGraphAccessReviewInstanceDecisionItem IdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisions(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId, accessReviewInstanceDecisionItemId).Select_(select_).Expand(expand).Execute()

Get decisions from identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
    accessReviewInstanceDecisionItemId := "accessReviewInstanceDecisionItemId_example" // string | key: id of accessReviewInstanceDecisionItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisions(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId, accessReviewInstanceDecisionItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisions`: MicrosoftGraphAccessReviewInstanceDecisionItem
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 
**accessReviewInstanceDecisionItemId** | **string** | key: id of accessReviewInstanceDecisionItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsInstancesGetDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAccessReviewInstanceDecisionItem**](MicrosoftGraphAccessReviewInstanceDecisionItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsInstancesListDecisions

> CollectionOfAccessReviewInstanceDecisionItem IdentityGovernanceAccessReviewsDefinitionsInstancesListDecisions(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get decisions from identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
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
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesListDecisions(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesListDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsInstancesListDecisions`: CollectionOfAccessReviewInstanceDecisionItem
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesListDecisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsInstancesListDecisionsRequest struct via the builder pattern


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

[**CollectionOfAccessReviewInstanceDecisionItem**](CollectionOfAccessReviewInstanceDecisionItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsInstancesUpdateDecisions

> IdentityGovernanceAccessReviewsDefinitionsInstancesUpdateDecisions(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId, accessReviewInstanceDecisionItemId).MicrosoftGraphAccessReviewInstanceDecisionItem(microsoftGraphAccessReviewInstanceDecisionItem).Execute()

Update the navigation property decisions in identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
    accessReviewInstanceDecisionItemId := "accessReviewInstanceDecisionItemId_example" // string | key: id of accessReviewInstanceDecisionItem
    microsoftGraphAccessReviewInstanceDecisionItem := *openapiclient.NewMicrosoftGraphAccessReviewInstanceDecisionItem() // MicrosoftGraphAccessReviewInstanceDecisionItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesUpdateDecisions(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId, accessReviewInstanceDecisionItemId).MicrosoftGraphAccessReviewInstanceDecisionItem(microsoftGraphAccessReviewInstanceDecisionItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsInstancesUpdateDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 
**accessReviewInstanceDecisionItemId** | **string** | key: id of accessReviewInstanceDecisionItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsInstancesUpdateDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphAccessReviewInstanceDecisionItem** | [**MicrosoftGraphAccessReviewInstanceDecisionItem**](MicrosoftGraphAccessReviewInstanceDecisionItem.md) | New navigation property values | 

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


## IdentityGovernanceAccessReviewsDefinitionsListInstances

> CollectionOfAccessReviewInstance IdentityGovernanceAccessReviewsDefinitionsListInstances(ctx, accessReviewScheduleDefinitionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get instances from identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
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
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsListInstances(context.Background(), accessReviewScheduleDefinitionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsListInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsListInstances`: CollectionOfAccessReviewInstance
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsListInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsListInstancesRequest struct via the builder pattern


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

[**CollectionOfAccessReviewInstance**](CollectionOfAccessReviewInstance.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsUpdateInstances

> IdentityGovernanceAccessReviewsDefinitionsUpdateInstances(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).MicrosoftGraphAccessReviewInstance(microsoftGraphAccessReviewInstance).Execute()

Update the navigation property instances in identityGovernance



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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
    microsoftGraphAccessReviewInstance := *openapiclient.NewMicrosoftGraphAccessReviewInstance() // MicrosoftGraphAccessReviewInstance | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsUpdateInstances(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).MicrosoftGraphAccessReviewInstance(microsoftGraphAccessReviewInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDefinitionsUpdateInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsUpdateInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAccessReviewInstance** | [**MicrosoftGraphAccessReviewInstance**](MicrosoftGraphAccessReviewInstance.md) | New navigation property values | 

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


## IdentityGovernanceAccessReviewsDeleteDefinitions

> IdentityGovernanceAccessReviewsDeleteDefinitions(ctx, accessReviewScheduleDefinitionId).IfMatch(ifMatch).Execute()

Delete navigation property definitions for identityGovernance

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDeleteDefinitions(context.Background(), accessReviewScheduleDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsDeleteDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDeleteDefinitionsRequest struct via the builder pattern


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


## IdentityGovernanceAccessReviewsGetDefinitions

> MicrosoftGraphAccessReviewScheduleDefinition IdentityGovernanceAccessReviewsGetDefinitions(ctx, accessReviewScheduleDefinitionId).Select_(select_).Expand(expand).Execute()

Get definitions from identityGovernance

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsGetDefinitions(context.Background(), accessReviewScheduleDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsGetDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsGetDefinitions`: MicrosoftGraphAccessReviewScheduleDefinition
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsGetDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsGetDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAccessReviewScheduleDefinition**](MicrosoftGraphAccessReviewScheduleDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsListDefinitions

> CollectionOfAccessReviewScheduleDefinition IdentityGovernanceAccessReviewsListDefinitions(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get definitions from identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsListDefinitions(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsListDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsListDefinitions`: CollectionOfAccessReviewScheduleDefinition
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsListDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsListDefinitionsRequest struct via the builder pattern


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

[**CollectionOfAccessReviewScheduleDefinition**](CollectionOfAccessReviewScheduleDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsUpdateDefinitions

> IdentityGovernanceAccessReviewsUpdateDefinitions(ctx, accessReviewScheduleDefinitionId).MicrosoftGraphAccessReviewScheduleDefinition(microsoftGraphAccessReviewScheduleDefinition).Execute()

Update the navigation property definitions in identityGovernance

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    microsoftGraphAccessReviewScheduleDefinition := *openapiclient.NewMicrosoftGraphAccessReviewScheduleDefinition() // MicrosoftGraphAccessReviewScheduleDefinition | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsUpdateDefinitions(context.Background(), accessReviewScheduleDefinitionId).MicrosoftGraphAccessReviewScheduleDefinition(microsoftGraphAccessReviewScheduleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceAccessReviewsUpdateDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsUpdateDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAccessReviewScheduleDefinition** | [**MicrosoftGraphAccessReviewScheduleDefinition**](MicrosoftGraphAccessReviewScheduleDefinition.md) | New navigation property values | 

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


## IdentityGovernanceDeleteAccessReviews

> IdentityGovernanceDeleteAccessReviews(ctx).IfMatch(ifMatch).Execute()

Delete navigation property accessReviews for identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceDeleteAccessReviews(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceDeleteAccessReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceDeleteAccessReviewsRequest struct via the builder pattern


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


## IdentityGovernanceGetAccessReviews

> MicrosoftGraphAccessReviewSet IdentityGovernanceGetAccessReviews(ctx).Select_(select_).Expand(expand).Execute()

Get accessReviews from identityGovernance

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
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceGetAccessReviews(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceGetAccessReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceGetAccessReviews`: MicrosoftGraphAccessReviewSet
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceGetAccessReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceGetAccessReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAccessReviewSet**](MicrosoftGraphAccessReviewSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceUpdateAccessReviews

> IdentityGovernanceUpdateAccessReviews(ctx).MicrosoftGraphAccessReviewSet(microsoftGraphAccessReviewSet).Execute()

Update the navigation property accessReviews in identityGovernance

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
    microsoftGraphAccessReviewSet := *openapiclient.NewMicrosoftGraphAccessReviewSet() // MicrosoftGraphAccessReviewSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceAccessReviewSetApi.IdentityGovernanceUpdateAccessReviews(context.Background()).MicrosoftGraphAccessReviewSet(microsoftGraphAccessReviewSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceAccessReviewSetApi.IdentityGovernanceUpdateAccessReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceUpdateAccessReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAccessReviewSet** | [**MicrosoftGraphAccessReviewSet**](MicrosoftGraphAccessReviewSet.md) | New navigation property values | 

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

