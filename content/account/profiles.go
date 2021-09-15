package account

import "io"

type Profiles []Profile

func DecodeProfiles(readCloser io.ReadCloser) (profiles []Profile, err error) {
	return
}

func ScanProfiles(scanner interface{ Scan(...interface{}) error }) (profiles []Profile, err error) {
	return
}
