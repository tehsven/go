import com.sun.jna.Library;
import com.sun.jna.Native;

public class sinwave {
	public interface sinwave_go extends Library {
		sinwave_go instance = (sinwave_go) Native.loadLibrary("_cgo_", sinwave_go.class);
		void sinwave(double freq, double duration);
	}

	public static void main(String[] args) {
		System.out.println("playing sinwave implemented in go");
		sinwave_go.instance.sinwave(440.0, 4.0);
	}
}