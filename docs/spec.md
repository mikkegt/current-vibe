# Current Vibe - 仕様書

## 概要

macOS用Discord Rich Presence CLIツール。アクティブウィンドウに応じてDiscordステータスを自動更新する。

## 動作環境

- macOS専用（AppleScriptでアクティブウィンドウを検出）
- Discordデスクトップアプリが起動している必要あり

## 起動・終了条件

| 条件 | 動作 |
|------|------|
| `DISCORD_APP_ID` 未設定 | エラー終了 |
| Discord未起動 | 接続失敗でエラー終了 |
| 実行中にDiscord終了 | 「接続が切れました」と表示して終了 |
| `Ctrl+C` | 正常終了 |

## アクティブウィンドウ検出

AppleScriptを使用してフォアグラウンドアプリを検出。

### 対応アプリ

| アプリ | 検出方法 | 絵文字 | 表示 |
|--------|----------|--------|------|
| GoLand | バンドルID `com.jetbrains.goland` | 🐹 | GoLand Play中 |
| CLion | バンドルID `com.jetbrains.CLion` | 🦁 | CLion Play中 |
| IntelliJ IDEA | バンドルID `com.jetbrains.intellij` | ☕ | IntelliJ Play中 |
| PyCharm | バンドルID `com.jetbrains.pycharm` | 🐍 | PyCharm Play中 |
| WebStorm | バンドルID `com.jetbrains.WebStorm` | 🌀 | WebStorm Play中 |
| VSCode | バンドルID `com.microsoft.VSCode` | 💻 | VSCode Play中 |
| Ghostty | バンドルID `com.mitchellh.ghostty` | 👻 | Ghostty Play中 |

### 対象外アプリ

対象外のアプリがアクティブな場合、`messages.go`で定義されたメッセージからランダムに表示。

カテゴリ: 動物系、食べ物系、気持ち系、テック系、ファンタジー系

## 更新間隔

5秒ごとにステータスを更新。

## ファイル構成

```
current-vibe/
├── main.go        # エントリーポイント、メインループ
├── presence.go    # Discord Rich Presence操作
├── process.go     # アクティブウィンドウ検出（AppleScript）
├── messages.go    # ランダムメッセージ定義
├── go.mod
├── go.sum
├── README.md      # ユーザー向けドキュメント
└── docs/
    └── spec.md    # 本仕様書
```

## 使用ライブラリ

| ライブラリ | 用途 |
|-----------|------|
| [hugolgst/rich-go](https://github.com/hugolgst/rich-go) | Discord Rich Presence |
