# PrintService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]MicrosoftGraphPrintServiceEndpoint**](MicrosoftGraphPrintServiceEndpoint.md) | Endpoints that can be used to access the service. Read-only. Nullable. | [optional] 

## Methods

### NewPrintService

`func NewPrintService() *PrintService`

NewPrintService instantiates a new PrintService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintServiceWithDefaults

`func NewPrintServiceWithDefaults() *PrintService`

NewPrintServiceWithDefaults instantiates a new PrintService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *PrintService) GetEndpoints() []MicrosoftGraphPrintServiceEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PrintService) GetEndpointsOk() (*[]MicrosoftGraphPrintServiceEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PrintService) SetEndpoints(v []MicrosoftGraphPrintServiceEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *PrintService) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


