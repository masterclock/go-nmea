package nmea

const (
	// TypeRMC type for RMC sentences
	TypeRMC = "RMC"
	// ValidRMC character
	ValidRMC = "A"
	// InvalidRMC character
	InvalidRMC = "V"
)

// RMC is the Recommended Minimum Specific GNSS data.
// http://aprs.gids.nl/nmea/#rmc
type RMC struct {
	BaseSentence
	Time      Time    // Time Stamp
	Validity  string  // validity - A-ok, V-invalid
	Latitude  float64 // Latitude
	Longitude float64 // Longitude
	Speed     float64 // Speed in knots
	Course    float64 // True course
	Date      Date    // Date
	Variation float64 // Magnetic variation
}

func (s RMC) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"time":       s.Time.String(),
		"time_valid": s.Time.Valid,
		"validity":   s.Validity,
		"latitude":   s.Latitude,
		"longitude":  s.Longitude,
		"speed":      s.Speed,
		"course":     s.Course,
		"date":       s.Date.String(),
		"date_valid": s.Date.Valid,
		"variation":  s.Variation,
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

// newRMC constructor
func newRMC(s BaseSentence) (RMC, error) {
	p := NewParser(s)
	p.AssertType(TypeRMC)
	m := RMC{
		BaseSentence: s,
		Time:         p.Time(0, "time"),
		Validity:     p.EnumString(1, "validity", ValidRMC, InvalidRMC),
		Latitude:     p.LatLong(2, 3, "latitude"),
		Longitude:    p.LatLong(4, 5, "longitude"),
		Speed:        p.Float64(6, "speed"),
		Course:       p.Float64(7, "course"),
		Date:         p.Date(8, "date"),
		Variation:    p.Float64(9, "variation"),
	}
	if p.EnumString(10, "direction", West, East) == West {
		m.Variation = 0 - m.Variation
	}
	return m, p.Err()
}
