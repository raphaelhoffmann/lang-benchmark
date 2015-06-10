#!/usr/bin/env python3

import fileinput
import io
import sys

set = []

def readDict():
  global set
  with open("../data/english_words.tsv") as f:
    for line in f:
      set.append(line.rstrip())
  set = frozenset(set)

def main():
  readDict()
  bufsize = 1024*1024
  with open('out', 'w', bufsize) as fout:
    with open('../data/sample.txt', 'r', bufsize) as fin:
      for line in fin:
        toks = line[:-1].split(' ') 
        for tok in toks:
          if tok in set:
            print("match", file=fout)

if __name__ == '__main__':
    main()
