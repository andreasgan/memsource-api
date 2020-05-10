# \MachineTranslationSettingsApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetList**](MachineTranslationSettingsApi.md#GetList) | **Get** /api2/v1/machineTranslateSettings | List machine translate settings
[**GetMTSettings**](MachineTranslationSettingsApi.md#GetMTSettings) | **Get** /api2/v1/machineTranslateSettings/{mtsUid} | Get machine translate settings
[**GetMachineTranslateSettingsForProjectTemplate**](MachineTranslationSettingsApi.md#GetMachineTranslateSettingsForProjectTemplate) | **Get** /api2/v1/projectTemplates/{projectTemplateId}/mtSettings | Get machine translate settings
[**GetStatus**](MachineTranslationSettingsApi.md#GetStatus) | **Get** /api2/v1/machineTranslateSettings/{mtsUid}/status | Get status of machine translate engine
[**GetTranslationResources**](MachineTranslationSettingsApi.md#GetTranslationResources) | **Get** /api2/v1/projects/{projectUid}/jobs/{jobUid}/translationResources | Get translation resources


# **GetList**
> PageDtoMachineTranslateSettingsPbmDto GetList(ctx, optional)
List machine translate settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MachineTranslationSettingsApiGetListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MachineTranslationSettingsApiGetListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoMachineTranslateSettingsPbmDto**](PageDtoMachineTranslateSettingsPbmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMTSettings**
> MachineTranslateSettingsPbmDto GetMTSettings(ctx, mtsUid)
Get machine translate settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mtsUid** | **string**|  | 

### Return type

[**MachineTranslateSettingsPbmDto**](MachineTranslateSettingsPbmDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMachineTranslateSettingsForProjectTemplate**
> MtSettingsPerLanguageListDto GetMachineTranslateSettingsForProjectTemplate(ctx, projectTemplateId)
Get machine translate settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTemplateId** | **string**|  | 

### Return type

[**MtSettingsPerLanguageListDto**](MTSettingsPerLanguageListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatus**
> MachineTranslateStatusDto GetStatus(ctx, mtsUid)
Get status of machine translate engine



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mtsUid** | **string**|  | 

### Return type

[**MachineTranslateStatusDto**](MachineTranslateStatusDto.md)

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

