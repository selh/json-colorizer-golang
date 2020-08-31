package tokenizer

import (
  "unicode"
  "strings"
)

/*iota sets the constants to incrementing values starting at 0*/
const (
  OPEN_CURL = iota
  CLOSE_CURL
  OPEN_BRAC
  CLOSE_BRAC
  SEMICOLON
  COMMA
  COLON
  BOOL_NULL
  STRING_CHAR
  NUMERIC
)

var TokenMap = map[string]int {
  "{" : OPEN_CURL,
  "}" : CLOSE_CURL,
  "[" : OPEN_BRAC,
  "]" : CLOSE_BRAC,
  ";" : SEMICOLON,
  "," : COMMA,
  ":" : COLON,
}

type Token struct{
  TokenType int
  Lexeme string
}

/* Check for valid JSON numeric structure
    including 1.0001, 10e-10, 10e10, 10e+10, 10E-10, etc... */
func isValidDigit( s byte ) bool{
  if (unicode.IsDigit(rune(s)) || s == '.' || s == 'e' || s == 'E') {
    return true
  } else {
    return false
  }
}

/* Label special JSON keywords to proper TokenType
(eg. true, false, null ) */
func keywordMapper( slice string, index *int,
                     s     string, key_len int ) (Token, bool) {
    equals := strings.Compare(s, slice)
    if equals == 0 {
      *index+=key_len-1
      return Token{BOOL_NULL, s}, true
    }
    return Token{}, false
}

/*Reads a string input and seperates each token into a
Token struct that contains its token type and token value
Returns an array of Token structures */
func Tokenize(s string) (result []Token) {

  n := len(s)

  for i:=0; i < n; i++ {
    switch {
    case s[i] == ' ', s[i] == '\n', s[i] == '\t':
        //do nothing
    case s[i] == '"': //If string read until closing "
      i++
      begin := i
      for i < n && (s[i] != '"' || s[i-1] == '\\') {
        i++
      }
      result = append(result, Token{STRING_CHAR, s[begin:i]})
    case s[i] == '-', unicode.IsDigit(rune(s[i])):
      begin := i
      i++
      if unicode.IsDigit(rune(s[i+1])){
        for i < n && isValidDigit(s[i]){
          i++
        }
      }
      result = append(result, Token{NUMERIC, s[begin:i]})
      i--
    case s[i] == 't': //check for keyword: true
      key_len := 4
      begin := i
      if i+key_len < n {
        tok, ok := keywordMapper(s[begin:i+key_len], &i, "true", key_len)
        if ok{
          result = append(result, tok)
        }
      }
    case s[i] == 'f': //check for keyword: false
      key_len := 5
      begin := i
      if i+key_len < n {
        tok, ok := keywordMapper(s[begin:i+key_len], &i, "false", key_len)
        if ok{
          result = append(result, tok)
        }
      }
    case s[i] == 'n': //check for keyword: null
      key_len := 4
      begin := i
      if i+key_len < n {
        tok, ok := keywordMapper(s[begin:i+key_len], &i, "null", key_len)
        if ok{
          result = append(result, tok)
        }
      }
    default:
      tok := string(s[i])
      result = append(result, Token{TokenMap[tok], tok})
    }
  }
  return result
}
