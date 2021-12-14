# MicrosoftGraphPrinterCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BottomMargins** | Pointer to **[]int32** | A list of supported bottom margins(in microns) for the printer. | [optional] 
**Collation** | Pointer to **NullableBool** | True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise. | [optional] 
**ColorModes** | Pointer to [**[]AnyOfmicrosoftGraphPrintColorMode**](AnyOfmicrosoftGraphPrintColorMode.md) | The color modes supported by the printer. Valid values are described in the following table. | [optional] 
**ContentTypes** | Pointer to **[]string** | A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types. | [optional] 
**CopiesPerJob** | Pointer to [**AnyOfmicrosoftGraphIntegerRange**](anyOf&lt;microsoft.graph.integerRange&gt;.md) | The range of copies per job supported by the printer. | [optional] 
**Dpis** | Pointer to **[]int32** | The list of print resolutions in DPI that are supported by the printer. | [optional] 
**DuplexModes** | Pointer to [**[]AnyOfmicrosoftGraphPrintDuplexMode**](AnyOfmicrosoftGraphPrintDuplexMode.md) | The list of duplex modes that are supported by the printer. Valid values are described in the following table. | [optional] 
**FeedOrientations** | Pointer to [**[]AnyOfmicrosoftGraphPrinterFeedOrientation**](AnyOfmicrosoftGraphPrinterFeedOrientation.md) | The list of feed orientations that are supported by the printer. | [optional] 
**Finishings** | Pointer to [**[]AnyOfmicrosoftGraphPrintFinishing**](AnyOfmicrosoftGraphPrintFinishing.md) | Finishing processes the printer supports for a printed document. | [optional] 
**InputBins** | Pointer to **[]string** | Supported input bins for the printer. | [optional] 
**IsColorPrintingSupported** | Pointer to **NullableBool** | True if color printing is supported by the printer; false otherwise. Read-only. | [optional] 
**IsPageRangeSupported** | Pointer to **NullableBool** | True if the printer supports printing by page ranges; false otherwise. | [optional] 
**LeftMargins** | Pointer to **[]int32** | A list of supported left margins(in microns) for the printer. | [optional] 
**MediaColors** | Pointer to **[]string** | The media (i.e., paper) colors supported by the printer. | [optional] 
**MediaSizes** | Pointer to **[]string** | The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table. | [optional] 
**MediaTypes** | Pointer to **[]string** | The media types supported by the printer. | [optional] 
**MultipageLayouts** | Pointer to [**[]AnyOfmicrosoftGraphPrintMultipageLayout**](AnyOfmicrosoftGraphPrintMultipageLayout.md) | The presentation directions supported by the printer. Supported values are described in the following table. | [optional] 
**Orientations** | Pointer to [**[]AnyOfmicrosoftGraphPrintOrientation**](AnyOfmicrosoftGraphPrintOrientation.md) | The print orientations supported by the printer. Valid values are described in the following table. | [optional] 
**OutputBins** | Pointer to **[]string** | The printer&#39;s supported output bins (trays). | [optional] 
**PagesPerSheet** | Pointer to **[]int32** | Supported number of Input Pages to impose upon a single Impression. | [optional] 
**Qualities** | Pointer to [**[]AnyOfmicrosoftGraphPrintQuality**](AnyOfmicrosoftGraphPrintQuality.md) | The print qualities supported by the printer. | [optional] 
**RightMargins** | Pointer to **[]int32** | A list of supported right margins(in microns) for the printer. | [optional] 
**Scalings** | Pointer to [**[]AnyOfmicrosoftGraphPrintScaling**](AnyOfmicrosoftGraphPrintScaling.md) | Supported print scalings. | [optional] 
**SupportsFitPdfToPage** | Pointer to **NullableBool** | True if the printer supports scaling PDF pages to match the print media size; false otherwise. | [optional] 
**TopMargins** | Pointer to **[]int32** | A list of supported top margins(in microns) for the printer. | [optional] 

## Methods

### NewMicrosoftGraphPrinterCapabilities

`func NewMicrosoftGraphPrinterCapabilities() *MicrosoftGraphPrinterCapabilities`

NewMicrosoftGraphPrinterCapabilities instantiates a new MicrosoftGraphPrinterCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrinterCapabilitiesWithDefaults

`func NewMicrosoftGraphPrinterCapabilitiesWithDefaults() *MicrosoftGraphPrinterCapabilities`

NewMicrosoftGraphPrinterCapabilitiesWithDefaults instantiates a new MicrosoftGraphPrinterCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBottomMargins

`func (o *MicrosoftGraphPrinterCapabilities) GetBottomMargins() []*int32`

GetBottomMargins returns the BottomMargins field if non-nil, zero value otherwise.

### GetBottomMarginsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetBottomMarginsOk() (*[]*int32, bool)`

GetBottomMarginsOk returns a tuple with the BottomMargins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomMargins

`func (o *MicrosoftGraphPrinterCapabilities) SetBottomMargins(v []*int32)`

SetBottomMargins sets BottomMargins field to given value.

### HasBottomMargins

`func (o *MicrosoftGraphPrinterCapabilities) HasBottomMargins() bool`

HasBottomMargins returns a boolean if a field has been set.

### GetCollation

`func (o *MicrosoftGraphPrinterCapabilities) GetCollation() bool`

GetCollation returns the Collation field if non-nil, zero value otherwise.

### GetCollationOk

`func (o *MicrosoftGraphPrinterCapabilities) GetCollationOk() (*bool, bool)`

GetCollationOk returns a tuple with the Collation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollation

`func (o *MicrosoftGraphPrinterCapabilities) SetCollation(v bool)`

SetCollation sets Collation field to given value.

### HasCollation

`func (o *MicrosoftGraphPrinterCapabilities) HasCollation() bool`

HasCollation returns a boolean if a field has been set.

### SetCollationNil

`func (o *MicrosoftGraphPrinterCapabilities) SetCollationNil(b bool)`

 SetCollationNil sets the value for Collation to be an explicit nil

### UnsetCollation
`func (o *MicrosoftGraphPrinterCapabilities) UnsetCollation()`

UnsetCollation ensures that no value is present for Collation, not even an explicit nil
### GetColorModes

`func (o *MicrosoftGraphPrinterCapabilities) GetColorModes() []*AnyOfmicrosoftGraphPrintColorMode`

GetColorModes returns the ColorModes field if non-nil, zero value otherwise.

### GetColorModesOk

`func (o *MicrosoftGraphPrinterCapabilities) GetColorModesOk() (*[]*AnyOfmicrosoftGraphPrintColorMode, bool)`

GetColorModesOk returns a tuple with the ColorModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorModes

`func (o *MicrosoftGraphPrinterCapabilities) SetColorModes(v []*AnyOfmicrosoftGraphPrintColorMode)`

SetColorModes sets ColorModes field to given value.

### HasColorModes

`func (o *MicrosoftGraphPrinterCapabilities) HasColorModes() bool`

HasColorModes returns a boolean if a field has been set.

### GetContentTypes

`func (o *MicrosoftGraphPrinterCapabilities) GetContentTypes() []*string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *MicrosoftGraphPrinterCapabilities) GetContentTypesOk() (*[]*string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *MicrosoftGraphPrinterCapabilities) SetContentTypes(v []*string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *MicrosoftGraphPrinterCapabilities) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetCopiesPerJob

`func (o *MicrosoftGraphPrinterCapabilities) GetCopiesPerJob() AnyOfmicrosoftGraphIntegerRange`

GetCopiesPerJob returns the CopiesPerJob field if non-nil, zero value otherwise.

### GetCopiesPerJobOk

`func (o *MicrosoftGraphPrinterCapabilities) GetCopiesPerJobOk() (*AnyOfmicrosoftGraphIntegerRange, bool)`

GetCopiesPerJobOk returns a tuple with the CopiesPerJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopiesPerJob

`func (o *MicrosoftGraphPrinterCapabilities) SetCopiesPerJob(v AnyOfmicrosoftGraphIntegerRange)`

SetCopiesPerJob sets CopiesPerJob field to given value.

### HasCopiesPerJob

`func (o *MicrosoftGraphPrinterCapabilities) HasCopiesPerJob() bool`

HasCopiesPerJob returns a boolean if a field has been set.

### SetCopiesPerJobNil

`func (o *MicrosoftGraphPrinterCapabilities) SetCopiesPerJobNil(b bool)`

 SetCopiesPerJobNil sets the value for CopiesPerJob to be an explicit nil

### UnsetCopiesPerJob
`func (o *MicrosoftGraphPrinterCapabilities) UnsetCopiesPerJob()`

UnsetCopiesPerJob ensures that no value is present for CopiesPerJob, not even an explicit nil
### GetDpis

`func (o *MicrosoftGraphPrinterCapabilities) GetDpis() []*int32`

GetDpis returns the Dpis field if non-nil, zero value otherwise.

### GetDpisOk

`func (o *MicrosoftGraphPrinterCapabilities) GetDpisOk() (*[]*int32, bool)`

GetDpisOk returns a tuple with the Dpis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpis

`func (o *MicrosoftGraphPrinterCapabilities) SetDpis(v []*int32)`

SetDpis sets Dpis field to given value.

### HasDpis

`func (o *MicrosoftGraphPrinterCapabilities) HasDpis() bool`

HasDpis returns a boolean if a field has been set.

### GetDuplexModes

`func (o *MicrosoftGraphPrinterCapabilities) GetDuplexModes() []*AnyOfmicrosoftGraphPrintDuplexMode`

GetDuplexModes returns the DuplexModes field if non-nil, zero value otherwise.

### GetDuplexModesOk

`func (o *MicrosoftGraphPrinterCapabilities) GetDuplexModesOk() (*[]*AnyOfmicrosoftGraphPrintDuplexMode, bool)`

GetDuplexModesOk returns a tuple with the DuplexModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplexModes

`func (o *MicrosoftGraphPrinterCapabilities) SetDuplexModes(v []*AnyOfmicrosoftGraphPrintDuplexMode)`

SetDuplexModes sets DuplexModes field to given value.

### HasDuplexModes

`func (o *MicrosoftGraphPrinterCapabilities) HasDuplexModes() bool`

HasDuplexModes returns a boolean if a field has been set.

### GetFeedOrientations

`func (o *MicrosoftGraphPrinterCapabilities) GetFeedOrientations() []*AnyOfmicrosoftGraphPrinterFeedOrientation`

GetFeedOrientations returns the FeedOrientations field if non-nil, zero value otherwise.

### GetFeedOrientationsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetFeedOrientationsOk() (*[]*AnyOfmicrosoftGraphPrinterFeedOrientation, bool)`

GetFeedOrientationsOk returns a tuple with the FeedOrientations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedOrientations

`func (o *MicrosoftGraphPrinterCapabilities) SetFeedOrientations(v []*AnyOfmicrosoftGraphPrinterFeedOrientation)`

SetFeedOrientations sets FeedOrientations field to given value.

### HasFeedOrientations

`func (o *MicrosoftGraphPrinterCapabilities) HasFeedOrientations() bool`

HasFeedOrientations returns a boolean if a field has been set.

### GetFinishings

`func (o *MicrosoftGraphPrinterCapabilities) GetFinishings() []*AnyOfmicrosoftGraphPrintFinishing`

GetFinishings returns the Finishings field if non-nil, zero value otherwise.

### GetFinishingsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetFinishingsOk() (*[]*AnyOfmicrosoftGraphPrintFinishing, bool)`

GetFinishingsOk returns a tuple with the Finishings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishings

`func (o *MicrosoftGraphPrinterCapabilities) SetFinishings(v []*AnyOfmicrosoftGraphPrintFinishing)`

SetFinishings sets Finishings field to given value.

### HasFinishings

`func (o *MicrosoftGraphPrinterCapabilities) HasFinishings() bool`

HasFinishings returns a boolean if a field has been set.

### GetInputBins

`func (o *MicrosoftGraphPrinterCapabilities) GetInputBins() []*string`

GetInputBins returns the InputBins field if non-nil, zero value otherwise.

### GetInputBinsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetInputBinsOk() (*[]*string, bool)`

GetInputBinsOk returns a tuple with the InputBins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBins

`func (o *MicrosoftGraphPrinterCapabilities) SetInputBins(v []*string)`

SetInputBins sets InputBins field to given value.

### HasInputBins

`func (o *MicrosoftGraphPrinterCapabilities) HasInputBins() bool`

HasInputBins returns a boolean if a field has been set.

### GetIsColorPrintingSupported

`func (o *MicrosoftGraphPrinterCapabilities) GetIsColorPrintingSupported() bool`

GetIsColorPrintingSupported returns the IsColorPrintingSupported field if non-nil, zero value otherwise.

### GetIsColorPrintingSupportedOk

`func (o *MicrosoftGraphPrinterCapabilities) GetIsColorPrintingSupportedOk() (*bool, bool)`

GetIsColorPrintingSupportedOk returns a tuple with the IsColorPrintingSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsColorPrintingSupported

`func (o *MicrosoftGraphPrinterCapabilities) SetIsColorPrintingSupported(v bool)`

SetIsColorPrintingSupported sets IsColorPrintingSupported field to given value.

### HasIsColorPrintingSupported

`func (o *MicrosoftGraphPrinterCapabilities) HasIsColorPrintingSupported() bool`

HasIsColorPrintingSupported returns a boolean if a field has been set.

### SetIsColorPrintingSupportedNil

`func (o *MicrosoftGraphPrinterCapabilities) SetIsColorPrintingSupportedNil(b bool)`

 SetIsColorPrintingSupportedNil sets the value for IsColorPrintingSupported to be an explicit nil

### UnsetIsColorPrintingSupported
`func (o *MicrosoftGraphPrinterCapabilities) UnsetIsColorPrintingSupported()`

UnsetIsColorPrintingSupported ensures that no value is present for IsColorPrintingSupported, not even an explicit nil
### GetIsPageRangeSupported

`func (o *MicrosoftGraphPrinterCapabilities) GetIsPageRangeSupported() bool`

GetIsPageRangeSupported returns the IsPageRangeSupported field if non-nil, zero value otherwise.

### GetIsPageRangeSupportedOk

`func (o *MicrosoftGraphPrinterCapabilities) GetIsPageRangeSupportedOk() (*bool, bool)`

GetIsPageRangeSupportedOk returns a tuple with the IsPageRangeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPageRangeSupported

`func (o *MicrosoftGraphPrinterCapabilities) SetIsPageRangeSupported(v bool)`

SetIsPageRangeSupported sets IsPageRangeSupported field to given value.

### HasIsPageRangeSupported

`func (o *MicrosoftGraphPrinterCapabilities) HasIsPageRangeSupported() bool`

HasIsPageRangeSupported returns a boolean if a field has been set.

### SetIsPageRangeSupportedNil

`func (o *MicrosoftGraphPrinterCapabilities) SetIsPageRangeSupportedNil(b bool)`

 SetIsPageRangeSupportedNil sets the value for IsPageRangeSupported to be an explicit nil

### UnsetIsPageRangeSupported
`func (o *MicrosoftGraphPrinterCapabilities) UnsetIsPageRangeSupported()`

UnsetIsPageRangeSupported ensures that no value is present for IsPageRangeSupported, not even an explicit nil
### GetLeftMargins

`func (o *MicrosoftGraphPrinterCapabilities) GetLeftMargins() []*int32`

GetLeftMargins returns the LeftMargins field if non-nil, zero value otherwise.

### GetLeftMarginsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetLeftMarginsOk() (*[]*int32, bool)`

GetLeftMarginsOk returns a tuple with the LeftMargins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftMargins

`func (o *MicrosoftGraphPrinterCapabilities) SetLeftMargins(v []*int32)`

SetLeftMargins sets LeftMargins field to given value.

### HasLeftMargins

`func (o *MicrosoftGraphPrinterCapabilities) HasLeftMargins() bool`

HasLeftMargins returns a boolean if a field has been set.

### GetMediaColors

`func (o *MicrosoftGraphPrinterCapabilities) GetMediaColors() []*string`

GetMediaColors returns the MediaColors field if non-nil, zero value otherwise.

### GetMediaColorsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetMediaColorsOk() (*[]*string, bool)`

GetMediaColorsOk returns a tuple with the MediaColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaColors

`func (o *MicrosoftGraphPrinterCapabilities) SetMediaColors(v []*string)`

SetMediaColors sets MediaColors field to given value.

### HasMediaColors

`func (o *MicrosoftGraphPrinterCapabilities) HasMediaColors() bool`

HasMediaColors returns a boolean if a field has been set.

### GetMediaSizes

`func (o *MicrosoftGraphPrinterCapabilities) GetMediaSizes() []*string`

GetMediaSizes returns the MediaSizes field if non-nil, zero value otherwise.

### GetMediaSizesOk

`func (o *MicrosoftGraphPrinterCapabilities) GetMediaSizesOk() (*[]*string, bool)`

GetMediaSizesOk returns a tuple with the MediaSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSizes

`func (o *MicrosoftGraphPrinterCapabilities) SetMediaSizes(v []*string)`

SetMediaSizes sets MediaSizes field to given value.

### HasMediaSizes

`func (o *MicrosoftGraphPrinterCapabilities) HasMediaSizes() bool`

HasMediaSizes returns a boolean if a field has been set.

### GetMediaTypes

`func (o *MicrosoftGraphPrinterCapabilities) GetMediaTypes() []*string`

GetMediaTypes returns the MediaTypes field if non-nil, zero value otherwise.

### GetMediaTypesOk

`func (o *MicrosoftGraphPrinterCapabilities) GetMediaTypesOk() (*[]*string, bool)`

GetMediaTypesOk returns a tuple with the MediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypes

`func (o *MicrosoftGraphPrinterCapabilities) SetMediaTypes(v []*string)`

SetMediaTypes sets MediaTypes field to given value.

### HasMediaTypes

`func (o *MicrosoftGraphPrinterCapabilities) HasMediaTypes() bool`

HasMediaTypes returns a boolean if a field has been set.

### GetMultipageLayouts

`func (o *MicrosoftGraphPrinterCapabilities) GetMultipageLayouts() []*AnyOfmicrosoftGraphPrintMultipageLayout`

GetMultipageLayouts returns the MultipageLayouts field if non-nil, zero value otherwise.

### GetMultipageLayoutsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetMultipageLayoutsOk() (*[]*AnyOfmicrosoftGraphPrintMultipageLayout, bool)`

GetMultipageLayoutsOk returns a tuple with the MultipageLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipageLayouts

`func (o *MicrosoftGraphPrinterCapabilities) SetMultipageLayouts(v []*AnyOfmicrosoftGraphPrintMultipageLayout)`

SetMultipageLayouts sets MultipageLayouts field to given value.

### HasMultipageLayouts

`func (o *MicrosoftGraphPrinterCapabilities) HasMultipageLayouts() bool`

HasMultipageLayouts returns a boolean if a field has been set.

### GetOrientations

`func (o *MicrosoftGraphPrinterCapabilities) GetOrientations() []*AnyOfmicrosoftGraphPrintOrientation`

GetOrientations returns the Orientations field if non-nil, zero value otherwise.

### GetOrientationsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetOrientationsOk() (*[]*AnyOfmicrosoftGraphPrintOrientation, bool)`

GetOrientationsOk returns a tuple with the Orientations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientations

`func (o *MicrosoftGraphPrinterCapabilities) SetOrientations(v []*AnyOfmicrosoftGraphPrintOrientation)`

SetOrientations sets Orientations field to given value.

### HasOrientations

`func (o *MicrosoftGraphPrinterCapabilities) HasOrientations() bool`

HasOrientations returns a boolean if a field has been set.

### GetOutputBins

`func (o *MicrosoftGraphPrinterCapabilities) GetOutputBins() []*string`

GetOutputBins returns the OutputBins field if non-nil, zero value otherwise.

### GetOutputBinsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetOutputBinsOk() (*[]*string, bool)`

GetOutputBinsOk returns a tuple with the OutputBins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputBins

`func (o *MicrosoftGraphPrinterCapabilities) SetOutputBins(v []*string)`

SetOutputBins sets OutputBins field to given value.

### HasOutputBins

`func (o *MicrosoftGraphPrinterCapabilities) HasOutputBins() bool`

HasOutputBins returns a boolean if a field has been set.

### GetPagesPerSheet

`func (o *MicrosoftGraphPrinterCapabilities) GetPagesPerSheet() []*int32`

GetPagesPerSheet returns the PagesPerSheet field if non-nil, zero value otherwise.

### GetPagesPerSheetOk

`func (o *MicrosoftGraphPrinterCapabilities) GetPagesPerSheetOk() (*[]*int32, bool)`

GetPagesPerSheetOk returns a tuple with the PagesPerSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesPerSheet

`func (o *MicrosoftGraphPrinterCapabilities) SetPagesPerSheet(v []*int32)`

SetPagesPerSheet sets PagesPerSheet field to given value.

### HasPagesPerSheet

`func (o *MicrosoftGraphPrinterCapabilities) HasPagesPerSheet() bool`

HasPagesPerSheet returns a boolean if a field has been set.

### GetQualities

`func (o *MicrosoftGraphPrinterCapabilities) GetQualities() []*AnyOfmicrosoftGraphPrintQuality`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *MicrosoftGraphPrinterCapabilities) GetQualitiesOk() (*[]*AnyOfmicrosoftGraphPrintQuality, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *MicrosoftGraphPrinterCapabilities) SetQualities(v []*AnyOfmicrosoftGraphPrintQuality)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *MicrosoftGraphPrinterCapabilities) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetRightMargins

`func (o *MicrosoftGraphPrinterCapabilities) GetRightMargins() []*int32`

GetRightMargins returns the RightMargins field if non-nil, zero value otherwise.

### GetRightMarginsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetRightMarginsOk() (*[]*int32, bool)`

GetRightMarginsOk returns a tuple with the RightMargins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightMargins

`func (o *MicrosoftGraphPrinterCapabilities) SetRightMargins(v []*int32)`

SetRightMargins sets RightMargins field to given value.

### HasRightMargins

`func (o *MicrosoftGraphPrinterCapabilities) HasRightMargins() bool`

HasRightMargins returns a boolean if a field has been set.

### GetScalings

`func (o *MicrosoftGraphPrinterCapabilities) GetScalings() []*AnyOfmicrosoftGraphPrintScaling`

GetScalings returns the Scalings field if non-nil, zero value otherwise.

### GetScalingsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetScalingsOk() (*[]*AnyOfmicrosoftGraphPrintScaling, bool)`

GetScalingsOk returns a tuple with the Scalings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalings

`func (o *MicrosoftGraphPrinterCapabilities) SetScalings(v []*AnyOfmicrosoftGraphPrintScaling)`

SetScalings sets Scalings field to given value.

### HasScalings

`func (o *MicrosoftGraphPrinterCapabilities) HasScalings() bool`

HasScalings returns a boolean if a field has been set.

### GetSupportsFitPdfToPage

`func (o *MicrosoftGraphPrinterCapabilities) GetSupportsFitPdfToPage() bool`

GetSupportsFitPdfToPage returns the SupportsFitPdfToPage field if non-nil, zero value otherwise.

### GetSupportsFitPdfToPageOk

`func (o *MicrosoftGraphPrinterCapabilities) GetSupportsFitPdfToPageOk() (*bool, bool)`

GetSupportsFitPdfToPageOk returns a tuple with the SupportsFitPdfToPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFitPdfToPage

`func (o *MicrosoftGraphPrinterCapabilities) SetSupportsFitPdfToPage(v bool)`

SetSupportsFitPdfToPage sets SupportsFitPdfToPage field to given value.

### HasSupportsFitPdfToPage

`func (o *MicrosoftGraphPrinterCapabilities) HasSupportsFitPdfToPage() bool`

HasSupportsFitPdfToPage returns a boolean if a field has been set.

### SetSupportsFitPdfToPageNil

`func (o *MicrosoftGraphPrinterCapabilities) SetSupportsFitPdfToPageNil(b bool)`

 SetSupportsFitPdfToPageNil sets the value for SupportsFitPdfToPage to be an explicit nil

### UnsetSupportsFitPdfToPage
`func (o *MicrosoftGraphPrinterCapabilities) UnsetSupportsFitPdfToPage()`

UnsetSupportsFitPdfToPage ensures that no value is present for SupportsFitPdfToPage, not even an explicit nil
### GetTopMargins

`func (o *MicrosoftGraphPrinterCapabilities) GetTopMargins() []*int32`

GetTopMargins returns the TopMargins field if non-nil, zero value otherwise.

### GetTopMarginsOk

`func (o *MicrosoftGraphPrinterCapabilities) GetTopMarginsOk() (*[]*int32, bool)`

GetTopMarginsOk returns a tuple with the TopMargins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopMargins

`func (o *MicrosoftGraphPrinterCapabilities) SetTopMargins(v []*int32)`

SetTopMargins sets TopMargins field to given value.

### HasTopMargins

`func (o *MicrosoftGraphPrinterCapabilities) HasTopMargins() bool`

HasTopMargins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


