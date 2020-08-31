//JSON Colorizer + Format Package
package colorizer

import(
  tokens "format/tokenize"
  "fmt"
  "strconv"
  //"unicode
)

//How many percent to increment different levels of indents
const IndentFactor = 1

/*Use special html format for these characters*/
var escSpecialHtml = map[rune]string {
  '<': "&lt;",
  '>': "&gt;",
  '&': "&amp;",
  '"': "&quot;",
  '\'': "&apos;",
}

/*Wrap each new indent level with a div to control json indentation
 openDiv contains the opening declaration for the div while
 closeDiv contains the closing declaration
 These two functions must be used together for proper html formating*/
func openDiv( indent_lvl int ) string {
  wrapper := "<div style=\"margin:0 auto;\"><div style=\"margin-left:" +
              strconv.Itoa(indent_lvl) + "%;\">"
  return wrapper
}

func closeDiv() string {
  return "</div></div>"
}

/*If you want to color special characters like
 {}, [], , , :, ; [0-9], [a-z]
different colors inside string use this*/
// func colorString( str string ) {
//   //fmt.Print("<span style=\"color:royalblue\">&quot;" + v.Lexeme + "&quot;</span>")
//   fmt.Print("<span style=\"color:royalblue\">&quot;</span>")
//   for _, value := range str {
//       switch{
//       case value == '{' || value == '}':
//         fmt.Print("<span style=\"color:midnightblue\">"+ string(value) +"</span>")
//       case value == '[' || value == ']':
//         fmt.Print("<span style=\"color:darkmagenta\">"+ string(value) +"</span>")
//       case value == ',':
//         fmt.Print("<span style=\"color:orchid\">"+ string(value) +"</span>")
//       case value == ':':
//         fmt.Print("<span style=\"color:orangered\">"+ string(value) +"</span>")
//       case value == ';':
//         fmt.Print("<span style=\"color:steelblue\">"+ string(value) +"</span>")
//       case unicode.IsDigit(value):
//         fmt.Print("<span style=\"color:mediumvioletred\">"+ string(value) +"</span>")
//       case unicode.IsLetter(value):
//         fmt.Print("<span style=\"color:royalblue\">"+ string(value) +"</span>")
//       default:
//         escaped_val, exists := escSpecialHtml[value]
//         if exists {
//           fmt.Print("<span style=\"color:green\">"+ string(escaped_val) +"</span>")
//         } else {
//           fmt.Print("<span style=\"color:green\">"+ string(value) +"</span>")
//         }
//       }
//   }
//   fmt.Print("<span style=\"color:royalblue\">&quot;</span>")
// }

/*Takes input pointer to slice of Token objects and prints them in
reader-friendly format. Also colors tokens according to Colors map
commented above */
func ColorizeTokens( t *[]tokens.Token ) {

  fmt.Print("<span style=\"font-family:monospace; white-space:pre\">")
  array_flag := false //used to track if comma should have \n after it
  close_flag := false //for the case where we have [ {...} ]
  indent_lvl := 0 //track the current indentation level

  for i, v := range *t {
    switch v.TokenType {

    case tokens.OPEN_CURL:

      array_flag = false
      indent_lvl += IndentFactor
      html := openDiv(indent_lvl)
      if ( i > 0 && (*t)[i-1].TokenType != tokens.OPEN_BRAC ){
        html += "<span style=\"color:midnightblue\">{</span><br/>"
      } else {
        html += "<br/><span style=\"color:midnightblue\">{</span><br/>"
      }
      fmt.Print(html)
      fmt.Print(openDiv(indent_lvl))

    case tokens.CLOSE_CURL:

      fmt.Print(closeDiv())
      html := "<br/><span style=\"color:midnightblue\">}</span><br/></div></div>"
      //if next token is a comma keep it on the same line as the closing brace
      if (i+1 < len(*t) && (*t)[i+1].TokenType == tokens.COMMA) {
        html = "<br/><span style=\"color:midnightblue\">}</span>"
        close_flag = true
      }
      fmt.Print(html)
      indent_lvl -= IndentFactor

    case tokens.OPEN_BRAC:
      array_flag = true
      fmt.Print("<span style=\"color:darkmagenta\">[</span>")
    case tokens.CLOSE_BRAC:
      array_flag = false
      fmt.Print("<span style=\"color:darkmagenta\">]</span>")
    case tokens.SEMICOLON:
      fmt.Print("<span style=\"color:steelblue\">;</span><br/>")
    case tokens.COMMA:
      if array_flag {
        fmt.Print("<span style=\"color:orchid\">, </span>")
      } else if close_flag {
        fmt.Print("<span style=\"color:orchid\">, </span></div></div>")
        close_flag = false
      } else {
        fmt.Print("<span style=\"color:orchid\">,</span><br/>")
      }
    case tokens.COLON:
      fmt.Print("<span style=\"color:orangered\"> : </span>")
    case tokens.BOOL_NULL:
      fmt.Print("<span style=\"color:gold\">" + v.Lexeme + "</span>")
    case tokens.STRING_CHAR:
      //colorString(v.Lexeme)
      fmt.Print("<span style=\"color:royalblue\">&quot;" + v.Lexeme + "&quot;</span>")
    case tokens.NUMERIC:
      fmt.Print("<span style=\"color:mediumvioletred\">" + v.Lexeme + "</span>")
    }
  }
  fmt.Print("</span>")
}
