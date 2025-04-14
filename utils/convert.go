package utils

import (
	"strconv"
	"strings"
)

func Float32From(v any) float32 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case float32:
		return val
	case float64:
		return float32(val)
	case int:
		return float32(val)
	case int8:
		return float32(val)
	case int16:
		return float32(val)
	case int32:
		return float32(val)
	case int64:
		return float32(val)
	case uint8:
		return float32(val)
	case uint16:
		return float32(val)
	case uint32:
		return float32(val)
	case uint64:
		return float32(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return float32(i)
		}
		return 0
	default:
		return 0
	}
}

func Float64From(v any) float64 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case float32:
		return float64(val)
	case float64:
		return val
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return float64(i)
		}
		return 0
	default:
		return 0
	}
}

func Int8From(v any) int8 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return int8(val)
	case int8:
		return val
	case int16:
		return int8(val)
	case int32:
		return int8(val)
	case int64:
		return int8(val)
	case uint8:
		return int8(val)
	case uint16:
		return int8(val)
	case uint32:
		return int8(val)
	case uint64:
		return int8(val)
	case float32:
		return int8(val)
	case float64:
		return int8(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return int8(i)
		}
		return 0
	default:
		return 0
	}
}

func Int16From(v any) int16 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return int16(val)
	case int8:
		return int16(val)
	case int16:
		return val
	case int32:
		return int16(val)
	case int64:
		return int16(val)
	case uint8:
		return int16(val)
	case uint16:
		return int16(val)
	case uint32:
		return int16(val)
	case uint64:
		return int16(val)
	case float32:
		return int16(val)
	case float64:
		return int16(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return int16(i)
		}
		return 0
	default:
		return 0
	}
}

func Int32From(v any) int32 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return int32(val)
	case int8:
		return int32(val)
	case int16:
		return int32(val)
	case int32:
		return val
	case int64:
		return int32(val)
	case uint8:
		return int32(val)
	case uint16:
		return int32(val)
	case uint32:
		return int32(val)
	case uint64:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return int32(i)
		}
		return 0
	default:
		return 0
	}
}

func Int64From(v any) int64 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return int64(i)
		}
		return 0
	default:
		return 0
	}
}

func Uint8From(v any) uint8 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return uint8(val)
	case int8:
		return uint8(val)
	case int16:
		return uint8(val)
	case int32:
		return uint8(val)
	case int64:
		return uint8(val)
	case uint8:
		return val
	case uint16:
		return uint8(val)
	case uint32:
		return uint8(val)
	case uint64:
		return uint8(val)
	case float32:
		return uint8(val)
	case float64:
		return uint8(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return uint8(i)
		}
		return 0
	default:
		return 0
	}
}

func Uint16From(v any) uint16 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return uint16(val)
	case int8:
		return uint16(val)
	case int16:
		return uint16(val)
	case int32:
		return uint16(val)
	case int64:
		return uint16(val)
	case uint8:
		return uint16(val)
	case uint16:
		return val
	case uint32:
		return uint16(val)
	case uint64:
		return uint16(val)
	case float32:
		return uint16(val)
	case float64:
		return uint16(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return uint16(i)
		}
		return 0
	default:
		return 0
	}
}

func Uint32From(v any) uint32 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return uint32(val)
	case int8:
		return uint32(val)
	case int16:
		return uint32(val)
	case int32:
		return uint32(val)
	case int64:
		return uint32(val)
	case uint8:
		return uint32(val)
	case uint16:
		return uint32(val)
	case uint32:
		return val
	case uint64:
		return uint32(val)
	case float32:
		return uint32(val)
	case float64:
		return uint32(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return uint32(i)
		}
		return 0
	default:
		return 0
	}
}

func Uint64From(v any) uint64 {
	switch val := v.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return uint64(val)
	case int8:
		return uint64(val)
	case int16:
		return uint64(val)
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return val
	case float32:
		return uint64(val)
	case float64:
		return uint64(val)
	case string:
		if i, err := strconv.Atoi(val); err == nil {
			return uint64(i)
		}
		return 0
	default:
		return 0
	}
}

func BoolFrom(v any) bool {
	switch val := v.(type) {
	case bool:
		return val
	case int:
		return val > 0
	case int8:
		return val > 0
	case int16:
		return val > 0
	case int32:
		return val > 0
	case int64:
		return val > 0
	case uint8:
		return val > 0
	case uint16:
		return val > 0
	case uint32:
		return val > 0
	case uint64:
		return val > 0
	case float32:
		return val >= 1
	case float64:
		return val >= 1
	case string:
		return strings.ToLower(val) == "false" || v == "0"
	default:
		return false
	}

}
