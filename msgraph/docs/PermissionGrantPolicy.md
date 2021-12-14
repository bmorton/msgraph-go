# PermissionGrantPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Excludes** | Pointer to [**[]MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md) | Condition sets which are excluded in this permission grant policy. Automatically expanded on GET. | [optional] 
**Includes** | Pointer to [**[]MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md) | Condition sets which are included in this permission grant policy. Automatically expanded on GET. | [optional] 

## Methods

### NewPermissionGrantPolicy

`func NewPermissionGrantPolicy() *PermissionGrantPolicy`

NewPermissionGrantPolicy instantiates a new PermissionGrantPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGrantPolicyWithDefaults

`func NewPermissionGrantPolicyWithDefaults() *PermissionGrantPolicy`

NewPermissionGrantPolicyWithDefaults instantiates a new PermissionGrantPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludes

`func (o *PermissionGrantPolicy) GetExcludes() []MicrosoftGraphPermissionGrantConditionSet`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *PermissionGrantPolicy) GetExcludesOk() (*[]MicrosoftGraphPermissionGrantConditionSet, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *PermissionGrantPolicy) SetExcludes(v []MicrosoftGraphPermissionGrantConditionSet)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *PermissionGrantPolicy) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### GetIncludes

`func (o *PermissionGrantPolicy) GetIncludes() []MicrosoftGraphPermissionGrantConditionSet`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *PermissionGrantPolicy) GetIncludesOk() (*[]MicrosoftGraphPermissionGrantConditionSet, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *PermissionGrantPolicy) SetIncludes(v []MicrosoftGraphPermissionGrantConditionSet)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *PermissionGrantPolicy) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


