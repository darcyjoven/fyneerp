package theme

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

//go:embed HarmonyOS_Sans_SC_Regular.ttf
var hmTTf []byte

var resourceHarmonyOSSansSCRegularTtf = &fyne.StaticResource{
	StaticName:    "HarmonyOS_Sans_SC_Regular.ttf",
	StaticContent: hmTTf,
}

// Font return bundled font resource
func (*MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceHarmonyOSSansSCRegularTtf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceHarmonyOSSansSCRegularTtf

}

func (*MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
