import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.Trees;

public class TestComment {
    public static void main(String[] args) throws Exception {
        ANTLRInputStream input = new ANTLRInputStream(System.in);
        CommentLexer lexer = new CommentLexer(input);
        CommonTokenStream tokens = new CommonTokenStream(lexer);
        CommentParser parser = new CommentParser(tokens);
        parser.setBuildParseTree(true);
        parser.file();
    }
}
