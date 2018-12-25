package main

import (
	"testing"
)

func Test_fight(t *testing.T) {
	type args struct {
		immuneSystem []*group
		infection    []*group
	}
	tests := []struct {
		name         string
		args         args
		immuneSystem []*group
		infection    []*group
		count        int
	}{
		{
			"test",
			args{
				immuneSystem: []*group{
					&group{
						id:    1,
						units: 17,
						hp:    5390,
						weaknesses: map[string]bool{
							"radiation":   true,
							"bludgeoning": true,
						},
						attackPower: 4507,
						initiative:  2,
						attackType:  "fire"},
					&group{
						id:    2,
						units: 989,
						hp:    1274,
						weaknesses: map[string]bool{
							"bludgeoning": true,
							"slashing":    true,
						},
						immunities: map[string]bool{
							"fire": true,
						},
						attackPower: 25,
						initiative:  3,
						attackType:  "slashing"},
				},
				infection: []*group{
					&group{
						id:    1,
						units: 801,
						hp:    4706,
						weaknesses: map[string]bool{
							"radiation": true,
						},
						attackPower: 116,
						initiative:  1,
						attackType:  "bludgeoning"},
					&group{
						id:    2,
						units: 4485,
						hp:    2961,
						weaknesses: map[string]bool{
							"fire": true,
							"cold": true,
						},
						immunities: map[string]bool{
							"radiation": true,
						},
						attackPower: 12,
						initiative:  4,
						attackType:  "slashing"},
				},
			},
			[]*group{},
			[]*group{
				&group{
					id:    1,
					units: 782,
					hp:    4706,
					weaknesses: map[string]bool{
						"radiation": true,
					},
					attackPower: 116,
					initiative:  1,
					attackType:  "bludgeoning"},
				&group{
					id:    2,
					units: 4434,
					hp:    2961,
					weaknesses: map[string]bool{
						"fire": true,
						"cold": true,
					},
					immunities: map[string]bool{
						"radiation": true,
					},
					attackPower: 12,
					initiative:  4,
					attackType:  "slashing"},
			},
			5216,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fight(tt.args.immuneSystem, tt.args.infection)
			if !equals(got, tt.immuneSystem) {
				t.Errorf("fight() got = %v, want %v", got, tt.immuneSystem)
			}
			if !equals(got1, tt.infection) {
				t.Errorf("fight() got1 = %v, want %v", got1, tt.infection)
			}

			gotSum := countUnits(got) + countUnits(got1)
			if tt.count != gotSum {
				t.Errorf("countUnits() got1 = %d, want %d", gotSum, tt.count)
			}
		})
	}
}

func Test_fightWithBoost(t *testing.T) {
	type args struct {
		boost        int
		immuneSystem []*group
		infection    []*group
	}
	tests := []struct {
		name         string
		args         args
		immuneSystem []*group
		infection    []*group
		count        int
	}{
		{
			"test",
			args{
				boost: 1570,
				immuneSystem: []*group{
					&group{
						id:    1,
						units: 17,
						hp:    5390,
						weaknesses: map[string]bool{
							"radiation":   true,
							"bludgeoning": true,
						},
						attackPower: 4507,
						initiative:  2,
						attackType:  "fire"},
					&group{
						id:    2,
						units: 989,
						hp:    1274,
						weaknesses: map[string]bool{
							"bludgeoning": true,
							"slashing":    true,
						},
						immunities: map[string]bool{
							"fire": true,
						},
						attackPower: 25,
						initiative:  3,
						attackType:  "slashing"},
				},
				infection: []*group{
					&group{
						id:    1,
						units: 801,
						hp:    4706,
						weaknesses: map[string]bool{
							"radiation": true,
						},
						attackPower: 116,
						initiative:  1,
						attackType:  "bludgeoning"},
					&group{
						id:    2,
						units: 4485,
						hp:    2961,
						weaknesses: map[string]bool{
							"fire": true,
							"cold": true,
						},
						immunities: map[string]bool{
							"radiation": true,
						},
						attackPower: 12,
						initiative:  4,
						attackType:  "slashing"},
				},
			},
			[]*group{
				&group{
					id:    2,
					units: 51,
					hp:    1274,
					weaknesses: map[string]bool{
						"bludgeoning": true,
						"slashing":    true,
					},
					immunities: map[string]bool{
						"fire": true,
					},
					attackPower: 25,
					initiative:  3,
					attackType:  "slashing"},
			},
			[]*group{},
			51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boost(tt.args.boost, tt.args.immuneSystem)
			got, got1 := fight(tt.args.immuneSystem, tt.args.infection)
			if !equals(got, tt.immuneSystem) {
				t.Errorf("fight() got = %v, want %v", got, tt.immuneSystem)
			}
			if !equals(got1, tt.infection) {
				t.Errorf("fight() got1 = %v, want %v", got1, tt.infection)
			}

			gotSum := countUnits(got) + countUnits(got1)
			if tt.count != gotSum {
				t.Errorf("countUnits() got1 = %d, want %d", gotSum, tt.count)
			}
		})
	}
}

func equals(a, b []*group) bool {
	if len(a) != len(b) {
		return false
	}
	for i, g := range a {
		if g.id != b[i].id || g.units != b[i].units {
			return false
		}
	}
	return true
}
