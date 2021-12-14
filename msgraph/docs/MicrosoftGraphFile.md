# MicrosoftGraphFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashes** | Pointer to [**AnyOfmicrosoftGraphHashes**](anyOf&lt;microsoft.graph.hashes&gt;.md) | Hashes of the file&#39;s binary content, if available. Read-only. | [optional] 
**MimeType** | Pointer to **NullableString** | The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only. | [optional] 
**ProcessingMetadata** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMicrosoftGraphFile

`func NewMicrosoftGraphFile() *MicrosoftGraphFile`

NewMicrosoftGraphFile instantiates a new MicrosoftGraphFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFileWithDefaults

`func NewMicrosoftGraphFileWithDefaults() *MicrosoftGraphFile`

NewMicrosoftGraphFileWithDefaults instantiates a new MicrosoftGraphFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashes

`func (o *MicrosoftGraphFile) GetHashes() AnyOfmicrosoftGraphHashes`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *MicrosoftGraphFile) GetHashesOk() (*AnyOfmicrosoftGraphHashes, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *MicrosoftGraphFile) SetHashes(v AnyOfmicrosoftGraphHashes)`

SetHashes sets Hashes field to given value.

### HasHashes

`func (o *MicrosoftGraphFile) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### SetHashesNil

`func (o *MicrosoftGraphFile) SetHashesNil(b bool)`

 SetHashesNil sets the value for Hashes to be an explicit nil

### UnsetHashes
`func (o *MicrosoftGraphFile) UnsetHashes()`

UnsetHashes ensures that no value is present for Hashes, not even an explicit nil
### GetMimeType

`func (o *MicrosoftGraphFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *MicrosoftGraphFile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *MicrosoftGraphFile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *MicrosoftGraphFile) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *MicrosoftGraphFile) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *MicrosoftGraphFile) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetProcessingMetadata

`func (o *MicrosoftGraphFile) GetProcessingMetadata() bool`

GetProcessingMetadata returns the ProcessingMetadata field if non-nil, zero value otherwise.

### GetProcessingMetadataOk

`func (o *MicrosoftGraphFile) GetProcessingMetadataOk() (*bool, bool)`

GetProcessingMetadataOk returns a tuple with the ProcessingMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMetadata

`func (o *MicrosoftGraphFile) SetProcessingMetadata(v bool)`

SetProcessingMetadata sets ProcessingMetadata field to given value.

### HasProcessingMetadata

`func (o *MicrosoftGraphFile) HasProcessingMetadata() bool`

HasProcessingMetadata returns a boolean if a field has been set.

### SetProcessingMetadataNil

`func (o *MicrosoftGraphFile) SetProcessingMetadataNil(b bool)`

 SetProcessingMetadataNil sets the value for ProcessingMetadata to be an explicit nil

### UnsetProcessingMetadata
`func (o *MicrosoftGraphFile) UnsetProcessingMetadata()`

UnsetProcessingMetadata ensures that no value is present for ProcessingMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


