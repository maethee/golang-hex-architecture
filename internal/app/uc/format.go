package uc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	validation "github.axa.com/maethee-chakkuchantorn/psp-improvement/internal/pkg/excel.validation"
	"github.axa.com/maethee-chakkuchantorn/psp-improvement/internal/pkg/reader"
)

type ConfigPSPFormat struct {
	Sheet int
	Type  string
	Rules []validation.RuleProps
}

func LoadConfiguration(file string) ConfigPSPFormat {
	var config ConfigPSPFormat
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func (i interactor) FormatTCP(inputPath string, outputPath string, config ConfigPSPFormat) (*os.File, error) {
	var validationFormatTcp = &validation.Validation{}

	err, rules := validation.CreateRules(config.Rules)
	if err != nil {
		panic(err)
	}

	for _, v := range rules {
		fmt.Println(v)
		validationFormatTcp.Add(v)
	}

	file, err := reader.Read(inputPath)
	if err != nil {
		panic(err)
	}

	validationFormatTcp.FormatSheet(file, validation.Sheet(config.Sheet), validation.File(file))
	err = file.Save(outputPath)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(outputPath)
	if err != nil {
		panic(err)
	}

	return f, nil
}

func (i interactor) ImportTCP(inputPath string, tempPath string) ([][][]string, error) {
	fmt.Println(inputPath)
	fmt.Println(tempPath)
	file, err := reader.Read(inputPath)
	if err != nil {
		log.Panicln(err)
	}

	err = file.Save(tempPath)
	if err != nil {
		log.Panicln(err)
	}

	return file.ToSlice()
}
