# MicrosoftGraphPersonOrGroupColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleSelection** | Pointer to **NullableBool** | Indicates whether multiple values can be selected from the source. | [optional] 
**ChooseFromType** | Pointer to **NullableString** | Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly. | [optional] 
**DisplayAs** | Pointer to **NullableString** | How to display the information about the person or group chosen. See below. | [optional] 

## Methods

### NewMicrosoftGraphPersonOrGroupColumn

`func NewMicrosoftGraphPersonOrGroupColumn() *MicrosoftGraphPersonOrGroupColumn`

NewMicrosoftGraphPersonOrGroupColumn instantiates a new MicrosoftGraphPersonOrGroupColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPersonOrGroupColumnWithDefaults

`func NewMicrosoftGraphPersonOrGroupColumnWithDefaults() *MicrosoftGraphPersonOrGroupColumn`

NewMicrosoftGraphPersonOrGroupColumnWithDefaults instantiates a new MicrosoftGraphPersonOrGroupColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleSelection

`func (o *MicrosoftGraphPersonOrGroupColumn) GetAllowMultipleSelection() bool`

GetAllowMultipleSelection returns the AllowMultipleSelection field if non-nil, zero value otherwise.

### GetAllowMultipleSelectionOk

`func (o *MicrosoftGraphPersonOrGroupColumn) GetAllowMultipleSelectionOk() (*bool, bool)`

GetAllowMultipleSelectionOk returns a tuple with the AllowMultipleSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSelection

`func (o *MicrosoftGraphPersonOrGroupColumn) SetAllowMultipleSelection(v bool)`

SetAllowMultipleSelection sets AllowMultipleSelection field to given value.

### HasAllowMultipleSelection

`func (o *MicrosoftGraphPersonOrGroupColumn) HasAllowMultipleSelection() bool`

HasAllowMultipleSelection returns a boolean if a field has been set.

### SetAllowMultipleSelectionNil

`func (o *MicrosoftGraphPersonOrGroupColumn) SetAllowMultipleSelectionNil(b bool)`

 SetAllowMultipleSelectionNil sets the value for AllowMultipleSelection to be an explicit nil

### UnsetAllowMultipleSelection
`func (o *MicrosoftGraphPersonOrGroupColumn) UnsetAllowMultipleSelection()`

UnsetAllowMultipleSelection ensures that no value is present for AllowMultipleSelection, not even an explicit nil
### GetChooseFromType

`func (o *MicrosoftGraphPersonOrGroupColumn) GetChooseFromType() string`

GetChooseFromType returns the ChooseFromType field if non-nil, zero value otherwise.

### GetChooseFromTypeOk

`func (o *MicrosoftGraphPersonOrGroupColumn) GetChooseFromTypeOk() (*string, bool)`

GetChooseFromTypeOk returns a tuple with the ChooseFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChooseFromType

`func (o *MicrosoftGraphPersonOrGroupColumn) SetChooseFromType(v string)`

SetChooseFromType sets ChooseFromType field to given value.

### HasChooseFromType

`func (o *MicrosoftGraphPersonOrGroupColumn) HasChooseFromType() bool`

HasChooseFromType returns a boolean if a field has been set.

### SetChooseFromTypeNil

`func (o *MicrosoftGraphPersonOrGroupColumn) SetChooseFromTypeNil(b bool)`

 SetChooseFromTypeNil sets the value for ChooseFromType to be an explicit nil

### UnsetChooseFromType
`func (o *MicrosoftGraphPersonOrGroupColumn) UnsetChooseFromType()`

UnsetChooseFromType ensures that no value is present for ChooseFromType, not even an explicit nil
### GetDisplayAs

`func (o *MicrosoftGraphPersonOrGroupColumn) GetDisplayAs() string`

GetDisplayAs returns the DisplayAs field if non-nil, zero value otherwise.

### GetDisplayAsOk

`func (o *MicrosoftGraphPersonOrGroupColumn) GetDisplayAsOk() (*string, bool)`

GetDisplayAsOk returns a tuple with the DisplayAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAs

`func (o *MicrosoftGraphPersonOrGroupColumn) SetDisplayAs(v string)`

SetDisplayAs sets DisplayAs field to given value.

### HasDisplayAs

`func (o *MicrosoftGraphPersonOrGroupColumn) HasDisplayAs() bool`

HasDisplayAs returns a boolean if a field has been set.

### SetDisplayAsNil

`func (o *MicrosoftGraphPersonOrGroupColumn) SetDisplayAsNil(b bool)`

 SetDisplayAsNil sets the value for DisplayAs to be an explicit nil

### UnsetDisplayAs
`func (o *MicrosoftGraphPersonOrGroupColumn) UnsetDisplayAs()`

UnsetDisplayAs ensures that no value is present for DisplayAs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


