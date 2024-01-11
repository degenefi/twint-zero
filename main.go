package main

import (
	"twint-zero/Core"
	"twint-zero/InputParser"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	Arguments := InputParser.InputParser()
	Core.Main(&(Arguments.Query), &(Arguments.Instance), &(Arguments.Format))
}
