# MicrosoftGraphManagedAppOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The operation name. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The last time the app operation was modified. | [optional] 
**State** | Pointer to **NullableString** | The current state of the operation | [optional] 
**Version** | Pointer to **NullableString** | Version of the entity. | [optional] 

## Methods

### NewMicrosoftGraphManagedAppOperation

`func NewMicrosoftGraphManagedAppOperation() *MicrosoftGraphManagedAppOperation`

NewMicrosoftGraphManagedAppOperation instantiates a new MicrosoftGraphManagedAppOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphManagedAppOperationWithDefaults

`func NewMicrosoftGraphManagedAppOperationWithDefaults() *MicrosoftGraphManagedAppOperation`

NewMicrosoftGraphManagedAppOperationWithDefaults instantiates a new MicrosoftGraphManagedAppOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphManagedAppOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedAppOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphManagedAppOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphManagedAppOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphManagedAppOperation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedAppOperation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedAppOperation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphManagedAppOperation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphManagedAppOperation) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphManagedAppOperation) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppOperation) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedAppOperation) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppOperation) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppOperation) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetState

`func (o *MicrosoftGraphManagedAppOperation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphManagedAppOperation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphManagedAppOperation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphManagedAppOperation) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphManagedAppOperation) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphManagedAppOperation) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetVersion

`func (o *MicrosoftGraphManagedAppOperation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedAppOperation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphManagedAppOperation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphManagedAppOperation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphManagedAppOperation) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphManagedAppOperation) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


