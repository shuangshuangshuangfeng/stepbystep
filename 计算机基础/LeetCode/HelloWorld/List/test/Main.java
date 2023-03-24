import java.util.Date;
import java.util.List;
import java.util.stream.Collectors;
import java.util.ArrayList;

/**
 * java8的例子，使用lambda表达式
 * @author zhangchao
 *
 */
public class Main {

    public static void main(String[] args) throws Exception{
        System.out.println(caculateTimeout(100, 60));



    }

    private static int caculateTimeout(int limit, int spend) {
        double percent = (spend * 1.0) / (limit * 1.0);

        if (percent >= 2.0 && percent <2.1) {
            return 1;
        } else if (percent >= 1.5 && percent < 1.6) {
            return 2;
        } else if (percent >= 1.0 && percent < 1.1) {
            return 3;
        } else if (percent >= 0.8 && percent < 0.9) {
            return 4;
        } else if (percent >= 0.5 && percent < 0.6) {
            return 5;
        } else { // 不需要提醒
            return -1;
        }
    }

}
