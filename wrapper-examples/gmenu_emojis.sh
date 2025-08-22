#!/bin/env bash

set -x

iDIR="$HOME/.config/swaync/images"
BIN="$HOME/.config/rofi/backend/bin/gmenu"

SCRIPT_DIR=$(dirname "$0")
source "$SCRIPT_DIR/notificationsHelpers.sh"

send_notification() {
	local message=$1
	local formatted_message
	formatted_message=$(format_message "$message")
	printf -v message_expanded "%b" "$formatted_message"

	notify-send -u normal -i "$iDIR/rofi-error.svg" -a "Rofi Menu" "$message_expanded"
}

ICON="${2:--1}"
PROMPT="${3:--1}"

declare -a ROFI_ARGS_FOR_ROFI=()

if [[ "$ICON" != "-1" ]]; then
	ROFI_ARGS_FOR_ROFI+=("-theme-str" "textbox-prompt-icon {str: \"$ICON\";}")
fi

if [[ "$PROMPT" != "-1" ]]; then
	ROFI_ARGS_FOR_ROFI+=("-theme-str" "entry { placeholder: \"$PROMPT\"; }")
fi

if [[ ! $(pidof rofi) ]]; then
	MENU_CONFIG_PATH="$1"
	if [[ -z "$MENU_CONFIG_PATH" ]]; then
		echo "Error: No menu config path provided."
		exit 1
	fi

	if [[ ! -f "$MENU_CONFIG_PATH" ]]; then
		echo "Error: Menu config file does not exist at $MENU_CONFIG_PATH."
		exit 1
	fi

	if [[ "$ICON" != "-1" && "$PROMPT" != "-1" ]]; then
		FULL_ROFI_ARGS=("${ROFI_ARGS_FOR_ROFI[@]}")
		if ! RESULT=$("$BIN" menu --menu-config "$MENU_CONFIG_PATH" --join | rofi -dmenu "${FULL_ROFI_ARGS[@]}"); then
			exit 1
		fi
		"$BIN" menu --menu-config "$MENU_CONFIG_PATH" --result "$RESULT" --default-exec "wl-copy"
	elif [[ "$ICON" != "-1" ]]; then
		ICON_ROFI_ARGS=("${ROFI_ARGS_FOR_ROFI[0]}")
		if ! RESULT=$("$BIN" menu --menu-config "$MENU_CONFIG_PATH" --join | rofi -dmenu "${ICON_ROFI_ARGS[@]}"); then
			exit 1
		fi
		"$BIN" menu --menu-config "$MENU_CONFIG_PATH" --result "$RESULT" --default-exec "wl-copy"
	elif [[ "$PROMPT" != "-1" ]]; then
		PROMPT_ROFI_ARGS=("${ROFI_ARGS_FOR_ROFI[1]}")
		if ! RESULT=$("$BIN" menu --menu-config "$MENU_CONFIG_PATH" --join | rofi -dmenu "${PROMPT_ROFI_ARGS[@]}"); then
			exit 1
		fi
		"$BIN" menu --menu-config "$MENU_CONFIG_PATH" --result "$RESULT" --default-exec "wl-copy"
	else
		if ! RESULT=$("$BIN" menu --menu-config "$MENU_CONFIG_PATH" --join | rofi -dmenu); then
			exit 1
		fi
		"$BIN" menu --menu-config "$MENU_CONFIG_PATH" --result "$RESULT" --default-exec "wl-copy"
	fi
	exit 0
else
	pkill rofi
	exit 0
fi
