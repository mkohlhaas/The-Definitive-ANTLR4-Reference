import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.Trees;

public class TestTags {
    public static void main(String[] args) throws Exception {
        ANTLRInputStream input = new ANTLRInputStream(System.in);
        Tags lexer = new Tags(input);
	Token t = lexer.nextToken();
	while ( t.getType()!=Token.EOF ) {
		System.out.println(t);
		t = lexer.nextToken();
	}
    }
}
