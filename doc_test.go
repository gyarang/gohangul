package gohangul

func ExampleDisassemble() {
	Disassemble("안녕") // Default.
	Disassemble("꿳")  // Complete hangul with complex consonants and vowels.
	Disassemble("ㄳㅙ") // Complex consonant and vowel.
	// Output: []rune{'ㅇ', 'ㅏ', 'ㄴ', 'ㄴ', 'ㅕ', 'ㅇ'}
	// []rune{'ㄲ', 'ㅜ', 'ㅔ', 'ㄱ', 'ㅅ'}
	// []rune{'ㄱ', 'ㅅ', 'ㅗ', 'ㅐ'}
}

func ExampleDisassembleAsGroup() {
	DisassembleAsGroup("안녕") // Default
	// Output: []rune{{'ㅇ', 'ㅏ', 'ㄴ'}, {'ㄴ', 'ㅕ', 'ㅇ'}}
}

func ExampleAssemble() {
	disa := Disassemble("안녕") // []rune{'ㅇ', 'ㅏ', 'ㄴ', 'ㄴ', 'ㅕ', 'ㅇ'}
	Assemble(disa)

	caution := Disassemble("옽ㅏ") // []rune{'ㅇ', 'ㅗ', 'ㅌ', 'ㅏ'}
	Assemble(caution)            // caution: Assemble is not a inverse function of Disassemble.
	// Output: 안녕
	// 오타
}

func ExampleEndsWith() {
	EndsWith("숯불", 'ㄹ') // 이 고기는 숯불로
	EndsWith("볏짚", 'ㄹ') // 아니 볏짚으로
	// Output: true
	// false
}

func ExampleEndsWithConsonant() {
	EndsWithConsonant("숯불") // 숯불은 향이
	EndsWithConsonant("가스") // 가스는 편리
	// Output: true
	// false
}

func ExampleSearch() {
	Search("달걀", "닭")
	Search("달걀", "알")
	// Output: 0
	// -1
}

func ExampleSearcher_Search() {
	searcher := NewSearcher("달걀")
	searcher.Search("닭")
	searcher.Search("알")
	// Output: 0
	// -1
}

func ExampleRangeSearch() {
	haystack := "간장공장공장장"
	RangeSearch(haystack, "장")
	RangeSearch(haystack, "ㅇ공ㅈ")
	RangeSearch(haystack, "갠")
	// Output: []Range{{1, 1}, {3, 3}, {5, 5}, {6, 6}}
	// []Range{{1, 3}, {3, 5}}
	// nil
}
