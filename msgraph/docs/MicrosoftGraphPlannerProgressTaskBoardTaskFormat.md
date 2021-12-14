# MicrosoftGraphPlannerProgressTaskBoardTaskFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**OrderHint** | Pointer to **NullableString** | Hint value used to order the task on the Progress view of the Task Board. The format is defined as outlined here. | [optional] 

## Methods

### NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat

`func NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() *MicrosoftGraphPlannerProgressTaskBoardTaskFormat`

NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat instantiates a new MicrosoftGraphPlannerProgressTaskBoardTaskFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPlannerProgressTaskBoardTaskFormatWithDefaults

`func NewMicrosoftGraphPlannerProgressTaskBoardTaskFormatWithDefaults() *MicrosoftGraphPlannerProgressTaskBoardTaskFormat`

NewMicrosoftGraphPlannerProgressTaskBoardTaskFormatWithDefaults instantiates a new MicrosoftGraphPlannerProgressTaskBoardTaskFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderHint

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) GetOrderHintOk() (*string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHint

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) SetOrderHint(v string)`

SetOrderHint sets OrderHint field to given value.

### HasOrderHint

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHintNil

`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) SetOrderHintNil(b bool)`

 SetOrderHintNil sets the value for OrderHint to be an explicit nil

### UnsetOrderHint
`func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) UnsetOrderHint()`

UnsetOrderHint ensures that no value is present for OrderHint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


