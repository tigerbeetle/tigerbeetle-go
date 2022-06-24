#!/usr/bin/env sh
ZIG_EXE=./tigerbeetle/zig/zig.exe
echo "#!/usr/bin/env sh\n $ZIG_EXE cc \$@" > zigcc.sh
CC=$(pwd)/zigcc.sh
go $@