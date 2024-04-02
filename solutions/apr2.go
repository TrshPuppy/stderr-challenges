package main

import (
	"fmt"
	"strings"
)

type EncodingMap map[rune]string

const (
	A         = ".-"
	B         = "-..."
	C         = "-.-."
	D         = "-.."
	E         = "."
	AccentedE = "..-.."
	F         = "..-."
	G         = "--."
	H         = "...."
	I         = ".."
	J         = ".---"
	K         = "-.-"
	L         = ".-.."
	M         = "--"
	N         = "-."
	O         = "---"
	P         = ".--."
	Q         = "--.-"
	R         = ".-."
	S         = "..."
	T         = "-"
	U         = "..-"
	V         = "...-"
	W         = ".--"
	X         = "-..-"
	Y         = "-.--"
	Z         = "--.."

	One   = ".----"
	Two   = "..---"
	Three = "...--"
	Four  = "....-"
	Five  = "....."
	Six   = "-...."
	Seven = "--..."
	Eight = "---.."
	Nine  = "----."
	Zero  = "-----"

	Period       = ".-.-.-" //.
	Comma        = "--..--" //,
	Colon        = "---..." //:
	QuestionMark = "..--.." //?
	Apostrophe   = ".----." //'
	Hyphen       = "-....-" //-
	Division     = "-..-."  ///
	LeftBracket  = "-.--."  //(
	RightBracket = "-.--.-" //)
	IvertedComma = ".-..-." //“ ”
	DoubleHyphen = "-...-"  //=
	Cross        = ".-.-."  //+
	CommercialAt = ".--.-." //@

	Understood           = "...-."
	Error                = "........"
	InvitationToTransmit = "-.-"
	Wait                 = ".-..."
	EndOfWork            = "...-.-"
	StartingSignal       = "-.-.-"

	Space = " "
)

var DefaultMorse = EncodingMap{
	'A': A,
	'B': B,
	'C': C,
	'D': D,
	'E': E,
	'F': F,
	'G': G,
	'H': H,
	'I': I,
	'J': J,
	'K': K,
	'L': L,
	'M': M,
	'N': N,
	'O': O,
	'P': P,
	'Q': Q,
	'R': R,
	'S': S,
	'T': T,
	'U': U,
	'V': V,
	'W': W,
	'X': X,
	'Y': Y,
	'Z': Z,

	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",

	'.':  ".-.-.-",
	',':  "--..--",
	':':  "---...",
	'?':  "..--..",
	'\'': ".----.",
	'-':  "-....-",
	'/':  "-..-.",
	'(':  "-.--.",
	')':  "-.--.-",
	'“':  ".-..-.",
	'=':  "-...-",
	'+':  ".-.-.",
	'@':  ".--.-.",

	' ': Space,
}

func main() {
	// Split string into slice breaking on the space b/w words (triple space)
	morseCode := ". -   . - . .   . . - .   . -     . - - .   . -   . - - .   . -     . - .   - - -   - -   .   - - -     . .   - .   - . .   . .   . -     . - . .   . .   - -   . -          . . - .   - - -   - . . -   -   . - .   - - -   -     - - -   . . .   - . - .   . -   . - .     - - -   . . .   - . - .   . -   . - .     . - . .   . .   - -   . -"
	sentences := strings.Split(morseCode, "          ")

	var decoded []string = make([]string, 0)

	// Loop through array of morse characters from message
	for _, s := range sentences {
		words := strings.Split(s, "     ")

		for _, w := range words {
			characters := strings.Split(w, "   ")

			for _, c := range characters {

				trimmed := []string{}
				for _, d := range c {
					//fmt.Printf("d: %s\n", string(d))
					if string(d) != " " {
						trimmed = append(trimmed, string(d))
					}
				}

				tr := strings.Join(trimmed, "")
				// Loop through map of ASCII to morse and compare to current msg char
				for run := range DefaultMorse {
					//				fmt.Printf(tr + "\n")
					if DefaultMorse[run] == tr {
						decoded = append(decoded, string(run))
					}
				}
			}

			decoded = append(decoded, " ")
		}

		decoded = append(decoded, "\n")

	}

	fmt.Printf("%s\n", strings.Join(decoded, ""))

	return
}
