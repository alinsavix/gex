As part of a larger Gauntlet (II) reverse engineering project, I've often
wanted a decent tile/stamp extractor, but such a thing does not (AFAIK!)
actually exist publicly. I wanted one, so I made one.

This is also my first golang program (seemed like a good excuse to learn),
so it's probably double champion edition gold platnum++ lame. Totally
open to suggestions on improvements, though. Especially once the code is
more than an empty husk of its future self.

Right now the command line interface is a bit clunky. I admit that I haven't
been able to think of a *good* way to specify what to extract from amongst
all the different things in the game.

Some commands you could try:
* gex maze0
* gex item-key
* gex ghost-walk-up

Current limitations:
* You can't specify arbitrary tile numbers to render
* Rendering animations doesn't work (most of the code is there, but I switched to png, which doesn't support animations)
* Anims also don't work because I haven't loaded the data required for it. I really should read it from the ROMs, though.
* You can't tell what monsters a given generator creates when dumping a maze
* Some of the code is still quite crap
