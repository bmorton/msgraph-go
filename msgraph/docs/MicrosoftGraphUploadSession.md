# MicrosoftGraphUploadSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDateTime** | Pointer to **NullableTime** | The date and time in UTC that the upload session will expire. The complete file must be uploaded before this expiration time is reached. | [optional] 
**NextExpectedRanges** | Pointer to **[]string** | A collection of byte ranges that the server is missing for the file. These ranges are zero indexed and of the format &#39;start-end&#39; (e.g. &#39;0-26&#39; to indicate the first 27 bytes of the file). When uploading files as Outlook attachments, instead of a collection of ranges, this property always indicates a single value &#39;{start}&#39;, the location in the file where the next upload should begin. | [optional] 
**UploadUrl** | Pointer to **NullableString** | The URL endpoint that accepts PUT requests for byte ranges of the file. | [optional] 

## Methods

### NewMicrosoftGraphUploadSession

`func NewMicrosoftGraphUploadSession() *MicrosoftGraphUploadSession`

NewMicrosoftGraphUploadSession instantiates a new MicrosoftGraphUploadSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUploadSessionWithDefaults

`func NewMicrosoftGraphUploadSessionWithDefaults() *MicrosoftGraphUploadSession`

NewMicrosoftGraphUploadSessionWithDefaults instantiates a new MicrosoftGraphUploadSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDateTime

`func (o *MicrosoftGraphUploadSession) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphUploadSession) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphUploadSession) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphUploadSession) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *MicrosoftGraphUploadSession) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *MicrosoftGraphUploadSession) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetNextExpectedRanges

`func (o *MicrosoftGraphUploadSession) GetNextExpectedRanges() []*string`

GetNextExpectedRanges returns the NextExpectedRanges field if non-nil, zero value otherwise.

### GetNextExpectedRangesOk

`func (o *MicrosoftGraphUploadSession) GetNextExpectedRangesOk() (*[]*string, bool)`

GetNextExpectedRangesOk returns a tuple with the NextExpectedRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExpectedRanges

`func (o *MicrosoftGraphUploadSession) SetNextExpectedRanges(v []*string)`

SetNextExpectedRanges sets NextExpectedRanges field to given value.

### HasNextExpectedRanges

`func (o *MicrosoftGraphUploadSession) HasNextExpectedRanges() bool`

HasNextExpectedRanges returns a boolean if a field has been set.

### GetUploadUrl

`func (o *MicrosoftGraphUploadSession) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *MicrosoftGraphUploadSession) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *MicrosoftGraphUploadSession) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *MicrosoftGraphUploadSession) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### SetUploadUrlNil

`func (o *MicrosoftGraphUploadSession) SetUploadUrlNil(b bool)`

 SetUploadUrlNil sets the value for UploadUrl to be an explicit nil

### UnsetUploadUrl
`func (o *MicrosoftGraphUploadSession) UnsetUploadUrl()`

UnsetUploadUrl ensures that no value is present for UploadUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


