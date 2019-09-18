package nmea

const (
	// TypeVHW type for VHW sentences
	TypeVHW = "VHW"
)

// VHW is the Actual vessel heading in degrees True.
// http://aprs.gids.nl/nmea/#hdt
type VHW struct {
	BaseSentence
	HeadingTrue     float64
	True            string
	HeadingMagnetic float64
	Magnetic        string
	SpeedKnots      float64
	Knots           string
	SpeedKph        float64
	Kph             string
}

func (s VHW) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"heading_true":     s.HeadingTrue,
		"true":             s.True,
		"heading_magnetic": s.HeadingMagnetic,
		"magnetic":         s.Magnetic,
		"speed_knots":      s.SpeedKnots,
		"knots":            s.Knots,
		"speed_kph":        s.SpeedKph,
		"kph":              s.Kph,
	}
	bm, err := s.BaseSentence.toMap()
	if err != nil {
		return m, err
	}
	for k, v := range bm {
		m[k] = v
	}
	return m, nil
}

// newVHW constructor
func newVHW(s BaseSentence) (VHW, error) {
	p := NewParser(s)
	p.AssertType(TypeVHW)
	m := VHW{
		BaseSentence:    s,
		HeadingTrue:     p.Float64(0, "HeadingTrue"),
		True:            p.String(1, "True"),
		HeadingMagnetic: p.Float64(2, "HeadingMagnetic"),
		Magnetic:        p.String(3, "Magnetic"),
		SpeedKnots:      p.Float64(4, "speedKnots"),
		Knots:           p.String(5, "Knots"),
		SpeedKph:        p.Float64(6, "SpeedKph"),
		Kph:             p.String(7, "Kph"),
	}
	return m, p.Err()
}
