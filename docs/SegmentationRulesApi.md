# \SegmentationRulesApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegmentationRule**](SegmentationRulesApi.md#CreateSegmentationRule) | **Post** /api2/v1/segmentationRules | Create segmentation rule
[**DeletesSegmentationRule**](SegmentationRulesApi.md#DeletesSegmentationRule) | **Delete** /api2/v1/segmentationRules/{segRuleId} | Delete segmentation rule
[**GetListOfSegmentationRules**](SegmentationRulesApi.md#GetListOfSegmentationRules) | **Get** /api2/v1/segmentationRules | List segmentation rules
[**GetSegmentationRule**](SegmentationRulesApi.md#GetSegmentationRule) | **Get** /api2/v1/segmentationRules/{segRuleId} | Get segmentation rule
[**UpdatesSegmentationRule**](SegmentationRulesApi.md#UpdatesSegmentationRule) | **Put** /api2/v1/segmentationRules/{segRuleId} | Edit segmentation rule


# **CreateSegmentationRule**
> SegmentationRuleDto CreateSegmentationRule(ctx, body, segRule)
Create segmentation rule

Creates new Segmentation Rule with file and segRule JSON Object as header parameter. The same object is used for GET action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InputStream**](InputStream.md)| streamed file | 
  **segRule** | **string**|  | 

### Return type

[**SegmentationRuleDto**](SegmentationRuleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletesSegmentationRule**
> DeletesSegmentationRule(ctx, segRuleId)
Delete segmentation rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segRuleId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfSegmentationRules**
> PageDtoSegmentationRuleReference GetListOfSegmentationRules(ctx, optional)
List segmentation rules



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SegmentationRulesApiGetListOfSegmentationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SegmentationRulesApiGetListOfSegmentationRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoSegmentationRuleReference**](PageDtoSegmentationRuleReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentationRule**
> SegmentationRuleDto GetSegmentationRule(ctx, segRuleId)
Get segmentation rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segRuleId** | **int64**|  | 

### Return type

[**SegmentationRuleDto**](SegmentationRuleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatesSegmentationRule**
> SegmentationRuleDto UpdatesSegmentationRule(ctx, segRuleId, optional)
Edit segmentation rule



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segRuleId** | **int64**|  | 
 **optional** | ***SegmentationRulesApiUpdatesSegmentationRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SegmentationRulesApiUpdatesSegmentationRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EditSegmentationRuleDto**](EditSegmentationRuleDto.md)|  | 

### Return type

[**SegmentationRuleDto**](SegmentationRuleDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

