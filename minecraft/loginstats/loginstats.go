// Loginstats is a package to parse minecraft log files and provide statistics
// about the usage of a server.
package loginstats

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type UserAction struct {
	UserName string
	UUID     string
	Join     bool
	Part     bool
	Time     time.Time
}

type UserStat struct {
	UserName      string
	UUID          string
	LoginCount    int
	TotalPlayTime int
	LastLogin     time.Time
}

// Map from UserName -> UserStat
type StatMap map[string]*UserStat

var userStats StatMap

// Read a minecraft log and return the statistics from it.
func ReadLog(log string) (StatMap, error) {
	userStats := make(map[string]*UserStat)
	fh, err := os.Open(log)
	if err != nil {
		return userStats, fmt.Errorf("Error opening %s: %q", log, err)
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		action, err := parseLine(scanner.Text())
		if err != nil {
			return userStats, err
		}
		if action.UserName == "" {
			continue
		}
		stats := userStats[action.UserName]
		if stats == nil {
			stats = &UserStat{}
			stats.UserName = action.UserName
			userStats[action.UserName] = stats
			fmt.Printf("Added %q to the map\n", *stats)
		}
		if action.UUID != "" && stats.UUID == "" {
			stats.UUID = action.UUID
		}
		// If this is a join message, count the last login
		if action.Join {
			stats.LastLogin = action.Time
			fmt.Println("Set last login to: ", stats.LastLogin)
		}
		// If this is a part message, do the accounting
		if action.Part && !stats.LastLogin.IsZero() {
			stats.TotalPlayTime += int(action.Time.Sub(stats.LastLogin).Minutes())
			var zero time.Time
			stats.LastLogin = zero
			stats.LoginCount++
		}
		fmt.Println("Current stats: ", stats)
	}
	if err := scanner.Err(); err != nil {
		return userStats, fmt.Errorf("Error reading file %s: %q", log, err)
	}

	return userStats, nil
}

/*
 * Example log lines:
 [01:28:14] [Server thread/INFO]: Notch left the game
 [01:41:10] [User Authenticator #5/INFO]: UUID of player Notch is 11111111-1111-1111-1111-111111111111
 [01:41:10] [Server thread/INFO]: Notch[/1.1.1.1:49297] logged in with entity id 107534 at (21.41199872833552, 101.90931254763402, -41.42809467288776)
 [01:41:10] [Server thread/INFO]: Notch joined the game
 [02:01:35] [Server thread/INFO]: Notch fell out of the world
*/
// parseLine parses a single line of a minecraft server log file returning a
// UserAction.
// Every UserAction returned will have a UserName defined.
func parseLine(line string) (UserAction, error) {
	ret := UserAction{}
	t, err := time.Parse("[03:04:05]", line[:10])
	if err != nil {
		return ret, err
	}
	ret.Time = t
	parts := strings.Split(line, " ")
	if parts[1] == "[Server" {
		if parts[4] == "joined" || parts[4] == "left" {
			ret.UserName = parts[3]
			if parts[4] == "left" {
				ret.Part = true
			} else if parts[4] == "joined" {
				ret.Join = true
			}
		}
	} else {
		// UUID line
		ret.UserName = parts[7]
		ret.UUID = parts[9]
	}
	return ret, nil
}
