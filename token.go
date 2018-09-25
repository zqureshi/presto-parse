package presto

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT // fields, table_name

	// Misc characters
	ASTERISK // *
	COMMA    // ,
	DOT      // .

	// Keywords
	SELECT
	FROM
)
