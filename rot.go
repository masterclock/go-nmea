package nmea

const (
	// TypeROT type for ROT sentences
	TypeROT = "ROT"
)

// ROT rate of turn
// http://aprs.gids.nl/nmea/#hdt
type ROT struct {
	BaseSentence
	Rate   float64 // rate of turn, degrees/minute, "-" bow turns to port
	Status string
}

func (s ROT) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"rate":   s.Rate,
		"status": s.Status,
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

// newROT constructor
func newROT(s BaseSentence) (ROT, error) {
	p := NewParser(s)
	p.AssertType(TypeROT)
	m := ROT{
		BaseSentence: s,
		Rate:         p.Float64(0, "Rate"),
		Status:       p.String(1, "Status"),
	}
	return m, p.Err()
}
