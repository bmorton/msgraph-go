# TermsAndConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptanceStatement** | Pointer to **NullableString** | Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&amp;C policy. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**BodyText** | Pointer to **NullableString** | Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | DateTime the object was created. | [optional] 
**Description** | Pointer to **NullableString** | Administrator-supplied description of the T&amp;C policy. | [optional] 
**DisplayName** | Pointer to **string** | Administrator-supplied name for the T&amp;C policy. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | DateTime the object was last modified. | [optional] 
**Title** | Pointer to **NullableString** | Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&amp;C policy. | [optional] 
**Version** | Pointer to **int32** | Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&amp;C policy. | [optional] 
**AcceptanceStatuses** | Pointer to [**[]MicrosoftGraphTermsAndConditionsAcceptanceStatus**](MicrosoftGraphTermsAndConditionsAcceptanceStatus.md) | The list of acceptance statuses for this T&amp;C policy. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphTermsAndConditionsAssignment**](MicrosoftGraphTermsAndConditionsAssignment.md) | The list of assignments for this T&amp;C policy. | [optional] 

## Methods

### NewTermsAndConditions

`func NewTermsAndConditions() *TermsAndConditions`

NewTermsAndConditions instantiates a new TermsAndConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermsAndConditionsWithDefaults

`func NewTermsAndConditionsWithDefaults() *TermsAndConditions`

NewTermsAndConditionsWithDefaults instantiates a new TermsAndConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptanceStatement

`func (o *TermsAndConditions) GetAcceptanceStatement() string`

GetAcceptanceStatement returns the AcceptanceStatement field if non-nil, zero value otherwise.

### GetAcceptanceStatementOk

`func (o *TermsAndConditions) GetAcceptanceStatementOk() (*string, bool)`

GetAcceptanceStatementOk returns a tuple with the AcceptanceStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceStatement

`func (o *TermsAndConditions) SetAcceptanceStatement(v string)`

SetAcceptanceStatement sets AcceptanceStatement field to given value.

### HasAcceptanceStatement

`func (o *TermsAndConditions) HasAcceptanceStatement() bool`

HasAcceptanceStatement returns a boolean if a field has been set.

### SetAcceptanceStatementNil

`func (o *TermsAndConditions) SetAcceptanceStatementNil(b bool)`

 SetAcceptanceStatementNil sets the value for AcceptanceStatement to be an explicit nil

### UnsetAcceptanceStatement
`func (o *TermsAndConditions) UnsetAcceptanceStatement()`

UnsetAcceptanceStatement ensures that no value is present for AcceptanceStatement, not even an explicit nil
### GetBodyText

`func (o *TermsAndConditions) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *TermsAndConditions) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *TermsAndConditions) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *TermsAndConditions) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### SetBodyTextNil

`func (o *TermsAndConditions) SetBodyTextNil(b bool)`

 SetBodyTextNil sets the value for BodyText to be an explicit nil

### UnsetBodyText
`func (o *TermsAndConditions) UnsetBodyText()`

UnsetBodyText ensures that no value is present for BodyText, not even an explicit nil
### GetCreatedDateTime

`func (o *TermsAndConditions) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *TermsAndConditions) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *TermsAndConditions) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *TermsAndConditions) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *TermsAndConditions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TermsAndConditions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TermsAndConditions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TermsAndConditions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TermsAndConditions) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TermsAndConditions) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *TermsAndConditions) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TermsAndConditions) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TermsAndConditions) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TermsAndConditions) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *TermsAndConditions) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *TermsAndConditions) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *TermsAndConditions) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *TermsAndConditions) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetTitle

`func (o *TermsAndConditions) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TermsAndConditions) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TermsAndConditions) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TermsAndConditions) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TermsAndConditions) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TermsAndConditions) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetVersion

`func (o *TermsAndConditions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TermsAndConditions) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TermsAndConditions) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TermsAndConditions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAcceptanceStatuses

`func (o *TermsAndConditions) GetAcceptanceStatuses() []MicrosoftGraphTermsAndConditionsAcceptanceStatus`

GetAcceptanceStatuses returns the AcceptanceStatuses field if non-nil, zero value otherwise.

### GetAcceptanceStatusesOk

`func (o *TermsAndConditions) GetAcceptanceStatusesOk() (*[]MicrosoftGraphTermsAndConditionsAcceptanceStatus, bool)`

GetAcceptanceStatusesOk returns a tuple with the AcceptanceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceStatuses

`func (o *TermsAndConditions) SetAcceptanceStatuses(v []MicrosoftGraphTermsAndConditionsAcceptanceStatus)`

SetAcceptanceStatuses sets AcceptanceStatuses field to given value.

### HasAcceptanceStatuses

`func (o *TermsAndConditions) HasAcceptanceStatuses() bool`

HasAcceptanceStatuses returns a boolean if a field has been set.

### GetAssignments

`func (o *TermsAndConditions) GetAssignments() []MicrosoftGraphTermsAndConditionsAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *TermsAndConditions) GetAssignmentsOk() (*[]MicrosoftGraphTermsAndConditionsAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *TermsAndConditions) SetAssignments(v []MicrosoftGraphTermsAndConditionsAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *TermsAndConditions) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


