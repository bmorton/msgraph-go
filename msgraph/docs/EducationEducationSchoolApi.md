# \EducationEducationSchoolApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationCreateSchools**](EducationEducationSchoolApi.md#EducationCreateSchools) | **Post** /education/schools | Create new navigation property to schools for education
[**EducationDeleteSchools**](EducationEducationSchoolApi.md#EducationDeleteSchools) | **Delete** /education/schools/{educationSchool-id} | Delete navigation property schools for education
[**EducationGetSchools**](EducationEducationSchoolApi.md#EducationGetSchools) | **Get** /education/schools/{educationSchool-id} | Get schools from education
[**EducationListSchools**](EducationEducationSchoolApi.md#EducationListSchools) | **Get** /education/schools | Get schools from education
[**EducationSchoolsCreateRefClasses**](EducationEducationSchoolApi.md#EducationSchoolsCreateRefClasses) | **Post** /education/schools/{educationSchool-id}/classes/$ref | Create new navigation property ref to classes for education
[**EducationSchoolsCreateRefUsers**](EducationEducationSchoolApi.md#EducationSchoolsCreateRefUsers) | **Post** /education/schools/{educationSchool-id}/users/$ref | Create new navigation property ref to users for education
[**EducationSchoolsDeleteRefAdministrativeUnit**](EducationEducationSchoolApi.md#EducationSchoolsDeleteRefAdministrativeUnit) | **Delete** /education/schools/{educationSchool-id}/administrativeUnit/$ref | Delete ref of navigation property administrativeUnit for education
[**EducationSchoolsGetAdministrativeUnit**](EducationEducationSchoolApi.md#EducationSchoolsGetAdministrativeUnit) | **Get** /education/schools/{educationSchool-id}/administrativeUnit | Get administrativeUnit from education
[**EducationSchoolsGetRefAdministrativeUnit**](EducationEducationSchoolApi.md#EducationSchoolsGetRefAdministrativeUnit) | **Get** /education/schools/{educationSchool-id}/administrativeUnit/$ref | Get ref of administrativeUnit from education
[**EducationSchoolsListClasses**](EducationEducationSchoolApi.md#EducationSchoolsListClasses) | **Get** /education/schools/{educationSchool-id}/classes | Get classes from education
[**EducationSchoolsListRefClasses**](EducationEducationSchoolApi.md#EducationSchoolsListRefClasses) | **Get** /education/schools/{educationSchool-id}/classes/$ref | Get ref of classes from education
[**EducationSchoolsListRefUsers**](EducationEducationSchoolApi.md#EducationSchoolsListRefUsers) | **Get** /education/schools/{educationSchool-id}/users/$ref | Get ref of users from education
[**EducationSchoolsListUsers**](EducationEducationSchoolApi.md#EducationSchoolsListUsers) | **Get** /education/schools/{educationSchool-id}/users | Get users from education
[**EducationSchoolsUpdateRefAdministrativeUnit**](EducationEducationSchoolApi.md#EducationSchoolsUpdateRefAdministrativeUnit) | **Put** /education/schools/{educationSchool-id}/administrativeUnit/$ref | Update the ref of navigation property administrativeUnit in education
[**EducationUpdateSchools**](EducationEducationSchoolApi.md#EducationUpdateSchools) | **Patch** /education/schools/{educationSchool-id} | Update the navigation property schools in education



## EducationCreateSchools

> MicrosoftGraphEducationSchool EducationCreateSchools(ctx).MicrosoftGraphEducationSchool(microsoftGraphEducationSchool).Execute()

Create new navigation property to schools for education

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
    microsoftGraphEducationSchool := *openapiclient.NewMicrosoftGraphEducationSchool() // MicrosoftGraphEducationSchool | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationCreateSchools(context.Background()).MicrosoftGraphEducationSchool(microsoftGraphEducationSchool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationCreateSchools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationCreateSchools`: MicrosoftGraphEducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationCreateSchools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEducationCreateSchoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphEducationSchool** | [**MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md) | New navigation property | 

### Return type

[**MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationDeleteSchools

> EducationDeleteSchools(ctx, educationSchoolId).IfMatch(ifMatch).Execute()

Delete navigation property schools for education

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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationDeleteSchools(context.Background(), educationSchoolId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationDeleteSchools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationDeleteSchoolsRequest struct via the builder pattern


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


## EducationGetSchools

> MicrosoftGraphEducationSchool EducationGetSchools(ctx, educationSchoolId).Select_(select_).Expand(expand).Execute()

Get schools from education

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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationGetSchools(context.Background(), educationSchoolId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationGetSchools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationGetSchools`: MicrosoftGraphEducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationGetSchools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationGetSchoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationListSchools

> CollectionOfEducationSchool EducationListSchools(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get schools from education

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
    resp, r, err := api_client.EducationEducationSchoolApi.EducationListSchools(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationListSchools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationListSchools`: CollectionOfEducationSchool
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationListSchools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEducationListSchoolsRequest struct via the builder pattern


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

[**CollectionOfEducationSchool**](CollectionOfEducationSchool.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsCreateRefClasses

> map[string]interface{} EducationSchoolsCreateRefClasses(ctx, educationSchoolId).RequestBody(requestBody).Execute()

Create new navigation property ref to classes for education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsCreateRefClasses(context.Background(), educationSchoolId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsCreateRefClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsCreateRefClasses`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationSchoolsCreateRefClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsCreateRefClassesRequest struct via the builder pattern


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


## EducationSchoolsCreateRefUsers

> map[string]interface{} EducationSchoolsCreateRefUsers(ctx, educationSchoolId).RequestBody(requestBody).Execute()

Create new navigation property ref to users for education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsCreateRefUsers(context.Background(), educationSchoolId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsCreateRefUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsCreateRefUsers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationSchoolsCreateRefUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsCreateRefUsersRequest struct via the builder pattern


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


## EducationSchoolsDeleteRefAdministrativeUnit

> EducationSchoolsDeleteRefAdministrativeUnit(ctx, educationSchoolId).IfMatch(ifMatch).Execute()

Delete ref of navigation property administrativeUnit for education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsDeleteRefAdministrativeUnit(context.Background(), educationSchoolId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsDeleteRefAdministrativeUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsDeleteRefAdministrativeUnitRequest struct via the builder pattern


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


## EducationSchoolsGetAdministrativeUnit

> MicrosoftGraphAdministrativeUnit EducationSchoolsGetAdministrativeUnit(ctx, educationSchoolId).Select_(select_).Expand(expand).Execute()

Get administrativeUnit from education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsGetAdministrativeUnit(context.Background(), educationSchoolId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsGetAdministrativeUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsGetAdministrativeUnit`: MicrosoftGraphAdministrativeUnit
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationSchoolsGetAdministrativeUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsGetAdministrativeUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsGetRefAdministrativeUnit

> string EducationSchoolsGetRefAdministrativeUnit(ctx, educationSchoolId).Execute()

Get ref of administrativeUnit from education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsGetRefAdministrativeUnit(context.Background(), educationSchoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsGetRefAdministrativeUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsGetRefAdministrativeUnit`: string
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationSchoolsGetRefAdministrativeUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsGetRefAdministrativeUnitRequest struct via the builder pattern


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


## EducationSchoolsListClasses

> CollectionOfEducationClass EducationSchoolsListClasses(ctx, educationSchoolId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get classes from education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
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
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsListClasses(context.Background(), educationSchoolId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsListClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsListClasses`: CollectionOfEducationClass
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationSchoolsListClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsListClassesRequest struct via the builder pattern


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

[**CollectionOfEducationClass**](CollectionOfEducationClass.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsListRefClasses

> CollectionOfLinksOfEducationClass EducationSchoolsListRefClasses(ctx, educationSchoolId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of classes from education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsListRefClasses(context.Background(), educationSchoolId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsListRefClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsListRefClasses`: CollectionOfLinksOfEducationClass
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationSchoolsListRefClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsListRefClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfEducationClass**](CollectionOfLinksOfEducationClass.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsListRefUsers

> CollectionOfLinksOfEducationUser EducationSchoolsListRefUsers(ctx, educationSchoolId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of users from education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsListRefUsers(context.Background(), educationSchoolId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsListRefUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsListRefUsers`: CollectionOfLinksOfEducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationSchoolsListRefUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsListRefUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfEducationUser**](CollectionOfLinksOfEducationUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsListUsers

> CollectionOfEducationUser EducationSchoolsListUsers(ctx, educationSchoolId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get users from education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
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
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsListUsers(context.Background(), educationSchoolId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationSchoolsListUsers`: CollectionOfEducationUser
    fmt.Fprintf(os.Stdout, "Response from `EducationEducationSchoolApi.EducationSchoolsListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsListUsersRequest struct via the builder pattern


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

[**CollectionOfEducationUser**](CollectionOfEducationUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationSchoolsUpdateRefAdministrativeUnit

> EducationSchoolsUpdateRefAdministrativeUnit(ctx, educationSchoolId).RequestBody(requestBody).Execute()

Update the ref of navigation property administrativeUnit in education



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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationSchoolsUpdateRefAdministrativeUnit(context.Background(), educationSchoolId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationSchoolsUpdateRefAdministrativeUnit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationSchoolsUpdateRefAdministrativeUnitRequest struct via the builder pattern


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


## EducationUpdateSchools

> EducationUpdateSchools(ctx, educationSchoolId).MicrosoftGraphEducationSchool(microsoftGraphEducationSchool).Execute()

Update the navigation property schools in education

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
    educationSchoolId := "educationSchoolId_example" // string | key: id of educationSchool
    microsoftGraphEducationSchool := *openapiclient.NewMicrosoftGraphEducationSchool() // MicrosoftGraphEducationSchool | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationEducationSchoolApi.EducationUpdateSchools(context.Background(), educationSchoolId).MicrosoftGraphEducationSchool(microsoftGraphEducationSchool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationEducationSchoolApi.EducationUpdateSchools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationSchoolId** | **string** | key: id of educationSchool | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationUpdateSchoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphEducationSchool** | [**MicrosoftGraphEducationSchool**](MicrosoftGraphEducationSchool.md) | New navigation property values | 

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

