package majiang

import (
	"fmt"
	"testing"
)

func Print_test(p Player) {
	if p.CheckWin() {
		fmt.Println(p.Hand)
		fmt.Println("Player win!")
	} else {
		fmt.Println(p.Hand)
		fmt.Println("Player lose!")
	}
}

func TestPlay(t *testing.T) {
	//Finish Card Test
	Y13 := []Card{
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
		Card{Value: "A", Count: 1},
	}

	L7D := []Card{
		Card{Value: "A", Count: 1},
		Card{Value: "A", Count: 1},
		Card{Value: "B", Count: 1},
		Card{Value: "B", Count: 1},
		Card{Value: "C", Count: 1},
		Card{Value: "C", Count: 1},
		Card{Value: "D", Count: 1},
		Card{Value: "D", Count: 1},
		Card{Value: "E", Count: 1},
		Card{Value: "E", Count: 1},
		Card{Value: "F", Count: 1},
		Card{Value: "F", Count: 1},
		Card{Value: "G", Count: 1},
		Card{Value: "G", Count: 1},
	}

	Normal_Card := []Card{
		Card{Value: "A", Count: 1},
		Card{Value: "A", Count: 1},
		Card{Value: "A", Count: 1},
		Card{Value: "B", Count: 1},
		Card{Value: "B", Count: 1},
		Card{Value: "a", Count: 1},
		Card{Value: "b", Count: 1},
		Card{Value: "c", Count: 1},
		Card{Value: "1", Count: 1},
		Card{Value: "2", Count: 1},
		Card{Value: "3", Count: 1},
		Card{Value: "中", Count: 1},
		Card{Value: "中", Count: 1},
		Card{Value: "中", Count: 1},
	}

	YTL := []Card{
		Card{Value: "A", Count: 1},
		Card{Value: "B", Count: 1},
		Card{Value: "C", Count: 1},
		Card{Value: "D", Count: 1},
		Card{Value: "E", Count: 1},
		Card{Value: "F", Count: 1},
		Card{Value: "G", Count: 1},
		Card{Value: "H", Count: 1},
		Card{Value: "I", Count: 1},
		Card{Value: "中", Count: 1},
		Card{Value: "中", Count: 1},
		Card{Value: "中", Count: 1},
		Card{Value: "1", Count: 1},
		Card{Value: "1", Count: 1},
	}

	Test_Player := Player{}

	Test_Player.Hand = Y13
	Print_test(Test_Player)
	Test_Player.Hand = L7D
	Print_test(Test_Player)
	Test_Player.Hand = Normal_Card
	Print_test(Test_Player)
	Test_Player.Hand = YTL
	Print_test(Test_Player)
	count := 0
	//Random Card Test
	for checkWin := false; !checkWin && count <= 20; {

		deck := NewDeck()
		deck.Shuffle()

		for i := 0; i < 14; i++ {
			Test_Player.Draw(deck)
		}

		if Test_Player.CheckWin() {
			checkWin = true
			fmt.Println(Test_Player.Hand)
			fmt.Println("Player win!")

		} else {
			fmt.Println(Test_Player.Hand)
			fmt.Println("Player lose!")
		}
		Test_Player.Drop()
		count++
	}
}
