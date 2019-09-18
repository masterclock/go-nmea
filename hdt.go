package nmea

const (
	// TypeHDT type for HDT sentences
	TypeHDT = "HDT"
)

// HDT is the Actual vessel heading in degrees True.
// http://aprs.gids.nl/nmea/#hdt
type HDT struct {
	BaseSentence
	Heading float64 // Heading in degrees
	True    bool    // Heading is relative to true north
}

func (s HDT) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"heading": s.Heading,
		"true":    s.True,
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

// newHDT constructor
func newHDT(s BaseSentence) (HDT, error) {
	p := NewParser(s)
	p.AssertType(TypeHDT)
	m := HDT{
		BaseSentence: s,
		Heading:      p.Float64(0, "heading"),
		True:         p.EnumString(1, "true", "T") == "T",
	}
	return m, p.Err()
}
