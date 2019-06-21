# sway-cwd

> WIP! This was literally written in 15min to scratch my own itch. I'm sure it
> can be improved, use at own risk. It needs tests too.

This is a simple tool to get the current working directory in
[sway](https://swaywm.org/). The idea is that the current working dir can be
passed to a new terminal instance, so it opens in the same directory.

For example, is using [kitty](https://sw.kovidgoyal.net/kitty/), add this to
`~/.config/sway/config`:

```
# Start a terminal in same directory
bindsym $mod+Return exec kitty -d "$(sway-cwd)"
```

Note that `sway-cmd` must be on the `$PATH`, which won't include any additions
set by your shell. You could place the binary in `/usr/local/bin` or include
the full path:

```
bindsym $mod+Return exec kitty -d "$(/home/akupila/go/bin/sway-cwd)"
```

The inspiration came from
[here](https://github.com/swaywm/sway/issues/1973#issuecomment-419504059).
Unlike the shell script, this the `sway-cwd` binary is statically linked and
thus has no external dependencies, such as `jq`.
