import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;

import java.util.ArrayDeque;

public class CPPLexer extends CPPBaseLexer {
	protected ArrayDeque<Token> buffer = new ArrayDeque<Token>();

	public CPPLexer(CharStream input) { super(input); }

	@Override
	public Token nextToken() {
		if ( buffer.size()>0 ) {
			return buffer.removeFirst();
		}
		else {
			// matched rule adds at least one to buffer via emit(t)
			super.nextToken(); // ignore return value; we use buffer
			return buffer.removeFirst();
		}
	}

	@Override
	public Token getToken() {
		return buffer.peek();
	}

	@Override
	public void emit(Token token) {
		super.emit(token);
		buffer.add(token);
	}
}
