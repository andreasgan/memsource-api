# \EmailTemplateApi

All URIs are relative to *https://cloud.memsource.com/web*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgEmailTemplate**](EmailTemplateApi.md#GetOrgEmailTemplate) | **Get** /api2/v1/emailTemplates/{templateId} | Get email template
[**ListOrgEmailTemplates**](EmailTemplateApi.md#ListOrgEmailTemplates) | **Get** /api2/v1/emailTemplates | List email templates


# **GetOrgEmailTemplate**
> OrganizationEmailTemplateDto GetOrgEmailTemplate(ctx, templateId)
Get email template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**|  | 

### Return type

[**OrganizationEmailTemplateDto**](OrganizationEmailTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgEmailTemplates**
> PageDtoOrganizationEmailTemplateDto ListOrgEmailTemplates(ctx, optional)
List email templates



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailTemplateApiListOrgEmailTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailTemplateApiListOrgEmailTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | 
 **pageNumber** | **optional.Int32**| Page number, starting with 0, default 0 | [default to 0]
 **pageSize** | **optional.Int32**| Page size, accepts values between 1 and 50, default 50 | [default to 50]

### Return type

[**PageDtoOrganizationEmailTemplateDto**](PageDtoOrganizationEmailTemplateDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

