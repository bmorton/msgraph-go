# MicrosoftGraphUserAttributeValuesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** | Determines whether the value is set as the default. | [optional] 
**Name** | Pointer to **NullableString** | The display name of the property displayed to the user in the user flow. | [optional] 
**Value** | Pointer to **NullableString** | The value that is set when this item is selected. | [optional] 

## Methods

### NewMicrosoftGraphUserAttributeValuesItem

`func NewMicrosoftGraphUserAttributeValuesItem() *MicrosoftGraphUserAttributeValuesItem`

NewMicrosoftGraphUserAttributeValuesItem instantiates a new MicrosoftGraphUserAttributeValuesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserAttributeValuesItemWithDefaults

`func NewMicrosoftGraphUserAttributeValuesItemWithDefaults() *MicrosoftGraphUserAttributeValuesItem`

NewMicrosoftGraphUserAttributeValuesItemWithDefaults instantiates a new MicrosoftGraphUserAttributeValuesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *MicrosoftGraphUserAttributeValuesItem) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphUserAttributeValuesItem) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphUserAttributeValuesItem) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphUserAttributeValuesItem) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftGraphUserAttributeValuesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphUserAttributeValuesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphUserAttributeValuesItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphUserAttributeValuesItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphUserAttributeValuesItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphUserAttributeValuesItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphUserAttributeValuesItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphUserAttributeValuesItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphUserAttributeValuesItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphUserAttributeValuesItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphUserAttributeValuesItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphUserAttributeValuesItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


