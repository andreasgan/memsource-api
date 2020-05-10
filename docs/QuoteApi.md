# \QuoteApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuote**](QuoteApi.md#CreateQuote) | **Post** /api2/v1/quotes | Create quote
[**CreateQuoteV2**](QuoteApi.md#CreateQuoteV2) | **Post** /api2/v2/quotes | Create quote
[**DeleteQuote**](QuoteApi.md#DeleteQuote) | **Delete** /api2/v1/quotes/{quoteId} | Delete quote
[**Get1**](QuoteApi.md#Get1) | **Get** /api2/v1/quotes/{quoteId} | Get quote


# **CreateQuote**
> QuoteDto CreateQuote(ctx, optional)
Create quote



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuoteApiCreateQuoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuoteApiCreateQuoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of QuoteCreateDto**](QuoteCreateDto.md)|  | 

### Return type

[**QuoteDto**](QuoteDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateQuoteV2**
> QuoteV2Dto CreateQuoteV2(ctx, optional)
Create quote

Either WorkflowSettings or Units must be sent for billingUnit \"Hour\".

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuoteApiCreateQuoteV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuoteApiCreateQuoteV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of QuoteCreateV2Dto**](QuoteCreateV2Dto.md)|  | 

### Return type

[**QuoteV2Dto**](QuoteV2Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuote**
> DeleteQuote(ctx, quoteId)
Delete quote



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get1**
> QuoteDto Get1(ctx, quoteId)
Get quote



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quoteId** | **int64**|  | 

### Return type

[**QuoteDto**](QuoteDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

