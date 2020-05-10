# \WorkflowStepApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListWFSteps**](WorkflowStepApi.md#ListWFSteps) | **Get** /api2/v1/workflowSteps | List workflow steps
[**ListWorkflowSteps**](WorkflowStepApi.md#ListWorkflowSteps) | **Get** /api2/v1/users/{userId}/workflowSteps | List assigned workflow steps


# **ListWFSteps**
> PageDtoWorkflowStepDto ListWFSteps(ctx, optional)
List workflow steps



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowStepApiListWFStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowStepApiListWFStepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoWorkflowStepDto**](PageDtoWorkflowStepDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkflowSteps**
> PageDtoWorkflowStepReference ListWorkflowSteps(ctx, userId, optional)
List assigned workflow steps



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**|  | 
 **optional** | ***WorkflowStepApiListWorkflowStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowStepApiListWorkflowStepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of []string**](string.md)|  | 
 **projectUid** | **optional.String**|  | 
 **targetLang** | [**optional.Interface of []string**](string.md)|  | 
 **dueInHours** | **optional.Int32**| -1 for jobs that are overdue | 
 **filename** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]

### Return type

[**PageDtoWorkflowStepReference**](PageDtoWorkflowStepReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

