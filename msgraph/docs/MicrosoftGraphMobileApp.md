# MicrosoftGraphMobileApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The date and time the app was created. | [optional] 
**Description** | Pointer to **NullableString** | The description of the app. | [optional] 
**Developer** | Pointer to **NullableString** | The developer of the app. | [optional] 
**DisplayName** | Pointer to **NullableString** | The admin provided or imported title of the app. | [optional] 
**InformationUrl** | Pointer to **NullableString** | The more information Url. | [optional] 
**IsFeatured** | Pointer to **bool** | The value indicating whether the app is marked as featured by the admin. | [optional] 
**LargeIcon** | Pointer to [**AnyOfmicrosoftGraphMimeContent**](anyOf&lt;microsoft.graph.mimeContent&gt;.md) | The large icon, to be displayed in the app details and used for upload of the icon. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The date and time the app was last modified. | [optional] 
**Notes** | Pointer to **NullableString** | Notes for the app. | [optional] 
**Owner** | Pointer to **NullableString** | The owner of the app. | [optional] 
**PrivacyInformationUrl** | Pointer to **NullableString** | The privacy statement Url. | [optional] 
**Publisher** | Pointer to **NullableString** | The publisher of the app. | [optional] 
**PublishingState** | Pointer to [**AnyOfmicrosoftGraphMobileAppPublishingState**](anyOf&lt;microsoft.graph.mobileAppPublishingState&gt;.md) | The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md) | The list of group assignments for this mobile app. | [optional] 
**Categories** | Pointer to [**[]MicrosoftGraphMobileAppCategory**](MicrosoftGraphMobileAppCategory.md) | The list of categories for this app. | [optional] 

## Methods

### NewMicrosoftGraphMobileApp

`func NewMicrosoftGraphMobileApp() *MicrosoftGraphMobileApp`

NewMicrosoftGraphMobileApp instantiates a new MicrosoftGraphMobileApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphMobileAppWithDefaults

`func NewMicrosoftGraphMobileAppWithDefaults() *MicrosoftGraphMobileApp`

NewMicrosoftGraphMobileAppWithDefaults instantiates a new MicrosoftGraphMobileApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphMobileApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMobileApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphMobileApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphMobileApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphMobileApp) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMobileApp) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMobileApp) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphMobileApp) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphMobileApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphMobileApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphMobileApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphMobileApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphMobileApp) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphMobileApp) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDeveloper

`func (o *MicrosoftGraphMobileApp) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MicrosoftGraphMobileApp) GetDeveloperOk() (*string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *MicrosoftGraphMobileApp) SetDeveloper(v string)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *MicrosoftGraphMobileApp) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloperNil

`func (o *MicrosoftGraphMobileApp) SetDeveloperNil(b bool)`

 SetDeveloperNil sets the value for Developer to be an explicit nil

### UnsetDeveloper
`func (o *MicrosoftGraphMobileApp) UnsetDeveloper()`

UnsetDeveloper ensures that no value is present for Developer, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphMobileApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMobileApp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphMobileApp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphMobileApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphMobileApp) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphMobileApp) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetInformationUrl

`func (o *MicrosoftGraphMobileApp) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MicrosoftGraphMobileApp) GetInformationUrlOk() (*string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformationUrl

`func (o *MicrosoftGraphMobileApp) SetInformationUrl(v string)`

SetInformationUrl sets InformationUrl field to given value.

### HasInformationUrl

`func (o *MicrosoftGraphMobileApp) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrlNil

`func (o *MicrosoftGraphMobileApp) SetInformationUrlNil(b bool)`

 SetInformationUrlNil sets the value for InformationUrl to be an explicit nil

### UnsetInformationUrl
`func (o *MicrosoftGraphMobileApp) UnsetInformationUrl()`

UnsetInformationUrl ensures that no value is present for InformationUrl, not even an explicit nil
### GetIsFeatured

`func (o *MicrosoftGraphMobileApp) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MicrosoftGraphMobileApp) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *MicrosoftGraphMobileApp) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.

### HasIsFeatured

`func (o *MicrosoftGraphMobileApp) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### GetLargeIcon

`func (o *MicrosoftGraphMobileApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MicrosoftGraphMobileApp) GetLargeIconOk() (*AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeIcon

`func (o *MicrosoftGraphMobileApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon sets LargeIcon field to given value.

### HasLargeIcon

`func (o *MicrosoftGraphMobileApp) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIconNil

`func (o *MicrosoftGraphMobileApp) SetLargeIconNil(b bool)`

 SetLargeIconNil sets the value for LargeIcon to be an explicit nil

### UnsetLargeIcon
`func (o *MicrosoftGraphMobileApp) UnsetLargeIcon()`

UnsetLargeIcon ensures that no value is present for LargeIcon, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphMobileApp) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphMobileApp) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphMobileApp) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphMobileApp) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetNotes

`func (o *MicrosoftGraphMobileApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftGraphMobileApp) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MicrosoftGraphMobileApp) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MicrosoftGraphMobileApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *MicrosoftGraphMobileApp) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *MicrosoftGraphMobileApp) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOwner

`func (o *MicrosoftGraphMobileApp) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphMobileApp) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MicrosoftGraphMobileApp) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MicrosoftGraphMobileApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *MicrosoftGraphMobileApp) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *MicrosoftGraphMobileApp) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetPrivacyInformationUrl

`func (o *MicrosoftGraphMobileApp) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MicrosoftGraphMobileApp) GetPrivacyInformationUrlOk() (*string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyInformationUrl

`func (o *MicrosoftGraphMobileApp) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl sets PrivacyInformationUrl field to given value.

### HasPrivacyInformationUrl

`func (o *MicrosoftGraphMobileApp) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrlNil

`func (o *MicrosoftGraphMobileApp) SetPrivacyInformationUrlNil(b bool)`

 SetPrivacyInformationUrlNil sets the value for PrivacyInformationUrl to be an explicit nil

### UnsetPrivacyInformationUrl
`func (o *MicrosoftGraphMobileApp) UnsetPrivacyInformationUrl()`

UnsetPrivacyInformationUrl ensures that no value is present for PrivacyInformationUrl, not even an explicit nil
### GetPublisher

`func (o *MicrosoftGraphMobileApp) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphMobileApp) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *MicrosoftGraphMobileApp) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *MicrosoftGraphMobileApp) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *MicrosoftGraphMobileApp) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *MicrosoftGraphMobileApp) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetPublishingState

`func (o *MicrosoftGraphMobileApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MicrosoftGraphMobileApp) GetPublishingStateOk() (*AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingState

`func (o *MicrosoftGraphMobileApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState sets PublishingState field to given value.

### HasPublishingState

`func (o *MicrosoftGraphMobileApp) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingStateNil

`func (o *MicrosoftGraphMobileApp) SetPublishingStateNil(b bool)`

 SetPublishingStateNil sets the value for PublishingState to be an explicit nil

### UnsetPublishingState
`func (o *MicrosoftGraphMobileApp) UnsetPublishingState()`

UnsetPublishingState ensures that no value is present for PublishingState, not even an explicit nil
### GetAssignments

`func (o *MicrosoftGraphMobileApp) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphMobileApp) GetAssignmentsOk() (*[]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MicrosoftGraphMobileApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MicrosoftGraphMobileApp) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetCategories

`func (o *MicrosoftGraphMobileApp) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphMobileApp) GetCategoriesOk() (*[]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MicrosoftGraphMobileApp) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MicrosoftGraphMobileApp) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


