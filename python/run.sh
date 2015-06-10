#!/bin/bash

time $(./bench2.py)

time $(cat ../data/sample.txt | ./bench2-with-pipes.py > out)
