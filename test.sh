#!/bin/bash

go test ./ast
go test ./lexer
go test ./parser
go test ./runtime
go test ./object
go test ./compiler --timeout 2s