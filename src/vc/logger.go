/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 Ashok Shau
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/AshokShau/TgMusicBot
 */

package vc

import (
        "ashokshau/tgmusic/config"
        "ashokshau/tgmusic/src/utils"
        "fmt"

        td "github.com/AshokShau/gotdbot"
)

// sendLogger sends a formatted log message to the designated logger chat.
func sendLogger(client *td.Client, chatID int64, song *utils.CachedTrack) {
        if chatID == 0 || song == nil || chatID == config.LoggerId {
                return
        }

        // Fetching Chat Details for logs (Removed Username to fix compilation error)
        chatTitle := "Unknown"
        if chat, err := client.GetChat(chatID); err == nil && chat != nil {
                if chat.Title != "" {
                        chatTitle = chat.Title
                }
        }

        // Formatting Stream Type
        streamType := song.Platform
        if song.IsVideo {
                streamType += " (Video)"
        } else {
                streamType += " (Audio)"
        }

        // ShiviMusic Style Format
        text := fmt.Sprintf(
                "<b>❖ ᴘʟᴀʏ ʟᴏɢ</b>\n\n"+
                "<b>● ᴄʜᴀᴛ ɪᴅ ➠</b> <code>%d</code>\n"+
                "<b>● ᴄʜᴀᴛ ɴᴀᴍᴇ ➠</b> %s\n\n"+
                "<b>● ʀᴇǫᴜᴇsᴛᴇᴅ ʙʏ ➠</b> %s\n"+
                "<b>● ǫᴜᴇʀʏ ➠</b> <a href='%s'>%s</a>\n"+
                "<b>● ᴅᴜʀᴀᴛɪᴏɴ ➠</b> %s\n"+
                "<b>● sᴛʀᴇᴀᴍᴛʏᴘᴇ ➠</b> %s",
                chatID,
                chatTitle,
                song.User,
                song.URL,
                song.Name,
                utils.SecToMin(song.Duration),
                streamType,
        )

        _, err := client.SendTextMessage(config.LoggerId, text, &td.SendTextMessageOpts{
                DisableWebPagePreview: true,
                ParseMode:             "HTML",
        })
        
        if err != nil {
                logger.Warn("Failed to send the message", "error", err)
        }
}
