/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphSubscription struct for MicrosoftGraphSubscription
type MicrosoftGraphSubscription struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identifier of the application used to create the subscription. Read-only.
	ApplicationId NullableString `json:"applicationId,omitempty"`
	// Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list.Note: Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType.
	ChangeType *string `json:"changeType,omitempty"`
	// Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification.
	ClientState NullableString `json:"clientState,omitempty"`
	// Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only.
	CreatorId NullableString `json:"creatorId,omitempty"`
	// A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional. Required when includeResourceData is true.
	EncryptionCertificate NullableString `json:"encryptionCertificate,omitempty"`
	// A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Optional.
	EncryptionCertificateId NullableString `json:"encryptionCertificateId,omitempty"`
	// Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to.  See the table below for maximum supported subscription length of time.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// When set to true, change notifications include resource data (such as content of a chat message). Optional.
	IncludeResourceData NullableBool `json:"includeResourceData,omitempty"`
	// Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2.
	LatestSupportedTlsVersion NullableString `json:"latestSupportedTlsVersion,omitempty"`
	// The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol. Optional. Read more about how Outlook resources use lifecycle notifications.
	LifecycleNotificationUrl NullableString `json:"lifecycleNotificationUrl,omitempty"`
	// OData Query Options for specifying value for the targeting resource. Clients receive notifications when resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property eg  when the print job is completed, when a print job resource isFetchable property value becomes true etc.
	NotificationQueryOptions NullableString `json:"notificationQueryOptions,omitempty"`
	// Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol.
	NotificationUrl *string `json:"notificationUrl,omitempty"`
	NotificationUrlAppId NullableString `json:"notificationUrlAppId,omitempty"`
	// Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource.
	Resource *string `json:"resource,omitempty"`
}

// NewMicrosoftGraphSubscription instantiates a new MicrosoftGraphSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSubscription() *MicrosoftGraphSubscription {
	this := MicrosoftGraphSubscription{}
	return &this
}

// NewMicrosoftGraphSubscriptionWithDefaults instantiates a new MicrosoftGraphSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSubscriptionWithDefaults() *MicrosoftGraphSubscription {
	this := MicrosoftGraphSubscription{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphSubscription) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSubscription) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSubscription) SetId(v string) {
	o.Id = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetApplicationId() string {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasApplicationId() bool {
	if o != nil && o.ApplicationId.IsSet() {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given NullableString and assigns it to the ApplicationId field.
func (o *MicrosoftGraphSubscription) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}
// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil
func (o *MicrosoftGraphSubscription) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetChangeType returns the ChangeType field value if set, zero value otherwise.
func (o *MicrosoftGraphSubscription) GetChangeType() string {
	if o == nil || o.ChangeType == nil {
		var ret string
		return ret
	}
	return *o.ChangeType
}

// GetChangeTypeOk returns a tuple with the ChangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSubscription) GetChangeTypeOk() (*string, bool) {
	if o == nil || o.ChangeType == nil {
		return nil, false
	}
	return o.ChangeType, true
}

// HasChangeType returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasChangeType() bool {
	if o != nil && o.ChangeType != nil {
		return true
	}

	return false
}

// SetChangeType gets a reference to the given string and assigns it to the ChangeType field.
func (o *MicrosoftGraphSubscription) SetChangeType(v string) {
	o.ChangeType = &v
}

// GetClientState returns the ClientState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetClientState() string {
	if o == nil || o.ClientState.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientState.Get()
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetClientStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientState.Get(), o.ClientState.IsSet()
}

// HasClientState returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasClientState() bool {
	if o != nil && o.ClientState.IsSet() {
		return true
	}

	return false
}

// SetClientState gets a reference to the given NullableString and assigns it to the ClientState field.
func (o *MicrosoftGraphSubscription) SetClientState(v string) {
	o.ClientState.Set(&v)
}
// SetClientStateNil sets the value for ClientState to be an explicit nil
func (o *MicrosoftGraphSubscription) SetClientStateNil() {
	o.ClientState.Set(nil)
}

// UnsetClientState ensures that no value is present for ClientState, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetClientState() {
	o.ClientState.Unset()
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetCreatorId() string {
	if o == nil || o.CreatorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetCreatorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableString and assigns it to the CreatorId field.
func (o *MicrosoftGraphSubscription) SetCreatorId(v string) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *MicrosoftGraphSubscription) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetEncryptionCertificate returns the EncryptionCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetEncryptionCertificate() string {
	if o == nil || o.EncryptionCertificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EncryptionCertificate.Get()
}

// GetEncryptionCertificateOk returns a tuple with the EncryptionCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetEncryptionCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionCertificate.Get(), o.EncryptionCertificate.IsSet()
}

// HasEncryptionCertificate returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasEncryptionCertificate() bool {
	if o != nil && o.EncryptionCertificate.IsSet() {
		return true
	}

	return false
}

// SetEncryptionCertificate gets a reference to the given NullableString and assigns it to the EncryptionCertificate field.
func (o *MicrosoftGraphSubscription) SetEncryptionCertificate(v string) {
	o.EncryptionCertificate.Set(&v)
}
// SetEncryptionCertificateNil sets the value for EncryptionCertificate to be an explicit nil
func (o *MicrosoftGraphSubscription) SetEncryptionCertificateNil() {
	o.EncryptionCertificate.Set(nil)
}

// UnsetEncryptionCertificate ensures that no value is present for EncryptionCertificate, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetEncryptionCertificate() {
	o.EncryptionCertificate.Unset()
}

// GetEncryptionCertificateId returns the EncryptionCertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetEncryptionCertificateId() string {
	if o == nil || o.EncryptionCertificateId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EncryptionCertificateId.Get()
}

// GetEncryptionCertificateIdOk returns a tuple with the EncryptionCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetEncryptionCertificateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionCertificateId.Get(), o.EncryptionCertificateId.IsSet()
}

// HasEncryptionCertificateId returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasEncryptionCertificateId() bool {
	if o != nil && o.EncryptionCertificateId.IsSet() {
		return true
	}

	return false
}

// SetEncryptionCertificateId gets a reference to the given NullableString and assigns it to the EncryptionCertificateId field.
func (o *MicrosoftGraphSubscription) SetEncryptionCertificateId(v string) {
	o.EncryptionCertificateId.Set(&v)
}
// SetEncryptionCertificateIdNil sets the value for EncryptionCertificateId to be an explicit nil
func (o *MicrosoftGraphSubscription) SetEncryptionCertificateIdNil() {
	o.EncryptionCertificateId.Set(nil)
}

// UnsetEncryptionCertificateId ensures that no value is present for EncryptionCertificateId, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetEncryptionCertificateId() {
	o.EncryptionCertificateId.Unset()
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphSubscription) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSubscription) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDateTime == nil {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime != nil {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *MicrosoftGraphSubscription) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// GetIncludeResourceData returns the IncludeResourceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetIncludeResourceData() bool {
	if o == nil || o.IncludeResourceData.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeResourceData.Get()
}

// GetIncludeResourceDataOk returns a tuple with the IncludeResourceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetIncludeResourceDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludeResourceData.Get(), o.IncludeResourceData.IsSet()
}

// HasIncludeResourceData returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasIncludeResourceData() bool {
	if o != nil && o.IncludeResourceData.IsSet() {
		return true
	}

	return false
}

// SetIncludeResourceData gets a reference to the given NullableBool and assigns it to the IncludeResourceData field.
func (o *MicrosoftGraphSubscription) SetIncludeResourceData(v bool) {
	o.IncludeResourceData.Set(&v)
}
// SetIncludeResourceDataNil sets the value for IncludeResourceData to be an explicit nil
func (o *MicrosoftGraphSubscription) SetIncludeResourceDataNil() {
	o.IncludeResourceData.Set(nil)
}

// UnsetIncludeResourceData ensures that no value is present for IncludeResourceData, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetIncludeResourceData() {
	o.IncludeResourceData.Unset()
}

// GetLatestSupportedTlsVersion returns the LatestSupportedTlsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetLatestSupportedTlsVersion() string {
	if o == nil || o.LatestSupportedTlsVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.LatestSupportedTlsVersion.Get()
}

// GetLatestSupportedTlsVersionOk returns a tuple with the LatestSupportedTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetLatestSupportedTlsVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LatestSupportedTlsVersion.Get(), o.LatestSupportedTlsVersion.IsSet()
}

// HasLatestSupportedTlsVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasLatestSupportedTlsVersion() bool {
	if o != nil && o.LatestSupportedTlsVersion.IsSet() {
		return true
	}

	return false
}

// SetLatestSupportedTlsVersion gets a reference to the given NullableString and assigns it to the LatestSupportedTlsVersion field.
func (o *MicrosoftGraphSubscription) SetLatestSupportedTlsVersion(v string) {
	o.LatestSupportedTlsVersion.Set(&v)
}
// SetLatestSupportedTlsVersionNil sets the value for LatestSupportedTlsVersion to be an explicit nil
func (o *MicrosoftGraphSubscription) SetLatestSupportedTlsVersionNil() {
	o.LatestSupportedTlsVersion.Set(nil)
}

// UnsetLatestSupportedTlsVersion ensures that no value is present for LatestSupportedTlsVersion, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetLatestSupportedTlsVersion() {
	o.LatestSupportedTlsVersion.Unset()
}

// GetLifecycleNotificationUrl returns the LifecycleNotificationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetLifecycleNotificationUrl() string {
	if o == nil || o.LifecycleNotificationUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LifecycleNotificationUrl.Get()
}

// GetLifecycleNotificationUrlOk returns a tuple with the LifecycleNotificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetLifecycleNotificationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LifecycleNotificationUrl.Get(), o.LifecycleNotificationUrl.IsSet()
}

// HasLifecycleNotificationUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasLifecycleNotificationUrl() bool {
	if o != nil && o.LifecycleNotificationUrl.IsSet() {
		return true
	}

	return false
}

// SetLifecycleNotificationUrl gets a reference to the given NullableString and assigns it to the LifecycleNotificationUrl field.
func (o *MicrosoftGraphSubscription) SetLifecycleNotificationUrl(v string) {
	o.LifecycleNotificationUrl.Set(&v)
}
// SetLifecycleNotificationUrlNil sets the value for LifecycleNotificationUrl to be an explicit nil
func (o *MicrosoftGraphSubscription) SetLifecycleNotificationUrlNil() {
	o.LifecycleNotificationUrl.Set(nil)
}

// UnsetLifecycleNotificationUrl ensures that no value is present for LifecycleNotificationUrl, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetLifecycleNotificationUrl() {
	o.LifecycleNotificationUrl.Unset()
}

// GetNotificationQueryOptions returns the NotificationQueryOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetNotificationQueryOptions() string {
	if o == nil || o.NotificationQueryOptions.Get() == nil {
		var ret string
		return ret
	}
	return *o.NotificationQueryOptions.Get()
}

// GetNotificationQueryOptionsOk returns a tuple with the NotificationQueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetNotificationQueryOptionsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NotificationQueryOptions.Get(), o.NotificationQueryOptions.IsSet()
}

// HasNotificationQueryOptions returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasNotificationQueryOptions() bool {
	if o != nil && o.NotificationQueryOptions.IsSet() {
		return true
	}

	return false
}

// SetNotificationQueryOptions gets a reference to the given NullableString and assigns it to the NotificationQueryOptions field.
func (o *MicrosoftGraphSubscription) SetNotificationQueryOptions(v string) {
	o.NotificationQueryOptions.Set(&v)
}
// SetNotificationQueryOptionsNil sets the value for NotificationQueryOptions to be an explicit nil
func (o *MicrosoftGraphSubscription) SetNotificationQueryOptionsNil() {
	o.NotificationQueryOptions.Set(nil)
}

// UnsetNotificationQueryOptions ensures that no value is present for NotificationQueryOptions, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetNotificationQueryOptions() {
	o.NotificationQueryOptions.Unset()
}

// GetNotificationUrl returns the NotificationUrl field value if set, zero value otherwise.
func (o *MicrosoftGraphSubscription) GetNotificationUrl() string {
	if o == nil || o.NotificationUrl == nil {
		var ret string
		return ret
	}
	return *o.NotificationUrl
}

// GetNotificationUrlOk returns a tuple with the NotificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSubscription) GetNotificationUrlOk() (*string, bool) {
	if o == nil || o.NotificationUrl == nil {
		return nil, false
	}
	return o.NotificationUrl, true
}

// HasNotificationUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasNotificationUrl() bool {
	if o != nil && o.NotificationUrl != nil {
		return true
	}

	return false
}

// SetNotificationUrl gets a reference to the given string and assigns it to the NotificationUrl field.
func (o *MicrosoftGraphSubscription) SetNotificationUrl(v string) {
	o.NotificationUrl = &v
}

// GetNotificationUrlAppId returns the NotificationUrlAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubscription) GetNotificationUrlAppId() string {
	if o == nil || o.NotificationUrlAppId.Get() == nil {
		var ret string
		return ret
	}
	return *o.NotificationUrlAppId.Get()
}

// GetNotificationUrlAppIdOk returns a tuple with the NotificationUrlAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubscription) GetNotificationUrlAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NotificationUrlAppId.Get(), o.NotificationUrlAppId.IsSet()
}

// HasNotificationUrlAppId returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasNotificationUrlAppId() bool {
	if o != nil && o.NotificationUrlAppId.IsSet() {
		return true
	}

	return false
}

// SetNotificationUrlAppId gets a reference to the given NullableString and assigns it to the NotificationUrlAppId field.
func (o *MicrosoftGraphSubscription) SetNotificationUrlAppId(v string) {
	o.NotificationUrlAppId.Set(&v)
}
// SetNotificationUrlAppIdNil sets the value for NotificationUrlAppId to be an explicit nil
func (o *MicrosoftGraphSubscription) SetNotificationUrlAppIdNil() {
	o.NotificationUrlAppId.Set(nil)
}

// UnsetNotificationUrlAppId ensures that no value is present for NotificationUrlAppId, not even an explicit nil
func (o *MicrosoftGraphSubscription) UnsetNotificationUrlAppId() {
	o.NotificationUrlAppId.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *MicrosoftGraphSubscription) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSubscription) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *MicrosoftGraphSubscription) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *MicrosoftGraphSubscription) SetResource(v string) {
	o.Resource = &v
}

func (o MicrosoftGraphSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ApplicationId.IsSet() {
		toSerialize["applicationId"] = o.ApplicationId.Get()
	}
	if o.ChangeType != nil {
		toSerialize["changeType"] = o.ChangeType
	}
	if o.ClientState.IsSet() {
		toSerialize["clientState"] = o.ClientState.Get()
	}
	if o.CreatorId.IsSet() {
		toSerialize["creatorId"] = o.CreatorId.Get()
	}
	if o.EncryptionCertificate.IsSet() {
		toSerialize["encryptionCertificate"] = o.EncryptionCertificate.Get()
	}
	if o.EncryptionCertificateId.IsSet() {
		toSerialize["encryptionCertificateId"] = o.EncryptionCertificateId.Get()
	}
	if o.ExpirationDateTime != nil {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if o.IncludeResourceData.IsSet() {
		toSerialize["includeResourceData"] = o.IncludeResourceData.Get()
	}
	if o.LatestSupportedTlsVersion.IsSet() {
		toSerialize["latestSupportedTlsVersion"] = o.LatestSupportedTlsVersion.Get()
	}
	if o.LifecycleNotificationUrl.IsSet() {
		toSerialize["lifecycleNotificationUrl"] = o.LifecycleNotificationUrl.Get()
	}
	if o.NotificationQueryOptions.IsSet() {
		toSerialize["notificationQueryOptions"] = o.NotificationQueryOptions.Get()
	}
	if o.NotificationUrl != nil {
		toSerialize["notificationUrl"] = o.NotificationUrl
	}
	if o.NotificationUrlAppId.IsSet() {
		toSerialize["notificationUrlAppId"] = o.NotificationUrlAppId.Get()
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSubscription struct {
	value *MicrosoftGraphSubscription
	isSet bool
}

func (v NullableMicrosoftGraphSubscription) Get() *MicrosoftGraphSubscription {
	return v.value
}

func (v *NullableMicrosoftGraphSubscription) Set(val *MicrosoftGraphSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSubscription(val *MicrosoftGraphSubscription) *NullableMicrosoftGraphSubscription {
	return &NullableMicrosoftGraphSubscription{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


