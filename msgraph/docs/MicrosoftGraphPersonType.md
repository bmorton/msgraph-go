# MicrosoftGraphPersonType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **NullableString** | The type of data source, such as Person. | [optional] 
**Subclass** | Pointer to **NullableString** | The secondary type of data source, such as OrganizationUser. | [optional] 

## Methods

### NewMicrosoftGraphPersonType

`func NewMicrosoftGraphPersonType() *MicrosoftGraphPersonType`

NewMicrosoftGraphPersonType instantiates a new MicrosoftGraphPersonType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPersonTypeWithDefaults

`func NewMicrosoftGraphPersonTypeWithDefaults() *MicrosoftGraphPersonType`

NewMicrosoftGraphPersonTypeWithDefaults instantiates a new MicrosoftGraphPersonType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *MicrosoftGraphPersonType) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *MicrosoftGraphPersonType) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *MicrosoftGraphPersonType) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *MicrosoftGraphPersonType) HasClass() bool`

HasClass returns a boolean if a field has been set.

### SetClassNil

`func (o *MicrosoftGraphPersonType) SetClassNil(b bool)`

 SetClassNil sets the value for Class to be an explicit nil

### UnsetClass
`func (o *MicrosoftGraphPersonType) UnsetClass()`

UnsetClass ensures that no value is present for Class, not even an explicit nil
### GetSubclass

`func (o *MicrosoftGraphPersonType) GetSubclass() string`

GetSubclass returns the Subclass field if non-nil, zero value otherwise.

### GetSubclassOk

`func (o *MicrosoftGraphPersonType) GetSubclassOk() (*string, bool)`

GetSubclassOk returns a tuple with the Subclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclass

`func (o *MicrosoftGraphPersonType) SetSubclass(v string)`

SetSubclass sets Subclass field to given value.

### HasSubclass

`func (o *MicrosoftGraphPersonType) HasSubclass() bool`

HasSubclass returns a boolean if a field has been set.

### SetSubclassNil

`func (o *MicrosoftGraphPersonType) SetSubclassNil(b bool)`

 SetSubclassNil sets the value for Subclass to be an explicit nil

### UnsetSubclass
`func (o *MicrosoftGraphPersonType) UnsetSubclass()`

UnsetSubclass ensures that no value is present for Subclass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


