package rlyeh

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Style struct {
	GlobalBaseColor                 rl.Color
	GlobalBorderColor               rl.Color
	GlobalTextColor                 rl.Color
	GlobalTextFontsize              int
	GlobalBorderHeight              int
	GlobalBackgroundColor           rl.Color
	GlobalLinesColor                rl.Color
	GlobalPadding                   int
	DialogTitleFontsize             int
	DialogTitleBackgroundColor      rl.Color
	DialogTitleTextColor            rl.Color
	LabelBorderWidth                int
	LabelTextColor                  rl.Color
	LabelTextPadding                int
	ButtonBorderWidth               int
	ButtonTextPadding               int
	ButtonDefaultBorderColor        rl.Color
	ButtonDefaultInsideColor        rl.Color
	ButtonDefaultTextColor          rl.Color
	ButtonHoverBorderColor          rl.Color
	ButtonHoverInsideColor          rl.Color
	ButtonHoverTextColor            rl.Color
	ButtonPressedBorderColor        rl.Color
	ButtonPressedInsideColor        rl.Color
	ButtonPressedTextColor          rl.Color
	ComboboxPadding                 int
	ComboboxWidth                   int
	ComboboxHeight                  int
	ComboboxBorderWidth             int
	ComboboxDefaultBorderColor      rl.Color
	ComboboxDefaultInsideColor      rl.Color
	ComboboxDefaultTextColor        rl.Color
	ComboboxDefaultListTextColor    rl.Color
	ComboboxHoverBorderColor        rl.Color
	ComboboxHoverInsideColor        rl.Color
	ComboboxHoverTextColor          rl.Color
	ComboboxHoverListTextColor      rl.Color
	ComboboxPressedBorderColor      rl.Color
	ComboboxPressedInsideColor      rl.Color
	ComboboxPressedTextColor        rl.Color
	ComboboxPressedListBorderColor  rl.Color
	ComboboxPressedListInsideColor  rl.Color
	ComboboxPressedListTextColor    rl.Color
	CheckboxDefaultBorderColor      rl.Color
	CheckboxDefaultInsideColor      rl.Color
	CheckboxHoverBorderColor        rl.Color
	CheckboxHoverInsideColor        rl.Color
	CheckboxClickBorderColor        rl.Color
	CheckboxClickInsideColor        rl.Color
	CheckboxDefaultActiveColor      rl.Color
	CheckboxInsideWidth             int
	TextboxBorderWidth              int
	TextboxBorderColor              rl.Color
	TextboxActiveBorderColor        rl.Color
	TextboxInsideColor              rl.Color
	TextboxTextColor                rl.Color
	TextboxLineColor                rl.Color
	TextboxTextFontsize             int
	ListviewTextColor               rl.Color
	ListviewSelectedTextColor       rl.Color
	ListviewSelectedBackgroundColor rl.Color
}

func NewStyle() *Style {
	style := &Style{}

	style.GlobalBaseColor = rl.NewColor(0xf5, 0xf5, 0xf5, 0xff)
	style.GlobalBorderColor = rl.NewColor(0x90, 0xab, 0xb5, 0xff)
	style.GlobalTextColor = rl.NewColor(0xf5, 0xf5, 0xf5, 0xff)
	style.GlobalTextFontsize = 10
	style.GlobalBorderHeight = 5
	style.GlobalBackgroundColor = rl.NewColor(0xf5, 0xf5, 0xf5, 0xff)
	style.GlobalLinesColor = rl.NewColor(0x90, 0xab, 0xb5, 0xff)
	style.GlobalPadding = 2
	style.DialogTitleFontsize = 10
	style.DialogTitleBackgroundColor = rl.NewColor(0x90, 0xab, 0xb5, 0xff)
	style.DialogTitleTextColor = rl.White
	style.LabelBorderWidth = 1
	style.LabelTextColor = rl.NewColor(0x4d, 0x4d, 0x4d, 0xff)
	style.LabelTextPadding = 2
	style.ButtonBorderWidth = 2
	style.ButtonTextPadding = 20
	style.ButtonDefaultBorderColor = rl.NewColor(0x82, 0x82, 0x82, 0xff)
	style.ButtonDefaultInsideColor = rl.NewColor(0xc8, 0xc8, 0xc8, 0xff)
	style.ButtonDefaultTextColor = rl.NewColor(0x4d, 0x4d, 0x4d, 0xff)
	style.ButtonHoverBorderColor = rl.NewColor(0xc8, 0xc8, 0xc8, 0xff)
	style.ButtonHoverInsideColor = rl.White
	style.ButtonHoverTextColor = rl.NewColor(0x86, 0x86, 0x86, 0xff)
	style.ButtonPressedBorderColor = rl.NewColor(0x7b, 0xb0, 0xd6, 0xff)
	style.ButtonPressedInsideColor = rl.NewColor(0xbc, 0xec, 0xff, 0xff)
	style.ButtonPressedTextColor = rl.NewColor(0x5f, 0x9a, 0xa7, 0xff)
	style.ComboboxPadding = 2
	style.ComboboxWidth = 30
	style.ComboboxHeight = 20
	style.ComboboxBorderWidth = 1
	style.ComboboxDefaultBorderColor = rl.NewColor(0x82, 0x82, 0x82, 0xff)
	style.ComboboxDefaultInsideColor = rl.NewColor(0xc8, 0xc8, 0xc8, 0xff)
	style.ComboboxDefaultTextColor = rl.NewColor(0x82, 0x82, 0x82, 0xff)
	style.ComboboxDefaultListTextColor = rl.NewColor(0x82, 0x82, 0x82, 0xff)
	style.ComboboxHoverBorderColor = rl.NewColor(0xc8, 0xc8, 0xc8, 0xff)
	style.ComboboxHoverInsideColor = rl.White
	style.ComboboxHoverTextColor = rl.NewColor(0x82, 0x82, 0x82, 0xff)
	style.ComboboxHoverListTextColor = rl.NewColor(0x82, 0x82, 0x82, 0xff)
	style.ComboboxPressedBorderColor = rl.NewColor(0x7b, 0xb0, 0xd6, 0xff)
	style.ComboboxPressedInsideColor = rl.NewColor(0xbc, 0xec, 0xff, 0xff)
	style.ComboboxPressedTextColor = rl.NewColor(0x5f, 0x9a, 0xa7, 0xff)
	style.ComboboxPressedListBorderColor = rl.NewColor(0x00, 0x78, 0xac, 0xff)
	style.ComboboxPressedListInsideColor = rl.NewColor(0x66, 0xe7, 0xff, 0xff)
	style.ComboboxPressedListTextColor = rl.NewColor(0x00, 0x78, 0xac, 0xff)
	style.CheckboxDefaultBorderColor = rl.NewColor(0x82, 0x82, 0x82, 0xff)
	style.CheckboxDefaultInsideColor = rl.White
	style.CheckboxHoverBorderColor = rl.NewColor(0xc8, 0xc8, 0xc8, 0xff)
	style.CheckboxHoverInsideColor = rl.White
	style.CheckboxClickBorderColor = rl.NewColor(0x66, 0xe7, 0xff, 0xff)
	style.CheckboxClickInsideColor = rl.NewColor(0xdd, 0xf5, 0xff, 0xff)
	style.CheckboxDefaultActiveColor = rl.Black
	style.CheckboxInsideWidth = 1
	style.TextboxBorderWidth = 1
	style.TextboxBorderColor = rl.NewColor(0x82, 0x82, 0x82, 0xff)
	style.TextboxActiveBorderColor = rl.NewColor(0x7b, 0xb0, 0xd6, 0xff)
	style.TextboxInsideColor = rl.NewColor(0xf5, 0xf5, 0xf5, 0xff)
	style.TextboxTextColor = rl.Black
	style.TextboxLineColor = rl.Black
	style.TextboxTextFontsize = 10
	style.ListviewTextColor = rl.Black
	style.ListviewSelectedTextColor = rl.White
	style.ListviewSelectedBackgroundColor = rl.Blue
	return style
}

func (self *Style) Save(filename string) {
	bytes, _ := json.MarshalIndent(*self, "", " ")

	_ = ioutil.WriteFile(filename, bytes, 0644)
}

func (self *Style) Load(filename string) {
	jsonFile, err := os.Open(filename)
	if nil != err {
		return
	}

	bytes, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(bytes, self)
}

func (self *Style) Scale(scale float32) {
	value := reflect.ValueOf(self).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		x, ok := field.Interface().(int)
		x = int(float32(x) * scale)

		if ok && x > 0 {
			field.SetInt(int64(x))
		}

	}
}

func GetStyle() *Style {
	return style
}

var style = NewStyle()
