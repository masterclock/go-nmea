package nmea

const (
	// TypeHDG type for HDG sentences
	TypeHDG = "HDG"
)

// HDG is the Actual vessel heading in degrees True.
// http://aprs.gids.nl/nmea/#hdt
type HDG struct {
	BaseSentence
	Heading            float64
	Deviation          float64
	DeviationDirection string
	Variation          float64
	VariationDirection string
}

// newHDG constructor
func newHDG(s BaseSentence) (HDG, error) {
	p := newParser(s)
	p.AssertType(TypeHDG)
	m := HDG{
		BaseSentence:       s,
		Heading:            p.Float64(0, "Heading"),
		Deviation:          p.Float64(1, "Deviation"),
		DeviationDirection: p.String(2, "DeviationDirection"),
		Variation:          p.Float64(3, "Variation"),
		VariationDirection: p.String(4, "VariationDirection"),
	}
	return m, p.Err()
}
