# GMENU


---
### WIP

<p style="font-weight: bold; padding-left: 1rem; padding-right: 8rem; text-align:justify;">
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;I decided to rename what was just a color backend and fold it into one single binary comprehensive general purpose backend outputting data as lines for any picker capable of handling it (like dmenu, rofi, wofi, tofi, otherofi, walker, runner, crawler, swimmer and anything else under the sun that can fuzilly match lines).
    </p>
⚠️ 🚧 🔨 ⚠️ 🚧 🔨 ⚠️ 🚧 🔨 ⚠️ 🚧 🔨 ⚠️ 🚧 🔨 ⚠️ 🚧 🔨 ⚠️ 🚧 🔨 ⚠️ 🚧 🔨 ⚠️ 🚧 🔨 ⚠️ 🚧 🔨 
  


<small style="font-weight: normal; font-size: 0.8rem;">Previous: README below. New things to come...</small>
---
A command-line backend built in go specifically to pick colors using a color picker on linux (like default `hyprpicker`) and output them preformatted for easy use in CSS or other applications.  
Intended to feed dmenu | fzf and other fuzzy pickers like that.  

Formats:
- HEX
- RGB
- RGBA
- HSLA
- HSLA
- OKLAB
- OKLCH
- Closest named color

## Installation

1.  Ensure you have Go installed.
2.  Clone this repo `git clone https://github.com/tom-gora/rofiQuickColors.git && cd rofiQuickColors`
3.  Run `go build ./cmd/rofiQuickColors` to build the executable.
4.  Move the `rofiQuickColors` executable to a directory in your system's PATH, or to any other location you prefer, so that your menu application can call it.


## Usage

1. run rofiQuickColors
2. Provide optional arguments:

- [ -p | --picker-command ] STRING  
  Command to use to call color picker that must return hex color value to stdout.  
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

