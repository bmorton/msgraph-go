# ConnectedOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | The description of the connected organization. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the connected organization. | [optional] 
**IdentitySources** | Pointer to [**[]AnyOfobject**](AnyOfobject.md) | The identity sources in this connected organization, one of azureActiveDirectoryTenant, domainIdentitySource or externalDomainFederation. Nullable. | [optional] 
**ModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphConnectedOrganizationState**](anyOf&lt;microsoft.graph.connectedOrganizationState&gt;.md) | The state of a connected organization defines whether assignment policies with requestor scope type AllConfiguredConnectedOrganizationSubjects are applicable or not.  The possible values are: configured, proposed, unknownFutureValue. | [optional] 
**ExternalSponsors** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Nullable. | [optional] 
**InternalSponsors** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Nullable. | [optional] 

## Methods

### NewConnectedOrganization

`func NewConnectedOrganization() *ConnectedOrganization`

NewConnectedOrganization instantiates a new ConnectedOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedOrganizationWithDefaults

`func NewConnectedOrganizationWithDefaults() *ConnectedOrganization`

NewConnectedOrganizationWithDefaults instantiates a new ConnectedOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *ConnectedOrganization) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ConnectedOrganization) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ConnectedOrganization) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *ConnectedOrganization) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *ConnectedOrganization) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *ConnectedOrganization) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *ConnectedOrganization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectedOrganization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectedOrganization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectedOrganization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConnectedOrganization) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConnectedOrganization) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *ConnectedOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectedOrganization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectedOrganization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectedOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ConnectedOrganization) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ConnectedOrganization) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIdentitySources

`func (o *ConnectedOrganization) GetIdentitySources() []*AnyOfobject`

GetIdentitySources returns the IdentitySources field if non-nil, zero value otherwise.

### GetIdentitySourcesOk

`func (o *ConnectedOrganization) GetIdentitySourcesOk() (*[]*AnyOfobject, bool)`

GetIdentitySourcesOk returns a tuple with the IdentitySources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitySources

`func (o *ConnectedOrganization) SetIdentitySources(v []*AnyOfobject)`

SetIdentitySources sets IdentitySources field to given value.

### HasIdentitySources

`func (o *ConnectedOrganization) HasIdentitySources() bool`

HasIdentitySources returns a boolean if a field has been set.

### GetModifiedDateTime

`func (o *ConnectedOrganization) GetModifiedDateTime() time.Time`

GetModifiedDateTime returns the ModifiedDateTime field if non-nil, zero value otherwise.

### GetModifiedDateTimeOk

`func (o *ConnectedOrganization) GetModifiedDateTimeOk() (*time.Time, bool)`

GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDateTime

`func (o *ConnectedOrganization) SetModifiedDateTime(v time.Time)`

SetModifiedDateTime sets ModifiedDateTime field to given value.

### HasModifiedDateTime

`func (o *ConnectedOrganization) HasModifiedDateTime() bool`

HasModifiedDateTime returns a boolean if a field has been set.

### SetModifiedDateTimeNil

`func (o *ConnectedOrganization) SetModifiedDateTimeNil(b bool)`

 SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil

### UnsetModifiedDateTime
`func (o *ConnectedOrganization) UnsetModifiedDateTime()`

UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
### GetState

`func (o *ConnectedOrganization) GetState() AnyOfmicrosoftGraphConnectedOrganizationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectedOrganization) GetStateOk() (*AnyOfmicrosoftGraphConnectedOrganizationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectedOrganization) SetState(v AnyOfmicrosoftGraphConnectedOrganizationState)`

SetState sets State field to given value.

### HasState

`func (o *ConnectedOrganization) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ConnectedOrganization) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ConnectedOrganization) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetExternalSponsors

`func (o *ConnectedOrganization) GetExternalSponsors() []MicrosoftGraphDirectoryObject`

GetExternalSponsors returns the ExternalSponsors field if non-nil, zero value otherwise.

### GetExternalSponsorsOk

`func (o *ConnectedOrganization) GetExternalSponsorsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetExternalSponsorsOk returns a tuple with the ExternalSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSponsors

`func (o *ConnectedOrganization) SetExternalSponsors(v []MicrosoftGraphDirectoryObject)`

SetExternalSponsors sets ExternalSponsors field to given value.

### HasExternalSponsors

`func (o *ConnectedOrganization) HasExternalSponsors() bool`

HasExternalSponsors returns a boolean if a field has been set.

### GetInternalSponsors

`func (o *ConnectedOrganization) GetInternalSponsors() []MicrosoftGraphDirectoryObject`

GetInternalSponsors returns the InternalSponsors field if non-nil, zero value otherwise.

### GetInternalSponsorsOk

`func (o *ConnectedOrganization) GetInternalSponsorsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetInternalSponsorsOk returns a tuple with the InternalSponsors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSponsors

`func (o *ConnectedOrganization) SetInternalSponsors(v []MicrosoftGraphDirectoryObject)`

SetInternalSponsors sets InternalSponsors field to given value.

### HasInternalSponsors

`func (o *ConnectedOrganization) HasInternalSponsors() bool`

HasInternalSponsors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


