package rules

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	db "github.com/danielmachado86/contracts/db"
	"github.com/stretchr/testify/require"
)

func loadTemplateFile(t *testing.T, template_file string) (Specs, error) {
	file, err := ioutil.ReadFile(template_file)
	require.NoError(t, err)

	template := Specs{}
	err = json.Unmarshal([]byte(file), &template)
	return template, err
}

func TestParties(t *testing.T) {

	template, err := loadTemplateFile(t, "rental_contract.json")
	require.NoError(t, err)

	result := make(ResultContract)
	result["metadata"] = meta

	ruleRegistry := make(map[string]Rule)

	for _, spec := range template {
		ruleRegistry[spec.Name], err = CreateRule(spec)
		require.NoError(t, err)
	}

	rule := ruleRegistry["contract_parties"]
	err = rule.Run(ruleRegistry)
	require.NoError(t, err)

	got := rule.GetOutput("contract_parties").([]*db.Party)
	require.Len(t, got, 2)
}

func TestSignatures(t *testing.T) {

	template, err := loadTemplateFile(t, "rental_contract.json")
	require.NoError(t, err)

	party1 := &db.Party{
		Username:   "DanielM",
		ContractID: 1,
	}
	party2 := &db.Party{
		Username:   "JimenaL",
		ContractID: 1,
	}

	var parties []*db.Party
	parties = append(parties, party1)
	parties = append(parties, party2)

	var signatureList = []*db.Signature{}

	signature1 := &db.Signature{
		Username:   party1.Username,
		ContractID: party1.ContractID,
		CreatedAt:  time.Now(),
	}
	signature2 := &db.Signature{
		Username:   party2.Username,
		ContractID: party2.ContractID,
		CreatedAt:  time.Now(),
	}

	signatureList = append(signatureList, signature1)
	signatureList = append(signatureList, signature2)

	ruleRegistry := make(map[string]Rule)

	for _, attribute := range template.Attributes {
		ruleRegistry[attribute.Name], err = CreateRule(attribute)
		require.NoError(t, err)
	}

	ruleRegistry["contract_parties"].(*Parties).Parties = parties
	ruleRegistry["contract_signatures"].(*EnoughSignatures).Signatures = signatureList

	rule := ruleRegistry["contract_signatures"]
	err = rule.Run(ruleRegistry)
	require.NoError(t, err)

	got := rule.GetOutput("contract_signatures").([]*db.Signature)
	require.Len(t, got, 2)
}

func TestIsSignedRule(t *testing.T) {

	template, err := loadTemplateFile(t, "rental_contract.json")
	require.NoError(t, err)

	party1 := &db.Party{
		Username:   "DanielM",
		ContractID: 1,
	}
	party2 := &db.Party{
		Username:   "JimenaL",
		ContractID: 1,
	}

	var parties []*db.Party
	parties = append(parties, party1)
	parties = append(parties, party2)

	var signatureList = []*db.Signature{}

	signature1 := &db.Signature{
		Username:   party1.Username,
		ContractID: party1.ContractID,
		CreatedAt:  time.Now(),
	}
	signature2 := &db.Signature{
		Username:   party2.Username,
		ContractID: party2.ContractID,
		CreatedAt:  time.Now(),
	}

	signatureList = append(signatureList, signature1)
	signatureList = append(signatureList, signature2)

	ruleRegistry := make(map[string]Rule)

	for _, attribute := range template.Attributes {
		ruleRegistry[attribute.Name], err = CreateRule(attribute)
		require.NoError(t, err)
	}

	ruleRegistry["contract_parties"].(*Parties).Parties = parties
	ruleRegistry["contract_signatures"].(*EnoughSignatures).Signatures = signatureList

	rule := ruleRegistry["contract_is_signed"]
	err = rule.Run(ruleRegistry)
	require.NoError(t, err)

	got := rule.GetOutput("contract_is_signed").(bool)
	require.True(t, got)
}

func TestSignatureDateRule(t *testing.T) {

	template, err := loadTemplateFile(t, "rental_contract.json")
	require.NoError(t, err)

	party1 := &db.Party{
		Username:   "DanielM",
		ContractID: 1,
	}
	party2 := &db.Party{
		Username:   "JimenaL",
		ContractID: 1,
	}

	var parties []*db.Party
	parties = append(parties, party1)
	parties = append(parties, party2)

	var signatureList = []*db.Signature{}

	signature1 := &db.Signature{
		Username:   party1.Username,
		ContractID: party1.ContractID,
		CreatedAt:  time.Now(),
	}
	signature2 := &db.Signature{
		Username:   party2.Username,
		ContractID: party2.ContractID,
		CreatedAt:  time.Now(),
	}

	signatureList = append(signatureList, signature1)
	signatureList = append(signatureList, signature2)

	ruleRegistry := make(map[string]Rule)

	for _, attribute := range template.Attributes {
		ruleRegistry[attribute.Name], err = CreateRule(attribute)
		require.NoError(t, err)
	}

	ruleRegistry["contract_parties"].(*Parties).Parties = parties
	ruleRegistry["contract_signatures"].(*EnoughSignatures).Signatures = signatureList

	rule := ruleRegistry["contract_signature_date"]
	err = rule.Run(ruleRegistry)
	require.NoError(t, err)

	got := rule.GetOutput("contract_signature_date").(time.Time)
	require.WithinDuration(t, got, time.Now(), time.Second)

}

func TestSAllRules(t *testing.T) {

	template, err := loadTemplateFile(t, "rental_contract.json")
	require.NoError(t, err)

	party1 := &db.Party{
		Username:   "DanielM",
		ContractID: 1,
	}
	party2 := &db.Party{
		Username:   "JimenaL",
		ContractID: 1,
	}

	var partyList []*db.Party
	partyList = append(partyList, party1)
	partyList = append(partyList, party2)

	var signatureList []*db.Signature

	signature1 := &db.Signature{
		Username:   party1.Username,
		ContractID: party1.ContractID,
		CreatedAt:  time.Now(),
	}
	signature2 := &db.Signature{
		Username:   party2.Username,
		ContractID: party2.ContractID,
		CreatedAt:  time.Now(),
	}

	signatureList = append(signatureList, signature1)
	signatureList = append(signatureList, signature2)

	ruleRegistry := make(map[string]Rule)

	for _, attribute := range template.Attributes {
		ruleRegistry[attribute.Name], err = CreateRule(attribute)
		require.NoError(t, err)
	}

	ruleRegistry["contract_parties"].(*Parties).Parties = partyList
	ruleRegistry["contract_signatures"].(*EnoughSignatures).Signatures = signatureList

	for _, rule := range ruleRegistry {
		err = rule.Run(ruleRegistry)
		require.NoError(t, err)
	}

	rule := ruleRegistry["contract_parties"]
	parties := rule.GetOutput("contract_parties").([]*db.Party)
	require.Len(t, parties, 2)

	rule = ruleRegistry["contract_signatures"]
	signatures := rule.GetOutput("contract_signatures").([]*db.Signature)
	require.Len(t, signatures, 2)

	rule = ruleRegistry["contract_is_signed"]
	is_signed := rule.GetOutput("contract_is_signed").(bool)
	require.True(t, is_signed)

	rule = ruleRegistry["contract_signature_date"]
	signature_date := rule.GetOutput("contract_signature_date").(time.Time)
	require.WithinDuration(t, signature_date, time.Now(), time.Second)

}
