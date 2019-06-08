import java.util.Arrays;
import java.util.Scanner;
import java.util.stream.Collectors;

public class MaxNum {
    public static void main(String[] args) {
        int n;
        Scanner scanner = new Scanner(System.in);
        n = scanner.nextInt();
        int[] numbers = new int[n];
        String[] strs = new String[n];
        
        for(int i = 0; i < n;i++){
            numbers[i] = scanner.nextInt();
        }
        for(int i =0; i < n; i++){
            strs[i] = Integer.toString(numbers[i]);
        }

        String a,b;
        for(int i = 0; i < n-1;){
            b = strs[i+1] + strs[i];
            a = strs[i] + strs[i+1];
            if(b.compareTo(a) > 0){
                String tmp = strs[i+1];
                strs[i+1] = strs[i];
                strs[i] = tmp;
                i= 0;
            }else {
                i++;
            }
        }
        System.out.println(Arrays.stream(strs).collect(Collectors.joining("")));
    }
}
