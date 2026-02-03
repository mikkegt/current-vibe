package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	appID := os.Getenv("DISCORD_APP_ID")
	if appID == "" {
		log.Fatal("DISCORD_APP_ID environment variable is required")
	}

	fmt.Println("🎮 Discord Rich Presence を起動中...")

	if err := Connect(appID); err != nil {
		log.Fatalf("Discordへの接続に失敗しました: %v", err)
	}
	defer Disconnect()

	fmt.Println("✅ Discordに接続しました！")

	// シグナルハンドリング
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 初回更新
	if err := updatePresence(); err != nil {
		log.Fatalf("ステータス更新に失敗しました: %v", err)
	}

	// 5秒間隔でステータス更新
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := updatePresence(); err != nil {
				fmt.Println("\n⚠️  Discordとの接続が切れました。終了します...")
				return
			}
		case <-sigChan:
			fmt.Println("\n👋 終了します...")
			return
		}
	}
}

func updatePresence() error {
	app := DetectActiveApp()

	var state string

	if app != nil {
		state = fmt.Sprintf("%s %s", app.Emoji, app.DisplayName)
		fmt.Printf("🔄 %s\n", state)
	} else {
		state = GetRandomMessage()
		fmt.Printf("🔄 %s\n", state)
	}

	return UpdateStatus(state)
}
