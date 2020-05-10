# \WebhookApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebHook**](WebhookApi.md#CreateWebHook) | **Post** /api2/v1/webhooks | Create webhook
[**CreateWebHook1**](WebhookApi.md#CreateWebHook1) | **Post** /api2/v2/webhooks | Create webhook
[**DeleteWebHook**](WebhookApi.md#DeleteWebHook) | **Delete** /api2/v1/webhooks/{webHookId} | Delete webhook
[**DeleteWebHook1**](WebhookApi.md#DeleteWebHook1) | **Delete** /api2/v2/webhooks/{webHookId} | Delete webhook
[**GetWebHook**](WebhookApi.md#GetWebHook) | **Get** /api2/v1/webhooks/{webHookId} | Get webhook
[**GetWebHook1**](WebhookApi.md#GetWebHook1) | **Get** /api2/v2/webhooks/{webHookId} | Get webhook
[**GetWebHookList**](WebhookApi.md#GetWebHookList) | **Get** /api2/v1/webhooks | Lists webhooks
[**GetWebHookList1**](WebhookApi.md#GetWebHookList1) | **Get** /api2/v2/webhooks | Lists webhooks
[**UpdateWebHook**](WebhookApi.md#UpdateWebHook) | **Put** /api2/v1/webhooks/{webHookId} | Edit webhook
[**UpdateWebHook1**](WebhookApi.md#UpdateWebHook1) | **Put** /api2/v2/webhooks/{webHookId} | Edit webhook


# **CreateWebHook**
> WebHookDto CreateWebHook(ctx, optional)
Create webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhookApiCreateWebHookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiCreateWebHookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebHookDto**](WebHookDto.md)|  | 

### Return type

[**WebHookDto**](WebHookDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWebHook1**
> WebHookDtoV2 CreateWebHook1(ctx, optional)
Create webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhookApiCreateWebHook1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiCreateWebHook1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebHookDtoV2**](WebHookDtoV2.md)|  | 

### Return type

[**WebHookDtoV2**](WebHookDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebHook**
> DeleteWebHook(ctx, webHookId)
Delete webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webHookId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebHook1**
> DeleteWebHook1(ctx, webHookId)
Delete webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webHookId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebHook**
> WebHookDto GetWebHook(ctx, webHookId)
Get webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webHookId** | **int64**|  | 

### Return type

[**WebHookDto**](WebHookDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebHook1**
> WebHookDtoV2 GetWebHook1(ctx, webHookId)
Get webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webHookId** | **int64**|  | 

### Return type

[**WebHookDtoV2**](WebHookDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebHookList**
> PageDtoWebHookDto GetWebHookList(ctx, optional)
Lists webhooks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhookApiGetWebHookListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiGetWebHookListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoWebHookDto**](PageDtoWebHookDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebHookList1**
> PageDtoWebHookDtoV2 GetWebHookList1(ctx, optional)
Lists webhooks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhookApiGetWebHookList1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiGetWebHookList1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoWebHookDtoV2**](PageDtoWebHookDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebHook**
> WebHookDto UpdateWebHook(ctx, webHookId, optional)
Edit webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webHookId** | **int64**|  | 
 **optional** | ***WebhookApiUpdateWebHookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiUpdateWebHookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WebHookDto**](WebHookDto.md)|  | 

### Return type

[**WebHookDto**](WebHookDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebHook1**
> WebHookDtoV2 UpdateWebHook1(ctx, webHookId, optional)
Edit webhook



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webHookId** | **int64**|  | 
 **optional** | ***WebhookApiUpdateWebHook1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiUpdateWebHook1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WebHookDtoV2**](WebHookDtoV2.md)|  | 

### Return type

[**WebHookDtoV2**](WebHookDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

