import static java.lang.Math.min;
import static java.lang.System.out;

void main() throws IOException {
    var poly = new BufferedReader(new InputStreamReader(System.in)).readLine();
    var shortest = poly.length();
    var chars = "abcdefghijklmnopqrstuvwxyz";
    for (int i = 0; i < chars.length(); i++) {
        var p = poly.replaceAll(chars.substring(i, i + 1) + "|" + chars.substring(i, i + 1).toUpperCase(), "");
        shortest = min(shortest, react(p).length());
    }
    out.println(react(poly).length() + " " + shortest);
}

String react(String poly) {
    var p = poly;
    while (true) {
        var red = p;
        for (int i = 0; i < red.length() - 1; i++) {
            if (red.charAt(i) != red.charAt(i + 1) && red.substring(i, i + 1).equalsIgnoreCase(red.substring(i + 1, i + 2))) {
                red = red.substring(0, i) + red.substring(i + 2);
                i--;
            }
        }
        if (red.length() == p.length()) break;
        p = red;
    }
    return p;
}