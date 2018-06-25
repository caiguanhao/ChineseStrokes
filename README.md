# ChineseStrokes

Number of strokes for every Chinese character. 81k+ characters included currently. You can use this to sort Chinese characters by number of strokes.

每个中文的笔画数。收录了八万一千多个中文字。可用作按笔画排序。

Source (数据来源) - zidian.911cha.com

## Usage

You can retrieve the strokes data manually:

```
# specify a code range
ChineseStrokes 4e00 4e0a

# specify a file
seq 19968 19978 | awk '{printf "%x\n", $1}' > codes
ChineseStrokes codes
```

Or you can use existing data:

To generate Go code:

```
grep -h -E '^[0-9]+ [0-9]+$' data/*.txt | sort -n | awk '
  BEGIN{print "func getStrokes(charCode int) int {\n\tswitch charCode {"}
    {printf "\tcase %s:\n\t\treturn %s\n", $1, $2}
      END{print "\t}\n\treturn 0\n}"}' > code.go
```

To generate Ruby code:

```
grep -h -E '^[0-9]+ [0-9]+$' data/*.txt | sort -n | awk '
  BEGIN{print "charCodes = {"}
    {printf "  %s => %s,\n", $1, $2}
      END{print "}"}' > code.rb
```

## CharCode

| Language   | Get char code        | To string                    |
|------------|----------------------|------------------------------|
| JavaScript | `'永'.charCodeAt(0)` | `String.fromCharCode(27704)` |
| Ruby       | `'永'.ord`           | `27704.chr(Encoding::UTF_8)` |
| Python     | `ord(u'永')`         | `unichr(27704)`              |
| Go         | `[]rune("永")[0]`    | `string(27704)`              |
