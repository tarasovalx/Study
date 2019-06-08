import java.util.*;

enum Tag {
    IDENT,
    NUMBER,
    ADD,
    SUB,
    MULT,
    DIV,
    LPAREN,
    RPAREN
}

class Token {
    Tag tag;
    String value;

    Token(Tag t, String s) {
        tag = t;
        value = s;
    }
}

class Parser {
    private ArrayList<Token> tokens;
    private int pos = 0, res;
    private Scanner in;
    private boolean isComplete;
    private HashMap<String, Integer> m = new HashMap<>();

    public Parser(ArrayList<Token> tokens, Scanner in) {
        this.in = in;
        this.tokens = tokens;
        try {
            res = parse();
            isComplete = true;
        } catch (Exception e) {
            isComplete = false;
        }

    }

    private int getVar(Token t) {
        if (m.containsKey(t.value)) {
            return m.get(t.value);
        } else {
            int i = in.nextInt();
            m.put(t.value, i);
            return i;
        }
    }

    private int parse() throws Exception {
        res = parseE();
        if (pos == tokens.size()) return res;
        else throw new Exception();
    }

    private int parseE() throws Exception {
        res = parseT();
        parseE1();
        return res;
    }

    private int parseE1() throws Exception {
        if (pos >= tokens.size()) return res;
        if (tokens.get(pos).tag == Tag.ADD) {
            pos++;
            res += parseT();
            parseE1();

        } else if (tokens.get(pos).tag == Tag.SUB) {
            pos++;
            res -= parseT();
            parseE1();
        }
        return res;
    }

    private int parseT() throws Exception {
        res = parseF();
        parseT1();
        return res;
    }

    private int parseT1() throws Exception {
        if (pos >= tokens.size()) return res;
        if (tokens.get(pos).tag == Tag.MULT) {
            pos++;
            res *= parseF();
            parseT1();
        } else if (tokens.get(pos).tag == Tag.DIV) {
            pos++;
            res /= parseF();
            parseT1();
        }
        return res;
    }

    private int parseF() throws Exception {
        if (tokens.get(pos).tag == Tag.NUMBER) {
            pos++;
            return Integer.parseInt(tokens.get(pos - 1).value);
        } else if (tokens.get(pos).tag == Tag.IDENT) {
            pos++;
            return getVar(tokens.get(pos - 1));
        } else if (tokens.get(pos).tag == Tag.LPAREN) {
            pos++;
            int res = parseE();
            if (tokens.get(pos).tag != Tag.RPAREN)
                throw new Exception();
            pos++;
            return res;
        } else if (tokens.get(pos).tag == Tag.SUB) {
            pos++;
            return -parseF();
        } else {
            throw new Exception();
        }
    }

    public String toString() {
        if (isComplete) return Integer.toString(res);
        else return "error";
    }
}

public class Calc {

    private static ArrayList<Token> tokenize(String s) {
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
                }
            }
        }
        return tokens;
    }

    private static Scanner in = new Scanner(System.in);

    public static void main(String[] args) {
        ArrayList<Token> tokens = tokenize(in.nextLine());
        System.out.println(new Parser(tokens, in));
    }
}
