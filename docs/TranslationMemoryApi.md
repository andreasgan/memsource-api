# \TranslationMemoryApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTargetLangToTransMemory**](TranslationMemoryApi.md#AddTargetLangToTransMemory) | **Post** /api2/v1/transMemories/{transMemoryId}/targetLanguages | Add target language to translation memory
[**ClearTransMemory**](TranslationMemoryApi.md#ClearTransMemory) | **Delete** /api2/v1/transMemories/{transMemoryId}/segments | Delete all segments
[**CreateTransMemory**](TranslationMemoryApi.md#CreateTransMemory) | **Post** /api2/v1/transMemories | Create translation memory
[**DeleteSourceAndTranslations**](TranslationMemoryApi.md#DeleteSourceAndTranslations) | **Delete** /api2/v1/transMemories/{transMemoryId}/segments/{segmentId} | Delete both source and translation
[**DeleteTransMemory**](TranslationMemoryApi.md#DeleteTransMemory) | **Delete** /api2/v1/transMemories/{transMemoryId} | Delete translation memory
[**DeleteTranslation**](TranslationMemoryApi.md#DeleteTranslation) | **Delete** /api2/v1/transMemories/{transMemoryId}/segments/{segmentId}/lang/{lang} | Delete segment of given language
[**DownloadSearchResult**](TranslationMemoryApi.md#DownloadSearchResult) | **Get** /api2/v1/transMemories/downloadExport/{asyncRequestId} | Download export
[**EditTransMemory**](TranslationMemoryApi.md#EditTransMemory) | **Put** /api2/v1/transMemories/{transMemoryId} | Edit translation memory
[**Export**](TranslationMemoryApi.md#Export) | **Get** /api2/v1/transMemories/{transMemoryId}/export | Export translation memory
[**ExportByQueryAsync**](TranslationMemoryApi.md#ExportByQueryAsync) | **Post** /api2/v1/transMemories/{transMemoryId}/exportByQueryAsync | Search translation memory
[**ExportV2**](TranslationMemoryApi.md#ExportV2) | **Post** /api2/v2/transMemories/{transMemoryId}/export | Export translation memory
[**GetBackgroundTasks**](TranslationMemoryApi.md#GetBackgroundTasks) | **Get** /api2/v1/transMemories/{transMemoryId}/lastBackgroundTask | Get last task information
[**GetMetadata**](TranslationMemoryApi.md#GetMetadata) | **Get** /api2/v1/transMemories/{transMemoryId}/metadata | Get translation memory metadata
[**GetProjectTemplateTransMemories**](TranslationMemoryApi.md#GetProjectTemplateTransMemories) | **Get** /api2/v1/projectTemplates/{projectTemplateId}/transMemories | Get translation memories
[**GetRelatedProjects**](TranslationMemoryApi.md#GetRelatedProjects) | **Get** /api2/v1/transMemories/{transMemoryId}/relatedProjects | List related projects
[**GetTransMemory**](TranslationMemoryApi.md#GetTransMemory) | **Get** /api2/v1/transMemories/{transMemoryId} | Get translation memory
[**GetTranslationResources**](TranslationMemoryApi.md#GetTranslationResources) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/translationResources | Get translation resources
[**ImportTransMemory**](TranslationMemoryApi.md#ImportTransMemory) | **Post** /api2/v1/transMemories/{transMemoryId}/import | Import segments
[**InsertToTransMemory**](TranslationMemoryApi.md#InsertToTransMemory) | **Post** /api2/v1/transMemories/{transMemoryId}/segments | Insert segment
[**ListTransMemories**](TranslationMemoryApi.md#ListTransMemories) | **Get** /api2/v1/transMemories | List translation memories
[**Search**](TranslationMemoryApi.md#Search) | **Post** /api2/v1/transMemories/{transMemoryId}/search | Search translation memory (sync)
[**SearchByJob**](TranslationMemoryApi.md#SearchByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/transMemories/search | Search job&#39;s translation memories
[**SearchByJob3**](TranslationMemoryApi.md#SearchByJob3) | **Post** /api2/v3/projects/{projectUid}/jobs/{jobUid}/transMemories/search | Search job&#39;s translation memories
[**SearchSegment**](TranslationMemoryApi.md#SearchSegment) | **Post** /api2/v1/projects/{projectUid}/transMemories/searchSegmentInProject | Search translation memory for segment in the project
[**SearchSegmentByJob**](TranslationMemoryApi.md#SearchSegmentByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/transMemories/searchSegment | Search translation memory for segment by job
[**SetProjectTemplateTransMemories**](TranslationMemoryApi.md#SetProjectTemplateTransMemories) | **Put** /api2/v1/projectTemplates/{projectTemplateId}/transMemories | Edit translation memories
[**UpdateTranslation**](TranslationMemoryApi.md#UpdateTranslation) | **Put** /api2/v1/transMemories/{transMemoryId}/segments/{segmentId} | Edit segment
[**WildCardSearchByJob**](TranslationMemoryApi.md#WildCardSearchByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/transMemories/wildCardSearch | Wildcard search job&#39;s translation memories
[**WildCardSearchByJob2**](TranslationMemoryApi.md#WildCardSearchByJob2) | **Post** /api2/v2/projects/{projectUid}/jobs/{jobUid}/transMemories/wildCardSearch | Wildcard search job&#39;s translation memories
[**WildCardSearchByJob3**](TranslationMemoryApi.md#WildCardSearchByJob3) | **Post** /api2/v3/projects/{projectUid}/jobs/{jobUid}/transMemories/wildCardSearch | Wildcard search job&#39;s translation memories
[**WildcardSearch**](TranslationMemoryApi.md#WildcardSearch) | **Post** /api2/v1/transMemories/{transMemoryId}/wildCardSearch | Wildcard search


# **AddTargetLangToTransMemory**
> TransMemoryDto AddTargetLangToTransMemory(ctx, transMemoryId, optional)
Add target language to translation memory



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiAddTargetLangToTransMemoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiAddTargetLangToTransMemoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TargetLanguageDto**](TargetLanguageDto.md)|  | 

### Return type

[**TransMemoryDto**](TransMemoryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearTransMemory**
> ClearTransMemory(ctx, transMemoryId)
Delete all segments



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTransMemory**
> TransMemoryDto CreateTransMemory(ctx, optional)
Create translation memory



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TranslationMemoryApiCreateTransMemoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiCreateTransMemoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TransMemoryCreateDto**](TransMemoryCreateDto.md)|  | 

### Return type

[**TransMemoryDto**](TransMemoryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSourceAndTranslations**
> DeleteSourceAndTranslations(ctx, transMemoryId, segmentId)
Delete both source and translation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
  **segmentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransMemory**
> DeleteTransMemory(ctx, transMemoryId, optional)
Delete translation memory



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiDeleteTransMemoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiDeleteTransMemoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTranslation**
> DeleteTranslation(ctx, transMemoryId, segmentId, lang)
Delete segment of given language



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
  **segmentId** | **string**|  | 
  **lang** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadSearchResult**
> DownloadSearchResult(ctx, asyncRequestId, optional)
Download export



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asyncRequestId** | **string**| Request ID | 
 **optional** | ***TranslationMemoryApiDownloadSearchResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiDownloadSearchResultOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | [default to TMX]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTransMemory**
> TransMemoryDto EditTransMemory(ctx, transMemoryId, optional)
Edit translation memory



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiEditTransMemoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiEditTransMemoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TransMemoryEditDto**](TransMemoryEditDto.md)|  | 

### Return type

[**TransMemoryDto**](TransMemoryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Export**
> Export(ctx, transMemoryId, optional)
Export translation memory



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiExportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**|  | [default to TMX]
 **targetLang** | [**optional.Interface of []string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportByQueryAsync**
> AsyncExportTmByQueryResponseDto ExportByQueryAsync(ctx, transMemoryId, optional)
Search translation memory

Use [this API](#operation/downloadSearchResult) to download result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiExportByQueryAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiExportByQueryAsyncOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExportByQueryDto**](ExportByQueryDto.md)|  | 

### Return type

[**AsyncExportTmByQueryResponseDto**](AsyncExportTMByQueryResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportV2**
> AsyncExportTmResponseDto ExportV2(ctx, transMemoryId, optional)
Export translation memory

Use [this API](#operation/downloadSearchResult) to download result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiExportV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiExportV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExportTmDto**](ExportTmDto.md)|  | 

### Return type

[**AsyncExportTmResponseDto**](AsyncExportTMResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackgroundTasks**
> BackgroundTasksTmDto GetBackgroundTasks(ctx, transMemoryId)
Get last task information



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 

### Return type

[**BackgroundTasksTmDto**](BackgroundTasksTmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadata**
> MetadataResponse GetMetadata(ctx, transMemoryId, optional)
Get translation memory metadata



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiGetMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiGetMetadataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **byLanguage** | **optional.Bool**|  | [default to false]

### Return type

[**MetadataResponse**](MetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTemplateTransMemories**
> ProjectTemplateTransMemoryListDto GetProjectTemplateTransMemories(ctx, projectTemplateId)
Get translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 

### Return type

[**ProjectTemplateTransMemoryListDto**](ProjectTemplateTransMemoryListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelatedProjects**
> PageDtoAbstractProjectDto GetRelatedProjects(ctx, transMemoryId, optional)
List related projects



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiGetRelatedProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiGetRelatedProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoAbstractProjectDto**](PageDtoAbstractProjectDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransMemory**
> TransMemoryDto GetTransMemory(ctx, transMemoryId)
Get translation memory



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 

### Return type

[**TransMemoryDto**](TransMemoryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTranslationResources**
> TranslationResourcesDto GetTranslationResources(ctx, projectUid, jobUid)
Get translation resources



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

[**TranslationResourcesDto**](TranslationResourcesDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportTransMemory**
> ImportResponse ImportTransMemory(ctx, transMemoryId, contentDisposition, optional)
Import segments



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
  **contentDisposition** | **string**| must match pattern &#x60;((inline|attachment); )?filename\\*&#x3D;UTF-8&#39;&#39;(.+)&#x60; | 
 **optional** | ***TranslationMemoryApiImportTransMemoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiImportTransMemoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 
 **strictLangMatching** | **optional.Bool**|  | [default to false]
 **stripNativeCodes** | **optional.Bool**|  | [default to true]
 **excludeNotConfirmedSegments** | **optional.Bool**|  | [default to false]

### Return type

[**ImportResponse**](ImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InsertToTransMemory**
> InsertToTransMemory(ctx, transMemoryId, optional)
Insert segment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiInsertToTransMemoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiInsertToTransMemoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SegmentDto**](SegmentDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransMemories**
> PageDtoTransMemoryDto ListTransMemories(ctx, optional)
List translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TranslationMemoryApiListTransMemoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiListTransMemoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **sourceLang** | **optional.String**|  | 
 **targetLang** | **optional.String**|  | 
 **clientId** | **optional.String**|  | 
 **domainId** | **optional.String**|  | 
 **subDomainId** | **optional.String**|  | 
 **businessUnitId** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoTransMemoryDto**](PageDtoTransMemoryDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Search**
> SearchResponseListTmDto Search(ctx, transMemoryId, optional)
Search translation memory (sync)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SearchRequestDto**](SearchRequestDto.md)|  | 

### Return type

[**SearchResponseListTmDto**](SearchResponseListTmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchByJob**
> SearchResponseListTmDto SearchByJob(ctx, projectUid, jobUid, optional)
Search job's translation memories

<b>This API is incorrectly implemented, use             endpoint <a href=\"#operation/searchSegmentByJob\">Search translation memory for segment by job</a>.</b> </br>             Returns at most <i>maxSegments</i> records with <i>score >= scoreThreshold</i> and at most             <i>maxSubsegments</i> records which are subsegment, i.e. the source text is substring of the query text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***TranslationMemoryApiSearchByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiSearchByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTmByJobRequestDto**](SearchTmByJobRequestDto.md)|  | 

### Return type

[**SearchResponseListTmDto**](SearchResponseListTmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchByJob3**
> SearchResponseListTmDtoV3 SearchByJob3(ctx, projectUid, jobUid, optional)
Search job's translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***TranslationMemoryApiSearchByJob3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiSearchByJob3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTmByJobRequestDtoV3**](SearchTmByJobRequestDtoV3.md)|  | 

### Return type

[**SearchResponseListTmDtoV3**](SearchResponseListTmDtoV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSegment**
> SearchResponseListTmDto SearchSegment(ctx, projectUid, optional)
Search translation memory for segment in the project

Returns at most <i>maxSegments</i>             records with <i>score >= scoreThreshold</i> and at most <i>maxSubsegments</i> records which are subsegment,             i.e. the source text is substring of the query text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***TranslationMemoryApiSearchSegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiSearchSegmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SearchTmRequestDto**](SearchTmRequestDto.md)|  | 

### Return type

[**SearchResponseListTmDto**](SearchResponseListTmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSegmentByJob**
> SearchResponseListTmDto SearchSegmentByJob(ctx, projectUid, jobUid, optional)
Search translation memory for segment by job

Returns at most <i>maxSegments</i>             records with <i>score >= scoreThreshold</i> and at most <i>maxSubsegments</i> records which are subsegment,             i.e. the source text is substring of the query text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***TranslationMemoryApiSearchSegmentByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiSearchSegmentByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTmByJobRequestDto**](SearchTmByJobRequestDto.md)|  | 

### Return type

[**SearchResponseListTmDto**](SearchResponseListTmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetProjectTemplateTransMemories**
> ProjectTemplateTransMemoryListDto SetProjectTemplateTransMemories(ctx, projectTemplateId, optional)
Edit translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 
 **optional** | ***TranslationMemoryApiSetProjectTemplateTransMemoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiSetProjectTemplateTransMemoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SetProjectTemplateTransMemoriesDto**](SetProjectTemplateTransMemoriesDto.md)|  | 

### Return type

[**ProjectTemplateTransMemoryListDto**](ProjectTemplateTransMemoryListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTranslation**
> UpdateTranslation(ctx, transMemoryId, segmentId, optional)
Edit segment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
  **segmentId** | **string**|  | 
 **optional** | ***TranslationMemoryApiUpdateTranslationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiUpdateTranslationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TranslationDto**](TranslationDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WildCardSearchByJob**
> SearchResponseListTmDto WildCardSearchByJob(ctx, projectUid, jobUid, optional)
Wildcard search job's translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***TranslationMemoryApiWildCardSearchByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiWildCardSearchByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WildCardSearchByJobRequestDto**](WildCardSearchByJobRequestDto.md)|  | 

### Return type

[**SearchResponseListTmDto**](SearchResponseListTmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WildCardSearchByJob2**
> SearchResponseListTmDtoV2 WildCardSearchByJob2(ctx, projectUid, jobUid, optional)
Wildcard search job's translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***TranslationMemoryApiWildCardSearchByJob2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiWildCardSearchByJob2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WildCardSearchByJobRequestDtoV2**](WildCardSearchByJobRequestDtoV2.md)|  | 

### Return type

[**SearchResponseListTmDtoV2**](SearchResponseListTmDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WildCardSearchByJob3**
> SearchResponseListTmDtoV3 WildCardSearchByJob3(ctx, projectUid, jobUid, optional)
Wildcard search job's translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***TranslationMemoryApiWildCardSearchByJob3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiWildCardSearchByJob3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WildCardSearchByJobRequestDtoV3**](WildCardSearchByJobRequestDtoV3.md)|  | 

### Return type

[**SearchResponseListTmDtoV3**](SearchResponseListTmDtoV3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WildcardSearch**
> SearchResponseListTmDto WildcardSearch(ctx, transMemoryId, optional)
Wildcard search



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transMemoryId** | **int64**|  | 
 **optional** | ***TranslationMemoryApiWildcardSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TranslationMemoryApiWildcardSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WildCardSearchRequestDto**](WildCardSearchRequestDto.md)|  | 

### Return type

[**SearchResponseListTmDto**](SearchResponseListTmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

