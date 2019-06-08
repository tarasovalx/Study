import java.math.BigInteger;
import java.util.Scanner;

public class FastFib {

    private static BigInteger[][] mult2x2Matrixes(BigInteger[][] a, BigInteger[][] b){
        BigInteger[][] res = new BigInteger[2][2];
        res[0][0] = a[0][0].multiply(b[0][0]).add(a[0][1].multiply(b[1][0]));
        res[0][1] = a[0][0].multiply(b[0][1]).add(a[0][1].multiply(b[1][1]));
        res[1][0] = a[1][0].multiply(b[0][0]).add(a[1][1].multiply(b[1][0]));
        res[1][1] = a[1][0].multiply(b[0][1]).add(a[1][1].multiply(b[1][1]));
        return res;
    }

    private static BigInteger[][] FastPow2x2Matrixes(BigInteger[][] a, int n){
        BigInteger[][] res = {{BigInteger.ONE, BigInteger.ZERO},
                              {BigInteger.ZERO, BigInteger.ONE}};

        while (n != 0){
            if ((n&1) != 0){
                res = mult2x2Matrixes(res, a);
                --n;
            }else {
                a = mult2x2Matrixes(a,a);
                n = n>>1;
            }
        }

        return res;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        BigInteger[][] a = {{BigInteger.ONE, BigInteger.ONE},
                            {BigInteger.ONE, BigInteger.ZERO}};

        BigInteger[][] powered = FastPow2x2Matrixes(a, n-1);

        System.out.println(powered[1][0].add(powered[1][1]));
    }
}
