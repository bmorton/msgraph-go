# MicrosoftGraphAgreementAcceptance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AgreementFileId** | Pointer to **NullableString** | The identifier of the agreement file accepted by the user. | [optional] 
**AgreementId** | Pointer to **NullableString** | The identifier of the agreement. | [optional] 
**DeviceDisplayName** | Pointer to **NullableString** | The display name of the device used for accepting the agreement. | [optional] 
**DeviceId** | Pointer to **NullableString** | The unique identifier of the device used for accepting the agreement. | [optional] 
**DeviceOSType** | Pointer to **NullableString** | The operating system used to accept the agreement. | [optional] 
**DeviceOSVersion** | Pointer to **NullableString** | The operating system version of the device used to accept the agreement. | [optional] 
**ExpirationDateTime** | Pointer to **NullableTime** | The expiration date time of the acceptance. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**RecordedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphAgreementAcceptanceState**](anyOf&lt;microsoft.graph.agreementAcceptanceState&gt;.md) | The state of the agreement acceptance. Possible values are: accepted, declined. | [optional] 
**UserDisplayName** | Pointer to **NullableString** | Display name of the user when the acceptance was recorded. | [optional] 
**UserEmail** | Pointer to **NullableString** | Email of the user when the acceptance was recorded. | [optional] 
**UserId** | Pointer to **NullableString** | The identifier of the user who accepted the agreement. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | UPN of the user when the acceptance was recorded. | [optional] 

## Methods

### NewMicrosoftGraphAgreementAcceptance

`func NewMicrosoftGraphAgreementAcceptance() *MicrosoftGraphAgreementAcceptance`

NewMicrosoftGraphAgreementAcceptance instantiates a new MicrosoftGraphAgreementAcceptance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAgreementAcceptanceWithDefaults

`func NewMicrosoftGraphAgreementAcceptanceWithDefaults() *MicrosoftGraphAgreementAcceptance`

NewMicrosoftGraphAgreementAcceptanceWithDefaults instantiates a new MicrosoftGraphAgreementAcceptance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAgreementAcceptance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAgreementAcceptance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAgreementAcceptance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAgreementAcceptance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgreementFileId

`func (o *MicrosoftGraphAgreementAcceptance) GetAgreementFileId() string`

GetAgreementFileId returns the AgreementFileId field if non-nil, zero value otherwise.

### GetAgreementFileIdOk

`func (o *MicrosoftGraphAgreementAcceptance) GetAgreementFileIdOk() (*string, bool)`

GetAgreementFileIdOk returns a tuple with the AgreementFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementFileId

`func (o *MicrosoftGraphAgreementAcceptance) SetAgreementFileId(v string)`

SetAgreementFileId sets AgreementFileId field to given value.

### HasAgreementFileId

`func (o *MicrosoftGraphAgreementAcceptance) HasAgreementFileId() bool`

HasAgreementFileId returns a boolean if a field has been set.

### SetAgreementFileIdNil

`func (o *MicrosoftGraphAgreementAcceptance) SetAgreementFileIdNil(b bool)`

 SetAgreementFileIdNil sets the value for AgreementFileId to be an explicit nil

### UnsetAgreementFileId
`func (o *MicrosoftGraphAgreementAcceptance) UnsetAgreementFileId()`

UnsetAgreementFileId ensures that no value is present for AgreementFileId, not even an explicit nil
### GetAgreementId

`func (o *MicrosoftGraphAgreementAcceptance) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *MicrosoftGraphAgreementAcceptance) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *MicrosoftGraphAgreementAcceptance) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *MicrosoftGraphAgreementAcceptance) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *MicrosoftGraphAgreementAcceptance) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *MicrosoftGraphAgreementAcceptance) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetDeviceDisplayName

`func (o *MicrosoftGraphAgreementAcceptance) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *MicrosoftGraphAgreementAcceptance) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *MicrosoftGraphAgreementAcceptance) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.

### HasDeviceDisplayName

`func (o *MicrosoftGraphAgreementAcceptance) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### SetDeviceDisplayNameNil

`func (o *MicrosoftGraphAgreementAcceptance) SetDeviceDisplayNameNil(b bool)`

 SetDeviceDisplayNameNil sets the value for DeviceDisplayName to be an explicit nil

### UnsetDeviceDisplayName
`func (o *MicrosoftGraphAgreementAcceptance) UnsetDeviceDisplayName()`

UnsetDeviceDisplayName ensures that no value is present for DeviceDisplayName, not even an explicit nil
### GetDeviceId

`func (o *MicrosoftGraphAgreementAcceptance) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphAgreementAcceptance) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MicrosoftGraphAgreementAcceptance) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MicrosoftGraphAgreementAcceptance) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *MicrosoftGraphAgreementAcceptance) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *MicrosoftGraphAgreementAcceptance) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceOSType

`func (o *MicrosoftGraphAgreementAcceptance) GetDeviceOSType() string`

GetDeviceOSType returns the DeviceOSType field if non-nil, zero value otherwise.

### GetDeviceOSTypeOk

`func (o *MicrosoftGraphAgreementAcceptance) GetDeviceOSTypeOk() (*string, bool)`

GetDeviceOSTypeOk returns a tuple with the DeviceOSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOSType

`func (o *MicrosoftGraphAgreementAcceptance) SetDeviceOSType(v string)`

SetDeviceOSType sets DeviceOSType field to given value.

### HasDeviceOSType

`func (o *MicrosoftGraphAgreementAcceptance) HasDeviceOSType() bool`

HasDeviceOSType returns a boolean if a field has been set.

### SetDeviceOSTypeNil

`func (o *MicrosoftGraphAgreementAcceptance) SetDeviceOSTypeNil(b bool)`

 SetDeviceOSTypeNil sets the value for DeviceOSType to be an explicit nil

### UnsetDeviceOSType
`func (o *MicrosoftGraphAgreementAcceptance) UnsetDeviceOSType()`

UnsetDeviceOSType ensures that no value is present for DeviceOSType, not even an explicit nil
### GetDeviceOSVersion

`func (o *MicrosoftGraphAgreementAcceptance) GetDeviceOSVersion() string`

GetDeviceOSVersion returns the DeviceOSVersion field if non-nil, zero value otherwise.

### GetDeviceOSVersionOk

`func (o *MicrosoftGraphAgreementAcceptance) GetDeviceOSVersionOk() (*string, bool)`

GetDeviceOSVersionOk returns a tuple with the DeviceOSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOSVersion

`func (o *MicrosoftGraphAgreementAcceptance) SetDeviceOSVersion(v string)`

SetDeviceOSVersion sets DeviceOSVersion field to given value.

### HasDeviceOSVersion

`func (o *MicrosoftGraphAgreementAcceptance) HasDeviceOSVersion() bool`

HasDeviceOSVersion returns a boolean if a field has been set.

### SetDeviceOSVersionNil

`func (o *MicrosoftGraphAgreementAcceptance) SetDeviceOSVersionNil(b bool)`

 SetDeviceOSVersionNil sets the value for DeviceOSVersion to be an explicit nil

### UnsetDeviceOSVersion
`func (o *MicrosoftGraphAgreementAcceptance) UnsetDeviceOSVersion()`

UnsetDeviceOSVersion ensures that no value is present for DeviceOSVersion, not even an explicit nil
### GetExpirationDateTime

`func (o *MicrosoftGraphAgreementAcceptance) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphAgreementAcceptance) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphAgreementAcceptance) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphAgreementAcceptance) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *MicrosoftGraphAgreementAcceptance) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *MicrosoftGraphAgreementAcceptance) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetRecordedDateTime

`func (o *MicrosoftGraphAgreementAcceptance) GetRecordedDateTime() time.Time`

GetRecordedDateTime returns the RecordedDateTime field if non-nil, zero value otherwise.

### GetRecordedDateTimeOk

`func (o *MicrosoftGraphAgreementAcceptance) GetRecordedDateTimeOk() (*time.Time, bool)`

GetRecordedDateTimeOk returns a tuple with the RecordedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedDateTime

`func (o *MicrosoftGraphAgreementAcceptance) SetRecordedDateTime(v time.Time)`

SetRecordedDateTime sets RecordedDateTime field to given value.

### HasRecordedDateTime

`func (o *MicrosoftGraphAgreementAcceptance) HasRecordedDateTime() bool`

HasRecordedDateTime returns a boolean if a field has been set.

### SetRecordedDateTimeNil

`func (o *MicrosoftGraphAgreementAcceptance) SetRecordedDateTimeNil(b bool)`

 SetRecordedDateTimeNil sets the value for RecordedDateTime to be an explicit nil

### UnsetRecordedDateTime
`func (o *MicrosoftGraphAgreementAcceptance) UnsetRecordedDateTime()`

UnsetRecordedDateTime ensures that no value is present for RecordedDateTime, not even an explicit nil
### GetState

`func (o *MicrosoftGraphAgreementAcceptance) GetState() AnyOfmicrosoftGraphAgreementAcceptanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphAgreementAcceptance) GetStateOk() (*AnyOfmicrosoftGraphAgreementAcceptanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphAgreementAcceptance) SetState(v AnyOfmicrosoftGraphAgreementAcceptanceState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphAgreementAcceptance) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphAgreementAcceptance) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphAgreementAcceptance) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetUserDisplayName

`func (o *MicrosoftGraphAgreementAcceptance) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphAgreementAcceptance) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphAgreementAcceptance) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *MicrosoftGraphAgreementAcceptance) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayNameNil

`func (o *MicrosoftGraphAgreementAcceptance) SetUserDisplayNameNil(b bool)`

 SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil

### UnsetUserDisplayName
`func (o *MicrosoftGraphAgreementAcceptance) UnsetUserDisplayName()`

UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
### GetUserEmail

`func (o *MicrosoftGraphAgreementAcceptance) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *MicrosoftGraphAgreementAcceptance) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *MicrosoftGraphAgreementAcceptance) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *MicrosoftGraphAgreementAcceptance) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *MicrosoftGraphAgreementAcceptance) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *MicrosoftGraphAgreementAcceptance) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphAgreementAcceptance) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphAgreementAcceptance) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphAgreementAcceptance) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphAgreementAcceptance) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MicrosoftGraphAgreementAcceptance) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MicrosoftGraphAgreementAcceptance) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphAgreementAcceptance) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphAgreementAcceptance) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphAgreementAcceptance) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphAgreementAcceptance) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphAgreementAcceptance) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphAgreementAcceptance) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


