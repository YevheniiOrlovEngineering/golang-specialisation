#!/bin/bash

prog_file="${prog_file:-race.go}"
log_file="${log:-result.log}"
exec_n="${exec_n:-100}"

truncate -s0 "${log_file}"

for (( i=0; i<"${exec_n}"; i++ )); do
  go run "${prog_file}" >> "${log_file}"
done
