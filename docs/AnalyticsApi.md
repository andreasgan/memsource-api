# \AnalyticsApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Aggregate**](AnalyticsApi.md#Aggregate) | **Post** /api2/v1/analytics/{type} | Run analytics aggregation


# **Aggregate**
> AnalyticsResponseDto Aggregate(ctx, type_, optional)
Run analytics aggregation

For more information check <a target=\"_blank\" href=\"https://wiki.memsource.com/wiki/Analytics_Aggregations\">Memsource Wiki</a>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| This value is case sensitive | 
 **optional** | ***AnalyticsApiAggregateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalyticsApiAggregateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AnalyticsDto**](AnalyticsDto.md)|  | 

### Return type

[**AnalyticsResponseDto**](AnalyticsResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

