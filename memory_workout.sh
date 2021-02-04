#!/bin/bash

MY_DIR=$(dirname $0)

# Run through some memory exercises
for i in {1..100};
do
  $MY_DIR/derrick_tools memoryquiz presidents
done

for i in {1..100};
do
  $MY_DIR/derrick_tools memoryquiz countries
done

for i in {1..5};
do
  $MY_DIR/derrick_tools memoryquiz numbers
done
