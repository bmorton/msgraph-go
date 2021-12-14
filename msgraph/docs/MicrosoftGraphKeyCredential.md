# MicrosoftGraphKeyCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomKeyIdentifier** | Pointer to **NullableString** | Custom key identifier | [optional] 
**DisplayName** | Pointer to **NullableString** | Friendly name for the key. Optional. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | The date and time at which the credential expires.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**Key** | Pointer to **NullableString** | The certificate&#39;s raw data in byte array converted to Base64 string; for example, [System.Convert]::ToBase64String($Cert.GetRawCertData()). | [optional] 
**KeyId** | Pointer to **NullableString** | The unique identifier (GUID) for the key. | [optional] 
**StartDateTime** | Pointer to **NullableTime** | The date and time at which the credential becomes valid.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**Type** | Pointer to **NullableString** | The type of key credential; for example, Symmetric. | [optional] 
**Usage** | Pointer to **NullableString** | A string that describes the purpose for which the key can be used; for example, Verify. | [optional] 

## Methods

### NewMicrosoftGraphKeyCredential

`func NewMicrosoftGraphKeyCredential() *MicrosoftGraphKeyCredential`

NewMicrosoftGraphKeyCredential instantiates a new MicrosoftGraphKeyCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphKeyCredentialWithDefaults

`func NewMicrosoftGraphKeyCredentialWithDefaults() *MicrosoftGraphKeyCredential`

NewMicrosoftGraphKeyCredentialWithDefaults instantiates a new MicrosoftGraphKeyCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomKeyIdentifier

`func (o *MicrosoftGraphKeyCredential) GetCustomKeyIdentifier() string`

GetCustomKeyIdentifier returns the CustomKeyIdentifier field if non-nil, zero value otherwise.

### GetCustomKeyIdentifierOk

`func (o *MicrosoftGraphKeyCredential) GetCustomKeyIdentifierOk() (*string, bool)`

GetCustomKeyIdentifierOk returns a tuple with the CustomKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKeyIdentifier

`func (o *MicrosoftGraphKeyCredential) SetCustomKeyIdentifier(v string)`

SetCustomKeyIdentifier sets CustomKeyIdentifier field to given value.

### HasCustomKeyIdentifier

`func (o *MicrosoftGraphKeyCredential) HasCustomKeyIdentifier() bool`

HasCustomKeyIdentifier returns a boolean if a field has been set.

### SetCustomKeyIdentifierNil

`func (o *MicrosoftGraphKeyCredential) SetCustomKeyIdentifierNil(b bool)`

 SetCustomKeyIdentifierNil sets the value for CustomKeyIdentifier to be an explicit nil

### UnsetCustomKeyIdentifier
`func (o *MicrosoftGraphKeyCredential) UnsetCustomKeyIdentifier()`

UnsetCustomKeyIdentifier ensures that no value is present for CustomKeyIdentifier, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphKeyCredential) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphKeyCredential) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphKeyCredential) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphKeyCredential) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphKeyCredential) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphKeyCredential) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEndDateTime

`func (o *MicrosoftGraphKeyCredential) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphKeyCredential) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphKeyCredential) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphKeyCredential) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphKeyCredential) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphKeyCredential) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetKey

`func (o *MicrosoftGraphKeyCredential) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MicrosoftGraphKeyCredential) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MicrosoftGraphKeyCredential) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MicrosoftGraphKeyCredential) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *MicrosoftGraphKeyCredential) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *MicrosoftGraphKeyCredential) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetKeyId

`func (o *MicrosoftGraphKeyCredential) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *MicrosoftGraphKeyCredential) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *MicrosoftGraphKeyCredential) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *MicrosoftGraphKeyCredential) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### SetKeyIdNil

`func (o *MicrosoftGraphKeyCredential) SetKeyIdNil(b bool)`

 SetKeyIdNil sets the value for KeyId to be an explicit nil

### UnsetKeyId
`func (o *MicrosoftGraphKeyCredential) UnsetKeyId()`

UnsetKeyId ensures that no value is present for KeyId, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphKeyCredential) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphKeyCredential) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphKeyCredential) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphKeyCredential) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphKeyCredential) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphKeyCredential) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetType

`func (o *MicrosoftGraphKeyCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphKeyCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphKeyCredential) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphKeyCredential) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphKeyCredential) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphKeyCredential) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsage

`func (o *MicrosoftGraphKeyCredential) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MicrosoftGraphKeyCredential) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MicrosoftGraphKeyCredential) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *MicrosoftGraphKeyCredential) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *MicrosoftGraphKeyCredential) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *MicrosoftGraphKeyCredential) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


