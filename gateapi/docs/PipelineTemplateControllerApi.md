# \PipelineTemplateControllerApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublishPipelineTemplateUsingPOST**](PipelineTemplateControllerApi.md#PublishPipelineTemplateUsingPOST) | **Post** /pipelineTemplates/{id} | Publish a Pipeline Template


# **PublishPipelineTemplateUsingPOST**
> map[string]interface{} PublishPipelineTemplateUsingPOST(ctx, id, template)
Publish a Pipeline Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 
  **template** | [**interface{}**](interface{}.md)| Pipeline Template in json format | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

