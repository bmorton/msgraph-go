# \PrivacyFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachment**](PrivacyFunctionsApi.md#PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachment) | **Get** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/microsoft.graph.getFinalAttachment() | Invoke function getFinalAttachment
[**PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReport**](PrivacyFunctionsApi.md#PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReport) | **Get** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/microsoft.graph.getFinalReport() | Invoke function getFinalReport



## PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachment

> string PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachment(ctx, subjectRightsRequestId).Execute()

Invoke function getFinalAttachment

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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacyFunctionsApi.PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachment(context.Background(), subjectRightsRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacyFunctionsApi.PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachment`: string
    fmt.Fprintf(os.Stdout, "Response from `PrivacyFunctionsApi.PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsSubjectRightsRequestGetFinalAttachmentRequest struct via the builder pattern


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


## PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReport

> string PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReport(ctx, subjectRightsRequestId).Execute()

Invoke function getFinalReport

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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacyFunctionsApi.PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReport(context.Background(), subjectRightsRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacyFunctionsApi.PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReport`: string
    fmt.Fprintf(os.Stdout, "Response from `PrivacyFunctionsApi.PrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsSubjectRightsRequestGetFinalReportRequest struct via the builder pattern


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

