#!/bin/bash

go test ./ast --timeout 2s
go test ./lexer --timeout 2s
go test ./parser --timeout 2s
go test ./runtime --timeout 2s
go test ./object --timeout 2s
go test ./compiler --timeout 2s