# \ImportSettingsApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImportSettings**](ImportSettingsApi.md#CreateImportSettings) | **Post** /api2/v1/importSettings | Create import settings
[**DeleteImportSettings**](ImportSettingsApi.md#DeleteImportSettings) | **Delete** /api2/v1/importSettings/{uid} | Delete import settings
[**GetImportSettings**](ImportSettingsApi.md#GetImportSettings) | **Get** /api2/v1/importSettings/{uid} | Get import settings
[**GetImportSettings1**](ImportSettingsApi.md#GetImportSettings1) | **Get** /api2/v1/importSettings/default | Get organization&#39;s default import settings
[**ListImportSettings**](ImportSettingsApi.md#ListImportSettings) | **Get** /api2/v1/importSettings | List import settings


# **CreateImportSettings**
> ImportSettingsDto CreateImportSettings(ctx, optional)
Create import settings

Pre-defined import settings is handy for [Create Job](#operation/createJob).                   See [supported file types](https://wiki.memsource.com/wiki/API_File_Type_List)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportSettingsApiCreateImportSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportSettingsApiCreateImportSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ImportSettingsCreateDto**](ImportSettingsCreateDto.md)|  | 

### Return type

[**ImportSettingsDto**](ImportSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImportSettings**
> DeleteImportSettings(ctx, uid)
Delete import settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportSettings**
> ImportSettingsDto GetImportSettings(ctx, uid)
Get import settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 

### Return type

[**ImportSettingsDto**](ImportSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportSettings1**
> ImportSettingsDto GetImportSettings1(ctx, )
Get organization's default import settings



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ImportSettingsDto**](ImportSettingsDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListImportSettings**
> PageDtoImportSettingsReference ListImportSettings(ctx, optional)
List import settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportSettingsApiListImportSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportSettingsApiListImportSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoImportSettingsReference**](PageDtoImportSettingsReference.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

