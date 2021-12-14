# Directory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministrativeUnits** | Pointer to [**[]MicrosoftGraphAdministrativeUnit**](MicrosoftGraphAdministrativeUnit.md) | Conceptual container for user and group directory objects. | [optional] 
**DeletedItems** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Recently deleted items. Read-only. Nullable. | [optional] 

## Methods

### NewDirectory

`func NewDirectory() *Directory`

NewDirectory instantiates a new Directory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryWithDefaults

`func NewDirectoryWithDefaults() *Directory`

NewDirectoryWithDefaults instantiates a new Directory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrativeUnits

`func (o *Directory) GetAdministrativeUnits() []MicrosoftGraphAdministrativeUnit`

GetAdministrativeUnits returns the AdministrativeUnits field if non-nil, zero value otherwise.

### GetAdministrativeUnitsOk

`func (o *Directory) GetAdministrativeUnitsOk() (*[]MicrosoftGraphAdministrativeUnit, bool)`

GetAdministrativeUnitsOk returns a tuple with the AdministrativeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeUnits

`func (o *Directory) SetAdministrativeUnits(v []MicrosoftGraphAdministrativeUnit)`

SetAdministrativeUnits sets AdministrativeUnits field to given value.

### HasAdministrativeUnits

`func (o *Directory) HasAdministrativeUnits() bool`

HasAdministrativeUnits returns a boolean if a field has been set.

### GetDeletedItems

`func (o *Directory) GetDeletedItems() []MicrosoftGraphDirectoryObject`

GetDeletedItems returns the DeletedItems field if non-nil, zero value otherwise.

### GetDeletedItemsOk

`func (o *Directory) GetDeletedItemsOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetDeletedItemsOk returns a tuple with the DeletedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedItems

`func (o *Directory) SetDeletedItems(v []MicrosoftGraphDirectoryObject)`

SetDeletedItems sets DeletedItems field to given value.

### HasDeletedItems

`func (o *Directory) HasDeletedItems() bool`

HasDeletedItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


