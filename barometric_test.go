package barometric

import "testing"

func TestWorks(t *testing.T) {
	altitude := 1000.0
	pressure := Pressure(altitude)
	expected := 89874.57050221058

	if !(expected-pressure <= 0.00001) {
		t.Errorf("Expected pressure at %d to be %g, got %g", altitude, expected, pressure)
	}

	altitude = 12000.0
	pressure = Pressure(altitude)
	expected = 19330.435819390874

	if !(expected-pressure < 0.00001) {
		t.Errorf("Expected pressure at %d to be %g, got %g", altitude, expected, pressure)
	}

	altitude = 21000.0
	pressure = Pressure(altitude)
	expected = 4677.887313388965

	if !(expected-pressure < 0.00001) {
		t.Errorf("Expected pressure at %d to be %g, got %g", altitude, expected, pressure)
	}

	altitude = 33000.0
	pressure = Pressure(altitude)
	expected = 748.2293709369584

	if !(expected-pressure < 0.00001) {
		t.Errorf("Expected pressure at %d to be %g, got %g", altitude, expected, pressure)
	}

	altitude = 48000.0
	pressure = Pressure(altitude)
	expected = 97.75776072165853

	if !(expected-pressure < 0.00001) {
		t.Errorf("Expected pressure at %d to be %g, got %g", altitude, expected, pressure)
	}

	altitude = 52000.0
	pressure = Pressure(altitude)
	expected = 58.963154521258964

	if !(expected-pressure < 0.00001) {
		t.Errorf("Expected pressure at %d to be %g, got %g", altitude, expected, pressure)
	}

	altitude = 72000.0
	pressure = Pressure(altitude)
	expected = 3.3748140877068526

	if !(expected-pressure < 0.00001) {
		t.Errorf("Expected pressure at %d to be %g, got %g", altitude, expected, pressure)
	}

}
