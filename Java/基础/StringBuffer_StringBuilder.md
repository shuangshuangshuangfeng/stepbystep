### ``StringBuffer``和``StringBuilder``

### ``StringBuffer``
``StringBuffer``和``String``用法差不多，只是它**是线程安全**的，它的每个方法中，必要的都加上了``synchronize``。

它的主要方法有:
- insert(), 向字符串的某个``index``下插入字符。
- append()，字符串结尾插入字符串。

每个``StringBuffer``都有一个``capacity``(容量), 只要字符串的长度没有超过这个容量，就没有必要分配新的缓冲区数组，如果缓冲区内部溢出了，它会自动变大。

从JDK5版本开始，就有了非线程安全的``StringBuilder``类，``StringBuilder``和``StringBuffer``等效，``StringBuilder``类支持所有的``StringBuffer``的操作，但是比``StringBuffer``更快，因为它不是线程安全的，即单线程使用的。








