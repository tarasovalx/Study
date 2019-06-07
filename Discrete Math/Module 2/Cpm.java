import java.util.*;
import java.util.stream.Collectors;

class Edge{
    private Node a;
    private Node b;
    private String key;

    public Edge(Node a, Node b) {
        this.a = a;
        this.b = b;
        key = a.toString() + b.toString();
    }

    @Override
    public boolean equals(Object obj) {
        Edge e = (Edge)obj;
        return e.key.equals(this.key);
    }

    @Override
    public int hashCode() {
        return this.key.hashCode();
    }

    public String toString() {
        String res = "";
        res += a.key + " -> " + b.key;
            if (a.color == 2 & b.color == 2) {
                res += " [ color = blue ]";
            } else if(a.redWith.contains(b) || b.redWith.contains(a)) {
                res += " [ color = red ]";
            }
        return res;
    }
}

class Node {
    public int weight;
    public String key;
    public int dist = Integer.MAX_VALUE;
    public int color = 0;
    public ArrayList<Node> redWith = new ArrayList<>();
    public ArrayList<Node> to = new ArrayList<>();

    public Node(int weight, String key) {
        this.weight = weight;
        this.key = key;
    }

    public String toString() {
        String res = "";
        res += key;
        res += "[ label = \"" + key + "(" + weight + ")" + "\"";
        if (color == 2 || color == 4) {
            res += ", color = ";
            if (color == 2) {
                res += "blue";
            } else if (color == 4) {
                res += "red";
            }
        }
        res += " ]";
        return res;
    }
}

public class Cpm {
    private static ArrayList<Node> graph = new ArrayList<>();
    private static Set<Edge> edges = new HashSet<>();
    private static Map<String, Node> k = new HashMap<>();
    private static int maxDist = Integer.MAX_VALUE;
    private static ArrayList<ArrayList<Node>> pathes = new ArrayList<>();

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String seq = "";
        while (in.hasNextLine()) {
            seq += in.nextLine();
        }
        seq = seq.replaceAll("\\s+", "");
        List<List<String>> sequences = Arrays.stream(seq.split(";"))
                .map(l -> Arrays.asList(l.replaceAll(" ", "").split("<")))
                .collect(Collectors.toList());
        for (List<String> group : sequences) {
            Node prev = null;
            for (String task : group) {
                int pos = task.indexOf('(');
                if (task.indexOf('(') >= 0) {
                    String name = task.substring(0, pos);
                    int w = Integer.parseInt(task.substring(pos + 1, task.length() - 1));
                    graph.add(new Node(w, name));
                    k.put(name, graph.get(graph.size() - 1));
                    if (prev != null) {
                        edges.add(new Edge(prev, graph.get(graph.size() - 1)));
                        prev.to.add(graph.get(graph.size() - 1));
                    }
                    prev = graph.get(graph.size() - 1);
                } else {
                    if (prev != null) {
                        edges.add(new Edge(prev, k.get(task)));
                        prev.to.add(k.get(task));
                    }
                    prev = k.get(task);
                }
            }
        }
        colorGraph(graph);
    }

    private static void colorGraph(ArrayList<Node> g) {
        for (Node v : graph) {
            if (v.color == 0) {
                DFS(v);
                v.dist = -v.weight;
                ArrayList<Node> l = new ArrayList<>();
                l.add(v);
                findPathes(v, v.dist, l);
            }
        }

        for(ArrayList<Node> path:pathes){
            for(int i = 1; i < path.size(); i++){
                path.get(i).redWith.add(path.get(i-1));
                path.get(i).color = 4;
            }
            if(!path.isEmpty() && path.get(0).color != 2){
                path.get(0).color = 4;
            }
        }

        System.out.println("digraph {");
        for (Node v : graph) {
            System.out.println(v.toString());
        }
        for (Edge e : edges) {
            System.out.println(e);
        }
        System.out.println("}");
    }

    private static void DFS(Node node) {
        if (node.color > 0) return;
        node.color = 1;
        for (Node u : node.to) {
            if (u.color == 0) DFS(u);
            else if (u.color == 1) {
                colorCycle(u);
                node.color = 2;
            }
        }
        if (node.color != 2) node.color = 3;
    }

    private static void findPathes(Node v, int m, ArrayList<Node> path) {
        if(m < maxDist){
            maxDist = m;
            pathes.clear();
            pathes.add(path);
        }else if(m == maxDist) {
            pathes.add(path);
        }
        if (v.color != 2){
            for(Node u: v.to){
                if (u.color != 2 && v.dist + u.weight * (-1) <= u.dist) {
                    u.dist = v.dist + u.weight * (-1);
                    ArrayList<Node> appendable = (ArrayList<Node>) path.clone();
                    if (m > u.dist) {
                        m = u.dist;
                    }
                    appendable.add(u);
                    findPathes(u, u.dist, appendable);
                }
            }
        }
    }

    private static void colorCycle(Node v) {
        v.color = 2;
        for (Node u : v.to) {
            if (u.color != 2) colorCycle(u);
        }
    }
}
