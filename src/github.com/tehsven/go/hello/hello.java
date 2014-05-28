import com.sun.jna.Native;

public class hello {

	public static class hello_c {
		static {
			Native.register("hello");
		}
		public native void hello();
	}

	public static void main(String[] args) {
		System.out.println("hello java");
		new hello_c().hello();
	}
}