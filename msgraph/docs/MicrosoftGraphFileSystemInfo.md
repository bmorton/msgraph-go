# MicrosoftGraphFileSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | The UTC date and time the file was created on a client. | [optional] 
**LastAccessedDateTime** | Pointer to **NullableTime** | The UTC date and time the file was last accessed. Available for the recent file list only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The UTC date and time the file was last modified on a client. | [optional] 

## Methods

### NewMicrosoftGraphFileSystemInfo

`func NewMicrosoftGraphFileSystemInfo() *MicrosoftGraphFileSystemInfo`

NewMicrosoftGraphFileSystemInfo instantiates a new MicrosoftGraphFileSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFileSystemInfoWithDefaults

`func NewMicrosoftGraphFileSystemInfoWithDefaults() *MicrosoftGraphFileSystemInfo`

NewMicrosoftGraphFileSystemInfoWithDefaults instantiates a new MicrosoftGraphFileSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *MicrosoftGraphFileSystemInfo) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphFileSystemInfo) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphFileSystemInfo) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphFileSystemInfo) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphFileSystemInfo) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphFileSystemInfo) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastAccessedDateTime

`func (o *MicrosoftGraphFileSystemInfo) GetLastAccessedDateTime() time.Time`

GetLastAccessedDateTime returns the LastAccessedDateTime field if non-nil, zero value otherwise.

### GetLastAccessedDateTimeOk

`func (o *MicrosoftGraphFileSystemInfo) GetLastAccessedDateTimeOk() (*time.Time, bool)`

GetLastAccessedDateTimeOk returns a tuple with the LastAccessedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedDateTime

`func (o *MicrosoftGraphFileSystemInfo) SetLastAccessedDateTime(v time.Time)`

SetLastAccessedDateTime sets LastAccessedDateTime field to given value.

### HasLastAccessedDateTime

`func (o *MicrosoftGraphFileSystemInfo) HasLastAccessedDateTime() bool`

HasLastAccessedDateTime returns a boolean if a field has been set.

### SetLastAccessedDateTimeNil

`func (o *MicrosoftGraphFileSystemInfo) SetLastAccessedDateTimeNil(b bool)`

 SetLastAccessedDateTimeNil sets the value for LastAccessedDateTime to be an explicit nil

### UnsetLastAccessedDateTime
`func (o *MicrosoftGraphFileSystemInfo) UnsetLastAccessedDateTime()`

UnsetLastAccessedDateTime ensures that no value is present for LastAccessedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphFileSystemInfo) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphFileSystemInfo) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphFileSystemInfo) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphFileSystemInfo) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphFileSystemInfo) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphFileSystemInfo) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


