package core

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Iso8601TimeRecord struct {
	Created Iso8601Time
}

func TimeMustParse(text string) time.Time {
	t, err := time.Parse(time.RFC3339, text)
	if err != nil {
		panic(err)
	}
	return t
}

func TestIso8601Time(t *testing.T) {
	pairs := map[string]time.Time{
		`{ "Created": "2021-07-30T19:14:33Z" }`:          TimeMustParse("2021-07-30T19:14:33Z"),
		`{ "Created": "2021-07-30T19:14:33-0100" }`:      TimeMustParse("2021-07-30T20:14:33Z"),
		`{ "Created": "2021-07-30T19:14:33+0100" }`:      TimeMustParse("2021-07-30T18:14:33Z"),
		`{ "Created": "2021-07-30T19:14:33.000-01:00" }`: TimeMustParse("2021-07-30T20:14:33Z"),
		`{ "Created": "2021-07-30T19:14:33.000+01:00" }`: TimeMustParse("2021-07-30T18:14:33Z"),
	}

	for input, expected := range pairs {
		var record Iso8601TimeRecord
		err := json.Unmarshal([]byte(input), &record)
		assert.Nil(t, err)
		assert.Equal(t, expected, record.Created.ToTime().UTC())
	}
}
