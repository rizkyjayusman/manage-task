package constants

const (
	ROLE_PM              = 1
	ROLE_SM              = 2
	ROLE_CM              = 3
	ROLE_SPV             = 4
	ROLE_INVENTORY_ADMIN = 5
	ROLE_PC              = 6
)

func RoleStr(code int) string {
	switch code {
	case ROLE_PM:
		return "PM"
	case ROLE_SM:
		return "SM"
	case ROLE_CM:
		return "CM"
	case ROLE_SPV:
		return "SPV"
	case ROLE_INVENTORY_ADMIN:
		return "INVADM"
	case ROLE_PC:
		return "PC"
	default:
		return ""
	}
}

func RoleSubOrdinate(code int) string {
	switch code {
	case ROLE_PM:
		return "SM"
	case ROLE_SM:
		return "CM"
	case ROLE_CM:
		return "SPV"
	case ROLE_SPV:
		return ""
	case ROLE_INVENTORY_ADMIN:
		return ""
	case ROLE_PC:
		return ""
	default:
		return ""
	}
}
