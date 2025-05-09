# CreateVcInfoDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** | (必須) VCの発行者の識別子として用いる文字列の値。現在はDIDのみがサポートされていますが、今後のアップデートでDID以外の文字列をサポートする予定です。 | [default to null]
**VcSchemaId** | **string** | (必須) 新規発行するVCの元となるVCスキーマの識別子 | [default to null]
**CredentialSubject** | [***interface{}**](interface{}.md) | (必須) VCの主張内容(クレーム)となる値のオブジェクト型の値 | [default to null]
**Label** | **string** | (任意) VCの識別や整理などの管理するにあったてメタデータとして任意で設定可能な単語のフレーズ。(例: vc-for-project1) | [optional] [default to null]
**Description** | **string** | (任意) VCの発行目的や用途など任意で設定可能な説明文 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

