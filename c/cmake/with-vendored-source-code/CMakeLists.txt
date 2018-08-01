cmake_minimum_required(VERSION 2.8.9)
project (cmakeexample)

add_executable(main src/main.cpp)
add_subdirectory(lib)
include_directories(lib)

target_link_libraries(main sqlite3 pthread dl)
