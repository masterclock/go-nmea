package nmea

const (
	// TypeDPT type for DPT sentences
	TypeDPT = "DPT"
)

// DPT depth below keel
// http://aprs.gids.nl/nmea/#DPT
type DPT struct {
	BaseSentence
	Depth  float64 // water depth relative to transducer, in meters
	Offset float64 // offset from transducer, in meters
	Range  float64 // maximum range scale in use
}

// newDPT constructor
func newDPT(s BaseSentence) (DPT, error) {
	p := newParser(s)
	p.AssertType(TypeDPT)
	m := DPT{
		BaseSentence: s,
		Depth:        p.Float64(0, "Depth"),
		Offset:       p.Float64(1, "Offset"),
		Range:        p.Float64(2, "Range"),
	}
	return m, p.Err()
}
