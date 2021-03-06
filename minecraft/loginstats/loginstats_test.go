// Tests for loginstats
package loginstats

import (
	"os"
	"reflect"
	"testing"
	"time"
)

const testLogFile = "test.log"

const lineUUID = "[01:41:10] [User Authenticator #5/INFO]: UUID of player Notch is 11111111-1111-1111-1111-111111111111"
const lineLoggedIn = "[01:41:10] [Server thread/INFO]: Notch[/1.1.1.1:49297] logged in with entity id 107534 at (21.41199872833552, 101.90931254763402, -41.42809467288776)"
const lineJoin = "[01:41:10] [Server thread/INFO]: Notch joined the game"
const lineLeft = "[01:28:14] [Server thread/INFO]: Notch left the game"
const lineFell = "[02:01:35] [Server thread/INFO]: Notch fell out of the world"

func TestparseLineUUID(t *testing.T) {
	exTime, _ := time.Parse("03:04:05", "01:41:10")
	expected := UserAction{UserName: "Notch",
		UUID: "11111111-1111-1111-1111-111111111111",
		Time: exTime}
	result, err := parseLine(lineUUID)
	if err != nil || expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}

func TestparseLinePart(t *testing.T) {
	exTime, _ := time.Parse("03:04:05", "01:28:14")
	expected := UserAction{UserName: "Notch",
		Time: exTime,
		Part: true,
	}
	result, err := parseLine(lineLeft)
	if err != nil || expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}

func TestparseLineJoin(t *testing.T) {
	exTime, _ := time.Parse("03:04:05", "01:41:10")
	expected := UserAction{UserName: "Notch",
		Time: exTime,
		Join: true,
	}
	result, err := parseLine(lineJoin)
	if err != nil || expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}

func TestReadLog(t *testing.T) {
	fh, err := os.Open(testLogFile)
	if err != nil {
		t.Errorf("Error opening %s: %q", testLogFile, err)
	}
	err = ReadLog(fh)
	if err != nil {
		t.Error(err)
	}
	exStats := make(StatMap)
	p1 := UserStat{
		UserName:      "Player1",
		UUID:          "1-1-1-1-1",
		LoginCount:    3,
		TotalPlayTime: 3,
	}
	exStats[p1.UserName] = &p1

	p2 := UserStat{
		UserName:      "Player2",
		UUID:          "2-2-2-2-2",
		LoginCount:    1,
		TotalPlayTime: 99,
	}
	exStats[p2.UserName] = &p2

	p3 := UserStat{
		UserName:      "Player3",
		UUID:          "3-3-3-3-3",
		LoginCount:    2,
		TotalPlayTime: 57,
	}
	exStats[p3.UserName] = &p3

	if !reflect.DeepEqual(userStats, exStats) {
		t.Errorf("Result: %q\nExpected: %q\n", userStats, exStats)
	}
}
