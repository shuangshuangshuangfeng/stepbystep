public class JNIDemo {
    static {
        System.loadLibrary("JNIDemo");
    }

    public native String helloWorld();

    public static void main(String[] args){
        JNIDemo demo = new JNIDemo();
        System.out.println(demo.helloWorld());
    }
}
