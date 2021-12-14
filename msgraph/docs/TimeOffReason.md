# TimeOffReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The name of the timeOffReason. Required. | [optional] 
**IconType** | Pointer to [**AnyOfmicrosoftGraphTimeOffReasonIconType**](anyOf&lt;microsoft.graph.timeOffReasonIconType&gt;.md) | Supported icon types: none; car; calendar; running; plane; firstAid; doctor; notWorking; clock; juryDuty; globe; cup; phone; weather; umbrella; piggyBank; dog; cake; trafficCone; pin; sunny. Required. | [optional] 
**IsActive** | Pointer to **NullableBool** | Indicates whether the timeOffReason can be used when creating new entities or updating existing ones. Required. | [optional] 

## Methods

### NewTimeOffReason

`func NewTimeOffReason() *TimeOffReason`

NewTimeOffReason instantiates a new TimeOffReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffReasonWithDefaults

`func NewTimeOffReasonWithDefaults() *TimeOffReason`

NewTimeOffReasonWithDefaults instantiates a new TimeOffReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TimeOffReason) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TimeOffReason) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TimeOffReason) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TimeOffReason) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *TimeOffReason) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TimeOffReason) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIconType

`func (o *TimeOffReason) GetIconType() AnyOfmicrosoftGraphTimeOffReasonIconType`

GetIconType returns the IconType field if non-nil, zero value otherwise.

### GetIconTypeOk

`func (o *TimeOffReason) GetIconTypeOk() (*AnyOfmicrosoftGraphTimeOffReasonIconType, bool)`

GetIconTypeOk returns a tuple with the IconType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconType

`func (o *TimeOffReason) SetIconType(v AnyOfmicrosoftGraphTimeOffReasonIconType)`

SetIconType sets IconType field to given value.

### HasIconType

`func (o *TimeOffReason) HasIconType() bool`

HasIconType returns a boolean if a field has been set.

### SetIconTypeNil

`func (o *TimeOffReason) SetIconTypeNil(b bool)`

 SetIconTypeNil sets the value for IconType to be an explicit nil

### UnsetIconType
`func (o *TimeOffReason) UnsetIconType()`

UnsetIconType ensures that no value is present for IconType, not even an explicit nil
### GetIsActive

`func (o *TimeOffReason) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TimeOffReason) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TimeOffReason) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TimeOffReason) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *TimeOffReason) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *TimeOffReason) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


