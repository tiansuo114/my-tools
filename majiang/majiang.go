package majiang

import (
	"math/rand"
)

type Card struct {
	Value string
	Count int
}

type Player struct {
	Hand []Card
}

func (p *Player) Draw(deck *Deck) {
	card := deck.Draw()
	if card.Value == "万能牌" {
		index := rand.Intn(len(deck.Cards) - 1)
		card = deck.Cards[index]
		deck.Cards = append(deck.Cards[:index], deck.Cards[index+1:]...)
	}

	p.Hand = append(p.Hand, card)
}

func containsAll(arr1 []Card, arr2 []Card) bool {
	for _, v2 := range arr2 {
		found := false
		for _, v1 := range arr1 {
			if v1 == v2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Drop函数会清空player的手牌
func (p *Player) Drop() {
	p.Hand = p.Hand[:0]

}

func (p *Player) CheckWin() bool {
	// TODO: Implement win checking logic.
	// 统计同类牌数量
	counters := map[string]int{}
	for _, card := range p.Hand {
		counters[card.Value]++
	}
	// 检查同色组
	var sets int
	for _, count := range counters {
		if count >= 3 {
			sets++
		}
	}

	// 检查副对子数
	var pairs int
	for _, count := range counters {
		if count == 2 {
			pairs++
		}
	}

	//十三幺
	if containsAll(p.Hand, []Card{
		Card{Value: "A", Count: 1},
		Card{Value: "I", Count: 1},
		Card{Value: "a", Count: 1},
		Card{Value: "i", Count: 1},
		Card{Value: "1", Count: 1},
		Card{Value: "9", Count: 1},
		Card{Value: "发", Count: 1},
		Card{Value: "中", Count: 1},
		Card{Value: "白", Count: 1},
		Card{Value: "东", Count: 1},
		Card{Value: "南", Count: 1},
		Card{Value: "西", Count: 1},
		Card{Value: "北", Count: 1},
	}) && pairs == 1 {
		return true
	}

	// 全对子
	if pairs == 7 {
		return true
	}

	//普通胡牌 3+3+3+3+2
	if sets >= 4 && pairs == 1 {
		return true
	}

	return false
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	values := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "a", "b", "c", "d", "e", "f", "g", "h", "i", "1", "2", "3", "4", "5", "6", "7", "8", "9", "发", "中", "白", "东", "南", "西", "北", "万能牌"}
	cards := make([]Card, 0, len(values)*4)
	for _, value := range values {
		for i := 0; i < 4; i++ {
			cards = append(cards, Card{Value: value, Count: 1})
		}
	}
	return &Deck{cards}
}

func (d *Deck) Draw() Card {
	index := rand.Intn(len(d.Cards))
	card := d.Cards[index]
	d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
	return card
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}
