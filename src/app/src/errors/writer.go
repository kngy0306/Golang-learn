package main

import (
	"fmt"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil{
		return err
	}
	defer f.Close()
	
	sw := safeWriter{w: f}
	// 書き込む段階でerrが発生した場合、sw.errにエラーがストアされる。
	// 最後にエラーがあれば、nilの代わりにストアされたエラーが返る。
	sw.writeln("this is kona")
	sw.writeln("this is kona")
	sw.writeln("this is kona")
	sw.writeln("this is kona")
	
	return sw.err
}

type safeWriter struct {
	w io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil{
		return
	}
	
	_, sw.err = fmt.Fprintln(sw.w, s)
}