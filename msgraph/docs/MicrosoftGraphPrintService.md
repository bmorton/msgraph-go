# MicrosoftGraphPrintService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Endpoints** | Pointer to [**[]MicrosoftGraphPrintServiceEndpoint**](MicrosoftGraphPrintServiceEndpoint.md) | Endpoints that can be used to access the service. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphPrintService

`func NewMicrosoftGraphPrintService() *MicrosoftGraphPrintService`

NewMicrosoftGraphPrintService instantiates a new MicrosoftGraphPrintService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintServiceWithDefaults

`func NewMicrosoftGraphPrintServiceWithDefaults() *MicrosoftGraphPrintService`

NewMicrosoftGraphPrintServiceWithDefaults instantiates a new MicrosoftGraphPrintService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrintService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrintService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrintService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrintService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEndpoints

`func (o *MicrosoftGraphPrintService) GetEndpoints() []MicrosoftGraphPrintServiceEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *MicrosoftGraphPrintService) GetEndpointsOk() (*[]MicrosoftGraphPrintServiceEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *MicrosoftGraphPrintService) SetEndpoints(v []MicrosoftGraphPrintServiceEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *MicrosoftGraphPrintService) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


