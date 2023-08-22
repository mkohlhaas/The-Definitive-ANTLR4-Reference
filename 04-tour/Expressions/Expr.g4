grammar Expr;

prog
   // The start rule; begin parsing here.
   : statement+
   ;

statement
   : expr NEWLINE
   | ID '=' expr NEWLINE
   | NEWLINE
   ;

expr
   : expr ('*' | '/') expr
   | expr ('+' | '-') expr
   | INT
   | ID
   | '(' expr ')'
   ;

ID
   // match identifiers
   : [a-zA-Z]+
   ;

INT
   // match integers
   : [0-9]+
   ;

NEWLINE
   // return newlines to parser (is end-statement signal)
   : '\r'? '\n'
   ;

WS
   // toss out whitespace
   : [ \t]+ -> skip
   ;

