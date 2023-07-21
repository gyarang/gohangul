package gohangul

import (
	"bytes"
)

var (
	choArr  = [19]rune{'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
	jungArr = [21]rune{'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ', 'ㅔ', 'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ', 'ㅙ', 'ㅚ', 'ㅛ', 'ㅜ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅠ', 'ㅡ', 'ㅢ', 'ㅣ'}
	jongArr = [28]rune{0, 'ㄱ', 'ㄲ', 'ㄳ', 'ㄴ', 'ㄵ', 'ㄶ', 'ㄷ', 'ㄹ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅁ', 'ㅂ', 'ㅄ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}

	choMap = map[rune]int{
		'ㄱ': 0,
		'ㄲ': 1,
		'ㄴ': 2,
		'ㄷ': 3,
		'ㄸ': 4,
		'ㄹ': 5,
		'ㅁ': 6,
		'ㅂ': 7,
		'ㅃ': 8,
		'ㅅ': 9,
		'ㅆ': 10,
		'ㅇ': 11,
		'ㅈ': 12,
		'ㅉ': 13,
		'ㅊ': 14,
		'ㅋ': 15,
		'ㅌ': 16,
		'ㅍ': 17,
		'ㅎ': 18,
	}
	jungMap = map[rune]int{
		'ㅏ': 0,
		'ㅐ': 1,
		'ㅑ': 2,
		'ㅒ': 3,
		'ㅓ': 4,
		'ㅔ': 5,
		'ㅕ': 6,
		'ㅖ': 7,
		'ㅗ': 8,
		'ㅘ': 9,
		'ㅙ': 10,
		'ㅚ': 11,
		'ㅛ': 12,
		'ㅜ': 13,
		'ㅝ': 14,
		'ㅞ': 15,
		'ㅟ': 16,
		'ㅠ': 17,
		'ㅡ': 18,
		'ㅢ': 19,
		'ㅣ': 20,
	}
	jongMap = map[rune]int{
		'ㄱ': 1,
		'ㄲ': 2,
		'ㄳ': 3,
		'ㄴ': 4,
		'ㄵ': 5,
		'ㄶ': 6,
		'ㄷ': 7,
		'ㄹ': 8,
		'ㄺ': 9,
		'ㄻ': 10,
		'ㄼ': 11,
		'ㄽ': 12,
		'ㄾ': 13,
		'ㄿ': 14,
		'ㅀ': 15,
		'ㅁ': 16,
		'ㅂ': 17,
		'ㅄ': 18,
		'ㅅ': 19,
		'ㅆ': 20,
		'ㅇ': 21,
		'ㅈ': 22,
		'ㅊ': 23,
		'ㅋ': 24,
		'ㅌ': 25,
		'ㅍ': 26,
		'ㅎ': 27,
	}

	consonantsMap = map[rune]struct{}{
		'ㄱ': {},
		'ㄲ': {},
		'ㄳ': {},
		'ㄴ': {},
		'ㄵ': {},
		'ㄶ': {},
		'ㄷ': {},
		'ㄸ': {},
		'ㄹ': {},
		'ㄺ': {},
		'ㄻ': {},
		'ㄼ': {},
		'ㄽ': {},
		'ㄾ': {},
		'ㄿ': {},
		'ㅀ': {},
		'ㅁ': {},
		'ㅂ': {},
		'ㅃ': {},
		'ㅄ': {},
		'ㅅ': {},
		'ㅆ': {},
		'ㅇ': {},
		'ㅈ': {},
		'ㅉ': {},
		'ㅊ': {},
		'ㅋ': {},
		'ㅌ': {},
		'ㅍ': {},
		'ㅎ': {},
	}
	mixedJaumMap = map[string]rune{
		"ㄱㅅ": 'ㄳ',
		"ㄴㅈ": 'ㄵ',
		"ㄴㅎ": 'ㄶ',
		"ㄹㄱ": 'ㄺ',
		"ㄹㅁ": 'ㄻ',
		"ㄹㅂ": 'ㄼ',
		"ㄹㅅ": 'ㄽ',
		"ㄹㅌ": 'ㄾ',
		"ㄹㅍ": 'ㄿ',
		"ㄹㅎ": 'ㅀ',
		"ㅂㅅ": 'ㅄ',
	}
	mixedMoumMap = map[string]rune{
		"ㅗㅏ": 'ㅘ',
		"ㅗㅐ": 'ㅙ',
		"ㅗㅣ": 'ㅚ',
		"ㅜㅓ": 'ㅝ',
		"ㅜㅔ": 'ㅞ',
		"ㅜㅣ": 'ㅟ',
		"ㅡㅣ": 'ㅢ',
	}
	splitMixedJamoMap = map[rune][]rune{
		'ㄳ': {'ㄱ', 'ㅅ'},
		'ㄵ': {'ㄴ', 'ㅈ'},
		'ㄶ': {'ㄴ', 'ㅎ'},
		'ㄺ': {'ㄹ', 'ㄱ'},
		'ㄻ': {'ㄹ', 'ㅁ'},
		'ㄼ': {'ㄹ', 'ㅂ'},
		'ㄽ': {'ㄹ', 'ㅅ'},
		'ㄾ': {'ㄹ', 'ㅌ'},
		'ㄿ': {'ㄹ', 'ㅍ'},
		'ㅀ': {'ㄹ', 'ㅎ'},
		'ㅄ': {'ㅂ', 'ㅅ'},
		'ㅘ': {'ㅗ', 'ㅏ'},
		'ㅙ': {'ㅗ', 'ㅐ'},
		'ㅚ': {'ㅗ', 'ㅣ'},
		'ㅝ': {'ㅜ', 'ㅓ'},
		'ㅞ': {'ㅜ', 'ㅔ'},
		'ㅟ': {'ㅜ', 'ㅣ'},
		'ㅢ': {'ㅡ', 'ㅣ'},
	}
)

// Disassemble 함수는 string 을 받아서 각 글자를 자모음으로 분리한 뒤, rune slice 로 반환합니다.
func Disassemble(input string) []rune {
	var result []rune
	for _, r := range input {
		if IsComplete(r) {
			result = append(result, disassembleHangul(r)...)
			continue
		}
		if !IsComplete(r) {
			result = append(result, splitMixedJamo(r)...)
			continue
		}
	}

	return result
}

// DisassembleAsGroup 함수는 Disassemble 함수와 동일한 기능을 수행하지만, 각 글자를 자모음으로 분리한 뒤, 글자별로 rune slice 로 묶어서 반환합니다.
func DisassembleAsGroup(input string) [][]rune {
	result := make([][]rune, 0, len(input))
	for _, r := range input {
		if IsComplete(r) {
			result = append(result, disassembleHangul(r))
			continue
		}
		split := splitMixedJamo(r)
		result = append(result, split)
	}
	return result
}

// Assemble 함수는 분리된 자모음이 포함된 rune slice 를 받아서, 이를 합쳐서 string 으로 반환합니다.
// * 주의: Assemble 함수는 Disassemble 함수의 역함수가 아닙니다.
func Assemble(input []rune) string {
	var buf bytes.Buffer
	for i := 0; i < len(input); {
		r, step := buildHangul(input[i:])
		buf.WriteRune(r)
		i += step
	}
	return buf.String()
}

// IsComplete 함수는 rune 이 완성형 한글인지 아닌지를 판단합니다.
func IsComplete(r rune) bool {
	return r >= '가' && r <= '힣'
}

func search(hd, nd []rune) int {
	if len(hd) < len(nd) {
		return -1
	}

	for i := 0; i < len(hd)-len(nd)+1; i++ {
		if compareRuneSlice(hd[i:i+len(nd)], nd) {
			return i
		}
	}

	return -1
}

// Search 함수는 haystack과 needle을 Disassemble 한 뒤 haystack 에서 needle과 동일한 부분을 찾아 시작 index를 반환합니다.
// 만약 찾지 못했을 경우 -1을 반환합니다.
func Search(haystack, needle string) int {
	hd := Disassemble(haystack)
	nd := Disassemble(needle)

	return search(hd, nd)
}

// Searcher 는 동일한 문자열에 Search 함수를 반복해서 사용시 불필요한 Disassemble 과정을 생략하기 위한 타입입니다.
type Searcher []rune

// NewSearcher 함수는 반복되는 Search 를 위한 문자열의 Searcher 를 생성합니다.
func NewSearcher(needle string) Searcher {
	return Disassemble(needle)
}

// Searcher.Search 함수는 Search 함수와 동일하게 동작하지만 haystack 값만 매개변수로 받습니다.
func (s Searcher) Search(haystack string) int {
	hd := Disassemble(haystack)
	return search(hd, s)
}

type Range struct {
	Start int
	End   int
}

// RangeSearch 함수는 Search 함수와 동일하게 동작하지만, needle이 포함된 모든 문자열의 실제 index 값을 기준으로 Range 를 반환합니다.
// 만약 찾지 못했을 경우 nil 을 반환합니다.
func RangeSearch(haystack string, needle string) []Range {
	hd := Disassemble(haystack)
	nd := Disassemble(needle)
	ghd := DisassembleAsGroup(haystack)

	var result []Range

	if len(hd) < len(nd) || len(nd) == 0 {
		return result
	}

	ghdIndex := 0
	ghdIndexStart := 0
	for i := 0; i < len(hd)-len(nd)+1; i++ {
		// 실제 string 의 start index 계산
		if i-ghdIndexStart >= len(ghd[ghdIndex]) {
			ghdIndex++
			ghdIndexStart = i
		}

		if compareRuneSlice(hd[i:i+len(nd)], nd) {
			// 실제 string 의 End index 계산
			endGhdIndex := ghdIndex
			endRemain := i - ghdIndexStart + len(nd) - len(ghd[ghdIndex])
			for endRemain > 0 {
				endGhdIndex++
				endRemain -= len(ghd[endGhdIndex])
			}

			result = append(result, Range{ghdIndex, endGhdIndex})
		}
	}

	return result
}

// IsCompleteAll 주어진 문자열이 각각 전부 완성형 한글인지 여부를 반환합니다.
func IsCompleteAll(s string) bool {
	for _, r := range s {
		if !IsComplete(r) {
			return false
		}
	}
	return true
}

// IsConsonant 주어진 rune 이 자음인지 여부를 반환합니다.
func IsConsonant(r rune) bool {
	_, ok := consonantsMap[r]
	return ok
}

// IsConsonantAll 주어진 문자열이 각각 전부 자음인지 여부를 반환합니다.
func IsConsonantAll(s string) bool {
	for _, r := range s {
		if !IsConsonant(r) {
			return false
		}
	}
	return true
}

// IsVowel 주어진 rune 이 모음인지 여부를 반환합니다.
func IsVowel(r rune) bool {
	_, ok := jungMap[r]
	return ok
}

// IsVowelAll 주어진 문자열이 각각 전부 모음인지 여부를 반환합니다.
func IsVowelAll(s string) bool {
	for _, r := range s {
		if !IsVowel(r) {
			return false
		}
	}
	return true
}

// IsCho 주어진 rune 이 초성으로 사용될 수 있는지 여부를 반환합니다.
func IsCho(r rune) bool {
	_, ok := choMap[r]
	return ok
}

// IsChoAll 주어진 문자열이 각각 전부 초성으로 사용될 수 있는지 여부를 반환합니다.
func IsChoAll(s string) bool {
	for _, r := range s {
		if !IsCho(r) {
			return false
		}
	}
	return true
}

// IsJong 주어진 rune 이 종성으로 사용될 수 있는지 여부를 반환합니다.
func IsJong(r rune) bool {
	_, ok := jongMap[r]
	return ok
}

// IsJongAll 주어진 문자열이 각각 전부 종성으로 사용될 수 있는지 여부를 반환합니다.
func IsJongAll(s string) bool {
	for _, r := range s {
		if !IsJong(r) {
			return false
		}
	}
	return true
}

// EndsWithConsonant 주어진 문자열이 각각 전부 자음으로 끝나는지 여부를 반환합니다. 은/는, 이/가 구분에 사용할 수 있습니다.
func EndsWithConsonant(s string) bool {
	runes := []rune(s)
	r := runes[len(runes)-1]
	if IsComplete(r) {
		disassembled := disassembleHangul(r)
		return IsConsonant(disassembled[len(disassembled)-1])
	}

	return IsConsonant(r)
}

// EndsWith 함수는 주어진 문자열 input 이 target 으로 끝나는지 여부를 반환합니다.
// 로/으로 구분에 사용할 수 있습니다.
func EndsWith(input string, target rune) bool {
	disassembled := Disassemble(input)
	return disassembled[len(disassembled)-1] == target
}

func buildHangul(arr []rune) (rune, int) {
	var cho, jung, jong rune

	var generateHangul = func() rune {
		hangul := 44032 + 28*(21*choMap[cho]) + 28*jungMap[jung] + jongMap[jong]
		return rune(hangul)
	}

	if len(arr) == 1 {
		return arr[0], 1
	}

	for step := 1; step <= len(arr); step++ {
		if step == 1 {
			if _, ok := choMap[arr[0]]; ok {
				cho = arr[0]
				continue
			} else if _, ok := jungMap[arr[0]]; ok {
				jung = arr[0]
				continue
			} else {
				return arr[0], step
			}
		} else if step == 2 {
			// 가능한 상황: [초성], [중성]
			if _, ok := jungMap[arr[1]]; ok {
				if cho == 0 {
					// [중성] + 모음 상황
					if isMixed, mixed := canMoumMixed(jung, arr[1]); isMixed {
						return mixed, step
					} else {
						return jung, step - 1
					}
				} else {
					// [초성] + 모음 상황
					jung = arr[1]
					continue
				}
			} else if _, ok := choMap[arr[1]]; ok {
				// [초성] + 자음 상황
				if isMixed, mixed := canJaumMixed(cho, arr[1]); isMixed {
					return mixed, step
				} else {
					return cho, step - 1
				}
			} else {
				return arr[0], step - 1
			}
		} else if step == 3 {
			// 가능한 상황: [초성, 중성], [중성, 중성]
			if _, ok := jungMap[arr[2]]; ok {
				if cho == 0 {
					// [중성, 중성] + 모음 상황 (종료)
					return jung, step - 1
				} else {
					// [초성, 중성] + 모음 상황
					if isMixed, mixed := canMoumMixed(jung, arr[2]); isMixed {
						// [초성, 중성] + 모음 조합 가능
						jung = mixed
						continue
					} else {
						// [초성, 중성] + 모음 조합 불가능 (종료)
						return generateHangul(), step - 1
					}
				}
			} else if _, ok := jongMap[arr[2]]; ok {
				// [초성, 중성] + 종성 상황
				jong = arr[2]
				continue
			} else {
				return generateHangul(), step - 1
			}
		} else if step == 4 {
			// 가능한 상황: [초성, 중성, 종성], [초성, (중성, 중성)]
			if _, ok := jongMap[arr[3]]; ok {
				if jong == 0 {
					// [초성, (중성, 중성)] + 종성 상황
					jong = arr[3]
					continue
				} else {
					// [초성, 중성, 종성] + 종성 상황
					if isMixed, mixed := canJaumMixed(jong, arr[3]); isMixed {
						// [초성, 중성, 종성] + 종성 조합 가능
						jong = mixed
						continue
					} else {
						// [초성, 중성, 종성] + 종성 조합 불가능 (종료)
						return generateHangul(), step - 1
					}
				}
			} else if _, ok := jungMap[arr[3]]; ok {
				if jong == 0 {
					// [초성, (중성, 중성)] + 모음 상황 (종료)
					return generateHangul(), step - 1
				} else {
					// [초성, 중성, 종성] + 모음 상황
					if _, ok := choMap[jong]; ok {
						// 종성을 다음 글자 초성으로 붙일 수 있음
						jong = 0
						return generateHangul(), step - 2
					} else {
						// 종성을 다음 글자 초성으로 붙일 수 없음
						return generateHangul(), step - 1
					}
				}
			} else {
				return generateHangul(), step - 1
			}
		} else if step == 5 {
			// 가능한 상황: [초성, 중성, (종성, 종성)], [초성, (중성, 중성), 종성]
			if _, ok := jongMap[arr[4]]; ok {
				if isMixed, mixed := canJaumMixed(jong, arr[4]); isMixed {
					// [초성, (중성, 중성), 종성] + 종성(조합 가능) 상황 (종료)
					jong = mixed
					continue
				} else {
					// [초성, 중성, (종성, 종성)] + 종성(조합 불가능) 상황 (종료)
					// [초성, (중성, 중성), 종성] + 종성(조합 불가능) 상황 (종료)
					return generateHangul(), step - 1
				}
			} else if _, ok := jungMap[arr[4]]; ok {
				splitJong := splitMixedJamo(jong)
				if len(splitJong) == 1 {
					// [초성, (중성, 중성), 종성] + 모음 상황 (종료)
					jong = 0
					return generateHangul(), step - 2
				} else {
					// [초성, 중성, (종성, 종성)] + 모음 상황 (종료)
					jong = splitJong[0]
					return generateHangul(), step - 2
				}
			} else {
				return generateHangul(), step - 1
			}
		} else if step == 6 {
			// 가능한 상황 [초성, (중성, 중성), (종성, 종성)]
			if _, ok := jungMap[arr[5]]; ok {
				splitJong := splitMixedJamo(jong)
				jong = splitJong[0]
				return generateHangul(), step - 2
			} else {
				return generateHangul(), step - 1
			}
		}
	}

	return generateHangul(), len(arr)
}

func disassembleHangul(r rune) []rune {
	var result []rune
	temp := r - 44032
	choTemp := temp / 588
	jongIdx := temp % 28

	result = append(result, splitMixedJamo(choArr[choTemp])...)
	result = append(result, splitMixedJamo(jungArr[(temp%588)/28])...)
	if jongArr[jongIdx] != 0 {
		result = append(result, splitMixedJamo(jongArr[temp%28])...)
	}

	return result
}

func splitMixedJamo(r rune) []rune {
	if split, ok := splitMixedJamoMap[r]; ok {
		return split
	}
	return []rune{r}
}

func canMoumMixed(r1, r2 rune) (bool, rune) {
	combination := string([]rune{r1, r2})
	if mixed, ok := mixedMoumMap[combination]; ok {
		return true, mixed
	}
	return false, 0
}

func canJaumMixed(r1, r2 rune) (bool, rune) {
	combination := string([]rune{r1, r2})
	if mixed, ok := mixedJaumMap[combination]; ok {
		return true, mixed
	}
	return false, 0
}

func compareRuneSlice(s1, s2 []rune) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
