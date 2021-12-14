# MicrosoftGraphSortProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDescending** | Pointer to **NullableBool** | True if the sort order is descending. Default is false, with the sort order as ascending. Optional. | [optional] 
**Name** | Pointer to **string** | The name of the property to sort on. Required. | [optional] 

## Methods

### NewMicrosoftGraphSortProperty

`func NewMicrosoftGraphSortProperty() *MicrosoftGraphSortProperty`

NewMicrosoftGraphSortProperty instantiates a new MicrosoftGraphSortProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSortPropertyWithDefaults

`func NewMicrosoftGraphSortPropertyWithDefaults() *MicrosoftGraphSortProperty`

NewMicrosoftGraphSortPropertyWithDefaults instantiates a new MicrosoftGraphSortProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDescending

`func (o *MicrosoftGraphSortProperty) GetIsDescending() bool`

GetIsDescending returns the IsDescending field if non-nil, zero value otherwise.

### GetIsDescendingOk

`func (o *MicrosoftGraphSortProperty) GetIsDescendingOk() (*bool, bool)`

GetIsDescendingOk returns a tuple with the IsDescending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDescending

`func (o *MicrosoftGraphSortProperty) SetIsDescending(v bool)`

SetIsDescending sets IsDescending field to given value.

### HasIsDescending

`func (o *MicrosoftGraphSortProperty) HasIsDescending() bool`

HasIsDescending returns a boolean if a field has been set.

### SetIsDescendingNil

`func (o *MicrosoftGraphSortProperty) SetIsDescendingNil(b bool)`

 SetIsDescendingNil sets the value for IsDescending to be an explicit nil

### UnsetIsDescending
`func (o *MicrosoftGraphSortProperty) UnsetIsDescending()`

UnsetIsDescending ensures that no value is present for IsDescending, not even an explicit nil
### GetName

`func (o *MicrosoftGraphSortProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphSortProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphSortProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphSortProperty) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


