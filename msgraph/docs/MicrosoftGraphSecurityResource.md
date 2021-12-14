# MicrosoftGraphSecurityResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **NullableString** | Name of the resource that is related to current alert. Required. | [optional] 
**ResourceType** | Pointer to [**AnyOfmicrosoftGraphSecurityResourceType**](anyOf&lt;microsoft.graph.securityResourceType&gt;.md) | Represents type of security resources related to an alert. Possible values are: attacked, related. | [optional] 

## Methods

### NewMicrosoftGraphSecurityResource

`func NewMicrosoftGraphSecurityResource() *MicrosoftGraphSecurityResource`

NewMicrosoftGraphSecurityResource instantiates a new MicrosoftGraphSecurityResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSecurityResourceWithDefaults

`func NewMicrosoftGraphSecurityResourceWithDefaults() *MicrosoftGraphSecurityResource`

NewMicrosoftGraphSecurityResourceWithDefaults instantiates a new MicrosoftGraphSecurityResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *MicrosoftGraphSecurityResource) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphSecurityResource) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphSecurityResource) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphSecurityResource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MicrosoftGraphSecurityResource) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MicrosoftGraphSecurityResource) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil
### GetResourceType

`func (o *MicrosoftGraphSecurityResource) GetResourceType() AnyOfmicrosoftGraphSecurityResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *MicrosoftGraphSecurityResource) GetResourceTypeOk() (*AnyOfmicrosoftGraphSecurityResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *MicrosoftGraphSecurityResource) SetResourceType(v AnyOfmicrosoftGraphSecurityResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *MicrosoftGraphSecurityResource) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *MicrosoftGraphSecurityResource) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *MicrosoftGraphSecurityResource) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


