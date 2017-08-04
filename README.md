## lncounter

### Build
```bash
go build lncounter/main.go
```

### Usage
If filenames provided as commandline arguments:
```bash
./main 6_lines.txt lncounter.go .gitignore

# Output
6 6_lines.txt
28 lncounter.go
12 .gitignore
```

If there is no filenames provided, then reads stdin:
```bash
cat 6_lines.txt | ./main

# Output
6
```