# MicrosoftGraphIdentityUserFlowAttributeAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the identityUserFlowAttribute within a user flow. | [optional] 
**IsOptional** | Pointer to **bool** | Determines whether the identityUserFlowAttribute is optional. true means the user doesn&#39;t have to provide a value. false means the user cannot complete sign-up without providing a value. | [optional] 
**RequiresVerification** | Pointer to **bool** | Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user&#39;s phone number or email address. | [optional] 
**UserAttributeValues** | Pointer to [**[]AnyOfmicrosoftGraphUserAttributeValuesItem**](AnyOfmicrosoftGraphUserAttributeValuesItem.md) | The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect. | [optional] 
**UserInputType** | Pointer to [**AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType**](anyOf&lt;microsoft.graph.identityUserFlowAttributeInputType&gt;.md) | The input type of the user flow attribute. Possible values are: textBox, dateTimeDropdown, radioSingleSelect, dropdownSingleSelect, emailBox, checkboxMultiSelect. | [optional] 
**UserAttribute** | Pointer to [**AnyOfmicrosoftGraphIdentityUserFlowAttribute**](anyOf&lt;microsoft.graph.identityUserFlowAttribute&gt;.md) | The user attribute that you want to add to your user flow. | [optional] 

## Methods

### NewMicrosoftGraphIdentityUserFlowAttributeAssignment

`func NewMicrosoftGraphIdentityUserFlowAttributeAssignment() *MicrosoftGraphIdentityUserFlowAttributeAssignment`

NewMicrosoftGraphIdentityUserFlowAttributeAssignment instantiates a new MicrosoftGraphIdentityUserFlowAttributeAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIdentityUserFlowAttributeAssignmentWithDefaults

`func NewMicrosoftGraphIdentityUserFlowAttributeAssignmentWithDefaults() *MicrosoftGraphIdentityUserFlowAttributeAssignment`

NewMicrosoftGraphIdentityUserFlowAttributeAssignmentWithDefaults instantiates a new MicrosoftGraphIdentityUserFlowAttributeAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsOptional

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### GetRequiresVerification

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetRequiresVerification() bool`

GetRequiresVerification returns the RequiresVerification field if non-nil, zero value otherwise.

### GetRequiresVerificationOk

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetRequiresVerificationOk() (*bool, bool)`

GetRequiresVerificationOk returns a tuple with the RequiresVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresVerification

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetRequiresVerification(v bool)`

SetRequiresVerification sets RequiresVerification field to given value.

### HasRequiresVerification

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasRequiresVerification() bool`

HasRequiresVerification returns a boolean if a field has been set.

### GetUserAttributeValues

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserAttributeValues() []*AnyOfmicrosoftGraphUserAttributeValuesItem`

GetUserAttributeValues returns the UserAttributeValues field if non-nil, zero value otherwise.

### GetUserAttributeValuesOk

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserAttributeValuesOk() (*[]*AnyOfmicrosoftGraphUserAttributeValuesItem, bool)`

GetUserAttributeValuesOk returns a tuple with the UserAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeValues

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetUserAttributeValues(v []*AnyOfmicrosoftGraphUserAttributeValuesItem)`

SetUserAttributeValues sets UserAttributeValues field to given value.

### HasUserAttributeValues

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasUserAttributeValues() bool`

HasUserAttributeValues returns a boolean if a field has been set.

### GetUserInputType

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserInputType() AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType`

GetUserInputType returns the UserInputType field if non-nil, zero value otherwise.

### GetUserInputTypeOk

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserInputTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType, bool)`

GetUserInputTypeOk returns a tuple with the UserInputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputType

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetUserInputType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType)`

SetUserInputType sets UserInputType field to given value.

### HasUserInputType

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasUserInputType() bool`

HasUserInputType returns a boolean if a field has been set.

### SetUserInputTypeNil

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetUserInputTypeNil(b bool)`

 SetUserInputTypeNil sets the value for UserInputType to be an explicit nil

### UnsetUserInputType
`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) UnsetUserInputType()`

UnsetUserInputType ensures that no value is present for UserInputType, not even an explicit nil
### GetUserAttribute

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserAttribute() AnyOfmicrosoftGraphIdentityUserFlowAttribute`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) GetUserAttributeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttribute, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetUserAttribute(v AnyOfmicrosoftGraphIdentityUserFlowAttribute)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### SetUserAttributeNil

`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) SetUserAttributeNil(b bool)`

 SetUserAttributeNil sets the value for UserAttribute to be an explicit nil

### UnsetUserAttribute
`func (o *MicrosoftGraphIdentityUserFlowAttributeAssignment) UnsetUserAttribute()`

UnsetUserAttribute ensures that no value is present for UserAttribute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


