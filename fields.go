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

func projectMove(macro_index int, micro_index int) string {
	r := "place_move"

	var projection [9][9]string

	projection[0][0] = "0 0"
	projection[0][1] = "1 0"
	projection[0][2] = "2 0"
	projection[0][3] = "0 1"
	projection[0][4] = "1 1"
	projection[0][5] = "2 1"
	projection[0][6] = "0 2"
	projection[0][7] = "1 2"
	projection[0][8] = "2 2"

	projection[1][0] = "3 0"
	projection[1][1] = "4 0"
	projection[1][2] = "5 0"
	projection[1][3] = "3 1"
	projection[1][4] = "4 1"
	projection[1][5] = "5 1"
	projection[1][6] = "3 2"
	projection[1][7] = "4 2"
	projection[1][8] = "5 2"

	projection[2][0] = "6 0"
	projection[2][1] = "7 0"
	projection[2][2] = "8 0"
	projection[2][3] = "6 1"
	projection[2][4] = "7 1"
	projection[2][5] = "8 1"
	projection[2][6] = "6 2"
	projection[2][7] = "7 2"
	projection[2][8] = "8 2"

	projection[3][0] = "0 3"
	projection[3][1] = "1 3"
	projection[3][2] = "2 3"
	projection[3][3] = "0 4"
	projection[3][4] = "1 4"
	projection[3][5] = "2 4"
	projection[3][6] = "0 5"
	projection[3][7] = "1 5"
	projection[3][8] = "2 5"

	projection[4][0] = "3 3"
	projection[4][1] = "4 3"
	projection[4][2] = "5 3"
	projection[4][3] = "3 4"
	projection[4][4] = "4 4"
	projection[4][5] = "5 4"
	projection[4][6] = "3 5"
	projection[4][7] = "4 5"
	projection[4][8] = "5 5"

	projection[5][0] = "6 3"
	projection[5][1] = "7 3"
	projection[5][2] = "8 3"
	projection[5][3] = "6 4"
	projection[5][4] = "7 4"
	projection[5][5] = "8 4"
	projection[5][6] = "6 5"
	projection[5][7] = "7 5"
	projection[5][8] = "8 5"

	projection[6][0] = "0 6"
	projection[6][1] = "1 6"
	projection[6][2] = "2 6"
	projection[6][3] = "0 7"
	projection[6][4] = "1 7"
	projection[6][5] = "2 7"
	projection[6][6] = "0 8"
	projection[6][7] = "1 8"
	projection[6][8] = "2 8"

	projection[7][0] = "3 6"
	projection[7][1] = "4 6"
	projection[7][2] = "5 6"
	projection[7][3] = "3 7"
	projection[7][4] = "4 7"
	projection[7][5] = "5 7"
	projection[7][6] = "3 8"
	projection[7][7] = "4 8"
	projection[7][8] = "5 8"

	projection[8][0] = "6 6"
	projection[8][1] = "7 6"
	projection[8][2] = "8 6"
	projection[8][3] = "6 7"
	projection[8][4] = "7 7"
	projection[8][5] = "8 7"
	projection[8][6] = "6 8"
	projection[8][7] = "7 8"
	projection[8][8] = "8 8"

	r = r + " " + projection[macro_index][micro_index]
	return r
}
