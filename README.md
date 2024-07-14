# wc clone

A clone of the `wc` Linux CLI tool.

## Usage

```bash
wc-clone filename.txt

# output
# file: one-thousand-rows.txt
# lines: 1000

wc-clone -m filename.txt
# output
# file: one-thousand-rows.txt
# lines: 1000
# chars: 71000

wc-clone -mc filename.txt
# output
# file: one-thousand-rows.txt
# lines: 1000
# chars: 71000
# bytes: 72000
```

## Available options

- `-l` - Counts all the lines
- `-m` - Counts all the characters
- `-c` - Counts bytes
- `-w` - Counts all the words
- `-L` - Gets the length of the largest line
