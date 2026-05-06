package steampunk

// AllScenes returns a map of scene name to Scene factory functions
// using the default theme.
func AllScenes() map[string]*Scene {
	theme := DefaultTheme()
	return map[string]*Scene{
		"gear":         NewScene(GearFrames(), 8),
		"steam":        NewScene(SteamFrames(), 8),
		"piston":       PistonSceneWithTheme(theme),
		"smoke":        SmokeSceneWithTheme(theme),
		"gauge":        GaugeSceneWithTheme(theme),
		"valve":        ValveSceneWithTheme(theme),
		"boiler":       BoilerSceneWithTheme(theme),
		"clock":        ClockSceneWithTheme(theme),
		"pendulum":     PendulumSceneWithTheme(theme),
		"telegraph":    TelegraphSceneWithTheme(theme),
		"furnace":      FurnaceSceneWithTheme(theme),
		"pipe":         PipeSceneWithTheme(theme),
		"condenser":    CondenserSceneWithTheme(theme),
		"flywheel":     FlywheelSceneWithTheme(theme),
		"lever":        LeverSceneWithTheme(theme),
		"turbine":      TurbineSceneWithTheme(theme),
		"pressure":     PressureSceneWithTheme(theme),
		"relief_valve": ReliefValveSceneWithTheme(theme),
		"bellows":      BellowsSceneWithTheme(theme),
		"thermometer":  ThermometerSceneWithTheme(theme),
		"anemometer":   AnemometerSceneWithTheme(theme),
		"barometer":    BarometerSceneWithTheme(theme),
		"escapement":   EscapementSceneWithTheme(theme),
		"dynamo":       DynamoSceneWithTheme(theme),
		"sextant":      SextantSceneWithTheme(theme),
		"compass":      CompassSceneWithTheme(theme),
		"chronometer":  ChronometerSceneWithTheme(theme),
	}
}

// FindScene returns the Scene registered under name, or nil if not found.
func FindScene(name string) *Scene {
	scenes := AllScenes()
	s, ok := scenes[name]
	if !ok {
		return nil
	}
	return s
}
