package main

import "sort"


type Player struct {
    Name string
    Goals int
    Misses int
    Assists int
    Rating float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	return Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
		Rating: calculateRating(goals, misses, assists),
	}
}

func calculateRating(goals, misses, assists int) float64 {
	var rate float64
	if misses == 0 {
		rate = float64(goals) + float64(assists) / 2
		return rate
	}
	rate = (float64(goals) + float64(assists) / 2) / float64(misses)
	return rate
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals == players[j].Goals {
            return players[i].Name < players[j].Name
        }
        return players[i].Goals > players[j].Goals
    })
    return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Rating == players[j].Rating {
			return players[i].Name < players[j].Name
		}
		return players[i].Rating > players[j].Rating
	})
	return players
}

func gmSort(players []Player) []Player {
    sort.Slice(players, func(i, j int) bool {
        var ratioI, ratioJ float64
        if players[i].Misses == 0 {
            ratioI = float64(players[i].Goals) * 1000
        } else {
            ratioI = float64(players[i].Goals) / float64(players[i].Misses)
        }
        if players[j].Misses == 0 {
            ratioJ = float64(players[j].Goals) * 1000
        } else {
            ratioJ = float64(players[j].Goals) / float64(players[j].Misses)
        }

        if ratioI == ratioJ {
            return players[i].Name < players[j].Name
        }
        return ratioI > ratioJ
    })
    return players
}