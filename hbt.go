package nmea

const (
	// TypeHBT type for HBT sentences
	TypeHBT = "HBT"
)

// HBT heartheat supervision sentence
// http://aprs.gids.nl/nmea/#HBT
type HBT struct {
	BaseSentence
	Interval float64 // configured repeat interval (50s)
	Status   string  // equipment status A=normal
	ID       string  // sequential sequence identifier 0-9
}

func (s HBT) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"interval": s.Interval,
		"status":   s.Status,
		"id":       s.ID,
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

// newHBT constructor
func newHBT(s BaseSentence) (HBT, error) {
	p := NewParser(s)
	p.AssertType(TypeHBT)
	m := HBT{
		BaseSentence: s,
		Interval:     p.Float64(0, "Interval"),
		Status:       p.String(1, "Status"),
		ID:           p.String(2, "ID"),
	}
	return m, p.Err()
}
