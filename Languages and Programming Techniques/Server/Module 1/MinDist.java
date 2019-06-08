import java.util.Scanner;

public class MinDist {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String s = in.nextLine();
        char x = in.next().charAt(0), y = in.next().charAt(0);
        int xPos = -1; int yPos = -1;
        int minDist = s.length();
        for(int i = 0; i < s.length(); i++){
            if (s.charAt(i) == x) xPos = i;
            if (s.charAt(i) == y) yPos = i;
            if (xPos != -1 && yPos != -1){
                minDist = Math.min(minDist, Math.abs(xPos - yPos));
            }
        }
        System.out.println(--minDist);
    }
}