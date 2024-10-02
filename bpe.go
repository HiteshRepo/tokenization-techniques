package main

import (
	"fmt"
	"strings"
)

const numMerges = 10

// getPairs counts the frequency of consecutive pairs of characters in a given word.
//
// The function takes a slice of strings (word) as input, where each string represents a word.
// It iterates over the characters in the word, creating pairs of consecutive characters.
// The function then counts the frequency of each pair and stores the results in a map.
// The keys of the map are 2-element arrays of strings representing the pairs,
// and the values are integers representing the frequency of each pair.
//
// The function returns a map containing the frequency of each pair.
func getPairs(word []string) map[[2]string]int {
	pairs := make(map[[2]string]int)
	for i := 0; i < len(word)-1; i++ {
		pair := [2]string{word[i], word[i+1]}
		pairs[pair]++
	}
	return pairs
}

// tokenize takes a slice of strings (words) as input and tokenizes each word into a slice of individual characters.
//
// Parameters:
// - words: A slice of strings representing the words to be tokenized.
//
// Returns:
// - A 2D slice of strings representing the tokenized words. Each inner slice contains the individual characters of a word.
func tokenize(words []string) [][]string {
	tokenized := make([][]string, len(words))
	for i, word := range words {
		tokenized[i] = strings.Split(word, "")
	}
	return tokenized
}

// bpe performs Byte Pair Encoding (BPE) on a given slice of words.
// BPE is a simple data compression technique that iteratively merges pairs of consecutive characters
// based on their frequency of occurrence in the corpus.
//
// Parameters:
// - words: A slice of strings representing the words to be tokenized.
// - numMerges: An integer representing the number of merge operations to perform.
//
// Returns:
//   - A 2D slice of strings representing the tokenized words after performing the specified number of merges.
//     Each inner slice contains the individual characters of a word.
func bpe(words []string) [][]string {
	tokenized := tokenize(words)
	for merge := 0; merge < numMerges; merge++ {
		// Count all pairs across all tokenized words
		pairFreqs := make(map[[2]string]int)
		for _, word := range tokenized {
			pairs := getPairs(word)
			for pair, count := range pairs {
				pairFreqs[pair] += count
			}
		}

		// Find the most frequent pair
		var bestPair [2]string
		maxCount := 0
		for pair, count := range pairFreqs {
			if count > maxCount {
				bestPair = pair
				maxCount = count
			}
		}

		if maxCount == 0 {
			break // No more pairs to merge
		}

		// Merge the most frequent pair in all words
		for i, word := range tokenized {
			tokenized[i] = mergePair(word, bestPair)
		}

		fmt.Printf("After merge %d: %v\n", merge+1, tokenized)
	}

	return tokenized
}

// mergePair merges consecutive pairs of characters in a given word based on a specified pair.
// It iterates over the characters in the word, checking for consecutive pairs that match the specified pair.
// If a match is found, the pair is merged into a single token and added to the mergedWord slice.
// If no match is found, the character is added to the mergedWord slice as is.
//
// Parameters:
// - word: A slice of strings representing the word to be processed. Each string represents a character.
// - pair: A 2-element array of strings representing the pair of characters to be merged.
//
// Returns:
//   - A slice of strings representing the word after merging the specified pair.
//     Each string represents a character or a merged pair of characters.
func mergePair(word []string, pair [2]string) []string {
	mergedWord := []string{}
	i := 0
	for i < len(word) {
		if i < len(word)-1 && word[i] == pair[0] && word[i+1] == pair[1] {
			// Merge the pair into one token
			mergedWord = append(mergedWord, word[i]+word[i+1])
			i += 2 // Skip the next character since we merged
		} else {
			mergedWord = append(mergedWord, word[i])
			i++
		}
	}
	return mergedWord
}
