# MicrosoftGraphConditionalAccessGrantControls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltInControls** | Pointer to [**[]AnyOfmicrosoftGraphConditionalAccessGrantControl**](AnyOfmicrosoftGraphConditionalAccessGrantControl.md) | List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue. | [optional] 
**CustomAuthenticationFactors** | Pointer to **[]string** | List of custom controls IDs required by the policy. For more information, see Custom controls. | [optional] 
**Operator** | Pointer to **NullableString** | Defines the relationship of the grant controls. Possible values: AND, OR. | [optional] 
**TermsOfUse** | Pointer to **[]string** | List of terms of use IDs required by the policy. | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessGrantControls

`func NewMicrosoftGraphConditionalAccessGrantControls() *MicrosoftGraphConditionalAccessGrantControls`

NewMicrosoftGraphConditionalAccessGrantControls instantiates a new MicrosoftGraphConditionalAccessGrantControls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessGrantControlsWithDefaults

`func NewMicrosoftGraphConditionalAccessGrantControlsWithDefaults() *MicrosoftGraphConditionalAccessGrantControls`

NewMicrosoftGraphConditionalAccessGrantControlsWithDefaults instantiates a new MicrosoftGraphConditionalAccessGrantControls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltInControls

`func (o *MicrosoftGraphConditionalAccessGrantControls) GetBuiltInControls() []AnyOfmicrosoftGraphConditionalAccessGrantControl`

GetBuiltInControls returns the BuiltInControls field if non-nil, zero value otherwise.

### GetBuiltInControlsOk

`func (o *MicrosoftGraphConditionalAccessGrantControls) GetBuiltInControlsOk() (*[]AnyOfmicrosoftGraphConditionalAccessGrantControl, bool)`

GetBuiltInControlsOk returns a tuple with the BuiltInControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInControls

`func (o *MicrosoftGraphConditionalAccessGrantControls) SetBuiltInControls(v []AnyOfmicrosoftGraphConditionalAccessGrantControl)`

SetBuiltInControls sets BuiltInControls field to given value.

### HasBuiltInControls

`func (o *MicrosoftGraphConditionalAccessGrantControls) HasBuiltInControls() bool`

HasBuiltInControls returns a boolean if a field has been set.

### GetCustomAuthenticationFactors

`func (o *MicrosoftGraphConditionalAccessGrantControls) GetCustomAuthenticationFactors() []string`

GetCustomAuthenticationFactors returns the CustomAuthenticationFactors field if non-nil, zero value otherwise.

### GetCustomAuthenticationFactorsOk

`func (o *MicrosoftGraphConditionalAccessGrantControls) GetCustomAuthenticationFactorsOk() (*[]string, bool)`

GetCustomAuthenticationFactorsOk returns a tuple with the CustomAuthenticationFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAuthenticationFactors

`func (o *MicrosoftGraphConditionalAccessGrantControls) SetCustomAuthenticationFactors(v []string)`

SetCustomAuthenticationFactors sets CustomAuthenticationFactors field to given value.

### HasCustomAuthenticationFactors

`func (o *MicrosoftGraphConditionalAccessGrantControls) HasCustomAuthenticationFactors() bool`

HasCustomAuthenticationFactors returns a boolean if a field has been set.

### GetOperator

`func (o *MicrosoftGraphConditionalAccessGrantControls) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MicrosoftGraphConditionalAccessGrantControls) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MicrosoftGraphConditionalAccessGrantControls) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MicrosoftGraphConditionalAccessGrantControls) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *MicrosoftGraphConditionalAccessGrantControls) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *MicrosoftGraphConditionalAccessGrantControls) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetTermsOfUse

`func (o *MicrosoftGraphConditionalAccessGrantControls) GetTermsOfUse() []string`

GetTermsOfUse returns the TermsOfUse field if non-nil, zero value otherwise.

### GetTermsOfUseOk

`func (o *MicrosoftGraphConditionalAccessGrantControls) GetTermsOfUseOk() (*[]string, bool)`

GetTermsOfUseOk returns a tuple with the TermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUse

`func (o *MicrosoftGraphConditionalAccessGrantControls) SetTermsOfUse(v []string)`

SetTermsOfUse sets TermsOfUse field to given value.

### HasTermsOfUse

`func (o *MicrosoftGraphConditionalAccessGrantControls) HasTermsOfUse() bool`

HasTermsOfUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


