# MicrosoftGraphSignInFrequencySessionControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **NullableBool** | Specifies whether the session control is enabled. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphSigninFrequencyType**](anyOf&lt;microsoft.graph.signinFrequencyType&gt;.md) | Possible values are: days, hours. | [optional] 
**Value** | Pointer to **NullableInt32** | The number of days or hours. | [optional] 

## Methods

### NewMicrosoftGraphSignInFrequencySessionControl

`func NewMicrosoftGraphSignInFrequencySessionControl() *MicrosoftGraphSignInFrequencySessionControl`

NewMicrosoftGraphSignInFrequencySessionControl instantiates a new MicrosoftGraphSignInFrequencySessionControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSignInFrequencySessionControlWithDefaults

`func NewMicrosoftGraphSignInFrequencySessionControlWithDefaults() *MicrosoftGraphSignInFrequencySessionControl`

NewMicrosoftGraphSignInFrequencySessionControlWithDefaults instantiates a new MicrosoftGraphSignInFrequencySessionControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *MicrosoftGraphSignInFrequencySessionControl) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphSignInFrequencySessionControl) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphSignInFrequencySessionControl) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphSignInFrequencySessionControl) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *MicrosoftGraphSignInFrequencySessionControl) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *MicrosoftGraphSignInFrequencySessionControl) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetType

`func (o *MicrosoftGraphSignInFrequencySessionControl) GetType() AnyOfmicrosoftGraphSigninFrequencyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphSignInFrequencySessionControl) GetTypeOk() (*AnyOfmicrosoftGraphSigninFrequencyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphSignInFrequencySessionControl) SetType(v AnyOfmicrosoftGraphSigninFrequencyType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphSignInFrequencySessionControl) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphSignInFrequencySessionControl) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphSignInFrequencySessionControl) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphSignInFrequencySessionControl) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphSignInFrequencySessionControl) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphSignInFrequencySessionControl) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphSignInFrequencySessionControl) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphSignInFrequencySessionControl) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphSignInFrequencySessionControl) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


