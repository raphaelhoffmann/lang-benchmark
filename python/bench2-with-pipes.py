#!/usr/bin/env python3

import fileinput
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
  for line in sys.stdin:
    toks = line[:-1].split(' ') 
    for tok in toks:
      if tok in set:
        print ("match")

if __name__ == '__main__':
    main()
