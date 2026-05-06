package steampunk

// AllScenes returns a map of all named scenes available in the steampunk package.
func AllScenes() map[string]*Scene {
	return map[string]*Scene{
		"gear":        NewScene(GearFrames(), DefaultFPS),
		"steam":       NewScene(SteamFrames(), DefaultFPS),
		"piston":      PistonScene(),
		"smoke":       SmokeScene(),
		"gauge":       GaugeScene(),
		"valve":       ValveScene(),
		"boiler":      BoilerScene(),
		"clock":       ClockScene(),
		"pendulum":    PendulumScene(),
		"telegraph":   TelegraphScene(),
		"furnace":     FurnaceScene(),
		"pipe":        PipeScene(),
		"condenser":   CondenserScene(),
		"flywheel":    FlywheelScene(),
		"lever":       LeverScene(),
		"turbine":     TurbineScene(),
		"pressure":    PressureScene(),
		"relief":      ReliefValveScene(),
		"bellows":     BellowsScene(),
		"thermometer": ThermometerScene(),
		"anemometer":  AnemometerScene(),
		"barometer":   BarometerScene(),
		"escapement":  EscapementScene(),
		"dynamo":      DynamoScene(),
	}
}

// FindScene looks up a scene by name, returning nil if not found.
func FindScene(name string) *Scene {
	scenes := AllScenes()
	if s, ok := scenes[name]; ok {
		return s
	}
	return nil
}
