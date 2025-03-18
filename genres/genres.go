package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if err := os.MkdirAll("gores", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := genRes("res", "gores\\res.go"); err != nil {
		log.Fatal(err)
	}
}

const variable = "const %s = []byte{%s}\n"

func genRes(resDirPath, outPath string) error {
	out, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	if _, err := out.WriteString("package gores\n\n"); err != nil {
		return err
	}

	err = filepath.Walk(resDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			resBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			if _, err := fmt.Fprintf(out, variable, strings.ReplaceAll(info.Name(), ".", "_"), joinBytes(resBytes)); err != nil {
				return err
			}

		}

		return nil
	})

	return err
}

func joinBytes(p []byte) string {
	var s string

	if len(p) == 0 {
		return s
	}

	for _, v := range p {
		s += strconv.Itoa(int(v)) + ","
	}

	return s[:len(s)-1]
}

func filenameToVarName(filename string) string {
	var s string

	return s
}
