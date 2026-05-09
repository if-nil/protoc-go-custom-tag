package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	var inputFiles, xxxTags string
	var removeTagComment, removeJSONOmitEmpty bool
	flag.StringVar(&inputFiles, "input", "", "pattern to match input file(s)")
	flag.StringVar(&xxxTags, "XXX_skip", "", "tags that should be skipped (applies 'tag:\"-\"') for unknown fields (deprecated since protoc-gen-go v1.4.0)")
	flag.BoolVar(&removeTagComment, "remove_tag_comment", false, "removes tag comments from the generated file(s)")
	flag.BoolVar(&removeJSONOmitEmpty, "remove_json_omitempty", false, "removes omitempty from json tags in the generated file(s)")
	flag.BoolVar(&verbose, "verbose", false, "verbose logging")

	flag.Parse()

	var xxxSkipSlice []string
	if len(xxxTags) > 0 {
		logf("warn: deprecated flag '-XXX_skip' used")
		xxxSkipSlice = strings.Split(xxxTags, ",")
	}

	if inputFiles == "" {
		log.Fatal("input file is mandatory, see: -help")
	}

	if err := processMatchedFiles(inputFiles, xxxSkipSlice, removeTagComment, removeJSONOmitEmpty); err != nil {
		log.Fatal(err)
	}
}
