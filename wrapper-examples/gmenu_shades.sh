#!/bin/env bash

# set -x

iDIR="$HOME/.config/swaync/images"
BIN="$HOME/.config/rofi/backend/bin/gmenu"
FIRST_STEP_MENU="$HOME/.config/rofi/backend/menus/pick_or_clip.json"
ROFI_ICON_MENU_CONFIG="$HOME/.config/rofi/config-icon-menu.rasi"
ROFI_SHADES_CONFIG="$HOME/.config/rofi/config-shades.rasi"

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
	ACTION=$(
		"$BIN" menu -m "$FIRST_STEP_MENU" | rofi -dmenu \
			-markup-rows \
			-config "$ROFI_ICON_MENU_CONFIG"
	)

	if [[ -z "$ACTION" ]]; then
		pkill rofi
		swaync-client --close-latest
		send_notification "Picking cancelled."
		exit 130
	fi

	if [[ "$ACTION" == *"use_picker"* ]]; then
		pkill rofi
		FULL_GRADIENT_LIST=$("$BIN" shades)
		if [[ -z "$FULL_GRADIENT_LIST" || "$FULL_GRADIENT_LIST" == "Cancelled by user" ]]; then
			swaync-client --close-latest
			send_notification "Picking cancelled."
			exit 130
		fi
	else
		FULL_GRADIENT_LIST=$("$BIN" shades -C)
	fi

	declare -a TEMPLATE

	i=0
	while IFS= read -r line; do
		# interpolate color hex code into predefined lines of pango markup
		if ((i == 10)); then
			#specifically mark the orginal input color in between 10 up and 10 down shades and tints
			TEMPLATED_LINE="<span size='16pt' fgcolor='#000' bgcolor='$line'></span><span rise='2pt' fgcolor='$line'>  $line</span>"
		else
			TEMPLATED_LINE="<span size='20pt' fgcolor='$line' bgcolor='$line'>O</span><span rise='4pt' fgcolor='$line'>  $line</span>"
		fi
		TEMPLATE+=("$TEMPLATED_LINE")
		((i++))
	done <<<"$FULL_GRADIENT_LIST"

	# prep line separated output for rofi to consume
	MARKUP=$(printf "%s\n" "${TEMPLATE[@]}")
	# get selection from rofi picker
	CHOICE=$(
		echo "$MARKUP" | rofi -dmenu \
			-markup-rows \
			-config "$ROFI_SHADES_CONFIG" |
			grep -oE "#[0-9a-fA-F]{6}" | head -n 1
	)

	# if selection made push it to clipboard
	if [[ -n "$CHOICE" ]]; then
		echo "$CHOICE" | tr -d '\n' | wl-copy
	fi

}
main
