# ApplePushNotificationCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppleIdentifier** | Pointer to **NullableString** | Apple Id of the account used to create the MDM push certificate. | [optional] 
**Certificate** | Pointer to **NullableString** | Not yet documented | [optional] 
**CertificateSerialNumber** | Pointer to **NullableString** | Certificate serial number. This property is read-only. | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | The expiration date and time for Apple push notification certificate. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Last modified date and time for Apple push notification certificate. | [optional] 
**TopicIdentifier** | Pointer to **NullableString** | Topic Id. | [optional] 

## Methods

### NewApplePushNotificationCertificate

`func NewApplePushNotificationCertificate() *ApplePushNotificationCertificate`

NewApplePushNotificationCertificate instantiates a new ApplePushNotificationCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePushNotificationCertificateWithDefaults

`func NewApplePushNotificationCertificateWithDefaults() *ApplePushNotificationCertificate`

NewApplePushNotificationCertificateWithDefaults instantiates a new ApplePushNotificationCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppleIdentifier

`func (o *ApplePushNotificationCertificate) GetAppleIdentifier() string`

GetAppleIdentifier returns the AppleIdentifier field if non-nil, zero value otherwise.

### GetAppleIdentifierOk

`func (o *ApplePushNotificationCertificate) GetAppleIdentifierOk() (*string, bool)`

GetAppleIdentifierOk returns a tuple with the AppleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleIdentifier

`func (o *ApplePushNotificationCertificate) SetAppleIdentifier(v string)`

SetAppleIdentifier sets AppleIdentifier field to given value.

### HasAppleIdentifier

`func (o *ApplePushNotificationCertificate) HasAppleIdentifier() bool`

HasAppleIdentifier returns a boolean if a field has been set.

### SetAppleIdentifierNil

`func (o *ApplePushNotificationCertificate) SetAppleIdentifierNil(b bool)`

 SetAppleIdentifierNil sets the value for AppleIdentifier to be an explicit nil

### UnsetAppleIdentifier
`func (o *ApplePushNotificationCertificate) UnsetAppleIdentifier()`

UnsetAppleIdentifier ensures that no value is present for AppleIdentifier, not even an explicit nil
### GetCertificate

`func (o *ApplePushNotificationCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ApplePushNotificationCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ApplePushNotificationCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ApplePushNotificationCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ApplePushNotificationCertificate) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ApplePushNotificationCertificate) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateSerialNumber

`func (o *ApplePushNotificationCertificate) GetCertificateSerialNumber() string`

GetCertificateSerialNumber returns the CertificateSerialNumber field if non-nil, zero value otherwise.

### GetCertificateSerialNumberOk

`func (o *ApplePushNotificationCertificate) GetCertificateSerialNumberOk() (*string, bool)`

GetCertificateSerialNumberOk returns a tuple with the CertificateSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSerialNumber

`func (o *ApplePushNotificationCertificate) SetCertificateSerialNumber(v string)`

SetCertificateSerialNumber sets CertificateSerialNumber field to given value.

### HasCertificateSerialNumber

`func (o *ApplePushNotificationCertificate) HasCertificateSerialNumber() bool`

HasCertificateSerialNumber returns a boolean if a field has been set.

### SetCertificateSerialNumberNil

`func (o *ApplePushNotificationCertificate) SetCertificateSerialNumberNil(b bool)`

 SetCertificateSerialNumberNil sets the value for CertificateSerialNumber to be an explicit nil

### UnsetCertificateSerialNumber
`func (o *ApplePushNotificationCertificate) UnsetCertificateSerialNumber()`

UnsetCertificateSerialNumber ensures that no value is present for CertificateSerialNumber, not even an explicit nil
### GetExpirationDateTime

`func (o *ApplePushNotificationCertificate) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *ApplePushNotificationCertificate) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *ApplePushNotificationCertificate) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *ApplePushNotificationCertificate) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *ApplePushNotificationCertificate) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *ApplePushNotificationCertificate) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *ApplePushNotificationCertificate) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *ApplePushNotificationCertificate) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetTopicIdentifier

`func (o *ApplePushNotificationCertificate) GetTopicIdentifier() string`

GetTopicIdentifier returns the TopicIdentifier field if non-nil, zero value otherwise.

### GetTopicIdentifierOk

`func (o *ApplePushNotificationCertificate) GetTopicIdentifierOk() (*string, bool)`

GetTopicIdentifierOk returns a tuple with the TopicIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicIdentifier

`func (o *ApplePushNotificationCertificate) SetTopicIdentifier(v string)`

SetTopicIdentifier sets TopicIdentifier field to given value.

### HasTopicIdentifier

`func (o *ApplePushNotificationCertificate) HasTopicIdentifier() bool`

HasTopicIdentifier returns a boolean if a field has been set.

### SetTopicIdentifierNil

`func (o *ApplePushNotificationCertificate) SetTopicIdentifierNil(b bool)`

 SetTopicIdentifierNil sets the value for TopicIdentifier to be an explicit nil

### UnsetTopicIdentifier
`func (o *ApplePushNotificationCertificate) UnsetTopicIdentifier()`

UnsetTopicIdentifier ensures that no value is present for TopicIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


