import java.io.*;
import java.util.*;
import java.util.stream.Collectors;

public class Sync {
    private static List<File> SList = new ArrayList<>();
    private static List<File> DList = new ArrayList<>();

    private static int sPathLenght;
    private static int dPathLenght;

    public static void getAllFiles(File start, List<File> dest) {
        if (!start.isDirectory()) return;
        List<File> curList = Arrays.stream(start.listFiles()).collect(Collectors.toList());
        dest.addAll(curList.stream().filter(File::isFile)
                .collect(Collectors.toList()));
        for (File file : curList) {
            if (file.isDirectory()) {
                getAllFiles(file, dest);
            }
        }
    }

    public static boolean isSameData(File a, File b) {
        try{
            BufferedInputStream aStream = new BufferedInputStream(new FileInputStream(a));
            BufferedInputStream bStream = new BufferedInputStream(new FileInputStream(b));
            if (aStream.available() == bStream.available()){
                while (true){
                    int s1 = aStream.read();
                    int s2 = bStream.read();
                    if (s1==-1){
                        aStream.close();
                        bStream.close();
                        return true;
                    }
                    if(s1!= s2){
                        aStream.close();
                        bStream.close();
                        return false;
                    }
                }
            }else {
                aStream.close();
                bStream.close();
                return false;
            }

        }catch (IOException ex){

        }
        return false;
    }

    public static boolean isSamePath(File a, File b) {
        return a.getPath().substring(sPathLenght + 1).compareTo(b.getPath().substring(dPathLenght + 1)) == 0;
    }

    public static void main(String[] args) {
        File S = new File(args[0]);
        File D = new File(args[1]);

        getAllFiles(S, SList);
        getAllFiles(D, DList);

        sPathLenght = S.getName().length();
        dPathLenght = D.getName().length();



        for (int i = 0; i < SList.size(); i++) {
            File val = SList.get(i);

            for (int j = 0; j < DList.size(); j++) {
                File val1 = DList.get(j);
                if (isSamePath(val, val1) && isSameData(val, val1)) {
                    SList.remove(val);
                    DList.remove(val1);
                    i--;
                    break;
                }
            }
        }

        SList.sort(new Comparator<File>() {
            @Override
            public int compare(File o1, File o2) {
                return o1.getPath().compareTo(o2.getPath());
            }
        });
        DList.sort(new Comparator<File>() {
            @Override
            public int compare(File o1, File o2) {
                return o1.getPath().compareTo(o2.getPath());
            }
        });

        if ((DList.isEmpty() && SList.isEmpty())) {
            System.out.println("IDENTICAL");
        } else {
            DList.forEach(x -> System.out.println("DELETE " + x.getPath().substring(dPathLenght + 1)));
            SList.forEach(x -> System.out.println("COPY " + x.getPath().substring(sPathLenght + 1)));
        }
    }
}
