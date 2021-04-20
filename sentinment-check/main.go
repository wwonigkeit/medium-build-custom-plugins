package main

import (
    "fmt"
    "github.com/jonreiter/govader"
)

func main() {
   analyzer := govader.NewSentimentIntensityAnalyzer()
   sentiment := analyzer.PolarityScores("I LOVE YOU so much I want to cut you up in pieces and eat your liver")

   fmt.Println("Compound score:", sentiment.Compound)
   fmt.Println("Positive score:", sentiment.Positive)
   fmt.Println("Neutral score:", sentiment.Neutral)
   fmt.Println("Negative score:", sentiment.Negative)
}
