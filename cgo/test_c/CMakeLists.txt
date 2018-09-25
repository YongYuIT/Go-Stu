cmake_minimum_required(VERSION 2.8)
project(test_c)

add_library( # Sets the name of the library.
    test
    # Sets the library as a shared library.
    SHARED
    # Provides a relative path to your source file(s).
    ./test.c
    )


add_executable(${PROJECT_NAME} "main.c")
target_link_libraries( # Specifies the target library.
    ${PROJECT_NAME}
    test)
