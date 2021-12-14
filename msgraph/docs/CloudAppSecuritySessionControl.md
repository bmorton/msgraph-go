# CloudAppSecuritySessionControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudAppSecurityType** | Pointer to [**AnyOfmicrosoftGraphCloudAppSecuritySessionControlType**](anyOf&lt;microsoft.graph.cloudAppSecuritySessionControlType&gt;.md) | Possible values are: mcasConfigured, monitorOnly, blockDownloads, unknownFutureValue. For more information, see Deploy Conditional Access App Control for featured apps. | [optional] 

## Methods

### NewCloudAppSecuritySessionControl

`func NewCloudAppSecuritySessionControl() *CloudAppSecuritySessionControl`

NewCloudAppSecuritySessionControl instantiates a new CloudAppSecuritySessionControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAppSecuritySessionControlWithDefaults

`func NewCloudAppSecuritySessionControlWithDefaults() *CloudAppSecuritySessionControl`

NewCloudAppSecuritySessionControlWithDefaults instantiates a new CloudAppSecuritySessionControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudAppSecurityType

`func (o *CloudAppSecuritySessionControl) GetCloudAppSecurityType() AnyOfmicrosoftGraphCloudAppSecuritySessionControlType`

GetCloudAppSecurityType returns the CloudAppSecurityType field if non-nil, zero value otherwise.

### GetCloudAppSecurityTypeOk

`func (o *CloudAppSecuritySessionControl) GetCloudAppSecurityTypeOk() (*AnyOfmicrosoftGraphCloudAppSecuritySessionControlType, bool)`

GetCloudAppSecurityTypeOk returns a tuple with the CloudAppSecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAppSecurityType

`func (o *CloudAppSecuritySessionControl) SetCloudAppSecurityType(v AnyOfmicrosoftGraphCloudAppSecuritySessionControlType)`

SetCloudAppSecurityType sets CloudAppSecurityType field to given value.

### HasCloudAppSecurityType

`func (o *CloudAppSecuritySessionControl) HasCloudAppSecurityType() bool`

HasCloudAppSecurityType returns a boolean if a field has been set.

### SetCloudAppSecurityTypeNil

`func (o *CloudAppSecuritySessionControl) SetCloudAppSecurityTypeNil(b bool)`

 SetCloudAppSecurityTypeNil sets the value for CloudAppSecurityType to be an explicit nil

### UnsetCloudAppSecurityType
`func (o *CloudAppSecuritySessionControl) UnsetCloudAppSecurityType()`

UnsetCloudAppSecurityType ensures that no value is present for CloudAppSecurityType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


