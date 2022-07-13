package models

type CloudAccount int

const (
	_ CloudAccount = iota
	HntCloud
	Fnd
)

func (ca CloudAccount) String() string {
	switch ca {
	case HntCloud:
		return "HNT"
	case Fnd:
		return "FND"
	}
	return ""
}
