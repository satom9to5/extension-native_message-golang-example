# 手順
## Firefox

### レジストリ追加
```
REG ADD "HKEY_CURRENT_USER\SOFTWARE\Mozilla\NativeMessagingHosts\example_golang" /ve /t REG_SZ /d "D:\vagrant_shared\firefox\native-node-example\app\example_golang.json" /f
```

## Chrome
一旦拡張機能のページで読み込む。
読み込むとIDが表示されるので、接続先アプリjsonに以下の形で追加

```
  "allowed_origins": [
    "chrome-extension://{ID}/"
  ]
```

### レジストリ追加
```
REG ADD "HKEY_CURRENT_USER\SOFTWARE\Google\Chrome\NativeMessagingHosts\example_golang" /ve /t REG_SZ /d "D:\vagrant_shared\firefox\native-node-example\app\example_golang.json" /f 
```

# 特徴
## connection_based
1回アクセスするとすぐにonDisconnectが発火して切断される

## connectionless
複数回アクセス可能

# 注意点
- メッセージ送信ごとに外部アプリが起動する
- メッセージが返却されると外部アプリはすぐ終了する

# 開発時の注意点等
## Chrome
### リロード
拡張機能ページ内で実行

### デバッグ
拡張機能の「バックグラウンドページ」をクリック

### エラー
#### 外部アプリへの接続が出来なかった場合のエラーメッセージ
```
Error: Attempting to use a disconnected port object
```
