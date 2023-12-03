package helpers

import (
	"testing"
)

func TestStartsAsNum(t *testing.T) {
	num, len, isNum := startsAsNum("123")
	if !isNum || num != 123 || len != 3 {
		t.Errorf("Expected 123, 3, true, got %d, %d, %t", num, len, isNum)
	}
	num, len, isNum = startsAsNum("123..")
	if !isNum || num != 123 || len != 3 {
		t.Errorf("Expected 123, 3, true, got %d, %d, %t", num, len, isNum)
	}
	num, len, isNum = startsAsNum("...123")
	if isNum {
		t.Errorf("Expected 123, 3, true, got %d, %d, %t", num, len, isNum)
	}
	num, len, isNum = startsAsNum("...123.")
	if isNum {
		t.Errorf("Expected 123, 3, true, got %d, %d, %t", num, len, isNum)
	}
	num, len, isNum = startsAsNum("467..114.."[5:])
	if !isNum || num != 114 || len != 3 {
		t.Errorf("Expected 123, 3, true, got %d, %d, %t", num, len, isNum)
	}
	num, len, isNum = startsAsNum("467..114.."[4:])
	if isNum {
		t.Errorf("Expected 123, 3, true, got %d, %d, %t", num, len, isNum)
	}
}

func TestNotTouches(t *testing.T) {
	lines := []string{
		"..........",
		"&.-109.&..",
		"..........",
	}
	if touches(1, 2, 6, lines) {
		t.Errorf("Expected false, got true")
	}
	lines = []string{
		"&......&..",
		"&.-109.&..",
		"&......*..",
	}
	if touches(1, 2, 6, lines) {
		t.Errorf("Expected false, got true")
	}
	lines = []string{
		"&.-109.&..",
		"&......*..",
	}
	if touches(0, 2, 6, lines) {
		t.Errorf("Expected false, got true")
	}
	lines = []string{
		"&......&..",
		"&.-109.&..",
	}
	if touches(1, 2, 6, lines) {
		t.Errorf("Expected false, got true")
	}
}

func TestTouches(t *testing.T) {
	var lines []string
	lines = []string{
		"..........",
		"..-109&...",
		"..........",
	}
	if !touches(1, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		"..........",
		".&-109....",
		"..........",
	}
	if !touches(1, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		"......&...",
		"..-109....",
		"..........",
	}
	if !touches(1, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		"..........",
		"..-109....",
		"......&...",
	}
	if !touches(1, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		".&........",
		"..-109....",
		"..........",
	}
	if !touches(1, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		"..........",
		"..-109....",
		".&........",
	}
	if !touches(1, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		"..-109....",
		".&........",
	}
	if !touches(0, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		"..-109....",
		"......&...",
	}
	if !touches(0, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		".&........",
		"..-109....",
	}
	if !touches(1, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
	lines = []string{
		"......&...",
		"..-109....",
	}
	if !touches(1, 2, 6, lines) {
		t.Errorf("Expected true, got false")
	}
}
