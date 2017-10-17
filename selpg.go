/*=================================================================
Program name:
	selpg (SELect PaGes)

Purpose:
	Sometimes one needs to extract only a specified range of
pages from an input text file. This program allows the user to do
that.
===================================================================*/
package main
/*================================= import  ======================*/
import (
    "fmt"
    "os"
    "io"
    "flag"
    "bufio"
)
/*================================= types =========================*/

type sp_args  struct{
	start_page int
	end_page int
    in_filename string
    out_filename string
    print_dest string
	page_len int
	/* default value, can be overriden by "-l number" on command line */
	page_type bool
	/* 'l' for lines-delimited, 'f' for form-feed-delimited */
	/* default is 'l=72' */
}

/*================================= main()=== =====================*/
var progname string
func main() {
	/* save name by which program is invoked, for error messages */
	progname = os.Args[0]
	var sa sp_args
	initalArgs(&sa)
	process_args(&sa)
	process_input(&sa)
}
/*==========================initalArgs(args *sp_args)==============*/
func initalArgs(args *sp_args) {
	flag.IntVar(&args.start_page, "s", -1, "start page")
    flag.IntVar(&args.end_page, "e", -1, "end page")
    flag.IntVar(&args.page_len, "l", -1, "page len")
    flag.BoolVar(&args.page_type, "f", false, "type of print")
    flag.StringVar(&args.print_dest, "d", "", "specify the printer")
    //调用flag.Parse() 来执行命令行解析。
    flag.Parse()
	flag.Usage = Usage
}
/*================================= usage() =======================*/

func Usage() {
  	fmt.Printf("\nUSAGE: %s -s=start_page -e=end_page [ -f | -l=lines_per_page ] [ -ddest ] [ filename ]\n", progname);
    fmt.Printf("The arguments are:\n\n")
    fmt.Printf("\t-s=Number\tStart from Page <number>.\n")
    fmt.Printf("\t-e=Number\tEnd to Page <number>.\n")
    fmt.Printf("\t-l=Number\t[options]Specify the number of line per page.Default is 72.\n")
    fmt.Printf("\t-f\t\t[options]Specify that the pages are sperated by \\f.\n")
    fmt.Printf("\t-d=lpnumber\t[options]Print this page at lpnumber printer.\n")
    fmt.Printf("\t[filename]\t[options]Read input from the file and write output to the file.\n\n")
    fmt.Printf("If no file specified, %s will read input from stdin.\n\n", progname)
}
/*================================= process_args() ================*/

func process_args(sa *sp_args) {
	if sa.start_page <= 0 || sa.end_page <= 0 || sa.start_page > sa.end_page {
        fmt.Printf("%s: args are error\nPlease ensure the form of args and the start Page is larger than the end,\nand both of them are bigger than 1\n", progname)
        flag.Usage()
        os.Exit(1)
    }
    if sa.page_type == false {
        if sa.page_len == -1 {
            sa.page_len = 5//for test
        }
    }
    if sa.page_type == true && sa.page_len != -1 {
        fmt.Printf("%s: -l and -f can't both exist.\n\n", progname)
        flag.Usage()
        os.Exit(1)
    }
}

/*================================= process_input() ===============*/
func process_input(sa *sp_args) {
    //flag提供了Arg(i),Args()来获取non-flag参数，NArg()来获取non-flag的个数
    if flag.NArg() > 0 {
		sa.in_filename = flag.Arg(0)
	}
	if flag.NArg() > 1 {
		sa.out_filename = flag.Arg(1)
	}
	write(sa)
}
func write(sa *sp_args) {
	In := sa.in_filename
	Out := sa.out_filename
    /*
    O_RDWR：读写模式
    O_CREATE：文件不存在就创建
    O_TRUNC：打开并清空文件
    */
    var err error
	var Ibuf *bufio.Reader
	if In != "" {
		inFile, err := os.OpenFile(In, os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "warning(open file)1 ", err.Error())
			return
		}
		Ibuf = bufio.NewReader(inFile)
	} else {
		Ibuf = bufio.NewReader(os.Stdin)
	}

	var Obuf *os.File
	if Out != "" {
		Obuf, err = os.OpenFile(Out, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "warning(open file) 2", err.Error())
			return
		}
	} else {
		Obuf = nil
	}

	var count int
	count = sa.end_page - sa.start_page + 1
	if !sa.page_type {
	    /*read all the char from the file begin to the the startpage*/
		for i := 1; i < sa.start_page; i++ {
			for j := 0; j < sa.page_len; j++ {
				Ibuf.ReadString('\n')
			}
		}
		for i := 0; i < count; i++ {
			for j := 0; j < sa.page_len; j++ {
				line, err := Ibuf.ReadString('\n')
				if err != nil {
					if !(err == io.EOF && i != count && j != sa.page_len) {
						fmt.Fprint(os.Stderr, "warning(file reading)3 ", err.Error())
					}
				}
				if Obuf != nil {
					Obuf.WriteString(line)
				} else {
					fmt.Print(line)
				}
				if len(line) == 0 {
				    fmt.Printf("the start page is larger than the total page of this file,so there are nothing\n")
			        return
			    }
			}
		}
	} else { /*the cut of the page*/
		for i := 1; i < sa.start_page; i++ {
			Ibuf.ReadString('\f')
		}
		for i := 0; i < count; i++ {
			line, err := Ibuf.ReadString('\f')
			if err != nil {
				if !(err == io.EOF && i != count) {
					fmt.Fprint(os.Stderr, "warning(file reading)4 ", err.Error())
				}
			}
			if Obuf != nil {
				Obuf.WriteString(line)
			} else {
				fmt.Print(line)
			}
			if len(line) == 0 {
				fmt.Printf("the start page is larger than the total page of this file,so there are nothing\n")
			    return
			}
		}
	}
}