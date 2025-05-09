# VerificationResponseDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 検証結果の識別子 | [default to null]
**Label** | **string** | 検証結果の識別や整理などの管理するにあったてメタデータとして任意で設定可能な単語のフレーズ。(例: verify-for-check) | [default to null]
**Result** | **bool** | (必須) VC/VPの検証結果(true: 検証に成功, false: 検証に失敗) | [default to null]
**CreatedBy** | [***AllOfVerificationResponseDtoCreatedBy**](AllOfVerificationResponseDtoCreatedBy.md) | グループを作成したユーザーアカウント | [default to null]
**VcInfo** | [***AllOfVerificationResponseDtoVcInfo**](AllOfVerificationResponseDtoVcInfo.md) | 検証を実行したVC情報の識別子 | [default to null]
**VpInfo** | [***AllOfVerificationResponseDtoVpInfo**](AllOfVerificationResponseDtoVpInfo.md) | 検証を実行したVP情報の識別子 | [default to null]
**CreatedAt** | **string** | 検証実行時の日時 | [default to null]
**UpdatedAt** | **string** | 検証結果の最終更新日 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

