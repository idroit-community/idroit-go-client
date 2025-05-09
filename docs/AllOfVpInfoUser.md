# AllOfVpInfoUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | アカウントの識別子 | [default to null]
**Name** | **string** | アカウントの名前 | [default to null]
**Email** | **string** | アカウントのメールアドレス | [default to null]
**Password** | **string** | アカウントのパスワード(8文字以上20字未満) | [default to null]
**NeedActivateFlow** | **bool** | (必須: {default: false}) アカウントの有効化フローが必要か否か(true: 必要 /false: 不要) | [default to null]
**Status** | **string** | ユーザーのアカウントステータス | [default to null]
**Role** | **string** | アカウントのロール(admin/user/clinet) | [default to null]
**CreatedBy** | [***Object**](.md) | (任意) このアカウントを作成したユーザー | [optional] [default to null]
**DidInfos** | [**[]DidInfo**](DidInfo.md) | (任意) アカウントに紐付けされたDID情報の配列 | [optional] [default to null]
**VcInfos** | [**[]VcInfo**](VcInfo.md) | (任意) アカウントに紐付けされたVC情報の配列 | [optional] [default to null]
**VpInfos** | [**[]VpInfo**](VpInfo.md) | (任意) アカウントに紐付けされたVP情報の配列 | [optional] [default to null]
**Groups** | [**[]Group**](Group.md) | (任意) アカウントに紐付けされたグループの配列 | [optional] [default to null]
**UserTokens** | [**[]UserToken**](UserToken.md) | (任意) ユーザーアカウントがアップロードしたファイルの配列。 | [default to null]
**CreatedUsers** | [**[]User**](User.md) | (任意) このユーザーによって作成されたアカウントの配列 | [optional] [default to null]
**CreatedVcInfos** | [**[]VcInfo**](VcInfo.md) | (任意) アカウントに紐付けされたVC情報の配列 | [optional] [default to null]
**CreatedVpInfos** | [**[]VpInfo**](VpInfo.md) | (任意) アカウントに紐付けされたVP情報の配列 | [optional] [default to null]
**CreatedGroups** | [**[]Group**](Group.md) | (任意) アカウントに作成したグループの配列 | [optional] [default to null]
**CreatedDidInfos** | [**[]DidInfo**](DidInfo.md) | (任意) アカウントが作成したDID情報の配列 | [optional] [default to null]
**CreatedVcSchemas** | [**[]VcSchema**](VcSchema.md) | (任意) ユーザーアカウントが作成したVCスキーマの配列。 | [default to null]
**CreatedVerifications** | [**[]Verification**](Verification.md) | (任意) ユーザーアカウントが実行した検証結果の配列。 | [default to null]
**CreatedMails** | [**[]Mail**](Mail.md) | (任意) Admin権限アカウントが送信したメールの配列(Adminロール以外の場合、関係のないカラム) | [optional] [default to null]
**CreatedFiles** | [**[]*os.File**](*os.File.md) | (任意) ユーザーアカウントがアップロードしたファイルの配列。 | [default to null]
**CreatedAt** | **string** | ユーザーアカウントの作成日時 | [default to null]
**UpdatedAt** | **string** | ユーザーアカウントの更新日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

