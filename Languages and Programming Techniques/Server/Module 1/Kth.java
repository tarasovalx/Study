import java.util.Scanner;

public class Kth {

    private static long powerOf10(long x) {
        return x == 0 ? 1 : 10 * powerOf10(x - 1);
    }

    private static long getCap(long x) {
        long added, skipped = 0, cap = 0;
        while (true) {
            added = 9 * powerOf10(cap) * (cap + 1);
            if (skipped + added > x) {
                return cap + 1;
            }
            skipped += added;
            cap++;
        }
    }

    private static long getNumber(long cap, long number) {
        return powerOf10(cap - 1) + number;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        long k = scanner.nextLong();
        long cap = getCap(k);

        for (long j = cap; j > 1; --j) {
            k -= 9 * powerOf10(j - 2) * (j - 1);
        }

        long number = getNumber(cap, k / cap);

        for (long j = 0; j < cap - k % cap - 1; ++j) {
            number /= 10;
        }

        System.out.println(number % 10);
    }
}
