package main

import (
	"fmt"
	"strings"
)

var corpus = []string{"low", "lower", "newest", "widest"}

func main() {
	fmt.Println("\n######### BPE #########\n")

	fmt.Println("Initial tokenized corpus:", tokenize(corpus))
	finalTokens := bpe(corpus)
	fmt.Println("\nFinal BPE tokenized corpus:", finalTokens)

	fmt.Println("\n######### WordPiece #########\n")
	wpVocab := prepareWordPieceVocabForCorpus()
	finalTokens = make([][]string, 0)
	for _, word := range corpus {
		fmt.Printf("Tokenizing word: %s\n", word)
		tokens := wordPieceTokenize(word, wpVocab)
		finalTokens = append(finalTokens, tokens)
		fmt.Printf("Tokens: %v\n\n", tokens)
	}
	fmt.Println("\nFinal WordPiece tokenized corpus:", finalTokens)

	fmt.Println("\n######### Unigram #########\n")
	ugVocab := prepareUnigramVocabForCorpus()
	finalTokens = make([][]string, 0)
	for _, word := range corpus {
		fmt.Printf("Tokenizing word: %s\n", word)
		tokens := unigramTokenize(strings.ToLower(word), ugVocab)
		finalTokens = append(finalTokens, tokens)
		fmt.Printf("Tokens: %v\n\n", tokens)
	}
	fmt.Println("\nFinal Unigram tokenized corpus:", finalTokens)
}
