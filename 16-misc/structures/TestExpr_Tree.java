import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.Trees;

public class TestExpr_Tree {
    public static void main(String[] args) throws Exception {
        ANTLRInputStream input = new ANTLRInputStream(System.in);
        ExprLexer lexer = new ExprLexer(input);
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        ExprParser parser = new ExprParser(tokens);
        parser.setBuildParseTree(true);
        ParserRuleContext tree = parser.expr();
	System.out.println(tree.toStringTree(parser));
        tree.save(parser, "/tmp/t.ps"); // Generate postscript
    }
}
