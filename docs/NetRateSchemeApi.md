# \NetRateSchemeApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDiscountScheme**](NetRateSchemeApi.md#CreateDiscountScheme) | **Post** /api2/v1/netRateSchemes | Create net rate scheme
[**GetDiscountScheme**](NetRateSchemeApi.md#GetDiscountScheme) | **Get** /api2/v1/netRateSchemes/{netRateSchemeId} | Get net rate scheme
[**GetDiscountSchemeWorkflowStep**](NetRateSchemeApi.md#GetDiscountSchemeWorkflowStep) | **Get** /api2/v1/netRateSchemes/{netRateSchemeId}/workflowStepNetSchemes/{netRateSchemeWorkflowStepId} | Get scheme for workflow step
[**GetDiscountSchemeWorkflowSteps**](NetRateSchemeApi.md#GetDiscountSchemeWorkflowSteps) | **Get** /api2/v1/netRateSchemes/{netRateSchemeId}/workflowStepNetSchemes | List schemes for workflow step
[**GetDiscountSchemes**](NetRateSchemeApi.md#GetDiscountSchemes) | **Get** /api2/v1/netRateSchemes | List net rate schemes


# **CreateDiscountScheme**
> NetRateScheme CreateDiscountScheme(ctx, optional)
Create net rate scheme



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetRateSchemeApiCreateDiscountSchemeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetRateSchemeApiCreateDiscountSchemeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DiscountSchemeCreateDto**](DiscountSchemeCreateDto.md)|  | 

### Return type

[**NetRateScheme**](NetRateScheme.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiscountScheme**
> NetRateScheme GetDiscountScheme(ctx, netRateSchemeId)
Get net rate scheme



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netRateSchemeId** | **int64**|  | 

### Return type

[**NetRateScheme**](NetRateScheme.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiscountSchemeWorkflowStep**
> NetRateSchemeWorkflowStep GetDiscountSchemeWorkflowStep(ctx, netRateSchemeId, netRateSchemeWorkflowStepId)
Get scheme for workflow step



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netRateSchemeId** | **int64**|  | 
  **netRateSchemeWorkflowStepId** | **int64**|  | 

### Return type

[**NetRateSchemeWorkflowStep**](NetRateSchemeWorkflowStep.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiscountSchemeWorkflowSteps**
> PageDtoNetRateSchemeWorkflowStepReference GetDiscountSchemeWorkflowSteps(ctx, netRateSchemeId, optional)
List schemes for workflow step



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **netRateSchemeId** | **int64**|  | 
 **optional** | ***NetRateSchemeApiGetDiscountSchemeWorkflowStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetRateSchemeApiGetDiscountSchemeWorkflowStepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoNetRateSchemeWorkflowStepReference**](PageDtoNetRateSchemeWorkflowStepReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiscountSchemes**
> PageDtoNetRateSchemeReference GetDiscountSchemes(ctx, optional)
List net rate schemes



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetRateSchemeApiGetDiscountSchemesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetRateSchemeApiGetDiscountSchemesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoNetRateSchemeReference**](PageDtoNetRateSchemeReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

