package main

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

func readDict() (set map[string]bool) {
  set = make(map[string]bool)
  file, err := os.Open("../data/english_words.tsv")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    set[scanner.Text()] = true
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return
}

func main() {
    filename := "../data/sample.txt"

    set := readDict()

    fo, _ := os.OpenFile("./out", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
    w := bufio.NewWriterSize(fo, 1024*1024)

    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    r := bufio.NewReaderSize(f, 1024*1024)
    line, isPrefix, err := r.ReadLine()
    for err == nil && !isPrefix {
        s := string(line)

        toks := strings.Split(s, " ")

        for _, tok := range toks {
          if set[tok] {
            w.WriteString("match\n")
          }
        }

        line, isPrefix, err = r.ReadLine()
    }
    if isPrefix {
        fmt.Println(errors.New("buffer size to small"))
        return
    }
    if err != io.EOF {
        fmt.Println(err)
        return
    }
}

