import org.antlr.v4.runtime.*;

public class TestR {
	public static void main(String[] args) throws Exception {
		ANTLRInputStream input = new ANTLRFileStream(args[0]);
		RLexer lexer = new RLexer(input);
		CommonTokenStream tokens = new CommonTokenStream(lexer);
		RParser parser = new RParser(tokens);
		parser.setBuildParseTree(true);
		RuleContext tree = parser.prog();
		tree.inspect(parser); // show in gui
		//tree.save(parser, "/tmp/R.ps"); // Generate postscript
		System.out.println(tree.toStringTree(parser));
	}
}
