package uc

import (
	"os"
)

//handle func export to implementation
type Handler interface {
	FormatLogic
}

type FormatLogic interface {
	FormatTCP(inputPath string, outputPath string, config ConfigPSPFormat) (*os.File, error)
	ImportTCP(inputPath string, tempPath string) ([][][]string, error)
}

//handle implementation follow rule use case that is the same page
type HandlerConstructor struct {
	//	Format Format
}

func (c HandlerConstructor) New() Handler {

	return interactor{}
}
