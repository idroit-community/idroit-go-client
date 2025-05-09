# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerificationsControllerCreate**](VerificationsApi.md#VerificationsControllerCreate) | **Post** /api/v1/verifications | VC/VP検証
[**VerificationsControllerFindAll**](VerificationsApi.md#VerificationsControllerFindAll) | **Get** /api/v1/verifications | VC/VP検証結果一覧取得
[**VerificationsControllerFindOne**](VerificationsApi.md#VerificationsControllerFindOne) | **Get** /api/v1/verifications/{id} | VC/VP検証結果詳細取得

# **VerificationsControllerCreate**
> VerificationResponseDto VerificationsControllerCreate(ctx, body)
VC/VP検証

VC/VPの検証を実行します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerificationDto**](VerificationDto.md)|  | 

### Return type

[**VerificationResponseDto**](VerificationResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerificationsControllerFindAll**
> VerificationsResponseDto VerificationsControllerFindAll(ctx, optional)
VC/VP検証結果一覧取得

アプリケーションが管理するVC/VP検証結果を一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerificationsApiVerificationsControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerificationsApiVerificationsControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **label** | **optional.String**| (任意) 例: \&quot;verification-for-project1\&quot; | 
 **result** | **optional.Bool**| (任意) 例: true | 

### Return type

[**VerificationsResponseDto**](VerificationsResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerificationsControllerFindOne**
> VerificationResponseDto VerificationsControllerFindOne(ctx, id)
VC/VP検証結果詳細取得

リクエストパラメータのidで指定された単一のVC/VP検証結果の詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**VerificationResponseDto**](VerificationResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

