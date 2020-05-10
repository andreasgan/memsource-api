# \PriceListApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLanguagePair**](PriceListApi.md#CreateLanguagePair) | **Post** /api2/v1/priceLists/{priceListId}/priceSets | Add language pairs
[**CreatePriceList**](PriceListApi.md#CreatePriceList) | **Post** /api2/v1/priceLists | Create price list
[**DeleteLanguagePair**](PriceListApi.md#DeleteLanguagePair) | **Delete** /api2/v1/priceLists/{priceListId}/priceSets/{sourceLanguage}/{targetLanguage} | Remove language pair
[**DeleteLanguagePairs**](PriceListApi.md#DeleteLanguagePairs) | **Delete** /api2/v1/priceLists/{priceListId}/priceSets | Remove language pairs
[**DeletePriceList**](PriceListApi.md#DeletePriceList) | **Delete** /api2/v1/priceLists/{priceListId} | Delete price list
[**GetListOfPriceList**](PriceListApi.md#GetListOfPriceList) | **Get** /api2/v1/priceLists | List price lists
[**GetPriceList**](PriceListApi.md#GetPriceList) | **Get** /api2/v1/priceLists/{priceListId} | Get price list
[**GetPricesWithWorkflowSteps**](PriceListApi.md#GetPricesWithWorkflowSteps) | **Get** /api2/v1/priceLists/{priceListId}/priceSets | List price sets
[**SetMinimumPriceForSet**](PriceListApi.md#SetMinimumPriceForSet) | **Post** /api2/v1/priceLists/{priceListId}/priceSets/minimumPrices | Edit minimum prices
[**SetPrices**](PriceListApi.md#SetPrices) | **Post** /api2/v1/priceLists/{priceListId}/priceSets/prices | Edit prices
[**UpdatePriceList**](PriceListApi.md#UpdatePriceList) | **Put** /api2/v1/priceLists/{priceListId} | Update price list


# **CreateLanguagePair**
> TranslationPriceSetListDto CreateLanguagePair(ctx, priceListId, optional)
Add language pairs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 
 **optional** | ***PriceListApiCreateLanguagePairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PriceListApiCreateLanguagePairOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TranslationPriceSetCreateDto**](TranslationPriceSetCreateDto.md)|  | 

### Return type

[**TranslationPriceSetListDto**](TranslationPriceSetListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePriceList**
> TranslationPriceListDto CreatePriceList(ctx, optional)
Create price list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PriceListApiCreatePriceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PriceListApiCreatePriceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TranslationPriceListCreateDto**](TranslationPriceListCreateDto.md)|  | 

### Return type

[**TranslationPriceListDto**](TranslationPriceListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLanguagePair**
> DeleteLanguagePair(ctx, priceListId, sourceLanguage, targetLanguage)
Remove language pair



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 
  **sourceLanguage** | **string**|  | 
  **targetLanguage** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLanguagePairs**
> DeleteLanguagePairs(ctx, priceListId, optional)
Remove language pairs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 
 **optional** | ***PriceListApiDeleteLanguagePairsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PriceListApiDeleteLanguagePairsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TranslationPriceSetBulkDeleteDto**](TranslationPriceSetBulkDeleteDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePriceList**
> DeletePriceList(ctx, priceListId)
Delete price list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListOfPriceList**
> PageDtoTranslationPriceListDto GetListOfPriceList(ctx, optional)
List price lists



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PriceListApiGetListOfPriceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PriceListApiGetListOfPriceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoTranslationPriceListDto**](PageDtoTranslationPriceListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriceList**
> TranslationPriceListDto GetPriceList(ctx, priceListId)
Get price list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 

### Return type

[**TranslationPriceListDto**](TranslationPriceListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPricesWithWorkflowSteps**
> PageDtoTranslationPriceSetDto GetPricesWithWorkflowSteps(ctx, priceListId, optional)
List price sets



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 
 **optional** | ***PriceListApiGetPricesWithWorkflowStepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PriceListApiGetPricesWithWorkflowStepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]
 **sourceLanguages** | [**optional.Interface of []string**](string.md)|  | 
 **targetLanguages** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**PageDtoTranslationPriceSetDto**](PageDtoTranslationPriceSetDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMinimumPriceForSet**
> TranslationPriceListDto SetMinimumPriceForSet(ctx, priceListId, optional)
Edit minimum prices



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 
 **optional** | ***PriceListApiSetMinimumPriceForSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PriceListApiSetMinimumPriceForSetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TranslationPriceSetBulkMinimumPricesDto**](TranslationPriceSetBulkMinimumPricesDto.md)|  | 

### Return type

[**TranslationPriceListDto**](TranslationPriceListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPrices**
> TranslationPriceListDto SetPrices(ctx, priceListId, optional)
Edit prices

If object contains only price, all languages and workflow steps will be updated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 
 **optional** | ***PriceListApiSetPricesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PriceListApiSetPricesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TranslationPriceSetBulkPricesDto**](TranslationPriceSetBulkPricesDto.md)|  | 

### Return type

[**TranslationPriceListDto**](TranslationPriceListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePriceList**
> TranslationPriceListDto UpdatePriceList(ctx, priceListId, optional)
Update price list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **priceListId** | **int64**|  | 
 **optional** | ***PriceListApiUpdatePriceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PriceListApiUpdatePriceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TranslationPriceListCreateDto**](TranslationPriceListCreateDto.md)|  | 

### Return type

[**TranslationPriceListDto**](TranslationPriceListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

