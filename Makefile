.PHONY: bg

bg:
	cp /Users/mcs/Downloads/*.csv /Users/mcs/repos/go-brand/test/test.csv
	go test -v -count=1 test/format_suites_test.go
	tr -d '\n' < /Users/mcs/repos/go-brand/test/test.csv | pbcopy

line:
	#cp /Users/mcs/Downloads/*.csv /Users/mcs/repos/go-brand/test/test_line.csv
	go test -v -count=1 test/format_line_test.go
