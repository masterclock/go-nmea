package nmea

const (
	// TypeGLL type for GLL sentences
	TypeGLL = "GLL"
	// ValidGLL character
	ValidGLL = "A"
	// InvalidGLL character
	InvalidGLL = "V"
)

// GLL is Geographic Position, Latitude / Longitude and time.
// http://aprs.gids.nl/nmea/#gll
type GLL struct {
	BaseSentence
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Time      Time    // Time Stamp
	Validity  string  // validity - A-valid
}

func (s GLL) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"latitude":   s.Latitude,
		"longitude":  s.Longitude,
		"time":       s.Time.String(),
		"time_valid": s.Time.Valid,
		"validity":   s.Validity,
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

// newGLL constructor
func newGLL(s BaseSentence) (GLL, error) {
	p := NewParser(s)
	p.AssertType(TypeGLL)
	return GLL{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Time:         p.Time(4, "time"),
		Validity:     p.EnumString(5, "validity", ValidGLL, InvalidGLL),
	}, p.Err()
}
