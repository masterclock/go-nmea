package nmea

const (
	// TypeVTG type for VTG sentences
	TypeVTG = "VTG"
)

// VTG represents track & speed data.
// http://aprs.gids.nl/nmea/#vtg
type VTG struct {
	BaseSentence
	TrueTrack        float64
	MagneticTrack    float64
	GroundSpeedKnots float64
	GroundSpeedKPH   float64
}

func (s VTG) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"true_track":         s.TrueTrack,
		"magnetic_track":     s.MagneticTrack,
		"ground_speed_knots": s.GroundSpeedKnots,
		"ground_speed_kph":   s.GroundSpeedKPH,
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

// newVTG parses the VTG sentence into this struct.
// e.g: $GPVTG,360.0,T,348.7,M,000.0,N,000.0,K*43
func newVTG(s BaseSentence) (VTG, error) {
	p := NewParser(s)
	p.AssertType(TypeVTG)
	return VTG{
		BaseSentence:     s,
		TrueTrack:        p.Float64(0, "true track"),
		MagneticTrack:    p.Float64(2, "magnetic track"),
		GroundSpeedKnots: p.Float64(4, "ground speed (knots)"),
		GroundSpeedKPH:   p.Float64(6, "ground speed (km/h)"),
	}, p.Err()
}
