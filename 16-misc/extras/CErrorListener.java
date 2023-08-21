import org.antlr.v4.runtime.BaseErrorListener;
import org.antlr.v4.runtime.RecognitionException;
import org.antlr.v4.runtime.Recognizer;

/** Same same, just adds filename */
public class CErrorListener extends BaseErrorListener {
	@Override
	public void syntaxError(Recognizer<?, ?> recognizer,
							Object offendingSymbol,
							int line, int charPositionInLine,
							String msg, RecognitionException e)
	{
		CPPToken token = (CPPToken)offendingSymbol;
		System.err.println(token.filename +
						   " line " + line + ":" + charPositionInLine + " at " +
						   token.getText() + ": " + msg);
	}
}
