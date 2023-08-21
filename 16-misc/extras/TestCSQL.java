import org.antlr.v4.runtime.*;

public class TestCSQL {
	public static void main(String[] args) throws Exception {
		CharStream in = null;
		if ( args.length>0 ) in = new ANTLRFileStream(args[0]);
		else in = new ANTLRInputStream(System.in);

		CharsAsTokens fauxLexer = new CharsAsTokens(in, CSQL.tokenNames);
		CommonTokenStream tokens = new CommonTokenStream(fauxLexer);
		CSQL parser = new CSQL(tokens);
		ParserRuleContext tree = parser.prog();
		System.out.println(tree.toStringTree(parser));
		tree.inspect(parser);
	}
}
