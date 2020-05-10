# \QualityAssuranceApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIgnoredWarnings**](QualityAssuranceApi.md#AddIgnoredWarnings) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/qualityAssurances/ignoredWarnings | Add ignored warnings
[**DeleteIgnoredWarnings**](QualityAssuranceApi.md#DeleteIgnoredWarnings) | **Delete** /api2/v1/projects/{projectUid}/jobs/{jobUid}/qualityAssurances/ignoredWarnings | Delete ignored warnings
[**EnabledQualityChecksForJob**](QualityAssuranceApi.md#EnabledQualityChecksForJob) | **Get** /api2/v2/projects/{projectUid}/jobs/qualityAssurances/settings | Get QA settings
[**EnabledQualityChecksForJob1**](QualityAssuranceApi.md#EnabledQualityChecksForJob1) | **Get** /api2/v2/projects/{projectUid}/jobs/{jobUid}/qualityAssurances/settings | Get QA settings for job
[**RunQaForJobPart**](QualityAssuranceApi.md#RunQaForJobPart) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/qualityAssurances/run | Run quality assurance
[**RunQaForJobPartV2**](QualityAssuranceApi.md#RunQaForJobPartV2) | **Post** /api2/v2/projects/{projectUid}/jobs/{jobUid}/qualityAssurances/run | Run quality assurance
[**RunQaForJobPartV3**](QualityAssuranceApi.md#RunQaForJobPartV3) | **Post** /api2/v3/projects/{projectUid}/jobs/{jobUid}/qualityAssurances/run | Run quality assurance
[**RunQaForJobParts**](QualityAssuranceApi.md#RunQaForJobParts) | **Post** /api2/v1/projects/{projectUid}/jobs/qualityAssurances/run | Run quality assurance (batch)
[**RunQaForJobPartsV2**](QualityAssuranceApi.md#RunQaForJobPartsV2) | **Post** /api2/v2/projects/{projectUid}/jobs/qualityAssurances/run | Run quality assurance (batch)
[**RunQaForJobPartsV3**](QualityAssuranceApi.md#RunQaForJobPartsV3) | **Post** /api2/v3/projects/{projectUid}/jobs/qualityAssurances/run | Run quality assurance (batch)
[**RunQaForSegments**](QualityAssuranceApi.md#RunQaForSegments) | **Post** /api2/v2/projects/{projectUid}/jobs/qualityAssurances/segments/run | Run quality assurance on selected segments
[**RunQaForSegmentsV3**](QualityAssuranceApi.md#RunQaForSegmentsV3) | **Post** /api2/v3/projects/{projectUid}/jobs/qualityAssurances/segments/run | Run quality assurance on selected segments
[**UpdateIgnoredChecks**](QualityAssuranceApi.md#UpdateIgnoredChecks) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/qualityAssurances/ignoreChecks | Edit ignored checks


# **AddIgnoredWarnings**
> UpdateIgnoredWarningsDto AddIgnoredWarnings(ctx, projectUid, jobUid, optional)
Add ignored warnings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiAddIgnoredWarningsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiAddIgnoredWarningsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateIgnoredWarningsDto**](UpdateIgnoredWarningsDto.md)|  | 

### Return type

[**UpdateIgnoredWarningsDto**](UpdateIgnoredWarningsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIgnoredWarnings**
> DeleteIgnoredWarnings(ctx, projectUid, jobUid, optional)
Delete ignored warnings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiDeleteIgnoredWarningsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiDeleteIgnoredWarningsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateIgnoredWarningsDto**](UpdateIgnoredWarningsDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnabledQualityChecksForJob**
> QualityAssuranceChecksDto EnabledQualityChecksForJob(ctx, projectUid)
Get QA settings

Returns enabled quality assurance checks and settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 

### Return type

[**QualityAssuranceChecksDto**](QualityAssuranceChecksDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnabledQualityChecksForJob1**
> QualityAssuranceChecksDto EnabledQualityChecksForJob1(ctx, projectUid, jobUid)
Get QA settings for job

Returns enabled quality assurance checks and settings for job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

[**QualityAssuranceChecksDto**](QualityAssuranceChecksDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunQaForJobPart**
> QualityAssuranceResponse RunQaForJobPart(ctx, projectUid, jobUid, optional)
Run quality assurance



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiRunQaForJobPartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiRunQaForJobPartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of QualityAssuranceRunDto**](QualityAssuranceRunDto.md)|  | 

### Return type

[**QualityAssuranceResponse**](QualityAssuranceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunQaForJobPartV2**
> QualityAssuranceResponse RunQaForJobPartV2(ctx, projectUid, jobUid, optional)
Run quality assurance

Call \"Get QA settings\" endpoint to get the list of enabled QA checks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiRunQaForJobPartV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiRunQaForJobPartV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of QualityAssuranceRunDtoV2**](QualityAssuranceRunDtoV2.md)|  | 

### Return type

[**QualityAssuranceResponse**](QualityAssuranceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunQaForJobPartV3**
> QualityAssuranceResponseDto RunQaForJobPartV3(ctx, projectUid, jobUid, optional)
Run quality assurance

Call \"Get QA settings\" endpoint to get the list of enabled QA checks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiRunQaForJobPartV3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiRunQaForJobPartV3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of QualityAssuranceRunDtoV3**](QualityAssuranceRunDtoV3.md)|  | 

### Return type

[**QualityAssuranceResponseDto**](QualityAssuranceResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunQaForJobParts**
> QualityAssuranceResponse RunQaForJobParts(ctx, projectUid, optional)
Run quality assurance (batch)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiRunQaForJobPartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiRunQaForJobPartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QualityAssuranceBatchRunDto**](QualityAssuranceBatchRunDto.md)|  | 

### Return type

[**QualityAssuranceResponse**](QualityAssuranceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunQaForJobPartsV2**
> QualityAssuranceResponse RunQaForJobPartsV2(ctx, projectUid, optional)
Run quality assurance (batch)

Call \"Get QA settings\" endpoint to get the list of enabled QA checks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiRunQaForJobPartsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiRunQaForJobPartsV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QualityAssuranceBatchRunDtoV2**](QualityAssuranceBatchRunDtoV2.md)|  | 

### Return type

[**QualityAssuranceResponse**](QualityAssuranceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunQaForJobPartsV3**
> QualityAssuranceResponseDto RunQaForJobPartsV3(ctx, projectUid, optional)
Run quality assurance (batch)

Call \"Get QA settings\" endpoint to get the list of enabled QA checks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiRunQaForJobPartsV3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiRunQaForJobPartsV3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QualityAssuranceBatchRunDtoV3**](QualityAssuranceBatchRunDtoV3.md)|  | 

### Return type

[**QualityAssuranceResponseDto**](QualityAssuranceResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunQaForSegments**
> QualityAssuranceResponse RunQaForSegments(ctx, projectUid, optional)
Run quality assurance on selected segments

Source and target language of jobs have to match

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiRunQaForSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiRunQaForSegmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QualityAssuranceSegmentsRunDto**](QualityAssuranceSegmentsRunDto.md)|  | 

### Return type

[**QualityAssuranceResponse**](QualityAssuranceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunQaForSegmentsV3**
> QualityAssuranceResponseDto RunQaForSegmentsV3(ctx, projectUid, optional)
Run quality assurance on selected segments

By default runs only fast running checks. Source and target language of jobs have to match.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiRunQaForSegmentsV3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiRunQaForSegmentsV3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QualityAssuranceSegmentsRunDtoV3**](QualityAssuranceSegmentsRunDtoV3.md)|  | 

### Return type

[**QualityAssuranceResponseDto**](QualityAssuranceResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIgnoredChecks**
> UpdateIgnoredChecks(ctx, projectUid, jobUid, optional)
Edit ignored checks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***QualityAssuranceApiUpdateIgnoredChecksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualityAssuranceApiUpdateIgnoredChecksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateIgnoredChecksDto**](UpdateIgnoredChecksDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

