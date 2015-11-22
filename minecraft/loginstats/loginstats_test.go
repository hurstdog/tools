// Tests for loginstats
package loginstats

import (
	"testing"
	"time"
)

const lineUUID = "[01:41:10] [User Authenticator #5/INFO]: UUID of player Notch is 11111111-1111-1111-1111-111111111111"
const lineLoggedIn = "[01:41:10] [Server thread/INFO]: Notch[/1.1.1.1:49297] logged in with entity id 107534 at (21.41199872833552, 101.90931254763402, -41.42809467288776)"
const lineJoin = "[01:41:10] [Server thread/INFO]: Notch joined the game"
const lineLeft = "[01:28:14] [Server thread/INFO]: Notch left the game"
const lineFell = "[02:01:35] [Server thread/INFO]: Notch fell out of the world"

func xTestParseLine(t *testing.T) {
	exTime, _ := time.Parse("03:04:05", "01:41:10")
	expected := UserAction{UserName: "Notch",
		UUID: "11111111-1111-1111-1111-111111111111",
		Time: exTime}
	result, err := ParseLine(lineUUID)
	if err != nil || expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}

func TestParseLinePart(t *testing.T) {
	exTime, _ := time.Parse("03:04:05", "01:28:14")
	expected := UserAction{UserName: "Notch",
		Time: exTime,
		Part: true,
	}
	result, err := ParseLine(lineLeft)
	if err != nil || expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}

func TestParseLineJoin(t *testing.T) {
	exTime, _ := time.Parse("03:04:05", "01:41:10")
	expected := UserAction{UserName: "Notch",
		Time: exTime,
		Join: true,
	}
	result, err := ParseLine(lineJoin)
	if err != nil || expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}
