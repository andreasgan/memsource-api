# \MappingApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMappingForTask**](MappingApi.md#GetMappingForTask) | **Get** /api2/v1/mappings/tasks/{id} | Returns mapping for taskId (mxliff)


# **GetMappingForTask**
> TaskMappingDto GetMappingForTask(ctx, id, optional)
Returns mapping for taskId (mxliff)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***MappingApiGetMappingForTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MappingApiGetMappingForTaskOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowLevel** | **optional.Int32**|  | [default to 1]

### Return type

[**TaskMappingDto**](TaskMappingDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

