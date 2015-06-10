package main

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "log"
    "os"
    "runtime"
    "strings"
    "sync"
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
    //fmt.Println(scanner.Text())
    set[scanner.Text()] = true
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return
}


func main() {
    //set := readDict() 
    another()
    //oldReadLine()
    //ReadLine("../data/sample.txt")
    //Readln(os.Stdin)
}

func merge(cs ...<-chan []string) <-chan []string {
    var wg sync.WaitGroup
    out := make(chan []string)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan []string) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func read(filename string) <-chan []string {
    out := make(chan []string)
    go func() {
	f, err := os.Open(filename)
    	if err != nil {
        	fmt.Println(err)
        	return
    	}
    	defer f.Close()
    	r := bufio.NewReaderSize(f, 1024*1024)
        
    	line, isPrefix, err := r.ReadLine()
        batch := make([]string, 50, 50)
    	for err == nil && !isPrefix {
                s := string(line)
                if (len(batch) < 50) {
		  batch = append(batch, s)
                } else {
		  out <- batch
                  batch = make([]string, 50,50)
                }
       
                line, isPrefix, err = r.ReadLine()
        }
        if (len(batch) > 0) {
            out <- batch
        }
        
    	if isPrefix {
        	fmt.Println(errors.New("buffer size to small"))
        	return
    	}
	if err != io.EOF {
        	fmt.Println(err)
        	return
	}
        close(out)
    }()
    return out
}

func write(filename string, in<- chan []string) {
    fo, _ := os.OpenFile("./out", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
    w := bufio.NewWriterSize(fo, 1024*1024)
    w.WriteString("hell world")   
 
    for batch := range in {

       for _, line := range batch {
         w.WriteString(line)
         w.WriteString("\n")
       }
    }
    // cleanup
    w.Flush()
    return
}

func worker(in <-chan []string, out chan<- []string) {
        set := readDict()
        
        for batch := range in {
            for _, line := range batch {
              //fmt.Print(line)
              toks := strings.Split(line, " ")

              for _, tok := range toks {
                if set[tok] {
                  out <- []string{"match"}
                }
              }
           }
        }
        close(out)
        return
}

func run(par int, in <-chan []string, w func(<-chan []string, chan<- []string)) <-chan []string {
    runtime.GOMAXPROCS(par)
    workers := make([]<-chan []string, par)
    for i := range workers {
      outi := make(chan []string)
      workers[i] = outi
      go w(in, outi)
    }
    out := merge(workers...)
    return out
}



func another() {
    in := read("../data/sample.txt")
    out := run(8, in, worker) 
    write("out", out)


}

