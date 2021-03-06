package nmea

const (
	// TypeDBT type for DBT sentences
	TypeDBT = "DBT"
)

// DBT depth below keel
// http://aprs.gids.nl/nmea/#DBT
type DBT struct {
	BaseSentence
	DepthFeet   float64 // depth in feet
	Feet        string  // unit 'f'
	DepthMeters float64 // depth in meters
	Meters      string  // unit 'M'
	DepthFathom float64 // depth in fathom
	Fathom      string  // unit 'F'
}

func (s DBT) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"depth_feed":   s.DepthFeet,
		"feet":         s.Feet,
		"depth_meters": s.DepthMeters,
		"meters":       s.Meters,
		"depth_fathom": s.DepthFathom,
		"fathom":       s.Fathom,
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

// newDBT constructor
func newDBT(s BaseSentence) (DBT, error) {
	p := NewParser(s)
	p.AssertType(TypeDBT)
	m := DBT{
		BaseSentence: s,
		DepthFeet:    p.Float64(0, "DepthFeet"),
		Feet:         p.String(1, "Feet"),
		DepthMeters:  p.Float64(2, "DepthMeters"),
		Meters:       p.String(3, "Meters"),
		DepthFathom:  p.Float64(4, "DepthFathom"),
		Fathom:       p.String(5, "Fathom"),
	}
	return m, p.Err()
}
