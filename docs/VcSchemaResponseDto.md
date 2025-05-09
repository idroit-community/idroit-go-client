# VcSchemaResponseDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VCスキーマの識別子 | [default to null]
**Title** | **string** | VCスキーマのタイトル | [default to null]
**Version** | **string** | VCスキーマのバージョン | [default to null]
**Description** | **string** | VCスキーマの説明文 | [default to null]
**IsBadgeSchema** | **string** | VCスキーマがバッジのスキーマか否か | [default to null]
**VcInfos** | [**[]VcInfo**](VcInfo.md) |  | [default to null]
**VcContexts** | [**[]VcContext**](VcContext.md) | VCのスキーマのJSONスキーマコンテキスト | [default to null]
**VcSchemaProperties** | [**[]VcSchemaProperty**](VcSchemaProperty.md) | VCのスキーマの各項目における項目名と項目型のオブジェクトの配列 | [default to null]
**Groups** | [**[]Group**](Group.md) |  | [default to null]
**Badges** | [**[]Badge**](Badge.md) | VCスキーマを紐付けたバッジ | [default to null]
**File** | [***AllOfVcSchemaResponseDtoFile**](AllOfVcSchemaResponseDtoFile.md) | VCスキーマに紐付けるバッジ用の画像。&#x60;isBadgeSchema&#x60;プロパティが&#x60;true&#x60;の場合必須。 | [default to null]
**CreatedBy** | [***AllOfVcSchemaResponseDtoCreatedBy**](AllOfVcSchemaResponseDtoCreatedBy.md) | VCスキーマを作成したユーザーアカウント | [default to null]
**CreatedAt** | **string** | VCスキーマの作成日時 | [default to null]
**UpdatedAt** | **string** | VCスキーマの最終更新日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

