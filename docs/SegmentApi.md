# \SegmentApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSegmentsCount**](SegmentApi.md#GetSegmentsCount) | **Post** /api2/v1/projects/{projectUid}/jobs/segmentsCount | Get segments count
[**ListSegments**](SegmentApi.md#ListSegments) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/segments | Get segments


# **GetSegmentsCount**
> SegmentsCountsResponseListDto GetSegmentsCount(ctx, projectUid, optional)
Get segments count

Provides segments count (progress data)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***SegmentApiGetSegmentsCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SegmentApiGetSegmentsCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartReadyReferences**](JobPartReadyReferences.md)|  | 

### Return type

[**SegmentsCountsResponseListDto**](SegmentsCountsResponseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSegments**
> SegmentListDto ListSegments(ctx, projectUid, jobUid, optional)
Get segments



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***SegmentApiListSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SegmentApiListSegmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **beginIndex** | **optional.Int32**|  | [default to 0]
 **endIndex** | **optional.Int32**|  | [default to 0]

### Return type

[**SegmentListDto**](SegmentListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

