/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// UserActivity struct for UserActivity
type UserActivity struct {
	// Required. URL used to launch the activity in the best native experience represented by the appId. Might launch a web-based app if no native app exists.
	ActivationUrl *string `json:"activationUrl,omitempty"`
	// Required. URL for the domain representing the cross-platform identity mapping for the app. Mapping is stored either as a JSON file hosted on the domain or configurable via Windows Dev Center. The JSON file is named cross-platform-app-identifiers and is hosted at root of your HTTPS domain, either at the top level domain or include a sub domain. For example: https://contoso.com or https://myapp.contoso.com but NOT https://myapp.contoso.com/somepath. You must have a unique file and domain (or sub domain) per cross-platform app identity. For example, a separate file and domain is needed for Word vs. PowerPoint.
	ActivitySourceHost *string `json:"activitySourceHost,omitempty"`
	// Required. The unique activity ID in the context of the app - supplied by caller and immutable thereafter.
	AppActivityId *string `json:"appActivityId,omitempty"`
	// Optional. Short text description of the app used to generate the activity for use in cases when the app is not installed on the user’s local device.
	AppDisplayName NullableString `json:"appDisplayName,omitempty"`
	// Optional. A custom piece of data - JSON-LD extensible description of content according to schema.org syntax.
	ContentInfo AnyOfobject `json:"contentInfo,omitempty"`
	// Optional. Used in the event the content can be rendered outside of a native or web-based app experience (for example, a pointer to an item in an RSS feed).
	ContentUrl NullableString `json:"contentUrl,omitempty"`
	// Set by the server. DateTime in UTC when the object was created on the server.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Set by the server. DateTime in UTC when the object expired on the server.
	ExpirationDateTime NullableTime `json:"expirationDateTime,omitempty"`
	// Optional. URL used to launch the activity in a web-based app, if available.
	FallbackUrl NullableString `json:"fallbackUrl,omitempty"`
	// Set by the server. DateTime in UTC when the object was modified on the server.
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// Set by the server. A status code used to identify valid objects. Values: active, updated, deleted, ignored.
	Status AnyOfmicrosoftGraphStatus `json:"status,omitempty"`
	// Optional. The timezone in which the user's device used to generate the activity was located at activity creation time; values supplied as Olson IDs in order to support cross-platform representation.
	UserTimezone NullableString `json:"userTimezone,omitempty"`
	VisualElements *MicrosoftGraphVisualInfo `json:"visualElements,omitempty"`
	// Optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
	HistoryItems *[]MicrosoftGraphActivityHistoryItem `json:"historyItems,omitempty"`
}

// NewUserActivity instantiates a new UserActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActivity() *UserActivity {
	this := UserActivity{}
	return &this
}

// NewUserActivityWithDefaults instantiates a new UserActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActivityWithDefaults() *UserActivity {
	this := UserActivity{}
	return &this
}

// GetActivationUrl returns the ActivationUrl field value if set, zero value otherwise.
func (o *UserActivity) GetActivationUrl() string {
	if o == nil || o.ActivationUrl == nil {
		var ret string
		return ret
	}
	return *o.ActivationUrl
}

// GetActivationUrlOk returns a tuple with the ActivationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetActivationUrlOk() (*string, bool) {
	if o == nil || o.ActivationUrl == nil {
		return nil, false
	}
	return o.ActivationUrl, true
}

// HasActivationUrl returns a boolean if a field has been set.
func (o *UserActivity) HasActivationUrl() bool {
	if o != nil && o.ActivationUrl != nil {
		return true
	}

	return false
}

// SetActivationUrl gets a reference to the given string and assigns it to the ActivationUrl field.
func (o *UserActivity) SetActivationUrl(v string) {
	o.ActivationUrl = &v
}

// GetActivitySourceHost returns the ActivitySourceHost field value if set, zero value otherwise.
func (o *UserActivity) GetActivitySourceHost() string {
	if o == nil || o.ActivitySourceHost == nil {
		var ret string
		return ret
	}
	return *o.ActivitySourceHost
}

// GetActivitySourceHostOk returns a tuple with the ActivitySourceHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetActivitySourceHostOk() (*string, bool) {
	if o == nil || o.ActivitySourceHost == nil {
		return nil, false
	}
	return o.ActivitySourceHost, true
}

// HasActivitySourceHost returns a boolean if a field has been set.
func (o *UserActivity) HasActivitySourceHost() bool {
	if o != nil && o.ActivitySourceHost != nil {
		return true
	}

	return false
}

// SetActivitySourceHost gets a reference to the given string and assigns it to the ActivitySourceHost field.
func (o *UserActivity) SetActivitySourceHost(v string) {
	o.ActivitySourceHost = &v
}

// GetAppActivityId returns the AppActivityId field value if set, zero value otherwise.
func (o *UserActivity) GetAppActivityId() string {
	if o == nil || o.AppActivityId == nil {
		var ret string
		return ret
	}
	return *o.AppActivityId
}

// GetAppActivityIdOk returns a tuple with the AppActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetAppActivityIdOk() (*string, bool) {
	if o == nil || o.AppActivityId == nil {
		return nil, false
	}
	return o.AppActivityId, true
}

// HasAppActivityId returns a boolean if a field has been set.
func (o *UserActivity) HasAppActivityId() bool {
	if o != nil && o.AppActivityId != nil {
		return true
	}

	return false
}

// SetAppActivityId gets a reference to the given string and assigns it to the AppActivityId field.
func (o *UserActivity) SetAppActivityId(v string) {
	o.AppActivityId = &v
}

// GetAppDisplayName returns the AppDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetAppDisplayName() string {
	if o == nil || o.AppDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppDisplayName.Get()
}

// GetAppDisplayNameOk returns a tuple with the AppDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetAppDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppDisplayName.Get(), o.AppDisplayName.IsSet()
}

// HasAppDisplayName returns a boolean if a field has been set.
func (o *UserActivity) HasAppDisplayName() bool {
	if o != nil && o.AppDisplayName.IsSet() {
		return true
	}

	return false
}

// SetAppDisplayName gets a reference to the given NullableString and assigns it to the AppDisplayName field.
func (o *UserActivity) SetAppDisplayName(v string) {
	o.AppDisplayName.Set(&v)
}
// SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil
func (o *UserActivity) SetAppDisplayNameNil() {
	o.AppDisplayName.Set(nil)
}

// UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
func (o *UserActivity) UnsetAppDisplayName() {
	o.AppDisplayName.Unset()
}

// GetContentInfo returns the ContentInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetContentInfo() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.ContentInfo
}

// GetContentInfoOk returns a tuple with the ContentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetContentInfoOk() (*AnyOfobject, bool) {
	if o == nil || o.ContentInfo == nil {
		return nil, false
	}
	return &o.ContentInfo, true
}

// HasContentInfo returns a boolean if a field has been set.
func (o *UserActivity) HasContentInfo() bool {
	if o != nil && o.ContentInfo != nil {
		return true
	}

	return false
}

// SetContentInfo gets a reference to the given AnyOfobject and assigns it to the ContentInfo field.
func (o *UserActivity) SetContentInfo(v AnyOfobject) {
	o.ContentInfo = v
}

// GetContentUrl returns the ContentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetContentUrl() string {
	if o == nil || o.ContentUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentUrl.Get()
}

// GetContentUrlOk returns a tuple with the ContentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetContentUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentUrl.Get(), o.ContentUrl.IsSet()
}

// HasContentUrl returns a boolean if a field has been set.
func (o *UserActivity) HasContentUrl() bool {
	if o != nil && o.ContentUrl.IsSet() {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given NullableString and assigns it to the ContentUrl field.
func (o *UserActivity) SetContentUrl(v string) {
	o.ContentUrl.Set(&v)
}
// SetContentUrlNil sets the value for ContentUrl to be an explicit nil
func (o *UserActivity) SetContentUrlNil() {
	o.ContentUrl.Set(nil)
}

// UnsetContentUrl ensures that no value is present for ContentUrl, not even an explicit nil
func (o *UserActivity) UnsetContentUrl() {
	o.ContentUrl.Unset()
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *UserActivity) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *UserActivity) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *UserActivity) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *UserActivity) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime.Get()
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDateTime.Get(), o.ExpirationDateTime.IsSet()
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *UserActivity) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime.IsSet() {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given NullableTime and assigns it to the ExpirationDateTime field.
func (o *UserActivity) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime.Set(&v)
}
// SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil
func (o *UserActivity) SetExpirationDateTimeNil() {
	o.ExpirationDateTime.Set(nil)
}

// UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
func (o *UserActivity) UnsetExpirationDateTime() {
	o.ExpirationDateTime.Unset()
}

// GetFallbackUrl returns the FallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetFallbackUrl() string {
	if o == nil || o.FallbackUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.FallbackUrl.Get()
}

// GetFallbackUrlOk returns a tuple with the FallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetFallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FallbackUrl.Get(), o.FallbackUrl.IsSet()
}

// HasFallbackUrl returns a boolean if a field has been set.
func (o *UserActivity) HasFallbackUrl() bool {
	if o != nil && o.FallbackUrl.IsSet() {
		return true
	}

	return false
}

// SetFallbackUrl gets a reference to the given NullableString and assigns it to the FallbackUrl field.
func (o *UserActivity) SetFallbackUrl(v string) {
	o.FallbackUrl.Set(&v)
}
// SetFallbackUrlNil sets the value for FallbackUrl to be an explicit nil
func (o *UserActivity) SetFallbackUrlNil() {
	o.FallbackUrl.Set(nil)
}

// UnsetFallbackUrl ensures that no value is present for FallbackUrl, not even an explicit nil
func (o *UserActivity) UnsetFallbackUrl() {
	o.FallbackUrl.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *UserActivity) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *UserActivity) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *UserActivity) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *UserActivity) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetStatus() AnyOfmicrosoftGraphStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetStatusOk() (*AnyOfmicrosoftGraphStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserActivity) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphStatus and assigns it to the Status field.
func (o *UserActivity) SetStatus(v AnyOfmicrosoftGraphStatus) {
	o.Status = v
}

// GetUserTimezone returns the UserTimezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserActivity) GetUserTimezone() string {
	if o == nil || o.UserTimezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserTimezone.Get()
}

// GetUserTimezoneOk returns a tuple with the UserTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserActivity) GetUserTimezoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserTimezone.Get(), o.UserTimezone.IsSet()
}

// HasUserTimezone returns a boolean if a field has been set.
func (o *UserActivity) HasUserTimezone() bool {
	if o != nil && o.UserTimezone.IsSet() {
		return true
	}

	return false
}

// SetUserTimezone gets a reference to the given NullableString and assigns it to the UserTimezone field.
func (o *UserActivity) SetUserTimezone(v string) {
	o.UserTimezone.Set(&v)
}
// SetUserTimezoneNil sets the value for UserTimezone to be an explicit nil
func (o *UserActivity) SetUserTimezoneNil() {
	o.UserTimezone.Set(nil)
}

// UnsetUserTimezone ensures that no value is present for UserTimezone, not even an explicit nil
func (o *UserActivity) UnsetUserTimezone() {
	o.UserTimezone.Unset()
}

// GetVisualElements returns the VisualElements field value if set, zero value otherwise.
func (o *UserActivity) GetVisualElements() MicrosoftGraphVisualInfo {
	if o == nil || o.VisualElements == nil {
		var ret MicrosoftGraphVisualInfo
		return ret
	}
	return *o.VisualElements
}

// GetVisualElementsOk returns a tuple with the VisualElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetVisualElementsOk() (*MicrosoftGraphVisualInfo, bool) {
	if o == nil || o.VisualElements == nil {
		return nil, false
	}
	return o.VisualElements, true
}

// HasVisualElements returns a boolean if a field has been set.
func (o *UserActivity) HasVisualElements() bool {
	if o != nil && o.VisualElements != nil {
		return true
	}

	return false
}

// SetVisualElements gets a reference to the given MicrosoftGraphVisualInfo and assigns it to the VisualElements field.
func (o *UserActivity) SetVisualElements(v MicrosoftGraphVisualInfo) {
	o.VisualElements = &v
}

// GetHistoryItems returns the HistoryItems field value if set, zero value otherwise.
func (o *UserActivity) GetHistoryItems() []MicrosoftGraphActivityHistoryItem {
	if o == nil || o.HistoryItems == nil {
		var ret []MicrosoftGraphActivityHistoryItem
		return ret
	}
	return *o.HistoryItems
}

// GetHistoryItemsOk returns a tuple with the HistoryItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetHistoryItemsOk() (*[]MicrosoftGraphActivityHistoryItem, bool) {
	if o == nil || o.HistoryItems == nil {
		return nil, false
	}
	return o.HistoryItems, true
}

// HasHistoryItems returns a boolean if a field has been set.
func (o *UserActivity) HasHistoryItems() bool {
	if o != nil && o.HistoryItems != nil {
		return true
	}

	return false
}

// SetHistoryItems gets a reference to the given []MicrosoftGraphActivityHistoryItem and assigns it to the HistoryItems field.
func (o *UserActivity) SetHistoryItems(v []MicrosoftGraphActivityHistoryItem) {
	o.HistoryItems = &v
}

func (o UserActivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivationUrl != nil {
		toSerialize["activationUrl"] = o.ActivationUrl
	}
	if o.ActivitySourceHost != nil {
		toSerialize["activitySourceHost"] = o.ActivitySourceHost
	}
	if o.AppActivityId != nil {
		toSerialize["appActivityId"] = o.AppActivityId
	}
	if o.AppDisplayName.IsSet() {
		toSerialize["appDisplayName"] = o.AppDisplayName.Get()
	}
	if o.ContentInfo != nil {
		toSerialize["contentInfo"] = o.ContentInfo
	}
	if o.ContentUrl.IsSet() {
		toSerialize["contentUrl"] = o.ContentUrl.Get()
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.ExpirationDateTime.IsSet() {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime.Get()
	}
	if o.FallbackUrl.IsSet() {
		toSerialize["fallbackUrl"] = o.FallbackUrl.Get()
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserTimezone.IsSet() {
		toSerialize["userTimezone"] = o.UserTimezone.Get()
	}
	if o.VisualElements != nil {
		toSerialize["visualElements"] = o.VisualElements
	}
	if o.HistoryItems != nil {
		toSerialize["historyItems"] = o.HistoryItems
	}
	return json.Marshal(toSerialize)
}

type NullableUserActivity struct {
	value *UserActivity
	isSet bool
}

func (v NullableUserActivity) Get() *UserActivity {
	return v.value
}

func (v *NullableUserActivity) Set(val *UserActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActivity(val *UserActivity) *NullableUserActivity {
	return &NullableUserActivity{value: val, isSet: true}
}

func (v NullableUserActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


