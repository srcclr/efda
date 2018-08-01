cmake_minimum_required(VERSION 2.8.9)
project (cmakeexample)

set(CMAKE_MODULE_PATH ${CMAKE_SOURCE_DIR}/cmake_modules)
find_package(SQLite3 3.22 REQUIRED)
set(${SQLITE3_INCLUDE_DIRS})

add_executable(main src/main.cpp)
target_link_libraries(main ${SQLITE3_LIBRARIES})
