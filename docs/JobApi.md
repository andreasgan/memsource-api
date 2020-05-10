# \JobApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComparePart**](JobApi.md#ComparePart) | **Post** /api2/v1/projects/{projectUid}/jobs/compare | Compare jobs on workflow levels
[**CompletedFile**](JobApi.md#CompletedFile) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/targetFile | Download target file
[**CreateJob**](JobApi.md#CreateJob) | **Post** /api2/v1/projects/{projectUid}/jobs | Create job
[**CreateTermByJob**](JobApi.md#CreateTermByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/termBases/createByJob | Create term in job&#39;s term bases
[**DeleteAllTranslations**](JobApi.md#DeleteAllTranslations) | **Delete** /api2/v1/projects/{projectUid}/jobs/translations | Delete all translations
[**DeleteParts**](JobApi.md#DeleteParts) | **Delete** /api2/v1/projects/{projectUid}/jobs/batch | Delete job (batch)
[**EditPart**](JobApi.md#EditPart) | **Put** /api2/v1/projects/{projectUid}/jobs/{jobUid} | Edit job
[**EditParts**](JobApi.md#EditParts) | **Put** /api2/v1/projects/{projectUid}/jobs/batch | Edit jobs (batch)
[**FilePreview**](JobApi.md#FilePreview) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/preview | Download preview file
[**FilePreviewJob**](JobApi.md#FilePreviewJob) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/preview | Download preview file
[**GetBilingualFile**](JobApi.md#GetBilingualFile) | **Post** /api2/v1/projects/{projectUid}/jobs/bilingualFile | Download bilingual file
[**GetCompletedFileWarnings**](JobApi.md#GetCompletedFileWarnings) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/targetFileWarnings | Get target file&#39;s warnings
[**GetOriginalFile**](JobApi.md#GetOriginalFile) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/original | Download original file
[**GetPart**](JobApi.md#GetPart) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid} | Get job
[**GetPartsWorkflowStep**](JobApi.md#GetPartsWorkflowStep) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/workflowStep | Get job&#39;s workflowStep
[**GetSegmentsCount**](JobApi.md#GetSegmentsCount) | **Post** /api2/v1/projects/{projectUid}/jobs/segmentsCount | Get segments count
[**GetTranslationResources**](JobApi.md#GetTranslationResources) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/translationResources | Get translation resources
[**ListPartAnalyse**](JobApi.md#ListPartAnalyse) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/analyses | List analyses
[**ListPartAnalyse1**](JobApi.md#ListPartAnalyse1) | **Get** /api2/v2/projects/{projectUid}/jobs/{jobUid}/analyses | List analyses
[**ListParts**](JobApi.md#ListParts) | **Get** /api2/v1/projects/{projectUid}/jobs | List jobs
[**ListPartsV2**](JobApi.md#ListPartsV2) | **Get** /api2/v2/projects/{projectUid}/jobs | List jobs
[**ListProviders1**](JobApi.md#ListProviders1) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/providers/suggest | Get suggested providers
[**ListSegments**](JobApi.md#ListSegments) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/segments | Get segments
[**NotifyAssigned**](JobApi.md#NotifyAssigned) | **Post** /api2/v1/projects/{projectUid}/jobs/notifyAssigned | Notify assigned users
[**PreviewUrls**](JobApi.md#PreviewUrls) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/previewUrl | Get PDF preview
[**PseudoTranslate**](JobApi.md#PseudoTranslate) | **Post** /api2/v1/projects/{projectUid}/jobs/pseudoTranslate | Pseudo-translate job
[**PseudoTranslateJobPart**](JobApi.md#PseudoTranslateJobPart) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/pseudoTranslate | Pseudo-translates job
[**SearchByJob**](JobApi.md#SearchByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/transMemories/search | Search job&#39;s translation memories
[**SearchByJob2**](JobApi.md#SearchByJob2) | **Post** /api2/v2/projects/{projectUid}/jobs/{jobUid}/transMemories/search | Search job&#39;s translation memories
[**SearchByJob3**](JobApi.md#SearchByJob3) | **Post** /api2/v3/projects/{projectUid}/jobs/{jobUid}/transMemories/search | Search job&#39;s translation memories
[**SearchPartsInProject**](JobApi.md#SearchPartsInProject) | **Post** /api2/v1/projects/{projectUid}/jobs/search | Search jobs in project
[**SearchSegmentByJob**](JobApi.md#SearchSegmentByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/transMemories/searchSegment | Search translation memory for segment by job
[**SearchTermsByJob**](JobApi.md#SearchTermsByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/termBases/searchByJob | Search job&#39;s term bases
[**SearchTermsByJob1**](JobApi.md#SearchTermsByJob1) | **Post** /api2/v2/projects/{projectUid}/jobs/{jobUid}/termBases/searchByJob | Search job&#39;s term bases
[**SearchTermsInTextByJob**](JobApi.md#SearchTermsInTextByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/termBases/searchInTextByJob | Search terms in text
[**SearchTermsInTextByJobV2**](JobApi.md#SearchTermsInTextByJobV2) | **Post** /api2/v2/projects/{projectUid}/jobs/{jobUid}/termBases/searchInTextByJob | Search terms in text
[**SetStatus**](JobApi.md#SetStatus) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/setStatus | Edit job status
[**Split**](JobApi.md#Split) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/split | Split job
[**StatusChanges**](JobApi.md#StatusChanges) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/statusChanges | Get status changes
[**UploadBilingualFile**](JobApi.md#UploadBilingualFile) | **Put** /api2/v1/bilingualFiles | Upload bilingual file
[**WebEditorLink**](JobApi.md#WebEditorLink) | **Post** /api2/v1/projects/{projectUid}/jobs/webEditor | Get Web Editor URL
[**WebEditorLinkV2**](JobApi.md#WebEditorLinkV2) | **Post** /api2/v2/projects/{projectUid}/jobs/webEditor | Get Web Editor URL
[**WildCardSearchByJob**](JobApi.md#WildCardSearchByJob) | **Post** /api2/v1/projects/{projectUid}/jobs/{jobUid}/transMemories/wildCardSearch | Wildcard search job&#39;s translation memories
[**WildCardSearchByJob2**](JobApi.md#WildCardSearchByJob2) | **Post** /api2/v2/projects/{projectUid}/jobs/{jobUid}/transMemories/wildCardSearch | Wildcard search job&#39;s translation memories
[**WildCardSearchByJob3**](JobApi.md#WildCardSearchByJob3) | **Post** /api2/v3/projects/{projectUid}/jobs/{jobUid}/transMemories/wildCardSearch | Wildcard search job&#39;s translation memories


# **ComparePart**
> ComparedSegmentsDto ComparePart(ctx, projectUid, optional)
Compare jobs on workflow levels



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiComparePartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiComparePartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartReadyReferences**](JobPartReadyReferences.md)|  | 
 **atWorkflowLevel** | **optional.Int32**|  | [default to 1]
 **withWorkflowLevel** | **optional.Int32**|  | [default to 1]

### Return type

[**ComparedSegmentsDto**](ComparedSegmentsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompletedFile**
> CompletedFile(ctx, projectUid, jobUid, optional)
Download target file

This call will return target file with translation. This means even for other jobs that were created via 'split jobs' etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiCompletedFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiCompletedFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.String**|  | [default to ORIGINAL]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateJob**
> JobListDto CreateJob(ctx, projectUid, optional)
Create job

 Job file can be provided directly in the message body or downloaded from connector.   Please supply job metadata in `Memsource` header.   For file in the request body provide also the filename in `Content-Disposition` header.  Accepted metadata:     - `targetLangs` - **required**   - `due` - ISO 8601   - `workflowSettings` - project with workflow - see examples bellow   - `assignments` - project without workflows - see examples bellow   - `importSettings` - re-usable import settings - see [Create import settings](#operation/createImportSettings)   - `useProjectFileImportSettings` - mutually exclusive with importSettings   - `callbackUrl` - consumer callback   - `path` - original destination directory      for remote file jobs also `remoteFile` - see examples bellow:   - `connectorToken` - remote connector token   - `remoteFolder`    - `remoteFileName`   - `continuous` - true for continuous job  Create and assign job in project without workflow step: ```  {   \"targetLangs\": [     \"cs_cz\"   ],   \"callbackUrl\": \"https://my-shiny-service.com/consumeCallback\",   \"importSettings\": {     \"uid\": \"abcd123\"   },   \"due\": \"2007-12-03T10:15:30.00Z\",   \"path\": \"destination directory\",   \"assignments\": [     {       \"targetLang\": \"cs_cz\",       \"providers\": [         {           \"id\": \"4321\",           \"type\": \"USER\"         }       ]     }   ],   \"notifyProvider\": {     \"organizationEmailTemplate\": {       \"id\": \"39\"     },     \"notificationIntervalInMinutes\": \"10\"   } } ```  Create job from remote file without workflow steps: ```  {   \"remoteFile\": {     \"connectorToken\": \"948123ef-e1ef-4cd3-a90e-af1617848af3\",     \"remoteFolder\": \"/\",     \"remoteFileName\": \"Few words.docx\",     \"continuous\": false   },   \"assignments\": [],   \"workflowSettings\": [],   \"targetLangs\": [     \"cs\"   ] } ```  Create and assign job in project with workflow step: ```  {   \"targetLangs\": [     \"de\"   ],   \"useProjectFileImportSettings\": \"true\",   \"workflowSettings\": [     {       \"id\": \"64\",       \"due\": \"2007-12-03T10:15:30.00Z\",       \"assignments\": [         {           \"targetLang\": \"de\",           \"providers\": [             {               \"id\": \"3\",               \"type\": \"VENDOR\"             }           ]         }       ],       \"notifyProvider\": {         \"organizationEmailTemplate\": {           \"id\": \"39\"         },         \"notificationIntervalInMinutes\": \"10\"       }     }   ] } ```     

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiCreateJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiCreateJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memsource** | **optional.String**|  | 
 **contentDisposition** | **optional.String**| must match pattern &#x60;((inline|attachment); )?(filename\\*&#x3D;UTF-8&#39;&#39;(.+)|filename&#x3D;\&quot;?(.+)\&quot;?)&#x60; | 
 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 

### Return type

[**JobListDto**](JobListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTermByJob**
> TermPairDto CreateTermByJob(ctx, jobUid, projectUid, optional)
Create term in job's term bases

Create new term in the write term base assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***JobApiCreateTermByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiCreateTermByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CreateTermsDto**](CreateTermsDto.md)|  | 

### Return type

[**TermPairDto**](TermPairDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllTranslations**
> DeleteAllTranslations(ctx, projectUid, optional)
Delete all translations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiDeleteAllTranslationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiDeleteAllTranslationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartReadyReferences**](JobPartReadyReferences.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteParts**
> DeleteParts(ctx, projectUid, optional)
Delete job (batch)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiDeletePartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiDeletePartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartReferences**](JobPartReferences.md)|  | 
 **purge** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPart**
> JobPartExtendedDto EditPart(ctx, projectUid, jobUid, optional)
Edit job



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiEditPartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiEditPartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of JobPartUpdateSingleDto**](JobPartUpdateSingleDto.md)|  | 

### Return type

[**JobPartExtendedDto**](JobPartExtendedDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditParts**
> JobPartsDto EditParts(ctx, projectUid, optional)
Edit jobs (batch)

 Returns only jobs which were updated by the batch operation. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiEditPartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiEditPartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartUpdateBatchDto**](JobPartUpdateBatchDto.md)|  | 

### Return type

[**JobPartsDto**](JobPartsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilePreview**
> FilePreview(ctx, projectUid, jobUid, optional)
Download preview file

Takes bilingual file (.mxliff only) as argument. If not passed, data will be taken from database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiFilePreviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiFilePreviewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilePreviewJob**
> FilePreviewJob(ctx, projectUid, jobUid)
Download preview file



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBilingualFile**
> GetBilingualFile(ctx, projectUid, optional)
Download bilingual file



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiGetBilingualFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiGetBilingualFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartReadyReferences**](JobPartReadyReferences.md)|  | 
 **format** | **optional.String**|  | [default to MXLF]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompletedFileWarnings**
> TargetFileWarningsDto GetCompletedFileWarnings(ctx, projectUid, jobUid)
Get target file's warnings

 This call will return target file's warnings. This means even for other jobs that were created via 'split jobs' etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

[**TargetFileWarningsDto**](TargetFileWarningsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOriginalFile**
> GetOriginalFile(ctx, projectUid, jobUid)
Download original file



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPart**
> JobPartExtendedDto GetPart(ctx, projectUid, jobUid)
Get job



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

[**JobPartExtendedDto**](JobPartExtendedDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPartsWorkflowStep**
> ProjectWorkflowStepDto GetPartsWorkflowStep(ctx, projectUid, jobUid)
Get job's workflowStep



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

[**ProjectWorkflowStepDto**](ProjectWorkflowStepDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentsCount**
> SegmentsCountsResponseListDto GetSegmentsCount(ctx, projectUid, optional)
Get segments count

Provides segments count (progress data)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiGetSegmentsCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiGetSegmentsCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartReadyReferences**](JobPartReadyReferences.md)|  | 

### Return type

[**SegmentsCountsResponseListDto**](SegmentsCountsResponseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
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

# **ListPartAnalyse**
> PageDtoAnalyseDto ListPartAnalyse(ctx, projectUid, jobUid, optional)
List analyses



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiListPartAnalyseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiListPartAnalyseOpts struct

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
 **optional** | ***JobApiListPartAnalyse1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiListPartAnalyse1Opts struct

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

# **ListParts**
> PageDtoJobPartReference ListParts(ctx, projectUid, optional)
List jobs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiListPartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiListPartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]
 **count** | **optional.Bool**|  | [default to false]
 **workflowLevel** | **optional.Int32**|  | [default to 1]
 **status** | [**optional.Interface of []string**](string.md)|  | 
 **assignedLinguist** | **optional.Int64**|  | 
 **dueInHours** | **optional.Int32**|  | 
 **filename** | **optional.String**|  | 
 **targetLang** | **optional.String**|  | 

### Return type

[**PageDtoJobPartReference**](PageDtoJobPartReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPartsV2**
> PageDtoJobPartReferenceV2 ListPartsV2(ctx, projectUid, optional)
List jobs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiListPartsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiListPartsV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 50]
 **count** | **optional.Bool**|  | [default to false]
 **workflowLevel** | **optional.Int32**|  | [default to 1]
 **status** | [**optional.Interface of []string**](string.md)|  | 
 **assignedUser** | **optional.Int64**|  | 
 **dueInHours** | **optional.Int32**|  | 
 **filename** | **optional.String**|  | 
 **targetLang** | **optional.String**|  | 
 **assignedVendor** | **optional.Int64**|  | 

### Return type

[**PageDtoJobPartReferenceV2**](PageDtoJobPartReferenceV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProviders1**
> ProviderListDto ListProviders1(ctx, projectUid, jobUid)
Get suggested providers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

[**ProviderListDto**](ProviderListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSegments**
> SegmentListDto ListSegments(ctx, projectUid, jobUid, optional)
Get segments



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiListSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiListSegmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **beginIndex** | **optional.Int32**|  | [default to 0]
 **endIndex** | **optional.Int32**|  | [default to 0]

### Return type

[**SegmentListDto**](SegmentListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotifyAssigned**
> NotifyAssigned(ctx, projectUid, optional)
Notify assigned users



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiNotifyAssignedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiNotifyAssignedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NotifyJobPartsRequestDto**](NotifyJobPartsRequestDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreviewUrls**
> PreviewUrlsDto PreviewUrls(ctx, projectUid, jobUid, optional)
Get PDF preview



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiPreviewUrlsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiPreviewUrlsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workflowLevel** | **optional.Int32**|  | [default to 1]

### Return type

[**PreviewUrlsDto**](PreviewUrlsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PseudoTranslate**
> PseudoTranslate(ctx, projectUid, optional)
Pseudo-translate job



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiPseudoTranslateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiPseudoTranslateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JobPartReadyReferences**](JobPartReadyReferences.md)|  | 
 **replacement** | **optional.String**|  | [default to $]
 **prefix** | **optional.String**|  | 
 **suffix** | **optional.String**|  | 
 **length** | **optional.Float64**|  | [default to 1]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PseudoTranslateJobPart**
> PseudoTranslateJobPart(ctx, projectUid, jobUid, optional)
Pseudo-translates job



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiPseudoTranslateJobPartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiPseudoTranslateJobPartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PseudoTranslateActionDto**](PseudoTranslateActionDto.md)|  | 

### Return type

 (empty response body)

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
 **optional** | ***JobApiSearchByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchByJobOpts struct

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

# **SearchByJob2**
> SearchResponseListTmDtoV2 SearchByJob2(ctx, projectUid, jobUid, optional)
Search job's translation memories



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiSearchByJob2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchByJob2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTmByJobRequestDtoV2**](SearchTmByJobRequestDtoV2.md)|  | 

### Return type

[**SearchResponseListTmDtoV2**](SearchResponseListTmDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
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
 **optional** | ***JobApiSearchByJob3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchByJob3Opts struct

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

# **SearchPartsInProject**
> SearchJobsDto SearchPartsInProject(ctx, projectUid, optional)
Search jobs in project



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiSearchPartsInProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchPartsInProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SearchJobsRequestDto**](SearchJobsRequestDto.md)|  | 

### Return type

[**SearchJobsDto**](SearchJobsDto.md)

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
 **optional** | ***JobApiSearchSegmentByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchSegmentByJobOpts struct

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

# **SearchTermsByJob**
> SearchResponseListTbDto SearchTermsByJob(ctx, jobUid, projectUid, optional)
Search job's term bases

Search all read term bases assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***JobApiSearchTermsByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchTermsByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTbByJobRequestDto**](SearchTbByJobRequestDto.md)|  | 

### Return type

[**SearchResponseListTbDto**](SearchResponseListTbDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTermsByJob1**
> SearchTbResponseListDto SearchTermsByJob1(ctx, jobUid, projectUid, optional)
Search job's term bases

Search all read term bases assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***JobApiSearchTermsByJob1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchTermsByJob1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTbByJobRequestDto**](SearchTbByJobRequestDto.md)|  | 

### Return type

[**SearchTbResponseListDto**](SearchTbResponseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTermsInTextByJob**
> SearchInTextResponseListDto SearchTermsInTextByJob(ctx, jobUid, projectUid, optional)
Search terms in text

Search in text in all read term bases assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***JobApiSearchTermsInTextByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchTermsInTextByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchInTextByJobRequestDto**](SearchInTextByJobRequestDto.md)|  | 

### Return type

[**SearchInTextResponseListDto**](SearchInTextResponseListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTermsInTextByJobV2**
> SearchInTextResponseList2Dto SearchTermsInTextByJobV2(ctx, jobUid, projectUid, optional)
Search terms in text

Search in text in all read term bases assigned to the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **projectUid** | **string**|  | 
 **optional** | ***JobApiSearchTermsInTextByJobV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSearchTermsInTextByJobV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SearchTbInTextByJobRequestDto**](SearchTbInTextByJobRequestDto.md)|  | 

### Return type

[**SearchInTextResponseList2Dto**](SearchInTextResponseList2Dto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetStatus**
> SetStatus(ctx, projectUid, jobUid, optional)
Edit job status



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiSetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSetStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of JobStatusChangeActionDto**](JobStatusChangeActionDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Split**
> JobPartsDto Split(ctx, projectUid, jobUid, optional)
Split job

 Splits job by one of the following methods: * **After specific segments** - fill in `segmentOrdinals` * **Into X parts** - fill in `partCount`  * **Into parts with specific size** - fill in `partSize`. partSize represents segment count in each part. * **Into parts with specific word count** - fill in `wordCount`   * **By document parts** - fill in `byDocumentParts`, works only with **PowerPoint** files   Only one option at a time is allowed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 
 **optional** | ***JobApiSplitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiSplitOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SplitJobActionDto**](SplitJobActionDto.md)|  | 

### Return type

[**JobPartsDto**](JobPartsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatusChanges**
> JobPartStatusChangesDto StatusChanges(ctx, projectUid, jobUid)
Get status changes



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
  **jobUid** | **string**|  | 

### Return type

[**JobPartStatusChangesDto**](JobPartStatusChangesDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBilingualFile**
> JobPartsDto UploadBilingualFile(ctx, optional)
Upload bilingual file

Returns updated job parts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***JobApiUploadBilingualFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiUploadBilingualFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InputStream**](InputStream.md)|  | 
 **format** | **optional.String**|  | [default to MXLF]
 **saveToTransMemory** | **optional.String**|  | [default to Confirmed]
 **setCompleted** | **optional.Bool**|  | [default to false]

### Return type

[**JobPartsDto**](JobPartsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebEditorLink**
> WebEditorLinkDto WebEditorLink(ctx, projectUid, optional)
Get Web Editor URL

 Possible warning codes are:    - `NOT_ACCEPTED_BY_LINGUIST` - Job is not accepted by linguist    - `NOT_ASSIGNED_TO_LINGUIST` - Job is not assigned to linguist    - `PDF` - One of requested jobs is PDF   - `PREVIOUS_WORKFLOW_NOT_COMPLETED` - Previous workflow step is not completed    - `PREVIOUS_WORKFLOW_NOT_COMPLETED_STRICT` - Previous workflow step is not completed and project has strictWorkflowFinish set to true  Possible error codes are:    - `ASSIGNED_TO_OTHER_USER` - Job is accepted by other user    - `NOT_UNIQUE_TARGET_LANG` - Requested jobs contains different target locales    - `TOO_MANY_SEGMENTS` - Count of requested job's segments is higher than **40000**    Warning response example: ``` {     \"warnings\": [         {             \"message\": \"Not accepted by linguist\",             \"args\": {                 \"jobs\": [                     \"abcd1234\"                 ]             },             \"code\": \"NOT_ACCEPTED_BY_LINGUIST\"         },         {             \"message\": \"Previous workflow step not completed\",             \"args\": {                 \"jobs\": [                     \"abcd1234\"                 ]             },             \"code\": \"PREVIOUS_WORKFLOW_NOT_COMPLETED\"         }     ],     \"url\": \"/web/job/abcd1234-efgh5678/translate\" } ```  Error response example: Status: `400 Bad Request` ``` {     \"errorCode\": \"NOT_UNIQUE_TARGET_LANG\",     \"errorDescription\": \"Only files with identical target languages can be joined\",     \"errorDetails\": [         {             \"code\": \"NOT_UNIQUE_TARGET_LANG\",             \"args\": {                 \"targetLocales\": [                     \"de\",                     \"en\"                 ]             },             \"message\": \"Only files with identical target languages can be joined\"         },         {             \"code\": \"TOO_MANY_SEGMENTS\",             \"args\": {                 \"maxSegments\": 40000,                 \"segments\": 400009             },             \"message\": \"Up to 40000 segments can be opened in the Memsource Web Editor, job has 400009 segments\"         }     ] } ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiWebEditorLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiWebEditorLinkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateWebEditorLinkDto**](CreateWebEditorLinkDto.md)|  | 

### Return type

[**WebEditorLinkDto**](WebEditorLinkDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebEditorLinkV2**
> WebEditorLinkDtoV2 WebEditorLinkV2(ctx, projectUid, optional)
Get Web Editor URL

 Possible warning codes are:   - `NOT_ACCEPTED_BY_LINGUIST` - Job is not accepted by linguist   - `NOT_ASSIGNED_TO_LINGUIST` - Job is not assigned to linguist   - `PDF` - One of requested jobs is PDF   - `PREVIOUS_WORKFLOW_NOT_COMPLETED` - Previous workflow step is not completed   - `PREVIOUS_WORKFLOW_NOT_COMPLETED_STRICT` - Previous workflow step is not completed and project has strictWorkflowFinish set to true   - `IN_DELIVERED_STATE` - Jobs in DELIVERED state   - `IN_COMPLETED_STATE` - Jobs in COMPLETED state   - `IN_REJECTED_STATE` - Jobs in REJECTED state  Possible error codes are:   - `ASSIGNED_TO_OTHER_USER` - Job is accepted by other user   - `NOT_UNIQUE_TARGET_LANG` - Requested jobs contains different target locales   - `TOO_MANY_SEGMENTS` - Count of requested job's segments is higher than **40000**   - `COMPLETED_JOINED_WITH_OTHER` - Jobs in COMPLETED state cannot be joined with jobs in other states   - `DELIVERED_JOINED_WITH_OTHER` - Jobs in DELIVERED state cannot be joined with jobs in other states   - `REJECTED_JOINED_WITH_OTHER` - Jobs in REJECTED state cannot be joined with jobs in other states  Warning response example: ``` {     \"warnings\": [         {             \"message\": \"Not accepted by linguist\",             \"args\": {                 \"jobs\": [                     \"abcd1234\"                 ]             },             \"code\": \"NOT_ACCEPTED_BY_LINGUIST\"         },         {             \"message\": \"Previous workflow step not completed\",             \"args\": {                 \"jobs\": [                     \"abcd1234\"                 ]             },             \"code\": \"PREVIOUS_WORKFLOW_NOT_COMPLETED\"         }     ],     \"url\": \"/web/job/abcd1234-efgh5678/translate\" } ```  Error response example: Status: `400 Bad Request` ``` {     \"errorCode\": \"NOT_UNIQUE_TARGET_LANG\",     \"errorDescription\": \"Only files with identical target languages can be joined\",     \"errorDetails\": [         {             \"code\": \"NOT_UNIQUE_TARGET_LANG\",             \"args\": {                 \"targetLocales\": [                     \"de\",                     \"en\"                 ]             },             \"message\": \"Only files with identical target languages can be joined\"         },         {             \"code\": \"TOO_MANY_SEGMENTS\",             \"args\": {                 \"maxSegments\": 40000,                 \"segments\": 400009             },             \"message\": \"Up to 40000 segments can be opened in the Memsource Web Editor, job has 400009 segments\"         }     ] } ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectUid** | **string**|  | 
 **optional** | ***JobApiWebEditorLinkV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiWebEditorLinkV2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateWebEditorLinkDtoV2**](CreateWebEditorLinkDtoV2.md)|  | 

### Return type

[**WebEditorLinkDtoV2**](WebEditorLinkDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

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
 **optional** | ***JobApiWildCardSearchByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiWildCardSearchByJobOpts struct

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
 **optional** | ***JobApiWildCardSearchByJob2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiWildCardSearchByJob2Opts struct

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
 **optional** | ***JobApiWildCardSearchByJob3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JobApiWildCardSearchByJob3Opts struct

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

