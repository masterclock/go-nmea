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
	SpeedKmh        float64
	Kmh             string
}

// newVHW constructor
func newVHW(s BaseSentence) (VHW, error) {
	p := newParser(s)
	p.AssertType(TypeVHW)
	m := VHW{
		BaseSentence:    s,
		HeadingTrue:     p.Float64(0, "HeadingTrue"),
		True:            p.String(1, "True"),
		HeadingMagnetic: p.Float64(2, "HeadingMagnetic"),
		Magnetic:        p.String(3, "Magnetic"),
		SpeedKnots:      p.Float64(4, "speedKnots"),
		Knots:           p.String(5, "Knots"),
		SpeedKmh:        p.Float64(6, "SpeedKmh"),
		Kmh:             p.String(7, "Kmh"),
	}
	return m, p.Err()
}
