package colorizer

import tokens "format/tokenize"

var Color = map[int]string {
  tokens.OPEN_CURL: "midnightblue",
  tokens.CLOSE_CURL: "midnightblue",

  tokens.OPEN_BRAC: "darkmagenta",
  tokens.CLOSE_BRAC: "darkmagenta",

  tokens.SEMICOLON: "steelblue",
  tokens.COMMA: "orchid",
  tokens.COLON: "orangered",
  tokens.BOOL_NULL: "gold",
  tokens.STRING_CHAR: "royalblue",
  tokens.NUMERIC: "mediumvioletred",
}
