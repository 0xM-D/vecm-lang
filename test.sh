#!/bin/bash

go test ./ast
go test ./lexer
go test ./parser/tests
go test ./evaluator/tests
go test ./object