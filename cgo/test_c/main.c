#include"test.h"

int main()
{
    char* str=test_hello("yuyong");
    printf("%s\n",str);
    free(str);
}
