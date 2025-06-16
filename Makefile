.PHONY: bg

bg:
	go test -v -count=1 test/format_suites_test.go
