# MicrosoftGraphDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AdministrativeUnits** | Pointer to [**[]MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md) | Conceptual container for user and group directory objects. | [optional] 
**DeletedItems** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Recently deleted items. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphDirectory

`func NewMicrosoftGraphDirectory() *MicrosoftGraphDirectory`

NewMicrosoftGraphDirectory instantiates a new MicrosoftGraphDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDirectoryWithDefaults

`func NewMicrosoftGraphDirectoryWithDefaults() *MicrosoftGraphDirectory`

NewMicrosoftGraphDirectoryWithDefaults instantiates a new MicrosoftGraphDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDirectory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDirectory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDirectory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDirectory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAdministrativeUnits

`func (o *MicrosoftGraphDirectory) GetAdministrativeUnits() []MicrosoftGraphAdministrativeUnit`

GetAdministrativeUnits returns the AdministrativeUnits field if non-nil, zero value otherwise.

### GetAdministrativeUnitsOk

`func (o *MicrosoftGraphDirectory) GetAdministrativeUnitsOk() (*[]MicrosoftGraphAdministrativeUnit, bool)`

GetAdministrativeUnitsOk returns a tuple with the AdministrativeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeUnits

`func (o *MicrosoftGraphDirectory) SetAdministrativeUnits(v []MicrosoftGraphAdministrativeUnit)`

SetAdministrativeUnits sets AdministrativeUnits field to given value.

### HasAdministrativeUnits

`func (o *MicrosoftGraphDirectory) HasAdministrativeUnits() bool`

HasAdministrativeUnits returns a boolean if a field has been set.

### GetDeletedItems

`func (o *MicrosoftGraphDirectory) GetDeletedItems() []MicrosoftGraphDirectoryObject`

GetDeletedItems returns the DeletedItems field if non-nil, zero value otherwise.

### GetDeletedItemsOk

`func (o *MicrosoftGraphDirectory) GetDeletedItemsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetDeletedItemsOk returns a tuple with the DeletedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedItems

`func (o *MicrosoftGraphDirectory) SetDeletedItems(v []MicrosoftGraphDirectoryObject)`

SetDeletedItems sets DeletedItems field to given value.

### HasDeletedItems

`func (o *MicrosoftGraphDirectory) HasDeletedItems() bool`

HasDeletedItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


