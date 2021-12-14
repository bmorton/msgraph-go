# MicrosoftGraphManagedEBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The date and time when the eBook file was created. | [optional] 
**Description** | Pointer to **NullableString** | Description. | [optional] 
**DisplayName** | Pointer to **string** | Name of the eBook. | [optional] 
**InformationUrl** | Pointer to **NullableString** | The more information Url. | [optional] 
**LargeCover** | Pointer to [**AnyOfmicrosoftGraphMimeContent**](anyOf&lt;microsoft.graph.mimeContent&gt;.md) | Book cover. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The date and time when the eBook was last modified. | [optional] 
**PrivacyInformationUrl** | Pointer to **NullableString** | The privacy statement Url. | [optional] 
**PublishedDateTime** | Pointer to **time.Time** | The date and time when the eBook was published. | [optional] 
**Publisher** | Pointer to **NullableString** | Publisher. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphManagedEBookAssignment**](MicrosoftGraphManagedEBookAssignment.md) | The list of assignments for this eBook. | [optional] 
**DeviceStates** | Pointer to [**[]MicrosoftGraphDeviceInstallState**](MicrosoftGraphDeviceInstallState.md) | The list of installation states for this eBook. | [optional] 
**InstallSummary** | Pointer to [**AnyOfmicrosoftGraphEBookInstallSummary**](anyOf&lt;microsoft.graph.eBookInstallSummary&gt;.md) | Mobile App Install Summary. | [optional] 
**UserStateSummary** | Pointer to [**[]MicrosoftGraphUserInstallStateSummary**](MicrosoftGraphUserInstallStateSummary.md) | The list of installation states for this eBook. | [optional] 

## Methods

### NewMicrosoftGraphManagedEBook

`func NewMicrosoftGraphManagedEBook() *MicrosoftGraphManagedEBook`

NewMicrosoftGraphManagedEBook instantiates a new MicrosoftGraphManagedEBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphManagedEBookWithDefaults

`func NewMicrosoftGraphManagedEBookWithDefaults() *MicrosoftGraphManagedEBook`

NewMicrosoftGraphManagedEBookWithDefaults instantiates a new MicrosoftGraphManagedEBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphManagedEBook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedEBook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphManagedEBook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphManagedEBook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphManagedEBook) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedEBook) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedEBook) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedEBook) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphManagedEBook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphManagedEBook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphManagedEBook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphManagedEBook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphManagedEBook) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphManagedEBook) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphManagedEBook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedEBook) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedEBook) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphManagedEBook) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInformationUrl

`func (o *MicrosoftGraphManagedEBook) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphManagedEBook) GetInformationUrlOk() (*string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformationUrl

`func (o *MicrosoftGraphManagedEBook) SetInformationUrl(v string)`

SetInformationUrl sets InformationUrl field to given value.

### HasInformationUrl

`func (o *MicrosoftGraphManagedEBook) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrlNil

`func (o *MicrosoftGraphManagedEBook) SetInformationUrlNil(b bool)`

 SetInformationUrlNil sets the value for InformationUrl to be an explicit nil

### UnsetInformationUrl
`func (o *MicrosoftGraphManagedEBook) UnsetInformationUrl()`

UnsetInformationUrl ensures that no value is present for InformationUrl, not even an explicit nil
### GetLargeCover

`func (o *MicrosoftGraphManagedEBook) GetLargeCover() AnyOfmicrosoftGraphMimeContent`

GetLargeCover returns the LargeCover field if non-nil, zero value otherwise.

### GetLargeCoverOk

`func (o *MicrosoftGraphManagedEBook) GetLargeCoverOk() (*AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeCoverOk returns a tuple with the LargeCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeCover

`func (o *MicrosoftGraphManagedEBook) SetLargeCover(v AnyOfmicrosoftGraphMimeContent)`

SetLargeCover sets LargeCover field to given value.

### HasLargeCover

`func (o *MicrosoftGraphManagedEBook) HasLargeCover() bool`

HasLargeCover returns a boolean if a field has been set.

### SetLargeCoverNil

`func (o *MicrosoftGraphManagedEBook) SetLargeCoverNil(b bool)`

 SetLargeCoverNil sets the value for LargeCover to be an explicit nil

### UnsetLargeCover
`func (o *MicrosoftGraphManagedEBook) UnsetLargeCover()`

UnsetLargeCover ensures that no value is present for LargeCover, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedEBook) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedEBook) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedEBook) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedEBook) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetPrivacyInformationUrl

`func (o *MicrosoftGraphManagedEBook) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphManagedEBook) GetPrivacyInformationUrlOk() (*string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphManagedEBook) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl sets PrivacyInformationUrl field to given value.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphManagedEBook) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrlNil

`func (o *MicrosoftGraphManagedEBook) SetPrivacyInformationUrlNil(b bool)`

 SetPrivacyInformationUrlNil sets the value for PrivacyInformationUrl to be an explicit nil

### UnsetPrivacyInformationUrl
`func (o *MicrosoftGraphManagedEBook) UnsetPrivacyInformationUrl()`

UnsetPrivacyInformationUrl ensures that no value is present for PrivacyInformationUrl, not even an explicit nil
### GetPublishedDateTime

`func (o *MicrosoftGraphManagedEBook) GetPublishedDateTime() time.Time`

GetPublishedDateTime returns the PublishedDateTime field if non-nil, zero value otherwise.

### GetPublishedDateTimeOk

`func (o *MicrosoftGraphManagedEBook) GetPublishedDateTimeOk() (*time.Time, bool)`

GetPublishedDateTimeOk returns a tuple with the PublishedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDateTime

`func (o *MicrosoftGraphManagedEBook) SetPublishedDateTime(v time.Time)`

SetPublishedDateTime sets PublishedDateTime field to given value.

### HasPublishedDateTime

`func (o *MicrosoftGraphManagedEBook) HasPublishedDateTime() bool`

HasPublishedDateTime returns a boolean if a field has been set.

### GetPublisher

`func (o *MicrosoftGraphManagedEBook) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphManagedEBook) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *MicrosoftGraphManagedEBook) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *MicrosoftGraphManagedEBook) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *MicrosoftGraphManagedEBook) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *MicrosoftGraphManagedEBook) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetAssignments

`func (o *MicrosoftGraphManagedEBook) GetAssignments() []MicrosoftGraphManagedEBookAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphManagedEBook) GetAssignmentsOk() (*[]MicrosoftGraphManagedEBookAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphManagedEBook) SetAssignments(v []MicrosoftGraphManagedEBookAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphManagedEBook) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetDeviceStates

`func (o *MicrosoftGraphManagedEBook) GetDeviceStates() []MicrosoftGraphDeviceInstallState`

GetDeviceStates returns the DeviceStates field if non-nil, zero value otherwise.

### GetDeviceStatesOk

`func (o *MicrosoftGraphManagedEBook) GetDeviceStatesOk() (*[]MicrosoftGraphDeviceInstallState, bool)`

GetDeviceStatesOk returns a tuple with the DeviceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStates

`func (o *MicrosoftGraphManagedEBook) SetDeviceStates(v []MicrosoftGraphDeviceInstallState)`

SetDeviceStates sets DeviceStates field to given value.

### HasDeviceStates

`func (o *MicrosoftGraphManagedEBook) HasDeviceStates() bool`

HasDeviceStates returns a boolean if a field has been set.

### GetInstallSummary

`func (o *MicrosoftGraphManagedEBook) GetInstallSummary() AnyOfmicrosoftGraphEBookInstallSummary`

GetInstallSummary returns the InstallSummary field if non-nil, zero value otherwise.

### GetInstallSummaryOk

`func (o *MicrosoftGraphManagedEBook) GetInstallSummaryOk() (*AnyOfmicrosoftGraphEBookInstallSummary, bool)`

GetInstallSummaryOk returns a tuple with the InstallSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallSummary

`func (o *MicrosoftGraphManagedEBook) SetInstallSummary(v AnyOfmicrosoftGraphEBookInstallSummary)`

SetInstallSummary sets InstallSummary field to given value.

### HasInstallSummary

`func (o *MicrosoftGraphManagedEBook) HasInstallSummary() bool`

HasInstallSummary returns a boolean if a field has been set.

### SetInstallSummaryNil

`func (o *MicrosoftGraphManagedEBook) SetInstallSummaryNil(b bool)`

 SetInstallSummaryNil sets the value for InstallSummary to be an explicit nil

### UnsetInstallSummary
`func (o *MicrosoftGraphManagedEBook) UnsetInstallSummary()`

UnsetInstallSummary ensures that no value is present for InstallSummary, not even an explicit nil
### GetUserStateSummary

`func (o *MicrosoftGraphManagedEBook) GetUserStateSummary() []MicrosoftGraphUserInstallStateSummary`

GetUserStateSummary returns the UserStateSummary field if non-nil, zero value otherwise.

### GetUserStateSummaryOk

`func (o *MicrosoftGraphManagedEBook) GetUserStateSummaryOk() (*[]MicrosoftGraphUserInstallStateSummary, bool)`

GetUserStateSummaryOk returns a tuple with the UserStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStateSummary

`func (o *MicrosoftGraphManagedEBook) SetUserStateSummary(v []MicrosoftGraphUserInstallStateSummary)`

SetUserStateSummary sets UserStateSummary field to given value.

### HasUserStateSummary

`func (o *MicrosoftGraphManagedEBook) HasUserStateSummary() bool`

HasUserStateSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


