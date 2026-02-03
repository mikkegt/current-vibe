# Discord Rich Presence Go

macOS用のDiscord Rich Presence CLIツール。アクティブなアプリに応じてDiscordのステータスを自動更新します。

## 対応アプリ

| アプリ | 表示 |
|--------|------|
| GoLand | 🐹 GoLand Play中 |
| CLion | 🦁 CLion Play中 |
| IntelliJ IDEA | ☕ IntelliJ Play中 |
| PyCharm | 🐍 PyCharm Play中 |
| WebStorm | 🌀 WebStorm Play中 |
| VSCode | 💻 VSCode Play中 |
| Ghostty | 👻 Ghostty Play中 |

対象外のアプリがアクティブな場合は、ランダムなメッセージが表示されます。

## セットアップ

### 1. Discord Developer Portalでアプリを作成

1. https://discord.com/developers/applications にアクセス
2. 「New Application」をクリック
3. アプリ名を入力（Discordのステータスに表示される名前）
4. Application IDをコピー

### 2. ビルド

```bash
go build -o discord-presence .
```

### 3. 実行

```bash
export DISCORD_APP_ID=あなたのアプリID
./discord-presence
```

または1行で：

```bash
DISCORD_APP_ID=あなたのアプリID ./discord-presence
```

## 終了

`Ctrl+C` で終了します。
