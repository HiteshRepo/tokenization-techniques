package main

import (
	"math"
)

func prepareUnigramVocabForCorpus() []Subword {
	// Define the vocabulary with subwords and their probabilities (in log-probabilities for simplicity)
	return []Subword{
		{"low", -1.0}, // higher probability for the whole word "low"
		{"er", -2.0},  // suffix in "lower"
		{"new", -1.5}, // prefix in "newest"
		{"est", -1.5}, // suffix in "newest" and "widest"
		{"wid", -1.8}, // prefix in "widest"
		{"e", -2.5},   // smaller subword for "e"
		{"t", -2.5},   // smaller subword for "t"
		{"i", -2.5},   // smaller subword for "i"
	}
}

// Subword includes the subword itself and its probability.
type Subword struct {
	Token       string
	Probability float64
}

// unigramTokenize tokenizes a given word into subwords based on a provided vocabulary using the Unigram model.
// The function finds the most likely subword sequence that maximizes the likelihood of the input word.
// If no valid subword is found, it returns an unknown token ("[UNK]").
//
// Parameters:
// - word: The input word to tokenize.
// - vocab: A slice of Subword structs representing the vocabulary of subwords and their probabilities.
//
// Returns:
// - A slice of strings representing the tokenized subwords.
func unigramTokenize(word string, vocab []Subword) []string {
	tokens := []string{}

	// Keep track of the start of the word
	start := 0

	// Iterate over the word and find subwords in the vocabulary
	for start < len(word) {
		bestSubword := ""
		bestProb := -math.MaxFloat64
		end := len(word)

		// Try to find the subword that maximizes the likelihood
		for start < end {
			part := word[start:end]

			// Check if the part exists in the vocabulary and has the highest probability
			for _, subword := range vocab {
				if subword.Token == part && subword.Probability > bestProb {
					bestSubword = subword.Token
					bestProb = subword.Probability
					break
				}
			}
			end-- // Decrease end to check smaller subwords
		}

		// If no valid subword is found, return an unknown token.
		if bestSubword == "" {
			tokens = append(tokens, "[UNK]")
			break
		}

		// Add the best subword to the token list
		tokens = append(tokens, bestSubword)
		start += len(bestSubword) // Move the start index to the next part of the word
	}

	return tokens
}
