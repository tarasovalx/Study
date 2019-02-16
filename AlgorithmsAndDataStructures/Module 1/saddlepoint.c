#include <stdio.h>

int main() {
    int n = 0, m = 0;
    scanf("%d%d", &n, &m);
    int mat[n][m];
    int max_row[n];
    int min_col[m];

    for(int i =0; i < n; i++){
        for (int j=0;j<m; j++) scanf("%d", &mat[i][j]);
    }

    for (int i =0; i < n; i++){
        max_row[i] = mat[i][0];
        for (int j =0; j<m; j++){
            if (mat[i][j] > max_row[i]) max_row[i] = mat[i][j];
        }
    }
    for (int i = 0; i < m; i++){
        min_col[i] = mat[0][i];
        for (int j = 0;j < n; j++){
            if (mat[j][i] < min_col[i]) min_col[i] = mat[j][i];
        }
    }
    for (int i = 0; i < n; i++){
        for (int j = 0; j < m;j++){
            if (max_row[i] == min_col[j]){
                printf("%d %d", i, j);
                return 0;
            }
        }
    }
    printf("%s", "none");
    return 0;
}
