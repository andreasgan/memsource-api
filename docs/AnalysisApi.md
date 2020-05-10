# \AnalysisApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalysesBatchEditV2**](AnalysisApi.md#AnalysesBatchEditV2) | **Put** /api2/v2/analyses/bulk | Edit analyses (batch)
[**BatchEditAnalyses**](AnalysisApi.md#BatchEditAnalyses) | **Put** /api2/v1/analyses/bulk | Edit analyses (batch)
[**BulkDeleteAnalyses**](AnalysisApi.md#BulkDeleteAnalyses) | **Delete** /api2/v1/analyses/bulk | Delete analyses (batch)
[**CreateAnalyseAsync**](AnalysisApi.md#CreateAnalyseAsync) | **Post** /api2/v1/analyses | Create analysis
[**CreateAnalyseAsync1**](AnalysisApi.md#CreateAnalyseAsync1) | **Post** /api2/v2/analyses | Create analysis
[**CreateAnalysesForLangs**](AnalysisApi.md#CreateAnalysesForLangs) | **Post** /api2/v1/analyses/byLanguages | Create analyses by languages
[**CreateAnalysesForProviders**](AnalysisApi.md#CreateAnalysesForProviders) | **Post** /api2/v1/analyses/byProviders | Create analyses by providers
[**Delete**](AnalysisApi.md#Delete) | **Delete** /api2/v1/analyses/{analyseId} | Delete analysis
[**DownloadAnalyse**](AnalysisApi.md#DownloadAnalyse) | **Get** /api2/v1/analyses/{analyseId}/download | Download analysis
[**GetAnalyse**](AnalysisApi.md#GetAnalyse) | **Get** /api2/v1/analyses/{analyseId} | Get analysis
[**GetAnalyseLanguagePart**](AnalysisApi.md#GetAnalyseLanguagePart) | **Get** /api2/v1/analyses/{analyseId}/analyseLanguageParts/{analyseLanguagePartId} | Get analysis language part
[**GetAnalyseV2**](AnalysisApi.md#GetAnalyseV2) | **Get** /api2/v2/analyses/{analyseId} | Get analysis
[**GetAnalyseV3**](AnalysisApi.md#GetAnalyseV3) | **Get** /api2/v3/analyses/{analyseId} | Get analysis
[**GetJobPartAnalyse**](AnalysisApi.md#GetJobPartAnalyse) | **Get** /api2/v1/analyses/{analyseId}/jobs/{jobUid} | Get jobs analysis
[**ListByProject**](AnalysisApi.md#ListByProject) | **Get** /api2/v1/projects/{projectUid}/analyses | List analyses by project
[**ListByProjectV2**](AnalysisApi.md#ListByProjectV2) | **Get** /api2/v2/projects/{projectUid}/analyses | List analyses by project
[**ListJobParts**](AnalysisApi.md#ListJobParts) | **Get** /api2/v1/analyses/{analyseId}/analyseLanguageParts/{analyseLanguagePartId}/jobs | List jobs of analyses
[**ListPartAnalyse**](AnalysisApi.md#ListPartAnalyse) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/analyses | List analyses
[**ListPartAnalyse1**](AnalysisApi.md#ListPartAnalyse1) | **Get** /api2/v2/projects/{projectUid}/jobs/{jobUid}/analyses | List analyses


# **AnalysesBatchEditV2**
> AnalysesV2Dto AnalysesBatchEditV2(ctx, optional)
Edit analyses (batch)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnalysisApiAnalysesBatchEditV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiAnalysesBatchEditV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BulkEditAnalyseV2Dto**](BulkEditAnalyseV2Dto.md)|  | 

### Return type

[**AnalysesV2Dto**](AnalysesV2Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchEditAnalyses**
> []AnalyseDto BatchEditAnalyses(ctx, optional)
Edit analyses (batch)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnalysisApiBatchEditAnalysesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiBatchEditAnalysesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BulkEditAnalyseDto**](BulkEditAnalyseDto.md)|  | 

### Return type

[**[]AnalyseDto**](AnalyseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkDeleteAnalyses**
> BulkDeleteAnalyses(ctx, optional)
Delete analyses (batch)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnalysisApiBulkDeleteAnalysesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiBulkDeleteAnalysesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BulkDeleteAnalyseDto**](BulkDeleteAnalyseDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAnalyseAsync**
> AsyncAnalyseResponseDto CreateAnalyseAsync(ctx, optional)
Create analysis



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnalysisApiCreateAnalyseAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiCreateAnalyseAsyncOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateAnalyseAsyncDto**](CreateAnalyseAsyncDto.md)|  | 

### Return type

[**AsyncAnalyseResponseDto**](AsyncAnalyseResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAnalyseAsync1**
> AsyncAnalyseListResponseV2Dto CreateAnalyseAsync1(ctx, optional)
Create analysis

Returns created analyses - batching analyses by number of segments (api.segment.count.approximation, default 100000), in case request contains more segments than maximum (api.segment.max.count, default 300000), returns 400 bad request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnalysisApiCreateAnalyseAsync1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiCreateAnalyseAsync1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateAnalyseAsyncV2Dto**](CreateAnalyseAsyncV2Dto.md)|  | 

### Return type

[**AsyncAnalyseListResponseV2Dto**](AsyncAnalyseListResponseV2Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAnalysesForLangs**
> AsyncAnalyseListResponseDto CreateAnalysesForLangs(ctx, optional)
Create analyses by languages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnalysisApiCreateAnalysesForLangsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiCreateAnalysesForLangsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateAnalyseListAsyncDto**](CreateAnalyseListAsyncDto.md)|  | 

### Return type

[**AsyncAnalyseListResponseDto**](AsyncAnalyseListResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAnalysesForProviders**
> AsyncAnalyseListResponseDto CreateAnalysesForProviders(ctx, optional)
Create analyses by providers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AnalysisApiCreateAnalysesForProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiCreateAnalysesForProvidersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateAnalyseListAsyncDto**](CreateAnalyseListAsyncDto.md)|  | 

### Return type

[**AsyncAnalyseListResponseDto**](AsyncAnalyseListResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, analyseId, optional)
Delete analysis



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **analyseId** | **string**|  | 
 **optional** | ***AnalysisApiDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadAnalyse**
> DownloadAnalyse(ctx, analyseId, format)
Download analysis



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **analyseId** | **int64**|  | 
  **format** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnalyse**
> AnalyseDto GetAnalyse(ctx, analyseId)
Get analysis



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **analyseId** | **int64**|  | 

### Return type

[**AnalyseDto**](AnalyseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnalyseLanguagePart**
> AnalyseLanguagePartDto GetAnalyseLanguagePart(ctx, analyseId, analyseLanguagePartId)
Get analysis language part

Returns analysis language pair

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **analyseId** | **int64**|  | 
  **analyseLanguagePartId** | **int64**|  | 

### Return type

[**AnalyseLanguagePartDto**](AnalyseLanguagePartDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnalyseV2**
> AnalyseV2Dto GetAnalyseV2(ctx, analyseId)
Get analysis



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **analyseId** | **int64**|  | 

### Return type

[**AnalyseV2Dto**](AnalyseV2Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnalyseV3**
> AnalyseV3Dto GetAnalyseV3(ctx, analyseId)
Get analysis



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **analyseId** | **int64**|  | 

### Return type

[**AnalyseV3Dto**](AnalyseV3Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobPartAnalyse**
> AnalyseJobDto GetJobPartAnalyse(ctx, analyseId, jobUid)
Get jobs analysis

Returns job's analyse

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **analyseId** | **int64**|  | 
  **jobUid** | **string**|  | 

### Return type

[**AnalyseJobDto**](AnalyseJobDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByProject**
> PageDtoAnalyseDto ListByProject(ctx, projectUid, optional)
List analyses by project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***AnalysisApiListByProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiListByProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoAnalyseDto**](PageDtoAnalyseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByProjectV2**
> PageDtoAnalyseV2Dto ListByProjectV2(ctx, projectUid, optional)
List analyses by project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***AnalysisApiListByProjectV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiListByProjectV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoAnalyseV2Dto**](PageDtoAnalyseV2Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJobParts**
> PageDtoAnalyseJobDto ListJobParts(ctx, analyseId, analyseLanguagePartId, optional)
List jobs of analyses

Returns list of job's analyses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **analyseId** | **int64**|  | 
  **analyseLanguagePartId** | **int64**|  | 
 **optional** | ***AnalysisApiListJobPartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiListJobPartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoAnalyseJobDto**](PageDtoAnalyseJobDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPartAnalyse**
> PageDtoAnalyseDto ListPartAnalyse(ctx, projectUid, jobUid, optional)
List analyses



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***AnalysisApiListPartAnalyseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiListPartAnalyseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]

### Return type

[**PageDtoAnalyseDto**](PageDtoAnalyseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPartAnalyse1**
> PageDtoAnalyseV2Dto ListPartAnalyse1(ctx, projectUid, jobUid, optional)
List analyses



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***AnalysisApiListPartAnalyse1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AnalysisApiListPartAnalyse1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]

### Return type

[**PageDtoAnalyseV2Dto**](PageDtoAnalyseV2Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

