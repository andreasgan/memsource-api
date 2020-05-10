# \SpellCheckApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWord**](SpellCheckApi.md#AddWord) | **Post** /api2/v1/spellCheck/words | Add word to dictionary
[**Check**](SpellCheckApi.md#Check) | **Post** /api2/v1/spellCheck/check | Spell check
[**CheckByJob**](SpellCheckApi.md#CheckByJob) | **Post** /api2/v1/spellCheck/check/{jobUid} | Spell check for job
[**Suggest**](SpellCheckApi.md#Suggest) | **Post** /api2/v1/spellCheck/suggest | Suggest a word


# **AddWord**
> AddWord(ctx, optional)
Add word to dictionary



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpellCheckApiAddWordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpellCheckApiAddWordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DictionaryItemDto**](DictionaryItemDto.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Check**
> SpellCheckResponseDto Check(ctx, optional)
Spell check

Spell check using the settings of the user's organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpellCheckApiCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpellCheckApiCheckOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SpellCheckRequestDto**](SpellCheckRequestDto.md)|  | 

### Return type

[**SpellCheckResponseDto**](SpellCheckResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckByJob**
> SpellCheckResponseDto CheckByJob(ctx, jobUid, optional)
Spell check for job

Spell check using the settings from the project of the job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
 **optional** | ***SpellCheckApiCheckByJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpellCheckApiCheckByJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SpellCheckRequestDto**](SpellCheckRequestDto.md)|  | 

### Return type

[**SpellCheckResponseDto**](SpellCheckResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Suggest**
> SuggestResponseDto Suggest(ctx, optional)
Suggest a word

Spell check suggest using the users's spell check dictionary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpellCheckApiSuggestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpellCheckApiSuggestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SuggestRequestDto**](SuggestRequestDto.md)|  | 

### Return type

[**SuggestResponseDto**](SuggestResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

