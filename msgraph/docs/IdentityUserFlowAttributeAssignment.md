# IdentityUserFlowAttributeAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The display name of the identityUserFlowAttribute within a user flow. | [optional] 
**IsOptional** | Pointer to **bool** | Determines whether the identityUserFlowAttribute is optional. true means the user doesn&#39;t have to provide a value. false means the user cannot complete sign-up without providing a value. | [optional] 
**RequiresVerification** | Pointer to **bool** | Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user&#39;s phone number or email address. | [optional] 
**UserAttributeValues** | Pointer to [**[]AnyOfmicrosoftGraphUserAttributeValuesItem**](AnyOfmicrosoftGraphUserAttributeValuesItem.md) | The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect. | [optional] 
**UserInputType** | Pointer to [**AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType**](anyOf&lt;microsoft.graph.identityUserFlowAttributeInputType&gt;.md) | The input type of the user flow attribute. Possible values are: textBox, dateTimeDropdown, radioSingleSelect, dropdownSingleSelect, emailBox, checkboxMultiSelect. | [optional] 
**UserAttribute** | Pointer to [**AnyOfmicrosoftGraphIdentityUserFlowAttribute**](anyOf&lt;microsoft.graph.identityUserFlowAttribute&gt;.md) | The user attribute that you want to add to your user flow. | [optional] 

## Methods

### NewIdentityUserFlowAttributeAssignment

`func NewIdentityUserFlowAttributeAssignment() *IdentityUserFlowAttributeAssignment`

NewIdentityUserFlowAttributeAssignment instantiates a new IdentityUserFlowAttributeAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUserFlowAttributeAssignmentWithDefaults

`func NewIdentityUserFlowAttributeAssignmentWithDefaults() *IdentityUserFlowAttributeAssignment`

NewIdentityUserFlowAttributeAssignmentWithDefaults instantiates a new IdentityUserFlowAttributeAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *IdentityUserFlowAttributeAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityUserFlowAttributeAssignment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityUserFlowAttributeAssignment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityUserFlowAttributeAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IdentityUserFlowAttributeAssignment) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IdentityUserFlowAttributeAssignment) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsOptional

`func (o *IdentityUserFlowAttributeAssignment) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *IdentityUserFlowAttributeAssignment) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *IdentityUserFlowAttributeAssignment) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *IdentityUserFlowAttributeAssignment) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### GetRequiresVerification

`func (o *IdentityUserFlowAttributeAssignment) GetRequiresVerification() bool`

GetRequiresVerification returns the RequiresVerification field if non-nil, zero value otherwise.

### GetRequiresVerificationOk

`func (o *IdentityUserFlowAttributeAssignment) GetRequiresVerificationOk() (*bool, bool)`

GetRequiresVerificationOk returns a tuple with the RequiresVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresVerification

`func (o *IdentityUserFlowAttributeAssignment) SetRequiresVerification(v bool)`

SetRequiresVerification sets RequiresVerification field to given value.

### HasRequiresVerification

`func (o *IdentityUserFlowAttributeAssignment) HasRequiresVerification() bool`

HasRequiresVerification returns a boolean if a field has been set.

### GetUserAttributeValues

`func (o *IdentityUserFlowAttributeAssignment) GetUserAttributeValues() []*AnyOfmicrosoftGraphUserAttributeValuesItem`

GetUserAttributeValues returns the UserAttributeValues field if non-nil, zero value otherwise.

### GetUserAttributeValuesOk

`func (o *IdentityUserFlowAttributeAssignment) GetUserAttributeValuesOk() (*[]*AnyOfmicrosoftGraphUserAttributeValuesItem, bool)`

GetUserAttributeValuesOk returns a tuple with the UserAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeValues

`func (o *IdentityUserFlowAttributeAssignment) SetUserAttributeValues(v []*AnyOfmicrosoftGraphUserAttributeValuesItem)`

SetUserAttributeValues sets UserAttributeValues field to given value.

### HasUserAttributeValues

`func (o *IdentityUserFlowAttributeAssignment) HasUserAttributeValues() bool`

HasUserAttributeValues returns a boolean if a field has been set.

### GetUserInputType

`func (o *IdentityUserFlowAttributeAssignment) GetUserInputType() AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType`

GetUserInputType returns the UserInputType field if non-nil, zero value otherwise.

### GetUserInputTypeOk

`func (o *IdentityUserFlowAttributeAssignment) GetUserInputTypeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType, bool)`

GetUserInputTypeOk returns a tuple with the UserInputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputType

`func (o *IdentityUserFlowAttributeAssignment) SetUserInputType(v AnyOfmicrosoftGraphIdentityUserFlowAttributeInputType)`

SetUserInputType sets UserInputType field to given value.

### HasUserInputType

`func (o *IdentityUserFlowAttributeAssignment) HasUserInputType() bool`

HasUserInputType returns a boolean if a field has been set.

### SetUserInputTypeNil

`func (o *IdentityUserFlowAttributeAssignment) SetUserInputTypeNil(b bool)`

 SetUserInputTypeNil sets the value for UserInputType to be an explicit nil

### UnsetUserInputType
`func (o *IdentityUserFlowAttributeAssignment) UnsetUserInputType()`

UnsetUserInputType ensures that no value is present for UserInputType, not even an explicit nil
### GetUserAttribute

`func (o *IdentityUserFlowAttributeAssignment) GetUserAttribute() AnyOfmicrosoftGraphIdentityUserFlowAttribute`

GetUserAttribute returns the UserAttribute field if non-nil, zero value otherwise.

### GetUserAttributeOk

`func (o *IdentityUserFlowAttributeAssignment) GetUserAttributeOk() (*AnyOfmicrosoftGraphIdentityUserFlowAttribute, bool)`

GetUserAttributeOk returns a tuple with the UserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttribute

`func (o *IdentityUserFlowAttributeAssignment) SetUserAttribute(v AnyOfmicrosoftGraphIdentityUserFlowAttribute)`

SetUserAttribute sets UserAttribute field to given value.

### HasUserAttribute

`func (o *IdentityUserFlowAttributeAssignment) HasUserAttribute() bool`

HasUserAttribute returns a boolean if a field has been set.

### SetUserAttributeNil

`func (o *IdentityUserFlowAttributeAssignment) SetUserAttributeNil(b bool)`

 SetUserAttributeNil sets the value for UserAttribute to be an explicit nil

### UnsetUserAttribute
`func (o *IdentityUserFlowAttributeAssignment) UnsetUserAttribute()`

UnsetUserAttribute ensures that no value is present for UserAttribute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


