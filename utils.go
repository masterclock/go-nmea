package nmea

func makeSentence(raw string) string {
	ck := xorChecksum(raw[1:])
	return raw + "*" + ck
}
