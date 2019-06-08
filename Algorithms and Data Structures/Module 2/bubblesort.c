void bubblesort(unsigned long nel, int(*compare)(unsigned long i, unsigned long j), void(*swap) (unsigned long i, unsigned long j)){
    unsigned long t = nel - 1;
    unsigned long bound_r = t, bound_l = 0;
    int cnt = 0;
    while (t > 0){
        t = 0;
        if (cnt){
            for (int i = bound_r; i > bound_l; i--) {
                if (compare(i - 1, i) == 1) {
                    swap(i - 1, i);
                    t = i;
                    cnt = 0;

                }
            }
            bound_l = t;
        }
        else {
            for (int i = bound_l; i < bound_r; i++) {
                if (compare(i, i + 1) == 1) {
                    swap(i, i + 1);
                    t = i;
                    cnt = 1;

                }
            }
            bound_r = t;
        }
    }
}
