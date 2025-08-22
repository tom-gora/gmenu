#!/bin/env bash

iDIR="$HOME/.config/swaync/images"
BIN="$HOME/.config/rofi/backend/bin/gmenu"

SCRIPT_DIR=$(dirname "$0")
source "$SCRIPT_DIR/notificationsHelpers.sh"

send_notification() {
	local message=$1
	local formatted_message
	formatted_message=$(format_message "$message")
	printf -v message_expanded "%b" "$formatted_message"

	notify-send -u normal -i "$iDIR/dropper.svg" -a "Color Picker" "$message_expanded"
}

main() {
	# dismiss prev notifications
	swaync-client --close-latest

	send_notification "Pick a color..." && sleep 0.5

	RESULT=$("$BIN" pick)
	# if result is empty or contains word "error" | "Error" then exit with 1
	if [[ "$RESULT" =~ [Cc]ancelled ]]; then
		swaync-client --close-latest
		send_notification "Picking cancelled."
		exit 130
	fi

	if [[ -z "$RESULT" || "$RESULT" =~ [Ee]rror ]]; then
		swaync-client --close-latest
		send_notification "Error:\n$RESULT"
		exit 1
	fi

	CHOICE=$(
		echo "$RESULT" | rofi -dmenu -sync \
			-theme-str 'inputbar { enabled: false; }' \
			-theme-str 'entry { enabled: false; }' \
			-theme-str 'imagebox { background-image: url("/tmp/color_thumb.png", height);}' \
			-theme-str 'listview { lines: 8;}'

	)
	if [[ -z "$CHOICE" ]]; then
		swaync-client --close-latest
		send_notification "Picking cancelled."
		exit 130
	fi

	swaync-client --close-latest
	send_notification "Copied color string:\n$CHOICE"
	echo "$CHOICE" | wl-copy
}

main
