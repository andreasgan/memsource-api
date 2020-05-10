# \ConversationsApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLQAComment**](ConversationsApi.md#AddLQAComment) | **Post** /api2/v1/jobs/{jobUid}/conversations/lqas/{conversationId}/comments | Add LQA comment
[**AddPlainComment**](ConversationsApi.md#AddPlainComment) | **Post** /api2/v1/jobs/{jobUid}/conversations/plains/{conversationId}/comments | Add plain comment
[**CreateLQAConversation**](ConversationsApi.md#CreateLQAConversation) | **Post** /api2/v1/jobs/{jobUid}/conversations/lqas | Create LQA conversation
[**CreateSegmentTargetConversation**](ConversationsApi.md#CreateSegmentTargetConversation) | **Post** /api2/v1/jobs/{jobUid}/conversations/plains | Create plain conversation
[**DeleteLQAComment**](ConversationsApi.md#DeleteLQAComment) | **Delete** /api2/v1/jobs/{jobUid}/conversations/lqas/{conversationId}/comments/{commentId} | Delete LQA comment
[**DeleteLQAConversation**](ConversationsApi.md#DeleteLQAConversation) | **Delete** /api2/v1/jobs/{jobUid}/conversations/lqas/{conversationId} | Delete conversation
[**DeletePlainComment**](ConversationsApi.md#DeletePlainComment) | **Delete** /api2/v1/jobs/{jobUid}/conversations/plains/{conversationId}/comments/{commentId} | Delete plain comment
[**DeletePlainConversation**](ConversationsApi.md#DeletePlainConversation) | **Delete** /api2/v1/jobs/{jobUid}/conversations/plains/{conversationId} | Delete plain conversation
[**FindConversations**](ConversationsApi.md#FindConversations) | **Post** /api2/v1/jobs/conversations/find | Find all conversation
[**GetLQAConversation**](ConversationsApi.md#GetLQAConversation) | **Get** /api2/v1/jobs/{jobUid}/conversations/lqas/{conversationId} | Get LQA conversation
[**GetPlainConversation**](ConversationsApi.md#GetPlainConversation) | **Get** /api2/v1/jobs/{jobUid}/conversations/plains/{conversationId} | Get plain conversation
[**ListAllConversations**](ConversationsApi.md#ListAllConversations) | **Get** /api2/v1/jobs/{jobUid}/conversations | List all conversations
[**ListLQAConversations**](ConversationsApi.md#ListLQAConversations) | **Get** /api2/v1/jobs/{jobUid}/conversations/lqas | List LQA conversations
[**ListPlainConversations**](ConversationsApi.md#ListPlainConversations) | **Get** /api2/v1/jobs/{jobUid}/conversations/plains | List plain conversations
[**UpdateLQAComment**](ConversationsApi.md#UpdateLQAComment) | **Put** /api2/v1/jobs/{jobUid}/conversations/lqas/{conversationId}/comments/{commentId} | Edit LQA comment
[**UpdateLQAConversation**](ConversationsApi.md#UpdateLQAConversation) | **Put** /api2/v1/jobs/{jobUid}/conversations/lqas/{conversationId} | Edit LQA conversation
[**UpdatePlainComment**](ConversationsApi.md#UpdatePlainComment) | **Put** /api2/v1/jobs/{jobUid}/conversations/plains/{conversationId}/comments/{commentId} | Edit plain comment
[**UpdatePlainConversation**](ConversationsApi.md#UpdatePlainConversation) | **Put** /api2/v1/jobs/{jobUid}/conversations/plains/{conversationId} | Edit plain conversation


# **AddLQAComment**
> LqaConversationDto AddLQAComment(ctx, jobUid, conversationId, optional)
Add LQA comment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 
 **optional** | ***ConversationsApiAddLQACommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiAddLQACommentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AddCommentDto**](AddCommentDto.md)|  | 

### Return type

[**LqaConversationDto**](LQAConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPlainComment**
> PlainConversationDto AddPlainComment(ctx, jobUid, conversationId, optional)
Add plain comment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 
 **optional** | ***ConversationsApiAddPlainCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiAddPlainCommentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AddCommentDto**](AddCommentDto.md)|  | 

### Return type

[**PlainConversationDto**](PlainConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLQAConversation**
> LqaConversationDto CreateLQAConversation(ctx, jobUid, optional)
Create LQA conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
 **optional** | ***ConversationsApiCreateLQAConversationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiCreateLQAConversationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateLqaConversationDto**](CreateLqaConversationDto.md)|  | 

### Return type

[**LqaConversationDto**](LQAConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSegmentTargetConversation**
> PlainConversationDto CreateSegmentTargetConversation(ctx, jobUid, optional)
Create plain conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
 **optional** | ***ConversationsApiCreateSegmentTargetConversationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiCreateSegmentTargetConversationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreatePlainConversationDto**](CreatePlainConversationDto.md)|  | 

### Return type

[**PlainConversationDto**](PlainConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLQAComment**
> DeleteLQAComment(ctx, jobUid, conversationId, commentId)
Delete LQA comment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 
  **commentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLQAConversation**
> DeleteLQAConversation(ctx, jobUid, conversationId)
Delete conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePlainComment**
> DeletePlainComment(ctx, jobUid, conversationId, commentId)
Delete plain comment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 
  **commentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePlainConversation**
> DeletePlainConversation(ctx, jobUid, conversationId)
Delete plain conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindConversations**
> ConversationListDto FindConversations(ctx, optional)
Find all conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsApiFindConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiFindConversationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of FindConversationsDto**](FindConversationsDto.md)|  | 

### Return type

[**ConversationListDto**](ConversationListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLQAConversation**
> LqaConversationDto GetLQAConversation(ctx, jobUid, conversationId)
Get LQA conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 

### Return type

[**LqaConversationDto**](LQAConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlainConversation**
> PlainConversationDto GetPlainConversation(ctx, jobUid, conversationId)
Get plain conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 

### Return type

[**PlainConversationDto**](PlainConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllConversations**
> ConversationListDto ListAllConversations(ctx, jobUid, optional)
List all conversations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
 **optional** | ***ConversationsApiListAllConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiListAllConversationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **optional.Bool**|  | [default to false]
 **since** | **optional.String**|  | 

### Return type

[**ConversationListDto**](ConversationListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLQAConversations**
> LqaConversationsListDto ListLQAConversations(ctx, jobUid, optional)
List LQA conversations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
 **optional** | ***ConversationsApiListLQAConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiListLQAConversationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **optional.Bool**|  | [default to false]
 **since** | **optional.String**|  | 

### Return type

[**LqaConversationsListDto**](LQAConversationsListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPlainConversations**
> PlainConversationsListDto ListPlainConversations(ctx, jobUid, optional)
List plain conversations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
 **optional** | ***ConversationsApiListPlainConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiListPlainConversationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **optional.Bool**|  | [default to false]
 **since** | **optional.String**|  | 

### Return type

[**PlainConversationsListDto**](PlainConversationsListDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLQAComment**
> LqaConversationDto UpdateLQAComment(ctx, jobUid, conversationId, commentId, optional)
Edit LQA comment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 
  **commentId** | **string**|  | 
 **optional** | ***ConversationsApiUpdateLQACommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiUpdateLQACommentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of AddCommentDto**](AddCommentDto.md)|  | 

### Return type

[**LqaConversationDto**](LQAConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLQAConversation**
> LqaConversationDto UpdateLQAConversation(ctx, jobUid, conversationId, optional)
Edit LQA conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 
 **optional** | ***ConversationsApiUpdateLQAConversationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiUpdateLQAConversationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EditLqaConversationDto**](EditLqaConversationDto.md)|  | 

### Return type

[**LqaConversationDto**](LQAConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePlainComment**
> PlainConversationDto UpdatePlainComment(ctx, jobUid, conversationId, commentId, optional)
Edit plain comment



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 
  **commentId** | **string**|  | 
 **optional** | ***ConversationsApiUpdatePlainCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiUpdatePlainCommentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of AddCommentDto**](AddCommentDto.md)|  | 

### Return type

[**PlainConversationDto**](PlainConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePlainConversation**
> PlainConversationDto UpdatePlainConversation(ctx, jobUid, conversationId, optional)
Edit plain conversation



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobUid** | **string**|  | 
  **conversationId** | **string**|  | 
 **optional** | ***ConversationsApiUpdatePlainConversationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsApiUpdatePlainConversationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EditPlainConversationDto**](EditPlainConversationDto.md)|  | 

### Return type

[**PlainConversationDto**](PlainConversationDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

