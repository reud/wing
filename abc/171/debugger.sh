#!/bin/bash

CONTEST_NAME="abc170"
QUESTION_ID=a

g++ ${QUESTION_ID}.cpp -o  ${QUESTION_ID}.out

oj d https://atcoder.jp/contests/${CONTEST_NAME}/tasks/${CONTEST_NAME}_${QUESTION_ID} -d ${QUESTION_ID}
oj t -d ${QUESTION_ID} -c ./${QUESTION_ID}.out

