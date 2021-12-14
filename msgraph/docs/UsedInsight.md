# UsedInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUsed** | Pointer to [**AnyOfmicrosoftGraphUsageDetails**](anyOf&lt;microsoft.graph.usageDetails&gt;.md) | Information about when the item was last viewed or modified by the user. Read only. | [optional] 
**ResourceReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) | Reference properties of the used document, such as the url and type of the document. Read-only | [optional] 
**ResourceVisualization** | Pointer to [**AnyOfmicrosoftGraphResourceVisualization**](anyOf&lt;microsoft.graph.resourceVisualization&gt;.md) | Properties that you can use to visualize the document in your experience. Read-only | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) | Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem. | [optional] 

## Methods

### NewUsedInsight

`func NewUsedInsight() *UsedInsight`

NewUsedInsight instantiates a new UsedInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsedInsightWithDefaults

`func NewUsedInsightWithDefaults() *UsedInsight`

NewUsedInsightWithDefaults instantiates a new UsedInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUsed

`func (o *UsedInsight) GetLastUsed() AnyOfmicrosoftGraphUsageDetails`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *UsedInsight) GetLastUsedOk() (*AnyOfmicrosoftGraphUsageDetails, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *UsedInsight) SetLastUsed(v AnyOfmicrosoftGraphUsageDetails)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *UsedInsight) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *UsedInsight) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *UsedInsight) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetResourceReference

`func (o *UsedInsight) GetResourceReference() AnyOfmicrosoftGraphResourceReference`

GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.

### GetResourceReferenceOk

`func (o *UsedInsight) GetResourceReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool)`

GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceReference

`func (o *UsedInsight) SetResourceReference(v AnyOfmicrosoftGraphResourceReference)`

SetResourceReference sets ResourceReference field to given value.

### HasResourceReference

`func (o *UsedInsight) HasResourceReference() bool`

HasResourceReference returns a boolean if a field has been set.

### SetResourceReferenceNil

`func (o *UsedInsight) SetResourceReferenceNil(b bool)`

 SetResourceReferenceNil sets the value for ResourceReference to be an explicit nil

### UnsetResourceReference
`func (o *UsedInsight) UnsetResourceReference()`

UnsetResourceReference ensures that no value is present for ResourceReference, not even an explicit nil
### GetResourceVisualization

`func (o *UsedInsight) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization`

GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.

### GetResourceVisualizationOk

`func (o *UsedInsight) GetResourceVisualizationOk() (*AnyOfmicrosoftGraphResourceVisualization, bool)`

GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVisualization

`func (o *UsedInsight) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization)`

SetResourceVisualization sets ResourceVisualization field to given value.

### HasResourceVisualization

`func (o *UsedInsight) HasResourceVisualization() bool`

HasResourceVisualization returns a boolean if a field has been set.

### SetResourceVisualizationNil

`func (o *UsedInsight) SetResourceVisualizationNil(b bool)`

 SetResourceVisualizationNil sets the value for ResourceVisualization to be an explicit nil

### UnsetResourceVisualization
`func (o *UsedInsight) UnsetResourceVisualization()`

UnsetResourceVisualization ensures that no value is present for ResourceVisualization, not even an explicit nil
### GetResource

`func (o *UsedInsight) GetResource() AnyOfmicrosoftGraphEntity`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *UsedInsight) GetResourceOk() (*AnyOfmicrosoftGraphEntity, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *UsedInsight) SetResource(v AnyOfmicrosoftGraphEntity)`

SetResource sets Resource field to given value.

### HasResource

`func (o *UsedInsight) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *UsedInsight) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *UsedInsight) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


