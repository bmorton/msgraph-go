# ConditionalAccessRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamedLocations** | Pointer to [**[]MicrosoftGraphNamedLocation**](MicrosoftGraphNamedLocation.md) | Read-only. Nullable. Returns a collection of the specified named locations. | [optional] 
**Policies** | Pointer to [**[]MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md) | Read-only. Nullable. Returns a collection of the specified Conditional Access (CA) policies. | [optional] 

## Methods

### NewConditionalAccessRoot

`func NewConditionalAccessRoot() *ConditionalAccessRoot`

NewConditionalAccessRoot instantiates a new ConditionalAccessRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalAccessRootWithDefaults

`func NewConditionalAccessRootWithDefaults() *ConditionalAccessRoot`

NewConditionalAccessRootWithDefaults instantiates a new ConditionalAccessRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamedLocations

`func (o *ConditionalAccessRoot) GetNamedLocations() []MicrosoftGraphNamedLocation`

GetNamedLocations returns the NamedLocations field if non-nil, zero value otherwise.

### GetNamedLocationsOk

`func (o *ConditionalAccessRoot) GetNamedLocationsOk() (*[]MicrosoftGraphNamedLocation, bool)`

GetNamedLocationsOk returns a tuple with the NamedLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedLocations

`func (o *ConditionalAccessRoot) SetNamedLocations(v []MicrosoftGraphNamedLocation)`

SetNamedLocations sets NamedLocations field to given value.

### HasNamedLocations

`func (o *ConditionalAccessRoot) HasNamedLocations() bool`

HasNamedLocations returns a boolean if a field has been set.

### GetPolicies

`func (o *ConditionalAccessRoot) GetPolicies() []MicrosoftGraphConditionalAccessPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ConditionalAccessRoot) GetPoliciesOk() (*[]MicrosoftGraphConditionalAccessPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ConditionalAccessRoot) SetPolicies(v []MicrosoftGraphConditionalAccessPolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ConditionalAccessRoot) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


