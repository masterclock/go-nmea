package nmea

const (
	// TypeVDM type for VDM sentences
	TypeVDM = "VDM"

	// TypeVDO type for VDO sentences
	TypeVDO = "VDO"
)

// VDMVDO is a format used to encapsulate generic binary payloads. It is most commonly used
// with AIS data.
// http://catb.org/gpsd/AIVDM.html
type VDMVDO struct {
	BaseSentence
	NumFragments   int64
	FragmentNumber int64
	MessageID      int64
	Channel        string
	Payload        []byte
}

func (s VDMVDO) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"num_fragments":   s.NumFragments,
		"fragment_number": s.FragmentNumber,
		"message_id":      s.MessageID,
		"channel":         s.Channel,
		"payload":         s.Payload,
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

// newVDMVDO constructor
func newVDMVDO(s BaseSentence) (VDMVDO, error) {
	p := NewParser(s)
	m := VDMVDO{
		BaseSentence:   s,
		NumFragments:   p.Int64(0, "number of fragments"),
		FragmentNumber: p.Int64(1, "fragment number"),
		MessageID:      p.Int64(2, "sequence number"),
		Channel:        p.String(3, "channel ID"),
		Payload:        p.SixBitASCIIArmour(4, int(p.Int64(5, "number of padding bits")), "payload"),
	}
	return m, p.Err()
}
