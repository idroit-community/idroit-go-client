# AllOfBadgeVcSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | (必須) VCスキーマの識別子 | [default to null]
**Title** | **string** | (必須) VCスキーマのタイトル | [default to null]
**Version** | **string** | (必須) VCスキーマのバージョン | [default to null]
**Description** | **string** | (任意) VCスキーマの説明文 | [default to null]
**IsBadgeSchema** | **string** | (必須) VCスキーマがバッジのスキーマか否か | [optional] [default to false]
**VcInfos** | [**[]VcInfo**](VcInfo.md) | (任意) このVCスキーマを用いて生成したVC情報の配列。 | [optional] [default to null]
**VcContexts** | [**[]VcContext**](VcContext.md) | (任意) VCスキーマに設定したcontext項目の配列 | [optional] [default to null]
**VcSchemaProperties** | [**[]VcSchemaProperty**](VcSchemaProperty.md) | (必須) VCスキーマの項目情報の配列 | [optional] [default to null]
**Groups** | [**[]Group**](Group.md) | (任意) VCスキーマを紐付けたグループの配列 | [optional] [default to null]
**Badges** | [**[]Badge**](Badge.md) | (任意) VCスキーマを紐付けたバッジ | [optional] [default to null]
**File** | [***Object**](.md) | (任意) VCスキーマに紐付けるバッジ用の画像。&#x60;isBadgeSchema&#x60;プロパティが&#x60;true&#x60;の場合必須。 | [optional] [default to null]
**CreatedBy** | [***Object**](.md) | (必須) VCスキーマを作成したユーザーアカウント。 | [default to null]
**CreatedAt** | **string** | (必須) VCスキーマの作成日時 | [default to null]
**UpdatedAt** | **string** | (必須) VCスキーマの更新日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

