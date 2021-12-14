# MicrosoftGraphEmployeeOrgData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostCenter** | Pointer to **NullableString** | The cost center associated with the user. Returned only on $select. Supports $filter. | [optional] 
**Division** | Pointer to **NullableString** | The name of the division in which the user works. Returned only on $select. Supports $filter. | [optional] 

## Methods

### NewMicrosoftGraphEmployeeOrgData

`func NewMicrosoftGraphEmployeeOrgData() *MicrosoftGraphEmployeeOrgData`

NewMicrosoftGraphEmployeeOrgData instantiates a new MicrosoftGraphEmployeeOrgData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEmployeeOrgDataWithDefaults

`func NewMicrosoftGraphEmployeeOrgDataWithDefaults() *MicrosoftGraphEmployeeOrgData`

NewMicrosoftGraphEmployeeOrgDataWithDefaults instantiates a new MicrosoftGraphEmployeeOrgData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostCenter

`func (o *MicrosoftGraphEmployeeOrgData) GetCostCenter() string`

GetCostCenter returns the CostCenter field if non-nil, zero value otherwise.

### GetCostCenterOk

`func (o *MicrosoftGraphEmployeeOrgData) GetCostCenterOk() (*string, bool)`

GetCostCenterOk returns a tuple with the CostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCenter

`func (o *MicrosoftGraphEmployeeOrgData) SetCostCenter(v string)`

SetCostCenter sets CostCenter field to given value.

### HasCostCenter

`func (o *MicrosoftGraphEmployeeOrgData) HasCostCenter() bool`

HasCostCenter returns a boolean if a field has been set.

### SetCostCenterNil

`func (o *MicrosoftGraphEmployeeOrgData) SetCostCenterNil(b bool)`

 SetCostCenterNil sets the value for CostCenter to be an explicit nil

### UnsetCostCenter
`func (o *MicrosoftGraphEmployeeOrgData) UnsetCostCenter()`

UnsetCostCenter ensures that no value is present for CostCenter, not even an explicit nil
### GetDivision

`func (o *MicrosoftGraphEmployeeOrgData) GetDivision() string`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *MicrosoftGraphEmployeeOrgData) GetDivisionOk() (*string, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *MicrosoftGraphEmployeeOrgData) SetDivision(v string)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *MicrosoftGraphEmployeeOrgData) HasDivision() bool`

HasDivision returns a boolean if a field has been set.

### SetDivisionNil

`func (o *MicrosoftGraphEmployeeOrgData) SetDivisionNil(b bool)`

 SetDivisionNil sets the value for Division to be an explicit nil

### UnsetDivision
`func (o *MicrosoftGraphEmployeeOrgData) UnsetDivision()`

UnsetDivision ensures that no value is present for Division, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


