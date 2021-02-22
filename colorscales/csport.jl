#!/usr/bin/env juliaplain.dash

using ColorSchemes
for cs in ColorSchemes.colorschemes
    println(first(cs))
    for color in last(cs).colors
        println(color)
    end
end
