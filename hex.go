package main

import "strconv"

func hex(Text string) string {
	total, err := strconv.ParseInt(Text, 16, 64)
	if err != nil {
		return err.Error()
	}
	return strconv.Itoa(int(total))
}
