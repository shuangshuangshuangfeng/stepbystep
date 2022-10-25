import java.util.List;
import java.util.stream.Collectors;
import java.util.ArrayList;

/**
 * java8的例子，使用lambda表达式
 * @author zhangchao
 *
 */
public class Main {

    /**
     * 准备书的列表数据
     * @return
     */
    public static List<Book> prepareData() {
        // 准备书的列表，id是从1到10
        List<Book> bookList = new ArrayList<Book>();
        for (int i = 1; i < 11; i++) {
            bookList.add(new Book(String.valueOf(i), "book"+i));
        }
        return bookList;
    }

    public static void main(String[] args) {
        List<Book> bookList = prepareData();

        // 要被找出的书的ID
        // 这个输入法是真的很好用
        ArrayList<String> ids = new ArrayList<String>();
        ids.add("3");
        ids.add("6");
        ids.add("8");
        ids.add("9");

        // 存放过滤结果的列表
        List<Book> result = null;

        // 使用lambda表达式过滤出结果并放到result列表里，written by zhangchao
//        result = bookList.stream()
//                .filter(((Book b) -> b.getId()) > 5)
//                .collect(Collectors.toList());

        // 打印结果列表
        if (result != null && !result.isEmpty()) {
            result.forEach((Book b) -> System.out.println(b.getId() + "  " + b.getName()));
        }
    }
}

class Book {
    private String id;
    private String name;

    public Book(String id, String name) {
        super();
        this.id = id;
        this.name = name;
    }

    public Book() {
        super();
    }

    public String getId() {
        return id;
    }
    public void setId(String id) {
        this.id = id;
    }
    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }

}
