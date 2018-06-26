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

To generate Ruby hash:

```
grep -h -E '^[0-9]+ [0-9]+$' data/*.txt | sort -n | awk '
  BEGIN{print "charCodes = {"}
    {printf "  %s => %s,\n", $1, $2}
      END{print "}"}' > code.rb
```

To generate JSON:

```
grep -h -E '^[0-9]+ [0-9]+$' data/*.txt | sort -n | awk '
  BEGIN{print "{"}
    NR>1{ print "," }
      {printf "  \"%s\": %s", $1, $2}
        END{print "\n}"}' > code.json
```

## CharCode

| Language   | Get char code        | To string                    |
|------------|----------------------|------------------------------|
| JavaScript | `'永'.charCodeAt(0)` | `String.fromCharCode(27704)` |
| Ruby       | `'永'.ord`           | `27704.chr(Encoding::UTF_8)` |
| Python     | `ord(u'永')`         | `unichr(27704)`              |
| Go         | `[]rune("永")[0]`    | `string(27704)`              |

## Statistics

Average:

```
grep -h -E '^[0-9]+ [0-9]+$' data/*.txt | awk '{total += $2} END{print total/NR}'
=> 13.961
```

Distribution:

```
grep -h -E '^[0-9]+ [0-9]+$' data/*.txt | sort -k 2 -n | uniq -f 1 -c
```

count  | char code | storkes
------ | --------- | -------
20     | 131273    | 1
77     | 131073    | 2
173    | 131075    | 3
449    | 131072    | 4
808    | 131086    | 5
1620   | 131096    | 6
2703   | 131105    | 7
3776   | 131116    | 8
4751   | 131125    | 9
5625   | 131133    | 10
6445   | 131137    | 11
7037   | 131142    | 12
6765   | 131149    | 13
6650   | 131151    | 14
6377   | 131155    | 15
5816   | 131158    | 16
4715   | 131164    | 17
3988   | 131167    | 18
3282   | 131194    | 19
2669   | 131168    | 20
2080   | 131486    | 21
1589   | 131195    | 22
1185   | 131272    | 23
882    | 131487    | 24
564    | 132209    | 25
392    | 132210    | 26
308    | 133483    | 27
220    | 131488    | 28
127    | 131489    | 29
91     | 136579    | 30
50     | 135583    | 31
41     | 138176    | 32
30     | 133644    | 33
12     | 133323    | 34
11     | 136472    | 35
21     | 136473    | 36
5      | 145843    | 37
4      | 161759    | 38
6      | 136474    | 39
2      | 168403    | 40
1      | 173258    | 41
1      | 160152    | 42
2      | 161969    | 44
1      | 169571    | 45
1      | 169572    | 47
2      | 158149    | 48
1      | 158148    | 49
1      | 40856     | 51
1      | 19003     | 52
1      | 181929    | 53
2      | 132411    | 64
