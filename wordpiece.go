package main

import (
	"fmt"
)

// Done for words in 'corpus'.
// This vocabulary is learned during the pre-training phase of models like BERT or ALBERT.
func prepareWordPieceVocabForCorpus() map[string]bool {
	return map[string]bool{
		"low":   true, // Full word
		"##er":  true, // Suffix in "lower"
		"new":   true, // Prefix in "newest"
		"##est": true, // Suffix in "newest", "widest"
		"wid":   true, // Prefix in "widest"
		"##i":   true, // Character subword for "widest"
		"##d":   true, // Character subword for "widest"
	}
}

// wordPieceTokenize tokenizes a given word into subwords using the WordPiece algorithm.
// It uses a provided vocabulary of subwords to find the longest matching subword at each position in the word.
// If no valid subword is found, it uses "[UNK]" to represent an unknown token.
//
// Parameters:
// - word: The input word to tokenize.
// - vocab: A map containing the vocabulary of subwords. The keys are the subwords and the values are boolean true.
//
// Returns:
// - A slice of strings representing the tokens.
func wordPieceTokenize(word string, vocab map[string]bool) []string {
	tokens := []string{}
	start := 0

	// Iterate through the word, finding the longest matching subword from the vocabulary
	for start < len(word) {
		end := len(word) // Start by trying to match the entire remaining part of the word
		subword := ""

		// Try to find the longest matching subword
		for start < end {
			// If it's not the start of the word, prefix the subword with "##" to mark it as a continuation
			part := word[start:end]
			if start > 0 {
				part = "##" + part
			}

			// Check if the subword is in the vocabulary
			if _, exists := vocab[part]; exists {
				subword = part
				break
			}
			end-- // Reduce the end to find a smaller subword
		}

		// If no valid subword is found, we use "[UNK]" to represent an unknown token
		if subword == "" {
			fmt.Printf("Unknown subword: %s\n", word[start:])
			tokens = append(tokens, "[UNK]")
			break
		}

		// Add the subword to the list of tokens and update the starting position
		tokens = append(tokens, subword)
		start = end
	}

	return tokens
}
