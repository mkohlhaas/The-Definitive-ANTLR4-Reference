import org.antlr.v4.runtime.CommonToken;
import org.antlr.v4.runtime.TokenSource;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.misc.Pair;

public class CToken extends CommonToken {
	public String filename;

	public CToken(int type, String text) {
		super(type, text);
	}

	public CToken(Pair<TokenSource, CharStream> source,
                  int type, int channel, int start, int stop)
    {
		super(source, type, channel, start, stop);
	}

	@Override
	public String toString() {
		String t = super.toString();
		return filename+":"+t;
	}
}
