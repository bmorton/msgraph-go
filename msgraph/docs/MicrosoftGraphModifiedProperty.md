# MicrosoftGraphModifiedProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | Indicates the property name of the target attribute that was changed. | [optional] 
**NewValue** | Pointer to **NullableString** | Indicates the updated value for the propery. | [optional] 
**OldValue** | Pointer to **NullableString** | Indicates the previous value (before the update) for the property. | [optional] 

## Methods

### NewMicrosoftGraphModifiedProperty

`func NewMicrosoftGraphModifiedProperty() *MicrosoftGraphModifiedProperty`

NewMicrosoftGraphModifiedProperty instantiates a new MicrosoftGraphModifiedProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphModifiedPropertyWithDefaults

`func NewMicrosoftGraphModifiedPropertyWithDefaults() *MicrosoftGraphModifiedProperty`

NewMicrosoftGraphModifiedPropertyWithDefaults instantiates a new MicrosoftGraphModifiedProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphModifiedProperty) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphModifiedProperty) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphModifiedProperty) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphModifiedProperty) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphModifiedProperty) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphModifiedProperty) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetNewValue

`func (o *MicrosoftGraphModifiedProperty) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *MicrosoftGraphModifiedProperty) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *MicrosoftGraphModifiedProperty) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *MicrosoftGraphModifiedProperty) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### SetNewValueNil

`func (o *MicrosoftGraphModifiedProperty) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *MicrosoftGraphModifiedProperty) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
### GetOldValue

`func (o *MicrosoftGraphModifiedProperty) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *MicrosoftGraphModifiedProperty) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *MicrosoftGraphModifiedProperty) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *MicrosoftGraphModifiedProperty) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### SetOldValueNil

`func (o *MicrosoftGraphModifiedProperty) SetOldValueNil(b bool)`

 SetOldValueNil sets the value for OldValue to be an explicit nil

### UnsetOldValue
`func (o *MicrosoftGraphModifiedProperty) UnsetOldValue()`

UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


