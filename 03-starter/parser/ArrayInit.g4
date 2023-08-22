/*  Grammars always start with a grammar header.
 *  This grammar is called 'ArrayInit' and must match the filename 'ArrayInit.g4'.
 */
grammar ArrayInit;

init
   // A rule called 'init' that matches comma-separated values between {...}.

   // Must match at least one value.
   : '{' value (',' value)* '}'
   ;

value
   // A value can be either a nested array/struct or a simple integer (INT).

   // Parser rules start with lowercase letters, lexer rules with uppercase.
   : init
   | INT
   ;

INT
   // Define token INT as one or more digits.
   : [0-9]+
   ;

WS
   // Define whitespace rule, toss it out.
   : [ \t\r\n]+ -> skip
   ;

