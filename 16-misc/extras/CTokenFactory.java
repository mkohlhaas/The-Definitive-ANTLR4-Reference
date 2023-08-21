import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.TokenFactory;
import org.antlr.v4.runtime.TokenSource;
import org.antlr.v4.runtime.misc.Pair;
import org.antlr.v4.runtime.misc.Interval;

public class CTokenFactory implements TokenFactory<CToken> {
	private final PreprocessedCharStream cinput;

	public CTokenFactory(PreprocessedCharStream cinput) {
		this.cinput = cinput;
	}

	@Override
	public CToken create(int type, String text) {
		return new CToken(type, text);
	}

	@Override
	public CToken create(Pair<TokenSource, CharStream> source, int type, String text,
						 int channel, int start, int stop, int line,
						 int charPositionInLine)
	{
		CToken t = new CToken(source, type, channel, start, stop);
		t.setLine(line);
		t.setCharPositionInLine(charPositionInLine);
		CharStream input = source.b;
		t.setText(input.getText(Interval.of(start, stop)));
		t.filename = cinput.getFilenameFromCharIndex(start);
		t.setLine(cinput.getLineFromCharIndex(start));
		return t;
	}
}
