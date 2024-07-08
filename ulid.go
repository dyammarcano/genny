package main

//https://github.com/oklog/ulid

//Usage: ulid [-hlqz] [-f <format>] [parameters ...]
// -f, --format=<format>  when parsing, show times in this format: default, rfc3339, unix, ms
// -h, --help             print this help text
// -l, --local            when parsing, show local time instead of UTC
// -q, --quick            when generating, use non-crypto-grade entropy
// -z, --zero             when generating, fix entropy to all-zeroes
