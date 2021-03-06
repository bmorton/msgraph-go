# MicrosoftGraphDirectoryObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewMicrosoftGraphDirectoryObject

`func NewMicrosoftGraphDirectoryObject() *MicrosoftGraphDirectoryObject`

NewMicrosoftGraphDirectoryObject instantiates a new MicrosoftGraphDirectoryObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDirectoryObjectWithDefaults

`func NewMicrosoftGraphDirectoryObjectWithDefaults() *MicrosoftGraphDirectoryObject`

NewMicrosoftGraphDirectoryObjectWithDefaults instantiates a new MicrosoftGraphDirectoryObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDirectoryObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDirectoryObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDirectoryObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDirectoryObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphDirectoryObject) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphDirectoryObject) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphDirectoryObject) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphDirectoryObject) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphDirectoryObject) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphDirectoryObject) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


