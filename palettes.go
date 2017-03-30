package main

import "image/color"

var teleFfPalettes = [][]color.Color{
	{
		IRGB{0x0000},
		IRGB{0x0000},
		IRGB{0x4FEC},
		IRGB{0x5F00},
		IRGB{0x8F00},
		IRGB{0xCF20},
		IRGB{0xFF20},
		IRGB{0xEFFF},
		IRGB{0x8F00},
		IRGB{0xCF21},
		IRGB{0xFF76},
		IRGB{0xFF98},
		IRGB{0xFFDD},
		IRGB{0xFFFF},
		IRGB{0x0000},
		IRGB{0x0000},
	},
}

var floorPalettes = [][]color.Color{
	{
		IRGB{0x0}, // palette 0
		IRGB{0x6631},
		IRGB{0x9631},
		IRGB{0xC631},
		IRGB{0xF631},
		IRGB{0xF643},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x8333},
		IRGB{0x7223},
		IRGB{0xA324},
		IRGB{0xF334},
		IRGB{0xC334},
		IRGB{0xC445},
		IRGB{0xE445},
		IRGB{0xF445},
	}, {
		IRGB{0x0}, // palette 1
		IRGB{0x7222},
		IRGB{0x9444},
		IRGB{0x7753},
		IRGB{0xF421},
		IRGB{0xB842},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x8421},
		IRGB{0x7422},
		IRGB{0x9532},
		IRGB{0xC532},
		IRGB{0xF422},
		IRGB{0x7974},
		IRGB{0xE532},
		IRGB{0xF643},
	}, {
		IRGB{0x0}, // palette 2
		IRGB{0x6633},
		IRGB{0x9633},
		IRGB{0xC734},
		IRGB{0x7746},
		IRGB{0xF643},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x8333},
		IRGB{0x7223},
		IRGB{0xA324},
		IRGB{0xF334},
		IRGB{0xC334},
		IRGB{0xC445},
		IRGB{0xE645},
		IRGB{0xF445},
	}, {
		IRGB{0x3700}, // palette 3
		IRGB{0xA520},
		IRGB{0xD533},
		IRGB{0xE633},
		IRGB{0x8A40},
		IRGB{0xC833},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x6733},
		IRGB{0x3522},
		IRGB{0x8522},
		IRGB{0xB633},
		IRGB{0xF522},
		IRGB{0x9B60},
		IRGB{0x5F43},
		IRGB{0xBA66},
	}, {
		IRGB{0x10}, // palette 4
		IRGB{0x4272},
		IRGB{0x69A2},
		IRGB{0x56A5},
		IRGB{0x35F3},
		IRGB{0x9353},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x10F0},
		IRGB{0x1030},
		IRGB{0x131},
		IRGB{0x6131},
		IRGB{0x8131},
		IRGB{0x23F0},
		IRGB{0xF121},
		IRGB{0xC252},
	}, {
		IRGB{0x7003}, // palette 5
		IRGB{0x8108},
		IRGB{0xF226},
		IRGB{0xC224},
		IRGB{0x9336},
		IRGB{0x9336},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x6337},
		IRGB{0x8004},
		IRGB{0xA124},
		IRGB{0x5226},
		IRGB{0xA336},
		IRGB{0x9368},
		IRGB{0x9558},
		IRGB{0xF447},
	}, {
		IRGB{0x0}, // palette 6
		IRGB{0x5551},
		IRGB{0xE552},
		IRGB{0x7663},
		IRGB{0x7774},
		IRGB{0x7885},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x7441},
		IRGB{0x3551},
		IRGB{0x4552},
		IRGB{0x7663},
		IRGB{0x7774},
		IRGB{0x7885},
		IRGB{0x7995},
		IRGB{0xA996},
	}, {
		IRGB{0x0}, // palette 7
		IRGB{0x5155},
		IRGB{0x7233},
		IRGB{0xA366},
		IRGB{0x7477},
		IRGB{0x7588},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x7144},
		IRGB{0x1155},
		IRGB{0x5255},
		IRGB{0x8355},
		IRGB{0x7477},
		IRGB{0x7588},
		IRGB{0x8750},
		IRGB{0xA699},
	}, {
		IRGB{0x0}, // palette 8
		IRGB{0x7515},
		IRGB{0xB520},
		IRGB{0x6636},
		IRGB{0x7747},
		IRGB{0x7858},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x7414},
		IRGB{0x2515},
		IRGB{0x5525},
		IRGB{0xA636},
		IRGB{0x7747},
		IRGB{0x7858},
		IRGB{0x5C70},
		IRGB{0x9969},
	}, {
		IRGB{0x7520}, // palette 9
		IRGB{0x3FA0},
		IRGB{0x9950},
		IRGB{0xBF90},
		IRGB{0x8F90},
		IRGB{0xB842},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0xE531},
		IRGB{0x8530},
		IRGB{0x8420},
		IRGB{0x6C92},
		IRGB{0x9B70},
		IRGB{0x8FA0},
		IRGB{0xF530},
		IRGB{0x9FF8},
	}, {
		IRGB{0x0}, // palette 10
		IRGB{0x4600},
		IRGB{0x5800},
		IRGB{0x4F00},
		IRGB{0xC522},
		IRGB{0xB842},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x9421},
		IRGB{0x2700},
		IRGB{0x7200},
		IRGB{0x4C32},
		IRGB{0x4F22},
		IRGB{0x5E52},
		IRGB{0xF532},
		IRGB{0xF743},
	}, {
		IRGB{0x0}, // palette 11
		IRGB{0x8008},
		IRGB{0xF006},
		IRGB{0xF004},
		IRGB{0x9006},
		IRGB{0x9006},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x2007},
		IRGB{0x1007},
		IRGB{0x4},
		IRGB{0x5006},
		IRGB{0xA005},
		IRGB{0x8009},
		IRGB{0x5008},
		IRGB{0xF108},
	}, {
		IRGB{0x2877}, // palette 12
		IRGB{0x5752},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x0},
		IRGB{0x7642},
		IRGB{0x5053},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0x4222}, // palette 13
		IRGB{0x4666},
		IRGB{0xD666},
		IRGB{0xA224},
		IRGB{0x0},
		IRGB{0x9647},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x0},
		IRGB{0x7AAA},
		IRGB{0x9555},
		IRGB{0x2461},
		IRGB{0x3777},
		IRGB{0x0},
		IRGB{0x5324},
		IRGB{0x2222},
	}, {
		IRGB{0x0}, // palette 14
		IRGB{0x3527},
		IRGB{0x8C3B},
		IRGB{0xF637},
		IRGB{0x5C3B},
		IRGB{0xE835},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x6736},
		IRGB{0x888},
		IRGB{0x526},
		IRGB{0x7637},
		IRGB{0xB527},
		IRGB{0x9A55},
		IRGB{0xC844},
		IRGB{0xBA66},
	}, {
		IRGB{0x0}, // palette 15
		IRGB{0x3253},
		IRGB{0x83C4},
		IRGB{0xA363},
		IRGB{0x43BB},
		IRGB{0xE395},
		IRGB{0x0},
		IRGB{0xFFFF},
		IRGB{0x5376},
		IRGB{0x0},
		IRGB{0x4272},
		IRGB{0xA388},
		IRGB{0xB253},
		IRGB{0x8852},
		IRGB{0xD842},
		IRGB{0x86A6},
	},
}

var wallPalettes = [][]color.Color{
	{
		IRGB{0xF410}, // palette 0
		IRGB{0xF010},
		IRGB{0xF110},
		IRGB{0xD420},
		IRGB{0xB520},
		IRGB{0x4531},
		IRGB{0xF531},
		IRGB{0xF831},
		IRGB{0xFC63},
		IRGB{0xF840},
		IRGB{0xFD84},
		IRGB{0xF335},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xAAAA}, // palette 1
		IRGB{0xF223},
		IRGB{0x9323},
		IRGB{0xC434},
		IRGB{0x7434},
		IRGB{0x9535},
		IRGB{0xD545},
		IRGB{0xE656},
		IRGB{0xF989},
		IRGB{0xC656},
		IRGB{0xFB9B},
		IRGB{0xF335},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF410}, // palette 2
		IRGB{0xF023},
		IRGB{0xC023},
		IRGB{0x9034},
		IRGB{0xF034},
		IRGB{0xC035},
		IRGB{0xE045},
		IRGB{0xF156},
		IRGB{0xF289},
		IRGB{0xE056},
		IRGB{0xF4BA},
		IRGB{0xF098},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF410}, // palette 3
		IRGB{0xF073},
		IRGB{0x8080},
		IRGB{0x4081},
		IRGB{0x8070},
		IRGB{0x7090},
		IRGB{0xE080},
		IRGB{0xA090},
		IRGB{0xB1D5},
		IRGB{0xB3A0},
		IRGB{0xC8F8},
		IRGB{0xF0E9},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0xF8BD},
		IRGB{0x0},
	}, {
		IRGB{0xD310}, // palette 4
		IRGB{0x5400},
		IRGB{0x9500},
		IRGB{0xA600},
		IRGB{0x7700},
		IRGB{0x9600},
		IRGB{0xB900},
		IRGB{0xE801},
		IRGB{0xFC01},
		IRGB{0xAA00},
		IRGB{0xFF00},
		IRGB{0xAA70},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF410}, // palette 5
		IRGB{0xF400},
		IRGB{0xF620},
		IRGB{0xF720},
		IRGB{0xF830},
		IRGB{0xF930},
		IRGB{0xCB32},
		IRGB{0xFB40},
		IRGB{0xFF60},
		IRGB{0x9D50},
		IRGB{0xFF87},
		IRGB{0xFFBA},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF410}, // palette 6
		IRGB{0xD403},
		IRGB{0xB604},
		IRGB{0xB705},
		IRGB{0xB806},
		IRGB{0xB906},
		IRGB{0xEB17},
		IRGB{0xCB08},
		IRGB{0xCF5C},
		IRGB{0xAD09},
		IRGB{0xDF77},
		IRGB{0xDF9E},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF410}, // palette 7
		IRGB{0xF022},
		IRGB{0xA022},
		IRGB{0xD032},
		IRGB{0xC032},
		IRGB{0xC032},
		IRGB{0xC053},
		IRGB{0xE052},
		IRGB{0xF173},
		IRGB{0xF053},
		IRGB{0xF0A7},
		IRGB{0xF2C8},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF410}, // palette 8
		IRGB{0x2222},
		IRGB{0x9999},
		IRGB{0x3CCC},
		IRGB{0x9AAA},
		IRGB{0x9777},
		IRGB{0xBBBB},
		IRGB{0x6EEE},
		IRGB{0xCEEE},
		IRGB{0xAAAA},
		IRGB{0xFFFF},
		IRGB{0xBBBB},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0x0}, // palette 9
		IRGB{0x6960},
		IRGB{0xB960},
		IRGB{0xC942},
		IRGB{0x9A50},
		IRGB{0xA953},
		IRGB{0xF741},
		IRGB{0xCC81},
		IRGB{0xEFA0},
		IRGB{0xAE94},
		IRGB{0xCFC0},
		IRGB{0xFFF8},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xD310}, // palette 10
		IRGB{0x0},
		IRGB{0x7444},
		IRGB{0x9222},
		IRGB{0x9222},
		IRGB{0x7222},
		IRGB{0x8333},
		IRGB{0xB333},
		IRGB{0x8444},
		IRGB{0x5111},
		IRGB{0xA666},
		IRGB{0xA666},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF310}, // palette 11
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0xA100},
		IRGB{0x9200},
		IRGB{0x9500},
		IRGB{0xF400},
		IRGB{0xF500},
		IRGB{0xF700},
		IRGB{0xC200},
		IRGB{0xF900},
		IRGB{0xFA00},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0x0}, // palette 12
		IRGB{0x6608},
		IRGB{0x4608},
		IRGB{0x6707},
		IRGB{0x7606},
		IRGB{0x590A},
		IRGB{0x6B1D},
		IRGB{0x9B0C},
		IRGB{0x9F5F},
		IRGB{0x7D0E},
		IRGB{0xBF5F},
		IRGB{0xEF5F},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0x0}, // palette 13
		IRGB{0x0},
		IRGB{0x303F},
		IRGB{0x603F},
		IRGB{0x8039},
		IRGB{0x603F},
		IRGB{0x803F},
		IRGB{0xB04F},
		IRGB{0xF05D},
		IRGB{0x877F},
		IRGB{0xF5AF},
		IRGB{0xF8FF},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0x0}, // palette 14
		IRGB{0x0},
		IRGB{0x3A30},
		IRGB{0x2C30},
		IRGB{0x2F70},
		IRGB{0x3F40},
		IRGB{0x4F50},
		IRGB{0x6F51},
		IRGB{0x7F72},
		IRGB{0x3F60},
		IRGB{0xDA31},
		IRGB{0xFB40},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF410}, // palette 15
		IRGB{0x828F},
		IRGB{0x558F},
		IRGB{0x56AF},
		IRGB{0x88AF},
		IRGB{0x857F},
		IRGB{0x9ABF},
		IRGB{0x77AF},
		IRGB{0xCABF},
		IRGB{0xA8AF},
		IRGB{0xDDDF},
		IRGB{0xBBBF},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
		IRGB{0x0},
	}, {
		IRGB{0xF333}, // mystery palette
		IRGB{0xF333},
		IRGB{0xF030},
		IRGB{0xF040},
		IRGB{0xF060},
		IRGB{0xF180},
		IRGB{0xF290},
		IRGB{0xF730},
		IRGB{0xFA60},
		IRGB{0xF000},
		IRGB{0xF000},
		IRGB{0xF000},
		IRGB{0xF000},
		IRGB{0xF000},
		IRGB{0xF555},
		IRGB{0xF444},
	},
}

var basePalettes = [][]color.Color{
	{
		IRGB{0x0},    // palette 0 - death, doors, maybe thief, trans, forcefield hub
		IRGB{0x0},    // 1
		IRGB{0x988F}, // 2
		IRGB{0xA99F}, // 3
		IRGB{0xC99F}, // 4
		IRGB{0xFBBF}, // 5
		IRGB{0xA888}, // 6
		IRGB{0x0},    // 7
		IRGB{0x9999}, // 8
		IRGB{0x7F8F}, // 9
		IRGB{0x711F}, // 10
		IRGB{0x5F93}, // 11
		IRGB{0x6F84}, // 12
		IRGB{0xEF86}, // 13
		IRGB{0xF0F0}, // 14
		IRGB{0xFFBB}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 1 - keys, treas, pots, food, acid, movable wall, pickups
		IRGB{0x0},    // 1
		IRGB{0x7F84}, // 2
		IRGB{0x8F85}, // 3
		IRGB{0xAF83}, // 4
		IRGB{0x40F7}, // 5
		IRGB{0xFF20}, // 6
		IRGB{0xFF82}, // 7
		IRGB{0xFFC0}, // 8
		IRGB{0xFFFF}, // 9
		IRGB{0x5F50}, // 10
		IRGB{0x81F8}, // 11
		IRGB{0xEF88}, // 12
		IRGB{0xD79F}, // 13
		IRGB{0xA17F}, // 14
		IRGB{0xA00F}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 2 - ghost 1, grunt 1
		IRGB{0x0},    // 1
		IRGB{0x0},    // 2
		IRGB{0x7165}, // 3
		IRGB{0x7387}, // 4
		IRGB{0x7598}, // 5
		IRGB{0x76BA}, // 6
		IRGB{0x78DC}, // 7
		IRGB{0x7AFE}, // 8
		IRGB{0x0},    // 9
		IRGB{0x0},    // 10
		IRGB{0x7E7A}, // 11
		IRGB{0x7FDF}, // 12
		IRGB{0x7FC0}, // 13
		IRGB{0x3FF0}, // 14
		IRGB{0x7F72}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 3 - ghost 2, grunt 2
		IRGB{0x0},    // 1
		IRGB{0xFFFF}, // 2
		IRGB{0xA165}, // 3
		IRGB{0xA387}, // 4
		IRGB{0xA598}, // 5
		IRGB{0xA6BA}, // 6
		IRGB{0xA8DC}, // 7
		IRGB{0xAAFE}, // 8
		IRGB{0x0},    // 9
		IRGB{0xFFFF}, // 10
		IRGB{0xCE7A}, // 11
		IRGB{0xAFDF}, // 12
		IRGB{0x8FC0}, // 13
		IRGB{0x5FF0}, // 14
		IRGB{0xBF72}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 4 - ghost 3, grunt 3
		IRGB{0x0},    // 1
		IRGB{0xFF00}, // 2
		IRGB{0xF165}, // 3
		IRGB{0xF387}, // 4
		IRGB{0xF598}, // 5
		IRGB{0xF6BA}, // 6
		IRGB{0xF8DC}, // 7
		IRGB{0xFAFE}, // 8
		IRGB{0x0},    // 9
		IRGB{0xFF00}, // 10
		IRGB{0xFE7A}, // 11
		IRGB{0xFFDF}, // 12
		IRGB{0xCFC0}, // 13
		IRGB{0x8FF0}, // 14
		IRGB{0xFF72}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 5 - generators
		IRGB{0x0},    // 1
		IRGB{0x6ABF}, // 2
		IRGB{0x8BBF}, // 3
		IRGB{0xBAAF}, // 4
		IRGB{0xDCBF}, // 5
		IRGB{0xFDDF}, // 6
		IRGB{0xFFFF}, // 7
		IRGB{0x5F77}, // 8
		IRGB{0x6F99}, // 9
		IRGB{0x8FAA}, // 10
		IRGB{0x9FBB}, // 11
		IRGB{0xAFCC}, // 12
		IRGB{0xCFDD}, // 13
		IRGB{0x5BBF}, // 14
		IRGB{0xFF00}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 6 - demon 1
		IRGB{0x0},    // 1
		IRGB{0x3FC5}, // 2
		IRGB{0x5F84}, // 3
		IRGB{0x5F96}, // 4
		IRGB{0x6F86}, // 5
		IRGB{0xAFC0}, // 6
		IRGB{0x6FB1}, // 7
		IRGB{0x6F9A}, // 8
		IRGB{0x6F30}, // 9
		IRGB{0x8F00}, // 10
		IRGB{0x6F55}, // 11
		IRGB{0x2F19}, // 12
		IRGB{0x4F00}, // 13
		IRGB{0x6F11}, // 14
		IRGB{0xAFFF}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 7 - demon 2
		IRGB{0x0},    // 1
		IRGB{0x4FC5}, // 2
		IRGB{0x7F84}, // 3
		IRGB{0x7F96}, // 4
		IRGB{0x8F86}, // 5
		IRGB{0xCFC0}, // 6
		IRGB{0x9FB1}, // 7
		IRGB{0x9F9A}, // 8
		IRGB{0x9F30}, // 9
		IRGB{0xCF00}, // 10
		IRGB{0x9F55}, // 11
		IRGB{0x3F19}, // 12
		IRGB{0x6F00}, // 13
		IRGB{0x8F11}, // 14
		IRGB{0xCFFF}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 8 - demon 3, IT, dragon
		IRGB{0x0},    // 1
		IRGB{0x5FC5}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xBF96}, // 4
		IRGB{0xDF86}, // 5
		IRGB{0xFFC0}, // 6
		IRGB{0xFFB1}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0xFF30}, // 9
		IRGB{0xFF00}, // 10
		IRGB{0xFF55}, // 11
		IRGB{0x4F19}, // 12
		IRGB{0x9F00}, // 13
		IRGB{0xDF11}, // 14
		IRGB{0xFFFF}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 9 - lobber1, sorc1
		IRGB{0x0},    // 1
		IRGB{0x0},    // 2
		IRGB{0x0},    // 3
		IRGB{0x9F9A}, // 4
		IRGB{0xAF9A}, // 5
		IRGB{0x2909}, // 6
		IRGB{0x6A08}, // 7
		IRGB{0x9806}, // 8
		IRGB{0x8888}, // 9
		IRGB{0x9CCF}, // 10
		IRGB{0x4F33}, // 11
		IRGB{0x5F89}, // 12
		IRGB{0x7F83}, // 13
		IRGB{0xBF94}, // 14
		IRGB{0x9F11}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 10 - lobs2, sorcs2
		IRGB{0x0},    // 1
		IRGB{0xFFFF}, // 2
		IRGB{0x0},    // 3
		IRGB{0x9F9A}, // 4
		IRGB{0xCF9A}, // 5
		IRGB{0x4909}, // 6
		IRGB{0x9A08}, // 7
		IRGB{0xC806}, // 8
		IRGB{0x8888}, // 9
		IRGB{0xACCF}, // 10
		IRGB{0x5F33}, // 11
		IRGB{0x6F89}, // 12
		IRGB{0x7F83}, // 13
		IRGB{0xBF94}, // 14
		IRGB{0xBF11}, // 15
	}, {
		IRGB{0x0},    // 0 // palette 11 - lobs3, sorcs3
		IRGB{0x0},    // 1
		IRGB{0xFF00}, // 2
		IRGB{0x0},    // 3
		IRGB{0x9F9A}, // 4
		IRGB{0xEF9A}, // 5
		IRGB{0x6909}, // 6
		IRGB{0xCA08}, // 7
		IRGB{0xF806}, // 8
		IRGB{0x8888}, // 9
		IRGB{0xBCCF}, // 10
		IRGB{0x6F33}, // 11
		IRGB{0x7F89}, // 12
		IRGB{0x7F83}, // 13
		IRGB{0xBF94}, // 14
		IRGB{0xDF11}, // 15
	},
}

// Warrior
var warriorPalettes = [][]color.Color{
	{
		IRGB{0x0},    // Red
		IRGB{0x0},    // 1
		IRGB{0x6944}, // 2
		IRGB{0x9B74}, // 3
		IRGB{0xBF96}, // 4
		IRGB{0xDFA8}, // 5
		IRGB{0xBFB6}, // 6
		IRGB{0x6F99}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0x8F00}, // 9
		IRGB{0xFF00}, // 10
		IRGB{0xFFB3}, // 11
		IRGB{0x4D74}, // 12
		IRGB{0x7DB0}, // 13
		IRGB{0xF6FF}, // 14
		IRGB{0xFFFF}, // 15
	}, {
		IRGB{0x0},    // Blue
		IRGB{0x0},    // 1
		IRGB{0x6944}, // 2
		IRGB{0x9B74}, // 3
		IRGB{0xBF96}, // 4
		IRGB{0xDFA8}, // 5
		IRGB{0xBFB6}, // 6
		IRGB{0x7F99}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0xA00F}, // 9
		IRGB{0xF24F}, // 10
		IRGB{0xD3BF}, // 11
		IRGB{0x4D74}, // 12
		IRGB{0x7DB0}, // 13
		IRGB{0xF6FF}, // 14
		IRGB{0xFFFF}, // 15
	}, {
		IRGB{0x0},    // Yellow
		IRGB{0x0},    // 1
		IRGB{0x6944}, // 2
		IRGB{0x9B74}, // 3
		IRGB{0xBF96}, // 4
		IRGB{0xDFA8}, // 5
		IRGB{0x9FC5}, // 6
		IRGB{0x4FA8}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0xBE90}, // 9
		IRGB{0xFFF0}, // 10
		IRGB{0xCFFB}, // 11
		IRGB{0x4D74}, // 12
		IRGB{0x7DB0}, // 13
		IRGB{0xF6FF}, // 14
		IRGB{0xFFFF}, // 15
	}, {
		IRGB{0x0},    // Green
		IRGB{0x0},    // 1
		IRGB{0x6944}, // 2
		IRGB{0x9B74}, // 3
		IRGB{0xBD96}, // 4
		IRGB{0xDFA8}, // 5
		IRGB{0xBFB6}, // 6
		IRGB{0x44F9}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0x90F0}, // 9
		IRGB{0xD3F0}, // 10
		IRGB{0xFBFB}, // 11
		IRGB{0x4D74}, // 12
		IRGB{0x7DB0}, // 13
		IRGB{0xF6FF}, // 14
		IRGB{0xFFFF}, // 15
	}}

// Valkyrie
var valkyriePalettes = [][]color.Color{
	{
		IRGB{0x0},    // Red
		IRGB{0x0},    // 1
		IRGB{0x6F93}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xCF83}, // 4
		IRGB{0xFF66}, // 5
		IRGB{0x9FB0}, // 6
		IRGB{0x8FBC}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0x8F22}, // 9
		IRGB{0xAF22}, // 10
		IRGB{0xFF22}, // 11
		IRGB{0x7F00}, // 12
		IRGB{0xEFB0}, // 13
		IRGB{0xD111}, // 14
		IRGB{0xFFFF}, // 15
	}, {
		IRGB{0x0},    // Blue
		IRGB{0x0},    // 1
		IRGB{0x6F93}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xCF83}, // 4
		IRGB{0xFF66}, // 5
		IRGB{0x9FB0}, // 6
		IRGB{0x8FBC}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0xF25F}, // 9
		IRGB{0xF48F}, // 10
		IRGB{0xF8BF}, // 11
		IRGB{0x700F}, // 12
		IRGB{0xEFB0}, // 13
		IRGB{0xD111}, // 14
		IRGB{0xFFFF}, // 15
	}, {
		IRGB{0x0},    // Yellow
		IRGB{0x0},    // 1
		IRGB{0x6F93}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xCF83}, // 4
		IRGB{0xFF66}, // 5
		IRGB{0x9F90}, // 6
		IRGB{0x8FBC}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0x7FF2}, // 9
		IRGB{0xAFF2}, // 10
		IRGB{0xEFF2}, // 11
		IRGB{0x7FF0}, // 12
		IRGB{0xEF90}, // 13
		IRGB{0xD111}, // 14
		IRGB{0xFFFF}, // 15
	}, {
		IRGB{0x0},    // Green
		IRGB{0x0},    // 1
		IRGB{0x6F93}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xCF83}, // 4
		IRGB{0xFF66}, // 5
		IRGB{0x9FB0}, // 6
		IRGB{0x8FBC}, // 7
		IRGB{0xFF9A}, // 8
		IRGB{0x60F0}, // 9
		IRGB{0x90F0}, // 10
		IRGB{0xF0F0}, // 11
		IRGB{0x70F0}, // 12
		IRGB{0xEFB0}, // 13
		IRGB{0xD111}, // 14
		IRGB{0xFFFF}, // 15
	}}

// Wizard
var wizardPalettes = [][]color.Color{
	{
		IRGB{0x0},    // Red
		IRGB{0x0},    // 1
		IRGB{0x6F93}, // 2
		IRGB{0x4F00}, // 3
		IRGB{0xDF86}, // 4
		IRGB{0xFF9A}, // 5
		IRGB{0x6F00}, // 6
		IRGB{0x9F00}, // 7
		IRGB{0xFF00}, // 8
		IRGB{0x8FFF}, // 9
		IRGB{0xDFFF}, // 10
		IRGB{0x0},    // 11
		IRGB{0x0},    // 12
		IRGB{0x0},    // 13
		IRGB{0x7F43}, // 14
		IRGB{0x0},
	}, {
		IRGB{0x0},    // Blue
		IRGB{0x0},    // 1
		IRGB{0x6F93}, // 2
		IRGB{0x222F}, // 3
		IRGB{0xDF86}, // 4
		IRGB{0xFF9A}, // 5
		IRGB{0x622F}, // 6
		IRGB{0xA22F}, // 7
		IRGB{0xF22F}, // 8
		IRGB{0x8FFF}, // 9
		IRGB{0xDFFF}, // 10
		IRGB{0x0},    // 11
		IRGB{0x0},    // 12
		IRGB{0x0},    // 13
		IRGB{0x7F43}, // 14
		IRGB{0x0},
	}, {
		IRGB{0x0},    // Yellow
		IRGB{0x0},    // 1
		IRGB{0x6F93}, // 2
		IRGB{0x7F84}, // 3
		IRGB{0xDF86}, // 4
		IRGB{0xFF9A}, // 5
		IRGB{0x9FD0}, // 6
		IRGB{0xCFC0}, // 7
		IRGB{0xFFD0}, // 8
		IRGB{0x8FFF}, // 9
		IRGB{0xDFFF}, // 10
		IRGB{0x0},    // 11
		IRGB{0x0},    // 12
		IRGB{0x0},    // 13
		IRGB{0x7F43}, // 14
		IRGB{0x0},
	}, {
		IRGB{0x0},    // Green
		IRGB{0x0},    // 1
		IRGB{0x6F93}, // 2
		IRGB{0x21F1}, // 3
		IRGB{0xDF86}, // 4
		IRGB{0xFF9A}, // 5
		IRGB{0x42F2}, // 6
		IRGB{0x82F2}, // 7
		IRGB{0xA2F2}, // 8
		IRGB{0x8FFF}, // 9
		IRGB{0xDFFF}, // 10
		IRGB{0x0},    // 11
		IRGB{0x0},    // 12
		IRGB{0x0},    // 13
		IRGB{0x7F43}, // 14
		IRGB{0x0},
	}}

// Elf
var elfPalettes = [][]color.Color{
	{
		IRGB{0x0},    // Red
		IRGB{0x0},    // 1
		IRGB{0x4F93}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xDF86}, // 4
		IRGB{0xFF8A}, // 5
		IRGB{0x9FD9}, // 6
		IRGB{0xCFC8}, // 7
		IRGB{0xFFF6}, // 8
		IRGB{0xDFFF}, // 9
		IRGB{0xFFFF}, // 10
		IRGB{0x5F01}, // 11
		IRGB{0x9F00}, // 12
		IRGB{0xDF32}, // 13
		IRGB{0x6F10}, // 14
		IRGB{0x4A8F}, // 15
	}, {
		IRGB{0x0},    // Blue
		IRGB{0x0},    // 1
		IRGB{0x4F93}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xDF86}, // 4
		IRGB{0xFF8A}, // 5
		IRGB{0x9FD9}, // 6
		IRGB{0xCFC8}, // 7
		IRGB{0xFFF6}, // 8
		IRGB{0xDFFF}, // 9
		IRGB{0xFFFF}, // 10
		IRGB{0x801F}, // 11
		IRGB{0xF11D}, // 12
		IRGB{0xC06F}, // 13
		IRGB{0x6F10}, // 14
		IRGB{0x4A8F}, // 15
	}, {
		IRGB{0x0},    // Yellow
		IRGB{0x0},    // 1
		IRGB{0x4F93}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xDF86}, // 4
		IRGB{0xFF8A}, // 5
		IRGB{0x9FD9}, // 6
		IRGB{0xCFC8}, // 7
		IRGB{0xFFF6}, // 8
		IRGB{0xDFFF}, // 9
		IRGB{0xFFFF}, // 10
		IRGB{0x9FC0}, // 11
		IRGB{0xCFD0}, // 12
		IRGB{0xFFF0}, // 13
		IRGB{0x6F10}, // 14
		IRGB{0x4A8F}, // 15
	}, {
		IRGB{0x0},    // Green
		IRGB{0x0},    // 1
		IRGB{0x4F93}, // 2
		IRGB{0x9F84}, // 3
		IRGB{0xDF86}, // 4
		IRGB{0xFF8A}, // 5
		IRGB{0x9FD9}, // 6
		IRGB{0xCFC8}, // 7
		IRGB{0xFFF6}, // 8
		IRGB{0xDFFF}, // 9
		IRGB{0xFFFF}, // 10
		IRGB{0x60F1}, // 11
		IRGB{0x90F0}, // 12
		IRGB{0xC0F6}, // 13
		IRGB{0x6F10}, // 14
		IRGB{0x4A8F}, // 15
	},
}

var gauntletPalettes = map[string][][]color.Color{
	"teleff":   teleFfPalettes,
	"floor":    floorPalettes,
	"wall":     wallPalettes,
	"base":     basePalettes,
	"warrior":  warriorPalettes,
	"valkyrie": valkyriePalettes,
	"wizard":   wizardPalettes,
	"elf":      elfPalettes,
}
