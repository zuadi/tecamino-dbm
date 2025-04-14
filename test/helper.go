package test

import (
	"math/rand"

	"github.com/zuadi/tecamino-dbm.git/models"
)

func RandomType() models.Type {
	n := rand.Intn(11) + 1

	switch n {
	case 1:
		return "BIT"
	case 2:
		return "BYU"
	case 3:
		return "BYS"
	case 4:
		return "WOS"
	case 5:
		return "WOU"
	case 6:
		return "DWS"
	case 7:
		return "DWU"
	case 8:
		return "LOS"
	case 9:
		return "LOU"
	case 10:
		return "F32"
	case 11:
		return "F64"
	case 12:
		return "STRING"
	default:
		return "NONE"
	}

}
