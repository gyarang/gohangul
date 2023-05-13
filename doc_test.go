package gohangul

import "fmt"

func printRunes(runes []rune) {
	for _, r := range runes {
		fmt.Print(string(r))
	}
	fmt.Println()
}

func printRuneGroups(runes [][]rune) {
	for _, slice := range runes {
		for _, r := range slice {
			fmt.Print(string(r))
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func ExampleDisassemble() {
	printRunes(Disassemble("안녕")) // Default.
	printRunes(Disassemble("꿳"))  // Complete hangul with complex consonants and vowels.
	printRunes(Disassemble("ㄳㅙ")) // Complex consonant and vowel.
	// Output: ㅇㅏㄴㄴㅕㅇ
	// ㄲㅜㅔㄱㅅ
	// ㄱㅅㅗㅐ
}

func ExampleDisassembleAsGroup() {
	dism := DisassembleAsGroup("안녕") // Default
	printRuneGroups(dism)            // [][]rune{{'ㅇ', 'ㅏ', 'ㄴ'}, {'ㄴ', 'ㅕ', 'ㅇ'}}
	// Output: ㅇㅏㄴ ㄴㅕㅇ
}

func ExampleAssemble() {
	disa := Disassemble("안녕") // []rune{'ㅇ', 'ㅏ', 'ㄴ', 'ㄴ', 'ㅕ', 'ㅇ'}
	fmt.Println(Assemble(disa))

	caution := Disassemble("옽ㅏ")   // []rune{'ㅇ', 'ㅗ', 'ㅌ', 'ㅏ'}
	fmt.Println(Assemble(caution)) // caution: Assemble is not a inverse function of Disassemble.
	// Output: 안녕
	// 오타
}

func ExampleEndsWith() {
	fmt.Println(EndsWith("숯불", 'ㄹ')) // 이 고기는 숯불로
	fmt.Println(EndsWith("볏짚", 'ㄹ')) // 아니 볏짚으로
	// Output: true
	// false
}

func ExampleEndsWithConsonant() {
	fmt.Println(EndsWithConsonant("숯불")) // 숯불은 향이
	fmt.Println(EndsWithConsonant("가스")) // 가스는 편리
	// Output: true
	// false
}

func ExampleSearch() {
	fmt.Println(Search("달걀", "닭"))
	fmt.Println(Search("달걀", "알"))
	// Output: 0
	// -1
}

func ExampleSearcher_Search() {
	searcher := NewSearcher("닭")
	fmt.Println(searcher.Search("달걀"))
	fmt.Println(searcher.Search("달려라"))
	// Output: 0
	// -1
}

func ExampleRangeSearch() {
	haystack := "간장공장공장장"
	fmt.Println(RangeSearch(haystack, "장"))
	fmt.Println(RangeSearch(haystack, "ㅇ공ㅈ"))
	fmt.Println(RangeSearch(haystack, "갠"))
	// Output: [{1 1} {3 3} {5 5} {6 6}]
	// [{1 3} {3 5}]
	// []
}
