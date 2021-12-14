# MicrosoftGraphConditionalAccessLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeLocations** | Pointer to **[]string** | Location IDs excluded from scope of policy. | [optional] 
**IncludeLocations** | Pointer to **[]string** | Location IDs in scope of policy unless explicitly excluded, All, or AllTrusted. | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessLocations

`func NewMicrosoftGraphConditionalAccessLocations() *MicrosoftGraphConditionalAccessLocations`

NewMicrosoftGraphConditionalAccessLocations instantiates a new MicrosoftGraphConditionalAccessLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessLocationsWithDefaults

`func NewMicrosoftGraphConditionalAccessLocationsWithDefaults() *MicrosoftGraphConditionalAccessLocations`

NewMicrosoftGraphConditionalAccessLocationsWithDefaults instantiates a new MicrosoftGraphConditionalAccessLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeLocations

`func (o *MicrosoftGraphConditionalAccessLocations) GetExcludeLocations() []string`

GetExcludeLocations returns the ExcludeLocations field if non-nil, zero value otherwise.

### GetExcludeLocationsOk

`func (o *MicrosoftGraphConditionalAccessLocations) GetExcludeLocationsOk() (*[]string, bool)`

GetExcludeLocationsOk returns a tuple with the ExcludeLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLocations

`func (o *MicrosoftGraphConditionalAccessLocations) SetExcludeLocations(v []string)`

SetExcludeLocations sets ExcludeLocations field to given value.

### HasExcludeLocations

`func (o *MicrosoftGraphConditionalAccessLocations) HasExcludeLocations() bool`

HasExcludeLocations returns a boolean if a field has been set.

### GetIncludeLocations

`func (o *MicrosoftGraphConditionalAccessLocations) GetIncludeLocations() []string`

GetIncludeLocations returns the IncludeLocations field if non-nil, zero value otherwise.

### GetIncludeLocationsOk

`func (o *MicrosoftGraphConditionalAccessLocations) GetIncludeLocationsOk() (*[]string, bool)`

GetIncludeLocationsOk returns a tuple with the IncludeLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocations

`func (o *MicrosoftGraphConditionalAccessLocations) SetIncludeLocations(v []string)`

SetIncludeLocations sets IncludeLocations field to given value.

### HasIncludeLocations

`func (o *MicrosoftGraphConditionalAccessLocations) HasIncludeLocations() bool`

HasIncludeLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


