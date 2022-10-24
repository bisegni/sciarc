cmake_minimum_required(VERSION 3.2)
project(dcomp)
IF(CMAKE_INSTALL_PREFIX_INITIALIZED_TO_DEFAULT)
  SET(CMAKE_INSTALL_PREFIX ${CMAKE_CURRENT_BINARY_DIR}/local CACHE PATH "default install directory" FORCE)
ENDIF(CMAKE_INSTALL_PREFIX_INITIALIZED_TO_DEFAULT)

include(ExternalProject)

set(CMAKE_CXX_STANDARD 17)
set(BOOST_VERSION 1.79.0)

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/depinstall/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/depinstall/lib)
set(CMAKE_BINARY_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/depinstall/bin)
option(BUILD_SHARED_LIBS "Build using shared libraries" off)

message (STATUS "Setup Boost ${BOOST_VERSION}.")
if (WIN32)
     set (BOOST_BOOTSTRAP_CMD ${BOOST_GIT_REPO_DIR}/bootstrap.bat)
     # It was bjam at one point; it is now b2
    set (BOOST_BUILD_CMD ${BOOST_GIT_REPO_DIR}/b2.exe)
elseif (UNIX)
    set (BOOST_BOOTSTRAP_CMD ./${BOOST_GIT_REPO_DIR}/bootstrap.sh)
    set (BOOST_BUILD_CMD ./${BOOST_GIT_REPO_DIR}/b2)
endif ()

ExternalProject_Add (
    boost
    PREFIX ${CMAKE_CURRENT_BINARY_DIR}/boost
    GIT_REPOSITORY https://github.com/boostorg/boost.git
    GIT_TAG boost-${BOOST_VERSION}
    GIT_PROGRESS 1
    BUILD_IN_SOURCE 1
    UPDATE_COMMAND ${BOOST_BOOTSTRAP_CMD}
    CONFIGURE_COMMAND ""
    BUILD_COMMAND ${BOOST_BUILD_CMD} toolset=clang cxxflags=-stdlib=libc++ linkflags=-stdlib=libc++ --with-system --with-filesystem --with-thread --with-regex --with-iostreams link=static install --prefix=${CMAKE_INSTALL_PREFIX}
    INSTALL_COMMAND ""
    LOG_DOWNLOAD 1
    LOG_UPDATE 1
    LOG_CONFIGURE 1
    LOG_BUILD 1
    LOG_INSTALL 1
    )

SET(BUILD_LIBRARY ${BUILD_LIBRARY} 
    boost_system
    boost_thread
    boost_filesystem
    boost_regex
    boost_iostreams)
ADD_DEPENDENCIES(${PROJECT_NAME} boost)
ADD_DEPENDENCIES(unit_test boost)

message (STATUS "Setup GTest.")
ExternalProject_Add(
     googletest
     GIT_REPOSITORY https://github.com/google/googletest.git
	 GIT_TAG release-1.8.0
     PREFIX "${CMAKE_CURRENT_BINARY_DIR}/gtest"
     CMAKE_ARGS
      -DCMAKE_INSTALL_PREFIX:PATH=${CMAKE_INSTALL_PREFIX}
     LOG_DOWNLOAD ON
     LOG_CONFIGURE ON
     LOG_BUILD ON
)
ADD_DEPENDENCIES(unit_test googletest)

IF(NOT EXISTS ${CMAKE_CURRENT_BINARY_DIR}/cling)
    MESSAGE (STATUS "Download cling")
    MESSAGE(STATUS "Create directory for cling download files")
    EXECUTE_PROCESS(
    COMMAND mkdir -p ${CMAKE_CURRENT_BINARY_DIR}/cling;
    RESULT_VARIABLE err
    OUTPUT_VARIABLE out
    )

    MESSAGE(STATUS "Executing ${PROJECT_SOURCE_DIR}/config/cling_clone.sh")
    EXECUTE_PROCESS(
    COMMAND ${PROJECT_SOURCE_DIR}/tools/cling_clone.sh ${CMAKE_CURRENT_BINARY_DIR}/cling;
    WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}/cling
    RESULT_VARIABLE err
    OUTPUT_VARIABLE out
    )
    IF(err)
  	    MESSAGE(FATAL_ERROR "CANNOT create ${CMAKE_CURRENT_BINARY_DIR}/cling: ${err}")
    ENDIF()
ENDIF()

MESSAGE(STATUS "setup cling")
ExternalProject_Add(
    cling
    DOWNLOAD_COMMAND ""
    PREFIX "${CMAKE_BINARY_DIR}/cling"
    SOURCE_DIR "${CMAKE_BINARY_DIR}/cling/src"
    CMAKE_ARGS -DCMAKE_INSTALL_PREFIX:PATH=${CMAKE_INSTALL_PREFIX}
    LOG_DOWNLOAD ON
    LOG_UPDATE ON
    LOG_PATCH ON
    LOG_INSTALL ON
    LOG_CONFIGURE ON
    LOG_BUILD ON
)
ADD_DEPENDENCIES(${PROJECT_NAME} cling)
ADD_DEPENDENCIES(unit_test cling)
SET(BUILD_LIBRARY ${BUILD_LIBRARY}
                    clingInterpreter)
include_directories(${CMAKE_INSTALL_PREFIX}/include)
link_directories(${CMAKE_INSTALL_PREFIX}/lib)

set(SOURCE_FILES
    engine/src/dcomp.cpp)

set(TEST_SOURCE_FILES
    engine/test/test.cpp)

ADD_LIBRARY(
    ${PROJECT_NAME} STATIC
    ${SOURCE_FILES}
    )

ADD_EXECUTABLE(
    unit_test
    ${TEST_SOURCE_FILES}
)

TARGET_LINK_LIBRARIES(${PROJECT_NAME} PRIVATE ${BUILD_LIBRARY})
TARGET_LINK_LIBRARIES(unit_test ${BUILD_LIBRARY} gtest)


MESSAGE(STATUS "Install prefix: ${CMAKE_INSTALL_PREFIX}")
# enable_testing()
# add_test(NAME    unit_test 
#          COMMAND unit_test)