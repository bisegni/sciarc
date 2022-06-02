cmake_minimum_required(VERSION 3.2)
project(dcomp)

set(CMAKE_CXX_STANDARD 14)
 include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)
 conan_basic_setup()

set(SOURCE_FILES
    src/main.cpp)

ADD_EXECUTABLE(${PROJECT_NAME} ${SOURCE_FILES})
target_link_libraries(${PROJECT_NAME} ${CONAN_LIBS})
