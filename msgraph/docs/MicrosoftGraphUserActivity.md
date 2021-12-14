# MicrosoftGraphUserActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ActivationUrl** | Pointer to **string** | Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a web-based app if no native app exists. | [optional] 
**ActivitySourceHost** | Pointer to **string** | Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app identity. For example, a separate file and domain is needed for Word vs. PowerPoint. | [optional] 
**AppActivityId** | Pointer to **string** | Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter. | [optional] 
**AppDisplayName** | Pointer to **NullableString** | Optional. Short text description of the app used to generate the activity for use in cases when the app is not installed on the user’s local device. | [optional] 
**ContentInfo** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Optional. A custom piece of data - JSON-LD extensible description of content according to schema.org syntax. | [optional] 
**ContentUrl** | Pointer to **NullableString** | Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example, a pointer to an item in an RSS feed). | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Set by the server. DateTime in UTC when the object was created on the server. | [optional] 
**ExpirationDateTime** | Pointer to **NullableTime** | Set by the server. DateTime in UTC when the object expired on the server. | [optional] 
**FallbackUrl** | Pointer to **NullableString** | Optional. URL used to launch the activity in a web-based app, if available. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Set by the server. DateTime in UTC when the object was modified on the server. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphStatus**](anyOf&lt;microsoft.graph.status&gt;.md) | Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored. | [optional] 
**UserTimezone** | Pointer to **NullableString** | Optional. The timezone in which the user&#39;s device used to generate the activity was located at activity creation time; values supplied as Olson IDs in order to support cross-platform representation. | [optional] 
**VisualElements** | Pointer to [**MicrosoftGraphVisualInfo**](MicrosoftGraphVisualInfo.md) |  | [optional] 
**HistoryItems** | Pointer to [**[]MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md) | Optional. NavigationProperty/Containment; navigation property to the activity&#39;s historyItems. | [optional] 

## Methods

### NewMicrosoftGraphUserActivity

`func NewMicrosoftGraphUserActivity() *MicrosoftGraphUserActivity`

NewMicrosoftGraphUserActivity instantiates a new MicrosoftGraphUserActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserActivityWithDefaults

`func NewMicrosoftGraphUserActivityWithDefaults() *MicrosoftGraphUserActivity`

NewMicrosoftGraphUserActivityWithDefaults instantiates a new MicrosoftGraphUserActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUserActivity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserActivity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUserActivity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUserActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivationUrl

`func (o *MicrosoftGraphUserActivity) GetActivationUrl() string`

GetActivationUrl returns the ActivationUrl field if non-nil, zero value otherwise.

### GetActivationUrlOk

`func (o *MicrosoftGraphUserActivity) GetActivationUrlOk() (*string, bool)`

GetActivationUrlOk returns a tuple with the ActivationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationUrl

`func (o *MicrosoftGraphUserActivity) SetActivationUrl(v string)`

SetActivationUrl sets ActivationUrl field to given value.

### HasActivationUrl

`func (o *MicrosoftGraphUserActivity) HasActivationUrl() bool`

HasActivationUrl returns a boolean if a field has been set.

### GetActivitySourceHost

`func (o *MicrosoftGraphUserActivity) GetActivitySourceHost() string`

GetActivitySourceHost returns the ActivitySourceHost field if non-nil, zero value otherwise.

### GetActivitySourceHostOk

`func (o *MicrosoftGraphUserActivity) GetActivitySourceHostOk() (*string, bool)`

GetActivitySourceHostOk returns a tuple with the ActivitySourceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitySourceHost

`func (o *MicrosoftGraphUserActivity) SetActivitySourceHost(v string)`

SetActivitySourceHost sets ActivitySourceHost field to given value.

### HasActivitySourceHost

`func (o *MicrosoftGraphUserActivity) HasActivitySourceHost() bool`

HasActivitySourceHost returns a boolean if a field has been set.

### GetAppActivityId

`func (o *MicrosoftGraphUserActivity) GetAppActivityId() string`

GetAppActivityId returns the AppActivityId field if non-nil, zero value otherwise.

### GetAppActivityIdOk

`func (o *MicrosoftGraphUserActivity) GetAppActivityIdOk() (*string, bool)`

GetAppActivityIdOk returns a tuple with the AppActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppActivityId

`func (o *MicrosoftGraphUserActivity) SetAppActivityId(v string)`

SetAppActivityId sets AppActivityId field to given value.

### HasAppActivityId

`func (o *MicrosoftGraphUserActivity) HasAppActivityId() bool`

HasAppActivityId returns a boolean if a field has been set.

### GetAppDisplayName

`func (o *MicrosoftGraphUserActivity) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphUserActivity) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphUserActivity) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *MicrosoftGraphUserActivity) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *MicrosoftGraphUserActivity) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *MicrosoftGraphUserActivity) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetContentInfo

`func (o *MicrosoftGraphUserActivity) GetContentInfo() AnyOfobject`

GetContentInfo returns the ContentInfo field if non-nil, zero value otherwise.

### GetContentInfoOk

`func (o *MicrosoftGraphUserActivity) GetContentInfoOk() (*AnyOfobject, bool)`

GetContentInfoOk returns a tuple with the ContentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentInfo

`func (o *MicrosoftGraphUserActivity) SetContentInfo(v AnyOfobject)`

SetContentInfo sets ContentInfo field to given value.

### HasContentInfo

`func (o *MicrosoftGraphUserActivity) HasContentInfo() bool`

HasContentInfo returns a boolean if a field has been set.

### SetContentInfoNil

`func (o *MicrosoftGraphUserActivity) SetContentInfoNil(b bool)`

 SetContentInfoNil sets the value for ContentInfo to be an explicit nil

### UnsetContentInfo
`func (o *MicrosoftGraphUserActivity) UnsetContentInfo()`

UnsetContentInfo ensures that no value is present for ContentInfo, not even an explicit nil
### GetContentUrl

`func (o *MicrosoftGraphUserActivity) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *MicrosoftGraphUserActivity) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *MicrosoftGraphUserActivity) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *MicrosoftGraphUserActivity) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrlNil

`func (o *MicrosoftGraphUserActivity) SetContentUrlNil(b bool)`

 SetContentUrlNil sets the value for ContentUrl to be an explicit nil

### UnsetContentUrl
`func (o *MicrosoftGraphUserActivity) UnsetContentUrl()`

UnsetContentUrl ensures that no value is present for ContentUrl, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphUserActivity) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphUserActivity) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphUserActivity) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphUserActivity) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphUserActivity) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphUserActivity) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetExpirationDateTime

`func (o *MicrosoftGraphUserActivity) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphUserActivity) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphUserActivity) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphUserActivity) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *MicrosoftGraphUserActivity) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *MicrosoftGraphUserActivity) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetFallbackUrl

`func (o *MicrosoftGraphUserActivity) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *MicrosoftGraphUserActivity) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *MicrosoftGraphUserActivity) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *MicrosoftGraphUserActivity) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### SetFallbackUrlNil

`func (o *MicrosoftGraphUserActivity) SetFallbackUrlNil(b bool)`

 SetFallbackUrlNil sets the value for FallbackUrl to be an explicit nil

### UnsetFallbackUrl
`func (o *MicrosoftGraphUserActivity) UnsetFallbackUrl()`

UnsetFallbackUrl ensures that no value is present for FallbackUrl, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphUserActivity) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphUserActivity) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphUserActivity) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphUserActivity) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphUserActivity) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphUserActivity) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphUserActivity) GetStatus() AnyOfmicrosoftGraphStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphUserActivity) GetStatusOk() (*AnyOfmicrosoftGraphStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphUserActivity) SetStatus(v AnyOfmicrosoftGraphStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphUserActivity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphUserActivity) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphUserActivity) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserTimezone

`func (o *MicrosoftGraphUserActivity) GetUserTimezone() string`

GetUserTimezone returns the UserTimezone field if non-nil, zero value otherwise.

### GetUserTimezoneOk

`func (o *MicrosoftGraphUserActivity) GetUserTimezoneOk() (*string, bool)`

GetUserTimezoneOk returns a tuple with the UserTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTimezone

`func (o *MicrosoftGraphUserActivity) SetUserTimezone(v string)`

SetUserTimezone sets UserTimezone field to given value.

### HasUserTimezone

`func (o *MicrosoftGraphUserActivity) HasUserTimezone() bool`

HasUserTimezone returns a boolean if a field has been set.

### SetUserTimezoneNil

`func (o *MicrosoftGraphUserActivity) SetUserTimezoneNil(b bool)`

 SetUserTimezoneNil sets the value for UserTimezone to be an explicit nil

### UnsetUserTimezone
`func (o *MicrosoftGraphUserActivity) UnsetUserTimezone()`

UnsetUserTimezone ensures that no value is present for UserTimezone, not even an explicit nil
### GetVisualElements

`func (o *MicrosoftGraphUserActivity) GetVisualElements() MicrosoftGraphVisualInfo`

GetVisualElements returns the VisualElements field if non-nil, zero value otherwise.

### GetVisualElementsOk

`func (o *MicrosoftGraphUserActivity) GetVisualElementsOk() (*MicrosoftGraphVisualInfo, bool)`

GetVisualElementsOk returns a tuple with the VisualElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualElements

`func (o *MicrosoftGraphUserActivity) SetVisualElements(v MicrosoftGraphVisualInfo)`

SetVisualElements sets VisualElements field to given value.

### HasVisualElements

`func (o *MicrosoftGraphUserActivity) HasVisualElements() bool`

HasVisualElements returns a boolean if a field has been set.

### GetHistoryItems

`func (o *MicrosoftGraphUserActivity) GetHistoryItems() []MicrosoftGraphActivityHistoryItem`

GetHistoryItems returns the HistoryItems field if non-nil, zero value otherwise.

### GetHistoryItemsOk

`func (o *MicrosoftGraphUserActivity) GetHistoryItemsOk() (*[]MicrosoftGraphActivityHistoryItem, bool)`

GetHistoryItemsOk returns a tuple with the HistoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryItems

`func (o *MicrosoftGraphUserActivity) SetHistoryItems(v []MicrosoftGraphActivityHistoryItem)`

SetHistoryItems sets HistoryItems field to given value.

### HasHistoryItems

`func (o *MicrosoftGraphUserActivity) HasHistoryItems() bool`

HasHistoryItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


