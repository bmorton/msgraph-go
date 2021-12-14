# MicrosoftGraphEmailAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **NullableString** | The email address of the person or entity. | [optional] 
**Name** | Pointer to **NullableString** | The display name of the person or entity. | [optional] 

## Methods

### NewMicrosoftGraphEmailAddress

`func NewMicrosoftGraphEmailAddress() *MicrosoftGraphEmailAddress`

NewMicrosoftGraphEmailAddress instantiates a new MicrosoftGraphEmailAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEmailAddressWithDefaults

`func NewMicrosoftGraphEmailAddressWithDefaults() *MicrosoftGraphEmailAddress`

NewMicrosoftGraphEmailAddressWithDefaults instantiates a new MicrosoftGraphEmailAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MicrosoftGraphEmailAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphEmailAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftGraphEmailAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftGraphEmailAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MicrosoftGraphEmailAddress) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MicrosoftGraphEmailAddress) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetName

`func (o *MicrosoftGraphEmailAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphEmailAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphEmailAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphEmailAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphEmailAddress) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphEmailAddress) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


