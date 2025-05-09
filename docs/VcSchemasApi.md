# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VcSchemasControllerCreate**](VcSchemasApi.md#VcSchemasControllerCreate) | **Post** /api/v1/vc-schemas | 新規VCスキーマ作成
[**VcSchemasControllerFindAll**](VcSchemasApi.md#VcSchemasControllerFindAll) | **Get** /api/v1/vc-schemas | VCスキーマ一覧取得
[**VcSchemasControllerFindOne**](VcSchemasApi.md#VcSchemasControllerFindOne) | **Get** /api/v1/vc-schemas/{id} | VCスキーマ情報詳細取得
[**VcSchemasControllerUpdate**](VcSchemasApi.md#VcSchemasControllerUpdate) | **Put** /api/v1/vc-schemas/{id} | VCスキーマへのグループ紐付け

# **VcSchemasControllerCreate**
> VcSchemaResponseDto VcSchemasControllerCreate(ctx, body)
新規VCスキーマ作成

新規VCスキーマを作成します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVcSchemaDto**](CreateVcSchemaDto.md)|  | 

### Return type

[**VcSchemaResponseDto**](VcSchemaResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcSchemasControllerFindAll**
> VcSchemasResponseDto VcSchemasControllerFindAll(ctx, optional)
VCスキーマ一覧取得

VCスキーマを一覧として値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VcSchemasApiVcSchemasControllerFindAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VcSchemasApiVcSchemasControllerFindAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| (必須) ページネーションを表示する際のページ数に当たるクエリパラメーター。全件取得する際は1を指定。(デフォルト: 1) | 
 **limit** | **optional.Float64**| (必須) ページネーションを表示する際のページ数あたりに表示する件数を指定するクエリパラメーター。全件取得する際は0を指定。(デフォルト: 10) | 
 **title** | **optional.String**| (任意) 例: \&quot;Schema Sample\&quot; | 
 **version** | **optional.String**| (任意) 例: \&quot;1.0.0\&quot; | 
 **description** | **optional.String**| (任意) 例:  | 
 **isBadgeSchema** | **optional.Bool**| (任意) 例: true | 

### Return type

[**VcSchemasResponseDto**](VcSchemasResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcSchemasControllerFindOne**
> VcSchemaResponseDto VcSchemasControllerFindOne(ctx, id)
VCスキーマ情報詳細取得

リクエストパラメータのidで指定された単一のVCスキーマの詳細情報の値を返します。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**VcSchemaResponseDto**](VcSchemaResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcSchemasControllerUpdate**
> VcSchemaResponseDto VcSchemasControllerUpdate(ctx, body, id)
VCスキーマへのグループ紐付け

VCスキーマに関連するグループを紐付けます。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVcSchemaDto**](UpdateVcSchemaDto.md)|  | 
  **id** | **string**|  | 

### Return type

[**VcSchemaResponseDto**](VcSchemaResponseDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

