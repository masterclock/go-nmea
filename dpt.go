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

func (s DPT) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"depth":  s.Depth,
		"offset": s.Offset,
		"range":  s.Range,
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

// newDPT constructor
func newDPT(s BaseSentence) (DPT, error) {
	p := NewParser(s)
	p.AssertType(TypeDPT)
	m := DPT{
		BaseSentence: s,
		Depth:        p.Float64(0, "Depth"),
		Offset:       p.Float64(1, "Offset"),
		Range:        p.Float64(2, "Range"),
	}
	return m, p.Err()
}
