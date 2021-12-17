package day06

import "fmt"

type LanternFish struct {
	SpawnTimer int
}

func NewLanternFish(st int) (lf *LanternFish) {
	return &LanternFish{
		SpawnTimer: st,
	}
}

func (lf *LanternFish) Advance() (child *LanternFish) {
	if lf.SpawnTimer == 0 {
		child = NewLanternFish(8)
		lf.SpawnTimer = 6
	} else {
		lf.SpawnTimer--
	}
	return child
}

func (lf *LanternFish) AllDescendantsAfter(day int, m map[string]int) int {
	directDescendants := lf.DirectDescendantsAfter(day)
	key := fmt.Sprintf("(%d,%d)", lf.SpawnTimer, day)

	if n, ok := m[key]; ok {
		return n
	}

	sum := 0
	spawnDay := day

	for i := 0; i < directDescendants; i++ {
		if i == 0 {
			spawnDay -= (lf.SpawnTimer + 1)
		} else {
			spawnDay -= 7
		}
		child := NewLanternFish(8)
		sum = sum + 1 + child.AllDescendantsAfter(spawnDay, m)
	}

	m[key] = sum

	return sum
}

func (lf *LanternFish) DirectDescendantsAfter(day int) int {
	actualDays := day - (lf.SpawnTimer + 1)

	// we cannot spawn within the amount of time
	if actualDays < 0 {
		return 0
	}

	// we can spawn at least once. how many direct descendants?
	remaining := (actualDays / 7)
	sum := 1 + remaining

	return sum
}

func (lf *LanternFish) String() string {
	return fmt.Sprintf("%d", lf.SpawnTimer)
}
