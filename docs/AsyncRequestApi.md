# \AsyncRequestApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAsyncRequest**](AsyncRequestApi.md#GetAsyncRequest) | **Get** /api2/v1/async/{asyncRequestId} | Get asynchronous request
[**ListPendingRequests**](AsyncRequestApi.md#ListPendingRequests) | **Get** /api2/v1/async | List pending requests


# **GetAsyncRequest**
> AsyncRequestDto GetAsyncRequest(ctx, asyncRequestId)
Get asynchronous request



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asyncRequestId** | **int64**|  | 

### Return type

[**AsyncRequestDto**](AsyncRequestDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPendingRequests**
> PageDtoAsyncRequestDto ListPendingRequests(ctx, optional)
List pending requests



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AsyncRequestApiListPendingRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsyncRequestApiListPendingRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **optional.Bool**| Pending requests for organization instead of current user. Only for ADMIN. | [default to false]
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoAsyncRequestDto**](PageDtoAsyncRequestDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

