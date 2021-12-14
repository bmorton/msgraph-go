# MicrosoftGraphTrending

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** |  | [optional] 
**ResourceReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) | Reference properties of the trending document, such as the url and type of the document. | [optional] 
**ResourceVisualization** | Pointer to [**AnyOfmicrosoftGraphResourceVisualization**](anyOf&lt;microsoft.graph.resourceVisualization&gt;.md) | Properties that you can use to visualize the document in your experience. | [optional] 
**Weight** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value. | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) | Used for navigating to the trending document. | [optional] 

## Methods

### NewMicrosoftGraphTrending

`func NewMicrosoftGraphTrending() *MicrosoftGraphTrending`

NewMicrosoftGraphTrending instantiates a new MicrosoftGraphTrending object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTrendingWithDefaults

`func NewMicrosoftGraphTrendingWithDefaults() *MicrosoftGraphTrending`

NewMicrosoftGraphTrendingWithDefaults instantiates a new MicrosoftGraphTrending object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTrending) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTrending) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTrending) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTrending) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphTrending) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTrending) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTrending) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTrending) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphTrending) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphTrending) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetResourceReference

`func (o *MicrosoftGraphTrending) GetResourceReference() AnyOfmicrosoftGraphResourceReference`

GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.

### GetResourceReferenceOk

`func (o *MicrosoftGraphTrending) GetResourceReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool)`

GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceReference

`func (o *MicrosoftGraphTrending) SetResourceReference(v AnyOfmicrosoftGraphResourceReference)`

SetResourceReference sets ResourceReference field to given value.

### HasResourceReference

`func (o *MicrosoftGraphTrending) HasResourceReference() bool`

HasResourceReference returns a boolean if a field has been set.

### SetResourceReferenceNil

`func (o *MicrosoftGraphTrending) SetResourceReferenceNil(b bool)`

 SetResourceReferenceNil sets the value for ResourceReference to be an explicit nil

### UnsetResourceReference
`func (o *MicrosoftGraphTrending) UnsetResourceReference()`

UnsetResourceReference ensures that no value is present for ResourceReference, not even an explicit nil
### GetResourceVisualization

`func (o *MicrosoftGraphTrending) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization`

GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.

### GetResourceVisualizationOk

`func (o *MicrosoftGraphTrending) GetResourceVisualizationOk() (*AnyOfmicrosoftGraphResourceVisualization, bool)`

GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVisualization

`func (o *MicrosoftGraphTrending) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization)`

SetResourceVisualization sets ResourceVisualization field to given value.

### HasResourceVisualization

`func (o *MicrosoftGraphTrending) HasResourceVisualization() bool`

HasResourceVisualization returns a boolean if a field has been set.

### SetResourceVisualizationNil

`func (o *MicrosoftGraphTrending) SetResourceVisualizationNil(b bool)`

 SetResourceVisualizationNil sets the value for ResourceVisualization to be an explicit nil

### UnsetResourceVisualization
`func (o *MicrosoftGraphTrending) UnsetResourceVisualization()`

UnsetResourceVisualization ensures that no value is present for ResourceVisualization, not even an explicit nil
### GetWeight

`func (o *MicrosoftGraphTrending) GetWeight() AnyOfnumberstringstring`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MicrosoftGraphTrending) GetWeightOk() (*AnyOfnumberstringstring, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *MicrosoftGraphTrending) SetWeight(v AnyOfnumberstringstring)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *MicrosoftGraphTrending) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *MicrosoftGraphTrending) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *MicrosoftGraphTrending) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetResource

`func (o *MicrosoftGraphTrending) GetResource() AnyOfmicrosoftGraphEntity`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphTrending) GetResourceOk() (*AnyOfmicrosoftGraphEntity, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphTrending) SetResource(v AnyOfmicrosoftGraphEntity)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphTrending) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MicrosoftGraphTrending) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MicrosoftGraphTrending) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


