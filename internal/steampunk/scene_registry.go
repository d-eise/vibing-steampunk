package steampunk

// SceneEntry describes a named scene with metadata.
type SceneEntry struct {
	Name        string
	Description string
	Scene       *Scene
}

// AllScenes returns a registry of all available named scenes using the default theme.
func AllScenes() []SceneEntry {
	theme := DefaultTheme()
	return []SceneEntry{
		{Name: "barometer", Description: "Barometer pressure gauge", Scene: BarometerSceneWithTheme(theme)},
		{Name: "boiler", Description: "Steam boiler", Scene: BoilerSceneWithTheme(theme)},
		{Name: "clock", Description: "Steampunk clock", Scene: ClockSceneWithTheme(theme)},
		{Name: "condenser", Description: "Steam condenser", Scene: CondenserSceneWithTheme(theme)},
		{Name: "flywheel", Description: "Spinning flywheel", Scene: FlywheelSceneWithTheme(theme)},
		{Name: "furnace", Description: "Coal furnace", Scene: FurnaceSceneWithTheme(theme)},
		{Name: "gauge", Description: "Pressure gauge", Scene: GaugeSceneWithTheme("PSI", theme)},
		{Name: "lever", Description: "Control lever", Scene: LeverSceneWithTheme(theme)},
		{Name: "pendulum", Description: "Pendulum clock", Scene: PendulumSceneWithTheme(theme)},
		{Name: "pipe", Description: "Steam pipe", Scene: PipeSceneWithTheme(theme)},
		{Name: "piston", Description: "Steam piston", Scene: PistonSceneWithTheme(theme)},
		{Name: "pressure", Description: "Pressure indicator", Scene: PressureSceneWithTheme(theme)},
		{Name: "smoke", Description: "Smoke plume", Scene: SmokeSceneWithTheme(theme)},
		{Name: "telegraph", Description: "Morse telegraph", Scene: TelegraphSceneWithTheme(theme)},
		{Name: "thermometer", Description: "Steam thermometer", Scene: ThermometerSceneWithTheme(theme)},
		{Name: "turbine", Description: "Steam turbine", Scene: TurbineSceneWithTheme(theme)},
		{Name: "valve", Description: "Steam valve", Scene: ValveSceneWithTheme(theme)},
		{Name: "anemometer", Description: "Wind speed meter", Scene: AnemometerSceneWithTheme(theme)},
		{Name: "bellows", Description: "Forge bellows", Scene: BellowsSceneWithTheme(theme)},
		{Name: "relief_valve", Description: "Pressure relief valve", Scene: ReliefValveSceneWithTheme(theme)},
	}
}

// FindScene looks up a scene by name from the registry.
// Returns the Scene and true if found, nil and false otherwise.
func FindScene(name string) (*Scene, bool) {
	for _, entry := range AllScenes() {
		if entry.Name == name {
			return entry.Scene, true
		}
	}
	return nil, false
}
