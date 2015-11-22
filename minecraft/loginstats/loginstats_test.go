// Tests for loginstats
package loginstats

import (
	"fmt"
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
	err := ReadLog(testLogFile)
	if err != nil {
		t.Error(err)
	}
	for k, v := range userStats {
		fmt.Printf("Username: %s\n", k)
		fmt.Printf("Login Count: %d\n", v.LoginCount)
		fmt.Printf("Total Play Time: %d minutes\n", v.TotalPlayTime)
	}
}
