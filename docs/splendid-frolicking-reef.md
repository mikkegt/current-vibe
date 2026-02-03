# Discord Rich Presence CLI Tool - 実装計画

## 概要
GoでDiscord Rich Presenceを使ってIDEのステータスを表示するCLIツール（macOS専用）

**Application ID:** 環境変数 `DISCORD_APP_ID` から読み込み

---

## 使用ライブラリ

| ライブラリ | 用途 |
|-----------|------|
| [hugolgst/rich-go](https://github.com/hugolgst/rich-go) | Discord Rich Presence |
| [shirou/gopsutil/v4](https://github.com/shirou/gopsutil) | プロセス監視 |

---

## macOS プロセス検出（調査済み）

| アプリ | 検出方法 | 絵文字 | 優先度 |
|--------|----------|--------|--------|
| GoLand | プロセス名 `goland` | 🐹 | 1（最高） |
| CLion | プロセス名 `clion` | 🦁 | 2 |
| IntelliJ IDEA | プロセス名 `idea` | ☕ | 3 |
| PyCharm | プロセス名 `pycharm` | 🐍 | 4 |
| WebStorm | プロセス名 `webstorm` | 🌀 | 5 |
| VSCode | パスに `Visual Studio Code.app` 含む | 💻 | 6 |
| Ghostty | プロセス名 `ghostty` | 👻 | 7 |

※ JetBrains系は小文字のプロセス名で検出
※ VSCodeはElectronアプリのため、実行パスで判定

---

## ファイル構成

```
discord-rich-presence-go/
├── go.mod
├── go.sum
├── main.go           # エントリーポイント、メインループ
├── presence.go       # Discord Rich Presence操作
├── process.go        # プロセス監視ロジック
└── messages.go       # ランダムメッセージ定義
```

---

## 実装詳細

### 1. main.go
```go
func main() {
    appID := os.Getenv("DISCORD_APP_ID")
    if appID == "" {
        log.Fatal("DISCORD_APP_ID environment variable is required")
    }

    // Discord接続
    // メインループ（5秒間隔）
    // Ctrl+C でgraceful shutdown
}
```

### 2. process.go
```go
type AppDef struct {
    ProcessName string // プロセス名で検出（空ならパスで検出）
    PathContains string // 実行パスに含まれる文字列で検出
    DisplayName string
    Emoji       string
}

var appList = []AppDef{
    {ProcessName: "goland", DisplayName: "GoLand", Emoji: "🐹"},
    {ProcessName: "clion", DisplayName: "CLion", Emoji: "🦁"},
    {ProcessName: "idea", DisplayName: "IntelliJ IDEA", Emoji: "☕"},
    {ProcessName: "pycharm", DisplayName: "PyCharm", Emoji: "🐍"},
    {ProcessName: "webstorm", DisplayName: "WebStorm", Emoji: "🌀"},
    {PathContains: "Visual Studio Code.app", DisplayName: "VSCode", Emoji: "💻"},
    {ProcessName: "ghostty", DisplayName: "Ghostty", Emoji: "👻"},
}

func DetectApp() *AppStatus  // 優先度順にチェック
```

### 3. presence.go
```go
func Connect(appID string) error
func UpdateStatus(state, details string) error
func Disconnect()
```

### 4. messages.go
```go
var funnyMessages = []string{
    // 動物系
    "にゃんこと会議中🐱",
    "ゴムのアヒルと相談中🦆",
    "しばいぬとお散歩中🐕",
    "ペンギンとペアプロ中🐧",
    "たこと踊り中🐙",
    "うさぎとデバッグ中🐰",

    // 食べ物系
    "コーヒー充填中☕",
    "ラーメン啜り中🍜",
    "おにぎり補給中🍙",
    "たいやき食べ中🐟",

    // 活動系
    "思考中🧐",
    "昼寝モード💤",
    "瞑想中🧘",
    "妄想中💭",

    // テック系
    "バグと格闘中🐛",
    "無限ループから脱出中🔄",
    "Stack Overflow巡回中📚",
    "キーボード叩き中⌨️",
    "AIに質問中🤖",

    // ファンタジー系
    "宇宙と交信中🛸",
    "異世界転生準備中✨",
    "魔法詠唱中🪄",
    "ユニコーンと冒険中🦄",
}

func GetRandomMessage() string
```

---

## 表示例

| 状態 | State | Details |
|------|-------|---------|
| GoLand起動中 | 🐹 GoLand Play中 | 🐹 コーディング中 |
| CLion起動中 | 🦁 CLion Play中 | 🦁 コーディング中 |
| IntelliJ起動中 | ☕ IntelliJ IDEA Play中 | ☕ コーディング中 |
| PyCharm起動中 | 🐍 PyCharm Play中 | 🐍 コーディング中 |
| WebStorm起動中 | 🌀 WebStorm Play中 | 🌀 コーディング中 |
| VSCode起動中 | 💻 VSCode Play中 | 💻 コーディング中 |
| Ghostty起動中 | 👻 Ghostty Play中 | 👻 ターミナル中 |
| アプリ未検出 | (ランダムメッセージ) | 💭 休憩中 |

---

## 使用方法

```bash
# ビルド
go build -o discord-presence .

# 実行
export DISCORD_APP_ID=1468233551391494293
./discord-presence
```

---

## 検証方法

1. `go build` でビルド
2. Discordデスクトップアプリを起動
3. 環境変数を設定して実行
4. Discordプロフィールでステータス確認
5. GoLand/PyCharm/WebStormを起動・終了してステータス変化を確認
