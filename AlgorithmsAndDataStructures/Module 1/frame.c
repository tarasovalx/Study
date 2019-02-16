#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    if (argc != 4)
        printf("Usage: frame <height> <width> <text>\n");
    else{
        int height = atoi(argv[1]);
        int width = atoi(argv[2]);
        char* s = argv[3];
        if ((height < 3 && s) || width-2 < strlen(s)){
            printf("Error");
        }
        else{
            int y_pos = (height -1)/2;
            int x_pos = (width - strlen(s))/2;
            char header[width +1];
            for (int i = 0; i < width; i++) header[i] = '*';
            char mid[width +1];
            for (int i = 0; i < width; i++) mid[i] = ' ';
            char *label = malloc(sizeof(char)* (width +1));
            for (int i = 1; i < width; i++){
                if (i == x_pos){
                    strncpy(label + i, s, width - i);
                    i += strnlen(s, width) -1;
                }
                else
                    label[i] = ' ';
            }
            mid[0] = '*';
            mid[width-1] = '*';
            mid[width] = '\0';
            label[0] = '*';
            label[width-1] = '*';
            label[width] = '\0';
            header[width] = '\0';

            for (int i = 0; i< height; i++){
                if (i == 0 || i == height-1)
                    printf("%s\n", header);
                else if (i == y_pos)
                    printf("%s\n", label);
                else
                    printf("%s\n", mid);
            }
            free(label);
        }
    }
}
