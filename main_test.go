package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func TestPasArticle(t *testing.T) {
	prix := CalculPrix(0, 0.0, 0.0)
	assert.EqualValues(t, 0.0, prix, "le prix doit etre nul quand il n'y a pas d'article")
}

func TestUnArticleSansTaxe(t *testing.T) {
	prix := CalculPrix(1, 1.21, 0.0)
	assert.EqualValues(t, 1.21, prix)
}

func TestTroisArticlesSansTaxe(t *testing.T) {
	prix := CalculPrix(3, 1.21, 0.0)
	assert.EqualValues(t, 3.63, prix)
}

func TestTroisArticlesAvecTaxeCinqPct(t *testing.T) {
	prix := CalculPrix(3, 1.21, 0.05)
	assert.EqualValues(t, 3.81, roundFloat(prix, 2))
}

func TestTroisArticlesAvecTaxeVingtPct(t *testing.T) {
	prix := CalculPrix(3, 1.21, 0.20)
	assert.EqualValues(t, 4.36, roundFloat(prix, 2))
}

func TestRemiseTroisPct(t *testing.T) {
	prix := CalculPrix(5, 345.0, 0.10)
	assert.EqualValues(t, 1840.58, roundFloat(prix, 2))
}

func TestRemiseCinqPct(t *testing.T) {
	prix := CalculPrix(5, 1299, 0.10)
	assert.EqualValues(t, 6787.28, roundFloat(prix, 2))
}
