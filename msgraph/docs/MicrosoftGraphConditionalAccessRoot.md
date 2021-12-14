# MicrosoftGraphConditionalAccessRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**NamedLocations** | Pointer to [**[]MicrosoftGraphNamedLocation**](MicrosoftGraphNamedLocation.md) | Read-only. Nullable. Returns a collection of the specified named locations. | [optional] 
**Policies** | Pointer to [**[]MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md) | Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies. | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessRoot

`func NewMicrosoftGraphConditionalAccessRoot() *MicrosoftGraphConditionalAccessRoot`

NewMicrosoftGraphConditionalAccessRoot instantiates a new MicrosoftGraphConditionalAccessRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessRootWithDefaults

`func NewMicrosoftGraphConditionalAccessRootWithDefaults() *MicrosoftGraphConditionalAccessRoot`

NewMicrosoftGraphConditionalAccessRootWithDefaults instantiates a new MicrosoftGraphConditionalAccessRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphConditionalAccessRoot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphConditionalAccessRoot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphConditionalAccessRoot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphConditionalAccessRoot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNamedLocations

`func (o *MicrosoftGraphConditionalAccessRoot) GetNamedLocations() []MicrosoftGraphNamedLocation`

GetNamedLocations returns the NamedLocations field if non-nil, zero value otherwise.

### GetNamedLocationsOk

`func (o *MicrosoftGraphConditionalAccessRoot) GetNamedLocationsOk() (*[]MicrosoftGraphNamedLocation, bool)`

GetNamedLocationsOk returns a tuple with the NamedLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedLocations

`func (o *MicrosoftGraphConditionalAccessRoot) SetNamedLocations(v []MicrosoftGraphNamedLocation)`

SetNamedLocations sets NamedLocations field to given value.

### HasNamedLocations

`func (o *MicrosoftGraphConditionalAccessRoot) HasNamedLocations() bool`

HasNamedLocations returns a boolean if a field has been set.

### GetPolicies

`func (o *MicrosoftGraphConditionalAccessRoot) GetPolicies() []MicrosoftGraphConditionalAccessPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *MicrosoftGraphConditionalAccessRoot) GetPoliciesOk() (*[]MicrosoftGraphConditionalAccessPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *MicrosoftGraphConditionalAccessRoot) SetPolicies(v []MicrosoftGraphConditionalAccessPolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *MicrosoftGraphConditionalAccessRoot) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


