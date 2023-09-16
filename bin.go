package main

import "strconv"

func bin(Text string) string {
	Total, err := strconv.ParseInt(Text, 2, 64)
	if err != nil {
		return err.Error()
	}
	return strconv.Itoa(int(Total))
}
