# UserActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewUserActivity

`func NewUserActivity() *UserActivity`

NewUserActivity instantiates a new UserActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityWithDefaults

`func NewUserActivityWithDefaults() *UserActivity`

NewUserActivityWithDefaults instantiates a new UserActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationUrl

`func (o *UserActivity) GetActivationUrl() string`

GetActivationUrl returns the ActivationUrl field if non-nil, zero value otherwise.

### GetActivationUrlOk

`func (o *UserActivity) GetActivationUrlOk() (*string, bool)`

GetActivationUrlOk returns a tuple with the ActivationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationUrl

`func (o *UserActivity) SetActivationUrl(v string)`

SetActivationUrl sets ActivationUrl field to given value.

### HasActivationUrl

`func (o *UserActivity) HasActivationUrl() bool`

HasActivationUrl returns a boolean if a field has been set.

### GetActivitySourceHost

`func (o *UserActivity) GetActivitySourceHost() string`

GetActivitySourceHost returns the ActivitySourceHost field if non-nil, zero value otherwise.

### GetActivitySourceHostOk

`func (o *UserActivity) GetActivitySourceHostOk() (*string, bool)`

GetActivitySourceHostOk returns a tuple with the ActivitySourceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitySourceHost

`func (o *UserActivity) SetActivitySourceHost(v string)`

SetActivitySourceHost sets ActivitySourceHost field to given value.

### HasActivitySourceHost

`func (o *UserActivity) HasActivitySourceHost() bool`

HasActivitySourceHost returns a boolean if a field has been set.

### GetAppActivityId

`func (o *UserActivity) GetAppActivityId() string`

GetAppActivityId returns the AppActivityId field if non-nil, zero value otherwise.

### GetAppActivityIdOk

`func (o *UserActivity) GetAppActivityIdOk() (*string, bool)`

GetAppActivityIdOk returns a tuple with the AppActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppActivityId

`func (o *UserActivity) SetAppActivityId(v string)`

SetAppActivityId sets AppActivityId field to given value.

### HasAppActivityId

`func (o *UserActivity) HasAppActivityId() bool`

HasAppActivityId returns a boolean if a field has been set.

### GetAppDisplayName

`func (o *UserActivity) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *UserActivity) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *UserActivity) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *UserActivity) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *UserActivity) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *UserActivity) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetContentInfo

`func (o *UserActivity) GetContentInfo() AnyOfobject`

GetContentInfo returns the ContentInfo field if non-nil, zero value otherwise.

### GetContentInfoOk

`func (o *UserActivity) GetContentInfoOk() (*AnyOfobject, bool)`

GetContentInfoOk returns a tuple with the ContentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentInfo

`func (o *UserActivity) SetContentInfo(v AnyOfobject)`

SetContentInfo sets ContentInfo field to given value.

### HasContentInfo

`func (o *UserActivity) HasContentInfo() bool`

HasContentInfo returns a boolean if a field has been set.

### SetContentInfoNil

`func (o *UserActivity) SetContentInfoNil(b bool)`

 SetContentInfoNil sets the value for ContentInfo to be an explicit nil

### UnsetContentInfo
`func (o *UserActivity) UnsetContentInfo()`

UnsetContentInfo ensures that no value is present for ContentInfo, not even an explicit nil
### GetContentUrl

`func (o *UserActivity) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *UserActivity) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *UserActivity) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *UserActivity) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrlNil

`func (o *UserActivity) SetContentUrlNil(b bool)`

 SetContentUrlNil sets the value for ContentUrl to be an explicit nil

### UnsetContentUrl
`func (o *UserActivity) UnsetContentUrl()`

UnsetContentUrl ensures that no value is present for ContentUrl, not even an explicit nil
### GetCreatedDateTime

`func (o *UserActivity) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *UserActivity) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *UserActivity) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *UserActivity) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *UserActivity) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *UserActivity) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetExpirationDateTime

`func (o *UserActivity) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *UserActivity) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *UserActivity) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *UserActivity) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *UserActivity) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *UserActivity) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetFallbackUrl

`func (o *UserActivity) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *UserActivity) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *UserActivity) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *UserActivity) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### SetFallbackUrlNil

`func (o *UserActivity) SetFallbackUrlNil(b bool)`

 SetFallbackUrlNil sets the value for FallbackUrl to be an explicit nil

### UnsetFallbackUrl
`func (o *UserActivity) UnsetFallbackUrl()`

UnsetFallbackUrl ensures that no value is present for FallbackUrl, not even an explicit nil
### GetLastModifiedDateTime

`func (o *UserActivity) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *UserActivity) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *UserActivity) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *UserActivity) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *UserActivity) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *UserActivity) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetStatus

`func (o *UserActivity) GetStatus() AnyOfmicrosoftGraphStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserActivity) GetStatusOk() (*AnyOfmicrosoftGraphStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserActivity) SetStatus(v AnyOfmicrosoftGraphStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserActivity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UserActivity) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UserActivity) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserTimezone

`func (o *UserActivity) GetUserTimezone() string`

GetUserTimezone returns the UserTimezone field if non-nil, zero value otherwise.

### GetUserTimezoneOk

`func (o *UserActivity) GetUserTimezoneOk() (*string, bool)`

GetUserTimezoneOk returns a tuple with the UserTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTimezone

`func (o *UserActivity) SetUserTimezone(v string)`

SetUserTimezone sets UserTimezone field to given value.

### HasUserTimezone

`func (o *UserActivity) HasUserTimezone() bool`

HasUserTimezone returns a boolean if a field has been set.

### SetUserTimezoneNil

`func (o *UserActivity) SetUserTimezoneNil(b bool)`

 SetUserTimezoneNil sets the value for UserTimezone to be an explicit nil

### UnsetUserTimezone
`func (o *UserActivity) UnsetUserTimezone()`

UnsetUserTimezone ensures that no value is present for UserTimezone, not even an explicit nil
### GetVisualElements

`func (o *UserActivity) GetVisualElements() MicrosoftGraphVisualInfo`

GetVisualElements returns the VisualElements field if non-nil, zero value otherwise.

### GetVisualElementsOk

`func (o *UserActivity) GetVisualElementsOk() (*MicrosoftGraphVisualInfo, bool)`

GetVisualElementsOk returns a tuple with the VisualElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualElements

`func (o *UserActivity) SetVisualElements(v MicrosoftGraphVisualInfo)`

SetVisualElements sets VisualElements field to given value.

### HasVisualElements

`func (o *UserActivity) HasVisualElements() bool`

HasVisualElements returns a boolean if a field has been set.

### GetHistoryItems

`func (o *UserActivity) GetHistoryItems() []MicrosoftGraphActivityHistoryItem`

GetHistoryItems returns the HistoryItems field if non-nil, zero value otherwise.

### GetHistoryItemsOk

`func (o *UserActivity) GetHistoryItemsOk() (*[]MicrosoftGraphActivityHistoryItem, bool)`

GetHistoryItemsOk returns a tuple with the HistoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryItems

`func (o *UserActivity) SetHistoryItems(v []MicrosoftGraphActivityHistoryItem)`

SetHistoryItems sets HistoryItems field to given value.

### HasHistoryItems

`func (o *UserActivity) HasHistoryItems() bool`

HasHistoryItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


