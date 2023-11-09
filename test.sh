#!/bin/bash

go test ./ast
go test ./lexer
go test ./parser
go test ./evaluator
go test ./object