package main

import (
	"flag"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

var (
	// 回答に改行を含めるか否か
	ansNewLine = flag.Bool("newline", false, "this flag is answer contains \n")
)

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	outC := make(chan string)
	defer close(outC)
	go func() {
		var buf strings.Builder
		io.Copy(&buf, r)
		r.Close()
		outC <- buf.String()
	}()

	f()

	os.Stdout = stdout
	w.Close()

	return <-outC
}

func Test_main(t *testing.T) {
	flag.Parse()

	testCases := []struct {
		args []string
		ans  []string
	}{
		{
			[]string{
				"9",
				"1",
				"2",
				"3",
				"4",
				"5",
				"6",
				"7",
				"8",
				"9",
			},
			[]string{
				"45",
				"362880",
			},
		},
		{
			[]string{
				"4",
				"1",
				"2",
				"3",
				"4",
			},
			[]string{
				"10",
				"24",
			},
		},
		// TODO: Add test cases.
	}
	for i, tc := range testCases {
		os.Args = nil
		si := strconv.Itoa(i + 1)
		t.Run("Case "+si, func(t *testing.T) {
			os.Args = append(os.Args, strconv.Itoa(len(tc.args)))
			for _, arg := range tc.args {
				os.Args = append(os.Args, arg)
			}

			var ans string
			if len(tc.ans) != 0 {
				ans = strings.Join(tc.ans, "\n")
			} else {
				ans = tc.ans[0]
			}
			if *ansNewLine {
				ans += "\n"
			}

			ret := captureStdout(main)
			if ret != ans {
				t.Errorf("Unexpected output: %s Need: %s", ret, ans)
			}
		})
	}
}
