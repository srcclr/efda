#include <iostream>
#include <sqlite3.h>

using namespace std;


int main(int argc, char* argv[])
{
    std::cout << "SQLite3 version " << sqlite3_libversion() << std::endl;

    sqlite3* hDB = 0;
    int nRowCount = 0, nColCount = 0;
    char** papszResult = 0;
    sqlite3_open(argv[1], &hDB);
    if( !hDB )
        return 1;
    sqlite3_get_table( hDB, "SELECT * FROM sqlite_master",
        &papszResult, &nRowCount, &nColCount,
    0 );
    sqlite3_free_table(papszResult);
    sqlite3_close(hDB);
    return 0;
}

