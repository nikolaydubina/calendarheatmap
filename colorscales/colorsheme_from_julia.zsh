#!/usr/bin/env zsh

###
## Tested with commit f5d75c93541479f48d2c62f9108e2eac0dafed44 of https://github.com/JuliaGraphics/ColorSchemes.jl
## Personal comments:
## Usage:
## `re 'zshplain.dash ./colorsheme_from_julia.zsh' ~/code/misc/ColorSchemes.jl/data/allcolorschemes.jl ~/code/misc/ColorSchemes.jl/data/cmocean.jl ~/code/misc/ColorSchemes.jl/data/colorbrewerschemes.jl ~/code/misc/ColorSchemes.jl/data/colorcetdata.jl ~/code/misc/ColorSchemes.jl/data/cvd.jl ~/code/misc/ColorSchemes.jl/data/gnu.jl ~/code/misc/ColorSchemes.jl/data/matplotlib.jl ~/code/misc/ColorSchemes.jl/data/sampledcolorschemes.jl ~/code/misc/ColorSchemes.jl/data/scicolor.jl ~/code/misc/ColorSchemes.jl/data/seaborn.jl ~/code/misc/ColorSchemes.jl/data/tableau.jl > s.txt`
###

setopt multios re_match_pcre extendedglob pipefail

local input="$1"
test -e "$input" || return 1
local output="./${input:t:r}.go"

alias maybenot='(( ${csport_mode:-0} == 0 )) || return 0'

function eco() {
    maybenot
    print -r -- "$*" >> $output
}
function eco2() {
    maybenot
    print -r -- "$*"
}
function eco3() {
    (( ${csport_mode:-0} == 1 )) || return 0
    print -r -- "$*"
}
function ecerr() {
    print -r -- "$*" >&2
}

function colorscheme-port() {
    (( ${csport_mode:-0} == 0 )) && test -e "$output" && mv "$output" "${output}.bak"

    eco 'package colorscales'

    local i=0 name
    for line in "${(@f)$(cat $input)}" ; do
        i=$(( i + 1 ))
        # https://discourse.julialang.org/t/how-to-remove-comments-from-a-julia-file/55783
        if [[ "$line" =~ '^\s*#' ]] ; then
            # line="$match[1]"
            continue
        fi
        if [[ "$line" =~ '^\s*loadcolorscheme\(:(\S+)\s*,' ]] ; then
            name="${match[1]}"
            eco "var ${match[1]} = ColorScaleX{"
            eco2 "    case \"${match[1]}\":"$'\n'"        return ${match[1]}"
            eco3 "#+begin_src bsh.dash :exports both :results verbatim file
plot-basic $name
#+end_src

"
        elif [[ "$line" =~ '^\s*(?:Colors\.)?RGB(?:\{Float64\})?\(([^]]*)\)(\]?)' ]] ; then
            eco "RGB(${match[1]}),"
            if [[ "$match[2]" == ']' ]] ; then
                eco '}'
            fi
        elif [[ "$line" =~ '^\s*colorant"([^"]*)"\s*,?\s*(?:colorant"([^"]*)"\s*,?\s*)?(?:#.*)?$' ]] ; then
            eco "Hex(\"${match[1]}\"),"
            if [[ "$match[2]" =~ '\S' ]] ; then
                eco "Hex(\"${match[2]}\"),"
            fi
        elif [[ "$line" =~ '^\s*\]' ]] ; then
            eco "}"
        elif [[ "$line" =~ '^\s*$' ]] ; then
            eco
        else
            ecerr "UNKNOWN LINE at '${input}:${i}':"$'\n'"$line"$'\n'
        fi
    done

    gofmt -w $output
}

colorscheme-port
