package main

// maze object ids (names matching names in IDA)
const (
	MAZEOBJ_TILE_FLOOR = iota
	MAZEOBJ_TILE_STUN
	MAZEOBJ_WALL_REGULAR
	MAZEOBJ_WALL_MOVABLE
	MAZEOBJ_WALL_SECRET
	MAZEOBJ_WALL_DESTRUCTABLE
	MAZEOBJ_WALL_RANDOM
	MAZEOBJ_WALL_TRAPCYC1
	MAZEOBJ_WALL_TRAPCYC2
	MAZEOBJ_WALL_TRAPCYC3
	MAZEOBJ_TILE_TRAP1
	MAZEOBJ_TILE_TRAP2
	MAZEOBJ_TILE_TRAP3
	MAZEOBJ_DOOR_HORIZ
	MAZEOBJ_DOOR_VERT
	MAZEOBJ_PLAYERSTART
	MAZEOBJ_EXIT
	MAZEOBJ_EXITTO6
	MAZEOBJ_MONST_GHOST
	MAZEOBJ_MONST_GRUNT
	MAZEOBJ_MONST_DEMON
	MAZEOBJ_MONST_LOBBER
	MAZEOBJ_MONST_SORC
	MAZEOBJ_MONST_AUX_GRUNT
	MAZEOBJ_MONST_DEATH
	MAZEOBJ_MONST_ACID
	MAZEOBJ_MONST_SUPERSORC
	MAZEOBJ_MONST_IT
	MAZEOBJ_GEN_GHOST1
	MAZEOBJ_GEN_GHOST2
	MAZEOBJ_GEN_GHOST3
	MAZEOBJ_GEN_GRUNT1
	MAZEOBJ_GEN_GRUNT2
	MAZEOBJ_GEN_GRUNT3
	MAZEOBJ_GEN_DEMON1
	MAZEOBJ_GEN_DEMON2
	MAZEOBJ_GEN_DEMON3
	MAZEOBJ_GEN_LOBBER1
	MAZEOBJ_GEN_LOBBER2
	MAZEOBJ_GEN_LOBBER3
	MAZEOBJ_GEN_SORC1
	MAZEOBJ_GEN_SORC2
	MAZEOBJ_GEN_SORC3
	MAZEOBJ_GEN_AUX_GRUNT1
	MAZEOBJ_GEN_AUX_GRUNT2
	MAZEOBJ_GEN_AUX_GRUNT3
	MAZEOBJ_TREASURE
	MAZEOBJ_TREASURE_LOCKED
	MAZEOBJ_TREASURE_BAG
	MAZEOBJ_FOOD_DESTRUCTABLE
	MAZEOBJ_FOOD_INVULN
	MAZEOBJ_POT_DESTRUCTABLE
	MAZEOBJ_POT_INVULN
	MAZEOBJ_KEY
	MAZEOBJ_POWER_INVIS
	MAZEOBJ_POWER_REPULSE
	MAZEOBJ_POWER_REFLECT
	MAZEOBJ_POWER_TRANSPORT
	MAZEOBJ_POWER_SUPERSHOT
	MAZEOBJ_POWER_INVULN
	MAZEOBJ_MONST_DRAGON
	MAZEOBJ_HIDDENPOT
	MAZEOBJ_TRANSPORTER
	MAZEOBJ_FORCEFIELDHUB
)

// Flags for levels
const (
	LFLAG4_SHOTS_STUN       = 0x01
	LFLAG4_SHOTS_HURT       = 0x02
	LFLAG4_TRAPS_LOCAL      = 0x04
	LFLAG4_TRAPS_RANDOM     = 0x08
	LFLAG4_WRAP_V           = 0x10
	LFLAG4_WRAP_H           = 0x20
	LFLAG4_EXIT_FAKE        = 0x40
	LFLAG4_PLAYER_OFFSCREEN = 0x80

	LFLAG3_RANDOMFOOD_MASK  = 0x0700
	LFLAG3_WALLS_CYCLIC     = 0x0800
	LFLAG3_WALLS_DELETABLE1 = 0x1000
	LFLAG3_WALLS_DELETABLE2 = 0x2000
	LFLAG3_EXIT_MOVES       = 0x4000
	LFLAG3_EXIT_CHOOSEONE   = 0x8000

	LFLAG2_FAST_GHOSTS     = 0x010000
	LFLAG2_FAST_GRUNTS     = 0x020000
	LFLAG2_FAST_DEMONS     = 0x040000
	LFLAG2_FAST_LOBBERS    = 0x080000
	LFLAG2_FAST_SORCERERS  = 0x100000
	LFLAG2_FAST_AUX_GRUNTS = 0x200000
	LFLAG2_FAST_DEATHS     = 0x400000
	LFLAG2_INVIS_ALLWALLS  = 0x800000

	LFLAG1_ODDANGLE_GHOSTS     = 0x01000000
	LFLAG1_ODDANGLE_GRUNTS     = 0x02000000
	LFLAG1_ODDANGLE_DEMONS     = 0x04000000
	LFLAG1_ODDANGLE_LOBBERS    = 0x08000000
	LFLAG1_ODDANGLE_SORCERERS  = 0x10000000
	LFLAG1_ODDANGLE_AUX_GRUNTS = 0x20000000
	LFLAG1_ODDANGLE_DEATHS     = 0x40000000
	LFLAG1_INVIS_TRAPWALLS     = 0x80000000
)

var mazeLflag1strings = map[int]string{
	LFLAG1_ODDANGLE_GHOSTS:     "ODDANGLE_GHOSTS",
	LFLAG1_ODDANGLE_GRUNTS:     "ODDANGLE_GRUNTS",
	LFLAG1_ODDANGLE_DEMONS:     "ODDANGLE_DEMONS",
	LFLAG1_ODDANGLE_LOBBERS:    "ODDANGLE_LOBBERS",
	LFLAG1_ODDANGLE_SORCERERS:  "ODDANGLE_SORCERERS",
	LFLAG1_ODDANGLE_AUX_GRUNTS: "ODDANGLE_GRUNTS",
	LFLAG1_ODDANGLE_DEATHS:     "ODDANGLE_DEATHS",
	LFLAG1_INVIS_TRAPWALLS:     "INVIS_TRAPWALLS",
}
