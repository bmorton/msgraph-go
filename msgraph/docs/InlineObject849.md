# InlineObject849

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**MailNickname** | Pointer to **NullableString** |  | [optional] 
**Classification** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to [**AnyOfmicrosoftGraphTeamVisibilityType**](anyOf&lt;microsoft.graph.teamVisibilityType&gt;.md) |  | [optional] 
**PartsToClone** | Pointer to [**AnyOfmicrosoftGraphClonableTeamParts**](anyOf&lt;microsoft.graph.clonableTeamParts&gt;.md) |  | [optional] 

## Methods

### NewInlineObject849

`func NewInlineObject849() *InlineObject849`

NewInlineObject849 instantiates a new InlineObject849 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject849WithDefaults

`func NewInlineObject849WithDefaults() *InlineObject849`

NewInlineObject849WithDefaults instantiates a new InlineObject849 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *InlineObject849) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineObject849) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineObject849) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InlineObject849) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *InlineObject849) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *InlineObject849) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *InlineObject849) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject849) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject849) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject849) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineObject849) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineObject849) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMailNickname

`func (o *InlineObject849) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *InlineObject849) GetMailNicknameOk() (*string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNickname

`func (o *InlineObject849) SetMailNickname(v string)`

SetMailNickname sets MailNickname field to given value.

### HasMailNickname

`func (o *InlineObject849) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNicknameNil

`func (o *InlineObject849) SetMailNicknameNil(b bool)`

 SetMailNicknameNil sets the value for MailNickname to be an explicit nil

### UnsetMailNickname
`func (o *InlineObject849) UnsetMailNickname()`

UnsetMailNickname ensures that no value is present for MailNickname, not even an explicit nil
### GetClassification

`func (o *InlineObject849) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *InlineObject849) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *InlineObject849) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *InlineObject849) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *InlineObject849) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *InlineObject849) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetVisibility

`func (o *InlineObject849) GetVisibility() AnyOfmicrosoftGraphTeamVisibilityType`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineObject849) GetVisibilityOk() (*AnyOfmicrosoftGraphTeamVisibilityType, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineObject849) SetVisibility(v AnyOfmicrosoftGraphTeamVisibilityType)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineObject849) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *InlineObject849) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *InlineObject849) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetPartsToClone

`func (o *InlineObject849) GetPartsToClone() AnyOfmicrosoftGraphClonableTeamParts`

GetPartsToClone returns the PartsToClone field if non-nil, zero value otherwise.

### GetPartsToCloneOk

`func (o *InlineObject849) GetPartsToCloneOk() (*AnyOfmicrosoftGraphClonableTeamParts, bool)`

GetPartsToCloneOk returns a tuple with the PartsToClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsToClone

`func (o *InlineObject849) SetPartsToClone(v AnyOfmicrosoftGraphClonableTeamParts)`

SetPartsToClone sets PartsToClone field to given value.

### HasPartsToClone

`func (o *InlineObject849) HasPartsToClone() bool`

HasPartsToClone returns a boolean if a field has been set.

### SetPartsToCloneNil

`func (o *InlineObject849) SetPartsToCloneNil(b bool)`

 SetPartsToCloneNil sets the value for PartsToClone to be an explicit nil

### UnsetPartsToClone
`func (o *InlineObject849) UnsetPartsToClone()`

UnsetPartsToClone ensures that no value is present for PartsToClone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


