# UserResponseDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ユーザーアカウントの識別子 | [default to null]
**Name** | **string** | ユーザーアカウントの名前 | [default to null]
**Email** | **string** | ユーザーアカウントのメールアドレス | [default to null]
**NeedActivateFlow** | **bool** | ユーザーアクティベーション(true: アクティブ /false: 停止中) | [default to null]
**Status** | **string** | ユーザーのアカウントステータス(\&quot;inactive\&quot;: アクティブ未完了, \&quot;acrivating\&quot;: アクティブ作業途中, \&quot;active\&quot;: アクティブ中, \&quot;deactive\&quot;: 論理削除済) | [default to null]
**Role** | **string** | ユーザーのアカウントロール(\&quot;admin\&quot;, \&quot;user\&quot;, \&quot;client\&quot;) | [default to null]
**UserTokens** | [**[]UserToken**](UserToken.md) | (任意) ユーザーアカウントがアップロードしたファイルの配列。 | [default to null]
**CreatedBy** | [***AllOfUserResponseDtoCreatedBy**](AllOfUserResponseDtoCreatedBy.md) | ユーザーアカウントを作成したユーザーアカウント | [default to null]
**DidInfos** | [**[]DidInfo**](DidInfo.md) | ユーザーアカウントに紐付けられたDID情報の配列 | [optional] [default to null]
**VcInfos** | [**[]VcInfo**](VcInfo.md) | ユーザーアカウントに紐付けられたVC情報の配列 | [optional] [default to null]
**VpInfos** | [**[]VpInfo**](VpInfo.md) | ユーザーアカウントに紐付けられたVP情報の配列 | [optional] [default to null]
**Groups** | [**[]Group**](Group.md) | ユーザーアカウントが作成したグループの配列 | [optional] [default to null]
**CreatedUsers** | [**[]User**](User.md) | ユーザーアカウントに紐付けされたVP情報の配列 | [optional] [default to null]
**CreatedVcInfos** | [**[]VcInfo**](VcInfo.md) | ユーザーが作成したVC情報の配列 | [optional] [default to null]
**CreatedVpInfos** | [**[]VpInfo**](VpInfo.md) | ユーザーが作成したVP情報の配列 | [optional] [default to null]
**CreatedGroups** | [**[]Group**](Group.md) | ユーザーアカウントが作成したグループの配列 | [optional] [default to null]
**CreatedDidInfos** | [**[]DidInfo**](DidInfo.md) | ユーザーが作成したDID情報の配列 | [optional] [default to null]
**CreatedVcSchemas** | [**[]VcSchema**](VcSchema.md) | ユーザーが作成したVCスキーマの配列 | [optional] [default to null]
**CreatedVerifications** | [**[]Verification**](Verification.md) | ユーザーアカウントが実行した検証結果の配列 | [optional] [default to null]
**CreatedMails** | [**[]Mail**](Mail.md) | Admin権限アカウントが送信したメールの配列 | [optional] [default to null]
**CreatedFiles** | [**[]*os.File**](*os.File.md) | ユーザーアカウントに紐付けされたファイルの配列 | [optional] [default to null]
**CreatedAt** | **string** | ユーザーアカウントの作成日時 | [default to null]
**UpdatedAt** | **string** | ユーザーアカウントの更新日時 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

