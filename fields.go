package main

import (
	"strings"
)

func getFieldValue(sVal string) *bool {
	if val := getInt(sVal); val == pid {
		b := true
		return &b
	} else if val == opp_id {
		b := false
		return &b
	} else {
		return nil
	}
}

func projectFields(fieldString string) Board {
	fs := strings.Split(fieldString, ",")
	fields := Board{}

	fields[0] = getFieldValue(fs[0])
	fields[1] = getFieldValue(fs[1])
	fields[2] = getFieldValue(fs[2])
	fields[3] = getFieldValue(fs[9])
	fields[4] = getFieldValue(fs[10])
	fields[5] = getFieldValue(fs[11])
	fields[6] = getFieldValue(fs[18])
	fields[7] = getFieldValue(fs[19])
	fields[8] = getFieldValue(fs[20])

	fields[9] = getFieldValue(fs[3])
	fields[10] = getFieldValue(fs[4])
	fields[11] = getFieldValue(fs[5])
	fields[12] = getFieldValue(fs[12])
	fields[13] = getFieldValue(fs[13])
	fields[14] = getFieldValue(fs[14])
	fields[15] = getFieldValue(fs[21])
	fields[16] = getFieldValue(fs[22])
	fields[17] = getFieldValue(fs[23])

	fields[18] = getFieldValue(fs[6])
	fields[19] = getFieldValue(fs[7])
	fields[20] = getFieldValue(fs[8])
	fields[21] = getFieldValue(fs[15])
	fields[22] = getFieldValue(fs[16])
	fields[23] = getFieldValue(fs[17])
	fields[24] = getFieldValue(fs[24])
	fields[25] = getFieldValue(fs[25])
	fields[26] = getFieldValue(fs[26])

	fields[27] = getFieldValue(fs[27])
	fields[28] = getFieldValue(fs[28])
	fields[29] = getFieldValue(fs[29])
	fields[30] = getFieldValue(fs[36])
	fields[31] = getFieldValue(fs[37])
	fields[32] = getFieldValue(fs[38])
	fields[33] = getFieldValue(fs[45])
	fields[34] = getFieldValue(fs[46])
	fields[35] = getFieldValue(fs[47])

	fields[36] = getFieldValue(fs[30])
	fields[37] = getFieldValue(fs[31])
	fields[38] = getFieldValue(fs[32])
	fields[39] = getFieldValue(fs[39])
	fields[40] = getFieldValue(fs[40])
	fields[41] = getFieldValue(fs[41])
	fields[42] = getFieldValue(fs[48])
	fields[43] = getFieldValue(fs[49])
	fields[44] = getFieldValue(fs[50])

	fields[45] = getFieldValue(fs[33])
	fields[46] = getFieldValue(fs[34])
	fields[47] = getFieldValue(fs[35])
	fields[48] = getFieldValue(fs[42])
	fields[49] = getFieldValue(fs[43])
	fields[50] = getFieldValue(fs[44])
	fields[51] = getFieldValue(fs[51])
	fields[52] = getFieldValue(fs[52])
	fields[53] = getFieldValue(fs[53])

	fields[54] = getFieldValue(fs[54])
	fields[55] = getFieldValue(fs[55])
	fields[56] = getFieldValue(fs[56])
	fields[57] = getFieldValue(fs[63])
	fields[58] = getFieldValue(fs[64])
	fields[59] = getFieldValue(fs[65])
	fields[60] = getFieldValue(fs[72])
	fields[61] = getFieldValue(fs[73])
	fields[62] = getFieldValue(fs[74])

	fields[63] = getFieldValue(fs[57])
	fields[64] = getFieldValue(fs[58])
	fields[65] = getFieldValue(fs[59])
	fields[66] = getFieldValue(fs[66])
	fields[67] = getFieldValue(fs[67])
	fields[68] = getFieldValue(fs[68])
	fields[69] = getFieldValue(fs[75])
	fields[70] = getFieldValue(fs[76])
	fields[71] = getFieldValue(fs[77])

	fields[72] = getFieldValue(fs[60])
	fields[73] = getFieldValue(fs[61])
	fields[74] = getFieldValue(fs[62])
	fields[75] = getFieldValue(fs[69])
	fields[76] = getFieldValue(fs[70])
	fields[77] = getFieldValue(fs[71])
	fields[78] = getFieldValue(fs[78])
	fields[79] = getFieldValue(fs[79])
	fields[80] = getFieldValue(fs[80])

	return fields
}

var moveStrings = [81]string{
	"0 0",
	"1 0",
	"2 0",
	"0 1",
	"1 1",
	"2 1",
	"0 2",
	"1 2",
	"2 2",
	"3 0",
	"4 0",
	"5 0",
	"3 1",
	"4 1",
	"5 1",
	"3 2",
	"4 2",
	"5 2",
	"6 0",
	"7 0",
	"8 0",
	"6 1",
	"7 1",
	"8 1",
	"6 2",
	"7 2",
	"8 2",
	"0 3",
	"1 3",
	"2 3",
	"0 4",
	"1 4",
	"2 4",
	"0 5",
	"1 5",
	"2 5",
	"3 3",
	"4 3",
	"5 3",
	"3 4",
	"4 4",
	"5 4",
	"3 5",
	"4 5",
	"5 5",
	"6 3",
	"7 3",
	"8 3",
	"6 4",
	"7 4",
	"8 4",
	"6 5",
	"7 5",
	"8 5",
	"0 6",
	"1 6",
	"2 6",
	"0 7",
	"1 7",
	"2 7",
	"0 8",
	"1 8",
	"2 8",
	"3 6",
	"4 6",
	"5 6",
	"3 7",
	"4 7",
	"5 7",
	"3 8",
	"4 8",
	"5 8",
	"6 6",
	"7 6",
	"8 6",
	"6 7",
	"7 7",
	"8 7",
	"6 8",
	"7 8",
	"8 8"}

func projectMove(index int) string {
	r := "place_move"
	r += " " + moveStrings[index]
	return r
}
