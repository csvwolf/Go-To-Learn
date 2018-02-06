package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type scoreBoard map[string]*team

type team struct {
	name                            string
	played, win, loss, draw, points int
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	board := make(scoreBoard)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if err := board.addGame(line); err != nil {
			return err
		}
	}

	teams := board.getTeams()
	sort.Sort(byScore(teams))
	header := "Team                           | MP |  W |  D |  L |  P\n"
	io.WriteString(writer, header)
	for _, team := range teams {
		io.WriteString(writer, team.String()+"\n")
	}
	return nil
}

func (b scoreBoard) addGame(game string) error {
	fields := strings.Split(game, ";")
	if len(fields) != 3 {
		return fmt.Errorf("Game not properly formatted: %s", game)
	}
	home, homeOk := b[fields[0]]
	away, awayOk := b[fields[1]]
	if !homeOk {
		home = &team{name: fields[0]}
		b[fields[0]] = home
	}
	if !awayOk {
		away = &team{name: fields[1]}
		b[fields[1]] = away
	}
	switch fields[2] {
	case "win":
		home.addWin()
		away.addLoss()
	case "loss":
		home.addLoss()
		away.addWin()
	case "draw":
		home.addDraw()
		away.addDraw()
	default:
		return fmt.Errorf("Unknow win conditions: %s", game)
	}
	return nil
}

func (b scoreBoard) getTeams() []team {
	var teams []team
	for _, team := range b {
		teams = append(teams, *team)
	}
	return teams
}

type byScore []team

func (t byScore) Len() int      { return len(t) }
func (t byScore) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t byScore) Less(i, j int) bool {
	if t[i].points != t[j].points {
		return t[i].points > t[j].points
	} else if t[i].win != t[j].win {
		return t[i].win > t[j].win
	}
	return t[i].name < t[j].name
}

func (t *team) addWin() {
	t.played++
	t.win++
	t.points += 3
}

func (t *team) addLoss() {
	t.played++
	t.loss++
}

func (t *team) addDraw() {
	t.played++
	t.draw++
	t.points++
}

func (t *team) String() string {
	fmtString := "%-31s| %2d | %2d | %2d | %2d | %2d"
	return fmt.Sprintf(fmtString, t.name, t.played, t.win, t.draw, t.loss, t.points)
}
