/*
 * ParseUint.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"strconv"
)

func ParseUint(s string) (uint, error) {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(u64), nil
}
