import java.util.HashSet;
import java.util.Scanner;

public class Econom {
    private static HashSet<String> uniqExprs = new HashSet<>();

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String polandNot = scanner.nextLine();
        getUniqExprs(polandNot);
        System.out.println(uniqExprs.size());

    }

    private static void getUniqExprs(String expr) {
        if (expr.length() >= 5) {
            uniqExprs.add(expr);
            if(expr.length() >5) {
                String next = expr.substring(2, expr.length() - 1);
                int start = next.indexOf('('), end = start + 1;
                int left = 1, right = 0;
                for (; left != right; end++) {
                    if (next.charAt(end) == '(') left++;
                    else if (next.charAt(end) == ')') right++;
                }
                getUniqExprs(next.substring(start, end));
                getUniqExprs(next.substring(end));
            }
        }
    }
}
