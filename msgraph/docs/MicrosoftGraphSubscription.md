# MicrosoftGraphSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ApplicationId** | Pointer to **NullableString** | Identifier of the application used to create the subscription. Read-only. | [optional] 
**ChangeType** | Pointer to **string** | Required. Indicates the type of change in the subscribed resource that will raise a change notification. The supported values are: created, updated, deleted. Multiple values can be combined using a comma-separated list.Note: Drive root item and list change notifications support only the updated changeType. User and group change notifications support updated and deleted changeType. | [optional] 
**ClientState** | Pointer to **NullableString** | Optional. Specifies the value of the clientState property sent by the service in each change notification. The maximum length is 128 characters. The client can check that the change notification came from the service by comparing the value of the clientState property sent with the subscription with the value of the clientState property received with each change notification. | [optional] 
**CreatorId** | Pointer to **NullableString** | Identifier of the user or service principal that created the subscription. If the app used delegated permissions to create the subscription, this field contains the id of the signed-in user the app called on behalf of. If the app used application permissions, this field contains the id of the service principal corresponding to the app. Read-only. | [optional] 
**EncryptionCertificate** | Pointer to **NullableString** | A base64-encoded representation of a certificate with a public key used to encrypt resource data in change notifications. Optional. Required when includeResourceData is true. | [optional] 
**EncryptionCertificateId** | Pointer to **NullableString** | A custom app-provided identifier to help identify the certificate needed to decrypt resource data. Optional. | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | Required. Specifies the date and time when the webhook subscription expires. The time is in UTC, and can be an amount of time from subscription creation that varies for the resource subscribed to.  See the table below for maximum supported subscription length of time. | [optional] 
**IncludeResourceData** | Pointer to **NullableBool** | When set to true, change notifications include resource data (such as content of a chat message). Optional. | [optional] 
**LatestSupportedTlsVersion** | Pointer to **NullableString** | Specifies the latest version of Transport Layer Security (TLS) that the notification endpoint, specified by notificationUrl, supports. The possible values are: v1_0, v1_1, v1_2, v1_3. For subscribers whose notification endpoint supports a version lower than the currently recommended version (TLS 1.2), specifying this property by a set timeline allows them to temporarily use their deprecated version of TLS before completing their upgrade to TLS 1.2. For these subscribers, not setting this property per the timeline would result in subscription operations failing. For subscribers whose notification endpoint already supports TLS 1.2, setting this property is optional. In such cases, Microsoft Graph defaults the property to v1_2. | [optional] 
**LifecycleNotificationUrl** | Pointer to **NullableString** | The URL of the endpoint that receives lifecycle notifications, including subscriptionRemoved and missed notifications. This URL must make use of the HTTPS protocol. Optional. Read more about how Outlook resources use lifecycle notifications. | [optional] 
**NotificationQueryOptions** | Pointer to **NullableString** | OData Query Options for specifying value for the targeting resource. Clients receive notifications when resource reaches the state matching the query options provided here. With this new property in the subscription creation payload along with all existing properties, Webhooks will deliver notifications whenever a resource reaches the desired state mentioned in the notificationQueryOptions property eg  when the print job is completed, when a print job resource isFetchable property value becomes true etc. | [optional] 
**NotificationUrl** | Pointer to **string** | Required. The URL of the endpoint that will receive the change notifications. This URL must make use of the HTTPS protocol. | [optional] 
**NotificationUrlAppId** | Pointer to **NullableString** |  | [optional] 
**Resource** | Pointer to **string** | Required. Specifies the resource that will be monitored for changes. Do not include the base URL (https://graph.microsoft.com/v1.0/). See the possible resource path values for each supported resource. | [optional] 

## Methods

### NewMicrosoftGraphSubscription

`func NewMicrosoftGraphSubscription() *MicrosoftGraphSubscription`

NewMicrosoftGraphSubscription instantiates a new MicrosoftGraphSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSubscriptionWithDefaults

`func NewMicrosoftGraphSubscriptionWithDefaults() *MicrosoftGraphSubscription`

NewMicrosoftGraphSubscriptionWithDefaults instantiates a new MicrosoftGraphSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApplicationId

`func (o *MicrosoftGraphSubscription) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MicrosoftGraphSubscription) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *MicrosoftGraphSubscription) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *MicrosoftGraphSubscription) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *MicrosoftGraphSubscription) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *MicrosoftGraphSubscription) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetChangeType

`func (o *MicrosoftGraphSubscription) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *MicrosoftGraphSubscription) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *MicrosoftGraphSubscription) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *MicrosoftGraphSubscription) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetClientState

`func (o *MicrosoftGraphSubscription) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *MicrosoftGraphSubscription) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *MicrosoftGraphSubscription) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *MicrosoftGraphSubscription) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### SetClientStateNil

`func (o *MicrosoftGraphSubscription) SetClientStateNil(b bool)`

 SetClientStateNil sets the value for ClientState to be an explicit nil

### UnsetClientState
`func (o *MicrosoftGraphSubscription) UnsetClientState()`

UnsetClientState ensures that no value is present for ClientState, not even an explicit nil
### GetCreatorId

`func (o *MicrosoftGraphSubscription) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *MicrosoftGraphSubscription) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *MicrosoftGraphSubscription) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *MicrosoftGraphSubscription) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *MicrosoftGraphSubscription) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *MicrosoftGraphSubscription) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetEncryptionCertificate

`func (o *MicrosoftGraphSubscription) GetEncryptionCertificate() string`

GetEncryptionCertificate returns the EncryptionCertificate field if non-nil, zero value otherwise.

### GetEncryptionCertificateOk

`func (o *MicrosoftGraphSubscription) GetEncryptionCertificateOk() (*string, bool)`

GetEncryptionCertificateOk returns a tuple with the EncryptionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionCertificate

`func (o *MicrosoftGraphSubscription) SetEncryptionCertificate(v string)`

SetEncryptionCertificate sets EncryptionCertificate field to given value.

### HasEncryptionCertificate

`func (o *MicrosoftGraphSubscription) HasEncryptionCertificate() bool`

HasEncryptionCertificate returns a boolean if a field has been set.

### SetEncryptionCertificateNil

`func (o *MicrosoftGraphSubscription) SetEncryptionCertificateNil(b bool)`

 SetEncryptionCertificateNil sets the value for EncryptionCertificate to be an explicit nil

### UnsetEncryptionCertificate
`func (o *MicrosoftGraphSubscription) UnsetEncryptionCertificate()`

UnsetEncryptionCertificate ensures that no value is present for EncryptionCertificate, not even an explicit nil
### GetEncryptionCertificateId

`func (o *MicrosoftGraphSubscription) GetEncryptionCertificateId() string`

GetEncryptionCertificateId returns the EncryptionCertificateId field if non-nil, zero value otherwise.

### GetEncryptionCertificateIdOk

`func (o *MicrosoftGraphSubscription) GetEncryptionCertificateIdOk() (*string, bool)`

GetEncryptionCertificateIdOk returns a tuple with the EncryptionCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionCertificateId

`func (o *MicrosoftGraphSubscription) SetEncryptionCertificateId(v string)`

SetEncryptionCertificateId sets EncryptionCertificateId field to given value.

### HasEncryptionCertificateId

`func (o *MicrosoftGraphSubscription) HasEncryptionCertificateId() bool`

HasEncryptionCertificateId returns a boolean if a field has been set.

### SetEncryptionCertificateIdNil

`func (o *MicrosoftGraphSubscription) SetEncryptionCertificateIdNil(b bool)`

 SetEncryptionCertificateIdNil sets the value for EncryptionCertificateId to be an explicit nil

### UnsetEncryptionCertificateId
`func (o *MicrosoftGraphSubscription) UnsetEncryptionCertificateId()`

UnsetEncryptionCertificateId ensures that no value is present for EncryptionCertificateId, not even an explicit nil
### GetExpirationDateTime

`func (o *MicrosoftGraphSubscription) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphSubscription) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphSubscription) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphSubscription) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetIncludeResourceData

`func (o *MicrosoftGraphSubscription) GetIncludeResourceData() bool`

GetIncludeResourceData returns the IncludeResourceData field if non-nil, zero value otherwise.

### GetIncludeResourceDataOk

`func (o *MicrosoftGraphSubscription) GetIncludeResourceDataOk() (*bool, bool)`

GetIncludeResourceDataOk returns a tuple with the IncludeResourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResourceData

`func (o *MicrosoftGraphSubscription) SetIncludeResourceData(v bool)`

SetIncludeResourceData sets IncludeResourceData field to given value.

### HasIncludeResourceData

`func (o *MicrosoftGraphSubscription) HasIncludeResourceData() bool`

HasIncludeResourceData returns a boolean if a field has been set.

### SetIncludeResourceDataNil

`func (o *MicrosoftGraphSubscription) SetIncludeResourceDataNil(b bool)`

 SetIncludeResourceDataNil sets the value for IncludeResourceData to be an explicit nil

### UnsetIncludeResourceData
`func (o *MicrosoftGraphSubscription) UnsetIncludeResourceData()`

UnsetIncludeResourceData ensures that no value is present for IncludeResourceData, not even an explicit nil
### GetLatestSupportedTlsVersion

`func (o *MicrosoftGraphSubscription) GetLatestSupportedTlsVersion() string`

GetLatestSupportedTlsVersion returns the LatestSupportedTlsVersion field if non-nil, zero value otherwise.

### GetLatestSupportedTlsVersionOk

`func (o *MicrosoftGraphSubscription) GetLatestSupportedTlsVersionOk() (*string, bool)`

GetLatestSupportedTlsVersionOk returns a tuple with the LatestSupportedTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSupportedTlsVersion

`func (o *MicrosoftGraphSubscription) SetLatestSupportedTlsVersion(v string)`

SetLatestSupportedTlsVersion sets LatestSupportedTlsVersion field to given value.

### HasLatestSupportedTlsVersion

`func (o *MicrosoftGraphSubscription) HasLatestSupportedTlsVersion() bool`

HasLatestSupportedTlsVersion returns a boolean if a field has been set.

### SetLatestSupportedTlsVersionNil

`func (o *MicrosoftGraphSubscription) SetLatestSupportedTlsVersionNil(b bool)`

 SetLatestSupportedTlsVersionNil sets the value for LatestSupportedTlsVersion to be an explicit nil

### UnsetLatestSupportedTlsVersion
`func (o *MicrosoftGraphSubscription) UnsetLatestSupportedTlsVersion()`

UnsetLatestSupportedTlsVersion ensures that no value is present for LatestSupportedTlsVersion, not even an explicit nil
### GetLifecycleNotificationUrl

`func (o *MicrosoftGraphSubscription) GetLifecycleNotificationUrl() string`

GetLifecycleNotificationUrl returns the LifecycleNotificationUrl field if non-nil, zero value otherwise.

### GetLifecycleNotificationUrlOk

`func (o *MicrosoftGraphSubscription) GetLifecycleNotificationUrlOk() (*string, bool)`

GetLifecycleNotificationUrlOk returns a tuple with the LifecycleNotificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleNotificationUrl

`func (o *MicrosoftGraphSubscription) SetLifecycleNotificationUrl(v string)`

SetLifecycleNotificationUrl sets LifecycleNotificationUrl field to given value.

### HasLifecycleNotificationUrl

`func (o *MicrosoftGraphSubscription) HasLifecycleNotificationUrl() bool`

HasLifecycleNotificationUrl returns a boolean if a field has been set.

### SetLifecycleNotificationUrlNil

`func (o *MicrosoftGraphSubscription) SetLifecycleNotificationUrlNil(b bool)`

 SetLifecycleNotificationUrlNil sets the value for LifecycleNotificationUrl to be an explicit nil

### UnsetLifecycleNotificationUrl
`func (o *MicrosoftGraphSubscription) UnsetLifecycleNotificationUrl()`

UnsetLifecycleNotificationUrl ensures that no value is present for LifecycleNotificationUrl, not even an explicit nil
### GetNotificationQueryOptions

`func (o *MicrosoftGraphSubscription) GetNotificationQueryOptions() string`

GetNotificationQueryOptions returns the NotificationQueryOptions field if non-nil, zero value otherwise.

### GetNotificationQueryOptionsOk

`func (o *MicrosoftGraphSubscription) GetNotificationQueryOptionsOk() (*string, bool)`

GetNotificationQueryOptionsOk returns a tuple with the NotificationQueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationQueryOptions

`func (o *MicrosoftGraphSubscription) SetNotificationQueryOptions(v string)`

SetNotificationQueryOptions sets NotificationQueryOptions field to given value.

### HasNotificationQueryOptions

`func (o *MicrosoftGraphSubscription) HasNotificationQueryOptions() bool`

HasNotificationQueryOptions returns a boolean if a field has been set.

### SetNotificationQueryOptionsNil

`func (o *MicrosoftGraphSubscription) SetNotificationQueryOptionsNil(b bool)`

 SetNotificationQueryOptionsNil sets the value for NotificationQueryOptions to be an explicit nil

### UnsetNotificationQueryOptions
`func (o *MicrosoftGraphSubscription) UnsetNotificationQueryOptions()`

UnsetNotificationQueryOptions ensures that no value is present for NotificationQueryOptions, not even an explicit nil
### GetNotificationUrl

`func (o *MicrosoftGraphSubscription) GetNotificationUrl() string`

GetNotificationUrl returns the NotificationUrl field if non-nil, zero value otherwise.

### GetNotificationUrlOk

`func (o *MicrosoftGraphSubscription) GetNotificationUrlOk() (*string, bool)`

GetNotificationUrlOk returns a tuple with the NotificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUrl

`func (o *MicrosoftGraphSubscription) SetNotificationUrl(v string)`

SetNotificationUrl sets NotificationUrl field to given value.

### HasNotificationUrl

`func (o *MicrosoftGraphSubscription) HasNotificationUrl() bool`

HasNotificationUrl returns a boolean if a field has been set.

### GetNotificationUrlAppId

`func (o *MicrosoftGraphSubscription) GetNotificationUrlAppId() string`

GetNotificationUrlAppId returns the NotificationUrlAppId field if non-nil, zero value otherwise.

### GetNotificationUrlAppIdOk

`func (o *MicrosoftGraphSubscription) GetNotificationUrlAppIdOk() (*string, bool)`

GetNotificationUrlAppIdOk returns a tuple with the NotificationUrlAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUrlAppId

`func (o *MicrosoftGraphSubscription) SetNotificationUrlAppId(v string)`

SetNotificationUrlAppId sets NotificationUrlAppId field to given value.

### HasNotificationUrlAppId

`func (o *MicrosoftGraphSubscription) HasNotificationUrlAppId() bool`

HasNotificationUrlAppId returns a boolean if a field has been set.

### SetNotificationUrlAppIdNil

`func (o *MicrosoftGraphSubscription) SetNotificationUrlAppIdNil(b bool)`

 SetNotificationUrlAppIdNil sets the value for NotificationUrlAppId to be an explicit nil

### UnsetNotificationUrlAppId
`func (o *MicrosoftGraphSubscription) UnsetNotificationUrlAppId()`

UnsetNotificationUrlAppId ensures that no value is present for NotificationUrlAppId, not even an explicit nil
### GetResource

`func (o *MicrosoftGraphSubscription) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphSubscription) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphSubscription) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphSubscription) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


