package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type group struct {
	id          int
	units       int
	hp          int
	attackPower int
	attackType  string
	initiative  int
	weaknesses  map[string]bool
	immunities  map[string]bool
	target      *group
	targetBy    *group
}

func (g *group) attackSelectedEnemy() bool {
	if g.target == nil {
		return false
	}
	attacked := false
	damage := g.damageTo(g.target) * g.units
	if damage > 0 {
		killedUnits := damage / g.target.hp
		g.target.units -= killedUnits
		attacked = true
	}
	//After an attack, target is lost
	g.target.targetBy = nil
	g.target = nil
	return attacked
}
func (g *group) clone() *group {
	return &group{
		id:          g.id,
		units:       g.units,
		hp:          g.hp,
		attackPower: g.attackPower,
		attackType:  g.attackType,
		initiative:  g.initiative,
		weaknesses:  g.weaknesses,
		immunities:  g.immunities,
	}
}
func (g *group) damageTo(other *group) int {
	if other.weakTo(g.attackType) {
		return g.attackPower * 2
	}
	return g.attackPower
}
func (g *group) effectivePower() int {
	return g.units * g.attackPower
}

func (g *group) hasUnits() bool {
	return g.units > 0
}
func (g *group) immuneTo(attackType string) bool {
	return g.immunities[attackType]
}

func (g *group) targetEnemy(enemies []*group) {
	var target *group
	maxDamage := 0
	for _, enemy := range enemies {

		if enemy.targetBy != nil || enemy.immuneTo(g.attackType) {
			continue
		}
		damage := g.damageTo(enemy)
		if damage > maxDamage {
			target = enemy
			maxDamage = damage
		} else if damage == maxDamage {
			if enemy.effectivePower() > target.effectivePower() {
				target = enemy
				maxDamage = damage
			} else if enemy.effectivePower() == target.effectivePower() {
				if enemy.initiative > target.initiative {
					target = enemy
					maxDamage = damage
				}
			}
		}
	}
	if target != nil {
		g.target = target
		target.targetBy = g
	}
}

func (g *group) weakTo(attackType string) bool {
	return g.weaknesses[attackType]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	immuneSystem, infection := readInput(f)
	immuneSystemCopy := clone(immuneSystem)
	infectionCopy := clone(infection)
	immuneSystemCopy, infectionCopy = fight(immuneSystemCopy, infectionCopy)
	//print("Immune System", immuneSystemCopy)
	//print("Infection", infectionCopy)

	fmt.Printf("Immune System has %d units left after combat\n", countUnits(immuneSystemCopy))
	fmt.Printf("Infection has %d units left after combat\n", countUnits(infectionCopy))

	fmt.Println("Trying to find optimal boost...")

	is, in := fightUntilImmuneSystemWins(immuneSystem, infection)
	fmt.Printf("Immune System has %d units left after combat\n", countUnits(is))
	fmt.Printf("Infection has %d units left after combat\n", countUnits(in))

}
func fightUntilImmuneSystemWins(immuneSystem, infection []*group) ([]*group, []*group) {

	for i := 2; ; i++ {
		immuneSystemCopy := clone(immuneSystem)
		infectionCopy := clone(infection)
		boost(i, immuneSystemCopy)
		immuneSystemCopy, infectionCopy = fight(immuneSystemCopy, infectionCopy)
		if len(immuneSystemCopy) > 0 && len(infectionCopy) == 0 {
			return immuneSystemCopy, infectionCopy
		}

	}
}

func fight(immuneSystem, infection []*group) ([]*group, []*group) {
	lastImmuneCount := 0
	lastInfectionCount := 0

	for len(immuneSystem) > 0 && len(infection) > 0 {
		//target Selection
		sortByEffectivePower(immuneSystem)
		for _, i := range immuneSystem {
			i.targetEnemy(infection)
		}
		sortByEffectivePower(infection)
		for _, i := range infection {
			i.targetEnemy(immuneSystem)
		}
		//attack
		var groups []*group
		groups = append(groups, immuneSystem...)
		groups = append(groups, infection...)
		sortByInitiative(groups)

		for _, g := range groups {
			g.attackSelectedEnemy()
		}
		infection = filterAlive(infection)
		immuneSystem = filterAlive(immuneSystem)
		infections := countUnits(infection)
		immune := countUnits(immuneSystem)
		if infections == lastInfectionCount && immune == lastImmuneCount {
			break
		}
		lastInfectionCount = infections
		lastImmuneCount = immune
	}
	return immuneSystem, infection
}

func boost(boost int, groups []*group) {
	for _, g := range groups {
		g.attackPower += boost
	}
}

func clone(groups []*group) []*group {
	c := make([]*group, len(groups))
	for i, g := range groups {
		c[i] = g.clone()
	}
	return c
}
func countUnits(groups []*group) int {
	count := 0
	for _, g := range groups {
		count += g.units
	}
	return count
}
func filterAlive(groups []*group) []*group {
	var alive []*group
	for _, g := range groups {
		if g.hasUnits() {
			alive = append(alive, g)
		}
	}
	return alive
}
func sortByInitiative(groups []*group) {
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].initiative > groups[j].initiative
	})
}
func sortByEffectivePower(groups []*group) {
	sort.Slice(groups, func(i, j int) bool {
		ei := groups[i].effectivePower()
		ej := groups[j].effectivePower()
		if ei == ej {
			return groups[i].initiative > groups[j].initiative
		}
		return ei > ej
	})
}

func readInput(r io.Reader) ([]*group, []*group) {
	var immuneSystem []*group
	var infection []*group
	var curGroup *[]*group
	var id int
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		if line == "Immune System:" {
			curGroup = &immuneSystem
			id = 1
			continue
		}
		if line == "Infection:" {
			curGroup = &infection
			id = 1
			continue
		}
		g := newGroup(id, line)
		*curGroup = append(*curGroup, &g)
		id++
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return immuneSystem, infection
}
func newGroup(id int, s string) group {
	g := group{
		id: id,
	}
	open := strings.Index(s, "(")
	close := strings.Index(s, ")")
	line := s
	if open > 0 {
		line = s[:open] + s[close+1:]
	}

	fmt.Sscanf(line, "%d units each with %d hit points with an attack that does %d %s damage at initiative %d", &g.units, &g.hp, &g.attackPower, &g.attackType, &g.initiative)
	if open > 0 {
		rest := s[open : close+1]
		g.weaknesses, g.immunities = parseWI(rest)
	}

	return g
}
func parseWI(s string) (map[string]bool, map[string]bool) {
	weaknesses := make(map[string]bool)
	immunities := make(map[string]bool)

	split := strings.Split(s[1:len(s)-1], ";")
	if len(split) == 2 {
		ww := split[0]
		ii := split[1]
		ww = ww[strings.Index(ww, "to ")+3:]
		ii = ii[strings.Index(ii, "to ")+3:]
		for _, w := range strings.Split(ww, ",") {
			weaknesses[strings.TrimSpace(w)] = true
		}
		for _, i := range strings.Split(ii, ",") {
			immunities[strings.TrimSpace(i)] = true
		}

	} else if strings.Contains(split[0], "weak to") {
		ww := split[0][strings.Index(split[0], "to ")+3:]
		for _, w := range strings.Split(ww, ",") {
			weaknesses[strings.TrimSpace(w)] = true
		}
	} else if strings.Contains(split[0], "immune to") {
		ii := split[0][strings.Index(split[0], "to ")+3:]
		for _, i := range strings.Split(ii, ",") {
			immunities[strings.TrimSpace(i)] = true
		}
	}

	return weaknesses, immunities
}
func print(title string, groups []*group) {
	fmt.Printf("%s:\n", title)

	if len(groups) == 0 {
		fmt.Println("No groups remain.")
	}
	for _, g := range groups {
		fmt.Printf("Group %d contains %d units\n", g.id, g.units)
	}
}
