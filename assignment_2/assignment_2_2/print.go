package assignment_2_2

import "fmt"

func ShowSportsDetails() {
	headFactory, _ := GetTennisFactory("Head")
	wilsonFactory, _ := GetTennisFactory("Wilson")

	headBall := headFactory.makeBall()
	wilsonBall := wilsonFactory.makeBall()

	headRacket := headFactory.makeRacket()
	wilsonRacket := wilsonFactory.makeRacket()

	printBallDetails(headBall)
	printBallDetails(wilsonBall)

	printRacketDetails(headRacket)
	printRacketDetails(wilsonRacket)

}

func printBallDetails(b IBall) {
	fmt.Printf("Logo: %s", b.getLogo())
	fmt.Println()
	fmt.Printf("Material: %s", b.getMaterial())
	fmt.Println()
}

func printRacketDetails(r IRacket) {
	fmt.Printf("Logo: %s", r.getLogo())
	fmt.Println()
	fmt.Printf("Material: %s", r.getMaterial())
	fmt.Println()
}
