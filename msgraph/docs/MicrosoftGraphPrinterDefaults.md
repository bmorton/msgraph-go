# MicrosoftGraphPrinterDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorMode** | Pointer to [**AnyOfmicrosoftGraphPrintColorMode**](anyOf&lt;microsoft.graph.printColorMode&gt;.md) | The default color mode to use when printing the document. Valid values are described in the following table. | [optional] 
**ContentType** | Pointer to **NullableString** | The default content (MIME) type to use when processing documents. | [optional] 
**CopiesPerJob** | Pointer to **NullableInt32** | The default number of copies printed per job. | [optional] 
**Dpi** | Pointer to **NullableInt32** | The default resolution in DPI to use when printing the job. | [optional] 
**DuplexMode** | Pointer to [**AnyOfmicrosoftGraphPrintDuplexMode**](anyOf&lt;microsoft.graph.printDuplexMode&gt;.md) | The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table. | [optional] 
**Finishings** | Pointer to [**[]AnyOfmicrosoftGraphPrintFinishing**](AnyOfmicrosoftGraphPrintFinishing.md) | The default set of finishings to apply to print jobs. Valid values are described in the following table. | [optional] 
**FitPdfToPage** | Pointer to **NullableBool** | The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions. | [optional] 
**InputBin** | Pointer to **NullableString** | The default input bin that serves as the paper source. | [optional] 
**MediaColor** | Pointer to **NullableString** | The default media (such as paper) color to print the document on. | [optional] 
**MediaSize** | Pointer to **NullableString** | The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic. | [optional] 
**MediaType** | Pointer to **NullableString** | The default media (such as paper) type to print the document on. | [optional] 
**MultipageLayout** | Pointer to [**AnyOfmicrosoftGraphPrintMultipageLayout**](anyOf&lt;microsoft.graph.printMultipageLayout&gt;.md) | The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table. | [optional] 
**Orientation** | Pointer to [**AnyOfmicrosoftGraphPrintOrientation**](anyOf&lt;microsoft.graph.printOrientation&gt;.md) | The default orientation to use when printing the document. Valid values are described in the following table. | [optional] 
**OutputBin** | Pointer to **NullableString** | The default output bin to place completed prints into. See the printer&#39;s capabilities for a list of supported output bins. | [optional] 
**PagesPerSheet** | Pointer to **NullableInt32** | The default number of document pages to print on each sheet. | [optional] 
**Quality** | Pointer to [**AnyOfmicrosoftGraphPrintQuality**](anyOf&lt;microsoft.graph.printQuality&gt;.md) | The default quality to use when printing the document. Valid values are described in the following table. | [optional] 
**Scaling** | Pointer to [**AnyOfmicrosoftGraphPrintScaling**](anyOf&lt;microsoft.graph.printScaling&gt;.md) | Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table. | [optional] 

## Methods

### NewMicrosoftGraphPrinterDefaults

`func NewMicrosoftGraphPrinterDefaults() *MicrosoftGraphPrinterDefaults`

NewMicrosoftGraphPrinterDefaults instantiates a new MicrosoftGraphPrinterDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrinterDefaultsWithDefaults

`func NewMicrosoftGraphPrinterDefaultsWithDefaults() *MicrosoftGraphPrinterDefaults`

NewMicrosoftGraphPrinterDefaultsWithDefaults instantiates a new MicrosoftGraphPrinterDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColorMode

`func (o *MicrosoftGraphPrinterDefaults) GetColorMode() AnyOfmicrosoftGraphPrintColorMode`

GetColorMode returns the ColorMode field if non-nil, zero value otherwise.

### GetColorModeOk

`func (o *MicrosoftGraphPrinterDefaults) GetColorModeOk() (*AnyOfmicrosoftGraphPrintColorMode, bool)`

GetColorModeOk returns a tuple with the ColorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorMode

`func (o *MicrosoftGraphPrinterDefaults) SetColorMode(v AnyOfmicrosoftGraphPrintColorMode)`

SetColorMode sets ColorMode field to given value.

### HasColorMode

`func (o *MicrosoftGraphPrinterDefaults) HasColorMode() bool`

HasColorMode returns a boolean if a field has been set.

### SetColorModeNil

`func (o *MicrosoftGraphPrinterDefaults) SetColorModeNil(b bool)`

 SetColorModeNil sets the value for ColorMode to be an explicit nil

### UnsetColorMode
`func (o *MicrosoftGraphPrinterDefaults) UnsetColorMode()`

UnsetColorMode ensures that no value is present for ColorMode, not even an explicit nil
### GetContentType

`func (o *MicrosoftGraphPrinterDefaults) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphPrinterDefaults) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MicrosoftGraphPrinterDefaults) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MicrosoftGraphPrinterDefaults) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *MicrosoftGraphPrinterDefaults) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *MicrosoftGraphPrinterDefaults) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetCopiesPerJob

`func (o *MicrosoftGraphPrinterDefaults) GetCopiesPerJob() int32`

GetCopiesPerJob returns the CopiesPerJob field if non-nil, zero value otherwise.

### GetCopiesPerJobOk

`func (o *MicrosoftGraphPrinterDefaults) GetCopiesPerJobOk() (*int32, bool)`

GetCopiesPerJobOk returns a tuple with the CopiesPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopiesPerJob

`func (o *MicrosoftGraphPrinterDefaults) SetCopiesPerJob(v int32)`

SetCopiesPerJob sets CopiesPerJob field to given value.

### HasCopiesPerJob

`func (o *MicrosoftGraphPrinterDefaults) HasCopiesPerJob() bool`

HasCopiesPerJob returns a boolean if a field has been set.

### SetCopiesPerJobNil

`func (o *MicrosoftGraphPrinterDefaults) SetCopiesPerJobNil(b bool)`

 SetCopiesPerJobNil sets the value for CopiesPerJob to be an explicit nil

### UnsetCopiesPerJob
`func (o *MicrosoftGraphPrinterDefaults) UnsetCopiesPerJob()`

UnsetCopiesPerJob ensures that no value is present for CopiesPerJob, not even an explicit nil
### GetDpi

`func (o *MicrosoftGraphPrinterDefaults) GetDpi() int32`

GetDpi returns the Dpi field if non-nil, zero value otherwise.

### GetDpiOk

`func (o *MicrosoftGraphPrinterDefaults) GetDpiOk() (*int32, bool)`

GetDpiOk returns a tuple with the Dpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpi

`func (o *MicrosoftGraphPrinterDefaults) SetDpi(v int32)`

SetDpi sets Dpi field to given value.

### HasDpi

`func (o *MicrosoftGraphPrinterDefaults) HasDpi() bool`

HasDpi returns a boolean if a field has been set.

### SetDpiNil

`func (o *MicrosoftGraphPrinterDefaults) SetDpiNil(b bool)`

 SetDpiNil sets the value for Dpi to be an explicit nil

### UnsetDpi
`func (o *MicrosoftGraphPrinterDefaults) UnsetDpi()`

UnsetDpi ensures that no value is present for Dpi, not even an explicit nil
### GetDuplexMode

`func (o *MicrosoftGraphPrinterDefaults) GetDuplexMode() AnyOfmicrosoftGraphPrintDuplexMode`

GetDuplexMode returns the DuplexMode field if non-nil, zero value otherwise.

### GetDuplexModeOk

`func (o *MicrosoftGraphPrinterDefaults) GetDuplexModeOk() (*AnyOfmicrosoftGraphPrintDuplexMode, bool)`

GetDuplexModeOk returns a tuple with the DuplexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplexMode

`func (o *MicrosoftGraphPrinterDefaults) SetDuplexMode(v AnyOfmicrosoftGraphPrintDuplexMode)`

SetDuplexMode sets DuplexMode field to given value.

### HasDuplexMode

`func (o *MicrosoftGraphPrinterDefaults) HasDuplexMode() bool`

HasDuplexMode returns a boolean if a field has been set.

### SetDuplexModeNil

`func (o *MicrosoftGraphPrinterDefaults) SetDuplexModeNil(b bool)`

 SetDuplexModeNil sets the value for DuplexMode to be an explicit nil

### UnsetDuplexMode
`func (o *MicrosoftGraphPrinterDefaults) UnsetDuplexMode()`

UnsetDuplexMode ensures that no value is present for DuplexMode, not even an explicit nil
### GetFinishings

`func (o *MicrosoftGraphPrinterDefaults) GetFinishings() []*AnyOfmicrosoftGraphPrintFinishing`

GetFinishings returns the Finishings field if non-nil, zero value otherwise.

### GetFinishingsOk

`func (o *MicrosoftGraphPrinterDefaults) GetFinishingsOk() (*[]*AnyOfmicrosoftGraphPrintFinishing, bool)`

GetFinishingsOk returns a tuple with the Finishings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishings

`func (o *MicrosoftGraphPrinterDefaults) SetFinishings(v []*AnyOfmicrosoftGraphPrintFinishing)`

SetFinishings sets Finishings field to given value.

### HasFinishings

`func (o *MicrosoftGraphPrinterDefaults) HasFinishings() bool`

HasFinishings returns a boolean if a field has been set.

### GetFitPdfToPage

`func (o *MicrosoftGraphPrinterDefaults) GetFitPdfToPage() bool`

GetFitPdfToPage returns the FitPdfToPage field if non-nil, zero value otherwise.

### GetFitPdfToPageOk

`func (o *MicrosoftGraphPrinterDefaults) GetFitPdfToPageOk() (*bool, bool)`

GetFitPdfToPageOk returns a tuple with the FitPdfToPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitPdfToPage

`func (o *MicrosoftGraphPrinterDefaults) SetFitPdfToPage(v bool)`

SetFitPdfToPage sets FitPdfToPage field to given value.

### HasFitPdfToPage

`func (o *MicrosoftGraphPrinterDefaults) HasFitPdfToPage() bool`

HasFitPdfToPage returns a boolean if a field has been set.

### SetFitPdfToPageNil

`func (o *MicrosoftGraphPrinterDefaults) SetFitPdfToPageNil(b bool)`

 SetFitPdfToPageNil sets the value for FitPdfToPage to be an explicit nil

### UnsetFitPdfToPage
`func (o *MicrosoftGraphPrinterDefaults) UnsetFitPdfToPage()`

UnsetFitPdfToPage ensures that no value is present for FitPdfToPage, not even an explicit nil
### GetInputBin

`func (o *MicrosoftGraphPrinterDefaults) GetInputBin() string`

GetInputBin returns the InputBin field if non-nil, zero value otherwise.

### GetInputBinOk

`func (o *MicrosoftGraphPrinterDefaults) GetInputBinOk() (*string, bool)`

GetInputBinOk returns a tuple with the InputBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBin

`func (o *MicrosoftGraphPrinterDefaults) SetInputBin(v string)`

SetInputBin sets InputBin field to given value.

### HasInputBin

`func (o *MicrosoftGraphPrinterDefaults) HasInputBin() bool`

HasInputBin returns a boolean if a field has been set.

### SetInputBinNil

`func (o *MicrosoftGraphPrinterDefaults) SetInputBinNil(b bool)`

 SetInputBinNil sets the value for InputBin to be an explicit nil

### UnsetInputBin
`func (o *MicrosoftGraphPrinterDefaults) UnsetInputBin()`

UnsetInputBin ensures that no value is present for InputBin, not even an explicit nil
### GetMediaColor

`func (o *MicrosoftGraphPrinterDefaults) GetMediaColor() string`

GetMediaColor returns the MediaColor field if non-nil, zero value otherwise.

### GetMediaColorOk

`func (o *MicrosoftGraphPrinterDefaults) GetMediaColorOk() (*string, bool)`

GetMediaColorOk returns a tuple with the MediaColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaColor

`func (o *MicrosoftGraphPrinterDefaults) SetMediaColor(v string)`

SetMediaColor sets MediaColor field to given value.

### HasMediaColor

`func (o *MicrosoftGraphPrinterDefaults) HasMediaColor() bool`

HasMediaColor returns a boolean if a field has been set.

### SetMediaColorNil

`func (o *MicrosoftGraphPrinterDefaults) SetMediaColorNil(b bool)`

 SetMediaColorNil sets the value for MediaColor to be an explicit nil

### UnsetMediaColor
`func (o *MicrosoftGraphPrinterDefaults) UnsetMediaColor()`

UnsetMediaColor ensures that no value is present for MediaColor, not even an explicit nil
### GetMediaSize

`func (o *MicrosoftGraphPrinterDefaults) GetMediaSize() string`

GetMediaSize returns the MediaSize field if non-nil, zero value otherwise.

### GetMediaSizeOk

`func (o *MicrosoftGraphPrinterDefaults) GetMediaSizeOk() (*string, bool)`

GetMediaSizeOk returns a tuple with the MediaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSize

`func (o *MicrosoftGraphPrinterDefaults) SetMediaSize(v string)`

SetMediaSize sets MediaSize field to given value.

### HasMediaSize

`func (o *MicrosoftGraphPrinterDefaults) HasMediaSize() bool`

HasMediaSize returns a boolean if a field has been set.

### SetMediaSizeNil

`func (o *MicrosoftGraphPrinterDefaults) SetMediaSizeNil(b bool)`

 SetMediaSizeNil sets the value for MediaSize to be an explicit nil

### UnsetMediaSize
`func (o *MicrosoftGraphPrinterDefaults) UnsetMediaSize()`

UnsetMediaSize ensures that no value is present for MediaSize, not even an explicit nil
### GetMediaType

`func (o *MicrosoftGraphPrinterDefaults) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *MicrosoftGraphPrinterDefaults) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *MicrosoftGraphPrinterDefaults) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *MicrosoftGraphPrinterDefaults) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *MicrosoftGraphPrinterDefaults) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *MicrosoftGraphPrinterDefaults) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetMultipageLayout

`func (o *MicrosoftGraphPrinterDefaults) GetMultipageLayout() AnyOfmicrosoftGraphPrintMultipageLayout`

GetMultipageLayout returns the MultipageLayout field if non-nil, zero value otherwise.

### GetMultipageLayoutOk

`func (o *MicrosoftGraphPrinterDefaults) GetMultipageLayoutOk() (*AnyOfmicrosoftGraphPrintMultipageLayout, bool)`

GetMultipageLayoutOk returns a tuple with the MultipageLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipageLayout

`func (o *MicrosoftGraphPrinterDefaults) SetMultipageLayout(v AnyOfmicrosoftGraphPrintMultipageLayout)`

SetMultipageLayout sets MultipageLayout field to given value.

### HasMultipageLayout

`func (o *MicrosoftGraphPrinterDefaults) HasMultipageLayout() bool`

HasMultipageLayout returns a boolean if a field has been set.

### SetMultipageLayoutNil

`func (o *MicrosoftGraphPrinterDefaults) SetMultipageLayoutNil(b bool)`

 SetMultipageLayoutNil sets the value for MultipageLayout to be an explicit nil

### UnsetMultipageLayout
`func (o *MicrosoftGraphPrinterDefaults) UnsetMultipageLayout()`

UnsetMultipageLayout ensures that no value is present for MultipageLayout, not even an explicit nil
### GetOrientation

`func (o *MicrosoftGraphPrinterDefaults) GetOrientation() AnyOfmicrosoftGraphPrintOrientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *MicrosoftGraphPrinterDefaults) GetOrientationOk() (*AnyOfmicrosoftGraphPrintOrientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *MicrosoftGraphPrinterDefaults) SetOrientation(v AnyOfmicrosoftGraphPrintOrientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *MicrosoftGraphPrinterDefaults) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### SetOrientationNil

`func (o *MicrosoftGraphPrinterDefaults) SetOrientationNil(b bool)`

 SetOrientationNil sets the value for Orientation to be an explicit nil

### UnsetOrientation
`func (o *MicrosoftGraphPrinterDefaults) UnsetOrientation()`

UnsetOrientation ensures that no value is present for Orientation, not even an explicit nil
### GetOutputBin

`func (o *MicrosoftGraphPrinterDefaults) GetOutputBin() string`

GetOutputBin returns the OutputBin field if non-nil, zero value otherwise.

### GetOutputBinOk

`func (o *MicrosoftGraphPrinterDefaults) GetOutputBinOk() (*string, bool)`

GetOutputBinOk returns a tuple with the OutputBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputBin

`func (o *MicrosoftGraphPrinterDefaults) SetOutputBin(v string)`

SetOutputBin sets OutputBin field to given value.

### HasOutputBin

`func (o *MicrosoftGraphPrinterDefaults) HasOutputBin() bool`

HasOutputBin returns a boolean if a field has been set.

### SetOutputBinNil

`func (o *MicrosoftGraphPrinterDefaults) SetOutputBinNil(b bool)`

 SetOutputBinNil sets the value for OutputBin to be an explicit nil

### UnsetOutputBin
`func (o *MicrosoftGraphPrinterDefaults) UnsetOutputBin()`

UnsetOutputBin ensures that no value is present for OutputBin, not even an explicit nil
### GetPagesPerSheet

`func (o *MicrosoftGraphPrinterDefaults) GetPagesPerSheet() int32`

GetPagesPerSheet returns the PagesPerSheet field if non-nil, zero value otherwise.

### GetPagesPerSheetOk

`func (o *MicrosoftGraphPrinterDefaults) GetPagesPerSheetOk() (*int32, bool)`

GetPagesPerSheetOk returns a tuple with the PagesPerSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesPerSheet

`func (o *MicrosoftGraphPrinterDefaults) SetPagesPerSheet(v int32)`

SetPagesPerSheet sets PagesPerSheet field to given value.

### HasPagesPerSheet

`func (o *MicrosoftGraphPrinterDefaults) HasPagesPerSheet() bool`

HasPagesPerSheet returns a boolean if a field has been set.

### SetPagesPerSheetNil

`func (o *MicrosoftGraphPrinterDefaults) SetPagesPerSheetNil(b bool)`

 SetPagesPerSheetNil sets the value for PagesPerSheet to be an explicit nil

### UnsetPagesPerSheet
`func (o *MicrosoftGraphPrinterDefaults) UnsetPagesPerSheet()`

UnsetPagesPerSheet ensures that no value is present for PagesPerSheet, not even an explicit nil
### GetQuality

`func (o *MicrosoftGraphPrinterDefaults) GetQuality() AnyOfmicrosoftGraphPrintQuality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *MicrosoftGraphPrinterDefaults) GetQualityOk() (*AnyOfmicrosoftGraphPrintQuality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *MicrosoftGraphPrinterDefaults) SetQuality(v AnyOfmicrosoftGraphPrintQuality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *MicrosoftGraphPrinterDefaults) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### SetQualityNil

`func (o *MicrosoftGraphPrinterDefaults) SetQualityNil(b bool)`

 SetQualityNil sets the value for Quality to be an explicit nil

### UnsetQuality
`func (o *MicrosoftGraphPrinterDefaults) UnsetQuality()`

UnsetQuality ensures that no value is present for Quality, not even an explicit nil
### GetScaling

`func (o *MicrosoftGraphPrinterDefaults) GetScaling() AnyOfmicrosoftGraphPrintScaling`

GetScaling returns the Scaling field if non-nil, zero value otherwise.

### GetScalingOk

`func (o *MicrosoftGraphPrinterDefaults) GetScalingOk() (*AnyOfmicrosoftGraphPrintScaling, bool)`

GetScalingOk returns a tuple with the Scaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaling

`func (o *MicrosoftGraphPrinterDefaults) SetScaling(v AnyOfmicrosoftGraphPrintScaling)`

SetScaling sets Scaling field to given value.

### HasScaling

`func (o *MicrosoftGraphPrinterDefaults) HasScaling() bool`

HasScaling returns a boolean if a field has been set.

### SetScalingNil

`func (o *MicrosoftGraphPrinterDefaults) SetScalingNil(b bool)`

 SetScalingNil sets the value for Scaling to be an explicit nil

### UnsetScaling
`func (o *MicrosoftGraphPrinterDefaults) UnsetScaling()`

UnsetScaling ensures that no value is present for Scaling, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


