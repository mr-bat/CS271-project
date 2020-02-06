package main

import (
	"reflect"
	"testing"
)

const testLength = 4
const testId = 1

func TestTimetableInitialization(t *testing.T) {
	initializeTimetable(4)

	for i := 0; i < testLength; i++ {
		for j := 0; j < testLength; j++ {
			if timetable[i][j] != 0 {
				t.Errorf("Wrong initialization of Timetable for i@%d j@%d", i, j)
			}
		}
	}

	setId(testId)
	if getId() != 1 {
		t.Error("Wrong id initialization")
	}
}

func TestTimetableUpdate(t *testing.T) {
	if incTime() != 1 {
		t.Error("Wrong incTime")
	}
	if getTime() != 1 {
		t.Error("Wrong getTime")
	}

	_table := make([][]int, testLength)
	_rows := make([]int, testLength * testLength)
	for i := 0; i < testLength; i++ {
		_table[i] = _rows[i*testLength : (i+1)*testLength]
	}
	_id := 2
	__id := 3
	_table[_id][_id] = 2
	_table[_id][__id] = 1
	_table[__id][__id] = 1


	updateTimetable(_table, _id)

	_result := [testLength][testLength]int{
		{0, 0, 0, 0},
		{0, 1, 2, 1},
		{0, 0, 2, 1},
		{0, 0, 0, 1},
	}
	for i := 0; i < testLength; i++ {
		for j := 0; j < testLength; j++ {
			if timetable[i][j] != _result[i][j] {
				t.Errorf("Wrong update for i@%d j@%d, current: %d, expected: %d", i, j, timetable[i][j], _result[i][j])
			}
		}
	}
}

func TestTimetableParser(t *testing.T) {
	initializeTimetable(2)
	incTime()

	if convertTtToString() != "0&0\n0&1" {
		t.Error("Wrong timetable conversion")
		t.Error("\tconverted:")
		t.Error(convertTtToString())
	}
	if !reflect.DeepEqual(parseTt("0&0\n0&1"), timetable) {
		t.Error("Wrong timetable parsing")
		t.Error("\tparsed table:")
		t.Error(parseTt("0&0\n0&1"))
		t.Error("\ttimetable:")
		t.Error(timetable)
	}
}