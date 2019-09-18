package nmea

const (
	// TypeRTE type for RTE sentences
	TypeRTE = "RTE"

	// ActiveRoute active route
	ActiveRoute = "c"

	// WaypointList list containing waypoints
	WaypointList = "w"
)

// RTE is a route of waypoints
type RTE struct {
	BaseSentence
	NumberOfSentences         int64    // Number of sentences in sequence
	SentenceNumber            int64    // Sentence number
	ActiveRouteOrWaypointList string   // Current active route or waypoint list
	Name                      string   // Name or number of active route
	Idents                    []string // List of ident of waypoints
}

func (s RTE) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"number_of_sentences":           s.NumberOfSentences,
		"sentence_number":               s.SentenceNumber,
		"active_route_or_waypoint_list": s.ActiveRouteOrWaypointList,
		"name":                          s.Name,
		"idents":                        s.Idents,
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

// newRTE constructor
func newRTE(s BaseSentence) (RTE, error) {
	p := NewParser(s)
	p.AssertType(TypeRTE)
	return RTE{
		BaseSentence:              s,
		NumberOfSentences:         p.Int64(0, "number of sentences"),
		SentenceNumber:            p.Int64(1, "sentence number"),
		ActiveRouteOrWaypointList: p.EnumString(2, "active route or waypoint list", ActiveRoute, WaypointList),
		Name:                      p.String(3, "name or number"),
		Idents:                    p.ListString(4, "ident of waypoints"),
	}, p.Err()
}
