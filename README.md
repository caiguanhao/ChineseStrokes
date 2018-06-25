# ChineseStrokes

Number of strokes for every Chinese character. 81k+ characters included currently. You can use this to sort Chinese characters by number of strokes.

每个中文的笔画数。收录了八万一千多个中文字。可用作按笔画排序。

Source (数据来源) - zidian.911cha.com

Generate Go code:

```
grep -h -E '^[0-9]+ [0-9]+$' data/*.txt | sort -n | awk '
  BEGIN{print "func getStrokes(charCode int) int {\n\tswitch charCode {"}
    {printf "\tcase %s:\n\t\treturn %s\n", $1, $2}
      END{print "\t}\n\treturn 0\n}"}' > code.go
```

Generate Ruby code:

```
grep -h -E '^[0-9]+ [0-9]+$' data/*.txt | sort -n | awk '
  BEGIN{print "charCodes = {"}
    {printf "  %s => %s,\n", $1, $2}
      END{print "}"}' > code.rb
```
