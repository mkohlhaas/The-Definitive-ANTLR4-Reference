import org.antlr.v4.runtime.TokenSource;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.misc.Pair;
import org.antlr.v4.runtime.misc.Interval;

import java.util.List;

public class CPPToken extends CToken {
	public List<Interval> lineIntervals;

	public CPPToken(int type, String text) {
		super(type, text);
	}

	public CPPToken(Pair<TokenSource, CharStream> source, int type, int channel, int start, int stop) {
		super(source, type, channel, start, stop);
	}
}
