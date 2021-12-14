# MicrosoftGraphApplePushNotificationCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AppleIdentifier** | Pointer to **NullableString** | Apple Id of the account used to create the MDM push certificate. | [optional] 
**Certificate** | Pointer to **NullableString** | Not yet documented | [optional] 
**CertificateSerialNumber** | Pointer to **NullableString** | Certificate serial number. This property is read-only. | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | The expiration date and time for Apple push notification certificate. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | Last modified date and time for Apple push notification certificate. | [optional] 
**TopicIdentifier** | Pointer to **NullableString** | Topic Id. | [optional] 

## Methods

### NewMicrosoftGraphApplePushNotificationCertificate

`func NewMicrosoftGraphApplePushNotificationCertificate() *MicrosoftGraphApplePushNotificationCertificate`

NewMicrosoftGraphApplePushNotificationCertificate instantiates a new MicrosoftGraphApplePushNotificationCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphApplePushNotificationCertificateWithDefaults

`func NewMicrosoftGraphApplePushNotificationCertificateWithDefaults() *MicrosoftGraphApplePushNotificationCertificate`

NewMicrosoftGraphApplePushNotificationCertificateWithDefaults instantiates a new MicrosoftGraphApplePushNotificationCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphApplePushNotificationCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppleIdentifier

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetAppleIdentifier() string`

GetAppleIdentifier returns the AppleIdentifier field if non-nil, zero value otherwise.

### GetAppleIdentifierOk

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetAppleIdentifierOk() (*string, bool)`

GetAppleIdentifierOk returns a tuple with the AppleIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleIdentifier

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetAppleIdentifier(v string)`

SetAppleIdentifier sets AppleIdentifier field to given value.

### HasAppleIdentifier

`func (o *MicrosoftGraphApplePushNotificationCertificate) HasAppleIdentifier() bool`

HasAppleIdentifier returns a boolean if a field has been set.

### SetAppleIdentifierNil

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetAppleIdentifierNil(b bool)`

 SetAppleIdentifierNil sets the value for AppleIdentifier to be an explicit nil

### UnsetAppleIdentifier
`func (o *MicrosoftGraphApplePushNotificationCertificate) UnsetAppleIdentifier()`

UnsetAppleIdentifier ensures that no value is present for AppleIdentifier, not even an explicit nil
### GetCertificate

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *MicrosoftGraphApplePushNotificationCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *MicrosoftGraphApplePushNotificationCertificate) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateSerialNumber

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetCertificateSerialNumber() string`

GetCertificateSerialNumber returns the CertificateSerialNumber field if non-nil, zero value otherwise.

### GetCertificateSerialNumberOk

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetCertificateSerialNumberOk() (*string, bool)`

GetCertificateSerialNumberOk returns a tuple with the CertificateSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSerialNumber

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetCertificateSerialNumber(v string)`

SetCertificateSerialNumber sets CertificateSerialNumber field to given value.

### HasCertificateSerialNumber

`func (o *MicrosoftGraphApplePushNotificationCertificate) HasCertificateSerialNumber() bool`

HasCertificateSerialNumber returns a boolean if a field has been set.

### SetCertificateSerialNumberNil

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetCertificateSerialNumberNil(b bool)`

 SetCertificateSerialNumberNil sets the value for CertificateSerialNumber to be an explicit nil

### UnsetCertificateSerialNumber
`func (o *MicrosoftGraphApplePushNotificationCertificate) UnsetCertificateSerialNumber()`

UnsetCertificateSerialNumber ensures that no value is present for CertificateSerialNumber, not even an explicit nil
### GetExpirationDateTime

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphApplePushNotificationCertificate) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphApplePushNotificationCertificate) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetTopicIdentifier

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetTopicIdentifier() string`

GetTopicIdentifier returns the TopicIdentifier field if non-nil, zero value otherwise.

### GetTopicIdentifierOk

`func (o *MicrosoftGraphApplePushNotificationCertificate) GetTopicIdentifierOk() (*string, bool)`

GetTopicIdentifierOk returns a tuple with the TopicIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicIdentifier

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetTopicIdentifier(v string)`

SetTopicIdentifier sets TopicIdentifier field to given value.

### HasTopicIdentifier

`func (o *MicrosoftGraphApplePushNotificationCertificate) HasTopicIdentifier() bool`

HasTopicIdentifier returns a boolean if a field has been set.

### SetTopicIdentifierNil

`func (o *MicrosoftGraphApplePushNotificationCertificate) SetTopicIdentifierNil(b bool)`

 SetTopicIdentifierNil sets the value for TopicIdentifier to be an explicit nil

### UnsetTopicIdentifier
`func (o *MicrosoftGraphApplePushNotificationCertificate) UnsetTopicIdentifier()`

UnsetTopicIdentifier ensures that no value is present for TopicIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


