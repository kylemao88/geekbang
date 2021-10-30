// You can edit this code!
package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

const SEPARATOR_LINE = "-----------"

func fn1() error {
	err := fn2()
	// 非根因处使用WithMessage，无需生成堆栈信息
	return errors.WithMessage(err, "fn2")
}

func fn2() error {
	// do some db opr
	//  ...

	// 根因处使用Wrap，生成堆栈信息
	err := errors.Wrap(sql.ErrNoRows, "fn1")
	return err
}

func main() {
	err := fn1()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(SEPARATOR_LINE)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(SEPARATOR_LINE)

	if errors.Is(errors.Cause(err), sql.ErrNoRows) {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(SEPARATOR_LINE)

	return
}
