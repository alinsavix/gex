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
	LFLAG1_ODDANGLE_GHOSTS     = 0x01000000
	LFLAG1_ODDANGLE_GRUNTS     = 0x02000000
	LFLAG1_ODDANGLE_DEMONS     = 0x04000000
	LFLAG1_ODDANGLE_LOBBERS    = 0x08000000
	LFLAG1_ODDANGLE_SORCERERS  = 0x10000000
	LFLAG1_ODDANGLE_AUX_GRUNTS = 0x20000000
	LFLAG1_ODDANGLE_DEATHS     = 0x40000000
	LFLAG1_INVIS_TRAPWALLS     = 0x80000000

	LFLAG2_FAST_GHOSTS     = 0x010000
	LFLAG2_FAST_GRUNTS     = 0x020000
	LFLAG2_FAST_DEMONS     = 0x040000
	LFLAG2_FAST_LOBBERS    = 0x080000
	LFLAG2_FAST_SORCERERS  = 0x100000
	LFLAG2_FAST_AUX_GRUNTS = 0x200000
	LFLAG2_FAST_DEATHS     = 0x400000
	LFLAG2_INVIS_ALLWALLS  = 0x800000

	LFLAG3_RANDOMFOOD_MASK  = 0x0700
	LFLAG3_WALLS_CYCLIC     = 0x0800
	LFLAG3_WALLS_DELETABLE1 = 0x1000
	LFLAG3_WALLS_DELETABLE2 = 0x2000
	LFLAG3_EXIT_MOVES       = 0x4000
	LFLAG3_EXIT_CHOOSEONE   = 0x8000

	LFLAG4_SHOTS_STUN       = 0x01
	LFLAG4_SHOTS_HURT       = 0x02
	LFLAG4_TRAPS_LOCAL      = 0x04
	LFLAG4_TRAPS_RANDOM     = 0x08
	LFLAG4_WRAP_V           = 0x10
	LFLAG4_WRAP_H           = 0x20
	LFLAG4_EXIT_FAKE        = 0x40
	LFLAG4_PLAYER_OFFSCREEN = 0x80
)

var mazeFlagStrings = map[int]string{
	LFLAG1_ODDANGLE_GHOSTS:     "ODDANGLE_GHOSTS",
	LFLAG1_ODDANGLE_GRUNTS:     "ODDANGLE_GRUNTS",
	LFLAG1_ODDANGLE_DEMONS:     "ODDANGLE_DEMONS",
	LFLAG1_ODDANGLE_LOBBERS:    "ODDANGLE_LOBBERS",
	LFLAG1_ODDANGLE_SORCERERS:  "ODDANGLE_SORCERERS",
	LFLAG1_ODDANGLE_AUX_GRUNTS: "ODDANGLE_GRUNTS",
	LFLAG1_ODDANGLE_DEATHS:     "ODDANGLE_DEATHS",
	LFLAG1_INVIS_TRAPWALLS:     "INVIS_TRAPWALLS",

	LFLAG2_FAST_GHOSTS:     "FAST_GHOSTS",
	LFLAG2_FAST_GRUNTS:     "FAST_GRUNTS",
	LFLAG2_FAST_DEMONS:     "FAST_DEMONS",
	LFLAG2_FAST_LOBBERS:    "FAST_LOBBERS",
	LFLAG2_FAST_SORCERERS:  "FAST_SORCERERS",
	LFLAG2_FAST_AUX_GRUNTS: "FAST_AUX_GRUNTS",
	LFLAG2_FAST_DEATHS:     "FAST_DEATHS",
	LFLAG2_INVIS_ALLWALLS:  "INVIS_ALLWALLS",

	LFLAG3_WALLS_CYCLIC:     "WALLS_CYCLIC",
	LFLAG3_WALLS_DELETABLE1: "WALLS_DELETABLE1",
	LFLAG3_WALLS_DELETABLE2: "WALLS_DELETABLE2",
	LFLAG3_EXIT_MOVES:       "EXIT_MOVES",
	LFLAG3_EXIT_CHOOSEONE:   "EXIT_CHOOSE_ONE",

	LFLAG4_SHOTS_STUN:       "SHOTS_STUN",
	LFLAG4_SHOTS_HURT:       "SHOTS_HURT",
	LFLAG4_TRAPS_LOCAL:      "TRAPS_LOCAL",
	LFLAG4_TRAPS_RANDOM:     "TRAPS_RANDOM",
	LFLAG4_WRAP_V:           "WRAP_V",
	LFLAG4_WRAP_H:           "WRAP_H",
	LFLAG4_EXIT_FAKE:        "EXIT_FAKE",
	LFLAG4_PLAYER_OFFSCREEN: "PLAYER_OFFSCREEN",
}

const (
	TRICK_NONE           = 0x00
	TRICK_TRANSPORT1     = 0x01
	TRICK_TRANSPORT2     = 0x02
	TRICK_TRANSPORT3     = 0x03
	TRICK_TRANSPORT4     = 0x04
	TRICK_WATCHSHOOT1    = 0x05
	TRICK_WATCHSHOOT2    = 0x06
	TRICK_SAVESUPERSHOTS = 0x07
	TRICK_NOUSEINVUL     = 0x08
	TRICK_NOGETHIT       = 0x09
	TRICK_PUSHWALL       = 0x0a
	TRICK_NOFOOLED       = 0x0b
	TRICK_NOGREEDY1      = 0x0c
	TRICK_NOGREEDY2      = 0x0d
	TRICK_DIET           = 0x0e
	TRICK_BEPUSHY        = 0x0f
	TRICK_IT             = 0x10
	TRICK_NOHURTFRIENDS  = 0x11
)

var mazeSecretStrings = map[int]string{
	TRICK_NONE:           "No trick",
	TRICK_TRANSPORT1:     "Try Transportability (onto demon)",
	TRICK_TRANSPORT2:     "Try Transportability (onto death)",
	TRICK_TRANSPORT3:     "Try Transportability (into exit)",
	TRICK_TRANSPORT4:     "Try Transportability (onto secret wall)",
	TRICK_WATCHSHOOT1:    "Watch What You Shoot (shoot foods)",
	TRICK_WATCHSHOOT2:    "Watch What You Shoot (shoot secret walls)",
	TRICK_SAVESUPERSHOTS: "Save Super Shots",
	TRICK_NOUSEINVUL:     "Don't Use Invulnerability",
	TRICK_NOGETHIT:       "Don't Get Hit (while killing a dragon)",
	TRICK_PUSHWALL:       "Try Pushing a Wall",
	TRICK_NOFOOLED:       "Don't Be Fooled",
	TRICK_NOGREEDY1:      "Don't Be Greedy (no keys or potions)",
	TRICK_NOGREEDY2:      "Don't Be Greedy (no treasure)",
	TRICK_DIET:           "Go On a Diet (no food)",
	TRICK_BEPUSHY:        "Be Pushy",
	TRICK_IT:             "IT Could Be Nice",
	TRICK_NOHURTFRIENDS:  "Don't Hurt Friends",
}
