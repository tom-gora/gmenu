# rofi-quick-colors

A command-line backend built in go specifically to pick colors using a color picker on linux (like default `hyprpicker`) and output them preformatted for easy use in CSS or other applications.  
Intended to feed dmenu | fzf and other pickers like that.
Formats:

- HEX
- RGB
- RGBA
- HSLA
- HSLA
- OKLAB
- OKLCH
- Closest named color

## Usage

1. run rofiQuickColors
2. Provide optional arguments:

- [ -p | --picker-command ] STRING  
  Command to use to call picker that must return hex color value to stdout.  
  DEFAULT: "hyprpicker -n -f hex"

- [ -d | --draw-thumbnail] BOOLEAN  
  Should a temporarty PNG thumbnail filled with picked color be created in /tmp/color_thumb.png
  Could be used as color preview in gui pickers.

  DEFAULT: true

3. It will output the color in different formats to stdout.

## TODOs

- [x] Output as "\n" separated lines
- [ ] Output as JSON
- [ ] Output color shades and tints

## License

MIT

