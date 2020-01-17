package helpers

import (
	"github.com/velann21/roulette"
	"io/ioutil"
	"log"
)

func ReadFile(path string) []byte {
	ruleFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("ruleFile read err #%v  at path %v", err, path)
	}

	return ruleFile
}

type RuleSet struct {
	Executor roulette.SimpleExecute
}
var ruleSetObject *RuleSet
func (ruleSet *RuleSet)LoadRuleSet(path string) {
	config := roulette.TextTemplateParserConfig{}
	parser, err := roulette.NewParser(ReadFile(path), config)
	if err != nil {
		log.Fatal(err)
	}
	executor := roulette.NewSimpleExecutor(parser)
	ruleSet.Executor = executor
	ruleSetObject = ruleSet
}

func GetRuleSetObject() *RuleSet{
	return ruleSetObject
}





