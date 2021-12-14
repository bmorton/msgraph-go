# MicrosoftGraphSharedInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**LastShared** | Pointer to [**AnyOfmicrosoftGraphSharingDetail**](anyOf&lt;microsoft.graph.sharingDetail&gt;.md) | Details about the shared item. Read only. | [optional] 
**ResourceReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) | Reference properties of the shared document, such as the url and type of the document. Read-only | [optional] 
**ResourceVisualization** | Pointer to [**AnyOfmicrosoftGraphResourceVisualization**](anyOf&lt;microsoft.graph.resourceVisualization&gt;.md) | Properties that you can use to visualize the document in your experience. Read-only | [optional] 
**SharingHistory** | Pointer to [**[]AnyOfmicrosoftGraphSharingDetail**](AnyOfmicrosoftGraphSharingDetail.md) |  | [optional] 
**LastSharedMethod** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) |  | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) | Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem. | [optional] 

## Methods

### NewMicrosoftGraphSharedInsight

`func NewMicrosoftGraphSharedInsight() *MicrosoftGraphSharedInsight`

NewMicrosoftGraphSharedInsight instantiates a new MicrosoftGraphSharedInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSharedInsightWithDefaults

`func NewMicrosoftGraphSharedInsightWithDefaults() *MicrosoftGraphSharedInsight`

NewMicrosoftGraphSharedInsightWithDefaults instantiates a new MicrosoftGraphSharedInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSharedInsight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSharedInsight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSharedInsight) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSharedInsight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastShared

`func (o *MicrosoftGraphSharedInsight) GetLastShared() AnyOfmicrosoftGraphSharingDetail`

GetLastShared returns the LastShared field if non-nil, zero value otherwise.

### GetLastSharedOk

`func (o *MicrosoftGraphSharedInsight) GetLastSharedOk() (*AnyOfmicrosoftGraphSharingDetail, bool)`

GetLastSharedOk returns a tuple with the LastShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShared

`func (o *MicrosoftGraphSharedInsight) SetLastShared(v AnyOfmicrosoftGraphSharingDetail)`

SetLastShared sets LastShared field to given value.

### HasLastShared

`func (o *MicrosoftGraphSharedInsight) HasLastShared() bool`

HasLastShared returns a boolean if a field has been set.

### SetLastSharedNil

`func (o *MicrosoftGraphSharedInsight) SetLastSharedNil(b bool)`

 SetLastSharedNil sets the value for LastShared to be an explicit nil

### UnsetLastShared
`func (o *MicrosoftGraphSharedInsight) UnsetLastShared()`

UnsetLastShared ensures that no value is present for LastShared, not even an explicit nil
### GetResourceReference

`func (o *MicrosoftGraphSharedInsight) GetResourceReference() AnyOfmicrosoftGraphResourceReference`

GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.

### GetResourceReferenceOk

`func (o *MicrosoftGraphSharedInsight) GetResourceReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool)`

GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceReference

`func (o *MicrosoftGraphSharedInsight) SetResourceReference(v AnyOfmicrosoftGraphResourceReference)`

SetResourceReference sets ResourceReference field to given value.

### HasResourceReference

`func (o *MicrosoftGraphSharedInsight) HasResourceReference() bool`

HasResourceReference returns a boolean if a field has been set.

### SetResourceReferenceNil

`func (o *MicrosoftGraphSharedInsight) SetResourceReferenceNil(b bool)`

 SetResourceReferenceNil sets the value for ResourceReference to be an explicit nil

### UnsetResourceReference
`func (o *MicrosoftGraphSharedInsight) UnsetResourceReference()`

UnsetResourceReference ensures that no value is present for ResourceReference, not even an explicit nil
### GetResourceVisualization

`func (o *MicrosoftGraphSharedInsight) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization`

GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.

### GetResourceVisualizationOk

`func (o *MicrosoftGraphSharedInsight) GetResourceVisualizationOk() (*AnyOfmicrosoftGraphResourceVisualization, bool)`

GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVisualization

`func (o *MicrosoftGraphSharedInsight) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization)`

SetResourceVisualization sets ResourceVisualization field to given value.

### HasResourceVisualization

`func (o *MicrosoftGraphSharedInsight) HasResourceVisualization() bool`

HasResourceVisualization returns a boolean if a field has been set.

### SetResourceVisualizationNil

`func (o *MicrosoftGraphSharedInsight) SetResourceVisualizationNil(b bool)`

 SetResourceVisualizationNil sets the value for ResourceVisualization to be an explicit nil

### UnsetResourceVisualization
`func (o *MicrosoftGraphSharedInsight) UnsetResourceVisualization()`

UnsetResourceVisualization ensures that no value is present for ResourceVisualization, not even an explicit nil
### GetSharingHistory

`func (o *MicrosoftGraphSharedInsight) GetSharingHistory() []*AnyOfmicrosoftGraphSharingDetail`

GetSharingHistory returns the SharingHistory field if non-nil, zero value otherwise.

### GetSharingHistoryOk

`func (o *MicrosoftGraphSharedInsight) GetSharingHistoryOk() (*[]*AnyOfmicrosoftGraphSharingDetail, bool)`

GetSharingHistoryOk returns a tuple with the SharingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingHistory

`func (o *MicrosoftGraphSharedInsight) SetSharingHistory(v []*AnyOfmicrosoftGraphSharingDetail)`

SetSharingHistory sets SharingHistory field to given value.

### HasSharingHistory

`func (o *MicrosoftGraphSharedInsight) HasSharingHistory() bool`

HasSharingHistory returns a boolean if a field has been set.

### GetLastSharedMethod

`func (o *MicrosoftGraphSharedInsight) GetLastSharedMethod() AnyOfmicrosoftGraphEntity`

GetLastSharedMethod returns the LastSharedMethod field if non-nil, zero value otherwise.

### GetLastSharedMethodOk

`func (o *MicrosoftGraphSharedInsight) GetLastSharedMethodOk() (*AnyOfmicrosoftGraphEntity, bool)`

GetLastSharedMethodOk returns a tuple with the LastSharedMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSharedMethod

`func (o *MicrosoftGraphSharedInsight) SetLastSharedMethod(v AnyOfmicrosoftGraphEntity)`

SetLastSharedMethod sets LastSharedMethod field to given value.

### HasLastSharedMethod

`func (o *MicrosoftGraphSharedInsight) HasLastSharedMethod() bool`

HasLastSharedMethod returns a boolean if a field has been set.

### SetLastSharedMethodNil

`func (o *MicrosoftGraphSharedInsight) SetLastSharedMethodNil(b bool)`

 SetLastSharedMethodNil sets the value for LastSharedMethod to be an explicit nil

### UnsetLastSharedMethod
`func (o *MicrosoftGraphSharedInsight) UnsetLastSharedMethod()`

UnsetLastSharedMethod ensures that no value is present for LastSharedMethod, not even an explicit nil
### GetResource

`func (o *MicrosoftGraphSharedInsight) GetResource() AnyOfmicrosoftGraphEntity`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphSharedInsight) GetResourceOk() (*AnyOfmicrosoftGraphEntity, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphSharedInsight) SetResource(v AnyOfmicrosoftGraphEntity)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphSharedInsight) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MicrosoftGraphSharedInsight) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MicrosoftGraphSharedInsight) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


