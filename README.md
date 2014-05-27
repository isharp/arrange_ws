# arrange_ws.go

An [i3wm](http://i3wm.org) workspace sorter that uses the i3ipc Go library.

## About
There is a library for i3's interprocess communication interface written in Go that this script depends on. You can find it here: [i3ipc](https://github.com/proxypoke/i3ipc).
To use the script, simply start it when i3 starts by putting something like this in your `~/.i3/config` file:

``exec --no-startup-id "nohup /home/username/bin/arrange_ws >/dev/null 2>&1 &"``

## License
This program is free software under the terms of the WTFPL. It comes without any warranty, to the extent permitted by applicable law. For a copy of the license, see COPYING or head to [http://www.wtfpl.net/txt/copying](http://www.wtfpl.net/txt/copying/).
