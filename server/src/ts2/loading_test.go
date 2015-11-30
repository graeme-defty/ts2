package ts2

import (
	"encoding/json"
	"testing"
	"time"
)

func TestLoadOptions(t *testing.T) {
	var sim Simulation
	if err := json.Unmarshal([]byte(simJson), &sim); err != nil {
		t.Errorf("Options: error while loading JSON: %s", err)
	}
	assertEqual(t, sim.Options.CurrentScore, 12, "Options/currentScore")
	sixOClock, _ := time.Parse("15:04:05", "06:00:00")
	assertEqual(t, sim.Options.CurrentTime, Time{sixOClock}, "Options/currentTime")
	assertTrue(t, sim.Options.DefaultDelayAtEntry.equals(DelayGenerator{[]tuplet{{0, 0, 100}}}), "Options/defaultDelayAtEntry")
	assertTrue(t, sim.Options.DefaultMinimumStopTime.equals(DelayGenerator{[]tuplet{{20, 40, 90}, {40, 120, 10}}}), "Options/defaultMinimumStopTime")
	assertEqual(t, sim.Options.DefaultMaxSpeed, 18.06, "Options/defaultMaxSpeed")
	assertEqual(t, sim.Options.DefaultSignalVisibility, 100.0, "Options/defaultSignalVisibility")
	assertEqual(t, sim.Options.Description, "This simulation is a demo sim !", "Options/description")
	assertEqual(t, sim.Options.Title, "Demo Sim", "Options/title")
	assertEqual(t, sim.Options.TimeFactor, 5, "Options/timeFactor")
	assertEqual(t, sim.Options.Version, "0.7", "Options/version")
	assertEqual(t, sim.Options.WarningSpeed, 8.34, "Options/warningSpeed")
	assertEqual(t, sim.Options.TrackCircuitBased, false, "Options/trackCircuitBased")
}

func TestLoadTrackItems(t *testing.T) {
	var sim Simulation
	if err := json.Unmarshal([]byte(simJson), &sim); err != nil {
		t.Errorf("TrackItems: error while loading JSON: %s", err)
	}
	assertEqual(t, len(sim.TrackItems), 22, "TrackItems: Not all loaded")

	ei1, ok := sim.TrackItems[1].(EndItem)
	assertTrue(t, ok, "TrackItems: 1 not loaded")
	li2, ok := sim.TrackItems[2].(LineItem)
	assertTrue(t, ok, "TrackItems: 2 not loaded")
	si3, ok := sim.TrackItems[3].(SignalItem)
	assertTrue(t, ok, "TrackItems: 3 not loaded")
	li4, ok := sim.TrackItems[4].(LineItem)
	assertTrue(t, ok, "TrackItems: 4 not loaded")
	si5, ok := sim.TrackItems[5].(SignalItem)
	assertTrue(t, ok, "TrackItems: 5 not loaded")
	ili6, ok := sim.TrackItems[6].(InvisibleLinkItem)
	assertTrue(t, ok, "TrackItems: 6 not loaded")
	pi7, ok := sim.TrackItems[7].(PointsItem)
	assertTrue(t, ok, "TrackItems: 7 not loaded")
	li8, ok := sim.TrackItems[8].(LineItem)
	assertTrue(t, ok, "TrackItems: 8 not loaded")
	si9, ok := sim.TrackItems[9].(SignalItem)
	assertTrue(t, ok, "TrackItems: 9 not loaded")
	li10, ok := sim.TrackItems[10].(LineItem)
	assertTrue(t, ok, "TrackItems: 10 not loaded")
	si11, ok := sim.TrackItems[11].(SignalItem)
	assertTrue(t, ok, "TrackItems: 11 not loaded")
	li12, ok := sim.TrackItems[12].(LineItem)
	assertTrue(t, ok, "TrackItems: 12 not loaded")
	ei13, ok := sim.TrackItems[13].(EndItem)
	assertTrue(t, ok, "TrackItems: 13 not loaded")
	li14, ok := sim.TrackItems[14].(LineItem)
	assertTrue(t, ok, "TrackItems: 14 not loaded")
	si15, ok := sim.TrackItems[15].(SignalItem)
	assertTrue(t, ok, "TrackItems: 15 not loaded")
	li16, ok := sim.TrackItems[16].(LineItem)
	assertTrue(t, ok, "TrackItems: 16 not loaded")
	si17, ok := sim.TrackItems[17].(SignalItem)
	assertTrue(t, ok, "TrackItems: 15 not loaded")
	ei18, ok := sim.TrackItems[18].(EndItem)
	assertTrue(t, ok, "TrackItems: 18 not loaded")
	lft, ok := sim.Places["LFT"]
	assertTrue(t, ok, "Places: LFT not loaded")
	stn, ok := sim.Places["STN"]
	assertTrue(t, ok, "Places: STN not loaded")
	rgt, ok := sim.Places["RGT"]
	assertTrue(t, ok, "Places: RGT not loaded")
	pfi22, ok := sim.TrackItems[22].(PlatformItem)
	assertTrue(t, ok, "TrackItems: 22 not loaded")
	pfi23, ok := sim.TrackItems[23].(PlatformItem)
	assertTrue(t, ok, "TrackItems: 23 not loaded")
	txti24, ok := sim.TrackItems[24].(TextItem)
	assertTrue(t, ok, "TrackItems: 24 not loaded")
	txti25, ok := sim.TrackItems[25].(TextItem)
	assertTrue(t, ok, "TrackItems: 25 not loaded")

	assertEqual(t, ei1.Name(), "", "EndItem1/Name")
	assertEqual(t, ei1.NextItem(), nil, "EndItem1/NextItem")
	assertEqual(t, ei1.PreviousItem(), li2, "EndItem1/PreviousItem")
	assertEqual(t, ei1.Origin(), Point{0.0, 0.0}, "EndItem1/Origin")
	assertEqual(t, ei1.TiId(), 1, "EndItem/ID")
	assertEqual(t, li2.PreviousItem(), ei1, "LineItem2/PreviousItem")
	assertEqual(t, li2.TrackCode(), "", "LineItem2/TrackCode")
	assertEqual(t, li2.Place(), lft, "LineItem2/Place")
	assertEqual(t, li2.MaxSpeed(), 27.77, "LineItem2/MaxSpeed")
	assertEqual(t, li2.RealLength(), 400.0, "LineItem2/RealLength")

	assertEqual(t, si3.Origin(), li4.Origin(), "Items 3-4/Link")

	assertEqual(t, li4.Name(), "Sample Name", "Lineitem4/Name")
	assertEqual(t, li4.MaxSpeed(), 18.06, "LineItem4/MaxSpeed")

	assertEqual(t, ili6.MaxSpeed(), 10.0, "InvisibleLinkItem6/MaxSpeed")
	assertEqual(t, ili6.RealLength(), 200.0, "InvisibleLinkItem6/RealLength")
	assertEqual(t, ili6.Origin(), Point{200.0, 0.0}, "InvisibleLinkItem6/Origin")
	assertEqual(t, ili6.PreviousItem(), si5, "InvisibleLinkItem6/Previous")

	assertEqual(t, ili6, pi7.PreviousItem(), "Items 6-7/Link")
	assertEqual(t, li8, pi7.NextItem(), "Items 7-8/Link")

	assertEqual(t, si9.Reversed(), true, "SignalItem9/Reversed")

	assertEqual(t, li10.Place(), stn, "LineItem10/Place")
	assertEqual(t, li10.TrackCode(), "1", "LineItem10/TrackCode")

	assertEqual(t, si11.SignalType(), sim.SignalLib.Types["UK_3_ASPECTS"], "SignalItem11/Type")

	assertEqual(t, li12.Place(), rgt, "LineItem12/Place")
	assertEqual(t, li12.TrackCode(), "", "LineItem10/TrackCode")

	assertEqual(t, ei13.Origin(), Point{500.0, 0.0}, "EndItem13/Origin")

	assertEqual(t, pi7.ReverseItem(), li14, "Items 7-14/Link")
	assertEqual(t, pi7.CommonEnd(), Point{-5.0, 0.0}, "PointsItem7/CommonEnd")
	assertEqual(t, pi7.ReverseEnd(), Point{5.0, 5.0}, "PointsItem7/ReverseEnd")
	assertEqual(t, pi7.NormalEnd(), Point{5.0, 0.0}, "PointsItem7/NormalEnd")

	assertEqual(t, si15.Reversed(), true, "SignalItem9/Reversed")
	assertEqual(t, si15.PreviousItem(), li16, "Items 15-16/Link")

	assertEqual(t, li16.Place(), stn, "LineItem16/Place")
	assertEqual(t, li16.TrackCode(), "2", "LineItem16/TrackCode")

	assertEqual(t, si17.Reversed(), false, "SignalItem17/Reversed")
	assertEqual(t, si17.SignalType(), sim.SignalLib.Types["BUFFER"], "SignalItem17/Type")

	assertEqual(t, ei18.PreviousItem(), si17, "EndItem18/PreviousItem")
	assertEqual(t, ei18.NextItem(), nil, "EndItem18/NextItem")

	assertEqual(t, pfi22.Origin(), Point{300, 30}, "PlatformItem22/Origin")
	assertEqual(t, pfi22.End(), Point{400, 45}, "PlatformItem22/End")

	assertEqual(t, pfi23.Place(), stn, "PlatformItem23/Place")
	assertEqual(t, pfi23.TrackCode(), "1", "PlatformItem23/TrackCode")

	assertEqual(t, txti24.Name(), "2", "TextItem24/Name")
	assertEqual(t, txti25.Name(), "1", "TextItem25/Name")
}

func TestLoadSignalLibrary(t *testing.T) {
	var sim Simulation
	if err := json.Unmarshal([]byte(simJson), &sim); err != nil {
		t.Errorf("SignalLibrary: error while loading JSON: %s", err)
	}
	assertEqual(t, len(sim.SignalLib.Types), 2, "SignalLibrary: Not all types loaded")
	assertEqual(t, len(sim.SignalLib.Aspects), 4, "SignalLibrary: Not all aspects loaded")
	bufferAspect, ok := sim.SignalLib.Aspects["BUFFER"]
	if !ok {
		t.Errorf("SignalLibrary: no BUFFER in aspects")
	}
	assertEqual(t, bufferAspect.LineStyle, BUFFER, "SignalLibrary/Aspects/LineStyle")
	assertEqual(t, bufferAspect.OuterShapes, [6]signalShape{NONE, NONE, NONE, NONE, NONE, NONE}, "SignalLibrary/Aspects/OuterShapes")
	black, _ := FromHex("#000000")
	assertEqual(t, bufferAspect.OuterColors, [6]Color{black, black, black, black, black, black}, "SignalLibrary/Aspects/OuterColors")
	assertEqual(t, bufferAspect.Shapes, [6]signalShape{NONE, NONE, NONE, NONE, NONE, NONE}, "SignalLibrary/Aspects/Shapes")
	assertEqual(t, bufferAspect.ShapesColors, [6]Color{black, black, black, black, black, black}, "SignalLibrary/Aspects/Colors")
	assertEqual(t, len(bufferAspect.Actions), 1, "SignalLibrary/Aspects: Not all actions loaded")
	assertEqual(t, bufferAspect.Actions[0].Target, BEFORE_THIS_SIGNAL, "SignalLibrary/Aspects/Actions/Target")
	assertEqual(t, bufferAspect.Actions[0].Speed, 0.0, "SignalLibrary/Aspects/Actions/Speed")
	dangerAspect, ok := sim.SignalLib.Aspects["UK_DANGER"]
	if !ok {
		t.Errorf("SignalLibrary: no UK_DANGER in aspects")
	}
	assertEqual(t, dangerAspect.LineStyle, LINE, "SignalLibrary/Aspects/LineStyle")
	assertEqual(t, dangerAspect.Shapes, [6]signalShape{CIRCLE, NONE, NONE, NONE, NONE, NONE}, "SignalLibrary/Aspects/Shapes")
	red, _ := FromHex("#FF0000")
	assertEqual(t, dangerAspect.ShapesColors, [6]Color{red, black, black, black, black, black}, "SignalLibrary/Aspects/Colors")
	cautionAspect, ok := sim.SignalLib.Aspects["UK_CAUTION"]
	if !ok {
		t.Errorf("SignalLibrary: no UK_CAUTION in aspects")
	}
	assertEqual(t, cautionAspect.Actions[0].Target, BEFORE_NEXT_SIGNAL, "SignalLibrary/Aspects/Actions/Target")
	assertEqual(t, cautionAspect.Actions[0].Speed, 0.0, "SignalLibrary/Aspects/Actions/Speed")
}

func TestInitializeSimulation(t *testing.T) {
	var sim Simulation
	if err := json.Unmarshal([]byte(simJson), &sim); err != nil {
		t.Errorf("SignalLibrary: error while loading JSON: %s", err)
	}
	if err := sim.initialize(); err != nil {
		t.Errorf("Error while initializing simulation: %s", err)
	}
}
