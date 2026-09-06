/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 Ashok Shau
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/AshokShau/TgMusicBot
 */

package core

import (
	"ashokshau/tgmusic/config"
	"ashokshau/tgmusic/src/utils"
	"fmt"

	"github.com/AshokShau/gotdbot"
)

func cb(text, data string, style gotdbot.ButtonStyle) gotdbot.InlineKeyboardButton {
	return gotdbot.InlineKeyboardButton{
		Text: text,
		Type: &gotdbot.InlineKeyboardButtonTypeCallback{
			Data: []byte(data),
		},
		Style: style,
	}
}

func userId(text string, userId int64, style gotdbot.ButtonStyle) gotdbot.InlineKeyboardButton {
	return gotdbot.InlineKeyboardButton{
		Text:  text,
		Type:  &gotdbot.InlineKeyboardButtonTypeUser{UserId: userId},
		Style: style,
	}
}

func url(text, link string, style gotdbot.ButtonStyle) gotdbot.InlineKeyboardButton {
	return gotdbot.InlineKeyboardButton{
		Text: text,
		Type: &gotdbot.InlineKeyboardButtonTypeUrl{
			Url: link,
		},
		Style: style,
	}
}

// ---- ALL BUTTONS WITH CUSTOM FONT & EMOJIS ----
var CloseBtn = cb("𝐂ℓσѕє", "vcplay_close", gotdbot.ButtonStyleDanger{})
var HomeBtn = cb("𝐇σмє", "help_back", gotdbot.ButtonStylePrimary{})
var HelpBtn = cb("📥 𝐇єℓρ 𝐀η∂ 𝐂σммαη∂ѕ", "help_all", gotdbot.ButtonStylePrimary{})
var UserBtn = cb("𝐔ѕєяѕ", "help_user", gotdbot.ButtonStyleDefault{})
var AdminBtn = cb("𝐀∂мιηѕ", "help_admin", gotdbot.ButtonStyleDefault{})
var OwnerBtn = cb("🚫 𝐎ωηєя", "help_owner", gotdbot.ButtonStyleDanger{})
var DevsBtn = cb("✨ 𝐂ℓσηє", "help_devs", gotdbot.ButtonStyleDanger{}) // Kept function same, changed text to Clone
var PlaylistBtn = cb("𝐏ℓαуℓιѕт", "help_playlist", gotdbot.ButtonStyleDefault{})
var AutoplayBtn = cb("𝐀υтσρℓαу", "help_autoplay", gotdbot.ButtonStyleDefault{})

var SourceCodeBtn = url("📦 𝐒συя¢є", "https://t.me/SukkuBeatzBot", gotdbot.ButtonStyleDanger{})
var channelBtn = url("📢 𝐔ρ∂αтєѕ", config.SupportChannel, gotdbot.ButtonStyleDanger{})
var groupBtn = url("👥 𝐆яσυρ", config.SupportGroup, gotdbot.ButtonStyleDanger{})

func SupportKeyboard() *gotdbot.ReplyMarkupInlineKeyboard {
	return &gotdbot.ReplyMarkupInlineKeyboard{
		Rows: [][]gotdbot.InlineKeyboardButton{
			{channelBtn, groupBtn},
			{CloseBtn},
		},
	}
}

func SupportBtn() *gotdbot.ReplyMarkupInlineKeyboard {
	return &gotdbot.ReplyMarkupInlineKeyboard{
		Rows: [][]gotdbot.InlineKeyboardButton{
			{channelBtn, groupBtn},
		},
	}
}

func SettingsKeyboard(playMode, adminMode string, cmdDelete bool, language string) *gotdbot.ReplyMarkupInlineKeyboard {
	playText := "𝐄νєяуσηє"
	if playMode == utils.Admins {
		playText = "𝐀∂мιηѕ"
	}

	deleteText := "𝐅αℓѕє"
	if cmdDelete {
		deleteText = "𝐓яυє"
	}

	adminText := "𝐄νєяуσηє"
	if adminMode == utils.Admins {
		adminText = "𝐀∂мιηѕ"
	}

	langText := "𝐄ηgℓιѕн"
	if language != "en" && language != "" {
		langText = language
	}

	return &gotdbot.ReplyMarkupInlineKeyboard{
		Rows: [][]gotdbot.InlineKeyboardButton{
			{
				cb("𝐏ℓαу 𝐌σ∂є ➜", "settings_main", gotdbot.ButtonStyleDefault{}),
				cb(playText, "settings_play", gotdbot.ButtonStyleDefault{}),
			},
			{
				cb("𝐂σммαη∂ 𝐃єℓєтє ➜", "settings_main", gotdbot.ButtonStyleDefault{}),
				cb(deleteText, "settings_delete", gotdbot.ButtonStyleDefault{}),
			},
			{
				cb("𝐀∂мιη 𝐌σ∂є ➜", "settings_main", gotdbot.ButtonStyleDefault{}),
				cb(adminText, "settings_admin", gotdbot.ButtonStyleDefault{}),
			},
			{
				cb("𝐋αηgυαgє ➜", "settings_main", gotdbot.ButtonStyleDefault{}),
				cb(langText, "settings_lang", gotdbot.ButtonStyleDefault{}),
			},
			{CloseBtn},
		},
	}
}

func HelpMenuKeyboard() *gotdbot.ReplyMarkupInlineKeyboard {
	return &gotdbot.ReplyMarkupInlineKeyboard{
		Rows: [][]gotdbot.InlineKeyboardButton{
			{UserBtn, AdminBtn, OwnerBtn},
			{PlaylistBtn, DevsBtn, AutoplayBtn},
			{HomeBtn, CloseBtn},
		},
	}
}

func BackHelpMenuKeyboard() *gotdbot.ReplyMarkupInlineKeyboard {
	return &gotdbot.ReplyMarkupInlineKeyboard{
		Rows: [][]gotdbot.InlineKeyboardButton{
			{HelpBtn, HomeBtn},
			{CloseBtn, SourceCodeBtn},
		},
	}
}

func ControlButtons(mode string) *gotdbot.ReplyMarkupInlineKeyboard {
	skipBtn := cb("‣‣I", "play_skip", gotdbot.ButtonStyleDefault{})
	stopBtn := cb("▢", "play_stop", gotdbot.ButtonStyleDefault{})
	pauseBtn := cb("II", "play_pause", gotdbot.ButtonStyleDefault{})
	resumeBtn := cb("▷", "play_resume", gotdbot.ButtonStyleDefault{})
	muteBtn := cb("🔇", "play_mute", gotdbot.ButtonStyleDefault{})
	unmuteBtn := cb("🔊", "play_unmute", gotdbot.ButtonStyleDefault{})
	addToPlaylistBtn := cb("➕", "play_add_to_list", gotdbot.ButtonStylePrimary{})

	switch mode {

	case "play":
		return &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{skipBtn, stopBtn, pauseBtn},
				{addToPlaylistBtn, CloseBtn},
			},
		}

	case "pause":
		return &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{skipBtn, stopBtn, resumeBtn},
				{CloseBtn},
			},
		}

	case "resume":
		return &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{skipBtn, stopBtn, pauseBtn},
				{CloseBtn},
			},
		}

	case "mute":
		return &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{skipBtn, stopBtn, unmuteBtn},
				{CloseBtn},
			},
		}

	case "unmute":
		return &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{skipBtn, stopBtn, muteBtn},
				{CloseBtn},
			},
		}

	default:
		return &gotdbot.ReplyMarkupInlineKeyboard{
			Rows: [][]gotdbot.InlineKeyboardButton{
				{CloseBtn},
			},
		}
	}
}

// Sizzu Style Layout 
func AddMeMarkup(username string) *gotdbot.ReplyMarkupInlineKeyboard {

	addMeBtn := url(
		"✨ 𝐀∂∂ 𝐌є 𝐓σ 𝐘συя 𝐆яσυρ",
		fmt.Sprintf("https://t.me/%s?startgroup=true", username),
		gotdbot.ButtonStylePrimary{},
	)

	return &gotdbot.ReplyMarkupInlineKeyboard{
		Rows: [][]gotdbot.InlineKeyboardButton{
			{addMeBtn},                // Full-width (1st row)
			{OwnerBtn, DevsBtn},       // Half-width (2nd row)
			{channelBtn, SourceCodeBtn}, // Half-width (3rd row)
			{HelpBtn},                 // Full-width (4th row)
		},
	}
}

func PlayNowButton(trackID string) gotdbot.InlineKeyboardButton {
	return cb("▶️ 𝐏ℓαу 𝐍σω", fmt.Sprintf("play_now_%s", trackID), gotdbot.ButtonStyleDanger{})
}

func QueueMarkup(trackID string) *gotdbot.ReplyMarkupInlineKeyboard {
	return &gotdbot.ReplyMarkupInlineKeyboard{
		Rows: [][]gotdbot.InlineKeyboardButton{
			{PlayNowButton(trackID), CloseBtn},
		},
	}
}
