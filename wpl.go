package nmea

const (
	// TypeWPL type for WPL sentences
	TypeWPL = "WPL"
)

// WPL contains information about a waypoint location
type WPL struct {
	BaseSentence
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Ident     string  // Ident of nth waypoint
}

func (s WPL) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"latitude":  s.Latitude,
		"longitude": s.Longitude,
		"ident":     s.Ident,
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

// newWPL constructor
func newWPL(s BaseSentence) (WPL, error) {
	p := NewParser(s)
	p.AssertType(TypeWPL)
	return WPL{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Ident:        p.String(4, "ident of nth waypoint"),
	}, p.Err()
}
