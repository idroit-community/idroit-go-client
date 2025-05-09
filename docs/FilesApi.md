# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesControllerDownload**](FilesApi.md#FilesControllerDownload) | **Get** /api/v1/files/{id}/download | ファイルダウンロード
[**FilesControllerExecuteCSV**](FilesApi.md#FilesControllerExecuteCSV) | **Post** /api/v1/files/{id} | CSVファイル実行
[**FilesControllerFindAll**](FilesApi.md#FilesControllerFindAll) | **Get** /api/v1/files | ファイル一覧取得
[**FilesControllerFindOne**](FilesApi.md#FilesControllerFindOne) | **Get** /api/v1/files/{id} | ファイル詳細取得
[**FilesControllerRemove**](FilesApi.md#FilesControllerRemove) | **Delete** /api/v1/files/{id} | ファイル削除
[**FilesControllerUploadFile**](FilesApi.md#FilesControllerUploadFile) | **Post** /api/v1/files | ファイルアップロード

# **FilesControllerDownload**
> *os.File FilesControllerDownload(ctx, id)
ファイルダウンロード

リクエストパラメータのidで指定された単一のファイルのバイナリデータを返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesControllerExecuteCSV**
> FilesControllerExecuteCSV(ctx, body, id)
CSVファイル実行

(非推奨) ユーザー、クライアント情報を記載したCSVファイルを実行し、新規アカウントを作成します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDidDto**](CreateDidDto.md)|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesControllerFindAll**
> FilesResponseDto FilesControllerFindAll(ctx, optional)
ファイル一覧取得

ファイルを一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesApiFilesControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiFilesControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **filename** | **optional.String**| (任意) 例: \&quot;file-1732019975229-394515535.png\&quot; | 
 **originalname** | **optional.String**| (任意) 例: \&quot;english-badge.png\&quot; | 
 **executed** | **optional.Bool**| (任意) 例: true | 
 **type_** | **optional.String**| (任意) 例: \&quot;image/png\&quot; | 
 **status** | **optional.Float64**| (任意) 例: 0 | 

### Return type

[**FilesResponseDto**](FilesResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesControllerFindOne**
> FileResponseDto FilesControllerFindOne(ctx, id)
ファイル詳細取得

リクエストパラメータのidで指定された単一のファイルの詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**FileResponseDto**](FileResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesControllerRemove**
> FilesControllerRemove(ctx, id)
ファイル削除

リクエストパラメータのidで指定された単一のファイルを削除します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesControllerUploadFile**
> FileResponseDto FilesControllerUploadFile(ctx, file)
ファイルアップロード

新規ファイルアップロードを作成します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 

### Return type

[**FileResponseDto**](FileResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

