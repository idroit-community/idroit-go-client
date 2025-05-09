# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VcInfosControllerAddLabel**](VcInfosApi.md#VcInfosControllerAddLabel) | **Put** /api/v1/vc-infos/{id}/label | VC情報へのラベル追加
[**VcInfosControllerConnectUser**](VcInfosApi.md#VcInfosControllerConnectUser) | **Put** /api/v1/vc-infos/{id} | VC情報へのクライアントアカウント紐付け
[**VcInfosControllerCreate**](VcInfosApi.md#VcInfosControllerCreate) | **Post** /api/v1/vc-infos | 新規VC発行
[**VcInfosControllerDisconnectUser**](VcInfosApi.md#VcInfosControllerDisconnectUser) | **Delete** /api/v1/vc-infos/{id}/user | VC情報のクライアントアカウント紐付け解除
[**VcInfosControllerFindAll**](VcInfosApi.md#VcInfosControllerFindAll) | **Get** /api/v1/vc-infos | VC情報一覧取得
[**VcInfosControllerFindOne**](VcInfosApi.md#VcInfosControllerFindOne) | **Get** /api/v1/vc-infos/{id} | VC情報詳細取得
[**VcInfosControllerGenerateVp**](VcInfosApi.md#VcInfosControllerGenerateVp) | **Post** /api/v1/vc-infos/{id} | 新規VP情報生成
[**VcInfosControllerIssue**](VcInfosApi.md#VcInfosControllerIssue) | **Post** /api/v1/vc-infos/issue | 新規VC発行(スキーマなし)
[**VcInfosControllerUpload**](VcInfosApi.md#VcInfosControllerUpload) | **Post** /api/v1/vc-infos/upload | 新規VCアップロード

# **VcInfosControllerAddLabel**
> VcInfoResponseDto VcInfosControllerAddLabel(ctx, body, id)
VC情報へのラベル追加

リクエストパラメータのidで指定された単一のVC情報に対して、任意の管理用ラベルを追加します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddLabelToVcInfoDto**](AddLabelToVcInfoDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**VcInfoResponseDto**](VcInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcInfosControllerConnectUser**
> VcInfosControllerConnectUser(ctx, body, id)
VC情報へのクライアントアカウント紐付け

VC情報の所有者、関係者などの管理者としてクライアントアカウントを紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVcInfoDto**](UpdateVcInfoDto.md)|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcInfosControllerCreate**
> VcInfoResponseDto VcInfosControllerCreate(ctx, body)
新規VC発行

新規VCを発行します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVcInfoDto**](CreateVcInfoDto.md)|  | 

### Return type

[**VcInfoResponseDto**](VcInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcInfosControllerDisconnectUser**
> VcInfosControllerDisconnectUser(ctx, id)
VC情報のクライアントアカウント紐付け解除

VC情報に紐付いたクライアントアカウントの紐付けを解除します。

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

# **VcInfosControllerFindAll**
> VcInfosResponseDto VcInfosControllerFindAll(ctx, optional)
VC情報一覧取得

アプリケーションが管理するVC情報を一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VcInfosApiVcInfosControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VcInfosApiVcInfosControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **label** | **optional.String**| (任意) 例: \&quot;vc-for-project1\&quot; | 
 **description** | **optional.String**| (任意)  | 

### Return type

[**VcInfosResponseDto**](VcInfosResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcInfosControllerFindOne**
> VcInfoResponseDto VcInfosControllerFindOne(ctx, id)
VC情報詳細取得

リクエストパラメータのidで指定された単一のVC情報の詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**VcInfoResponseDto**](VcInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcInfosControllerGenerateVp**
> VcInfoResponseDto VcInfosControllerGenerateVp(ctx, body, id)
新規VP情報生成

リクエストパラメータのidで指定されたVC情報から新規VPを発行します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GenerateVpDto**](GenerateVpDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**VcInfoResponseDto**](VcInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcInfosControllerIssue**
> VcInfoResponseDto VcInfosControllerIssue(ctx, body)
新規VC発行(スキーマなし)

VCスキーマを指定せず直接新規VCを発行します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IssueVcInfoDto**](IssueVcInfoDto.md)|  | 

### Return type

[**VcInfoResponseDto**](VcInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcInfosControllerUpload**
> VcInfoResponseDto VcInfosControllerUpload(ctx, credentialObject, label, description)
新規VCアップロード

外部で発行された既存VCをアップロードし、本アプリケーションに保存します。。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialObject** | [**interface{}**](.md)|  | 
  **label** | **string**|  | 
  **description** | **string**|  | 

### Return type

[**VcInfoResponseDto**](VcInfoResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

