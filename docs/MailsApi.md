# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MailsControllerFindAll**](MailsApi.md#MailsControllerFindAll) | **Get** /api/v1/mails | 送信済みメール一覧取得
[**MailsControllerFindOne**](MailsApi.md#MailsControllerFindOne) | **Get** /api/v1/mails/{id} | 送信済みメール詳細取得
[**MailsControllerSend**](MailsApi.md#MailsControllerSend) | **Post** /api/v1/mails | 新規単一メール送信
[**MailsControllerSendBatch**](MailsApi.md#MailsControllerSendBatch) | **Post** /api/v1/mails/batch | 新規複数メールバッチ送信

# **MailsControllerFindAll**
> MailsResponseDto MailsControllerFindAll(ctx, optional)
送信済みメール一覧取得

送信済みメールを一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MailsApiMailsControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailsApiMailsControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 

### Return type

[**MailsResponseDto**](MailsResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailsControllerFindOne**
> Mail MailsControllerFindOne(ctx, id)
送信済みメール詳細取得

リクエストパラメータのidで指定された単一の送信済みメールの詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Mail**](Mail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailsControllerSend**
> MailsControllerSend(ctx, body)
新規単一メール送信

新規単一メールを作成し、送信します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendMailDto**](SendMailDto.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailsControllerSendBatch**
> MailsControllerSendBatch(ctx, body)
新規複数メールバッチ送信

新規メールを複数作成し、送信します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendMailBatchDto**](SendMailBatchDto.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

