package steampunk

import "fmt"

// GaugeFrames returns animation frames for a pressure gauge.
// Each frame represents the needle at a different position (W, NW, E, S).
func GaugeFrames() []string {
	return []string{
		`  .---.  ` + "\n" +
			` /  |  \ ` + "\n" +
			`|  -+   |` + "\n" +
			` \     / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` / /   \ ` + "\n" +
			`|  +    |` + "\n" +
			` \     / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /     \ ` + "\n" +
			`|    +-  |` + "\n" +
			` \     / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /     \ ` + "\n" +
			`|   +   |` + "\n" +
			` \ |   / ` + "\n" +
			`  '---'  `,
	}
}

// GaugeScene returns a Scene animating a pressure gauge with a label.
// Default frame delay is 8 ticks (slightly slower than before for readability).
func GaugeScene(label string) *Scene {
	return GaugeSceneWithTheme(label, DefaultTheme())
}

// GaugeSceneWithTheme returns a gauge Scene using the provided ColorTheme.
func GaugeSceneWithTheme(label string, theme ColorTheme) *Scene {
	raw := GaugeFrames()
	framed := make([]string, len(raw))
	for i, f := range raw {
		framed[i] = Colorize(fmt.Sprintf("%s\n%s", f, label), theme)
	}
	return NewScene(framed, 8)
}
