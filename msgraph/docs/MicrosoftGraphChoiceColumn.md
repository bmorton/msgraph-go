# MicrosoftGraphChoiceColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowTextEntry** | Pointer to **NullableBool** | If true, allows custom values that aren&#39;t in the configured choices. | [optional] 
**Choices** | Pointer to **[]string** | The list of values available for this column. | [optional] 
**DisplayAs** | Pointer to **NullableString** | How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons | [optional] 

## Methods

### NewMicrosoftGraphChoiceColumn

`func NewMicrosoftGraphChoiceColumn() *MicrosoftGraphChoiceColumn`

NewMicrosoftGraphChoiceColumn instantiates a new MicrosoftGraphChoiceColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChoiceColumnWithDefaults

`func NewMicrosoftGraphChoiceColumnWithDefaults() *MicrosoftGraphChoiceColumn`

NewMicrosoftGraphChoiceColumnWithDefaults instantiates a new MicrosoftGraphChoiceColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowTextEntry

`func (o *MicrosoftGraphChoiceColumn) GetAllowTextEntry() bool`

GetAllowTextEntry returns the AllowTextEntry field if non-nil, zero value otherwise.

### GetAllowTextEntryOk

`func (o *MicrosoftGraphChoiceColumn) GetAllowTextEntryOk() (*bool, bool)`

GetAllowTextEntryOk returns a tuple with the AllowTextEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTextEntry

`func (o *MicrosoftGraphChoiceColumn) SetAllowTextEntry(v bool)`

SetAllowTextEntry sets AllowTextEntry field to given value.

### HasAllowTextEntry

`func (o *MicrosoftGraphChoiceColumn) HasAllowTextEntry() bool`

HasAllowTextEntry returns a boolean if a field has been set.

### SetAllowTextEntryNil

`func (o *MicrosoftGraphChoiceColumn) SetAllowTextEntryNil(b bool)`

 SetAllowTextEntryNil sets the value for AllowTextEntry to be an explicit nil

### UnsetAllowTextEntry
`func (o *MicrosoftGraphChoiceColumn) UnsetAllowTextEntry()`

UnsetAllowTextEntry ensures that no value is present for AllowTextEntry, not even an explicit nil
### GetChoices

`func (o *MicrosoftGraphChoiceColumn) GetChoices() []*string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *MicrosoftGraphChoiceColumn) GetChoicesOk() (*[]*string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *MicrosoftGraphChoiceColumn) SetChoices(v []*string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *MicrosoftGraphChoiceColumn) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetDisplayAs

`func (o *MicrosoftGraphChoiceColumn) GetDisplayAs() string`

GetDisplayAs returns the DisplayAs field if non-nil, zero value otherwise.

### GetDisplayAsOk

`func (o *MicrosoftGraphChoiceColumn) GetDisplayAsOk() (*string, bool)`

GetDisplayAsOk returns a tuple with the DisplayAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAs

`func (o *MicrosoftGraphChoiceColumn) SetDisplayAs(v string)`

SetDisplayAs sets DisplayAs field to given value.

### HasDisplayAs

`func (o *MicrosoftGraphChoiceColumn) HasDisplayAs() bool`

HasDisplayAs returns a boolean if a field has been set.

### SetDisplayAsNil

`func (o *MicrosoftGraphChoiceColumn) SetDisplayAsNil(b bool)`

 SetDisplayAsNil sets the value for DisplayAs to be an explicit nil

### UnsetDisplayAs
`func (o *MicrosoftGraphChoiceColumn) UnsetDisplayAs()`

UnsetDisplayAs ensures that no value is present for DisplayAs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


