import org.antlr.v4.runtime.ANTLRFileStream;
import org.antlr.v4.runtime.ANTLRInputStream;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.ParserRuleContext;
import org.antlr.v4.runtime.tree.ParseTree;
import org.antlr.v4.runtime.tree.Tree;

import java.io.IOException;

public class TestSimple {
	public static void main(String[] args) throws IOException {
		ANTLRInputStream input = new ANTLRFileStream(args[0]);
		SimpleLexer lexer = new SimpleLexer(input);
		CommonTokenStream tokens = new CommonTokenStream(lexer);
		SimpleParser parser = new SimpleParser(tokens);
		ParseTree t = parser.s();
		System.out.println(t.toStringTree(parser));
	}
}
