# \TranslationApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HumanTranslate**](TranslationApi.md#HumanTranslate) | **Post** /api2/v1/projects/{projectUid}/jobs/humanTranslate | Human translate (Gengo or Unbabel)
[**MachineTranslationJob**](TranslationApi.md#MachineTranslationJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/translations/translateWithMachineTranslation | Translate using machine translation
[**PreTranslate**](TranslationApi.md#PreTranslate) | **Post** /api2/v1/projects/{projectUid}/jobs/preTranslate | Pre-translate job


# **HumanTranslate**
> AsyncRequestWrapperDto HumanTranslate(ctx, projectUid, optional)
Human translate (Gengo or Unbabel)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***TranslationApiHumanTranslateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationApiHumanTranslateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of HumanTranslateJobsDto**](HumanTranslateJobsDto.md)|  | 

### Return type

[**AsyncRequestWrapperDto**](AsyncRequestWrapperDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MachineTranslationJob**
> MachineTranslateResponse MachineTranslationJob(ctx, projectUid, jobUid, optional)
Translate using machine translation

Configured machine translate settings is used

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***TranslationApiMachineTranslationJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationApiMachineTranslationJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TranslationRequestDto**](TranslationRequestDto.md)|  | 

### Return type

[**MachineTranslateResponse**](MachineTranslateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreTranslate**
> AsyncRequestWrapperDto PreTranslate(ctx, projectUid, optional)
Pre-translate job



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***TranslationApiPreTranslateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationApiPreTranslateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PreTranslateJobsDto**](PreTranslateJobsDto.md)|  | 

### Return type

[**AsyncRequestWrapperDto**](AsyncRequestWrapperDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

