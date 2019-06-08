import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

class Fraction {
    public int numerator = 1;
    public int denumerator = 1;
    
    public static int lcm(int a, int b) {
        return a / gcd(a, b) * b;
    }

    public static int gcd(int a, int b) {
        a = Math.abs(a);
        b = Math.abs(b);
        if (a < b) {
            int temp = a;
            a = b;
            b = temp;
        }
        if (b == 0)
            return a;
        else
            return gcd(b, a % b);
    }

    public double getValue() {
        return ((double) numerator) / ((double) denumerator);
    }

    public Fraction add(Fraction b) {
        int lcm = lcm(denumerator, b.denumerator);
        return new Fraction(numerator * (lcm / denumerator) + b.numerator * (lcm / b.denumerator), lcm);
    }

    public Fraction mult(int n) {
        return new Fraction(numerator * n, denumerator);
    }

    public Fraction devide(int n) {
        return new Fraction(numerator, denumerator * n);
    }

    public Fraction mult(Fraction a) {
        return new Fraction(numerator * a.numerator, denumerator * a.denumerator);
    }

    public Fraction(int a) {
        numerator = a;
    }

    public Fraction Normalize() {
        int gcd = gcd(numerator, denumerator);
        if (numerator < 0 && denumerator < 0) {
            numerator = -numerator;
            denumerator = -denumerator;
        }
        if (denumerator < 0) {
            denumerator = -denumerator;
            numerator = -numerator;
        }
        return new Fraction(numerator / gcd, denumerator / gcd);
    }

    public Fraction(int num, int denum) {
        numerator = num;
        denumerator = denum;
    }

    public String toString() {
        return numerator + "/" + denumerator;
    }
}

public class Gauss {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int n = scanner.nextInt();
        ArrayList<List<Fraction>> equations = new ArrayList<>(n);
        for (int i = 0; i < n; i++) {
            equations.add(new ArrayList<>());
            for (int j = 0; j <= n; j++) {
                equations.get(i).add(new Fraction(scanner.nextInt()));
            }
        }

        //start Gauss Method

        for (int i = 0; i < n; i++) {
            int maxj = i;
            double max = equations.get(i).get(i).getValue();
            for (int j = i; j < equations.size(); j++) {
                if (equations.get(j).get(i).numerator != 0) {
                    if (Math.abs(equations.get(j).get(i).getValue()) > max) {
                        max = Math.abs(equations.get(j).get(i).getValue());
                        maxj = j;
                    }
                    break;
                }
            }
            List<Fraction> tmp = equations.get(i);
            equations.set(i, equations.get(maxj));
            equations.set(maxj, tmp);
            if (equations.get(i).get(i).numerator == 0) {
                System.out.println("No solution");
                System.exit(0);
            }
            equations.set(i, to1at(equations.get(i), i));
            for (int j = i + 1; j < n; j++) {
                Fraction k = equations.get(j).get(i);
                Fraction k1 = new Fraction(-k.numerator, k.denumerator);
                equations.set(j, sumRows(multRow(equations.get(i), k1), equations.get(j)));
            }
        }

        for (int i = n - 1; i > 0; i--) {
            for (int j = i - 1; j >= 0; j--) {
                Fraction k = equations.get(j).get(i);
                Fraction k1 = new Fraction(-k.numerator, k.denumerator);
                equations.set(j, sumRows(multRow(equations.get(i), k1), equations.get(j)));
            }
        }
        for (int i = 0; i < n; i++) {
            System.out.println(equations.get(i).get(n));
        }
    }

    public static List<Fraction> multRow(List<Fraction> row, Fraction n) {
        return IntStream.range(0, row.size())
                .mapToObj(x -> row.get(x).mult(n).Normalize())
                .collect(Collectors.toList());
    }

    public static List<Fraction> sumRows(List<Fraction> rowA, List<Fraction> rowB) {
        return IntStream.range(0, rowA.size())
                .mapToObj(x -> rowA.get(x).add(rowB.get(x)).Normalize())
                .collect(Collectors.toList());
    }

    public static List<Fraction> to1at(List<Fraction> row, int index) {
        int n = row.get(index).numerator;
        int d = row.get(index).denumerator;

        if (n == 0) {
            return IntStream.range(0, row.size())
                    .mapToObj(x -> row.get(x).Normalize())
                    .collect(Collectors.toList());
        }
        if (d == 1) {
            return IntStream.range(0, row.size())
                    .mapToObj(x -> row.get(x).devide(n).Normalize())
                    .collect(Collectors.toList());
        } else {
            return IntStream.range(0, row.size())
                    .mapToObj(x -> row.get(x).mult(d).devide(n).Normalize())
                    .collect(Collectors.toList());
        }
    }
}
