# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BadgesControllerAddLabel**](BadgesApi.md#BadgesControllerAddLabel) | **Put** /api/v1/badges/{id}/label | Badgeへのラベル追加
[**BadgesControllerCreate**](BadgesApi.md#BadgesControllerCreate) | **Post** /api/v1/badges | 新規バッジ発行
[**BadgesControllerDownload**](BadgesApi.md#BadgesControllerDownload) | **Get** /api/v1/badges/{id}/download | バッジダウンロード
[**BadgesControllerFileVerify**](BadgesApi.md#BadgesControllerFileVerify) | **Post** /api/v1/badges/file | バッジファイル検証
[**BadgesControllerFindAll**](BadgesApi.md#BadgesControllerFindAll) | **Get** /api/v1/badges | バッジ一覧取得
[**BadgesControllerFindOne**](BadgesApi.md#BadgesControllerFindOne) | **Get** /api/v1/badges/{id} | バッジ詳細取得
[**BadgesControllerVerify**](BadgesApi.md#BadgesControllerVerify) | **Put** /api/v1/badges/{id}/verify | バッジ検証

# **BadgesControllerAddLabel**
> BadgeResponseDto BadgesControllerAddLabel(ctx, body, id)
Badgeへのラベル追加

リクエストパラメータのidで指定された単一のBadgeに対して、任意の管理用ラベルを追加します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddLabelToBadgeDto**](AddLabelToBadgeDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**BadgeResponseDto**](BadgeResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BadgesControllerCreate**
> BadgeResponseDto BadgesControllerCreate(ctx, body)
新規バッジ発行

新規バッジを発行します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBadgeDto**](CreateBadgeDto.md)|  | 

### Return type

[**BadgeResponseDto**](BadgeResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BadgesControllerDownload**
> *os.File BadgesControllerDownload(ctx, id, vpInfoId)
バッジダウンロード

リクエストパラメータのidで指定された単一のバッジ画像のバイナリデータを返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **vpInfoId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BadgesControllerFileVerify**
> VerifiyBadgeFileReponseDto BadgesControllerFileVerify(ctx, file)
バッジファイル検証

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 

### Return type

[**VerifiyBadgeFileReponseDto**](VerifiyBadgeFileReponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BadgesControllerFindAll**
> BadgesResponseDto BadgesControllerFindAll(ctx, optional)
バッジ一覧取得

アプリケーションが管理するバッジ情報を一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BadgesApiBadgesControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BadgesApiBadgesControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **label** | **optional.String**| (任意) 例: \&quot;sample-badge-1\&quot; | 
 **filename** | **optional.String**| (任意) 例: \&quot;badge-12345-12345.png\&quot; | 
 **description** | **optional.String**| (任意)  | 
 **status** | **optional.Float64**| (任意) 例: 1 | 

### Return type

[**BadgesResponseDto**](BadgesResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BadgesControllerFindOne**
> BadgeResponseDto BadgesControllerFindOne(ctx, id)
バッジ詳細取得

リクエストパラメータのidで指定された単一のバッジ情報の詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**BadgeResponseDto**](BadgeResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BadgesControllerVerify**
> VerifiyBadgeReponseDto BadgesControllerVerify(ctx, body, id)
バッジ検証

バッジのVC/VPの検証を実行します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyBadgeDto**](VerifyBadgeDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**VerifiyBadgeReponseDto**](VerifiyBadgeReponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

