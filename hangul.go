package hangul_go

import (
	"bytes"
)

var (
	choArr        = [19]rune{'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
	jungArr       = [21]rune{'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ', 'ㅔ', 'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ', 'ㅙ', 'ㅚ', 'ㅛ', 'ㅜ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅠ', 'ㅡ', 'ㅢ', 'ㅣ'}
	jongArr       = [28]rune{0, 'ㄱ', 'ㄲ', 'ㄳ', 'ㄴ', 'ㄵ', 'ㄶ', 'ㄷ', 'ㄹ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅁ', 'ㅂ', 'ㅄ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
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
)

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

func Assemble(input []rune) string {
	var buf bytes.Buffer
	for i := 0; i < len(input); {
		r, step := buildHangul(input[i:])
		buf.WriteRune(r)
		i += step
	}
	return buf.String()
}

func IsComplete(r rune) bool {
	return r >= '가' && r <= '힣'
}

func Search(a, b string) int {
	ad := Disassemble(a)
	bd := Disassemble(b)

	if len(ad) < len(bd) {
		return -1
	}

	for i := 0; i < len(ad)-len(bd)+1; i++ {
		if string(ad[i:i+len(bd)]) == string(bd) {
			return i
		}
	}

	return -1
}

func IsCompleteAll(s string) bool {
	for _, r := range s {
		if !IsComplete(r) {
			return false
		}
	}
	return true
}

func IsConsonant(r rune) bool {
	_, ok := consonantsMap[r]
	return ok
}

func IsConsonantAll(s string) bool {
	for _, r := range s {
		if !IsConsonant(r) {
			return false
		}
	}
	return true
}

func IsVowel(r rune) bool {
	_, ok := jungMap[r]
	return ok
}

func IsVowelAll(s string) bool {
	for _, r := range s {
		if !IsVowel(r) {
			return false
		}
	}
	return true
}

func IsCho(r rune) bool {
	_, ok := choMap[r]
	return ok
}

func IsChoAll(s string) bool {
	for _, r := range s {
		if !IsCho(r) {
			return false
		}
	}
	return true
}

func IsJong(r rune) bool {
	_, ok := jongMap[r]
	return ok
}

func IsJongAll(s string) bool {
	for _, r := range s {
		if !IsJong(r) {
			return false
		}
	}
	return true
}

func EndsWithConsonant(s string) bool {
	runes := []rune(s)
	r := runes[len(runes)-1]
	if IsComplete(r) {
		disassembled := disassembleHangul(r)
		return IsConsonant(disassembled[len(disassembled)-1])
	}

	return IsConsonant(r)
}

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
				// 초성 없음
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
	switch r {
	case 'ㄳ':
		return []rune{'ㄱ', 'ㅅ'}
	case 'ㄵ':
		return []rune{'ㄴ', 'ㅈ'}
	case 'ㄶ':
		return []rune{'ㄴ', 'ㅎ'}
	case 'ㄺ':
		return []rune{'ㄹ', 'ㄱ'}
	case 'ㄻ':
		return []rune{'ㄹ', 'ㅁ'}
	case 'ㄼ':
		return []rune{'ㄹ', 'ㅂ'}
	case 'ㄽ':
		return []rune{'ㄹ', 'ㅅ'}
	case 'ㄾ':
		return []rune{'ㄹ', 'ㅌ'}
	case 'ㄿ':
		return []rune{'ㄹ', 'ㅍ'}
	case 'ㅀ':
		return []rune{'ㄹ', 'ㅎ'}
	case 'ㅄ':
		return []rune{'ㅂ', 'ㅅ'}
	case 'ㅘ':
		return []rune{'ㅗ', 'ㅏ'}
	case 'ㅙ':
		return []rune{'ㅗ', 'ㅐ'}
	case 'ㅚ':
		return []rune{'ㅗ', 'ㅣ'}
	case 'ㅝ':
		return []rune{'ㅜ', 'ㅓ'}
	case 'ㅞ':
		return []rune{'ㅜ', 'ㅔ'}
	case 'ㅟ':
		return []rune{'ㅜ', 'ㅣ'}
	case 'ㅢ':
		return []rune{'ㅡ', 'ㅣ'}
	default:
		return []rune{r}
	}
}

func canMoumMixed(r1, r2 rune) (bool, rune) {
	if r1 == 'ㅗ' && r2 == 'ㅏ' {
		return true, 'ㅘ'
	} else if r1 == 'ㅜ' && r2 == 'ㅓ' {
		return true, 'ㅝ'
	} else if r1 == 'ㅗ' && r2 == 'ㅐ' {
		return true, 'ㅙ'
	} else if r1 == 'ㅜ' && r2 == 'ㅔ' {
		return true, 'ㅞ'
	} else if r1 == 'ㅡ' && r2 == 'ㅣ' {
		return true, 'ㅢ'
	} else if r1 == 'ㅗ' && r2 == 'ㅣ' {
		return true, 'ㅚ'
	} else if r1 == 'ㅜ' && r2 == 'ㅣ' {
		return true, 'ㅟ'
	} else {
		return false, 0
	}
}

func canJaumMixed(r1, r2 rune) (bool, rune) {
	if r1 == 'ㄱ' && r2 == 'ㅅ' {
		return true, 'ㄳ'
	} else if r1 == 'ㄴ' && r2 == 'ㅈ' {
		return true, 'ㄵ'
	} else if r1 == 'ㄴ' && r2 == 'ㅎ' {
		return true, 'ㄶ'
	} else if r1 == 'ㄹ' && r2 == 'ㄱ' {
		return true, 'ㄺ'
	} else if r1 == 'ㄹ' && r2 == 'ㅁ' {
		return true, 'ㄻ'
	} else if r1 == 'ㄹ' && r2 == 'ㅂ' {
		return true, 'ㄼ'
	} else if r1 == 'ㄹ' && r2 == 'ㅅ' {
		return true, 'ㄽ'
	} else if r1 == 'ㄹ' && r2 == 'ㅌ' {
		return true, 'ㄾ'
	} else if r1 == 'ㄹ' && r2 == 'ㅍ' {
		return true, 'ㄿ'
	} else if r1 == 'ㄹ' && r2 == 'ㅎ' {
		return true, 'ㅀ'
	} else if r1 == 'ㅂ' && r2 == 'ㅅ' {
		return true, 'ㅄ'
	} else {
		return false, 0
	}
}
