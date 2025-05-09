# CreateVcSchemaDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | (必須) 新規生成するVCスキーマのタイトル | [default to null]
**Version** | **string** | (必須) VCスキーマのバージョン(デフォルト: 1.0.0) | [default to null]
**Contexts** | **[]string** | (任意) VCスキーマに設定する@context文字列の配列 | [optional] [default to null]
**Description** | **string** | (任意) VCスキーマの説明文 | [optional] [default to null]
**IsBadgeSchema** | **bool** | (任意) このスキーマがバッジのスキーマか否か(デフォルト: false) | [default to null]
**FileId** | **string** | (任意) バッジの画像のファイル識別子。&#x27;&#x60;isBadgeSchema&#x60;プロパティが&#x60;true&#x60;であり、バッジスキーマとして利用する場合は必須のパラメータ。 | [default to null]
**PropArray** | [**[]VcSchemaProperty**](VcSchemaProperty.md) | (必須) VCのスキーマの各項目における項目名と項目型のオブジェクトの配列 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

