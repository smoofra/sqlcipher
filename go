
# build script for macosx 

# I have an appropriate build of libressl here, so I'm just going to use that
p=/AppleInternal/Library/Frameworks/Python.framework/Versions/3.7/

./configure LDFLAGS="-L$p/lib" CPPFLAGS="-I$p/include" CFLAGS='-DSQLITE_HAS_CODEC -DSQLITE_SHM_DIRECTORY=\"/tmp\" -g '
make clean
make

echo export  DYLD_LIBRARY_PATH=/Users/lawrence_danna/sqlcipher/.libs/:$p/lib

