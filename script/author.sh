#!/usr/bin/env bash

# Generate AUTHORS file in build directory

echo "Here is an inevitably incomplete list of MUCH-APPRECIATED CONTRIBUTORS --
people who have submitted patches, reported bugs, added translations, helped
answer newbie questions, and generally made 'go-xn' that much better:
" > ./build/AUTHORS

git log --format='%aN <%aE>' \
    | sort -u \
    | sed -e '/dependabot/d' \
    >> ./build/AUTHORS
