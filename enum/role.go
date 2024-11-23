package enum

type Role string

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
)

func (s Role) String() string {
	switch s {
	case USER:
		return "USER"
	case ADMIN:
		return "ADMIN"
	default:
		return "Unknown"
	}
}
