#include"test.h"
char* test_hello(const char* name){
    const char* hello=" -> hello";
    char* result=(char*)malloc(sizeof(char)*(strlen(name))+strlen(hello));
    strcpy(result,name);
    strcat(result,hello);
    return result;
}
