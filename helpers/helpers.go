package helpers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ReaderToString converts an io.ReadCloser to a string
func ReaderToString(reader io.ReadCloser) string {
	if reader == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return ""
	}
	newString := buf.String()
	_, err = fmt.Printf(newString)
	if err != nil {
		return ""
	}
	return newString
}

func getTimestampAndOffset(regex *regexp.Regexp, timeString string) (int64, int64, error) {
	separatedValues := regex.Split(timeString, 2)
	timestamp, err := strconv.ParseInt(separatedValues[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	offset, err := strconv.ParseInt(separatedValues[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return timestamp, offset, nil
}

// DotNetJSONTimeToRFC3339 Converts the .Net formatted time returned by the Xero API to a more readable format
func DotNetJSONTimeToRFC3339(jsonTime string, isUTC bool) (string, error) {
	//If jsonTime is empty we don't need to convert anything
	if jsonTime == "" {
		return "", nil
	}
	//The format returned looks like: /Date(1494201600000+0000)/
	//so first we need to strip out the unnecessary /'s, brackets, and letters
	numbersAndPlusSymbol := regexp.MustCompile("[0-9+-]")
	jsonSlice := numbersAndPlusSymbol.FindAllString(jsonTime, -1)
	//Then we join the resulting array into a string
	jsonTimeString := strings.Join(jsonSlice[:], "")
	//if the offset (the bit after the Unix timestamp) is positive (signalled by a + symbol)
	//then we need to add it to the timestamp and return the result
	var golangTime time.Time
	if strings.Contains(jsonTimeString, "+") {
		plusSymbol := regexp.MustCompile("\\+")
		timestamp, offset, err := getTimestampAndOffset(plusSymbol, jsonTimeString)
		if err != nil {
			return time.Now().Format(time.RFC3339), err
		}
		golangTime = time.Unix((timestamp/1000)+offset, 0)
	} else
	//if the offset (the bit after the Unix timestamp) is negative (signalled by a - symbol)
	//then we need to subtract it from the timestamp and return the result
	if strings.Contains(jsonTimeString, "-") {
		minusSymbol := regexp.MustCompile("\\-")
		timestamp, offset, err := getTimestampAndOffset(minusSymbol, jsonTimeString)
		if err != nil {
			return time.Now().Format(time.RFC3339), err
		}
		golangTime = time.Unix((timestamp/1000)-offset, 0)
	} else {
		//If there is no offset then we just return the converted timestamp
		timestamp, err := strconv.ParseInt(jsonTimeString, 10, 64)
		if err != nil {
			return time.Now().Format(time.RFC3339), err
		}
		golangTime = time.Unix((timestamp / 1000), 0)
	}
	formattedTime := golangTime.UTC().Format(time.RFC3339)
	//The Xero API does not expect an offset. We either need to supply the local time
	//or the UTC time. If we designate the time format as local golang will add the offset
	//but if we designate it as UTC it adds a Z suffix. To satisfy the API requirements we
	//will remove the Z from local times so they aren't seen as UTC times.
	if isUTC {
		return formattedTime, nil
	} else {
		return strings.TrimSuffix(formattedTime, "Z"), nil
	}
}

// TodayRFC3339 returns an RFC3339 formatted date
// with a 0 valued time as required by many Xero endpoints
func TodayRFC3339() string {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return strings.TrimSuffix(today.Format(time.RFC3339), "Z")
}

// FormatDate returns an RFC3339 formatted date
// with a 0 valued time as required by many Xero endpoints
func FormatDate(date time.Time) string {
	d := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	return strings.TrimSuffix(d.Format(time.RFC3339), "Z")
}

func ReadPrivateKeyFromPath(privateKeyFilePath string) string {
	if privateKeyFilePath == "" {
		return ""
	}

	privateKeyFileContents, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(privateKeyFileContents)
}

func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
