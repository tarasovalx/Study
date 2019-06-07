import java.util.*;

enum Tag {
    IDENT,
    NUMBER,
    ADD,
    SUB,
    MULT,
    DIV,
    LPAREN,
    RPAREN,
    COLON,
    LESS,
    MORE,
    LESSEQ,
    MOREEQ,
    NOTEQ,
    ASSIGN,
    INTERROGATION,
    EQUAL,
    COMMA,
    SEMICOMMA
}

class Token {
    Tag tag;
    String value;

    Token(Tag t, String s) {
        tag = t;
        value = s;
    }
}

class Call {
    String ident;
    int argsCount = 0;

    public Call(String i) {
        ident = i;
    }
}

class Parser {
    private ArrayList<Node> nodes = new ArrayList<>();
    private List<Token> tokens;
    private int pos = 0;
    private String ident;
    private HashSet<String> args = new HashSet<>();
    private ArrayList<Call> dependencies = new ArrayList<>();

    public Parser(List<Token> tokens) {
        this.tokens = tokens;
    }

    private Tag getCurrentTag() {
        return tokens.get(pos).tag;
    }

    public ArrayList<Node> parse() throws Exception {
        while (pos < tokens.size()) {
            parseProgramm();
        }
        return nodes;
    }

    private void parseProgramm() throws Exception {
        parseFunction();
        nodes.add(new Node(((ArrayList<Call>) this.dependencies.clone()), ident, args.size()));
        ident = "";
        args.clear();
        dependencies.clear();
    }

    private void parseFunction() throws Exception {
        if (getCurrentTag() != Tag.IDENT) throw new Exception();
        ident = tokens.get(pos).value;
        pos++;
        if (getCurrentTag() != Tag.LPAREN) throw new Exception();
        pos++;
        parseFormalArgsList();
        if (getCurrentTag() != Tag.RPAREN) throw new Exception();
        pos++;
        if (getCurrentTag() != Tag.ASSIGN) throw new Exception();
        pos++;
        parseE();
        if (getCurrentTag() != Tag.SEMICOMMA) throw new Exception();
        pos++;
    }

    private void parseE() throws Exception {
        parseCompE();
        if (pos >= tokens.size()) return;
        if (getCurrentTag() != Tag.INTERROGATION) return;
        pos++;
        parseCompE();
        if (getCurrentTag() != Tag.COLON) throw new Exception();
        pos++;
        parseE();
    }

    private boolean isCompOp(Tag x) {
        return x == Tag.EQUAL ||
                x == Tag.NOTEQ ||
                x == Tag.LESS ||
                x == Tag.MORE ||
                x == Tag.LESSEQ ||
                x == Tag.MOREEQ;
    }

    private void parseCompE() throws Exception {
        parseE1();
        if (pos < tokens.size() && isCompOp(getCurrentTag())) {
            pos++;
            parseE1();
        }
    }

    private void parseE1() throws Exception {
        parseT();
        parseE2();
    }

    private void parseE2() throws Exception {
        if (getCurrentTag() == Tag.ADD) {
            pos++;
            parseT();
            parseE2();

        } else if (getCurrentTag() == Tag.SUB) {
            pos++;
            parseT();
            parseE2();
        }
    }

    private void parseT() throws Exception {
        parseF();
        parseT1();
    }

    private void parseT1() throws Exception {
        //if (pos >= tokens.size()) return;
        if (tokens.get(pos).tag == Tag.MULT) {
            pos++;
            parseF();
            parseT1();
        } else if (tokens.get(pos).tag == Tag.DIV) {
            pos++;
            parseF();
            parseT1();
        }
    }

    private void parseF() throws Exception {
        if (getCurrentTag() == Tag.NUMBER) {
            pos++;
        } else if (getCurrentTag() == Tag.IDENT) {
            if (pos + 1 < tokens.size() && tokens.get(pos + 1).tag != Tag.LPAREN) {
                if (!args.contains(tokens.get(pos).value)) throw new Exception();
            }
            pos++;
            if (pos < tokens.size() && getCurrentTag() == Tag.LPAREN) {
                String currentCallIdent = tokens.get(pos - 1).value;
                dependencies.add(new Call(currentCallIdent));
                pos++;
                parseActualArgsList(dependencies.get(dependencies.size() - 1));
                if (getCurrentTag() != Tag.RPAREN) throw new Exception();
                pos++;
            }
        } else if (tokens.get(pos).tag == Tag.LPAREN) {
            pos++;
            parseE();
            if (tokens.get(pos).tag != Tag.RPAREN) throw new Exception();
            pos++;
        } else if (tokens.get(pos).tag == Tag.SUB) {
            pos++;
            parseF();
        } else throw new Exception();
    }

    private void parseFormalArgsList() throws Exception {
        if (getCurrentTag() == Tag.RPAREN) return;
        parseIdentList();
    }

    private void parseIdentList() throws Exception {
        if (getCurrentTag() == Tag.IDENT) {
            args.add(tokens.get(pos).value);
            pos++;
        } else throw new Exception();
        if (getCurrentTag() == Tag.COMMA) {
            pos++;
            parseIdentList();
        }
    }


    private void parseActualArgsList(Call c) throws Exception {
        if (getCurrentTag() == Tag.RPAREN) return;
        parseExprList(c);
    }

    private void parseExprList(Call c) throws Exception {
        parseE();
        c.argsCount++;
        if (getCurrentTag() == Tag.COMMA) {
            pos++;
            parseExprList(c);
        } else if (getCurrentTag() != Tag.RPAREN) throw new Exception();
    }
}

class Node {
    public ArrayList<Call> dependecies;
    public String ident;
    public int low;
    public int T1 = 0;
    public ArrayList<Node> from = new ArrayList<>();
    int argsCount;
    int color = 0;

    public String toString() {
        return ident;
    }

    public Node(ArrayList<Call> dependecies, String ident, int cnt) {
        this.dependecies = dependecies;
        this.ident = ident;
        this.argsCount = cnt;
    }
}

public class Modules {
    private static int componentsCnt = 1;
    private static int time = 1;

    private static void tarjan() {
        Stack<Node> stack = new Stack<>();
        for (Node v : graph)
            if (v.T1 == 0) visitVertexTarjan(v, stack);
    }

    private static void visitVertexTarjan(Node v, Stack<Node> stack) {
        v.T1 = v.low = time++;
        stack.push(v);
        for(Node u : v.from) {
            if (u.T1 == 0) visitVertexTarjan(u, stack);
            if (u.color == 0 && v.low > u.low) v.low = u.low;
        }
        if (v.T1 == v.low) {
            Node u;
            do {
                u = stack.pop();
                u.color = componentsCnt;
            } while (u != v);
            componentsCnt++;
        }
    }

    private static ArrayList<Node> graph = new ArrayList<>();

    private static HashMap<String, Node> map = new HashMap<>();

    private static ArrayList<Token> tokenize(String s) throws Exception {
        ArrayList<Token> tokens = new ArrayList<>();
        for (int pos = 0; pos < s.length(); ++pos) {
            char currentChar = s.charAt(pos);
            String currentToken = "";
            if (Character.isAlphabetic(currentChar)) {
                for (; pos < s.length() && (Character.isDigit(s.charAt(pos)) || Character.isAlphabetic(s.charAt(pos))); ++pos) {
                    currentToken += s.charAt(pos);
                }
                pos--;
                tokens.add(new Token(Tag.IDENT, currentToken));
            } else if (Character.isDigit(currentChar)) {
                for (; pos < s.length() && (Character.isDigit(s.charAt(pos))); ++pos) {
                    currentToken += s.charAt(pos);
                }
                pos--;
                tokens.add(new Token(Tag.NUMBER, currentToken));
            } else {
                switch (currentChar) {
                    case '+':
                        tokens.add(new Token(Tag.ADD, "+"));
                        break;
                    case '-':
                        tokens.add(new Token(Tag.SUB, "-"));
                        break;
                    case '*':
                        tokens.add(new Token(Tag.MULT, "*"));
                        break;
                    case '/':
                        tokens.add(new Token(Tag.DIV, "/"));
                        break;
                    case '(':
                        tokens.add(new Token(Tag.LPAREN, "("));
                        break;
                    case ')':
                        tokens.add(new Token(Tag.RPAREN, ")"));
                        break;
                    case '=':
                        tokens.add(new Token(Tag.EQUAL, "="));
                        break;
                    case ',':
                        tokens.add(new Token(Tag.COMMA, ","));
                        break;
                    case ';':
                        tokens.add(new Token(Tag.SEMICOMMA, ";"));
                        break;
                    case ':':
                        if (s.charAt(pos + 1) == '=') {
                            tokens.add(new Token(Tag.ASSIGN, ":="));
                            pos++;
                        } else tokens.add(new Token(Tag.COLON, ":"));
                        break;
                    case '?':
                        tokens.add(new Token(Tag.INTERROGATION, "?"));
                        break;
                    case '<': {
                        if (s.charAt(pos + 1) == '=') {
                            tokens.add(new Token(Tag.LESSEQ, "<="));
                            pos++;
                        } else if (s.charAt(pos + 1) == '>') {
                            tokens.add(new Token(Tag.NOTEQ, "<>"));
                            pos++;
                        } else tokens.add(new Token(Tag.LESS, "<"));
                        break;
                    }
                    case '>': {
                        if (s.charAt(pos + 1) == '=') {
                            tokens.add(new Token(Tag.MOREEQ, ">="));
                            pos++;
                        } else tokens.add(new Token(Tag.MORE, ">"));
                        break;
                    }
                    case ' ':
                        break;
                    case '\t':
                        break;
                    case '\n':
                        break;
                    default:
                        throw new Exception();
                }
            }
        }
        return tokens;
    }

    private static Scanner in = new Scanner(System.in);

    public static void main(String[] args) {
        String s = "";
        while (in.hasNextLine()) { s += in.nextLine(); }

        try {
            Parser parser = new Parser(tokenize(s));
            graph = parser.parse();
            for (Node v : graph) {
                map.put(v.ident, v);
            }
        } catch (Exception e) {
            System.out.println("error");
            System.exit(0);
        }

        for (Node v : graph) {
            for (Call call : v.dependecies) {
                Node adding = map.get(call.ident);
                if (adding != null && adding.argsCount == call.argsCount) {
                    v.from.add(adding);
                } else {
                    System.out.println("error");
                    System.exit(0);
                }
            }
        }
        tarjan();
        System.out.println(componentsCnt - 1);
    }
}
