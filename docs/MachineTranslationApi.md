# \MachineTranslationApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MachineTranslation**](MachineTranslationApi.md#MachineTranslation) | **Post** /api2/v1/machineTranslations/{mtSettingsId}/translate | Translate with MT


# **MachineTranslation**
> MachineTranslateResponse MachineTranslation(ctx, mtSettingsId, optional)
Translate with MT



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mtSettingsId** | **int64**|  | 
 **optional** | ***MachineTranslationApiMachineTranslationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MachineTranslationApiMachineTranslationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TranslationRequestExtendedDto**](TranslationRequestExtendedDto.md)|  | 

### Return type

[**MachineTranslateResponse**](MachineTranslateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

