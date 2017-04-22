This is the output of a quick & dirty hack to make the G2 maze dumper
also dump G1 mazes. Source not available yet because of the sheer
ugliness of the hack.

Known issues:
  - some colors/textures may be off. Not 100% positive.
  - g1 lets mazes have level 1/2/3 monsters, but they all show as level 3
  - special potions (extra armor, etc) don't show
  - a couple of the mazes are obviously decoded incorrectly
  - probably other things -- it's a *really* ugly hack

If someone is bored and looks through the mazes and figures out what
specifically is decoded wrong on some of them, open an issue on github and
I'll try to fix the bad ones. If you're using MAME and get to one of the
mazes that decodes poorly, you can hit f4 and then enter twice to show
the playfield straight from playfield memory, and screenshot that to show
what the maze (the walls, at least) is supposed to look like.
