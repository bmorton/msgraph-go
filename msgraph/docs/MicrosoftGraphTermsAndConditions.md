# MicrosoftGraphTermsAndConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphTermsAndConditions

`func NewMicrosoftGraphTermsAndConditions() *MicrosoftGraphTermsAndConditions`

NewMicrosoftGraphTermsAndConditions instantiates a new MicrosoftGraphTermsAndConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTermsAndConditionsWithDefaults

`func NewMicrosoftGraphTermsAndConditionsWithDefaults() *MicrosoftGraphTermsAndConditions`

NewMicrosoftGraphTermsAndConditionsWithDefaults instantiates a new MicrosoftGraphTermsAndConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTermsAndConditions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTermsAndConditions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTermsAndConditions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTermsAndConditions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAcceptanceStatement

`func (o *MicrosoftGraphTermsAndConditions) GetAcceptanceStatement() string`

GetAcceptanceStatement returns the AcceptanceStatement field if non-nil, zero value otherwise.

### GetAcceptanceStatementOk

`func (o *MicrosoftGraphTermsAndConditions) GetAcceptanceStatementOk() (*string, bool)`

GetAcceptanceStatementOk returns a tuple with the AcceptanceStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceStatement

`func (o *MicrosoftGraphTermsAndConditions) SetAcceptanceStatement(v string)`

SetAcceptanceStatement sets AcceptanceStatement field to given value.

### HasAcceptanceStatement

`func (o *MicrosoftGraphTermsAndConditions) HasAcceptanceStatement() bool`

HasAcceptanceStatement returns a boolean if a field has been set.

### SetAcceptanceStatementNil

`func (o *MicrosoftGraphTermsAndConditions) SetAcceptanceStatementNil(b bool)`

 SetAcceptanceStatementNil sets the value for AcceptanceStatement to be an explicit nil

### UnsetAcceptanceStatement
`func (o *MicrosoftGraphTermsAndConditions) UnsetAcceptanceStatement()`

UnsetAcceptanceStatement ensures that no value is present for AcceptanceStatement, not even an explicit nil
### GetBodyText

`func (o *MicrosoftGraphTermsAndConditions) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *MicrosoftGraphTermsAndConditions) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *MicrosoftGraphTermsAndConditions) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *MicrosoftGraphTermsAndConditions) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### SetBodyTextNil

`func (o *MicrosoftGraphTermsAndConditions) SetBodyTextNil(b bool)`

 SetBodyTextNil sets the value for BodyText to be an explicit nil

### UnsetBodyText
`func (o *MicrosoftGraphTermsAndConditions) UnsetBodyText()`

UnsetBodyText ensures that no value is present for BodyText, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphTermsAndConditions) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTermsAndConditions) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTermsAndConditions) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTermsAndConditions) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphTermsAndConditions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTermsAndConditions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphTermsAndConditions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphTermsAndConditions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphTermsAndConditions) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphTermsAndConditions) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphTermsAndConditions) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTermsAndConditions) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphTermsAndConditions) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphTermsAndConditions) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphTermsAndConditions) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTermsAndConditions) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTermsAndConditions) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTermsAndConditions) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetTitle

`func (o *MicrosoftGraphTermsAndConditions) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphTermsAndConditions) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphTermsAndConditions) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphTermsAndConditions) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MicrosoftGraphTermsAndConditions) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MicrosoftGraphTermsAndConditions) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetVersion

`func (o *MicrosoftGraphTermsAndConditions) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphTermsAndConditions) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphTermsAndConditions) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphTermsAndConditions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAcceptanceStatuses

`func (o *MicrosoftGraphTermsAndConditions) GetAcceptanceStatuses() []MicrosoftGraphTermsAndConditionsAcceptanceStatus`

GetAcceptanceStatuses returns the AcceptanceStatuses field if non-nil, zero value otherwise.

### GetAcceptanceStatusesOk

`func (o *MicrosoftGraphTermsAndConditions) GetAcceptanceStatusesOk() (*[]MicrosoftGraphTermsAndConditionsAcceptanceStatus, bool)`

GetAcceptanceStatusesOk returns a tuple with the AcceptanceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceStatuses

`func (o *MicrosoftGraphTermsAndConditions) SetAcceptanceStatuses(v []MicrosoftGraphTermsAndConditionsAcceptanceStatus)`

SetAcceptanceStatuses sets AcceptanceStatuses field to given value.

### HasAcceptanceStatuses

`func (o *MicrosoftGraphTermsAndConditions) HasAcceptanceStatuses() bool`

HasAcceptanceStatuses returns a boolean if a field has been set.

### GetAssignments

`func (o *MicrosoftGraphTermsAndConditions) GetAssignments() []MicrosoftGraphTermsAndConditionsAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphTermsAndConditions) GetAssignmentsOk() (*[]MicrosoftGraphTermsAndConditionsAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphTermsAndConditions) SetAssignments(v []MicrosoftGraphTermsAndConditionsAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphTermsAndConditions) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


