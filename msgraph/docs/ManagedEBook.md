# ManagedEBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewManagedEBook

`func NewManagedEBook() *ManagedEBook`

NewManagedEBook instantiates a new ManagedEBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedEBookWithDefaults

`func NewManagedEBookWithDefaults() *ManagedEBook`

NewManagedEBookWithDefaults instantiates a new ManagedEBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *ManagedEBook) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ManagedEBook) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ManagedEBook) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *ManagedEBook) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *ManagedEBook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedEBook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedEBook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagedEBook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ManagedEBook) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ManagedEBook) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *ManagedEBook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManagedEBook) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ManagedEBook) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ManagedEBook) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInformationUrl

`func (o *ManagedEBook) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *ManagedEBook) GetInformationUrlOk() (*string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformationUrl

`func (o *ManagedEBook) SetInformationUrl(v string)`

SetInformationUrl sets InformationUrl field to given value.

### HasInformationUrl

`func (o *ManagedEBook) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrlNil

`func (o *ManagedEBook) SetInformationUrlNil(b bool)`

 SetInformationUrlNil sets the value for InformationUrl to be an explicit nil

### UnsetInformationUrl
`func (o *ManagedEBook) UnsetInformationUrl()`

UnsetInformationUrl ensures that no value is present for InformationUrl, not even an explicit nil
### GetLargeCover

`func (o *ManagedEBook) GetLargeCover() AnyOfmicrosoftGraphMimeContent`

GetLargeCover returns the LargeCover field if non-nil, zero value otherwise.

### GetLargeCoverOk

`func (o *ManagedEBook) GetLargeCoverOk() (*AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeCoverOk returns a tuple with the LargeCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeCover

`func (o *ManagedEBook) SetLargeCover(v AnyOfmicrosoftGraphMimeContent)`

SetLargeCover sets LargeCover field to given value.

### HasLargeCover

`func (o *ManagedEBook) HasLargeCover() bool`

HasLargeCover returns a boolean if a field has been set.

### SetLargeCoverNil

`func (o *ManagedEBook) SetLargeCoverNil(b bool)`

 SetLargeCoverNil sets the value for LargeCover to be an explicit nil

### UnsetLargeCover
`func (o *ManagedEBook) UnsetLargeCover()`

UnsetLargeCover ensures that no value is present for LargeCover, not even an explicit nil
### GetLastModifiedDateTime

`func (o *ManagedEBook) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *ManagedEBook) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *ManagedEBook) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *ManagedEBook) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetPrivacyInformationUrl

`func (o *ManagedEBook) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *ManagedEBook) GetPrivacyInformationUrlOk() (*string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyInformationUrl

`func (o *ManagedEBook) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl sets PrivacyInformationUrl field to given value.

### HasPrivacyInformationUrl

`func (o *ManagedEBook) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrlNil

`func (o *ManagedEBook) SetPrivacyInformationUrlNil(b bool)`

 SetPrivacyInformationUrlNil sets the value for PrivacyInformationUrl to be an explicit nil

### UnsetPrivacyInformationUrl
`func (o *ManagedEBook) UnsetPrivacyInformationUrl()`

UnsetPrivacyInformationUrl ensures that no value is present for PrivacyInformationUrl, not even an explicit nil
### GetPublishedDateTime

`func (o *ManagedEBook) GetPublishedDateTime() time.Time`

GetPublishedDateTime returns the PublishedDateTime field if non-nil, zero value otherwise.

### GetPublishedDateTimeOk

`func (o *ManagedEBook) GetPublishedDateTimeOk() (*time.Time, bool)`

GetPublishedDateTimeOk returns a tuple with the PublishedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDateTime

`func (o *ManagedEBook) SetPublishedDateTime(v time.Time)`

SetPublishedDateTime sets PublishedDateTime field to given value.

### HasPublishedDateTime

`func (o *ManagedEBook) HasPublishedDateTime() bool`

HasPublishedDateTime returns a boolean if a field has been set.

### GetPublisher

`func (o *ManagedEBook) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ManagedEBook) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ManagedEBook) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ManagedEBook) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *ManagedEBook) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *ManagedEBook) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetAssignments

`func (o *ManagedEBook) GetAssignments() []MicrosoftGraphManagedEBookAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *ManagedEBook) GetAssignmentsOk() (*[]MicrosoftGraphManagedEBookAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *ManagedEBook) SetAssignments(v []MicrosoftGraphManagedEBookAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *ManagedEBook) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetDeviceStates

`func (o *ManagedEBook) GetDeviceStates() []MicrosoftGraphDeviceInstallState`

GetDeviceStates returns the DeviceStates field if non-nil, zero value otherwise.

### GetDeviceStatesOk

`func (o *ManagedEBook) GetDeviceStatesOk() (*[]MicrosoftGraphDeviceInstallState, bool)`

GetDeviceStatesOk returns a tuple with the DeviceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStates

`func (o *ManagedEBook) SetDeviceStates(v []MicrosoftGraphDeviceInstallState)`

SetDeviceStates sets DeviceStates field to given value.

### HasDeviceStates

`func (o *ManagedEBook) HasDeviceStates() bool`

HasDeviceStates returns a boolean if a field has been set.

### GetInstallSummary

`func (o *ManagedEBook) GetInstallSummary() AnyOfmicrosoftGraphEBookInstallSummary`

GetInstallSummary returns the InstallSummary field if non-nil, zero value otherwise.

### GetInstallSummaryOk

`func (o *ManagedEBook) GetInstallSummaryOk() (*AnyOfmicrosoftGraphEBookInstallSummary, bool)`

GetInstallSummaryOk returns a tuple with the InstallSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallSummary

`func (o *ManagedEBook) SetInstallSummary(v AnyOfmicrosoftGraphEBookInstallSummary)`

SetInstallSummary sets InstallSummary field to given value.

### HasInstallSummary

`func (o *ManagedEBook) HasInstallSummary() bool`

HasInstallSummary returns a boolean if a field has been set.

### SetInstallSummaryNil

`func (o *ManagedEBook) SetInstallSummaryNil(b bool)`

 SetInstallSummaryNil sets the value for InstallSummary to be an explicit nil

### UnsetInstallSummary
`func (o *ManagedEBook) UnsetInstallSummary()`

UnsetInstallSummary ensures that no value is present for InstallSummary, not even an explicit nil
### GetUserStateSummary

`func (o *ManagedEBook) GetUserStateSummary() []MicrosoftGraphUserInstallStateSummary`

GetUserStateSummary returns the UserStateSummary field if non-nil, zero value otherwise.

### GetUserStateSummaryOk

`func (o *ManagedEBook) GetUserStateSummaryOk() (*[]MicrosoftGraphUserInstallStateSummary, bool)`

GetUserStateSummaryOk returns a tuple with the UserStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStateSummary

`func (o *ManagedEBook) SetUserStateSummary(v []MicrosoftGraphUserInstallStateSummary)`

SetUserStateSummary sets UserStateSummary field to given value.

### HasUserStateSummary

`func (o *ManagedEBook) HasUserStateSummary() bool`

HasUserStateSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


