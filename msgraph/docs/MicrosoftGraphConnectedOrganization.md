# MicrosoftGraphConnectedOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | The description of the connected organization. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the connected organization. | [optional] 
**IdentitySources** | Pointer to [**[]AnyOfobject**](AnyOfobject.md) | The identity sources in this connected organization, one of azureActiveDirectoryTenant, domainIdentitySource or externalDomainFederation. Nullable. | [optional] 
**ModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphConnectedOrganizationState**](anyOf&lt;microsoft.graph.connectedOrganizationState&gt;.md) | The state of a connected organization defines whether assignment policies with requestor scope type AllConfiguredConnectedOrganizationSubjects are applicable or not.  The possible values are: configured, proposed, unknownFutureValue. | [optional] 
**ExternalSponsors** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Nullable. | [optional] 
**InternalSponsors** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Nullable. | [optional] 

## Methods

### NewMicrosoftGraphConnectedOrganization

`func NewMicrosoftGraphConnectedOrganization() *MicrosoftGraphConnectedOrganization`

NewMicrosoftGraphConnectedOrganization instantiates a new MicrosoftGraphConnectedOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConnectedOrganizationWithDefaults

`func NewMicrosoftGraphConnectedOrganizationWithDefaults() *MicrosoftGraphConnectedOrganization`

NewMicrosoftGraphConnectedOrganizationWithDefaults instantiates a new MicrosoftGraphConnectedOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphConnectedOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphConnectedOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphConnectedOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphConnectedOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphConnectedOrganization) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphConnectedOrganization) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphConnectedOrganization) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphConnectedOrganization) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphConnectedOrganization) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphConnectedOrganization) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphConnectedOrganization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphConnectedOrganization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphConnectedOrganization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphConnectedOrganization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphConnectedOrganization) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphConnectedOrganization) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphConnectedOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphConnectedOrganization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphConnectedOrganization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphConnectedOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphConnectedOrganization) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphConnectedOrganization) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIdentitySources

`func (o *MicrosoftGraphConnectedOrganization) GetIdentitySources() []*AnyOfobject`

GetIdentitySources returns the IdentitySources field if non-nil, zero value otherwise.

### GetIdentitySourcesOk

`func (o *MicrosoftGraphConnectedOrganization) GetIdentitySourcesOk() (*[]*AnyOfobject, bool)`

GetIdentitySourcesOk returns a tuple with the IdentitySources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitySources

`func (o *MicrosoftGraphConnectedOrganization) SetIdentitySources(v []*AnyOfobject)`

SetIdentitySources sets IdentitySources field to given value.

### HasIdentitySources

`func (o *MicrosoftGraphConnectedOrganization) HasIdentitySources() bool`

HasIdentitySources returns a boolean if a field has been set.

### GetModifiedDateTime

`func (o *MicrosoftGraphConnectedOrganization) GetModifiedDateTime() time.Time`

GetModifiedDateTime returns the ModifiedDateTime field if non-nil, zero value otherwise.

### GetModifiedDateTimeOk

`func (o *MicrosoftGraphConnectedOrganization) GetModifiedDateTimeOk() (*time.Time, bool)`

GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDateTime

`func (o *MicrosoftGraphConnectedOrganization) SetModifiedDateTime(v time.Time)`

SetModifiedDateTime sets ModifiedDateTime field to given value.

### HasModifiedDateTime

`func (o *MicrosoftGraphConnectedOrganization) HasModifiedDateTime() bool`

HasModifiedDateTime returns a boolean if a field has been set.

### SetModifiedDateTimeNil

`func (o *MicrosoftGraphConnectedOrganization) SetModifiedDateTimeNil(b bool)`

 SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil

### UnsetModifiedDateTime
`func (o *MicrosoftGraphConnectedOrganization) UnsetModifiedDateTime()`

UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
### GetState

`func (o *MicrosoftGraphConnectedOrganization) GetState() AnyOfmicrosoftGraphConnectedOrganizationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphConnectedOrganization) GetStateOk() (*AnyOfmicrosoftGraphConnectedOrganizationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphConnectedOrganization) SetState(v AnyOfmicrosoftGraphConnectedOrganizationState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphConnectedOrganization) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphConnectedOrganization) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphConnectedOrganization) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetExternalSponsors

`func (o *MicrosoftGraphConnectedOrganization) GetExternalSponsors() []MicrosoftGraphDirectoryObject`

GetExternalSponsors returns the ExternalSponsors field if non-nil, zero value otherwise.

### GetExternalSponsorsOk

`func (o *MicrosoftGraphConnectedOrganization) GetExternalSponsorsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetExternalSponsorsOk returns a tuple with the ExternalSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSponsors

`func (o *MicrosoftGraphConnectedOrganization) SetExternalSponsors(v []MicrosoftGraphDirectoryObject)`

SetExternalSponsors sets ExternalSponsors field to given value.

### HasExternalSponsors

`func (o *MicrosoftGraphConnectedOrganization) HasExternalSponsors() bool`

HasExternalSponsors returns a boolean if a field has been set.

### GetInternalSponsors

`func (o *MicrosoftGraphConnectedOrganization) GetInternalSponsors() []MicrosoftGraphDirectoryObject`

GetInternalSponsors returns the InternalSponsors field if non-nil, zero value otherwise.

### GetInternalSponsorsOk

`func (o *MicrosoftGraphConnectedOrganization) GetInternalSponsorsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetInternalSponsorsOk returns a tuple with the InternalSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSponsors

`func (o *MicrosoftGraphConnectedOrganization) SetInternalSponsors(v []MicrosoftGraphDirectoryObject)`

SetInternalSponsors sets InternalSponsors field to given value.

### HasInternalSponsors

`func (o *MicrosoftGraphConnectedOrganization) HasInternalSponsors() bool`

HasInternalSponsors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


