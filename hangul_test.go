package hangul_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisassemble(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []rune
	}{
		{
			name:  "'가나다'-기본동작",
			input: "가나다",
			want:  []rune{'ㄱ', 'ㅏ', 'ㄴ', 'ㅏ', 'ㄷ', 'ㅏ'},
		},
		{
			name:  "'비행'-받침",
			input: "비행",
			want:  []rune{'ㅂ', 'ㅣ', 'ㅎ', 'ㅐ', 'ㅇ'},
		},
		{
			name:  "'쓸다'-초성에 쌍자음",
			input: "쓸다",
			want:  []rune{'ㅆ', 'ㅡ', 'ㄹ', 'ㄷ', 'ㅏ'},
		},
		{
			name:  "'의사'-중성에 복합모음",
			input: "의사",
			want:  []rune{'ㅇ', 'ㅡ', 'ㅣ', 'ㅅ', 'ㅏ'},
		},
		{
			name:  "닭고기-종성에 복합 자음",
			input: "닭고기",
			want:  []rune{'ㄷ', 'ㅏ', 'ㄹ', 'ㄱ', 'ㄱ', 'ㅗ', 'ㄱ', 'ㅣ'},
		},
		{
			name:  "옽ㅏ",
			input: "옽ㅏ",
			want:  []rune{'ㅇ', 'ㅗ', 'ㅌ', 'ㅏ'},
		},
		{
			name:  "AB삵e$@%2324sdf낄캌ㅋㅋㅋㅋ",
			input: "AB삵e$@%2324sdf낄캌ㅋㅋㅋㅋ",
			want:  []rune{'A', 'B', 'ㅅ', 'ㅏ', 'ㄹ', 'ㄱ', 'e', '$', '@', '%', '2', '3', '2', '4', 's', 'd', 'f', 'ㄲ', 'ㅣ', 'ㄹ', 'ㅋ', 'ㅏ', 'ㅋ', 'ㅋ', 'ㅋ', 'ㅋ', 'ㅋ'},
		},
		{
			name:  "뷁궬릪쯻튋",
			input: "뷁궬릪쯻튋",
			want:  []rune{'ㅂ', 'ㅜ', 'ㅔ', 'ㄹ', 'ㄱ', 'ㄱ', 'ㅜ', 'ㅔ', 'ㄹ', 'ㄹ', 'ㅡ', 'ㅣ', 'ㅍ', 'ㅉ', 'ㅡ', 'ㅣ', 'ㄹ', 'ㅂ', 'ㅌ', 'ㅜ', 'ㅣ', 'ㄹ', 'ㅂ'},
		},
		{
			name:  "ㄳ",
			input: "ㄳ",
			want:  []rune{'ㄱ', 'ㅅ'},
		},
		{
			name:  "ㅙ",
			input: "ㅙ",
			want:  []rune{'ㅗ', 'ㅐ'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Disassemble(tt.input)
			if !assert.Equal(t, tt.want, result) {
				t.Errorf("Disassemble() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestDisassembleAsGroup(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]rune
	}{
		{
			name:  "기본동작-기본동작",
			input: "가나다",
			want:  [][]rune{{'ㄱ', 'ㅏ'}, {'ㄴ', 'ㅏ'}, {'ㄷ', 'ㅏ'}},
		},
		{
			name:  "'비행'-받침",
			input: "비행",
			want:  [][]rune{{'ㅂ', 'ㅣ'}, {'ㅎ', 'ㅐ', 'ㅇ'}},
		},
		{
			name:  "'쓸다'-초성에 쌍자음",
			input: "쓸다",
			want:  [][]rune{{'ㅆ', 'ㅡ', 'ㄹ'}, {'ㄷ', 'ㅏ'}},
		},
		{
			name:  "'의사'-중성에 복합모음",
			input: "의사",
			want:  [][]rune{{'ㅇ', 'ㅡ', 'ㅣ'}, {'ㅅ', 'ㅏ'}},
		}, {
			name:  "'짧은'-종성에 복합자음",
			input: "짧은",
			want:  [][]rune{{'ㅉ', 'ㅏ', 'ㄹ', 'ㅂ'}, {'ㅇ', 'ㅡ', 'ㄴ'}},
		},
		{
			name:  "닭고기-종성에 복합 자음",
			input: "닭고기",
			want:  [][]rune{{'ㄷ', 'ㅏ', 'ㄹ', 'ㄱ'}, {'ㄱ', 'ㅗ'}, {'ㄱ', 'ㅣ'}},
		},
		{
			name:  "옽ㅏ",
			input: "옽ㅏ",
			want:  [][]rune{{'ㅇ', 'ㅗ', 'ㅌ'}, {'ㅏ'}},
		},
		{
			name:  "AB삵e$@%2s낄캌ㅋ",
			input: "AB삵e$@%2s낄캌ㅋ",
			want:  [][]rune{{'A'}, {'B'}, {'ㅅ', 'ㅏ', 'ㄹ', 'ㄱ'}, {'e'}, {'$'}, {'@'}, {'%'}, {'2'}, {'s'}, {'ㄲ', 'ㅣ', 'ㄹ'}, {'ㅋ', 'ㅏ', 'ㅋ'}, {'ㅋ'}},
		},
		{
			name:  "뷁궬릪쯻튋",
			input: "뷁궬릪쯻튋",
			want:  [][]rune{{'ㅂ', 'ㅜ', 'ㅔ', 'ㄹ', 'ㄱ'}, {'ㄱ', 'ㅜ', 'ㅔ', 'ㄹ'}, {'ㄹ', 'ㅡ', 'ㅣ', 'ㅍ'}, {'ㅉ', 'ㅡ', 'ㅣ', 'ㄹ', 'ㅂ'}, {'ㅌ', 'ㅜ', 'ㅣ', 'ㄹ', 'ㅂ'}},
		},
		{
			name:  "ㄳ",
			input: "ㄳ",
			want:  [][]rune{{'ㄱ', 'ㅅ'}},
		},
		{
			name:  "ㅙ",
			input: "ㅙ",
			want:  [][]rune{{'ㅗ', 'ㅐ'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !assert.Equal(t, tt.want, DisassembleAsGroup(tt.input)) {
				t.Errorf("TestDisassembleAsGroup() = %v, want %v", DisassembleAsGroup(tt.input), tt.want)
			}
		})
	}
}

func TestAssemble(t *testing.T) {
	tests := []struct {
		name  string
		input []rune
		want  string
	}{
		{
			name:  "기본동작",
			input: []rune{'ㄱ', 'ㅏ', 'ㄴ', 'ㅏ', 'ㄷ', 'ㅏ'},
			want:  "가나다",
		},
		{
			name:  "비행",
			input: []rune{'ㅂ', 'ㅣ', 'ㅎ', 'ㅐ', 'ㅇ'},
			want:  "비행",
		},
		{
			name:  "쓸다",
			input: []rune{'ㅆ', 'ㅡ', 'ㄹ', 'ㄷ', 'ㅏ'},
			want:  "쓸다",
		},
		{
			name:  "의사",
			input: []rune{'ㅇ', 'ㅡ', 'ㅣ', 'ㅅ', 'ㅏ'},
			want:  "의사",
		},
		{
			name:  "짧은",
			input: []rune{'ㅉ', 'ㅏ', 'ㄹ', 'ㅂ', 'ㅇ', 'ㅡ', 'ㄴ'},
			want:  "짧은",
		},
		{
			name:  "닭고기",
			input: []rune{'ㄷ', 'ㅏ', 'ㄹ', 'ㄱ', 'ㄱ', 'ㅗ', 'ㄱ', 'ㅣ'},
			want:  "닭고기",
		},
		{
			name:  "오타",
			input: []rune{'ㅇ', 'ㅗ', 'ㅌ', 'ㅏ'},
			want:  "오타",
		},
		{
			name:  "AB삵e$@%2324sdf낄캌ㅋㅋㅋㅋ",
			input: []rune{'A', 'B', 'ㅅ', 'ㅏ', 'ㄹ', 'ㄱ', 'e', '$', '@', '%', '2', '3', '2', '4', 's', 'd', 'f', 'ㄲ', 'ㅣ', 'ㄹ', 'ㅋ', 'ㅏ', 'ㅋ', 'ㅋ', 'ㅋ', 'ㅋ', 'ㅋ'},
			want:  "AB삵e$@%2324sdf낄캌ㅋㅋㅋㅋ",
		},
		{
			name:  "뷁궬릪쯻튋",
			input: []rune{'ㅂ', 'ㅜ', 'ㅔ', 'ㄹ', 'ㄱ', 'ㄱ', 'ㅜ', 'ㅔ', 'ㄹ', 'ㄹ', 'ㅡ', 'ㅣ', 'ㅍ', 'ㅉ', 'ㅡ', 'ㅣ', 'ㄹ', 'ㅂ', 'ㅌ', 'ㅜ', 'ㅣ', 'ㄹ', 'ㅂ'},
			want:  "뷁궬릪쯻튋",
		},
		{
			name:  "ㄳ",
			input: []rune{'ㄱ', 'ㅅ'},
			want:  "ㄳ",
		},
		{
			name:  "ㅙ",
			input: []rune{'ㅗ', 'ㅐ'},
			want:  "ㅙ",
		},
		{
			name:  "ㅈ사",
			input: []rune{'ㅈ', 'ㅅ', 'ㅏ'},
			want:  "ㅈ사",
		},
		{
			name:  "ㄳㄳ",
			input: []rune{'ㄱ', 'ㅅ', 'ㄱ', 'ㅅ'},
			want:  "ㄳㄳ",
		},
		{
			name:  "ㅙㅙ",
			input: []rune{'ㅗ', 'ㅐ', 'ㅗ', 'ㅐ'},
			want:  "ㅙㅙ",
		},
		{
			name:  "조ㅙ",
			input: []rune{'ㅈ', 'ㅗ', 'ㅗ', 'ㅐ'},
			want:  "조ㅙ",
		},
		{
			name:  "ㅣㅙ",
			input: []rune{'ㅣ', 'ㅗ', 'ㅐ'},
			want:  "ㅣㅙ",
		},
		{
			name:  "ㅃ짜ㄸ",
			input: []rune{'ㅃ', 'ㅉ', 'ㅏ', 'ㄸ'},
			want:  "ㅃ짜ㄸ",
		},
		{
			name:  "ㅒㅗㅒ",
			input: []rune{'ㅒ', 'ㅗ', 'ㅒ'},
			want:  "ㅒㅗㅒ",
		},
		{
			name:  "쀍ㅅ",
			input: []rune{'ㅃ', 'ㅜ', 'ㅔ', 'ㄹ', 'ㄱ', 'ㅅ'},
			want:  "쀍ㅅ",
		},
		{
			name:  "쀌가",
			input: []rune{'ㅃ', 'ㅜ', 'ㅔ', 'ㄹ', 'ㄱ', 'ㅏ'},
			want:  "쀌가",
		},
		{
			name:  "쀌궬궭ㅂ",
			input: []rune{'ㅃ', 'ㅞ', 'ㄹ', 'ㄱ', 'ㅞ', 'ㄹ', 'ㄱ', 'ㅞ', 'ㄹ', 'ㄱ', 'ㅂ'},
			want:  "쀌궬궭ㅂ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Assemble(tt.input)
			if !assert.Equal(t, tt.want, result) {
				t.Errorf("Assemble() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestIsComplete(t *testing.T) {
	tests := []struct {
		name   string
		input  rune
		result bool
	}{
		{
			name:   "한",
			input:  '한',
			result: true,
		},
		{
			name:   "ㄱ",
			input:  'ㄱ',
			result: false,
		},
		{
			name:   "ㅙ",
			input:  'ㅙ',
			result: false,
		},
		{
			name:   "a",
			input:  'a',
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsComplete(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsComplete() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsConsonant(t *testing.T) {
	tests := []struct {
		name   string
		input  rune
		result bool
	}{
		{
			name:   "한",
			input:  '한',
			result: false,
		},
		{
			name:   "ㄱ",
			input:  'ㄱ',
			result: true,
		},
		{
			name:   "ㅙ",
			input:  'ㅙ',
			result: false,
		},
		{
			name:   "a",
			input:  'a',
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsConsonant(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsConsonant() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsVowel(t *testing.T) {
	tests := []struct {
		name   string
		input  rune
		result bool
	}{
		{
			name:   "한",
			input:  '한',
			result: false,
		},
		{
			name:   "ㄱ",
			input:  'ㄱ',
			result: false,
		},
		{
			name:   "ㅙ",
			input:  'ㅙ',
			result: true,
		},
		{
			name:   "a",
			input:  'a',
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsVowel(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsVowel() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsCho(t *testing.T) {
	tests := []struct {
		name   string
		input  rune
		result bool
	}{
		{
			name:   "ㄱ",
			input:  'ㄱ',
			result: true,
		},
		{
			name:   "ㄸ",
			input:  'ㄸ',
			result: true,
		},
		{
			name:   "ㄳ",
			input:  'ㄳ',
			result: false,
		},
		{
			name:   "ㅏ",
			input:  'ㅏ',
			result: false,
		},
		{
			name:   "a",
			input:  'a',
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCho(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsCho() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsJong(t *testing.T) {
	tests := []struct {
		name   string
		input  rune
		result bool
	}{
		{
			name:   "ㄱ",
			input:  'ㄱ',
			result: true,
		},
		{
			name:   "ㄸ",
			input:  'ㄸ',
			result: false,
		},
		{
			name:   "ㄳ",
			input:  'ㄳ',
			result: true,
		},
		{
			name:   "ㅏ",
			input:  'ㅏ',
			result: false,
		},
		{
			name:   "a",
			input:  'a',
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsJong(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsJong() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsCompleteAll(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "한글",
			input:  "한글",
			result: true,
		},
		{
			name:   "한글ㄱ",
			input:  "한글ㄱ",
			result: false,
		},
		{
			name:   "한글a",
			input:  "한글a",
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCompleteAll(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsCompleteAll() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsConsonantAll(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "한ㄱ",
			input:  "한ㄱ",
			result: false,
		},
		{
			name:   "ㅎㄱ",
			input:  "ㅎㄱ",
			result: true,
		},
		{
			name:   "ㅁa",
			input:  "ㅁa",
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsConsonantAll(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsConsonantAll() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsVowelAll(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "한글",
			input:  "한글",
			result: false,
		},
		{
			name:   "ㅗㄱ",
			input:  "ㅗㄱ",
			result: false,
		},
		{
			name:   "ㅙㅜ",
			input:  "ㅙㅜ",
			result: true,
		},
		{
			name:   "ㅙa",
			input:  "ㅙa",
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsVowelAll(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsVowelAll() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsChoAll(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "ㄱㄴㄷ",
			input:  "ㄱㄴㄷ",
			result: true,
		},
		{
			name:   "ㄸㄲㅆ",
			input:  "ㄸㄲㅆ",
			result: true,
		},
		{
			name:   "ㄱㄴㄳ",
			input:  "ㄱㄴㄳ",
			result: false,
		},
		{
			name:   "ㄱㄴㅏ",
			input:  "ㄱㄴㅏ",
			result: false,
		},
		{
			name:   "ㄱㄴa",
			input:  "ㄱㄴa",
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsChoAll(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsChoAll() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestIsJongAll(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "ㄱㄴ",
			input:  "ㄱㄴ",
			result: true,
		},
		{
			name:   "ㄱㄸ",
			input:  "ㄱㄸ",
			result: false,
		},
		{
			name:   "ㄳㅄ",
			input:  "ㄳㅄ",
			result: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsJongAll(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("IsJongAll() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestEndsWithConsonant(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "ㄱ",
			result: true,
			input:  "ㄱ",
		},
		{
			name:   "가",
			result: false,
			input:  "가",
		},
		{
			name:   "ㅏ",
			result: false,
			input:  "ㅏ",
		},
		{
			name:   "각",
			result: true,
			input:  "각",
		},
		{
			name:   "가각",
			result: true,
			input:  "가각",
		},
		{
			name:   "abc각",
			result: true,
			input:  "abc각",
		},
		{
			name:   "abc",
			result: false,
			input:  "abc",
		},
		{
			name:   "各",
			result: false,
			input:  "各",
		},
		{
			name:   "!@#",
			result: false,
			input:  "!@#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndsWithConsonant(tt.input)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("EndsWithConsonant() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		target rune
		result bool
	}{
		{
			name:   "\"ㄱ\" 'ㄱ'",
			input:  "ㄱ",
			target: 'ㄱ',
			result: true,
		},
		{
			name:   "\"가\" 'ㄱ'",
			input:  "가",
			target: 'ㄱ',
			result: false,
		},
		{
			name:   "\"가\" 'ㅏ'",
			input:  "가",
			target: 'ㅏ',
			result: true,
		},
		{
			name:   "\"각\" 'ㄱ'",
			input:  "각",
			target: 'ㄱ',
			result: true,
		},
		{
			name:   "\"abc\" 'c'",
			input:  "abc",
			target: 'c',
			result: true,
		},
		{
			name:   "\"abc\" 'b'",
			input:  "abc",
			target: 'b',
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndsWith(tt.input, tt.target)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("EndsWith() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		target string
		needle string
		result int
	}{
		{
			name:   "\"도우미\", \"도움\"",
			target: "도우미",
			needle: "도움",
			result: 0,
		},
		{
			name:   "\"달걀\" \"닭\"",
			target: "달걀",
			needle: "닭",
			result: 0,
		},
		{
			name:   "\"도우미\" \"ㅜㅁ\"",
			target: "도우미",
			needle: "ㅜㅁ",
			result: 3,
		},
		{
			name:   "\"달맞이\" \"ㄹ마\"",
			target: "달맞이",
			needle: "ㄹ마",
			result: 2,
		},
		{
			name:   "\"달맞이\" \"ㅁㅈ\"",
			target: "달맞이",
			needle: "ㅁㅈ",
			result: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Search(tt.target, tt.needle)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("Search() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestRangeSearch(t *testing.T) {
	haystack := "간장공장공장장"

	tests := []struct {
		name   string
		needle string
		result []Range
	}{
		{
			name:   "갠",
			needle: "갠",
			result: nil,
		},
		{
			name:   "간",
			needle: "간",
			result: []Range{{0, 0}},
		},
		{
			name:   "장",
			needle: "장",
			result: []Range{{1, 1}, {3, 3}, {5, 5}, {6, 6}},
		},
		{
			name:   "공장",
			needle: "공장",
			result: []Range{{2, 3}, {4, 5}},
		},
		{
			name:   "ㅏㅇㄱ",
			needle: "ㅏㅇㄱ",
			result: []Range{{1, 2}, {3, 4}},
		},
		{
			name:   "ㅇ공ㅈ",
			needle: "ㅇ공ㅈ",
			result: []Range{{1, 3}, {3, 5}},
		},
		{
			name:   "empty",
			needle: "",
			result: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RangeSearch(haystack, tt.needle)
			if !assert.Equal(t, tt.result, result) {
				t.Errorf("RangeSearch() = %v, want %v", result, tt.result)
			}
		})
	}
}
