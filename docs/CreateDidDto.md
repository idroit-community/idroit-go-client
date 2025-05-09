# CreateDidDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DidMethod** | **string** | (必須) 新規生成するDIDのメソッド。(現在対応するDIDメソッド: \&quot;key\&quot;, \&quot;ether\&quot;, \&quot;ethr:sepolia\&quot;, \&quot;ethr:arbitrum:goerli\&quot;, \&quot;web\&quot;) | [default to null]
**Label** | **string** | (任意) DIDの識別や整理などの管理するにあったてメタデータとして任意で設定可能な単語のフレーズ。(例: did-for-project1) | [optional] [default to null]
**DomainName** | **string** | (任意) did:webメソッドでDIDを生成する際に必要なドメイン名。DIDによって指定されたドメインがドメインネームシステム(DNS)を通じて解決されるときのホスト名。did:webメソッド以外では必要のないパラメータなため任意の項目である。 | [optional] [default to null]
**Description** | **string** | (任意) DIDの生成目的や用途など任意で設定可能な説明文。 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

