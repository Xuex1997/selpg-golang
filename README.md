# Selpg--golang
Sometimes one needs to extract only a specified range of pages from an input text file. This program allows the user to do that.

## Usage

```shell
Usage of selpg:
  -s int
  		start page(default -1)
  -e int
  		end page(default -1)
  -f bool
  		use \f to seperate pages
  -l int
  		length of page to seperate pages(!!page_type and page_len can\'t both exist)
  filename
    	[Options]Read input from the file and write output to the file.
```
## Examples

**Assume there are two files, `t1` and `t2` and assume the default page len is 5(for test)**
```shell
$ cat t1.text
line 1 of file one
line 2 of file one
line 3 of file one
line 4 of file one
line 5 of file one
line 6 of file one
line 7 of file one
line 8 of file one
line 9 of file one
line 10 of file one
line 11 of file one
line 12 of file one
$ cat t2.text
line 1 of file two
line 2 of file two
line 3 of file two
line 4 of file two
line 5 of file two
line 6 of file two
line 7 of file two
line 8 of file two
line 9 of file two
line 10 of file two
line 11 of file two
line 12 of file two
```
**Invalid Input**

```shell
$ ./selpg -s 0 -e 2  < file1
./selpg: args are error
Please ensure the form of args and the start Page is larger than the end,
and both of them are bigger than 1
```
```shell
$ ./selpg -s 2 -e 1  < file1
./selpg: args are error
Please ensure the form of args and the start Page is larger than the end,
and both of them are bigger than 1
```
```shell
./selpg -s 1 -e 1 t6.text
warning(open file)1  open t6.text: no such file or directory
```
```shell
$ ./selpg -s 1 -e 2 -l 3 -f t1.text
./selpg: -l and -f can\'t both exist.
```
```shell
$ ./selpg -s 6 -e 7 -l 3  t1.text
the start page is larger than the total page of this file,so there are nothing
```

**Read from standard input**
```shell
$ ./selpg -s 1 -e 2  < file1
line 1 of file one
line 2 of file one
line 3 of file one
line 4 of file one
line 5 of file one
line 6 of file one
line 7 of file one
line 8 of file one
line 9 of file one
line 10 of file one
```

**Read from file1**
```shell
$./selpg -s 1 -e 2 -l 3 t1.text
line 1 of file one
line 2 of file one
line 3 of file one
line 4 of file one
line 5 of file one
line 6 of file one
```

**Read from file1 and write to file2**
```shell
./selpg -s 1 -e 2 -l 3  t1.text t2.text
$cat t2.text
line 1 of file one
line 2 of file one
line 3 of file one
line 4 of file one
line 5 of file one
line 6 of file one
```
```shell
cat t1.text | ./selpg -s 1 -e 1
line 1 of file one
line 2 of file one
line 3 of file one
line 4 of file one
line 5 of file one
```
`暂时没有实现打印机的功能,希望以后有时间可以实现`


