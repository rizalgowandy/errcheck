// This file exists so that we can check that multi-file packages work
package a

import "fmt"

type t struct{}

func (x t) a() error {
	fmt.Println("this method returns an error") // EXCLUDED
//line myfile.txt:100
	fmt.Println("this method also returns an error") // EXCLUDED
	return nil
}

type u struct {
	t t
}

type embedt struct {
	t
}

type embedtalias = embedt

type embedtptralias = *embedt
