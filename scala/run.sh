#!/bin/bash

scalac bench2.scala
scalac bench2withpipes.scala

time $(scala bench2)
time $(cat ../data/sample.txt | scala bench2withpipes > out)

