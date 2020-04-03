# gocolumns

Just a little program I wrote while learning.

It will take a CSV and convert the column headers to a formatted struct. Might save you some time/headaches if you have a ton of columns.


Example CSV

```
long_col_with_underscores, short_underscore , short-dashes,plain
1,2,3,4
1,2,3,4
1,2,3,4
1,2,3,4
```

Example Output

```
type TestThisStuff struct {
	LongColWithUnderscores string `csv:"LongColWithUnderscores"`
	ShortUnderscore        string `csv:"ShortUnderscore"`
	ShortDashes            string `csv:"ShortDashes"`
	Plain                  string `csv:"Plain"`
}
```
