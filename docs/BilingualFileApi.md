# \BilingualFileApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareBilingualFile**](BilingualFileApi.md#CompareBilingualFile) | **Post** /api2/v1/bilingualFiles/compare | Compare bilingual file
[**ConvertBilingualFile**](BilingualFileApi.md#ConvertBilingualFile) | **Post** /api2/v1/bilingualFiles/convert | Convert bilingual file
[**GetBilingualFile**](BilingualFileApi.md#GetBilingualFile) | **Post** /api2/v1/projects/{projectUid}/jobs/bilingualFile | Download bilingual file
[**GetPreviewFile**](BilingualFileApi.md#GetPreviewFile) | **Post** /api2/v1/bilingualFiles/preview | Download preview
[**UploadBilingualFile**](BilingualFileApi.md#UploadBilingualFile) | **Put** /api2/v1/bilingualFiles | Upload bilingual file


# **CompareBilingualFile**
> ComparedSegmentsDto CompareBilingualFile(ctx, optional)
Compare bilingual file

Compares bilingual file to job state. Returns list of compared segments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BilingualFileApiCompareBilingualFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BilingualFileApiCompareBilingualFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 
 **workflowLevel** | **optional.Int32**|  | [default to 1]

### Return type

[**ComparedSegmentsDto**](ComparedSegmentsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertBilingualFile**
> ConvertBilingualFile(ctx, from, to, optional)
Convert bilingual file



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **from** | **string**|  | 
  **to** | **string**|  | 
 **optional** | ***BilingualFileApiConvertBilingualFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BilingualFileApiConvertBilingualFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBilingualFile**
> GetBilingualFile(ctx, projectUid, optional)
Download bilingual file



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***BilingualFileApiGetBilingualFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BilingualFileApiGetBilingualFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartReadyReferences**](JobPartReadyReferences.md)|  | 
 **format** | **optional.String**|  | [default to MXLF]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPreviewFile**
> GetPreviewFile(ctx, optional)
Download preview

Supports mxliff format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BilingualFileApiGetPreviewFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BilingualFileApiGetPreviewFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBilingualFile**
> JobPartsDto UploadBilingualFile(ctx, optional)
Upload bilingual file

Returns updated job parts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BilingualFileApiUploadBilingualFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BilingualFileApiUploadBilingualFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 
 **format** | **optional.String**|  | [default to MXLF]
 **saveToTransMemory** | **optional.String**|  | [default to Confirmed]
 **setCompleted** | **optional.Bool**|  | [default to false]

### Return type

[**JobPartsDto**](JobPartsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

