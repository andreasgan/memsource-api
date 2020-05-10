# \WorkflowChangesApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadWorkflowChanges**](WorkflowChangesApi.md#DownloadWorkflowChanges) | **Post** /api2/v2/jobs/workflowChanges | Download workflow changes report


# **DownloadWorkflowChanges**
> Response DownloadWorkflowChanges(ctx, optional)
Download workflow changes report



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowChangesApiDownloadWorkflowChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowChangesApiDownloadWorkflowChangesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WorkflowChangesDto**](WorkflowChangesDto.md)|  | 

### Return type

[**Response**](Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

