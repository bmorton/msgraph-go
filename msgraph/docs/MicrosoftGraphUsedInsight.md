# MicrosoftGraphUsedInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**LastUsed** | Pointer to [**AnyOfmicrosoftGraphUsageDetails**](anyOf&lt;microsoft.graph.usageDetails&gt;.md) | Information about when the item was last viewed or modified by the user. Read only. | [optional] 
**ResourceReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) | Reference properties of the used document, such as the url and type of the document. Read-only | [optional] 
**ResourceVisualization** | Pointer to [**AnyOfmicrosoftGraphResourceVisualization**](anyOf&lt;microsoft.graph.resourceVisualization&gt;.md) | Properties that you can use to visualize the document in your experience. Read-only | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) | Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem. | [optional] 

## Methods

### NewMicrosoftGraphUsedInsight

`func NewMicrosoftGraphUsedInsight() *MicrosoftGraphUsedInsight`

NewMicrosoftGraphUsedInsight instantiates a new MicrosoftGraphUsedInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUsedInsightWithDefaults

`func NewMicrosoftGraphUsedInsightWithDefaults() *MicrosoftGraphUsedInsight`

NewMicrosoftGraphUsedInsightWithDefaults instantiates a new MicrosoftGraphUsedInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUsedInsight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUsedInsight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUsedInsight) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUsedInsight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsed

`func (o *MicrosoftGraphUsedInsight) GetLastUsed() AnyOfmicrosoftGraphUsageDetails`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *MicrosoftGraphUsedInsight) GetLastUsedOk() (*AnyOfmicrosoftGraphUsageDetails, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *MicrosoftGraphUsedInsight) SetLastUsed(v AnyOfmicrosoftGraphUsageDetails)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *MicrosoftGraphUsedInsight) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *MicrosoftGraphUsedInsight) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *MicrosoftGraphUsedInsight) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetResourceReference

`func (o *MicrosoftGraphUsedInsight) GetResourceReference() AnyOfmicrosoftGraphResourceReference`

GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.

### GetResourceReferenceOk

`func (o *MicrosoftGraphUsedInsight) GetResourceReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool)`

GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceReference

`func (o *MicrosoftGraphUsedInsight) SetResourceReference(v AnyOfmicrosoftGraphResourceReference)`

SetResourceReference sets ResourceReference field to given value.

### HasResourceReference

`func (o *MicrosoftGraphUsedInsight) HasResourceReference() bool`

HasResourceReference returns a boolean if a field has been set.

### SetResourceReferenceNil

`func (o *MicrosoftGraphUsedInsight) SetResourceReferenceNil(b bool)`

 SetResourceReferenceNil sets the value for ResourceReference to be an explicit nil

### UnsetResourceReference
`func (o *MicrosoftGraphUsedInsight) UnsetResourceReference()`

UnsetResourceReference ensures that no value is present for ResourceReference, not even an explicit nil
### GetResourceVisualization

`func (o *MicrosoftGraphUsedInsight) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization`

GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.

### GetResourceVisualizationOk

`func (o *MicrosoftGraphUsedInsight) GetResourceVisualizationOk() (*AnyOfmicrosoftGraphResourceVisualization, bool)`

GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVisualization

`func (o *MicrosoftGraphUsedInsight) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization)`

SetResourceVisualization sets ResourceVisualization field to given value.

### HasResourceVisualization

`func (o *MicrosoftGraphUsedInsight) HasResourceVisualization() bool`

HasResourceVisualization returns a boolean if a field has been set.

### SetResourceVisualizationNil

`func (o *MicrosoftGraphUsedInsight) SetResourceVisualizationNil(b bool)`

 SetResourceVisualizationNil sets the value for ResourceVisualization to be an explicit nil

### UnsetResourceVisualization
`func (o *MicrosoftGraphUsedInsight) UnsetResourceVisualization()`

UnsetResourceVisualization ensures that no value is present for ResourceVisualization, not even an explicit nil
### GetResource

`func (o *MicrosoftGraphUsedInsight) GetResource() AnyOfmicrosoftGraphEntity`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphUsedInsight) GetResourceOk() (*AnyOfmicrosoftGraphEntity, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphUsedInsight) SetResource(v AnyOfmicrosoftGraphEntity)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphUsedInsight) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MicrosoftGraphUsedInsight) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MicrosoftGraphUsedInsight) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


