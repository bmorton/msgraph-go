# MicrosoftGraphPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CameraMake** | Pointer to **NullableString** | Camera manufacturer. Read-only. | [optional] 
**CameraModel** | Pointer to **NullableString** | Camera model. Read-only. | [optional] 
**ExposureDenominator** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The denominator for the exposure time fraction from the camera. Read-only. | [optional] 
**ExposureNumerator** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The numerator for the exposure time fraction from the camera. Read-only. | [optional] 
**FNumber** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The F-stop value from the camera. Read-only. | [optional] 
**FocalLength** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The focal length from the camera. Read-only. | [optional] 
**Iso** | Pointer to **NullableInt32** | The ISO value from the camera. Read-only. | [optional] 
**Orientation** | Pointer to **NullableInt32** | The orientation value from the camera. Writable on OneDrive Personal. | [optional] 
**TakenDateTime** | Pointer to **NullableTime** | Represents the date and time the photo was taken. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphPhoto

`func NewMicrosoftGraphPhoto() *MicrosoftGraphPhoto`

NewMicrosoftGraphPhoto instantiates a new MicrosoftGraphPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPhotoWithDefaults

`func NewMicrosoftGraphPhotoWithDefaults() *MicrosoftGraphPhoto`

NewMicrosoftGraphPhotoWithDefaults instantiates a new MicrosoftGraphPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCameraMake

`func (o *MicrosoftGraphPhoto) GetCameraMake() string`

GetCameraMake returns the CameraMake field if non-nil, zero value otherwise.

### GetCameraMakeOk

`func (o *MicrosoftGraphPhoto) GetCameraMakeOk() (*string, bool)`

GetCameraMakeOk returns a tuple with the CameraMake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraMake

`func (o *MicrosoftGraphPhoto) SetCameraMake(v string)`

SetCameraMake sets CameraMake field to given value.

### HasCameraMake

`func (o *MicrosoftGraphPhoto) HasCameraMake() bool`

HasCameraMake returns a boolean if a field has been set.

### SetCameraMakeNil

`func (o *MicrosoftGraphPhoto) SetCameraMakeNil(b bool)`

 SetCameraMakeNil sets the value for CameraMake to be an explicit nil

### UnsetCameraMake
`func (o *MicrosoftGraphPhoto) UnsetCameraMake()`

UnsetCameraMake ensures that no value is present for CameraMake, not even an explicit nil
### GetCameraModel

`func (o *MicrosoftGraphPhoto) GetCameraModel() string`

GetCameraModel returns the CameraModel field if non-nil, zero value otherwise.

### GetCameraModelOk

`func (o *MicrosoftGraphPhoto) GetCameraModelOk() (*string, bool)`

GetCameraModelOk returns a tuple with the CameraModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraModel

`func (o *MicrosoftGraphPhoto) SetCameraModel(v string)`

SetCameraModel sets CameraModel field to given value.

### HasCameraModel

`func (o *MicrosoftGraphPhoto) HasCameraModel() bool`

HasCameraModel returns a boolean if a field has been set.

### SetCameraModelNil

`func (o *MicrosoftGraphPhoto) SetCameraModelNil(b bool)`

 SetCameraModelNil sets the value for CameraModel to be an explicit nil

### UnsetCameraModel
`func (o *MicrosoftGraphPhoto) UnsetCameraModel()`

UnsetCameraModel ensures that no value is present for CameraModel, not even an explicit nil
### GetExposureDenominator

`func (o *MicrosoftGraphPhoto) GetExposureDenominator() AnyOfnumberstringstring`

GetExposureDenominator returns the ExposureDenominator field if non-nil, zero value otherwise.

### GetExposureDenominatorOk

`func (o *MicrosoftGraphPhoto) GetExposureDenominatorOk() (*AnyOfnumberstringstring, bool)`

GetExposureDenominatorOk returns a tuple with the ExposureDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureDenominator

`func (o *MicrosoftGraphPhoto) SetExposureDenominator(v AnyOfnumberstringstring)`

SetExposureDenominator sets ExposureDenominator field to given value.

### HasExposureDenominator

`func (o *MicrosoftGraphPhoto) HasExposureDenominator() bool`

HasExposureDenominator returns a boolean if a field has been set.

### SetExposureDenominatorNil

`func (o *MicrosoftGraphPhoto) SetExposureDenominatorNil(b bool)`

 SetExposureDenominatorNil sets the value for ExposureDenominator to be an explicit nil

### UnsetExposureDenominator
`func (o *MicrosoftGraphPhoto) UnsetExposureDenominator()`

UnsetExposureDenominator ensures that no value is present for ExposureDenominator, not even an explicit nil
### GetExposureNumerator

`func (o *MicrosoftGraphPhoto) GetExposureNumerator() AnyOfnumberstringstring`

GetExposureNumerator returns the ExposureNumerator field if non-nil, zero value otherwise.

### GetExposureNumeratorOk

`func (o *MicrosoftGraphPhoto) GetExposureNumeratorOk() (*AnyOfnumberstringstring, bool)`

GetExposureNumeratorOk returns a tuple with the ExposureNumerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureNumerator

`func (o *MicrosoftGraphPhoto) SetExposureNumerator(v AnyOfnumberstringstring)`

SetExposureNumerator sets ExposureNumerator field to given value.

### HasExposureNumerator

`func (o *MicrosoftGraphPhoto) HasExposureNumerator() bool`

HasExposureNumerator returns a boolean if a field has been set.

### SetExposureNumeratorNil

`func (o *MicrosoftGraphPhoto) SetExposureNumeratorNil(b bool)`

 SetExposureNumeratorNil sets the value for ExposureNumerator to be an explicit nil

### UnsetExposureNumerator
`func (o *MicrosoftGraphPhoto) UnsetExposureNumerator()`

UnsetExposureNumerator ensures that no value is present for ExposureNumerator, not even an explicit nil
### GetFNumber

`func (o *MicrosoftGraphPhoto) GetFNumber() AnyOfnumberstringstring`

GetFNumber returns the FNumber field if non-nil, zero value otherwise.

### GetFNumberOk

`func (o *MicrosoftGraphPhoto) GetFNumberOk() (*AnyOfnumberstringstring, bool)`

GetFNumberOk returns a tuple with the FNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFNumber

`func (o *MicrosoftGraphPhoto) SetFNumber(v AnyOfnumberstringstring)`

SetFNumber sets FNumber field to given value.

### HasFNumber

`func (o *MicrosoftGraphPhoto) HasFNumber() bool`

HasFNumber returns a boolean if a field has been set.

### SetFNumberNil

`func (o *MicrosoftGraphPhoto) SetFNumberNil(b bool)`

 SetFNumberNil sets the value for FNumber to be an explicit nil

### UnsetFNumber
`func (o *MicrosoftGraphPhoto) UnsetFNumber()`

UnsetFNumber ensures that no value is present for FNumber, not even an explicit nil
### GetFocalLength

`func (o *MicrosoftGraphPhoto) GetFocalLength() AnyOfnumberstringstring`

GetFocalLength returns the FocalLength field if non-nil, zero value otherwise.

### GetFocalLengthOk

`func (o *MicrosoftGraphPhoto) GetFocalLengthOk() (*AnyOfnumberstringstring, bool)`

GetFocalLengthOk returns a tuple with the FocalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocalLength

`func (o *MicrosoftGraphPhoto) SetFocalLength(v AnyOfnumberstringstring)`

SetFocalLength sets FocalLength field to given value.

### HasFocalLength

`func (o *MicrosoftGraphPhoto) HasFocalLength() bool`

HasFocalLength returns a boolean if a field has been set.

### SetFocalLengthNil

`func (o *MicrosoftGraphPhoto) SetFocalLengthNil(b bool)`

 SetFocalLengthNil sets the value for FocalLength to be an explicit nil

### UnsetFocalLength
`func (o *MicrosoftGraphPhoto) UnsetFocalLength()`

UnsetFocalLength ensures that no value is present for FocalLength, not even an explicit nil
### GetIso

`func (o *MicrosoftGraphPhoto) GetIso() int32`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *MicrosoftGraphPhoto) GetIsoOk() (*int32, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *MicrosoftGraphPhoto) SetIso(v int32)`

SetIso sets Iso field to given value.

### HasIso

`func (o *MicrosoftGraphPhoto) HasIso() bool`

HasIso returns a boolean if a field has been set.

### SetIsoNil

`func (o *MicrosoftGraphPhoto) SetIsoNil(b bool)`

 SetIsoNil sets the value for Iso to be an explicit nil

### UnsetIso
`func (o *MicrosoftGraphPhoto) UnsetIso()`

UnsetIso ensures that no value is present for Iso, not even an explicit nil
### GetOrientation

`func (o *MicrosoftGraphPhoto) GetOrientation() int32`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *MicrosoftGraphPhoto) GetOrientationOk() (*int32, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *MicrosoftGraphPhoto) SetOrientation(v int32)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *MicrosoftGraphPhoto) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### SetOrientationNil

`func (o *MicrosoftGraphPhoto) SetOrientationNil(b bool)`

 SetOrientationNil sets the value for Orientation to be an explicit nil

### UnsetOrientation
`func (o *MicrosoftGraphPhoto) UnsetOrientation()`

UnsetOrientation ensures that no value is present for Orientation, not even an explicit nil
### GetTakenDateTime

`func (o *MicrosoftGraphPhoto) GetTakenDateTime() time.Time`

GetTakenDateTime returns the TakenDateTime field if non-nil, zero value otherwise.

### GetTakenDateTimeOk

`func (o *MicrosoftGraphPhoto) GetTakenDateTimeOk() (*time.Time, bool)`

GetTakenDateTimeOk returns a tuple with the TakenDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakenDateTime

`func (o *MicrosoftGraphPhoto) SetTakenDateTime(v time.Time)`

SetTakenDateTime sets TakenDateTime field to given value.

### HasTakenDateTime

`func (o *MicrosoftGraphPhoto) HasTakenDateTime() bool`

HasTakenDateTime returns a boolean if a field has been set.

### SetTakenDateTimeNil

`func (o *MicrosoftGraphPhoto) SetTakenDateTimeNil(b bool)`

 SetTakenDateTimeNil sets the value for TakenDateTime to be an explicit nil

### UnsetTakenDateTime
`func (o *MicrosoftGraphPhoto) UnsetTakenDateTime()`

UnsetTakenDateTime ensures that no value is present for TakenDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


