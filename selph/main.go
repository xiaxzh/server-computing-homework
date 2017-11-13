package main

import (
  "fmt"
  "flag"
  "os"
  "bufio"
  "io"
)

var (
  flagSet = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
  start_page = flagSet.Int("s", -1, "beginning page of this task")
  end_page = flagSet.Int("e", -1, "endding page of this task")
  page_len = flagSet.Int("l", 72, "lines in one page")
  page_type = flagSet.Bool("f", false, "split by '\\f'")
  print_dest = flagSet.String("d", "", "destination of this task")
)


func isValid() bool {
  // enter not enough arguments
  if flagSet.NFlag() < 2 {
    fmt.Fprintf(os.Stderr, "not enough arguments\n")
    return false
  }

  // invalid start page or forget to set start page
  if *start_page < 1 {
    fmt.Fprintf(os.Stderr, "invalid start page\n")
    return false
  }

  // invalid end page or forget to set end page
  if *end_page < 1 || *end_page < *start_page {
    fmt.Fprintf(os.Stderr, "invalid end page\n")
    return false
  }

  // invalid page_len
  if *page_len < 1 {
    fmt.Fprintf(os.Stderr, "invalid length -f \n")
    return false
  }

  // set -f and -l xxx at the same time
  if *page_type == true && *page_len != 72 {
    fmt.Fprintf(os.Stderr, "you should not enter -f while you have entered -l=xxx\n")
    return false
  }

  return true
}

func printFile(input string) {

  // declaration
  inBuf := bufio.NewReader(os.Stdin)
  if input != "" {
    f, err := os.Open(input)
    // error occurs
    if err != nil {
      fmt.Fprintln(os.Stderr, err.Error())
    }
    inBuf = bufio.NewReader(f)
  }

  // -f that means we count page by '\f'
  if *page_type == true {
    //  remove the unnecessery part of this file
    for i := 1; i < *start_page; i++ {
      inBuf.ReadString('\f')
    }

    // read the content we need in this file
    for i:=0; i < *end_page-*start_page+1; i++ {
      line, err := inBuf.ReadString('\f')
      if err != nil {
        fmt.Fprintf(os.Stderr, err.Error())
      }

      // output the content to the specific file
      if *print_dest != "" {
        output, err := os.OpenFile(*print_dest, os.O_RDWR|os.O_CREATE, 0766)
        if err != nil {
          if err == io.EOF {
            return
          }
          fmt.Fprintln(os.Stderr, err.Error())
        }
        output.WriteString(line)
      } else {
        fmt.Fprintf(os.Stdout, line)
      }
    }

  }

  // -l default 72, count page by lines (lines count is page_len)
  if *page_type ==false {
    // remove the unnecessery part of this file
    for i := 1; i < *start_page; i++ {
      for j := 0; j < *page_len; j++ {
        inBuf.ReadString('\n')
      }
    }

    // read the content we need
    for i := 0; i < *end_page-*start_page+1; i++ {
      for j := 0; j < *page_len; j++ {
        // fmt.Printf("times:%d", j)
        line, err := inBuf.ReadString('\n')
        if err != nil {
          if err == io.EOF {
            return
          }
          fmt.Fprintln(os.Stderr, err.Error())
        }

        if *print_dest != "" {
          // fmt.Printf("Here\n")
          output, err := os.OpenFile(*print_dest, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
          if err != nil {
            fmt.Fprintln(os.Stderr, err.Error())
          }
          output.WriteString(line)
        } else {
          // fmt.Printf("haha\n")
          fmt.Fprintf(os.Stdout, line)
        }
      }
    }
  }

}

func main()  {
  flagSet.Parse(os.Args[1:])

  /*
  fmt.Fprintln(os.Stdout, flagSet.NFlag())
  fmt.Fprintln(os.Stdout, flagSet.Args())
  */

  /*
  fmt.Printf("%d\n", *start_page)
  fmt.Printf("%d\n", *end_page)
  fmt.Printf("%d\n", *page_len)
  fmt.Printf("%s\n", *print_dest)
  */

  if isValid() {

    var input string
    if flagSet.NArg() > 0 {
      input = flagSet.Arg(0)
    }
    printFile(input)
  }
}
