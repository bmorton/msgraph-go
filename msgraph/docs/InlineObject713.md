# InlineObject713

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**PhysicalDeviceId** | Pointer to **NullableString** |  | [optional] 
**HasPhysicalDevice** | Pointer to **NullableBool** |  | [optional] [default to false]
**CertificateSigningRequest** | Pointer to [**MicrosoftGraphPrintCertificateSigningRequest**](MicrosoftGraphPrintCertificateSigningRequest.md) |  | [optional] 
**ConnectorId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInlineObject713

`func NewInlineObject713() *InlineObject713`

NewInlineObject713 instantiates a new InlineObject713 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject713WithDefaults

`func NewInlineObject713WithDefaults() *InlineObject713`

NewInlineObject713WithDefaults instantiates a new InlineObject713 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *InlineObject713) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineObject713) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineObject713) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InlineObject713) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetManufacturer

`func (o *InlineObject713) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InlineObject713) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InlineObject713) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InlineObject713) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *InlineObject713) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineObject713) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineObject713) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineObject713) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPhysicalDeviceId

`func (o *InlineObject713) GetPhysicalDeviceId() string`

GetPhysicalDeviceId returns the PhysicalDeviceId field if non-nil, zero value otherwise.

### GetPhysicalDeviceIdOk

`func (o *InlineObject713) GetPhysicalDeviceIdOk() (*string, bool)`

GetPhysicalDeviceIdOk returns a tuple with the PhysicalDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDeviceId

`func (o *InlineObject713) SetPhysicalDeviceId(v string)`

SetPhysicalDeviceId sets PhysicalDeviceId field to given value.

### HasPhysicalDeviceId

`func (o *InlineObject713) HasPhysicalDeviceId() bool`

HasPhysicalDeviceId returns a boolean if a field has been set.

### SetPhysicalDeviceIdNil

`func (o *InlineObject713) SetPhysicalDeviceIdNil(b bool)`

 SetPhysicalDeviceIdNil sets the value for PhysicalDeviceId to be an explicit nil

### UnsetPhysicalDeviceId
`func (o *InlineObject713) UnsetPhysicalDeviceId()`

UnsetPhysicalDeviceId ensures that no value is present for PhysicalDeviceId, not even an explicit nil
### GetHasPhysicalDevice

`func (o *InlineObject713) GetHasPhysicalDevice() bool`

GetHasPhysicalDevice returns the HasPhysicalDevice field if non-nil, zero value otherwise.

### GetHasPhysicalDeviceOk

`func (o *InlineObject713) GetHasPhysicalDeviceOk() (*bool, bool)`

GetHasPhysicalDeviceOk returns a tuple with the HasPhysicalDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPhysicalDevice

`func (o *InlineObject713) SetHasPhysicalDevice(v bool)`

SetHasPhysicalDevice sets HasPhysicalDevice field to given value.

### HasHasPhysicalDevice

`func (o *InlineObject713) HasHasPhysicalDevice() bool`

HasHasPhysicalDevice returns a boolean if a field has been set.

### SetHasPhysicalDeviceNil

`func (o *InlineObject713) SetHasPhysicalDeviceNil(b bool)`

 SetHasPhysicalDeviceNil sets the value for HasPhysicalDevice to be an explicit nil

### UnsetHasPhysicalDevice
`func (o *InlineObject713) UnsetHasPhysicalDevice()`

UnsetHasPhysicalDevice ensures that no value is present for HasPhysicalDevice, not even an explicit nil
### GetCertificateSigningRequest

`func (o *InlineObject713) GetCertificateSigningRequest() MicrosoftGraphPrintCertificateSigningRequest`

GetCertificateSigningRequest returns the CertificateSigningRequest field if non-nil, zero value otherwise.

### GetCertificateSigningRequestOk

`func (o *InlineObject713) GetCertificateSigningRequestOk() (*MicrosoftGraphPrintCertificateSigningRequest, bool)`

GetCertificateSigningRequestOk returns a tuple with the CertificateSigningRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSigningRequest

`func (o *InlineObject713) SetCertificateSigningRequest(v MicrosoftGraphPrintCertificateSigningRequest)`

SetCertificateSigningRequest sets CertificateSigningRequest field to given value.

### HasCertificateSigningRequest

`func (o *InlineObject713) HasCertificateSigningRequest() bool`

HasCertificateSigningRequest returns a boolean if a field has been set.

### GetConnectorId

`func (o *InlineObject713) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *InlineObject713) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *InlineObject713) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *InlineObject713) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### SetConnectorIdNil

`func (o *InlineObject713) SetConnectorIdNil(b bool)`

 SetConnectorIdNil sets the value for ConnectorId to be an explicit nil

### UnsetConnectorId
`func (o *InlineObject713) UnsetConnectorId()`

UnsetConnectorId ensures that no value is present for ConnectorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


