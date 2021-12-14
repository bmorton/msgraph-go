# MicrosoftGraphPrintJobConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collate** | Pointer to **NullableBool** | Whether the printer should collate pages wehen printing multiple copies of a multi-page document. | [optional] 
**ColorMode** | Pointer to [**AnyOfmicrosoftGraphPrintColorMode**](anyOf&lt;microsoft.graph.printColorMode&gt;.md) | The color mode the printer should use to print the job. Valid values are described in the table below. Read-only. | [optional] 
**Copies** | Pointer to **NullableInt32** | The number of copies that should be printed. Read-only. | [optional] 
**Dpi** | Pointer to **NullableInt32** | The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only. | [optional] 
**DuplexMode** | Pointer to [**AnyOfmicrosoftGraphPrintDuplexMode**](anyOf&lt;microsoft.graph.printDuplexMode&gt;.md) | The duplex mode the printer should use when printing the job. Valid values are described in the table below. Read-only. | [optional] 
**FeedOrientation** | Pointer to [**AnyOfmicrosoftGraphPrinterFeedOrientation**](anyOf&lt;microsoft.graph.printerFeedOrientation&gt;.md) | The orientation to use when feeding media into the printer. Valid values are described in the following table. Read-only. | [optional] 
**Finishings** | Pointer to [**[]AnyOfmicrosoftGraphPrintFinishing**](AnyOfmicrosoftGraphPrintFinishing.md) | Finishing processes to use when printing. | [optional] 
**FitPdfToPage** | Pointer to **NullableBool** |  | [optional] 
**InputBin** | Pointer to **NullableString** | The input bin (tray) to use when printing. See the printer&#39;s capabilities for a list of supported input bins. | [optional] 
**Margin** | Pointer to [**AnyOfmicrosoftGraphPrintMargin**](anyOf&lt;microsoft.graph.printMargin&gt;.md) | The margin settings to use when printing. | [optional] 
**MediaSize** | Pointer to **NullableString** | The media size to use when printing. Supports standard size names for ISO and ANSI media sizes. Valid values listed in the printerCapabilities topic. | [optional] 
**MediaType** | Pointer to **NullableString** | The default media (such as paper) type to print the document on. | [optional] 
**MultipageLayout** | Pointer to [**AnyOfmicrosoftGraphPrintMultipageLayout**](anyOf&lt;microsoft.graph.printMultipageLayout&gt;.md) | The direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table. | [optional] 
**Orientation** | Pointer to [**AnyOfmicrosoftGraphPrintOrientation**](anyOf&lt;microsoft.graph.printOrientation&gt;.md) | The orientation setting the printer should use when printing the job. Valid values are described in the following table. | [optional] 
**OutputBin** | Pointer to **NullableString** | The output bin to place completed prints into. See the printer&#39;s capabilities for a list of supported output bins. | [optional] 
**PageRanges** | Pointer to [**[]AnyOfmicrosoftGraphIntegerRange**](AnyOfmicrosoftGraphIntegerRange.md) | The page ranges to print. Read-only. | [optional] 
**PagesPerSheet** | Pointer to **NullableInt32** | The number of document pages to print on each sheet. | [optional] 
**Quality** | Pointer to [**AnyOfmicrosoftGraphPrintQuality**](anyOf&lt;microsoft.graph.printQuality&gt;.md) | The print quality to use when printing the job. Valid values are described in the table below. Read-only. | [optional] 
**Scaling** | Pointer to [**AnyOfmicrosoftGraphPrintScaling**](anyOf&lt;microsoft.graph.printScaling&gt;.md) | Specifies how the printer should scale the document data to fit the requested media. Valid values are described in the following table. | [optional] 

## Methods

### NewMicrosoftGraphPrintJobConfiguration

`func NewMicrosoftGraphPrintJobConfiguration() *MicrosoftGraphPrintJobConfiguration`

NewMicrosoftGraphPrintJobConfiguration instantiates a new MicrosoftGraphPrintJobConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintJobConfigurationWithDefaults

`func NewMicrosoftGraphPrintJobConfigurationWithDefaults() *MicrosoftGraphPrintJobConfiguration`

NewMicrosoftGraphPrintJobConfigurationWithDefaults instantiates a new MicrosoftGraphPrintJobConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollate

`func (o *MicrosoftGraphPrintJobConfiguration) GetCollate() bool`

GetCollate returns the Collate field if non-nil, zero value otherwise.

### GetCollateOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetCollateOk() (*bool, bool)`

GetCollateOk returns a tuple with the Collate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollate

`func (o *MicrosoftGraphPrintJobConfiguration) SetCollate(v bool)`

SetCollate sets Collate field to given value.

### HasCollate

`func (o *MicrosoftGraphPrintJobConfiguration) HasCollate() bool`

HasCollate returns a boolean if a field has been set.

### SetCollateNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetCollateNil(b bool)`

 SetCollateNil sets the value for Collate to be an explicit nil

### UnsetCollate
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetCollate()`

UnsetCollate ensures that no value is present for Collate, not even an explicit nil
### GetColorMode

`func (o *MicrosoftGraphPrintJobConfiguration) GetColorMode() AnyOfmicrosoftGraphPrintColorMode`

GetColorMode returns the ColorMode field if non-nil, zero value otherwise.

### GetColorModeOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetColorModeOk() (*AnyOfmicrosoftGraphPrintColorMode, bool)`

GetColorModeOk returns a tuple with the ColorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorMode

`func (o *MicrosoftGraphPrintJobConfiguration) SetColorMode(v AnyOfmicrosoftGraphPrintColorMode)`

SetColorMode sets ColorMode field to given value.

### HasColorMode

`func (o *MicrosoftGraphPrintJobConfiguration) HasColorMode() bool`

HasColorMode returns a boolean if a field has been set.

### SetColorModeNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetColorModeNil(b bool)`

 SetColorModeNil sets the value for ColorMode to be an explicit nil

### UnsetColorMode
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetColorMode()`

UnsetColorMode ensures that no value is present for ColorMode, not even an explicit nil
### GetCopies

`func (o *MicrosoftGraphPrintJobConfiguration) GetCopies() int32`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetCopiesOk() (*int32, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *MicrosoftGraphPrintJobConfiguration) SetCopies(v int32)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *MicrosoftGraphPrintJobConfiguration) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### SetCopiesNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetCopiesNil(b bool)`

 SetCopiesNil sets the value for Copies to be an explicit nil

### UnsetCopies
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetCopies()`

UnsetCopies ensures that no value is present for Copies, not even an explicit nil
### GetDpi

`func (o *MicrosoftGraphPrintJobConfiguration) GetDpi() int32`

GetDpi returns the Dpi field if non-nil, zero value otherwise.

### GetDpiOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetDpiOk() (*int32, bool)`

GetDpiOk returns a tuple with the Dpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpi

`func (o *MicrosoftGraphPrintJobConfiguration) SetDpi(v int32)`

SetDpi sets Dpi field to given value.

### HasDpi

`func (o *MicrosoftGraphPrintJobConfiguration) HasDpi() bool`

HasDpi returns a boolean if a field has been set.

### SetDpiNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetDpiNil(b bool)`

 SetDpiNil sets the value for Dpi to be an explicit nil

### UnsetDpi
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetDpi()`

UnsetDpi ensures that no value is present for Dpi, not even an explicit nil
### GetDuplexMode

`func (o *MicrosoftGraphPrintJobConfiguration) GetDuplexMode() AnyOfmicrosoftGraphPrintDuplexMode`

GetDuplexMode returns the DuplexMode field if non-nil, zero value otherwise.

### GetDuplexModeOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetDuplexModeOk() (*AnyOfmicrosoftGraphPrintDuplexMode, bool)`

GetDuplexModeOk returns a tuple with the DuplexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplexMode

`func (o *MicrosoftGraphPrintJobConfiguration) SetDuplexMode(v AnyOfmicrosoftGraphPrintDuplexMode)`

SetDuplexMode sets DuplexMode field to given value.

### HasDuplexMode

`func (o *MicrosoftGraphPrintJobConfiguration) HasDuplexMode() bool`

HasDuplexMode returns a boolean if a field has been set.

### SetDuplexModeNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetDuplexModeNil(b bool)`

 SetDuplexModeNil sets the value for DuplexMode to be an explicit nil

### UnsetDuplexMode
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetDuplexMode()`

UnsetDuplexMode ensures that no value is present for DuplexMode, not even an explicit nil
### GetFeedOrientation

`func (o *MicrosoftGraphPrintJobConfiguration) GetFeedOrientation() AnyOfmicrosoftGraphPrinterFeedOrientation`

GetFeedOrientation returns the FeedOrientation field if non-nil, zero value otherwise.

### GetFeedOrientationOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetFeedOrientationOk() (*AnyOfmicrosoftGraphPrinterFeedOrientation, bool)`

GetFeedOrientationOk returns a tuple with the FeedOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedOrientation

`func (o *MicrosoftGraphPrintJobConfiguration) SetFeedOrientation(v AnyOfmicrosoftGraphPrinterFeedOrientation)`

SetFeedOrientation sets FeedOrientation field to given value.

### HasFeedOrientation

`func (o *MicrosoftGraphPrintJobConfiguration) HasFeedOrientation() bool`

HasFeedOrientation returns a boolean if a field has been set.

### SetFeedOrientationNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetFeedOrientationNil(b bool)`

 SetFeedOrientationNil sets the value for FeedOrientation to be an explicit nil

### UnsetFeedOrientation
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetFeedOrientation()`

UnsetFeedOrientation ensures that no value is present for FeedOrientation, not even an explicit nil
### GetFinishings

`func (o *MicrosoftGraphPrintJobConfiguration) GetFinishings() []*AnyOfmicrosoftGraphPrintFinishing`

GetFinishings returns the Finishings field if non-nil, zero value otherwise.

### GetFinishingsOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetFinishingsOk() (*[]*AnyOfmicrosoftGraphPrintFinishing, bool)`

GetFinishingsOk returns a tuple with the Finishings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishings

`func (o *MicrosoftGraphPrintJobConfiguration) SetFinishings(v []*AnyOfmicrosoftGraphPrintFinishing)`

SetFinishings sets Finishings field to given value.

### HasFinishings

`func (o *MicrosoftGraphPrintJobConfiguration) HasFinishings() bool`

HasFinishings returns a boolean if a field has been set.

### GetFitPdfToPage

`func (o *MicrosoftGraphPrintJobConfiguration) GetFitPdfToPage() bool`

GetFitPdfToPage returns the FitPdfToPage field if non-nil, zero value otherwise.

### GetFitPdfToPageOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetFitPdfToPageOk() (*bool, bool)`

GetFitPdfToPageOk returns a tuple with the FitPdfToPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitPdfToPage

`func (o *MicrosoftGraphPrintJobConfiguration) SetFitPdfToPage(v bool)`

SetFitPdfToPage sets FitPdfToPage field to given value.

### HasFitPdfToPage

`func (o *MicrosoftGraphPrintJobConfiguration) HasFitPdfToPage() bool`

HasFitPdfToPage returns a boolean if a field has been set.

### SetFitPdfToPageNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetFitPdfToPageNil(b bool)`

 SetFitPdfToPageNil sets the value for FitPdfToPage to be an explicit nil

### UnsetFitPdfToPage
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetFitPdfToPage()`

UnsetFitPdfToPage ensures that no value is present for FitPdfToPage, not even an explicit nil
### GetInputBin

`func (o *MicrosoftGraphPrintJobConfiguration) GetInputBin() string`

GetInputBin returns the InputBin field if non-nil, zero value otherwise.

### GetInputBinOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetInputBinOk() (*string, bool)`

GetInputBinOk returns a tuple with the InputBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBin

`func (o *MicrosoftGraphPrintJobConfiguration) SetInputBin(v string)`

SetInputBin sets InputBin field to given value.

### HasInputBin

`func (o *MicrosoftGraphPrintJobConfiguration) HasInputBin() bool`

HasInputBin returns a boolean if a field has been set.

### SetInputBinNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetInputBinNil(b bool)`

 SetInputBinNil sets the value for InputBin to be an explicit nil

### UnsetInputBin
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetInputBin()`

UnsetInputBin ensures that no value is present for InputBin, not even an explicit nil
### GetMargin

`func (o *MicrosoftGraphPrintJobConfiguration) GetMargin() AnyOfmicrosoftGraphPrintMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetMarginOk() (*AnyOfmicrosoftGraphPrintMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *MicrosoftGraphPrintJobConfiguration) SetMargin(v AnyOfmicrosoftGraphPrintMargin)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *MicrosoftGraphPrintJobConfiguration) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### SetMarginNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetMarginNil(b bool)`

 SetMarginNil sets the value for Margin to be an explicit nil

### UnsetMargin
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetMargin()`

UnsetMargin ensures that no value is present for Margin, not even an explicit nil
### GetMediaSize

`func (o *MicrosoftGraphPrintJobConfiguration) GetMediaSize() string`

GetMediaSize returns the MediaSize field if non-nil, zero value otherwise.

### GetMediaSizeOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetMediaSizeOk() (*string, bool)`

GetMediaSizeOk returns a tuple with the MediaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSize

`func (o *MicrosoftGraphPrintJobConfiguration) SetMediaSize(v string)`

SetMediaSize sets MediaSize field to given value.

### HasMediaSize

`func (o *MicrosoftGraphPrintJobConfiguration) HasMediaSize() bool`

HasMediaSize returns a boolean if a field has been set.

### SetMediaSizeNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetMediaSizeNil(b bool)`

 SetMediaSizeNil sets the value for MediaSize to be an explicit nil

### UnsetMediaSize
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetMediaSize()`

UnsetMediaSize ensures that no value is present for MediaSize, not even an explicit nil
### GetMediaType

`func (o *MicrosoftGraphPrintJobConfiguration) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *MicrosoftGraphPrintJobConfiguration) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *MicrosoftGraphPrintJobConfiguration) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetMultipageLayout

`func (o *MicrosoftGraphPrintJobConfiguration) GetMultipageLayout() AnyOfmicrosoftGraphPrintMultipageLayout`

GetMultipageLayout returns the MultipageLayout field if non-nil, zero value otherwise.

### GetMultipageLayoutOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetMultipageLayoutOk() (*AnyOfmicrosoftGraphPrintMultipageLayout, bool)`

GetMultipageLayoutOk returns a tuple with the MultipageLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipageLayout

`func (o *MicrosoftGraphPrintJobConfiguration) SetMultipageLayout(v AnyOfmicrosoftGraphPrintMultipageLayout)`

SetMultipageLayout sets MultipageLayout field to given value.

### HasMultipageLayout

`func (o *MicrosoftGraphPrintJobConfiguration) HasMultipageLayout() bool`

HasMultipageLayout returns a boolean if a field has been set.

### SetMultipageLayoutNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetMultipageLayoutNil(b bool)`

 SetMultipageLayoutNil sets the value for MultipageLayout to be an explicit nil

### UnsetMultipageLayout
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetMultipageLayout()`

UnsetMultipageLayout ensures that no value is present for MultipageLayout, not even an explicit nil
### GetOrientation

`func (o *MicrosoftGraphPrintJobConfiguration) GetOrientation() AnyOfmicrosoftGraphPrintOrientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetOrientationOk() (*AnyOfmicrosoftGraphPrintOrientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *MicrosoftGraphPrintJobConfiguration) SetOrientation(v AnyOfmicrosoftGraphPrintOrientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *MicrosoftGraphPrintJobConfiguration) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### SetOrientationNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetOrientationNil(b bool)`

 SetOrientationNil sets the value for Orientation to be an explicit nil

### UnsetOrientation
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetOrientation()`

UnsetOrientation ensures that no value is present for Orientation, not even an explicit nil
### GetOutputBin

`func (o *MicrosoftGraphPrintJobConfiguration) GetOutputBin() string`

GetOutputBin returns the OutputBin field if non-nil, zero value otherwise.

### GetOutputBinOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetOutputBinOk() (*string, bool)`

GetOutputBinOk returns a tuple with the OutputBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputBin

`func (o *MicrosoftGraphPrintJobConfiguration) SetOutputBin(v string)`

SetOutputBin sets OutputBin field to given value.

### HasOutputBin

`func (o *MicrosoftGraphPrintJobConfiguration) HasOutputBin() bool`

HasOutputBin returns a boolean if a field has been set.

### SetOutputBinNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetOutputBinNil(b bool)`

 SetOutputBinNil sets the value for OutputBin to be an explicit nil

### UnsetOutputBin
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetOutputBin()`

UnsetOutputBin ensures that no value is present for OutputBin, not even an explicit nil
### GetPageRanges

`func (o *MicrosoftGraphPrintJobConfiguration) GetPageRanges() []*AnyOfmicrosoftGraphIntegerRange`

GetPageRanges returns the PageRanges field if non-nil, zero value otherwise.

### GetPageRangesOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetPageRangesOk() (*[]*AnyOfmicrosoftGraphIntegerRange, bool)`

GetPageRangesOk returns a tuple with the PageRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageRanges

`func (o *MicrosoftGraphPrintJobConfiguration) SetPageRanges(v []*AnyOfmicrosoftGraphIntegerRange)`

SetPageRanges sets PageRanges field to given value.

### HasPageRanges

`func (o *MicrosoftGraphPrintJobConfiguration) HasPageRanges() bool`

HasPageRanges returns a boolean if a field has been set.

### GetPagesPerSheet

`func (o *MicrosoftGraphPrintJobConfiguration) GetPagesPerSheet() int32`

GetPagesPerSheet returns the PagesPerSheet field if non-nil, zero value otherwise.

### GetPagesPerSheetOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetPagesPerSheetOk() (*int32, bool)`

GetPagesPerSheetOk returns a tuple with the PagesPerSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesPerSheet

`func (o *MicrosoftGraphPrintJobConfiguration) SetPagesPerSheet(v int32)`

SetPagesPerSheet sets PagesPerSheet field to given value.

### HasPagesPerSheet

`func (o *MicrosoftGraphPrintJobConfiguration) HasPagesPerSheet() bool`

HasPagesPerSheet returns a boolean if a field has been set.

### SetPagesPerSheetNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetPagesPerSheetNil(b bool)`

 SetPagesPerSheetNil sets the value for PagesPerSheet to be an explicit nil

### UnsetPagesPerSheet
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetPagesPerSheet()`

UnsetPagesPerSheet ensures that no value is present for PagesPerSheet, not even an explicit nil
### GetQuality

`func (o *MicrosoftGraphPrintJobConfiguration) GetQuality() AnyOfmicrosoftGraphPrintQuality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetQualityOk() (*AnyOfmicrosoftGraphPrintQuality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *MicrosoftGraphPrintJobConfiguration) SetQuality(v AnyOfmicrosoftGraphPrintQuality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *MicrosoftGraphPrintJobConfiguration) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### SetQualityNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetQualityNil(b bool)`

 SetQualityNil sets the value for Quality to be an explicit nil

### UnsetQuality
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetQuality()`

UnsetQuality ensures that no value is present for Quality, not even an explicit nil
### GetScaling

`func (o *MicrosoftGraphPrintJobConfiguration) GetScaling() AnyOfmicrosoftGraphPrintScaling`

GetScaling returns the Scaling field if non-nil, zero value otherwise.

### GetScalingOk

`func (o *MicrosoftGraphPrintJobConfiguration) GetScalingOk() (*AnyOfmicrosoftGraphPrintScaling, bool)`

GetScalingOk returns a tuple with the Scaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaling

`func (o *MicrosoftGraphPrintJobConfiguration) SetScaling(v AnyOfmicrosoftGraphPrintScaling)`

SetScaling sets Scaling field to given value.

### HasScaling

`func (o *MicrosoftGraphPrintJobConfiguration) HasScaling() bool`

HasScaling returns a boolean if a field has been set.

### SetScalingNil

`func (o *MicrosoftGraphPrintJobConfiguration) SetScalingNil(b bool)`

 SetScalingNil sets the value for Scaling to be an explicit nil

### UnsetScaling
`func (o *MicrosoftGraphPrintJobConfiguration) UnsetScaling()`

UnsetScaling ensures that no value is present for Scaling, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


