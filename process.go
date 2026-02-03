package main

import (
	"os/exec"
	"strings"
)

type AppDef struct {
	AppName     string // macOSのアプリ名（frontmost検出用）
	DisplayName string
	Emoji       string
}

type AppStatus struct {
	DisplayName string
	Emoji       string
}

// 検出対象アプリ一覧
var appList = []AppDef{
	{AppName: "GoLand", DisplayName: "GoLand Play中", Emoji: "🐹"},
	{AppName: "CLion", DisplayName: "CLion Play中", Emoji: "🦁"},
	{AppName: "IntelliJ IDEA", DisplayName: "IntelliJ Play中", Emoji: "☕"},
	{AppName: "PyCharm", DisplayName: "PyCharm Play中", Emoji: "🐍"},
	{AppName: "WebStorm", DisplayName: "WebStorm Play中", Emoji: "🌀"},
	{AppName: "Code", DisplayName: "VSCode Play中", Emoji: "💻"},
	{AppName: "Ghostty", DisplayName: "Ghostty Play中", Emoji: "👻"},
}

// DetectActiveApp はアクティブウィンドウのアプリを検出する
func DetectActiveApp() *AppStatus {
	// AppleScriptで最前面のアプリ名を取得
	cmd := exec.Command("osascript", "-e",
		`tell application "System Events" to get name of first application process whose frontmost is true`)
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	activeApp := strings.ToLower(strings.TrimSpace(string(output)))

	// 対象アプリか確認（大文字小文字無視）
	for _, app := range appList {
		if strings.ToLower(app.AppName) == activeApp {
			return &AppStatus{
				DisplayName: app.DisplayName,
				Emoji:       app.Emoji,
			}
		}
	}

	return nil
}
