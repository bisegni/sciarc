cmake_minimum_required(VERSION 3.2)
project(sciarc)
IF(CMAKE_INSTALL_PREFIX_INITIALIZED_TO_DEFAULT)
  SET(CMAKE_INSTALL_PREFIX ${CMAKE_CURRENT_BINARY_DIR}/local CACHE PATH "default install directory" FORCE)
ENDIF(CMAKE_INSTALL_PREFIX_INITIALIZED_TO_DEFAULT)

include(ExternalProject)

set(CMAKE_CXX_STANDARD 17)
set(BOOST_VERSION 1.77.0)

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_INSTALL_PREFIX}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_INSTALL_PREFIX}/lib)
set(CMAKE_BINARY_OUTPUT_DIRECTORY ${CMAKE_INSTALL_PREFIX}/bin)
option(BUILD_SHARED_LIBS "Build using shared libraries" off)

set(SOURCE_FILES
    engine/src/sciarc.cpp
    engine/src/epics/epics.cpp
    )

ADD_LIBRARY(
    ${PROJECT_NAME} STATIC
    ${SOURCE_FILES}
    )

SET_TARGET_PROPERTIES(${PROJECT_NAME} 
                        PROPERTIES PUBLIC_HEADER 
                        engine/src/sciarc.h
                        )
install(TARGETS ${PROJECT_NAME}
        PUBLIC_HEADER
        DESTINATION ${CMAKE_INSTALL_PREFIX}/include/sciarc
)
install(TARGETS ${PROJECT_NAME}
        #RUNTIME DESTINATION bin
        LIBRARY DESTINATION lib
        #ARCHIVE DESTINATION lib/static
        )

TARGET_INCLUDE_DIRECTORIES(${PROJECT_NAME} PUBLIC engine/src)
TARGET_INCLUDE_DIRECTORIES(${PROJECT_NAME} PUBLIC ${CMAKE_INSTALL_PREFIX}/include)
TARGET_LINK_DIRECTORIES(${PROJECT_NAME} PUBLIC ${CMAKE_INSTALL_PREFIX}/lib)

set(TEST_SOURCE_FILES
    engine/test/test.cpp
    engine/test/SciarcTest.cpp
    )

ADD_EXECUTABLE(
    unit_test
    ${TEST_SOURCE_FILES}
)
install(TARGETS unit_test
        RUNTIME DESTINATION bin
        #LIBRARY DESTINATION lib
        #ARCHIVE DESTINATION lib/static
        )

TARGET_INCLUDE_DIRECTORIES(unit_test PUBLIC ${CMAKE_INSTALL_PREFIX}/include)
TARGET_LINK_DIRECTORIES(unit_test PUBLIC ${CMAKE_INSTALL_PREFIX}/lib)
TARGET_LINK_DIRECTORIES(unit_test PUBLIC ${CMAKE_INSTALL_PREFIX}/lib/linux-x86_64)

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
    GIT_SUBMODULES_RECURSE true
    UPDATE_COMMAND ${BOOST_BOOTSTRAP_CMD}
    BUILD_COMMAND ${BOOST_BUILD_CMD} --with-system --with-filesystem --with-thread --with-regex --with-iostreams link=static install --prefix=${CMAKE_INSTALL_PREFIX}
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
#--------EPICS------
ExternalProject_Add(
 	epics
    PREFIX ${CMAKE_CURRENT_BINARY_DIR}/epics
 	GIT_REPOSITORY https://github.com/epics-base/epics-base.git
    GIT_TAG origin/7.0
    GIT_PROGRESS 1
    BUILD_IN_SOURCE 1
    PATCH_COMMAND ${PROJECT_SOURCE_DIR}/epics_patch.sh ${CMAKE_CURRENT_BINARY_DIR}
    CONFIGURE_COMMAND ""
	UPDATE_COMMAND    ""
	INSTALL_COMMAND   ""
    BUILD_COMMAND make -j 3  INSTALL_LOCATION=${CMAKE_INSTALL_PREFIX} install
  )


SET(BUILD_LIBRARY ${BUILD_LIBRARY} )
ADD_DEPENDENCIES(${PROJECT_NAME} epics)
ADD_DEPENDENCIES(unit_test epics)
TARGET_INCLUDE_DIRECTORIES(${PROJECT_NAME} PUBLIC ${CMAKE_INSTALL_PREFIX}/include/os/${CMAKE_SYSTEM_NAME})
TARGET_INCLUDE_DIRECTORIES(${PROJECT_NAME} PUBLIC ${CMAKE_INSTALL_PREFIX}/include/compiler/gcc)
TARGET_INCLUDE_DIRECTORIES(${PROJECT_NAME} PUBLIC ${CMAKE_INSTALL_PREFIX}/include/pv)
TARGET_INCLUDE_DIRECTORIES(${PROJECT_NAME} PUBLIC ${CMAKE_INSTALL_PREFIX}/include/pva)

#--------Google Test------
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

TARGET_LINK_LIBRARIES(${PROJECT_NAME} PUBLIC ${BUILD_LIBRARY} pvaClient pvAccess pthread)
TARGET_LINK_LIBRARIES(unit_test PUBLIC ${BUILD_LIBRARY} gtest sciarc )

MESSAGE(STATUS "Install prefix: ${CMAKE_INSTALL_PREFIX}")
# enable_testing()
# add_test(NAME    unit_test 
#          COMMAND unit_test)