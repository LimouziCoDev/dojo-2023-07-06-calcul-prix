package main

import "fmt"

func CalculPrix(nbrArticle int, prixUnitaire float64, taxe float64) float64 {
	prixHT := float64(nbrArticle) * prixUnitaire
	if prixHT >= 5000.0 {
		prixHT *= 0.95
	} else if prixHT >= 1000.0 {
		prixHT *= 0.97
	}
	prixTTC := prixHT + prixHT*taxe
	return prixTTC
}

func main() {
	fmt.Println("Hello world!")
}
