# MobileApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewMobileApp

`func NewMobileApp() *MobileApp`

NewMobileApp instantiates a new MobileApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileAppWithDefaults

`func NewMobileAppWithDefaults() *MobileApp`

NewMobileAppWithDefaults instantiates a new MobileApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *MobileApp) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MobileApp) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MobileApp) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MobileApp) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MobileApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MobileApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MobileApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MobileApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MobileApp) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MobileApp) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDeveloper

`func (o *MobileApp) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *MobileApp) GetDeveloperOk() (*string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *MobileApp) SetDeveloper(v string)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *MobileApp) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### SetDeveloperNil

`func (o *MobileApp) SetDeveloperNil(b bool)`

 SetDeveloperNil sets the value for Developer to be an explicit nil

### UnsetDeveloper
`func (o *MobileApp) UnsetDeveloper()`

UnsetDeveloper ensures that no value is present for Developer, not even an explicit nil
### GetDisplayName

`func (o *MobileApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileApp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileApp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MobileApp) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MobileApp) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetInformationUrl

`func (o *MobileApp) GetInformationUrl() string`

GetInformationUrl returns the InformationUrl field if non-nil, zero value otherwise.

### GetInformationUrlOk

`func (o *MobileApp) GetInformationUrlOk() (*string, bool)`

GetInformationUrlOk returns a tuple with the InformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformationUrl

`func (o *MobileApp) SetInformationUrl(v string)`

SetInformationUrl sets InformationUrl field to given value.

### HasInformationUrl

`func (o *MobileApp) HasInformationUrl() bool`

HasInformationUrl returns a boolean if a field has been set.

### SetInformationUrlNil

`func (o *MobileApp) SetInformationUrlNil(b bool)`

 SetInformationUrlNil sets the value for InformationUrl to be an explicit nil

### UnsetInformationUrl
`func (o *MobileApp) UnsetInformationUrl()`

UnsetInformationUrl ensures that no value is present for InformationUrl, not even an explicit nil
### GetIsFeatured

`func (o *MobileApp) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *MobileApp) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *MobileApp) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.

### HasIsFeatured

`func (o *MobileApp) HasIsFeatured() bool`

HasIsFeatured returns a boolean if a field has been set.

### GetLargeIcon

`func (o *MobileApp) GetLargeIcon() AnyOfmicrosoftGraphMimeContent`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *MobileApp) GetLargeIconOk() (*AnyOfmicrosoftGraphMimeContent, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeIcon

`func (o *MobileApp) SetLargeIcon(v AnyOfmicrosoftGraphMimeContent)`

SetLargeIcon sets LargeIcon field to given value.

### HasLargeIcon

`func (o *MobileApp) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### SetLargeIconNil

`func (o *MobileApp) SetLargeIconNil(b bool)`

 SetLargeIconNil sets the value for LargeIcon to be an explicit nil

### UnsetLargeIcon
`func (o *MobileApp) UnsetLargeIcon()`

UnsetLargeIcon ensures that no value is present for LargeIcon, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MobileApp) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MobileApp) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MobileApp) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MobileApp) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetNotes

`func (o *MobileApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MobileApp) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MobileApp) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MobileApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *MobileApp) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *MobileApp) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOwner

`func (o *MobileApp) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MobileApp) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MobileApp) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MobileApp) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *MobileApp) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *MobileApp) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetPrivacyInformationUrl

`func (o *MobileApp) GetPrivacyInformationUrl() string`

GetPrivacyInformationUrl returns the PrivacyInformationUrl field if non-nil, zero value otherwise.

### GetPrivacyInformationUrlOk

`func (o *MobileApp) GetPrivacyInformationUrlOk() (*string, bool)`

GetPrivacyInformationUrlOk returns a tuple with the PrivacyInformationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyInformationUrl

`func (o *MobileApp) SetPrivacyInformationUrl(v string)`

SetPrivacyInformationUrl sets PrivacyInformationUrl field to given value.

### HasPrivacyInformationUrl

`func (o *MobileApp) HasPrivacyInformationUrl() bool`

HasPrivacyInformationUrl returns a boolean if a field has been set.

### SetPrivacyInformationUrlNil

`func (o *MobileApp) SetPrivacyInformationUrlNil(b bool)`

 SetPrivacyInformationUrlNil sets the value for PrivacyInformationUrl to be an explicit nil

### UnsetPrivacyInformationUrl
`func (o *MobileApp) UnsetPrivacyInformationUrl()`

UnsetPrivacyInformationUrl ensures that no value is present for PrivacyInformationUrl, not even an explicit nil
### GetPublisher

`func (o *MobileApp) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MobileApp) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *MobileApp) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *MobileApp) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *MobileApp) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *MobileApp) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetPublishingState

`func (o *MobileApp) GetPublishingState() AnyOfmicrosoftGraphMobileAppPublishingState`

GetPublishingState returns the PublishingState field if non-nil, zero value otherwise.

### GetPublishingStateOk

`func (o *MobileApp) GetPublishingStateOk() (*AnyOfmicrosoftGraphMobileAppPublishingState, bool)`

GetPublishingStateOk returns a tuple with the PublishingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingState

`func (o *MobileApp) SetPublishingState(v AnyOfmicrosoftGraphMobileAppPublishingState)`

SetPublishingState sets PublishingState field to given value.

### HasPublishingState

`func (o *MobileApp) HasPublishingState() bool`

HasPublishingState returns a boolean if a field has been set.

### SetPublishingStateNil

`func (o *MobileApp) SetPublishingStateNil(b bool)`

 SetPublishingStateNil sets the value for PublishingState to be an explicit nil

### UnsetPublishingState
`func (o *MobileApp) UnsetPublishingState()`

UnsetPublishingState ensures that no value is present for PublishingState, not even an explicit nil
### GetAssignments

`func (o *MobileApp) GetAssignments() []MicrosoftGraphMobileAppAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MobileApp) GetAssignmentsOk() (*[]MicrosoftGraphMobileAppAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *MobileApp) SetAssignments(v []MicrosoftGraphMobileAppAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *MobileApp) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetCategories

`func (o *MobileApp) GetCategories() []MicrosoftGraphMobileAppCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MobileApp) GetCategoriesOk() (*[]MicrosoftGraphMobileAppCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MobileApp) SetCategories(v []MicrosoftGraphMobileAppCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MobileApp) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


