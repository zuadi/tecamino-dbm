package models

import (
	"fmt"

	"github.com/zuadi/tecamino-dbm.git/utils"
)

const (
	NONE Type = "NONE"
	BIT  Type = "BIT"
	BYU  Type = "BYU"    // UINT8
	BYS  Type = "BYS"    // INT8
	WOS  Type = "WOS"    // INT16
	WOU  Type = "WOU"    // UINT16
	DWS  Type = "DWS"    // INT32
	DWU  Type = "DWU"    // UINT32
	LOS  Type = "LOS"    // INT64
	LOU  Type = "LOU"    // UINT64
	F32  Type = "F32"    // FLOAT32
	F64  Type = "F64"    // FLOAT64
	STR  Type = "STRING" // STRING
)

type Type string

func (t *Type) DefaultValue(v any) any {
	switch *t {
	case BIT:
		return false
	case BYS, BYU, WOS, WOU, DWS, DWU, LOS, LOU, F32, F64:
		return 0
	case STR:
		return ""
	}
	return nil
}

func (t *Type) ConvertValue(v any) any {
	switch *t {
	case BIT:
		return utils.BoolFrom(v)
	case BYS:
		return utils.Int8From(v)
	case BYU:
		return utils.Uint8From(v)
	case WOS:
		return utils.Int16From(v)
	case WOU:
		return utils.Uint16From(v)
	case DWS:
		return utils.Int32From(v)
	case DWU:
		return utils.Uint32From(v)
	case LOS:
		return utils.Int64From(v)
	case LOU:
		return utils.Uint64From(v)
	case F32:
		return utils.Float32From(v)
	case F64:
		return utils.Float64From(v)
	case STR:
		return fmt.Sprintf("%v", v)
	}
	return nil
}
