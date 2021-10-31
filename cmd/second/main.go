// You can edit this code!
package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

const SEPARATOR_LINE = "-----------"

/// 用来解耦第3方库
var NotFound = errors.New("not found")

func oprDB() error {
	// do some db opr
	//  ...
	rand.Seed(time.Now().Unix())
	r := rand.Intn(2)
	if r == 0 {
		return nil
	}

	return sql.ErrNoRows
}

func innerDao(querySql string) error {
	err := oprDB()

	// 根因处使用Wrap，生成堆栈信息
	if err == sql.ErrNoRows {
		return errors.Wrap(NotFound, fmt.Sprintf("data  not found, sql[%s]", querySql))
	}

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("db query system err, sql[%s]", querySql))
	}

	// do something
	return nil
}

func Dao(querySql string) error {
	//
	err := innerDao(querySql)
	if err != nil {
		// 非根因处使用WithMessage，无需生成堆栈信息
		return errors.WithMessage(err, "Dao")
	}

	return nil
}

func main() {
	querySql := "this is a executable sql"
	err := Dao(querySql)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sql exec succ")
	}
	fmt.Println(SEPARATOR_LINE)

	if errors.Is(err, NotFound) {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(SEPARATOR_LINE)

	if errors.Is(errors.Cause(err), NotFound) {
		fmt.Printf("%+v\n", err)
	}

	return
}
