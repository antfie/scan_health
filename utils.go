package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func colorPrintf(format string) {
	color.New().Printf(format)
}

func isStringInStringArray(input string, list []string) bool {
	for _, item := range list {
		if input == item {
			return true
		}
	}

	return false
}

func stringToFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

func dedupeArray[T interface{}](array []T) []T {
	result := []T{}

	for _, item := range array {
		found := false
		for _, processedItem := range result {
			if !found && reflect.DeepEqual(item, processedItem) {
				found = true
			}
		}

		if !found {
			result = append(result, item)
		}
	}

	return result
}

func formatDuration(duration time.Duration) string {
	str := duration.String()

	if duration.Hours() > 24 {
		days := int(duration.Hours() / 24)
		remainingHours := duration.Hours() - float64(days*24)
		str = fmt.Sprintf("%dd %dh%s", days, int(remainingHours), strings.Split(str, "h")[1])
	}

	str = strings.Replace(str, "h", "h ", 1)
	str = strings.Replace(str, "m", "m ", 1)

	return str
}

func printTitle(title string) {
	color.HiCyan(title)
	fmt.Println(strings.Repeat("=", len(title)))
}

func pluralise(count int) string {
	if count > 1 {
		return "s"
	}

	return ""
}

func top5StringList(items []string) string {
	if len(items) > 5 {
		return strings.Join(items[0:5], ", ") + "..."
	}

	return strings.Join(items, ", ")
}
