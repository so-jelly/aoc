package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"time"
)

var playingCards = map[string]int{
	"2": 2, "3": 3, "4": 4,
	"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"T": 10,
	"J": 11, "Q": 12, "K": 13,
	"A": 14,
}

func init() {
	DayFunc[7] = Day7
}

type Hand struct {
	cards  []int
	points int
	bid    int
}

// map of points to hand, avoid duplicates
var hands = make(map[int]Hand)

func Day7(part int, r io.Reader) {
	stime := time.Now()
	if part == 2 {
		playingCards["J"] = 1
	}
	day7processdata(r, part)
	winners()
	fmt.Printf("%dms\n", time.Since(stime).Microseconds())
}

func day7processdata(r io.Reader, part int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		h := newHand(line, part)
		if _, ok := hands[h.points]; ok {
			panic(fmt.Sprintf("duplicate hand points: %+v, %+v", h, hands[h.points]))
		}
		hands[h.points] = h
	}
}

func newHand(s string, part int) Hand {
	hand := Hand{}
	hb := strings.Fields(s)
	for _, v := range hb[0] {
		hand.cards = append(hand.cards, playingCards[string(v)])
	}
	if part == 1 {
		hand = hand.calculatePoints()
	} else {
		hand = hand.d7p2calcHand()
	}
	bidStr := hb[1]
	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		panic(err)
	}
	hand.bid = bid
	return hand
}

func (h Hand) calculatePoints() Hand {
	// point system for hands
	// <kind [0-7]><cards as dealt>
	const (
		cardkinds = 10000000000 * iota
		highCard
		onePair
		twoPair
		threeOfAKind
		fullHouse
		fourOfAKind
		fiveOfAKind
	)

	var points int
	cardCounts := make(map[int]int)

	var maxGroup int // 3,4,5
	var pairs int

	for _, cardValue := range h.cards {
		cardCounts[cardValue]++
		if cardCounts[cardValue] > maxGroup {
			maxGroup = cardCounts[cardValue]
		}
		if cardCounts[cardValue] == 3 {
			pairs--
		}
		if cardCounts[cardValue] == 2 {
			pairs++
		}
	}

	x := 1
	for i := len(h.cards); i > 0; i-- {
		points += h.cards[i-1] * x
		x *= 100
	}
	switch maxGroup {
	case 5:
		points += fiveOfAKind
	case 4:
		points += fourOfAKind
	case 3:
		if pairs > 0 {
			points += fullHouse
		} else {
			points += threeOfAKind
		}
	case 2:
		if pairs > 1 {
			points += twoPair
		} else {
			points += onePair
		}
	default:
	}
	h.points = points
	return h
}

func winners() {
	leaderboard := make([]int, 0)
	for k := range hands {
		leaderboard = append(leaderboard, hands[k].points)
	}
	allWinnings := 0
	slices.Sort(leaderboard)
	// fmt.Println(leaderboard)
	for i, v := range leaderboard {
		h := hands[v]
		win := h.bid * (i + 1)
		fmt.Printf("win hand: %+v, rank: %d, win: %d \n", h, i, win)
		allWinnings += win
	}
	fmt.Println(allWinnings)
}

func (h Hand) d7p2calcHand() Hand {
	// point system for hands
	// <kind [0-7]><cards as dealt>
	const (
		cardkinds = 10000000000 * iota
		highCard
		onePair
		twoPair
		threeOfAKind
		fullHouse
		fourOfAKind
		fiveOfAKind
	)

	var points int
	cardCounts := make(map[int]int)

	var maxGroup int // 3,4,5
	var pairs int
	var wild int
	for _, cardValue := range h.cards {
		if cardValue == playingCards["J"] {
			wild++
		}
		cardCounts[cardValue]++
		if cardCounts[cardValue] > maxGroup {
			maxGroup = cardCounts[cardValue]
		}
		if cardCounts[cardValue] == 3 {
			pairs--
		}
		if cardCounts[cardValue] == 2 {
			pairs++
		}
	}

	x := 1
	for i := len(h.cards); i > 0; i-- {
		points += h.cards[i-1] * x
		x *= 100
	}

	switch maxGroup {
	case 5:
		points += fiveOfAKind
	case 4:
		if wild > 0 {
			points += fiveOfAKind
			break
		}
		points += fourOfAKind
	case 3:
		if pairs > 0 {
			if wild > 0 {
				points += fiveOfAKind
				break
			}
			points += fullHouse
		} else {
			if wild > 0 { // 1
				points += fourOfAKind
				break
			}
			points += threeOfAKind
		}
	case 2:
		if pairs > 1 {
			if wild == 1 {
				points += fullHouse
				break
			}
			if wild == 2 {
				points += fourOfAKind
				break
			}
			points += twoPair
		} else {
			if wild > 0 { // either 2 wild and another card or 1 wild and 1 pair
				points += threeOfAKind
				break
			}
			points += onePair
		}
	default:
		if wild > 0 {
			points += onePair
		}
	}
	h.points = points
	return h
}
