# MicrosoftGraphWorkbookSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id of the workbook session. | [optional] 
**PersistChanges** | Pointer to **NullableBool** | true for persistent session. false for non-persistent session (view mode) | [optional] 

## Methods

### NewMicrosoftGraphWorkbookSessionInfo

`func NewMicrosoftGraphWorkbookSessionInfo() *MicrosoftGraphWorkbookSessionInfo`

NewMicrosoftGraphWorkbookSessionInfo instantiates a new MicrosoftGraphWorkbookSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookSessionInfoWithDefaults

`func NewMicrosoftGraphWorkbookSessionInfoWithDefaults() *MicrosoftGraphWorkbookSessionInfo`

NewMicrosoftGraphWorkbookSessionInfoWithDefaults instantiates a new MicrosoftGraphWorkbookSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookSessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookSessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookSessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookSessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphWorkbookSessionInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphWorkbookSessionInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPersistChanges

`func (o *MicrosoftGraphWorkbookSessionInfo) GetPersistChanges() bool`

GetPersistChanges returns the PersistChanges field if non-nil, zero value otherwise.

### GetPersistChangesOk

`func (o *MicrosoftGraphWorkbookSessionInfo) GetPersistChangesOk() (*bool, bool)`

GetPersistChangesOk returns a tuple with the PersistChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistChanges

`func (o *MicrosoftGraphWorkbookSessionInfo) SetPersistChanges(v bool)`

SetPersistChanges sets PersistChanges field to given value.

### HasPersistChanges

`func (o *MicrosoftGraphWorkbookSessionInfo) HasPersistChanges() bool`

HasPersistChanges returns a boolean if a field has been set.

### SetPersistChangesNil

`func (o *MicrosoftGraphWorkbookSessionInfo) SetPersistChangesNil(b bool)`

 SetPersistChangesNil sets the value for PersistChanges to be an explicit nil

### UnsetPersistChanges
`func (o *MicrosoftGraphWorkbookSessionInfo) UnsetPersistChanges()`

UnsetPersistChanges ensures that no value is present for PersistChanges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


