#!/bin/bash
find -name "*.sh" -printf "%f\n" | sed -r 's/^(.+)\.[^.]+$/\1/'| sort -r
