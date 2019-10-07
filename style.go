package rlyeh

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Property - GUI property
type Property int32

// GUI properties enumeration
const (
	GlobalBaseColor Property = iota
	GlobalBorderColor
	GlobalTextColor
	GlobalTextFontsize
	GlobalBorderHeight
	GlobalBackgroundColor
	GlobalLinesColor
	GlobalPadding
	LabelBorderWidth
	LabelTextColor
	LabelTextPadding
	ButtonBorderWidth
	ButtonTextPadding
	ButtonDefaultBorderColor
	ButtonDefaultInsideColor
	ButtonDefaultTextColor
	ButtonHoverBorderColor
	ButtonHoverInsideColor
	ButtonHoverTextColor
	ButtonPressedBorderColor
	ButtonPressedInsideColor
	ButtonPressedTextColor
	ProgressbarBorderColor
	ProgressbarInsideColor
	ProgressbarProgressColor
	ProgressbarBorderWidth
	ComboboxPadding
	ComboboxWidth
	ComboboxHeight
	ComboboxBorderWidth
	ComboboxDefaultBorderColor
	ComboboxDefaultInsideColor
	ComboboxDefaultTextColor
	ComboboxDefaultListTextColor
	ComboboxHoverBorderColor
	ComboboxHoverInsideColor
	ComboboxHoverTextColor
	ComboboxHoverListTextColor
	ComboboxPressedBorderColor
	ComboboxPressedInsideColor
	ComboboxPressedTextColor
	ComboboxPressedListBorderColor
	ComboboxPressedListInsideColor
	ComboboxPressedListTextColor
	CheckboxDefaultBorderColor
	CheckboxDefaultInsideColor
	CheckboxHoverBorderColor
	CheckboxHoverInsideColor
	CheckboxClickBorderColor
	CheckboxClickInsideColor
	CheckboxDefaultActiveColor
	CheckboxInsideWidth
	TextboxBorderWidth
	TextboxBorderColor
	TextboxActiveBorderColor
	TextboxInsideColor
	TextboxTextColor
	TextboxLineColor
	TextboxTextFontsize
)

// Current GUI style (default light)
var style = []int64{
	0xf5f5f5ff, // GLOBAL_BASE_COLOR
	0x90abb5ff, // GLOBAL_BORDER_COLOR
	0xf5f5f5ff, // GLOBAL_TEXT_COLOR
	10,         // GLOBAL_TEXT_FONTSIZE
	5,          // GLOBAL_BORDER_HEIGHT
	0xf5f5f5ff, // GLOBAL_BACKGROUND_COLOR
	0x90abb5ff, // GLOBAL_LINES_COLOR
	2,          // GLOBAL_PADDING
	1,          // LABEL_BORDER_WIDTH
	0x4d4d4dff, // LABEL_TEXT_COLOR
	2,          // LABEL_TEXT_PADDING
	2,          // BUTTON_BORDER_WIDTH
	20,         // BUTTON_TEXT_PADDING
	0x828282ff, // BUTTON_DEFAULT_BORDER_COLOR
	0xc8c8c8ff, // BUTTON_DEFAULT_INSIDE_COLOR
	0x4d4d4dff, // BUTTON_DEFAULT_TEXT_COLOR
	0xc8c8c8ff, // BUTTON_HOVER_BORDER_COLOR
	0xffffffff, // BUTTON_HOVER_INSIDE_COLOR
	0x868686ff, // BUTTON_HOVER_TEXT_COLOR
	0x7bb0d6ff, // BUTTON_PRESSED_BORDER_COLOR
	0xbcecffff, // BUTTON_PRESSED_INSIDE_COLOR
	0x5f9aa7ff, // BUTTON_PRESSED_TEXT_COLOR
	0x828282ff, // PROGRESSBAR_BORDER_COLOR
	0xc8c8c8ff, // PROGRESSBAR_INSIDE_COLOR
	0xbcecffff, // PROGRESSBAR_PROGRESS_COLOR
	2,          // PROGRESSBAR_BORDER_WIDTH
	2,          // COMBOBOX_PADDING
	30,         // COMBOBOX_WIDTH
	20,         // COMBOBOX_HEIGHT
	1,          // COMBOBOX_BORDER_WIDTH
	0x828282ff, // COMBOBOX_DEFAULT_BORDER_COLOR
	0xc8c8c8ff, // COMBOBOX_DEFAULT_INSIDE_COLOR
	0x828282ff, // COMBOBOX_DEFAULT_TEXT_COLOR
	0x828282ff, // COMBOBOX_DEFAULT_LIST_TEXT_COLOR
	0xc8c8c8ff, // COMBOBOX_HOVER_BORDER_COLOR
	0xffffffff, // COMBOBOX_HOVER_INSIDE_COLOR
	0x828282ff, // COMBOBOX_HOVER_TEXT_COLOR
	0x828282ff, // COMBOBOX_HOVER_LIST_TEXT_COLOR
	0x7bb0d6ff, // COMBOBOX_PRESSED_BORDER_COLOR
	0xbcecffff, // COMBOBOX_PRESSED_INSIDE_COLOR
	0x5f9aa7ff, // COMBOBOX_PRESSED_TEXT_COLOR
	0x0078acff, // COMBOBOX_PRESSED_LIST_BORDER_COLOR
	0x66e7ffff, // COMBOBOX_PRESSED_LIST_INSIDE_COLOR
	0x0078acff, // COMBOBOX_PRESSED_LIST_TEXT_COLOR
	0x828282ff, // CHECKBOX_DEFAULT_BORDER_COLOR
	0xffffffff, // CHECKBOX_DEFAULT_INSIDE_COLOR
	0xc8c8c8ff, // CHECKBOX_HOVER_BORDER_COLOR
	0xffffffff, // CHECKBOX_HOVER_INSIDE_COLOR
	0x66e7ffff, // CHECKBOX_CLICK_BORDER_COLOR
	0xddf5ffff, // CHECKBOX_CLICK_INSIDE_COLOR
	0xbcecffff, // CHECKBOX_STATUS_ACTIVE_COLOR
	1,          // CHECKBOX_INSIDE_WIDTH
	1,          // TEXTBOX_BORDER_WIDTH
	0x828282ff, // TEXTBOX_BORDER_COLOR
	0x7bb0d6ff, // TEXTBOX_ACTIVE_BORDER_COLOR
	0xf5f5f5ff, // TEXTBOX_INSIDE_COLOR
	0x000000ff, // TEXTBOX_TEXT_COLOR
	0x000000ff, // TEXTBOX_LINE_COLOR
	10,         // TEXTBOX_TEXT_FONTSIZE
}

// GUI property names (to read/write style text files)
var propertyName = []string{
	"GLOBAL_BASE_COLOR",
	"GLOBAL_BORDER_COLOR",
	"GLOBAL_TEXT_COLOR",
	"GLOBAL_TEXT_FONTSIZE",
	"GLOBAL_BORDER_HEIGHT",
	"GLOBAL_BACKGROUND_COLOR",
	"GLOBAL_LINES_COLOR",
	"GLOBAL_PADDING",
	"LABEL_BORDER_WIDTH",
	"LABEL_TEXT_COLOR",
	"LABEL_TEXT_PADDING",
	"BUTTON_BORDER_WIDTH",
	"BUTTON_TEXT_PADDING",
	"BUTTON_DEFAULT_BORDER_COLOR",
	"BUTTON_DEFAULT_INSIDE_COLOR",
	"BUTTON_DEFAULT_TEXT_COLOR",
	"BUTTON_HOVER_BORDER_COLOR",
	"BUTTON_HOVER_INSIDE_COLOR",
	"BUTTON_HOVER_TEXT_COLOR",
	"BUTTON_PRESSED_BORDER_COLOR",
	"BUTTON_PRESSED_INSIDE_COLOR",
	"BUTTON_PRESSED_TEXT_COLOR",
	"PROGRESSBAR_BORDER_COLOR",
	"PROGRESSBAR_INSIDE_COLOR",
	"PROGRESSBAR_PROGRESS_COLOR",
	"PROGRESSBAR_BORDER_WIDTH",
	"COMBOBOX_PADDING",
	"COMBOBOX_WIDTH",
	"COMBOBOX_HEIGHT",
	"COMBOBOX_BORDER_WIDTH",
	"COMBOBOX_DEFAULT_BORDER_COLOR",
	"COMBOBOX_DEFAULT_INSIDE_COLOR",
	"COMBOBOX_DEFAULT_TEXT_COLOR",
	"COMBOBOX_DEFAULT_LIST_TEXT_COLOR",
	"COMBOBOX_HOVER_BORDER_COLOR",
	"COMBOBOX_HOVER_INSIDE_COLOR",
	"COMBOBOX_HOVER_TEXT_COLOR",
	"COMBOBOX_HOVER_LIST_TEXT_COLOR",
	"COMBOBOX_PRESSED_BORDER_COLOR",
	"COMBOBOX_PRESSED_INSIDE_COLOR",
	"COMBOBOX_PRESSED_TEXT_COLOR",
	"COMBOBOX_PRESSED_LIST_BORDER_COLOR",
	"COMBOBOX_PRESSED_LIST_INSIDE_COLOR",
	"COMBOBOX_PRESSED_LIST_TEXT_COLOR",
	"CHECKBOX_DEFAULT_BORDER_COLOR",
	"CHECKBOX_DEFAULT_INSIDE_COLOR",
	"CHECKBOX_HOVER_BORDER_COLOR",
	"CHECKBOX_HOVER_INSIDE_COLOR",
	"CHECKBOX_CLICK_BORDER_COLOR",
	"CHECKBOX_CLICK_INSIDE_COLOR",
	"CHECKBOX_STATUS_ACTIVE_COLOR",
	"CHECKBOX_INSIDE_WIDTH",
	"TEXTBOX_BORDER_WIDTH",
	"TEXTBOX_BORDER_COLOR",
	"TEXTBOX_ACTIVE_BORDER_COLOR",
	"TEXTBOX_INSIDE_COLOR",
	"TEXTBOX_TEXT_COLOR",
	"TEXTBOX_LINE_COLOR",
	"TEXTBOX_TEXT_FONTSIZE",
}

func GetColor(color Property) rl.Color {
	return rl.GetColor(int32(style[color]))
}

// BackgroundColor - Get background color
func BackgroundColor() rl.Color {
	return GetColor(GlobalBackgroundColor)
}

// LinesColor - Get lines color
func LinesColor() rl.Color {
	return GetColor(GlobalLinesColor)
}

// TextColor - Get text color for normal state
func TextColor() rl.Color {
	return GetColor(GlobalTextColor)
}

func SaveStyle(filename string) {
	var styleFile string
	for i := 0; i < len(propertyName); i++ {
		styleFile += fmt.Sprintf("%-40s0x%x\n", propertyName[i], GetStyleProperty(Property(i)))
	}

	ioutil.WriteFile(filename, []byte(styleFile), 0644)
}

func LoadStyle(filename string) {
	LoadStyleScaled(filename, 1)
}

func LoadStyleScaled(filename string, scale float32) {
	file, err := rl.OpenAsset(filename)
	if err != nil {
		rl.TraceLog(rl.LogWarning, "[%s] GUI style file could not be opened", filename)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}

		id := fields[0]
		value := fields[1]

		for i := 0; i < len(propertyName); i++ {
			if id == propertyName[i] {
				if strings.HasPrefix(value, "0x") {
					value = value[2:]
				}

				v, err := strconv.ParseInt(value, 16, 64)

				if strings.Contains(id, "WIDTH") || strings.Contains(id, "HEIGHT") ||
					strings.Contains(id, "PADDING") || strings.Contains(id, "SIZE") {
					v = int64(float32(v) * scale)
				}

				if err == nil {
					style[i] = v
				}
			}
		}
	}
}

// SetStyleProperty - Set one style property
func SetStyleProperty(guiProperty Property, value int64) {
	style[guiProperty] = value
}

// GetStyleProperty - Get one style property
func GetStyleProperty(guiProperty Property) int64 {
	return style[int(guiProperty)]
}
