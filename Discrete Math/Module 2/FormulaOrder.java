import java.util.*;
import java.util.stream.Collectors;

enum Tag {
    START,
    IDENT,
    CONST,
    ADD,
    SUB,
    MULT,
    DIV,
    LPAREN,
    RPAREN,
    COMMA,
    EQUAL

}

class Node {
    class Token {
        String value;
        Tag t;

        public Token(String v, Tag tt) {
            value = v;
            t = tt;
        }
    }

    ArrayList<Token> getTokens(String s) throws Exception {
        ArrayList<Token> tokens = new ArrayList<>();
        Tag last = Tag.START;
        for (int i = 0; i < s.length(); i++) {
            switch (s.charAt(i)) {
                case ' ':
                    break;
                case '+':
                    last = Tag.ADD;
                    tokens.add(new Token("+", Tag.ADD));
                    break;
                case '-':
                    if (last == Tag.IDENT || last == Tag.CONST || last == Tag.RPAREN) {
                        tokens.add(new Token("-", Tag.SUB));
                        last = Tag.SUB;
                    }
                    break;
                case '*':
                    last = Tag.MULT;
                    tokens.add(new Token("*", Tag.MULT));
                    break;
                case '/':
                    last = Tag.DIV;
                    tokens.add(new Token("/", Tag.DIV));
                    break;
                case '(':
                    last = Tag.LPAREN;
                    tokens.add(new Token("(", Tag.LPAREN));
                    break;
                case ')':
                    last = Tag.RPAREN;
                    tokens.add(new Token(")", Tag.RPAREN));
                    break;
                case ',':
                    last = Tag.COMMA;
                    tokens.add(new Token(",", Tag.COMMA));
                    break;
                case '=':
                    last = Tag.EQUAL;
                    tokens.add(new Token("=", Tag.EQUAL));
                    break;
                default:
                    int j = i;
                    if (Character.isDigit((s.charAt(i)))) {
                        String var = "";
                        while (i < s.length() && Character.isDigit(s.charAt(i))) {
                            var += s.charAt(i) + "";
                            i++;
                        }
                        i--;
                        last = Tag.CONST;
                        tokens.add(new Token(var, Tag.CONST));
                    } else if (Character.isAlphabetic(s.charAt(i))) {
                        String var = s.charAt(i) + "";
                        i++;
                        while (i < s.length() && (Character.isAlphabetic(s.charAt(i)) || Character.isDigit(s.charAt(i)))) {
                            var += s.charAt(i);
                            i++;
                        }
                        i--;
                        last = Tag.IDENT;
                        tokens.add(new Token(var, Tag.IDENT));
                    }
                    else throw new Exception();
            }
        }
        return tokens;
    }

    void parse() throws Exception {
        List<Token> left = new ArrayList<>();
        List<List<Token>> right = new ArrayList<>();
        int lcnt = 0;
        int rcnt = 0;
        right.add(new ArrayList<>());

        boolean check = false;
        for (Token t : tokens) {
            if(!check && t.t == Tag.COMMA){
                lcnt++;
            }

            if(check && t.t == Tag.COMMA){
                rcnt++;
                right.add(new ArrayList<>());
            }
            if (t.t == Tag.EQUAL) {
                check = true;
                continue;
            }
            if (check && t.t != Tag.COMMA) {
                right.get(rcnt).add(t);
            }
            if(!check) {
                left.add(t);
            }
        }
        if (left.stream().anyMatch(x->x.t != Tag.IDENT && x.t != Tag.COMMA)){
            throw new Exception();
        }

        if (right.size() == 0 || lcnt != rcnt || left.stream().filter(x->x.t == Tag.IDENT).count() != rcnt+1) {
            throw new Exception();
        }
        parseLeft(left);
        for(List<Token> expr : right){
            parseExpr(expr);
        }
    }

    void parseLeft(List<Token> tokens) throws Exception {
        idents = tokens.stream().filter(x -> x.t == Tag.IDENT).map(x -> x.value).collect(Collectors.toList());
        if (idents.size() == 0) {
            throw new Exception();
        }
    }

    void parseExpr(List<Token> tokens) throws Exception {
        if (tokens.isEmpty()){
            throw new Exception();
        }
        ArrayDeque<Token> q = new ArrayDeque<>(tokens);
        ArrayDeque<Token> vars = new ArrayDeque<>();
        Stack<Token> ops = new Stack<>();
        int opCnt = 0, varCnt = 0;
        while (!q.isEmpty()) {
            Token token = q.poll();
            if (token.t == Tag.IDENT || token.t == Tag.CONST) {
                if (token.t == Tag.IDENT) {
                    dependencies.add(token.value);
                }
                vars.push(token);
                varCnt++;
            } else if (token.t == Tag.LPAREN) {
                ops.push(token);
            } else if (token.t == Tag.RPAREN) {
                while (!ops.isEmpty() && ops.lastElement().t != Tag.LPAREN) {
                    vars.push(ops.pop());
                }
                if (!ops.isEmpty() && ops.lastElement().t == Tag.LPAREN) {
                    ops.pop();
                    if (!ops.isEmpty() && (ops.lastElement().t == Tag.DIV || ops.lastElement().t == Tag.MULT || ops.lastElement().t == Tag.ADD || ops.lastElement().t == Tag.SUB)) {
                        vars.push(ops.pop());
                    }
                } else if (!ops.isEmpty()) {
                    throw new Exception();
                }
            } else {
                ops.push(token);
                opCnt++;
            }
        }
        while (!ops.isEmpty()) {
            Token token = ops.pop();
            if (token.t == Tag.LPAREN) {
                throw new Exception();
            } else {
                vars.push(token);
            }
        }
        if (varCnt - opCnt != 1) {
            throw new Exception();
        }
    }

    int color = 0;
    String formula;
    ArrayList<Token> tokens;
    List<String> idents = new ArrayList<>();
    List<String> dependencies = new ArrayList<>();
    HashSet<Node> to = new HashSet<>();
    HashSet<Node> from = new HashSet<>();

    public Node(String s) {
        formula = s;
        try {
            tokens = getTokens(s);
            parse();
        } catch (Exception e) {
            System.out.println("syntax error");
            System.exit(1);
        }
    }

    public String toString() {
        return formula;
    }
}

public class FormulaOrder {

    public static void dfs(Queue<Node> queue, Node v) throws Exception {
        v.color = 1;
        for (Node u : v.to) {
            if (u.color == 0)
                dfs(queue, u);
            else if (u.color == 1)
                throw new Exception();
        }
        v.color = 2;
        queue.add(v);
    }

    public static void main(String[] args) {
        HashMap<String, Node> m = new HashMap<>();
        Scanner in = new Scanner(System.in);

        ArrayList<Node> graph = new ArrayList<>();

        while (in.hasNextLine()) {
            Node adding = new Node(in.nextLine());
            for (String key : adding.idents) {
                if (m.containsKey(key)) {
                    System.out.println("syntax error");
                    System.exit(1);
                }
                m.put(key, adding);
            }
            graph.add(adding);
        }

        for (Node v : graph) {
            for (String key : v.dependencies) {
                v.from.add(m.get(key));
            }
        }
        for (Node v : graph) {
            for (Node u : v.from) {
                if (u == null){
                    System.out.println("syntax error");
                    System.exit(1);
                }
                u.to.add(v);
            }
        }

        ArrayDeque<Node> q = new ArrayDeque<>();
        try {
            for (Node v : graph) {
                if (v.color == 0) {
                    dfs(q, v);
                }
            }

        } catch (Exception e) {
            System.out.println("cycle");
            System.exit(1);
        }
        while (!q.isEmpty()) {
            Node e =  q.pollLast();
            System.out.println(e);
        }
    }
}